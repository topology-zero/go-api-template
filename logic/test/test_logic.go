package test

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"go-api-template/query"
	"go-api-template/types/test"
)

// Test 测试测试
func Test(req *test.TestRequest) (resp test.TestResponse, err error) {

	roleModel := query.RoleModel
	allRole, _ := roleModel.Find()
	for _, v := range allRole {
		logrus.Infof("%+v", v)
	}

	z, _ := strconv.Atoi(req.Z)

	resp.X = strconv.Itoa(req.X)
	resp.Y = strconv.Itoa(req.Y)
	resp.Z = z
	return
}
