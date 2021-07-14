package app

import (
	"net/http"

	"github.com/al-gas/golang/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/user", controllers.GetUser)
}
