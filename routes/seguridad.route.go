package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/HackBBVA/models"
	"github.com/gorilla/mux"
)

var seguridad []models.SeguridadStruct

func GetSeguridadStructEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(seguridad)
}

func GetSeguridadEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for _, item := range seguridad {
		if item.IdSeguridad == params["idSeguridad"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&seguridad)
}
func CreateSeguridadEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	var Seguridads models.SeguridadStruct
	_ = json.NewDecoder(req.Body).Decode(&Seguridads)
	Seguridads.IdSeguridad = params["idSeguridad"]
	seguridad = append(seguridad, Seguridads)
	json.NewEncoder(w).Encode(Seguridads)
}
func DeleteSeguridadEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for index, item := range seguridad {
		if item.IdSeguridad == params["idSeguridad"] {
			seguridad = append(seguridad[:index], seguridad[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(banco)
}