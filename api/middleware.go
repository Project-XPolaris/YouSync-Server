package api

import (
	"github.com/allentom/haruka"
	"strings"
	"yousync/config"
	"yousync/service"
	"yousync/youplus"
)

var noAuthPath = []string{
	"/user/auth",
}

type AuthMiddleware struct {
}

func (a *AuthMiddleware) OnRequest(ctx *haruka.Context) {
	if !config.Instance.YouPlusAuth {
		ctx.Param["uid"] = service.PublicUid
		ctx.Param["username"] = service.PublicUsername
		ctx.Param["token"] = ""
		return
	}
	for _, targetPath := range noAuthPath {
		if ctx.Request.URL.Path == targetPath {
			return
		}
	}
	rawString := ctx.Request.Header.Get("Authorization")
	if len(rawString) == 0 {
		rawString = ctx.GetQueryString("token")
	}
	ctx.Param["token"] = rawString
	if len(rawString) > 0 {
		rawString = strings.Replace(rawString, "Bearer ", "", 1)
		response, err := youplus.DefaultClient.CheckAuth(rawString)
		if err == nil && response.Success {
			ctx.Param["uid"] = response.Uid
			ctx.Param["username"] = response.Username
		} else {
			ctx.Param["uid"] = service.PublicUid
			ctx.Param["username"] = service.PublicUsername
		}
	} else {
		ctx.Param["uid"] = service.PublicUid
		ctx.Param["username"] = service.PublicUsername
	}
}

type ReadUserMiddleware struct {
}

func (m *ReadUserMiddleware) OnRequest(ctx *haruka.Context) {
	if ctx.Param["uid"] == nil {
		return
	}
	user, _ := service.GetUserById(ctx.Param["uid"].(string))
	ctx.Param["user"] = user
}
