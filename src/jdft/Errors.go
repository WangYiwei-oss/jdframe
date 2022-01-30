package jdft

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				context.AbortWithStatusJSON(400, gin.H{
					"error": e,
				})
			}
		}()
		context.Next()
	}
}

func Throw(err string, code int, context *gin.Context) {
	context.Set("JDFT STATUS", code)
	panic(err)
}

func Error(err error, msg ...string) {
	if err == nil {
		return
	} else {
		errMsg := err.Error()
		panic(errMsg + "," + msg[0])
	}
}
