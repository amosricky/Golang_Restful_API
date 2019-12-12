package util

const (
	ErrorInvalidErrCode     = 10000
	ErrorInvalidID 			= 10001
	ErrorNotExistAuthor		= 10002
	ErrorNotExistArticle    = 10003
	ErrorNotExistUser		= 10004
	ErrorExistUser		    = 10005
	ErrorPasswdWrong        = 10006
	ErrorNoAuth             = 10007
	ErrorToken				= 10008
	ErrorExpired			= 10009
	ErrorImageNotExist	    = 10010
	ErrorImageFormat	    = 10011
)

var MsgFlags = map[int]string{
	ErrorInvalidErrCode:              "Invalid error code",
	ErrorInvalidID: 				  "Invalid ID",
	ErrorNotExistAuthor:  			  "Author not exist",
	ErrorNotExistArticle:             "Article not exist",
	ErrorNotExistUser:				  "User not exist",
	ErrorExistUser:                   "User account already exist",
	ErrorPasswdWrong:				  "Wrong password",
	ErrorNoAuth:					  "No auth token",
	ErrorToken:                       "Token error",
	ErrorExpired:					  "Token expired",
	ErrorImageNotExist:               "Image not exist",
	ErrorImageFormat:                 "Image format not correct",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ErrorInvalidErrCode]
}

