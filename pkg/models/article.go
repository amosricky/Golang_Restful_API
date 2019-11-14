package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func (a *Article) InitAriticleTable()  {
	if !db.HasTable(&Article{}) {
		if err := db.CreateTable(&Article{}).Error; err != nil {
			panic(err)
		}
	}
}

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id int) bool {
	var article Article
	db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]Article, error) {
	var articles []Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

// GetArticle Get a single article based on ID
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

// EditArticle modify a single article
func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddArticle add a single article
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

// DeleteArticle delete a single article
func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

