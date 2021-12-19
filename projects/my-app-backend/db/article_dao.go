package db

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	config2 "my-app-backend/config"
	"my-app-backend/entity"
	"my-app-backend/util"
	"time"
)

type ArticleDao interface {
	Insert(article *entity.Article) error

	Select(id string) (*entity.Article, error)

	SelectByAuthorName(author, name string) (*entity.Article, error)

	// ListByAuthor0 未分页版本
	ListByAuthor0(author string) ([]entity.Article, error)

	ListByAuthor(author string, pgNum, pgSize int64, sort string, desc bool) ([]entity.Article, error)

	Update(article *entity.Article) error

	Delete0(id string) error

	Delete(id string) error
}

type ArticleDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

func NewArticleDaoService() *ArticleDaoService {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	util.JustPanic(err)
	collection := client.Database(config.DatabaseConfig.DB).Collection(CollectionArticle)
	return &ArticleDaoService{
		collection,
		logger,
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
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	one := a.collection.FindOne(timeout, primitive.M{"_id": id})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		a.logger.Error(err)
		return nil, one.Err()
	}
	result := entity.Article{}
	err := one.Decode(&result)
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	return &result, nil
}

func (a *ArticleDaoService) SelectByAuthorName(author, name string) (*entity.Article, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	one := a.collection.FindOne(timeout, primitive.M{"author": author, "name": name})
	if err := one.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		a.logger.Error(err)
		return nil, one.Err()
	}
	result := entity.Article{}
	err := one.Decode(&result)
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	return &result, nil
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	skip := (pgNum - 1) * pgNum
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
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	_, err := a.collection.UpdateByID(timeout, article.Id, article)
	return err
}

func (a *ArticleDaoService) Delete0(id string) error {
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	_, err := a.collection.DeleteOne(timeout, primitive.M{"_id": id})
	return err
}

func (a *ArticleDaoService) Delete(id string) error {
	article, err := a.Select(id)
	if err != nil {
		return err
	}
	article.UpdatedTime = time.Now()
	article.Available = true
	return a.Update(article)
}
