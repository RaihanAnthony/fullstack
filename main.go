package main

import (
	"log"
	"net/http"

)

func main() {

	// middleware loggger
	// router.Handle("/", middleware.logger(http.HandlerFunc(controllers.index)))

	http.HandleFunc("/", controllers.Index)
	port := "8000"
	log.Println("server berjalan pada http://localhost:%\n", port)
	http.ListenAndServe(port, nil)
}
