package service

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"my-app-backend/db"
	"net/http"
)

type StaticSrcApi interface {
	ADownloadFile(context *gin.Context)
}

type StaticSrcApiHandler struct {
	gridFsDao db.GridFSDao
	logger    *logrus.Logger
}

func (s *StaticSrcApiHandler) ADownloadFile(context *gin.Context) {
	filename := context.Param("filename")
	data, _, err := s.gridFsDao.DownloadFile(filename)
	if err != nil {
		s.logger.Errorf("下载文件: %s 失败", filename)
		context.AbortWithStatus(500)
	}
	reader := bytes.NewReader(data)
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}

	context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	//TODO implement me
	panic("implement me")
}

func (*StaticSrcApiHandler) f() {}
