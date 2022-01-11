package config

type Config struct {
	Imgur Imgur `json:"imgur"`
	Proxy Proxy `json:"proxy"`
}

type Imgur struct {
	UseHttps bool `json:"useHttps"`
	ClientID string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	RefreshToken string `json:"refreshToken"`
	ApiDomain string `json:"apiDomain"`
	ApiVersion int `json:"apiVersion"`
	ImageDomain string `json:"imageDomain"`
	AlbumID string `json:"albumID"`
}


type Proxy struct {
	Port string `json:"port"`
	UseMediaHub bool `json:"useMediaHub"`
	UseAlbum bool `json:"useAlbum"`
	EnableResourceStats bool `json:"enableResourceStats"`
	DisablePublicResources bool `json:"disablePublicResources"`
}