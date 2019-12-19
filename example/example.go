package example

// Common
type OK struct {
	Code    int    `json:"error_code" example:"200"`
	Message string `json:"error_message" example:"OK"`
	Result  string `json:"result" example:"OK"`
}

type SuccessCommon struct {
	Code    int    `json:"error_code" example:"200"`
	Message string `json:"error_message" example:"OK"`
}

// default
type PostRegister struct {
	account 	string `json:"account" example:"edison1234"`
	name 		string `json:"name" example:"edison"`
	password 	string `json:"password" example:"test1234"`
}

type PostLogin struct {
	account 	string `json:"account" example:"edison1234"`
	password 	string `json:"password" example:"test1234"`
}

type UploadImage struct {
	SuccessCommon
	Result	UploadImageInfo
}

type UploadImageInfo struct {
	encode_file_name 	string `json:"encode_file_name" example:"fd456406745d816a45cae554c788e754.jpeg"`
	image_url 			string `json:"image_url" example:"http://127.0.0.1:8000/images/fd456406745d816a45cae554c788e754.jpeg"`
}

// Author
type GetAuthors struct {
	SuccessCommon
	Result  []AuthorInfo `json:"result"`
}

type GetAuthor struct {
	SuccessCommon
	Result  AuthorInfo `json:"result"`
}

type PostAuthor struct {
	name    		string `json:"name" example:"Thomas Alva Edison"`
	born    		int    `json:"born" example:"1847"`
}

type AuthorInfo struct {
	id    			int    `json:"id" example:"1"`
	created_on    	int    `json:"created_on" example:"1574952062"`
	modified_on    	int    `json:"modified_on" example:"1575189888"`
	deleted_on    	int    `json:"deleted_on" example:"0"`
	name    		string `json:"name" example:"William Shakespeare"`
	born    		int    `json:"born" example:"1564"`
}

// Article
type GetArticle struct {
	SuccessCommon
	Result  ArticleInfo `json:"result"`
}

type GetArticles struct {
	SuccessCommon
	Result  []ArticleInfo `json:"result"`
}

type PostArticle struct {
	author_id    	int    		`json:"author_id" example:"1"`
	title    		string 		`json:"title" example:"The Merchant Venice "`
	desc    		string 		`json:"desc" example:"The Merchant of Venice is a 16th-century play written by William Shakespeare"`
	content    		string 		`json:"content" example:"Although classified as a comedy...."`
	image_url    	string 		`json:"image_url" example:"http://xxx.xxx.xxx.jpeg"`
}

type ArticleInfo struct {
	id    			int    		`json:"id" example:"1"`
	created_on    	int    		`json:"created_on" example:"1574952062"`
	modified_on    	int    		`json:"modified_on" example:"1575189888"`
	deleted_on    	int    		`json:"deleted_on" example:"0"`
	author_id    	int    		`json:"author_id" example:"1"`
	author 			AuthorInfo 	`json:"author"`
	title    		string 		`json:"title" example:"The Merchant Venice "`
	desc    		string 		`json:"desc" example:"The Merchant of Venice is a 16th-century play written by William Shakespeare"`
	content    		string 		`json:"content" example:"Although classified as a comedy...."`
	image_url    	string 		`json:"image_url" example:"http://xxx.xxx.xxx.jpeg"`
}

// Error
type InvalidID struct {
	Code    int    `json:"error_code" example:"10001"`
	Message string `json:"error_message" example:"Invalid ID"`
	Result  string `json:"result"`
}

type AuthorNotExist struct {
	Code    int    `json:"error_code" example:"10002"`
	Message string `json:"error_message" example:"Author not exist"`
	Result  string `json:"result"`
}

type ArticleNotExist struct {
	Code    int    `json:"error_code" example:"10003"`
	Message string `json:"error_message" example:"Article not exist"`
	Result  string `json:"result"`
}

type ErrorExistUser struct {
	Code    int    `json:"error_code" example:"10005"`
	Message string `json:"error_message" example:"User account already exist"`
	Result  string `json:"result"`
}

type ErrorPasswdWrong struct {
	Code    int    `json:"error_code" example:"10006"`
	Message string `json:"error_message" example:"Wrong password""`
	Result  string `json:"result"`
}

type ErrorImageFormat struct {
	Code    int    `json:"error_code" example:"10011"`
	Message string `json:"error_message" example:"Image format not correct""`
	Result  string `json:"result"`
}