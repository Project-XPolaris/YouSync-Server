package database

import "gorm.io/gorm"

type SyncFolder struct {
	gorm.Model
	Path string
}
