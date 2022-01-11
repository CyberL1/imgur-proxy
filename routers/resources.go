package routers

import (
	"fmt"
	"imgurproxy/jsonerror"
	"imgurproxy/proxy"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ResourcesRouters(r *gin.Engine) {
	r.GET("/:code", func(c *gin.Context) {
		codeWithExt := strings.Split(c.Param("code"), ".")
		code := codeWithExt[0]
		var ext string
		if len(codeWithExt) != 1 { // Check if code includes extenstion
			ext = mime.TypeByExtension("." + codeWithExt[1])
		} else {
			ext = proxy.GetResourceByCode(code).Data.Type
		}

		userFetch := proxy.GetAccount()
		resourceFetch := proxy.GetResourceByCodeRaw(code)
		resourceFetchBody := proxy.GetResourceByCode(code)

		if resourceFetch.StatusCode == 404 {
			c.JSON(http.StatusNotFound, jsonerror.Response{
				Status: 404,
				Error: "Resource not found",
			})
			return
		}

		if cfg.Proxy.DisablePublicResources && resourceFetchBody.Data.AccountID != userFetch.Data.ID {
			c.JSON(http.StatusForbidden, jsonerror.Response{
				Status: 403,
				Error: "Public resources are disabled",
			})
			return
		}

		if resourceFetch.StatusCode == 403 && !cfg.Proxy.DisablePublicResources {
			c.JSON(http.StatusForbidden, jsonerror.Response{
				Status: 403,
				Error: "clientID token is invalid",
			})
			return
		}

		if resourceFetch.StatusCode == 403 && cfg.Proxy.DisablePublicResources {
			c.HTML(http.StatusForbidden, "invalidToken.html", nil)
			return
		}

		if string(c.Request.URL.Path[len(c.Request.URL.Path)-1]) == "." {
			if cfg.Proxy.EnableResourceStats {
				c.JSON(http.StatusOK, resourceFetchBody)
				return
			} else {
				c.Redirect(302, code)
				return
			}
		}
		
		var protocol string
		if cfg.Imgur.UseHttps {
			protocol = "https:"
			} else {
			protocol = "http:"
		}

		var resourceURL string
		if len(codeWithExt) != 1 {
			resourceURL = fmt.Sprintf("%v//%v/%v.%v", protocol, cfg.Imgur.ImageDomain, code, codeWithExt[1])
		} else {
			resourceURL = fmt.Sprintf("%v//%v/%v.%v", protocol, cfg.Imgur.ImageDomain, code, strings.Split(resourceFetchBody.Data.Type, "/")[1])
		}

		resourceData, _ := http.Get(resourceURL)
		resourceBody, _ := ioutil.ReadAll(resourceData.Body)
		c.Header("Content-Type", ext)
		c.String(http.StatusOK, string(resourceBody))
	})
}