package models

import (
	"Golang_Restful_API/pkg/setting"
	"Golang_Restful_API/pkg/util"
	"crypto/md5"
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type User struct {
	Model
	Account string `json:"account"`
	Name string `json:"name"`
	Salt int `json:"salt"`
	AuthCode string `json:"authcode"`
}

type NewUser struct {
	Account string `json:"account"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type LoginUser struct {
	Account string `json:"account"`
	Password string `json:"password"`
}

func (u *User) InitUserTable()  {
	if !db.HasTable(&User{}) {
		if err := db.CreateTable(&User{}).Error; err != nil {
			logrus.Panic(err)
		}
	}
}

func Register(newUser NewUser) util.ExtendError{

	var extendError util.ExtendError

	for{
		if ExistUserByAccount(newUser.Account){
			extendError = util.NewBaseError(util.ErrorExistUser, "")
			break
		}

		account := newUser.Account
		name := newUser.Name
		salt := GenerateSalt()
		authCode := GenerateAuthCode(salt, newUser.Password)

		db.Create(&User {
			Account: account,
			Name : name,
			Salt: salt,
			AuthCode: authCode,
		})

		break
	}
	return extendError
}

func Login(login LoginUser) (string, util.ExtendError) {

	var extendError util.ExtendError
	var loggingUser User
	var JWTToken string

	for{
		db.Where("account = ?", login.Account).First(&loggingUser)
		if loggingUser.ID == 0{
			extendError = util.NewBaseError(util.ErrorNotExistUser, "")
			break
		}

		loggingAuthCode := GenerateAuthCode(loggingUser.Salt, login.Password)
		trueAuthCode    := loggingUser.AuthCode

		if loggingAuthCode != trueAuthCode{
			extendError = util.NewBaseError(util.ErrorPasswdWrong, "")
			break
		}

		JWTNewToken, JWTExtendError := GenerateJWTToken(loggingUser.Name)
		if JWTExtendError != nil{
			extendError = JWTExtendError
			break
		}

		JWTToken = JWTNewToken
		break
	}

	return JWTToken, extendError
}

func GenerateAuthCode(salt int, password string) string {
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)

	return hex.EncodeToString(st)
}

func GenerateSalt() int{
	max := setting.ServerSetting.RandomMax
	min := setting.ServerSetting.RandomMin

	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min

	return randNum
}

func GenerateJWTToken(userName string) (string, util.ExtendError)  {

	JWTToken, extendError := util.GenerateToken(userName)

	return JWTToken, extendError

}

func ExistUserByAccount(account string) bool {
	var user User
	db.Where("account = ? AND deleted_on = ? ", account, 0).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}