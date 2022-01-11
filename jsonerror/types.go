package jsonerror

type Response struct {
	Status int `json:"status"`
	Error string `json:"error"`
}