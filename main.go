package main

import (
	"net/http"

	"github.com/alexferreira1796/api-go-rest/models"
	"github.com/alexferreira1796/api-go-rest/routes"
)

func main() {
	models.PersonalitiesData = [] models.Personalities{
		{
			Name: "Nome 2",
			History: "História 2",
		},
		{
			Name: "Nome 2",
			History: "História 2",
		},
	}

	routes.HandleRequest()

	http.ListenAndServe(":8000", nil)
}