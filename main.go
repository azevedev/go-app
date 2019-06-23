package main

import "fmt"

func test() int{
	return 10
}
func main(){
	varInit();
	fmt.Println("Main!")
 	
	fmt.Println("Calculo: ", calcula_salario_aposentadoria(65, 45, 190))



	//Funções funcionando 
	 	//simulador_usuarios
	 	users := criar_users(10)
		fmt.Println(users)

		random_atributes(&users)
		fmt.Println(users)


	//Funções em desenvolvimento
		// calculadora()
		// previdencia()

	
}