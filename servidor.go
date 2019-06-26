package main

import (
	//"fmt"
	//"net"
	"net/http"
    "log"

    "github.com/zserge/webview"
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
 
 
    go func() {
    http.Handle("/", router)
    log.Fatal(http.ListenAndServe(":8080", router))
    }()
    url := "http://" + "127.0.0.1:8080"
    w := webview.New(webview.Settings{
        Title:     "App_LP",
        Resizable: true,        
        URL:   url,
    })
    defer w.Exit()
    w.Run()
 
}

