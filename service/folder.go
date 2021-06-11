package service

import (
	"errors"
	"os"
	"yousync/database"
)

func NewSyncFolder(path string) (*database.SyncFolder, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New("target path is not directory")
	}
	folder := &database.SyncFolder{Path: path}
	err = database.Instance.Save(folder).Error
	return folder, err
}
