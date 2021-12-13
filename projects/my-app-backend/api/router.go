package api

import (
	"github.com/gin-gonic/gin"
	"my-app-backend/auth"
	"my-app-backend/entity/resp"
	"my-app-backend/service"
	"my-app-backend/util"
	"net/http"
	"strings"
)

var logger = util.NewLogger()

func Route() {
	router := gin.New()
	router.Use(gin.CustomRecovery(func(context *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			context.JSON(http.StatusInternalServerError, resp.NewIntervalError(err))
		} else {
			context.AbortWithStatusJSON(http.StatusInternalServerError, resp.NewIntervalError("unknown error occurred."))
		}
	}))
	// ApiHandler实例化
	staticSrcApi := service.NewStaticSrcApiHandler()
	userApiHandler := service.NewUserApiHandler()
	versionOne := router.Group("/v1")
	{
		versionOne.PUT("/user", userApiHandler.Login)
		versionOne.POST("/user", userApiHandler.SignUp)

		// 静态资源处理器
		versionOne.GET("/static/a/:filename", staticSrcApi.ADownloadFile)

		versionOne.Use(authFunc)
		versionOne.DELETE("/user", userApiHandler.Logout)
	}
	err := router.Run(":8190")
	util.JustPanic(err)
}

func authFunc(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		context.JSON(resp.NewMissToken().Code, resp.NewMissToken())
		return
	}
	token = token[7:]
	username, role, err := auth.ValidToken(token)
	if err != nil {
		context.JSON(resp.NewAuthFailed().Code, resp.NewAuthFailed())
		return
	}
	logger.Errorf("用户: %s 开始操作", username)
	context.Set("username", username)
	context.Set("role", role)
}
