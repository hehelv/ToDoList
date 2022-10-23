package service

import (
	"MyToDoList/model"
	"MyToDoList/pkg/e"
	"MyToDoList/pkg/util"
	"MyToDoList/serializer"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (srv *UserService) Register() *serializer.Response {
	code := e.SUCCESS
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", srv.UserName).First(&user).Count(&count)
	if count == 1 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.UserName = srv.UserName
	// 加密密码
	if err := user.SetPassword(srv.Password); err != nil {
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		util.LogrusObj.Error(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (srv *UserService) Login() *serializer.Response {
	var user model.User
	code := e.SUCCESS
	if err := model.DB.Where("user_name=?", srv.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			util.LogrusObj.Info(err)
			code = e.ErrorNotExistUser
			return &serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if !user.CheckPassword(srv.Password) {
		code = e.ErrorNotCompare
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	token, err := util.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorAuthToken
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}
