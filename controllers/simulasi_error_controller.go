package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MuhammadSyaugi13/simulasi-error/helper"
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
func DataError(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	// membuat simulasi error
	panic("Ups")

}

func DataCustomError(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

	result, err := func() (string, error) {
		status := "error"

		if status == "error" {
			return "", errors.New("ups terjadi error")
		}

		return "Berhasil", nil
	}()

	if err != nil {
		helper.HandleIfError(err)

		// Mengembalikan response error
		pesan := []map[string]interface{}{
			{"code": 500},
			{"status": "INTERNAL SERVER ERROR"},
			{"data": result},
		}

		// Convert slice data pesan ke JSON
		jsonData, err := json.Marshal(pesan)
		if err != nil {
			panic(err)
		}

		// memberikan response data users
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonData)
		return
	}

}
