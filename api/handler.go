package api

import (
	"github.com/allentom/haruka"
	"net/http"
	"net/http/httputil"
	"net/url"
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
	folder, err := service.NewSyncFolder(syncPath, requestBody.Path, context.Param["uid"].(string))
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
		"auth":    config.Instance.YouPlusAuth,
	})
}
var youPlusLoginHandler haruka.RequestHandler = func(context *haruka.Context) {
	url, err := url.Parse(config.Instance.YouPlusUrl)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	request := context.Request
	request.URL.Path = "/user/auth"
	proxy.ServeHTTP(context.Writer, request)
}

var getSyncFolderListHandler haruka.RequestHandler = func(context *haruka.Context) {
	list, err := service.GetSyncFolder(
		context.Param["uid"].(string),
		context.Param["page"].(int),
		context.Param["pageSize"].(int),
	)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	data := make([]BaseSyncFolder, 0)
	for _, folder := range list {
		template := BaseSyncFolder{}
		template.Assign(folder)
		data = append(data, template)
	}
	context.JSON(haruka.JSON{
		"success":  true,
		"page":     context.Param["page"].(int),
		"pageSize": context.Param["pageSize"].(int),
		"result":   data,
	})

}

var removeFolderHandler haruka.RequestHandler = func(context *haruka.Context) {
	id, err := context.GetPathParameterAsInt("id")
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	err = service.RemoveSyncFolder(context.Param["uid"].(string), id)
	if err != nil {
		AbortError(context, err, http.StatusBadRequest)
		return
	}
	context.JSON(haruka.JSON{
		"success": true,
	})
}
