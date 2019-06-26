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
        tmpl, _ := template.ParseFiles("templates/Previdencia.html")
        tmpl.Execute(w, " ")
    }
}

func PrevidenciaRecHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        tmpl, _ := template.ParseFiles("templates/PrevidenciaReceber.html")
        tmpl.Execute(w, " ")
    } else {
        Usr := Usuario{0, 0.0, 0, "nome", "perfil", 0, 0.0, 0.0}
        Usr.idade, _ = strconv.Atoi(r.FormValue("idade"))
        Usr.expec_idade_aposent, _ = strconv.Atoi(r.FormValue("idade_aposent"))
        Usr.expec_contribuicao, _ = strconv.ParseFloat(r.FormValue("contribuir"), 64)
        tmpl, _ := template.ParseFiles("templates/PrevidenciaReceber.html")
        str := fmt.Sprintf("%.2f R$", CalculadoraAposentadoria(Usr))
        data := Rstruct{str}
        tmpl.Execute(w, data)
    }
}

func PrevidenciaContHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        tmpl, _ := template.ParseFiles("templates/PrevidenciaContribuir.html")
        tmpl.Execute(w, " ")
    } else {
        Usr := Usuario{0, 0.0, 0, "nome", "perfil", 0, 0.0, 0.0}
        Usr.idade, _ = strconv.Atoi(r.FormValue("idade"))
        Usr.expec_idade_aposent, _ = strconv.Atoi(r.FormValue("idade_aposent"))
        Usr.expec_salario, _ = strconv.ParseFloat(r.FormValue("receber"), 64)
        tmpl, _ := template.ParseFiles("templates/PrevidenciaContribuir.html")
        str := fmt.Sprintf("%.2f R$", CalculadoraContribuicao(Usr))
        data := Rstruct{str}
        tmpl.Execute(w, data)
    }
}

//Investimentos
func InvestimentosHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        tmpl, _ := template.ParseFiles("templates/Investimentos.html")
        tmpl.Execute(w, " ")
    } else {
        tmpl, _ := template.ParseFiles("templates/Investimentos.html")
        tmpl.Execute(w, " ")
    }
}

func InvestimentosPerfilHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        tmpl, _ := template.ParseFiles("templates/InvestimentosPerfil.html")
        tmpl.Execute(w, " ")
    } else {
        Usr := Usuario{0, 0.0, 0, "nome", "perfil", 0, 0.0, 0.0}
        Usr.idade, _ = strconv.Atoi(r.FormValue("idade"))
        Usr.expec_idade_aposent, _ = strconv.Atoi(r.FormValue("idade_aposent"))
        Usr.expec_contribuicao, _ = strconv.ParseFloat(r.FormValue("contribuir"), 64)
        tmpl, _ := template.ParseFiles("templates/InvestimentosPerfil.html")
        str := fmt.Sprintf("%.2f R$", CalculadoraAposentadoria(Usr))
        data := Rstruct{str}
        tmpl.Execute(w, data)
    }
}

func InvestimentosRendHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        tmpl, _ := template.ParseFiles("templates/InvestimentosRendimento.html")
        tmpl.Execute(w, " ")
    } else {
        Usr := Usuario{0, 0.0, 0, "nome", "perfil", 0, 0.0, 0.0}
        Usr.idade, _ = strconv.Atoi(r.FormValue("idade"))
        Usr.expec_idade_aposent, _ = strconv.Atoi(r.FormValue("idade_aposent"))
        Usr.expec_contribuicao, _ = strconv.ParseFloat(r.FormValue("contribuir"), 64)
        tmpl, _ := template.ParseFiles("templates/InvestimentosRendimento.html")
        str := fmt.Sprintf("%.2f R$", CalculadoraAposentadoria(Usr))
        data := Rstruct{str}
        tmpl.Execute(w, data)
    }
}

//Custos
func CustosHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/Custos.html")
    tmpl.Execute(w, &Page{Title: "Custos"})
}
 
