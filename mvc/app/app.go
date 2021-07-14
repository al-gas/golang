package app

import (
	"net/http"

	"github.com/al-gas/golang/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/user", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
