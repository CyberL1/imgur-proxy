package token

import (
	"imgurproxy/config"
	"imgurproxy/request"

	"git.mills.io/prologic/bitcask"
)

func GetNew() {
	db, _ := bitcask.Open("./bearerToken")
	defer db.Close()
	cfg := config.GetConfig()

	req := request.Oauth("token", map[string]string{
		"refresh_token": cfg.Imgur.RefreshToken,
		"client_id": cfg.Imgur.ClientID,
		"client_secret": cfg.Imgur.ClientSecret,
		"grant_type": "refresh_token",
	})

	db.Put([]byte("bearerToken"), []byte(req.AccessToken))
}

func GetToken() string {
	db, _ := bitcask.Open("./bearerToken")
	defer db.Close()

	t, _ := db.Get([]byte("bearerToken"))
	return string(t)
}