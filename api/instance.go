package api

import (
	"github.com/allentom/haruka"
	"github.com/allentom/haruka/middleware"
	"github.com/rs/cors"
	"yousync/config"
)

func RunAPIService() {
	e := haruka.NewEngine()
	e.UseCors(cors.AllowAll())
	e.UseMiddleware(middleware.NewLoggerMiddleware())
	e.UseMiddleware(&AuthMiddleware{})
	e.UseMiddleware(&ReadUserMiddleware{})
	e.UseMiddleware(middleware.NewPaginationMiddleware("page", "pageSize", 1, 20))
	e.Router.DELETE("/sync/folder/{id:[0-9]+}", removeFolderHandler)
	e.Router.POST("/sync/folder", newSyncFolderHandler)
	e.Router.GET("/sync/folder", getSyncFolderListHandler)
	e.Router.GET("/explore/readdir", readDirectoryHandler)
	e.Router.GET("/info", serviceInfoHandler)
	e.Router.POST("/user/auth", youPlusLoginHandler)
	e.Router.GET("/user/auth", youPlusLoginHandler)
	e.RunAndListen(config.Instance.Addr)
}
