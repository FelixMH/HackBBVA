package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/HackBBVA/models"
)

var cliente []models.Clientes

func GetClientesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(cliente)
}