package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/HackBBVA/models"
	"github.com/gorilla/mux"
)

func GetTerminalEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for _, item := range terminal {
		if item.IdTerminal == params["idTerminal"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&terminal)
}
func CreateTerminalEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	var terminales models.Terminales
	_ = json.NewDecoder(req.Body).Decode(&terminales)
	terminales.IdTerminal = params["idBanco"]
	terminal = append(terminal, terminales)
	json.NewEncoder(w).Encode(terminal)
}
func DeleteTerminalEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	params := mux.Vars(req)
	for index, item := range terminal {
		if item.IdTerminal == params["idTerminal"] {
			terminal = append(terminal[:index], terminal[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(terminal)
}