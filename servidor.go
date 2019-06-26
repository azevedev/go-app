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
    router.HandleFunc("/PrevidenciaReceber", PrevidenciaRecHandler)
    router.HandleFunc("/PrevidenciaContribuir", PrevidenciaContHandler)
 
    router.HandleFunc("/Investimentos", InvestimentosHandler)
    router.HandleFunc("/InvestimentosPerfil", InvestimentosPerfilHandler)
    router.HandleFunc("/InvestimentosRendimento", InvestimentosRendHandler)

    router.HandleFunc("/Custos", CustosHandler)
 
 
    http.Handle("/", router)
    http.ListenAndServe(":8080", router)
 
}