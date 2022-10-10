package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code        int         `json:"-"`
	Status      string      `json:"status"`
	IsError     bool        `json:"isError"`
	Data        interface{} `json:"data,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

func (res *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Origin", "https://govehiclerental.netlify.app")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	if res.IsError {
		w.WriteHeader(res.Code)
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("Error When Encode respone"))
	}
}

func New(data interface{}, code int, isError bool) *Response {
	if isError {
		return &Response{
			Code:        code,
			Status:      getStatus(code),
			IsError:     isError,
			Description: data,
		}
	}
	return &Response{
		Code:    code,
		Status:  getStatus(code),
		IsError: isError,
		Data:    data,
	}
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 404:
		desc = "Not Found"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = "Undefined"
	}

	return desc
}
