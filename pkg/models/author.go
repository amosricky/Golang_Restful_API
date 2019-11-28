package models

import (
	"github.com/sirupsen/logrus"
)

type Author struct {
	Model
	Name string `json:"name"`
	Born int `json:"born"`
}

func (a *Author) InitTagTable()  {
	if !db.HasTable(&Author{}) {
		if err := db.CreateTable(&Author{}).Error; err != nil {
			logrus.Panic(err)
		}
	}
}

func GetAuthor(pageNum int, pageSize int, maps interface {}) (author []Author) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&author)

	return
}

func AddAuthor(name string, born int) bool{

	db.Create(&Author {
		Name : name,
		Born : born,
	})

	return true
}

func EditAuthor(id int, name string, born int) bool {
	var author Author

	db.Where("id = ?", id).First(&author)
	db.Model(&author).Updates(map[string]interface{}{"name": name, "born": born})

	return true
}

func DeleteAuthor(id int) bool {
	db.Where("id = ?", id).Delete(&Author{})

	return true
}

func ExistAuthorByName(name string) bool {
	var author Author
	db.Select("id").Where("name = ?", name).First(&author)
	if author.ID > 0 {
		return true
	}

	return false
}

func ExistAuthorByID(id int) bool {
	var author Author
	db.Select("id").Where("id = ?", id).First(&author)
	if author.ID > 0 {
		return true
	}

	return false
}