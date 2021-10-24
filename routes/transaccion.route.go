package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/HackBBVA/models"
	"github.com/gorilla/mux"
)

var transaccion []models.Transacciones

func GetTransaccionEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for _, item := range transaccion {
		if item.IdTransaccion == params["idTransaccion"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&transaccion)
}

func CreateTransaccionEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	var transacciones models.Transacciones
	_ = json.NewDecoder(req.Body).Decode(&transacciones)
	transacciones.IdTransaccion = params["idTransaccion"]
	transaccion = append(transaccion, transacciones)
	json.NewEncoder(w).Encode(transaccion)
}

func DeleteTransaccionEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for index, item := range transaccion {
		if item.IdTransaccion == params["idTransaccion"] {
			transaccion = append(transaccion[:index], transaccion[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(banco)
}