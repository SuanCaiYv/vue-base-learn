package config

import (
	"encoding/json"
	"my-app-backend/util"
	"os"
	"path/filepath"
)

func init() {
	absPath, err := filepath.Abs("../application.json")
	util.JustPanic(err)
	confFile, err := os.OpenFile(absPath, os.O_RDONLY, os.ModePerm)
	util.JustPanic(err)
	fileInfo, err := confFile.Stat()
	util.JustPanic(err)
	bytes := make([]byte, fileInfo.Size(), fileInfo.Size())
	_, err = confFile.Read(bytes)
	util.JustPanic(err)
	configObject := make(map[string]interface{})
	err = json.Unmarshal(bytes, &configObject)
	util.JustPanic(err)

	database := configObject["database"].(map[string]interface{})
	tmpRoles := configObject["roles"].([]interface{})
	roles := make([]Role, 0, len(tmpRoles))
	for _, val := range tmpRoles {
		role := val.(map[string]interface{})
		roles = append(roles, Role{
			Name: role["name"].(string),
			Desc: role["desc"].(string),
		})
	}
	redis := configObject["redis"].(map[string]interface{})
	tmpAccounts := configObject["accounts"].([]interface{})
	accounts := make([]Account, 0, len(tmpAccounts))
	for _, val := range tmpAccounts {
		tmpAccount := val.(map[string]interface{})
		accounts = append(accounts, Account{
			Username:   tmpAccount["username"].(string),
			Credential: tmpAccount["credential"].(string),
			VerCode:    tmpAccount["ver_code"].(string),
		})
	}
	config = &Configuration{
		Owner: configObject["owner"].(string),
		DatabaseConfig: &DatabaseConfig{
			Url:      database["url"].(string),
			Port:     int(database["port"].(float64)),
			DB:       database["db"].(string),
			GridFSDB: database["grid_fs_db"].(string),
			User:     database["user"].(string),
			Password: database["password"].(string),
		},
		RedisConfig: &RedisConfig{
			Url:      redis["url"].(string),
			Port:     int(redis["port"].(float64)),
			DB:       int(redis["db"].(float64)),
			User:     redis["user"].(string),
			Password: redis["password"].(string),
		},
		Roles:    roles,
		Accounts: accounts,
	}
}

type Configuration struct {
	Owner          string
	DatabaseConfig *DatabaseConfig
	RedisConfig    *RedisConfig
	Roles          []Role
	Accounts       []Account
}

type DatabaseConfig struct {
	Url      string
	Port     int
	DB       string
	GridFSDB string
	User     string
	Password string
}

type RedisConfig struct {
	Url      string
	Port     int
	DB       int
	User     string
	Password string
}

type Role struct {
	Name string
	Desc string
}

type Account struct {
	Username   string
	Credential string
	VerCode    string
}

var config *Configuration

func ApplicationConfiguration() *Configuration {
	return config
}