package routers

import (
	"imgurproxy/config"
	"imgurproxy/jsonerror"
	"imgurproxy/proxy"
	"imgurproxy/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

var cfg = config.GetConfig()

func PagesRouters(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		if cfg.Proxy.UseMediaHub {
			resourcesFetch := proxy.GetAllResourcesRaw()
			resourcesFetchBody := proxy.GetAllResources()
			albumInfo := proxy.GetAlbum()

			if cfg.Proxy.UseAlbum {
				if resourcesFetch.StatusCode == 403 {
					c.JSON(http.StatusForbidden, jsonerror.Response{
						Status: 403,
						Error: "clientID is invalid",
					})
					return
				}

				if albumInfo.StatusCode == 400 {
					c.JSON(http.StatusBadRequest, jsonerror.Response{
						Status: 400,
						Error: "albumID is required",
					})
					return
				}

				if albumInfo.StatusCode == 404 {
					c.JSON(http.StatusForbidden, jsonerror.Response{
						Status: 404,
						Error: "albumID is invalid",
					})
					return
				}
			}

			if resourcesFetch.StatusCode == 403 {
				c.HTML(http.StatusForbidden, "invalidToken.html", nil)
				return
			}

			c.HTML(http.StatusOK, "mediaHub.html", gin.H{
				"Resources": resourcesFetchBody,
			})
		} else {
			c.HTML(http.StatusOK, "index.html", nil)
		}
	})

	r.GET("/_renew", func(c *gin.Context) {
		token.GetNew()
	})
}