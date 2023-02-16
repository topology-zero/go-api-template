package test

import (
	"strconv"

	"go-api-template/svc"
	"go-api-template/types/test"
)

// Test 测试测试
func Test(req *test.TestRequest, ctx *svc.ServiceContext) (resp test.TestResponse, err error) {
	z, _ := strconv.Atoi(req.Z)

	resp.X = strconv.Itoa(req.X)
	resp.Y = strconv.Itoa(req.Y)
	resp.Z = z

	ctx.Log.Info("测试 traceId ")

	return
}
