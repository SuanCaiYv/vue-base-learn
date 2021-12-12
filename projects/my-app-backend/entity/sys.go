package entity

import (
	"my-app-backend/util"
	"time"
)

type SysUser struct {
	Id          string    `bson:"_id" json:"-"`
	Username    string    `bson:"username" json:"username"`
	Credential  string    `bson:"credential" json:"credential"`
	Salt        string    `bson:"salt" json:"salt"`
	Role        string    `bson:"role" json:"role"`
	Info        UserInfo  `bson:"info" json:"info"`
	Available   bool      `bson:"available" json:"-"`
	CreatedTime time.Time `bson:"created_time" json:"-"`
	UpdatedTime time.Time `bson:"updated_time" json:"-"`
}

type SysRole struct {
	Id          string    `bson:"_id" json:"-"`
	Name        string    `bson:"name" json:"name"`
	Desc        string    `bson:"desc" json:"desc"`
	Available   bool      `bson:"available" json:"-"`
	CreatedTime time.Time `bson:"created_time" json:"-"`
	UpdatedTime time.Time `bson:"updated_time" json:"-"`
}

func DefaultSysUser() *SysUser {
	return &SysUser{
		Username:   "default-username" + util.GenerateUUID(),
		Credential: "default-credential" + util.GenerateUUID(),
		Salt:       util.GenerateUUID(),
		Role:       "default-role",
		Info: UserInfo{
			Avatar:    "default-avatar",
			Nickname:  "default-nickname" + util.GenerateUUID(),
			Email:     "default-email" + util.GenerateUUID(),
			Phone:     "default-phone" + util.GenerateUUID(),
			WeChat:    "default-wechat" + util.GenerateUUID(),
			QQ:        "default-qq" + util.GenerateUUID(),
			GitHub:    "default-github" + util.GenerateUUID(),
			Location:  "default-location" + util.GenerateUUID(),
			Signature: "default-signature" + util.GenerateUUID(),
		},
	}
}
