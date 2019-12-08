package v1

import (
	"Golang_Restful_API/pkg/models"
	"Golang_Restful_API/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Register(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))

	for{
		var register models.NewUser
		err := c.BindJSON(&register)
		if err != nil{
			logrus.Errorf("Register :%v", err)
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			break
		}

		account := register.Account
		name := register.Name
		password := register.Password

		valid := validation.Validation{}
		valid.Required(account, "account").Message("[account] can't be empty!")
		valid.MinSize(account,6, "account" ).Message("[account] is too short! (min: 6, max:12)")
		valid.MaxSize(account,12, "account").Message("[account] is too long! (min: 6, max:12)")
		valid.Required(name, "name").Message("[name] can't be empty!")
		valid.MaxSize(name, 12, "name").Message("[name] is too long! (max:12)")
		valid.Required(password, "password").Message("[password] can't be empty!")
		valid.MinSize(password,6, "password" ).Message("[password] is too short! (min: 6, max:12)")
		valid.MaxSize(password,12, "password").Message("[password] is too long! (min: 6, max:12)")
		valid.AlphaNumeric(password, "password").Message("[password] must contain alphabet and integer")

		if valid.HasErrors() {
			errMsg := valid.Errors
			logrus.Errorf("Register :%v", errMsg[0].Message)
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, errMsg[0].Message))
		}else if extendError := models.Register(register); extendError!=nil{
			logrus.Errorf("Register :%v", extendError.Error())
			responseBody.SetExtendError(extendError)
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

func Login(c *gin.Context)  {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))
	data := make(map[string]interface{})

	for{
		var login models.LoginUser
		err := c.BindJSON(&login)
		if err != nil{
			logrus.Errorf("Login :%v", err)
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			break
		}

		account := login.Account
		password := login.Password

		valid := validation.Validation{}
		valid.Required(account, "account").Message("[account] can't be empty!")
		valid.Required(password, "password").Message("[password] can't be empty!")

		if valid.HasErrors() {
			errMsg := valid.Errors
			logrus.Errorf("Login :%v", errMsg[0].Message)
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, errMsg[0].Message))
			break
		}

		JWTToken, extendError := models.Login(login)
		if extendError != nil{
			logrus.Errorf("Login :%v", extendError.Error())
			responseBody.SetExtendError(extendError)
			break
		}

		data["token"] = JWTToken
		responseBody.Result = data
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

