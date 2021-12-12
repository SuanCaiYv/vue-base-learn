package db

import (
	context2 "context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	config2 "my-app-backend/config"
	"my-app-backend/util"
	"os"
	"time"
)

type GridFSDaoService struct {
	bucket *gridfs.Bucket
	logger *logrus.Logger
}

type GridFSDao interface {
	UploadFile(fileContent []byte, filename string) error

	ModifyFile(file os.File) error

	DownloadFile(filename string) ([]byte, error)

	DeleteFile(filename string) error
}

func NewGridFSDaoService() *GridFSDaoService {
	logger := util.NewLogger()
	config := config2.NewConfiguration()
	ctx, cancel := context2.WithTimeout(context2.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	util.JustPanic(err)
	bucket, err := gridfs.NewBucket(client.Database(config.DatabaseConfig.GridFSDB))
	util.JustPanic(err)
	return &GridFSDaoService{
		bucket,
		logger,
	}
}

func (g *GridFSDaoService) UploadFile(fileContent []byte, filename string) error {
	uploadStream, err := g.bucket.OpenUploadStream(filename)
	defer func(uploadStream *gridfs.UploadStream) {
		_ = uploadStream.Close()
	}(uploadStream)
	if err != nil {
		g.logger.Error("打开GridFS上传流失败")
		return err
	}
	_, err = uploadStream.Write(fileContent)
	if err != nil {
		g.logger.Error("上传文件至GridFS失败")
		return err
	}
	return nil
}

func (g *GridFSDaoService) ModifyFile(file os.File) error {
	//TODO implement me
	panic("implement me")
}

func (g *GridFSDaoService) DownloadFile(filename string) ([]byte, error) {
	downloadStream, err := g.bucket.OpenDownloadStreamByName(filename)
	defer func(downloadStream *gridfs.DownloadStream) {
		_ = downloadStream.Close()
	}(downloadStream)
	if err != nil {
		g.logger.Error("打开下载流失败")
		return nil, err
	}
	size := downloadStream.GetFile().Length
	data := make([]byte, size, size)
	_, err = downloadStream.Read(data)
	if err != nil {
		g.logger.Error("读取文件失败")
		return nil, err
	}
	return data, nil
}

func (g *GridFSDaoService) DeleteFile(filename string) error {
	ctx, cancel := context2.WithTimeout(context2.Background(), 5*time.Second)
	defer cancel()
	cursor, err := g.bucket.GetFilesCollection().Find(ctx, bson.M{"filename": filename})
	if err != nil {
		g.logger.Errorf("查找files失败，文件名: %s", filename)
		return err
	}
	var gFile *gridfs.File = nil
	for cursor.Next(ctx) {
		file := gridfs.File{}
		err := cursor.Decode(&file)
		if err != nil {
			return err
		}
		gFile = &file
	}
	err = g.bucket.Delete(gFile.ID)
	if err != nil {
		g.logger.Errorf("删除chunks失败，文件名: %s", filename)
		return err
	}
	return nil
}
