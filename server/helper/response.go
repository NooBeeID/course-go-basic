package helper

import (
	"encoding/json"
	"fmt"
	apps "go-web-template/server/apps/api"
	"log"
	"net/http"
)

func HandleNotMethodAllowed(w http.ResponseWriter, method string) {
	var response apps.ResponseFail

	msg := fmt.Sprintf("Method %s tidak diperbolehkan", method)
	response.Status = http.StatusMethodNotAllowed
	response.Message = msg

	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
}

func HandleBadRequest(w http.ResponseWriter, err error) {
	log.Println(err)
	var response apps.ResponseFail
	response.Status = http.StatusBadRequest
	response.Message = err.Error()

	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
}
