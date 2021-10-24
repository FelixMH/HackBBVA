package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/HackBBVA/models"
	"github.com/gorilla/mux"
)

//var banco []models.Bancos

func GetBancoEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for _, item := range banco {
		if item.IdBanco == params["idBanco"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&banco)
}


func CreateBancoEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	var Bancos models.Bancos
	_ = json.NewDecoder(req.Body).Decode(&Bancos)
	Bancos.IdBanco = params["idBanco"]
	banco = append(banco, Bancos)
	json.NewEncoder(w).Encode(banco)
}

func DeleteBancoEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for index, item := range banco {
		if item.IdBanco == params["idBanco"] {
			banco = append(banco[:index], banco[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(banco)
}