package api

import (
	"github.com/allentom/haruka"
	"net/http"
	"os"
	"path/filepath"
	"yousync/service"
)

type NewSyncFolderRequestBody struct {
	Path string `json:"path"`
}

var newSyncFolderHandler haruka.RequestHandler = func(context *haruka.Context) {
	var requestBody NewSyncFolderRequestBody
	err := context.ParseJson(&requestBody)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	folder, err := service.NewSyncFolder(requestBody.Path)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	template := BaseSyncFolder{}
	template.Assign(folder)
	context.JSON(template)
}

var readDirectoryHandler haruka.RequestHandler = func(context *haruka.Context) {
	rootPath := context.GetQueryString("path")
	if len(rootPath) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			AbortError(context, err, http.StatusInternalServerError)
			return
		}
		rootPath = homeDir
	}
	infos, err := service.ReadDirectory(rootPath)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	data := make([]BaseFileItemTemplate, 0)
	for _, info := range infos {
		template := BaseFileItemTemplate{}
		template.Assign(info, rootPath)
		data = append(data, template)
	}
	context.JSON(haruka.JSON{
		"path":  rootPath,
		"sep":   string(os.PathSeparator),
		"files": data,
		"back":  filepath.Dir(rootPath),
	})

}

var serviceInfoHandler haruka.RequestHandler = func(context *haruka.Context) {
	context.JSON(haruka.JSON{
		"success": true,
		"name":    "YouSync Service",
	})
}
