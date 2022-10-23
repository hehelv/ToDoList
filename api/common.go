package api

import (
	"MyToDoList/conf"
	"MyToDoList/serializer"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) *serializer.Response {
	// 验证不通过， 用翻译器翻译一下返回
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return &serializer.Response{
				Status: 40001,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	// json解析失败， 返回
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return &serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}
	// 参数错误
	return &serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
