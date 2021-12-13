package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"my-app-backend/db"
)

type ArticleApi interface {
	AddArticle(context *gin.Context)

	UpdateArticle(context *gin.Context)

	DeleteArticle(context *gin.Context)

	ListArticle(context *gin.Context)
}

type ArticleApiHandler struct {
	articleDao db.ArticleDao
	gridFsDao  db.GridFSDao
	logger     *logrus.Logger
}

func (*ArticleApiHandler) f() {}
