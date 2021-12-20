package middlewares

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/jdft"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"
)

func RBAC() gin.HandlerFunc {
	adap, err := gormadapter.NewAdapterByDB(jdft.NewGormAdapter().DB)
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
	return func(context *gin.Context) {
		user := context.Request.Header.Get("user")
		access, err := e.Enforce(user, context.Request.RequestURI, context.Request.Method)
		if !access || err != nil {
			context.AbortWithStatusJSON(403, gin.H{
				"success": false,
				"status":  "未通过权限验证",
				"data":    "",
			})
		} else {
			context.Set("user", user)
			context.Set("data", context.Request.Body)
			fmt.Println("通过验证")
			context.Next()
		}
	}
}
