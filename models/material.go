package models

import (
	"mathTube/base"
	"time"

	"gorm.io/gorm"
)

type Material struct {
	ID               int            `json:"id" gorm:"autoIncrement;primaryKey"`
	Address          string         `json:"address"`
	Name             string         `json:"name"`
	ShortDescription string         `json:"shortDescription"`
	FullDescription  string         `json:"fullDescription"`
	Votes            int            `json:"votes"`
	Status           int            `json:"status"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-"`
}

type Materials []Material

func (m *Material) Create() error {
	db := base.GetDB()
	return db.Model(Material{}).Create(m).Error
}

func (m *Materials) ReadAll() error {
	db := base.GetDB()
	return db.Model(Material{}).Find(m).Error
}

// 0, false, "" update кор намекунад
func (m *Material) Update() error {
	db := base.GetDB()
	return db.Model(Material{}).Where("id", m.ID).Updates(m).Error
}

// хатман майдони votes иваз мешад
func (m *Material) UpdateVote() error {
	db := base.GetDB()
	return db.Model(Material{}).Where("id", m.ID).Updates(map[string]interface{}{"votes": m.Votes}).Error
}

func (m *Material) Delete() error {
	db := base.GetDB()
	return db.Model(Material{}).Where("id", m.ID).Delete(m).Error
}
