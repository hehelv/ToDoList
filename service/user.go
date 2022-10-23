package service

import "MyToDoList/serializer"

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (srv *UserService) Register() *serializer.Response {

}

func (srv *UserService) Login() *serializer.Response {

}
