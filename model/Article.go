package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int" json:"cid"`
	Desc    string `gorm:"type:varchar(200);not null" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArt 创建文章
func CreateArt(data *Article) int {
	if CheckArt(data.Title) == errmsg.ErrorExistArticle {
		return errmsg.ErrorExistArticle
	}
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CheckArt(name string) int {
	var art Article
	db.Select("id").Where("title = ?", name).First(&art)
	if art.ID > 0 {
		return errmsg.ErrorExistArticle
	}
	return errmsg.SUCCESS
}

func GetArts(pageSize int, pageNum int) []Article {
	var arts []Article
	if pageNum == -1 {
		err = db.Limit(pageSize).Offset(-1).Find(&arts).Error
	} else {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return arts
}

func DeleteArt(id int) int {
	err = db.Where("id=?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func EditArt(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&Article{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
