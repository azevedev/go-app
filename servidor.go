package main

import (
	//"fmt"
	//"net"
	"net/http"

    "github.com/gorilla/mux"
)
 
 
func servidor() {

	router := mux.NewRouter()
    router.HandleFunc("/", HomePageHandler) // GET
 
    router.HandleFunc("/Previdencia", PrevidenciaHandler)
 
    router.HandleFunc("/Investimentos", InvestimentosPageHandler).Methods("GET")
    router.HandleFunc("/Investimentos", InvestimentosHandler).Methods("POST")

    router.HandleFunc("/Custos", CustosPageHandler).Methods("GET")
    router.HandleFunc("/Custos", CustosHandler).Methods("POST")
 
 
    http.Handle("/", router)
    http.ListenAndServe(":8080", router)
 
}