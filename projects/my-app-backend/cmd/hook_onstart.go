package main

import (
	config2 "my-app-backend/config"
	"my-app-backend/db"
	"my-app-backend/entity"
)

func InitSysRoleDB() {
	config := config2.NewConfiguration()
	sysRoleDao := db.NewSysRoleDaoService()
	for _, val := range config.Roles {
		result, _ := sysRoleDao.SelectByName(val.Name)
		if result == nil {
			role := &entity.SysRole{
				Name: val.Name,
				Desc: val.Desc,
			}
			_ = sysRoleDao.Insert(role)
		}
	}
}
