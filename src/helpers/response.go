package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res *Response) ResponseJSON(w http.ResponseWriter, status int, message string, data interface{}, err error) {

	switch status {
	case 200:
		res.Status = status
		res.Message = message
		res.Data = data
		w.WriteHeader(http.StatusOK)
	case 201:
		res.Status = status
		res.Message = message
		res.Data = data
		w.WriteHeader(http.StatusCreated)
	case 204:
		res.Status = status
		res.Message = message
		res.Data = data
		w.WriteHeader(http.StatusNoContent)
	case 300:
		res.Status = status
		res.Message = message
		res.Data = data
		w.WriteHeader(http.StatusMultipleChoices)
	case 304:
		res.Status = status
		res.Message = message
		res.Data = data
		w.WriteHeader(http.StatusNotModified)
	case 400:
		http.Error(w, "", http.StatusBadRequest)
		res.Status = status
		res.Message = err.Error()
		res.Data = data
	case 401:
		http.Error(w, "", http.StatusUnauthorized)
		res.Status = status
		res.Message = err.Error()
		res.Data = data
	case 403:
		http.Error(w, "", http.StatusForbidden)
		res.Status = status
		res.Message = err.Error()
		res.Data = data
	case 404:
		http.Error(w, "", http.StatusNotFound)
		res.Status = status
		res.Message = err.Error()
		res.Data = data
	case 500:
		http.Error(w, "", http.StatusInternalServerError)
		res.Status = status
		res.Message = err.Error()
		res.Data = data
	case 501:
		http.Error(w, "", http.StatusBadGateway)
		res.Status = status
		res.Message = err.Error()
		res.Data = data
	default:
		res.Status = status
		res.Message = ""
		res.Data = data
	}

	json.NewEncoder(w).Encode(res)
}
