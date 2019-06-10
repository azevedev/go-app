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
	s1 := rand.NewSource(time.Now().UnixNano())
 	r1 := rand.New(s1)
	for i := 0; i < len(*users); i++{
		r := r1.Intn(11)
		uid := (*users)[i].id
		switch(r){
		 case 00:
		 	/*
		 	go func(){
		 		Calcula_Salario_aposentadoria(uid)
		 		fmt.Println("Simulated 'Calcula_Salario_aposentadoria' with user_id: ",uid)
		 	}()
		 	*/
		 	break
		 case 01:
			/*
		 	go func(){
		 		Calcula_Idade_aposentadoria(uid)
		 		fmt.Println("Simulated 'Calcula_Idade_aposentadoria' with user_id: ",uid)
		 	}()
		 	*/
			break
		 case 02:
			/*
		 	go func(){
		 		Calcula_Investimento_em_previdencia_mensal(uid)
		 		fmt.Println("Simulated 'Calcula_Investimento_em_previdencia_mensal' with user_id: ",uid)
		 	}()
		 	*/
			break
		 case 03:
		 	/*
		 	go func(){
		 		LCI(uid)
		 		fmt.Println("Simulated 'LCI' with user_id: ",uid)
		 	}()
		 	*/
			break
		 case 04:	
			/*
		 	go func(){
		 		LCA(uid)
		 		fmt.Println("Simulated 'LCA' with user_id: ",uid)
		 	}()
		 	*/
			break
		 case 05:
		 	/*
		 	go func(){
		 		LC(uid)
		 		fmt.Println("Simulated 'LC' with user_id: ",uid)
		 	}()
		 	*/
		 	break
	 	 case 06:
	 	 	/*
		 	go func(){
		 		CDB(uid)
		 		fmt.Println("Simulated 'CDB' with user_id: ",uid)
		 	}()
		 	*/
		 	break
		 case 07:
		 	/*
		 	go func(){
		 		Acoes_Escriturais(uid)
		 		fmt.Println("Simulated 'Acoes_Escriturais' with user_id: ",uid)
		 	}()
		 	*/
		 	break
		 case 8:
		 	/*
		 	go func(){
		 		Rentabilidade_da_carteira(uid)
		 		fmt.Println("Simulated 'Rentabilidade_da_carteira' with user_id: ",uid)
		 	}()
		 	*/
		 	break	
		 case 9:
		 	/*
		 	go func(){
		 		Simulador_de_investimento(uid)
		 		fmt.Println("Simulated 'Simulador_de_investimento' with user_id: ",uid)
		 	}()
		 	*/
		 	break
		 case 10:
		 	/*
		 	go func(){
		 		Calcul_poupanca(uid)
		 		fmt.Println("Simulated 'Calculo_poupanca' with user_id: ",uid)
		 	}()
		 	*/
		 	break
		 default:
		 	/*
		 	go func(){
		 		fmt.Println("Simulation failed with user_id: ",uid)
		 	}()
		 	*/
		 	break


	
}
		 
		}
}