package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Method yang berfungsi dengan baik (tidak mengembalikan error)
func DataUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	// contoh slice yang berisikan data users
	users := []map[string]interface{}{
		{"name": "Alice", "age": 30},
		{"name": "Bob", "age": 25},
		{"name": "Charlie", "age": 35},
	}

	// Convert slice data users ke JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	// memberikan response data users
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

// Method yang mengembalikan error
func DataUserError(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	// membuat simulasi error
	panic("Ups")

}
