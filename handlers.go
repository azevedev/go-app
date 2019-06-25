package main
 
import (
    "fmt"
    "strconv"
    "net/http"
    "html/template"

)
type Page struct {
    Title string
}

type Rstruct struct {
    Resultado string
}


// Home
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("teste")
    tmpl, _ := template.ParseFiles("templates/Home.html")
    tmpl.Execute(w, &Page{Title: "Home"})
}

//Previdencia

func PrevidenciaHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        tmpl, _ := template.ParseFiles("templates/Previdencia.html")
        tmpl.Execute(w, " ")
    } else {
        fmt.Println("teste")
        Usr := Usuario{0, 0.0, 0, "nome", "perfil", 0, 0.0, 0.0}
        Usr.idade, _ = strconv.Atoi(r.FormValue("idade"))
        Usr.expec_idade_aposent, _ = strconv.Atoi(r.FormValue("espectativa"))
        Usr.expec_salario, _ = strconv.ParseFloat(r.FormValue("salario"), 64)
        tmpl, _ := template.ParseFiles("templates/Previdencia.html")
        data := Rstruct{"voce nunca vai aposentar"}
        tmpl.Execute(w, data)
    }
}

//Investimentos
func InvestimentosPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/Investimentos.html")
    tmpl.Execute(w, &Page{Title: "Investimentos"})
}

func InvestimentosHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/Investimentos.html")
    tmpl.Execute(w, &Page{Title: "Investimentos"})
}

//Custos
func CustosPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/Custos.html")
    tmpl.Execute(w, &Page{Title: "Custos"})
}

func CustosHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/Custos.html")
    tmpl.Execute(w, &Page{Title: "Custos"})
}
 
