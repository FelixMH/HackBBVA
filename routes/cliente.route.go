package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/HackBBVA/models"
	"github.com/gorilla/mux"
)

func GetClienteEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for _, item := range cliente {
		if item.IdCliente == params["idCliente"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&cliente)
}

func CreateClienteEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	var Clientes models.Clientes
	_ = json.NewDecoder(req.Body).Decode(&Clientes)
	Clientes.IdCliente = params["idCliente"]
	cliente = append(cliente, Clientes)
	json.NewEncoder(w).Encode(cliente)
}


func DeleteClienteEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for index, item := range cliente {
		if item.IdCliente == params["idClienet"] {
			cliente = append(cliente[:index], cliente[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cliente)
}