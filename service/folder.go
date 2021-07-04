package service

import (
	"errors"
	"os"
	"yousync/database"
)

func NewSyncFolder(path string, displayPath string, uid string) (*database.SyncFolder, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New("target path is not directory")
	}
	folder := &database.SyncFolder{Path: path, Uid: uid, DisplayPath: displayPath}
	err = database.Instance.Save(folder).Error
	return folder, err
}

func GetSyncFolder(uid string, page int, pageSize int) ([]*database.SyncFolder, error) {
	var list []*database.SyncFolder
	err := database.Instance.Where("uid = ?", uid).Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return list, err
}

func RemoveSyncFolder(uid string, id int) error {
	err := database.Instance.Unscoped().Where("uid = ?", uid).Where("id = ?", id).Delete(&database.SyncFolder{}).Error
	return err
}
