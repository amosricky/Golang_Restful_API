package models

import "github.com/sirupsen/logrus"

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Auth) InitAuthTable()  {
	if !db.HasTable(&Auth{}) {
		if err := db.CreateTable(&Auth{}).Error; err != nil {
			logrus.Panic(err)
		}
	}
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}