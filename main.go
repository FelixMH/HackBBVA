package main

import (
	"encoding/json"

	"log"

	"net/http"

	"github.com/FelixMH/HackBBVA/models"
	"github.com/gorilla/mux"
)





var banco []models.Bancos
var cliente []models.Clientes
var seguridad []models.SeguridadStruct
var terminal []models.Terminales
var transaccion []models.Transacciones

func getBancosEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
	json.NewEncoder(w).Encode(banco)
}

func getClientesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(cliente)
}

func getSeguridadStructEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(seguridad)
}

func getTerminalesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(terminal)
}

func getTransaccionesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(transaccion)
}

func getBancoEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
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

func getClienteEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
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

func getSeguridadEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
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

func getTerminalEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
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
	json.NewEncoder(w).Encode(transaccion)
}

func getTransaccionEndpoint(w http.ResponseWriter, req *http.Request) { //devuelve algo al navegador, deevuelve algo al servidor
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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/banco", getBancosEndpoint).Methods("GET")
	router.HandleFunc("/cliente", getClientesEndpoint).Methods("GET")
	router.HandleFunc("/seguridad", getSeguridadStructEndpoint).Methods("GET")
	router.HandleFunc("/terminal", getTerminalEndpoint).Methods("GET")
	router.HandleFunc("/transaccion", getTransaccionesEndpoint).Methods("GET")

	router.HandleFunc("/banco/{idBanco}", getBancoEndpoint).Methods("GET")
	router.HandleFunc("/banco/{idBanco}", CreateBancoEndpoint).Methods("POST")
	router.HandleFunc("/banco/{idBanco}", DeleteBancoEndpoint).Methods("DELETE")

	router.HandleFunc("/cliente/{idCliente}", getClienteEndpoint).Methods("GET")
	router.HandleFunc("/cliente/{idClienet}", CreateClienteEndpoint).Methods("POST")
	router.HandleFunc("/cliente/{idCliente}", DeleteClienteEndpoint).Methods("DELETE")

	router.HandleFunc("/seguridad/{idSeguridad}", getSeguridadEndpoint).Methods("GET")
	router.HandleFunc("/seguridad/{idSeguridad}", CreateSeguridadEndpoint).Methods("POST")
	router.HandleFunc("/seguridad/{idSeguridad}", DeleteSeguridadEndpoint).Methods("DELETE")

	router.HandleFunc("/terminal/{idTerminal}", getTerminalEndpoint).Methods("GET")
	router.HandleFunc("/terminal/{idTerminal}", CreateTerminalEndpoint).Methods("POST")
	router.HandleFunc("/terminal/{idTerminal}", DeleteTerminalEndpoint).Methods("DELETE")

	router.HandleFunc("/transaccion/{idTransaccion}", getTransaccionEndpoint).Methods("GET")
	router.HandleFunc("/transaccion/{idTransaccion}", CreateTransaccionEndpoint).Methods("POST")
	router.HandleFunc("/transaccion/{idTransaccion}", DeleteTransaccionEndpoint).Methods("DELETE")

	http.ListenAndServe(":3000", router)

	log.Fatal(http.ListenAndServe(":3000", router))
}
