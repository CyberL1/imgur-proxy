package request

type Resource struct {
	Data ResourceData `json:"data"`
	Success bool `json:"success"`
	Status int `json:"status"`
}

type Resources struct {
	Data []ResourceData `json:"data"`
	Success bool `json:"success"`
	Status int `json:"status"`
}

type ResourceData struct {
	ID string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	AccountID int `json:"account_id,omitempty"`
	Link string `json:"link,omitempty"`
	Error string `json:"error,omitempty"`
	Request string `json:"request,omitempty"`
	Method string `json:"method,omitempty"`
}

type Account struct {
	Data AccountData `json:"data"`
	Success bool `json:"success"`
	Status int `json:"status"`
}

type AccountData struct {
	ID int `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
	Request string `json:"request,omitempty"`
	Method string `json:"method,omitempty"`
}

type OAuth struct {
	AccessToken string `json:"access_token"`
}