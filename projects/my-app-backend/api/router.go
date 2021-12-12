package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.New()
	router.Use()
	versionOne := router.Group("/v1")
	versionOne.PUT("/user")
	versionOne.POST("/user")

	versionOne.Use(Auth)
	versionOne.DELETE("/user")
}

func Auth(context *gin.Context) {
	token := context.GetHeader("Authorization")
	fmt.Println(token)
}
