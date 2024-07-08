package routes

import (
	"net/http"

	"github.com/alexferreira1796/api-go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/personalidades", controllers.GetAllPersonalities)
}