package models

import (
	"mathTube/base"
)

type Person struct {
	ID        int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Fullname  string `json:"fullName" gorm:"size:250"`
	Login     string `json:"login" gorm:"unique"`
	Password  string `json:"password"`
	Gender    bool   `json:"gender"`
	Access    int    `json:"access"`
	LastVisit string `json:"lastVisit"`
}

func (p *Person) Authentication() error {
	db := base.GetDB()
	ps := []Person{}
	err := db.Model(p).Where(p).Find(&ps).Error
	if err != nil {
		return err
	}
	return nil
}
