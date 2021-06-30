package api

import (
	"github.com/allentom/haruka"
	"net/http"
	"os"
	"path/filepath"
	"yousync/config"
	"yousync/service"
	"yousync/youplus"
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
	syncPath := requestBody.Path
	if config.Instance.YouPlusPath {
		realPath, err := youplus.DefaultClient.GetRealPath(requestBody.Path, context.Param["token"].(string))
		if err != nil {
			AbortError(context, err, http.StatusBadRequest)
			return
		}
		syncPath = realPath
	}
	folder, err := service.NewSyncFolder(syncPath)
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
	if config.Instance.YouPlusPath {
		token := context.Param["token"].(string)
		items, err := youplus.DefaultClient.ReadDir(rootPath, token)
		if err != nil {
			AbortError(context, err, http.StatusInternalServerError)
			return
		}
		data := make([]BaseFileItemTemplate, 0)
		for _, item := range items {
			template := BaseFileItemTemplate{}
			template.AssignWithYouPlusItem(item)
			data = append(data, template)
		}
		context.JSON(haruka.JSON{
			"path":  rootPath,
			"sep":   "/",
			"files": data,
			"back":  filepath.Dir(rootPath),
		})
		return
	} else {
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

}

var serviceInfoHandler haruka.RequestHandler = func(context *haruka.Context) {
	context.JSON(haruka.JSON{
		"success": true,
		"name":    "YouSync Service",
	})
}
