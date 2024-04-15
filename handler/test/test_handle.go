package test

import (
	"go-api-template/internal/response"
	"go-api-template/logic/test"
	"go-api-template/svc"
	"go-api-template/types"

	"github.com/gin-gonic/gin"
)

// TestHandle 测试测试
func TestHandle(c *gin.Context) {
	var req types.TestRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := test.Test(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
