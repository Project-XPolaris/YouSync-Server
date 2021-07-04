package api

import (
	"os"
	"path/filepath"
	"yousync/database"
	"yousync/youplus"
)

type BaseSyncFolder struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func (f *BaseSyncFolder) Assign(folder *database.SyncFolder) {
	f.Id = folder.ID
	f.Name = filepath.Base(folder.Path)
	f.Path = folder.DisplayPath
}

type BaseFileItemTemplate struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func (t *BaseFileItemTemplate) Assign(info os.FileInfo, rootPath string) {
	if info.IsDir() {
		t.Type = "Directory"
	} else {
		t.Type = "File"
	}
	t.Name = info.Name()
	t.Path = filepath.Join(rootPath, info.Name())
}
func (t *BaseFileItemTemplate) AssignWithYouPlusItem(item youplus.ReadDirItem) {
	t.Type = item.Type
	t.Path = item.Path
	t.Name = filepath.Base(item.Path)
}
