package test

import (
	"go-api-template/internal/response"
	"go-api-template/logic/test"
	testType "go-api-template/types/test"

	"github.com/gin-gonic/gin"
)

// TestHandle 测试测试
func TestHandle(c *gin.Context) {
	var req testType.TestRequest
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := test.Test(&req)
	response.HandleResponse(c, resp, err)
}
