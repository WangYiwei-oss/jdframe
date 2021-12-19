package middlewares

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func RBAC() gin.HandlerFunc {
	return func(context *gin.Context) {
		e := casbin.NewEnforcer("src/casbin/model.conf", "src/casbin/p.csv")
		user := context.Request.Header.Get("user")
		if !e.Enforce(user, context.Request.RequestURI, context.Request.Method) {
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
