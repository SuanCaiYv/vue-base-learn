package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"my-app-backend/config"
	"my-app-backend/db"
	"my-app-backend/entity/resp"
	"my-app-backend/util"
	"strconv"
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

func NewArticleApiHandler() *ArticleApiHandler {
	return &ArticleApiHandler{
		articleDao: db.NewArticleDaoService(),
		gridFsDao:  db.NewGridFSDaoService(),
		logger:     util.NewLogger(),
	}
}

func (a *ArticleApiHandler) AddArticle(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleApiHandler) UpdateArticle(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleApiHandler) DeleteArticle(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleApiHandler) ListArticle(context *gin.Context) {
	pgSize, _ := strconv.Atoi(context.DefaultQuery("pgSize", "10"))
	pgNum, _ := strconv.Atoi(context.DefaultQuery("pgNum", "1"))
	sort := context.DefaultQuery("sort", "created_time")
	desc, _ := strconv.ParseBool(context.DefaultQuery("desc", "true"))
	owner := config.ApplicationConfiguration().Owner
	articles, err := a.articleDao.ListByAuthor(owner, int64(pgNum), int64(pgSize), sort, desc)
	if err != nil {
		a.logger.Errorf("获取文章列表失败: %v", err)
		context.JSON(200, resp.NewIntervalError("获取文章列表失败"))
		return
	}
	context.JSON(200, resp.NewOk(articles))
}

func (*ArticleApiHandler) f() {}
