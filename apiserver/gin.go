package apiserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toolapi/dao"
)

func InitGinServer(bundle *dao.OptBundle) *http.Server {
	// Create gin http server.
	gin.SetMode("release")
	router := gin.Default()
	router.Use(gin.Recovery())

	InitSdpApi(router, bundle)
	InitShareMapApi(router, bundle)
	InitRedirectApi(router, bundle)
	port := "8889"
	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
	}

	return httpServer
}

