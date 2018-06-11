package cartcontroller

import (
	"encoding/json"
	"models"
	"net/http"
  "utils"
  "repository"
  "github.com/gorilla/mux"
  "errors"
)


func parseCartJSON(request *http.Request) (error, models.Cart) {
	decoder := json.NewDecoder(request.Body)
	var cart models.Cart
	err := decoder.Decode(&cart)
	// TODO: Ideally should validate the JSON here
	return err, cart
}

func Create(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	err, cart := parseCartJSON(r)
	if err != nil {
		utils.SendErrorResponse(w, r, err)
		return
	}
  cart.Id = utils.NewId()
  repository.Get().Add(cart)
	utils.SendStatusCreated(w, r, utils.BuildCreateResponse(cart.Id, "carts"))
}

func Get(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  if len(id) == 0 {
    utils.SendErrorResponse(w, r, errors.New("Invalid id passed"))
    return
  }
  cart, ok := repository.Get().Get(id)
  if !ok {
    utils.SendNotFound(w, r)
    return
  }
  json, _ := json.Marshal(cart)
  utils.SendResult(w, r, json)
}

func Delete(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  if len(id) == 0 {
    utils.SendErrorResponse(w, r, errors.New("Invalid id passed"))
    return
  }
  ok := repository.Get().Delete(id)
  if !ok {
    utils.SendNotFound(w, r)
    return
  }
  utils.SendStatusOK(w, r)
}
