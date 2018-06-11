package utils

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func NewId() string {
	return uuid.UUID.String(uuid.Must(uuid.NewV4()))
}

type CreateResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
	Links   []Link `json:"_links"`
}

type Error struct {
	Message string `json:"message"`
}

type Link struct {
	Instance string `json:"instance"`
}

func BuildCreateResponse(id string, instanceType string) CreateResponse {
	//TODO: add http uri of current server url as link
	link := fmt.Sprintf("/%s/%s", instanceType, id)
	return CreateResponse{Message: "Created", Id: id, Links: []Link{Link{Instance: link}}}
}

func SendStatusCreated(w http.ResponseWriter, r *http.Request, response CreateResponse) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json, _ := json.Marshal(response)
	w.Write(json)
}

func SendResult(w http.ResponseWriter, r *http.Request, resultJSON []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(resultJSON)
}

func SendStatusOK(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func SendNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	json, _ := json.Marshal(Error{Message: "Not found"})
	w.Write(json)
}

func SendErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json, _ := json.Marshal(Error{Message: err.Error()})
	w.Write(json)
}
