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
		(*users)[i].expec_idade_aposent = r1.Intn(80) + 18
		(*users)[i].expec_contribuicao = r1.Float64() * 5000
		(*users)[i].expec_contribuicao = r1.Float64() * 3000
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
		for ii := 0; ii < len(*users); ii++ {
			fmt.Printf("i = %d\n", ii)
			fmt.Printf("len = %d\n", len(*users))

			r := r1.Intn(11)
			uid := ii + r
			switch r {
			case 00:
				arg1 := r1.Float64() * 1000
				arg2 := r1.Float64() * 100.0
				r := poupanca(arg1, arg2)
				fmt.Printf("Simulated 'poupanca' with user_id: %d| result: %f\n", uid, r)
				break
			case 01:

				arg1 := r1.Float64() * 1000
				arg2 := r1.Float64() * 100.0
				r := tesouroPrefixado(arg1, arg2)
				fmt.Printf("Simulated 'tesouroPrefixado' with user_id: %d| result: %f\n", uid, r)

				break
			case 02:

				arg1 := r1.Float64() * 1000
				arg2 := r1.Float64() * 100.0
				r := tesouroSelic(arg1, arg2)
				fmt.Printf("Simulated 'tesouroSelic' with user_id: %d| result: %f\n", uid, r)

				break
			case 03:

				arg1 := r1.Float64() * 1000
				arg2 := r1.Float64() * 100.0
				r := tesouroIPCA(arg1, arg2)
				fmt.Printf("Simulated 'tesouroIPCA' with user_id: %d| result: %f\n", uid, r)

				break
			case 04:

				arg1 := r1.Float64() * 1000
				arg2 := r1.Float64() * 100.0
				r := CDBeLC(arg1, arg2)
				fmt.Printf("Simulated 'CDBeLC' with user_id: %d| result: %f\n", uid, r)

				break
			case 05:

				arg1 := r1.Float64() * 1000
				arg2 := r1.Float64() * 100.0
				r := LCIeLCA(arg1, arg2)
				fmt.Printf("Simulated 'LCIeLCA' with user_id: %d| result: %f\n", uid, r)

				break
			case 06:

				fmt.Printf("Simulated 'analisePerfil' with user_id: %d|\n", uid)

				break
			case 07:

				r := r1.Float64() * 100
				fmt.Printf("Simulated 'CalculadoraAposentadoria' with user_id: %d| result: %f\n", uid, r)

				break
			case 8:

				r := r1.Float64() * 100
				fmt.Printf("Simulated 'CalculadoraAposentadoria' with user_id: %d| result: %f\n", uid, r)

				break
			case 9:

				r := r1.Float64() * 100
				fmt.Printf("Simulated 'CalculadoraContribuicao' with user_id: %d| result: %f\n", uid, r)

				break
			case 10:

				fmt.Printf("Simulated 'custos' with user_id: %d| result: %f\n", uid, r)

				break
			default:

				fmt.Printf("Simulation failed with user_id: \n", uid)

				break

			}

		}
	}
}
