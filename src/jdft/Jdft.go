package jdft

import (
	mcasbin "github.com/WangYiwei-oss/jdframe/src/casbin"
	"github.com/WangYiwei-oss/jdframe/src/configinjector"
	"github.com/WangYiwei-oss/jdframe/src/configparser"
	"github.com/WangYiwei-oss/jdframe/src/configs"
	"github.com/WangYiwei-oss/jdframe/src/models"
	"github.com/WangYiwei-oss/jdframe/wscore"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
)

//暴露一些类型给用户
type User models.User
type UserRole models.UserRole

var Jedis *configs.Jedis
var QrCode *configs.QrCode
var Gorm *configs.GormAdapter
var WebSocketFactory *wscore.ClientMap
var CasbinEnforcer *casbin.Enforcer
var GlobalLogger *configs.Mlogger

//单例
var onceCron sync.Once
var taskCron *cron.Cron

func init() {
	GlobalLogger = configs.GetLogger("Global")
	CasbinEnforcer = mcasbin.E
	WebSocketFactory = wscore.WebSocketFactory
	Gorm = configs.NewGormAdapter()
	Jedis = configs.NewJedis()
	QrCode = configs.NewQrCode()
}

func GetGlobalSettings() map[string]interface{} {
	return configparser.GlobalSettings
}

func GetLogger(name string) *configs.Mlogger {
	return configs.GetLogger(name)
}

func getCronTask() *cron.Cron { //创建定时任务的单例模式
	onceCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}

type Jdft struct {
	*gin.Engine
	Router      *gin.RouterGroup
	beanfactory *BeanFactory
}

func NewJdft() *Jdft {
	getCronTask().Start()
	return &Jdft{Engine: gin.New(), beanfactory: NewBeanFactory()}
}

func (j *Jdft) DefaultBean() *Jdft {
	j.Beans(configinjector.GetBeans()...) //注入默认的依赖
	return j
}

func (j *Jdft) Handle(httpMethod, relativePath string, handler interface{}) *Jdft {
	if handlerfunc := ConvertResponder(handler); handlerfunc != nil {
		router := j.Router
		router.Handle(httpMethod, relativePath, handlerfunc)
	} else {
		log.Printf("路由%s, %s没有被解析并注入", httpMethod, relativePath)
	}
	return j
}

func (j *Jdft) Attach(f ...gin.HandlerFunc) *Jdft {
	j.Use(f...)
	return j
}

func (j *Jdft) Launch() {
	log.Println("[INFO] 迁移角色表")
	//角色数据表迁移
	//err := Gorm.AutoMigrate(&models.User{})
	//if err != nil {
	//	log.Fatalln("用户表迁移错误:", err)
	//}
	err := Gorm.AutoMigrate(&models.Role{})
	if err != nil {
		log.Fatalln("权限表迁移错误:", err)
	}
	err = Gorm.AutoMigrate(&models.Router{})
	if err != nil {
		log.Fatalln("router表迁移错误:", err)
	}
	err = Gorm.AutoMigrate(&models.UserRole{})
	if err != nil {
		log.Fatalln("user_role表迁移错误:", err)
	}
	ip, _ := configparser.GlobalSettings["IP"].(string)
	err = j.Run(ip)
	if err != nil {
		log.Fatalf("服务器启动失败, %s", err)
	}
}

func (j *Jdft) Mount(version string, controllers ...JdController) *Jdft {
	j.Router = j.Group(version)
	for _, controller := range controllers {
		controller.Build(j)
		j.beanfactory.inject(controller)
	}
	return j
}

// Beans 进行依赖注入
func (j *Jdft) Beans(beans ...interface{}) *Jdft {
	if beans == nil || len(beans) == 0 {
		return j
	}
	for _, bean := range beans {
		j.beanfactory.addBean(bean)
	}
	return j
}

//增加定时任务
func (j *Jdft) CronTask(expr string, f func()) *Jdft {
	_, err := getCronTask().AddFunc(expr, f)
	if err != nil {
		log.Fatalf("启动定时任务失败, %s", err)
	}
	return j
}
