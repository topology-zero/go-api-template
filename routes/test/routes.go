// Code generated by goctl. DO NOT EDIT.
package test

import (
	"go-api-template/handler/test"

	"github.com/gin-gonic/gin"
)

func RegisterTestRoute(e *gin.Engine) {
	g := e.Group("/t")
	g.POST("/test", test.TestHandle)
}
