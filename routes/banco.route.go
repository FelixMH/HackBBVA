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

	// guardar := models.Bancos{
	// 	IdBanco: params["idBanco"],
	// 	Nombre: params["Nombre"],
	// 	IdSeguridad: &models.SeguridadStruct{
	// 		IdSeguridad: params["idSeguridad"],
	// 		CodigoPaquete: params["CodigoPaquete"],
	// 	},
	// 	IdCliente: &models.Clientes{
	// 		IdCliente: params["idCliente"],
	// 	},
	// 	IdTransaccion: &models.Transacciones{
	// 		IdTransaccion: params["idTransaccion"],
	// 	},
	// }

	// err := bd.InsertBancoBD(guardar)
	// if err != nil {
	// 	http.Error(w,"Ocurrió un error al intentar enviar el registro.", 400)
	// 	return
	// }
	Bancos.IdBanco = params["idBanco"]
	banco = append(banco, Bancos)
	json.NewEncoder(w).Encode(banco)
	
	w.WriteHeader(http.StatusCreated)


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