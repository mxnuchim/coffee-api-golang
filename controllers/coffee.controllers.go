package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mxnuchim/coffee-api-golang/services"
	"github.com/mxnuchim/coffee-api-golang/utils"
)


var coffee services.Coffee

// GET/coffees
func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCoffees()

	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelop{"coffees": all})
}

// GET/coffees/coffee/{id}
func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    coffee, err := coffee.GetCoffeeById(id)
    if err != nil {
        utils.MessageLogs.ErrorLog.Println(err)
        return
    }
    utils.WriteJSON(w, http.StatusOK, coffee)
}

//POST/coffees/coffee
func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)

	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}

	coffeeCreated, err := coffee.CreateCoffee(coffeeData)

	if err != nil {
		utils.MessageLogs.ErrorLog.Println(err)
		return
	}
	
	utils.WriteJSON(w, http.StatusOK, coffeeCreated)
}

// PUT/coffees/coffee/{id}
func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
    var coffeeData services.Coffee
    id := chi.URLParam(r, "id")
    err := json.NewDecoder(r.Body).Decode(&coffeeData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    coffeUpdated, err := coffee.UpdateCoffee(id, coffeeData)
    if err != nil {
        utils.MessageLogs.ErrorLog.Println(err)
    }
    utils.WriteJSON(w, http.StatusOK, coffeUpdated)
}

// DELETE/coffees/coffee/{id}
func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    err := coffee.DeleteCoffee(id)
    if err != nil {
       utils.MessageLogs.ErrorLog.Println(err) 
    }
    utils.WriteJSON(w, http.StatusOK, utils.Envelop{"message": "coffee deleted successfully"})
}