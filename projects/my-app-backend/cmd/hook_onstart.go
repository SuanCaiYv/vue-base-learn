package main

import (
	"io/ioutil"
	config2 "my-app-backend/config"
	"my-app-backend/db"
	"my-app-backend/entity"
	"my-app-backend/util"
	"os"
	"path/filepath"
	"strings"
)

func BeforeStart() {
	config := config2.NewConfiguration()
	sysRoleDao := db.NewSysRoleDaoService()
	for _, val := range config.Roles {
		result, err := sysRoleDao.SelectByName(val.Name)
		util.JustPanic(err)
		if result == nil {
			role := &entity.SysRole{
				Name: val.Name,
				Desc: val.Desc,
			}
			err := sysRoleDao.Insert(role)
			util.JustPanic(err)
		}
	}
	sysUserDao := db.NewSysUserDaoService()
	for _, val := range config.Accounts {
		tmpSysUser, err := sysUserDao.SelectByUsername(val.Username)
		util.JustPanic(err)
		if tmpSysUser != nil {
			continue
		}
		sysUser := entity.DefaultSysUser()
		sysUser.Username = val.Username
		sysUser.Credential = val.Credential
		if config.Owner == val.Username {
			sysUser.Role = entity.RoleOwner
		} else {
			sysUser.Role = entity.RoleReader
		}
		sysUser.Salt = strings.ReplaceAll(util.GenerateUUID(), "-", "")[:6]
		sysUser.Info.Email = val.Username
		sysUser.Info.Nickname = val.Username
		err = sysUserDao.Insert(sysUser)
		util.JustPanic(err)
	}
	gridFsDao := db.NewGridFSDaoService()
	defaultAvatarPath, err := filepath.Abs("static/default-avatar.png")
	util.JustPanic(err)
	defaultAvatar, err := os.OpenFile(defaultAvatarPath, os.O_RDONLY, os.ModePerm)
	util.JustPanic(err)
	data, err := ioutil.ReadAll(defaultAvatar)
	util.JustPanic(err)
	err = gridFsDao.UploadFile(data, "default-avatar.png", make(map[string]interface{}))
	util.JustPanic(err)
}
