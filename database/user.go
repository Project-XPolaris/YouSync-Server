package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid        string
	SyncFolder []SyncFolder `gorm:"foreignKey:Uid;references:Uid"`
}
