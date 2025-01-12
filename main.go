package main

import (
	"fmt"
	"net/http"

	"github.com/MuhammadSyaugi13/simulasi-error/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/api/users", controllers.DataUser)
	router.GET("/api/data-error", controllers.DataError)
	router.GET("/api/data-custom-error", controllers.DataCustomError)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Menjalankan Server......")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
