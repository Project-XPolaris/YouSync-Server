package api

import (
	"github.com/allentom/haruka"
	"net/http"
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
	_, err = service.NewSyncFolder(requestBody.Path)
	if err != nil {
		AbortError(context, err, http.StatusInternalServerError)
		return
	}
	context.JSON(haruka.JSON{
		"success": true,
	})
}
