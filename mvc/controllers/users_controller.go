package controllers

import (
	"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId := req.URL.Query().Get("user_id")
	log.Printf("About to process user_id %v", userId)
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		return
	}
}
