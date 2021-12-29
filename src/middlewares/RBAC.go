package middlewares

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/casbin"
	"github.com/gin-gonic/gin"
)

func RBAC() gin.HandlerFunc {
	return func(context *gin.Context) {
		user := context.Request.Header.Get("user")
		access, err := casbin.E.Enforce(user, context.Request.RequestURI, context.Request.Method)
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
