package proxy

import (
	"imgurproxy/config"
	"imgurproxy/request"
	"imgurproxy/token"
	"net/http"
)

var cfg = config.GetConfig()

func GetResourceByCodeRaw(code string) (*http.Response) {
	var authMethod string
	if cfg.Proxy.DisablePublicResources {
		authMethod = "Bearer " + token.GetToken()
	} else {
		authMethod = "Client-ID " + cfg.Imgur.ClientID
	}
	return request.Make("image/" + code, authMethod)
}

func GetResourceByCode(code string) request.Resource {
	var authMethod string
	if cfg.Proxy.DisablePublicResources {
		authMethod = "Bearer " + token.GetToken()
	} else {
		authMethod = "Client-ID " + cfg.Imgur.ClientID
	}
	return request.GetResource("image/" + code, authMethod)
}

func GetAllResourcesRaw() (*http.Response) {
	var endpoint string
	var authMethod string
	if cfg.Proxy.UseAlbum {
		endpoint = "album/" + cfg.Imgur.AlbumID + "/images"
		authMethod = "Client-ID " + cfg.Imgur.ClientID
	} else {
		endpoint = "account/me/images"
		authMethod = "Bearer " + token.GetToken()
	}
	return request.Make(endpoint, authMethod)
}

func GetAllResources() request.Resources {
	var endpoint string
	var authMethod string
	if cfg.Proxy.UseAlbum {
		endpoint = "album/" + cfg.Imgur.AlbumID + "/images"
		authMethod = "Client-ID " + cfg.Imgur.ClientID
	} else {
		endpoint = "account/me/images"
		authMethod = "Bearer " + token.GetToken()
	}
	return request.GetResourcesArray(endpoint, authMethod)
}

func GetAlbum() (*http.Response) {
	return request.Make("album/" + cfg.Imgur.AlbumID, "Client-ID " + cfg.Imgur.ClientID)
}

func GetAccount() request.Account {
	return request.GetAccount("account/me", "Bearer " + token.GetToken())
}