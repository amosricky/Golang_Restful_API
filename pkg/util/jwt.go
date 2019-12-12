package util

import (
	"Golang_Restful_API/pkg/setting"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtSalt = []byte(setting.AppSetting.JWTSalt)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, ExtendError) {

	var extendError ExtendError
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSalt)
	if err!=nil{
		extendError = NewBaseError(http.StatusBadRequest, err.Error())
	}

	return token, extendError
}

func ParseToken(token string) (*Claims, ExtendError) {
	var extendError ExtendError
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSalt, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	if err!=nil{
		fmt.Println(err.Error())
		extendError = NewBaseError(ErrorToken, "")
	}

	return nil, extendError
}

func JWTAuth() gin.HandlerFunc{
	return func(c *gin.Context) {

		responseBody := NewResponseBody(NewBaseError(http.StatusOK, ""))

		for{
			authorization := c.Request.Header.Get("Authorization")

			if len(authorization) == 0 {
				responseBody.SetExtendError(NewBaseError(ErrorNoAuth, ""))
				break
			}

			claims, extendError := ParseToken(authorization)
			if extendError != nil{
				responseBody.SetExtendError(extendError)
				break
			}

			if time.Now().Unix() > claims.ExpiresAt{
				responseBody.SetExtendError(NewBaseError(ErrorExpired, ""))
				break
			}
			break
		}

		if responseBody.Code != http.StatusOK{
			c.JSON(responseBody.StatusCode(), responseBody)
			c.Abort()
		}else {
			c.Next()
		}
	}
}
