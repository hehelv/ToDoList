# 备忘录
> 基于Gin + Gorm ，符合Restful API设计规范，实现一个备忘录


## 主要功能
- 用户登录注册（jwt-go鉴权）
- 备忘录CRUD
- 存储备忘录浏览次数
- 分页

## 主要依赖
Golang V1.17
- Gin
- Gorm
- mysql
- redis
- ini
- jwt-go
- logrus
- go-swagger

## 项目结构
```
MyToDoList/
|—— api
|—— cache
|—— conf
|—— middleware
|—— model
|—— pkg
|   |—— e
|   |—— util
|—— routes
|—— serializer
|—— service
```
- api: 用于存放接口函数


## 简要说明
- mysql: 存储主要数据
- redis: 存储备忘录浏览记录

## 项目运行
### 项目使用 go mod 管理

### 下载依赖
```
go mod tidy
```

### 运行
```
go run main.go
```