package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/FelixMH/HackBBVA/routes"
)

func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/banco", routes.GetBancosEndpoint).Methods("GET")
	router.HandleFunc("/cliente", routes.GetClientesEndpoint).Methods("GET")
	router.HandleFunc("/seguridad", routes.GetSeguridadStructEndpoint).Methods("GET")
	router.HandleFunc("/terminal", routes.GetTerminalEndpoint).Methods("GET")
	router.HandleFunc("/transaccion", routes.GetTransaccionesEndpoint).Methods("GET")

	router.HandleFunc("/banco/{idBanco}", routes.GetBancoEndpoint).Methods("GET")
	router.HandleFunc("/banco/{idBanco}", routes.CreateBancoEndpoint).Methods("POST")
	router.HandleFunc("/banco/{idBanco}", routes.DeleteBancoEndpoint).Methods("DELETE")

	router.HandleFunc("/cliente/{idCliente}", routes.GetClienteEndpoint).Methods("GET")
	router.HandleFunc("/cliente/{idClienet}", routes.CreateClienteEndpoint).Methods("POST")
	router.HandleFunc("/cliente/{idCliente}", routes.DeleteClienteEndpoint).Methods("DELETE")

	router.HandleFunc("/seguridad/{idSeguridad}", routes.GetSeguridadEndpoint).Methods("GET")
	router.HandleFunc("/seguridad/{idSeguridad}", routes.CreateSeguridadEndpoint).Methods("POST")
	router.HandleFunc("/seguridad/{idSeguridad}", routes.DeleteSeguridadEndpoint).Methods("DELETE")

	router.HandleFunc("/terminal/{idTerminal}", routes.GetTerminalEndpoint).Methods("GET")
	router.HandleFunc("/terminal/{idTerminal}", routes.CreateTerminalEndpoint).Methods("POST")
	router.HandleFunc("/terminal/{idTerminal}", routes.DeleteTerminalEndpoint).Methods("DELETE")

	router.HandleFunc("/transaccion/{idTransaccion}", routes.GetTransaccionEndpoint).Methods("GET")
	router.HandleFunc("/transaccion/{idTransaccion}", routes.CreateTransaccionEndpoint).Methods("POST")
	router.HandleFunc("/transaccion/{idTransaccion}", routes.DeleteTransaccionEndpoint).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}



	//http.ListenAndServe(":3000", router)

	//log.Fatal(http.ListenAndServe(":3000", router))