package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"my-app-backend/entity"
	"time"
)

type ArticleDao interface {
	Insert(article *entity.Article) error

	Select(id string) (*entity.Article, error)

	SelectByName(name string) (*entity.Article, error)

	// ListByAuthor0 未分页版本
	ListByAuthor0(author string) ([]entity.Article, error)

	ListByAuthor(author string, pgNum, pgSize int, sort string, desc bool) ([]entity.Article, error)

	Update(article *entity.Article) error
}

type ArticleDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

func (a *ArticleDaoService) Insert(article *entity.Article) error {
	article.Available = true
	article.CreatedTime = time.Now()
	article.UpdatedTime = time.Now()
	article.ReleaseTime = time.Now()
	article.Id = primitive.NewObjectID().Hex()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := a.collection.InsertOne(ctx, article)
	return err
}

func (a *ArticleDaoService) Select(id string) (*entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleDaoService) SelectByName(name string) (*entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleDaoService) ListByAuthor0(author string) ([]entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleDaoService) ListByAuthor(author string, pgNum, pgSize int, sort string, desc bool) ([]entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleDaoService) Update(article *entity.Article) error {
	//TODO implement me
	panic("implement me")
}
