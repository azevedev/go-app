package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
)

//Cria uma lista de X (qtd) usuários com atributos nulos, depois retorna essa lista.
func criarUsers(qtd int) []Usuario {
	users := make([]Usuario, qtd)
	for i := 0; i < qtd; i++ {
		users[i] = userNew()
	}
	return users
}

//Recebe uma lista de usuário e atribui valores aleatórios aos mesmos.
func randomAtributes(users *[]Usuario) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < len(*users); i++ {

		(*users)[i].id = -1                                                   //número de controle para diferenciarmos usuários reais de usuários simulados
		(*users)[i].nome = randomdata.FullName(randomdata.RandomGender)       //gera um nome aleatório
		(*users)[i].idade = r1.Intn(30) + 18                                  //gera uma idade aleatória
		(*users)[i].renda = float64(int(((r1.Float64()*5000)+100)*100)) / 100 //código para deixar a 'renda' com 2 casas decimais
		pf := ""
		try := r1.Intn(3) + 1
		switch try {
		case 01:
			pf = "Conservador"
			break
		case 02:
			pf = "Moderado"
			break
		case 03:
			pf = "Arrojado"
		}
		(*users)[i].perfilInvestimento = pf
		// Prints para teste de funcionamento
		// fmt.Print("i:",i)
		// fmt.Print(" nome: ",(*users)[i].nome)
		// fmt.Print(" idade: ",(*users)[i].idade)
		// fmt.Print(" renda: ",(*users)[i].renda)
		// fmt.Print(" perfil: ",(*users)[i].perfil_investimento)
		// fmt.Println()
	}

}

//Recebe uma lista de
func randomQueries(users *[]Usuario) {
	for {
		time.Sleep(800 * time.Millisecond)
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		for i := 0; i < len(*users); i++ {
			r := r1.Intn(11)
			uid := i + r
			switch r {
			case 00:
				go func() {
					fmt.Println("Simulated 'poupanca' with user_id: ", uid)
				}()
				break
			case 01:
				go func() {
					fmt.Println("Simulated 'tesouroPrefixado' with user_id: ", uid)
				}()

				break
			case 02:
				go func() {
					fmt.Println("Simulated 'tesouroSelic' with user_id: ", uid)
				}()
				break
			case 03:
				go func() {
					fmt.Println("Simulated 'tesouroIPCA' with user_id: ", uid)
				}()

				break
			case 04:
				go func() {
					fmt.Println("Simulated 'CDBeLC' with user_id: ", uid)
				}()
				break
			case 05:
				go func() {
					fmt.Println("Simulated 'LCIeLCA' with user_id: ", uid)
				}()
				break
			case 06:
				go func() {
					fmt.Println("Simulated 'analisePerfil' with user_id: ", uid)
				}()
				break
			case 07:
				go func() {
					fmt.Println("Simulated 'CalculadoraAposentadoria' with user_id: ", uid)
				}()
				break
			case 8:
				go func() {
					fmt.Println("Simulated 'CalculadoraAposentadoria' with user_id: ", uid)
				}()
				break
			case 9:
				go func() {
					fmt.Println("Simulated 'CalculadoraContribuicao' with user_id: ", uid)
				}()
				break
			case 10:
				go func() {
					fmt.Println("Simulated 'custos' with user_id: ", uid)
				}()
				break
			default:
				go func() {
					fmt.Println("Simulation failed with user_id: ", uid)
				}()
				break

			}

		}
	}
}
