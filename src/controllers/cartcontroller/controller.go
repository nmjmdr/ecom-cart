package cartcontroller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"models"
	"net/http"
	"promocache"
	"promocalc"
	"repository"
	"utils"
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

type PromofyRequest struct {
	// ids of promos to apply
	Promos []string `json:"promos"`
}

func parsePromofyRequest(request *http.Request) (error, PromofyRequest) {
	decoder := json.NewDecoder(request.Body)
	var req PromofyRequest
	err := decoder.Decode(&req)
	// TODO: Ideally should validate the JSON here
	return err, req
}

func ApplyPromos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["cartId"]
	if len(id) == 0 {
		utils.SendErrorResponse(w, r, errors.New("Invalid cart id passed"))
		return
	}
	err, promofyRequest := parsePromofyRequest(r)
	if err != nil {
		utils.SendErrorResponse(w, r, err)
		return
	}

	cart, ok := repository.Get().Get(id)
	if !ok {
		utils.SendNotFound(w, r)
		return
	}
	promos := make([]models.Promo, 0)
	for _, promoId := range promofyRequest.Promos {
		promo, ok := promocache.GetPromoCache().Get(promoId)
		if !ok {
			utils.SendErrorResponse(w, r, errors.New("Non existant promo id passed"))
			return
		}
		promos = append(promos, promo)
	}

	var calculator = promocalc.NewCalculator()
	promofiedCart := calculator.ApplyPromos(promos, &cart)
	json, _ := json.Marshal(promofiedCart)
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
