// Code generated by goctl. DO NOT EDIT.
package types

type TestRequest struct {
	X int    `json:"x" label:"测试X"` // 测试X
	Y int    `json:"y" label:"测试Y"` // 测试Y
	Z string `json:"z" label:"测试Z"` // 测试Z
}

type TestResponse struct {
	X string `json:"x"` // 测试X
	Y string `json:"y"` // 测试Y
	Z int    `json:"z"` // 测试Z
}
