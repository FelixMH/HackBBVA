package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/HackBBVA/models"
)

var terminal []models.Terminales

func GetTerminalesEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(terminal)
}