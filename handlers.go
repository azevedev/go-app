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
        q1, _ := strconv.Atoi(r.FormValue("questao1"))
        q2, _ := strconv.Atoi(r.FormValue("questao2"))
        q3, _ := strconv.Atoi(r.FormValue("questao3"))
        q4, _ := strconv.Atoi(r.FormValue("questao4"))
        q5, _ := strconv.Atoi(r.FormValue("questao5"))
        q6, _ := strconv.Atoi(r.FormValue("questao6"))
        SomaPerfil := q1+q2+q3+q4+q5+q6
        tmpl, _ := template.ParseFiles("templates/InvestimentosPerfil.html")
        str := analisePerfil(SomaPerfil)
        data := Rstruct{str}
        tmpl.Execute(w, data)
    }
}


func InvestimentosRendHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        tmpl, _ := template.ParseFiles("templates/InvestimentosRendimento.html")
        tmpl.Execute(w, " ")
    } else {
        Deposito, _ := strconv.ParseFloat(r.FormValue("deposito"), 64)
        Meses, _ := strconv.ParseFloat(r.FormValue("meses"), 64)
        tmpl, _ := template.ParseFiles("templates/InvestimentosRendimento.html")
        r1, r2, r3, r4, r5, r6 := Rendimentos(Deposito, Meses)
        data := struct{ Resultado1 string
                        Resultado2 string
                        Resultado3 string
                        Resultado4 string
                        Resultado5 string
                        Resultado6 string
                    }{r1, r2, r3, r4, r5, r6}
        tmpl.Execute(w, data)
    }
}

//Custos
func CustosHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/Custos.html")
    tmpl.Execute(w, &Page{Title: "Custos"})
}
 
