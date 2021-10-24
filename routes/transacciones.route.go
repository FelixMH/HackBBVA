package routes

import (
	"encoding/json"
	"net/http"
)

func GetTransaccionesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(transaccion)
}