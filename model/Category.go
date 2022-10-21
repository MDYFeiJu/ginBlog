package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
}

// CreateCategory 新增分类
func CreateCategory(cg *Category) int {
	if CheckCategory(cg.Name) != errmsg.SUCCESS {
		return errmsg.ErrorCategoryNameUsed
	}
	err := db.Create(&cg).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) int {
	var cg Category
	db.Select("id").Where("name = ?", name).First(&cg)
	if cg.ID > 0 {
		return errmsg.ErrorCategoryNameUsed
	}
	return errmsg.SUCCESS
}

// GetCategory 获取分类列表
func GetCategory(pageSize int, pageNum int) []Category {
	var cg []Category
	if pageNum == -1 {
		err = db.Limit(pageSize).Offset(-1).Find(&cg).Error
	} else {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cg).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cg
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	err = db.Where("id=?", id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditCategory 编辑分类信息
func EditCategory(id int, data *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&Category{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
