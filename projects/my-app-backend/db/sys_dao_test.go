package db

import (
	"fmt"
	"my-app-backend/entity"
	"testing"
)

func TestSystemDaoService_Insert(t *testing.T) {
	instance := NewSysUserDaoService()
	err := instance.Insert(entity.DefaultSysUser())
	fmt.Println(err)
}
