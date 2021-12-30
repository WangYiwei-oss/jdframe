package casbin

import (
	"github.com/WangYiwei-oss/jdframe/src/configs"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

var E *casbin.Enforcer

func init() {
	adap, err := gormadapter.NewAdapterByDB(configs.NewGormAdapter().DB)
	if err != nil {
		log.Fatalln("casbin: 读取Gorm Adapter失败")
	}
	e, err := casbin.NewEnforcer("model.conf", adap)
	if err != nil {
		log.Fatalln("casbin: 读取配置文件失败")
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatalln("casbin: 读取policy失败")
	}
	E = e
	initPolicy()
}

func initPolicy() {
	m := make([]*RoleRel, 0)
	GetRoles(0, &m, "") //获取角色 对应
	for _, r := range m {
		_, err := E.AddRoleForUser(r.PRole, r.Role)
		if err != nil {
			log.Fatalln("casbin: 添加角色规则失败", err)
		}
	}
	userRoles := GetUserRoles()
	for _, user_role := range userRoles {
		_, err := E.AddRoleForUser(user_role.UserName, user_role.RoleName)
		if err != nil {
			log.Fatalln("casbin: 添加用户角色规则失败", err)
		}
	}
	uriRoles := GetRouterRoles()
	for _, uriRole := range uriRoles {
		_, err := E.AddPolicy(uriRole.RoleName, uriRole.RouterUri, uriRole.RouterMethod)
		if err != nil {
			log.Fatalln("casbin: 添加uri角色规则失败", err)
		}
	}
}
