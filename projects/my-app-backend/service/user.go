package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"my-app-backend/db"
	"my-app-backend/entity/resp"
)

type UserApi interface {
	SignUp(context *gin.Context)

	Login(context *gin.Context)

	Logout(context *gin.Context)

	GetUserInfo(context *gin.Context)

	UpdateUserInfo(context *gin.Context)
}

type UserApiHandler struct {
	sysUserDao db.SysUserDao
	gridFsDao  db.GridFSDao
	logger     *logrus.Logger
}

type sign struct {
	Username   string `json:"username"`
	Credential string `json:"credential"`
	VerCode    string `json:"verCode"`
}

func (u *UserApiHandler) SignUp(context *gin.Context) {
	signup := &sign{}
	err := context.BindJSON(signup)
	if err != nil {
		u.logger.Error("参数解析失败")
		context.JSON(400, resp.NewBadRequest("参数解析失败"))
	}
	//TODO implement me
	panic("implement me")
}

func (u *UserApiHandler) Login(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserApiHandler) Logout(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserApiHandler) GetUserInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserApiHandler) UpdateUserInfo(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}
