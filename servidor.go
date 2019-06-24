package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"

	"github.com/zserge/webview"
)

// Person struct que define person
type Person struct {
	Fname string
	Lname string
}

const (
	windowWidth  = 480
	windowHeight = 500
)

func myhandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles("index.html")
		data := struct {
			NomeCompleto string
		}{
			NomeCompleto: "",
		}
		tmpl.Execute(w, data)
	case "POST":
		P := Person{Fname: "Sean", Lname: "50"}
		P.Fname = r.FormValue("firstname")
		P.Lname = r.FormValue("lastname")
		fmt.Println("POST", P.Fname, P.Lname)
		tmpl, _ := template.ParseFiles("index.html")
		data := struct {
			NomeCompleto string
		}{
			NomeCompleto: P.Fname + P.Lname,
		}
		tmpl.Execute(w, data)
	default:
		tmpl, _ := template.ParseFiles("index.html")
		data := struct {
			NomeCompleto string
		}{
			NomeCompleto: "",
		}
		tmpl.Execute(w, data)
	}
}

func startServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", myhandler)
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}

func servidor() {
	url := startServer()
	w := webview.New(webview.Settings{
		Width:     windowWidth,
		Height:    windowHeight,
		Title:     "App_LP",
		Resizable: true,
		URL:       url,
	})
	w.SetColor(255, 255, 255, 255)
	defer w.Exit()
	w.Run()
}
