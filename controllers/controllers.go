package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexferreira1796/api-go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá mundo")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	// Codifica os dados de personalidades em JSON e escreve na resposta HTTP
	json.NewEncoder(w).Encode(models.PersonalitiesData)
}