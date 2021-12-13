package db

import (
	context2 "context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	config2 "my-app-backend/config"
	"my-app-backend/entity"
	"my-app-backend/util"
	"time"
)

type SysUserDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

type SysRoleDaoService struct {
	collection *mongo.Collection
	logger     *logrus.Logger
}

type SysUserDao interface {
	Insert(sysUser *entity.SysUser) error

	Select(id string) (*entity.SysUser, error)

	Update(sysUser *entity.SysUser) error

	Delete(id string) error

	SelectByUsername(username string) (*entity.SysUser, error)

	SelectByNickname(nickname string) (*entity.SysUser, error)
}

type SysRoleDao interface {
	Insert(sysRole *entity.SysRole) error

	Select(id string) (*entity.SysRole, error)

	SelectByName(name string) (*entity.SysRole, error)
}

func NewSysUserDaoService() *SysUserDaoService {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context2.WithTimeout(context2.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	util.JustPanic(err)
	collection := client.Database(config.DatabaseConfig.DB).Collection(CollectionSysUser)
	return &SysUserDaoService{
		collection,
		logger,
	}
}

func (s *SysUserDaoService) Insert(sysUser *entity.SysUser) error {
	sysUser.Id = primitive.NewObjectID().Hex()
	sysUser.Available = true
	sysUser.CreatedTime = time.Time{}
	sysUser.UpdatedTime = time.Time{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := s.collection.InsertOne(ctx, sysUser)
	return err
}

func (s *SysUserDaoService) Select(id string) (*entity.SysUser, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SysUserDaoService) Update(sysUser *entity.SysUser) error {
	//TODO implement me
	panic("implement me")
}

func (s *SysUserDaoService) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SysUserDaoService) SelectByUsername(username string) (*entity.SysUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := s.collection.Find(ctx, primitive.M{"username": username})
	if err != nil {
		s.logger.Error(err)
	}
	defer func(cursor *mongo.Cursor, ctx context2.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			s.logger.Error(err)
		}
	}(cursor, ctx)
	var result *entity.SysUser = nil
	for cursor.Next(ctx) {
		sysUser := entity.SysUser{}
		err := cursor.Decode(&sysUser)
		if err != nil {
			return nil, err
		}
		result = &sysUser
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SysUserDaoService) SelectByNickname(nickname string) (*entity.SysUser, error) {
	//TODO implement me
	panic("implement me")
}

func NewSysRoleDaoService() *SysRoleDaoService {
	logger := util.NewLogger()
	config := config2.ApplicationConfiguration()
	ctx, cancel := context2.WithTimeout(context2.Background(), 2*time.Second)
	defer cancel()
	url := fmt.Sprintf("%s:%d", config.DatabaseConfig.Url, config.DatabaseConfig.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	util.JustPanic(err)
	collection := client.Database(config.DatabaseConfig.DB).Collection(CollectionSysRole)
	return &SysRoleDaoService{
		collection,
		logger,
	}
}

func (s *SysRoleDaoService) Insert(sysRole *entity.SysRole) error {
	sysRole.Id = primitive.NewObjectID().Hex()
	sysRole.Available = true
	sysRole.CreatedTime = time.Time{}
	sysRole.UpdatedTime = time.Time{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := s.collection.InsertOne(ctx, sysRole)
	return err
}

func (s *SysRoleDaoService) Select(id string) (*entity.SysRole, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SysRoleDaoService) SelectByName(name string) (*entity.SysRole, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := s.collection.Find(ctx, primitive.M{"name": name})
	if err != nil {
		log.Fatal(err)
	}
	defer func(cursor *mongo.Cursor, ctx context2.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			s.logger.Error(err)
		}
	}(cursor, ctx)
	var result *entity.SysRole = nil
	for cursor.Next(ctx) {
		sysRole := entity.SysRole{}
		err := cursor.Decode(&sysRole)
		if err != nil {
			return nil, err
		}
		result = &sysRole
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
