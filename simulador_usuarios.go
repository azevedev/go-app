package main

import "time"
import "math/rand"
import "github.com/Pallinder/go-randomdata"


//Cria uma lista de X (qtd) usuários com atributos nulos, depois retorna essa lista.
func criar_users(qtd int) []Usuario{
	users := make([]Usuario, qtd)
	for i := 0; i < qtd; i++ {
		users[i] = user_new()
	}
	return users
}
//Recebe uma lista de usuário e atribui valores aleatórios aos mesmos.
func random_atributes(users *[]Usuario){
	s1 := rand.NewSource(time.Now().UnixNano())
 	r1 := rand.New(s1)
	
	for i := 0; i < len(*users); i++{
		
		(*users)[i].id = -1 //número de controle para diferenciarmos usuários reais de usuários simulados
		(*users)[i].nome = randomdata.FullName(randomdata.RandomGender) //gera um nome aleatório
		(*users)[i].idade = r1.Intn(30) + 18 //gera uma idade aleatória
		(*users)[i].renda = float64(int(((r1.Float64() * 5000) + 100) * 100)) / 100 //código para deixar a 'renda' com 2 casas decimais
		pf := ""
		try := r1.Intn(3)+1
		switch(try){
			case 01:
				pf = "Conservador"
				break
			case 02:
				pf = "Moderado"
				break
			case 03: 
				pf = "Arrojado"
		}
		(*users)[i].perfil_investimento = pf
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
func random_queries(users *[]Usuario){

}