package util

const (
	ErrorInvalidErrCode     = 10000
	ErrorInvalidID 			= 10001
	ErrorNotExistAuthor		= 10002
	ErrorNotExistArticle    = 10003
)

var MsgFlags = map[int]string{
	ErrorInvalidErrCode:              "Invalid Error Code",
	ErrorInvalidID: 				  "Invalid ID",
	ErrorNotExistAuthor:  			  "Author not exist",
	ErrorNotExistArticle:             "Article not exist",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ErrorInvalidErrCode]
}

