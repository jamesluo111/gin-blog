package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		err  error
		tags []Tag
	)
	fmt.Println("参数", pageNum, pageSize)
	fmt.Printf("%v", maps)

	if pageNum > 0 && pageSize > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func GetTagTotal(maps interface{}) (count int, err error) {
	if err = db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.Id > 0 {
		return true, nil
	}
	return false, nil
}

func AddTag(maps map[string]interface{}) error {
	tag := Tag{
		Name:      maps["name"].(string),
		CreatedBy: maps["created_by"].(string),
		State:     maps["state"].(int),
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func ExistTagById(id int) (bool, error) {
	var tag Tag
	if err := db.Select("id").Where("id = ?", id).First(&tag).Error; err != nil {
		return false, err
	}
	if tag.Id > 0 {
		return true, nil
	}
	return false, nil
}

func EditTag(id int, data map[string]interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})
	return true
}
