package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"my-app-backend/entity"
	"time"
)

type ArticleDao interface {
	Insert(article *entity.Article) error

	Select(id string) (*entity.Article, error)

	SelectByName(name string) (*entity.Article, error)

	// ListByAuthor0 未分页版本
	ListByAuthor0(author string) ([]entity.Article, error)

	ListByAuthor(author string, pgNum, pgSize int64, sort string, desc bool) ([]entity.Article, error)

	Update(article *entity.Article) error
}

type ArticleDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

func NewArticleDaoService() *ArticleDaoService {
	return &ArticleDaoService{
		collection: nil,
		logger:     nil,
	}
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

func (a *ArticleDaoService) ListByAuthor(author string, pgNum, pgSize int64, sort string, desc bool) ([]entity.Article, error) {
	descInt := 1
	if desc {
		descInt = -1
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	skip := (pgNum-1) * pgNum
	cursor, err := a.collection.Find(ctx, primitive.M{"author": author}, &options.FindOptions{
		Limit: &pgSize,
		Skip:  &skip,
		Sort:  primitive.M{sort: descInt},
	})
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			a.logger.Error(err)
		}
	}(cursor, ctx)
	if err != nil {
		return nil, err
	}
	results := make([]entity.Article, 0, pgSize)
	for cursor.Next(ctx) {
		tmp := entity.Article{}
		err := cursor.Decode(&tmp)
		if err != nil {
			a.logger.Error(err)
			return nil, err
		}
		results = append(results, tmp)
	}
	if err := cursor.Err(); err != nil {
		a.logger.Error(err)
		return nil, err
	}
	return results, nil
}

func (a *ArticleDaoService) Update(article *entity.Article) error {
	//TODO implement me
	panic("implement me")
}
