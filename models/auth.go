package models

import "github.com/jinzhu/gorm"

type Auth struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckoutAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where("username = ? and password = ?", username, password).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.Id > 0 {
		return true, nil
	}

	return false, nil
}
