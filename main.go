package main

import "fmt"

func test() int{
	return 10
}
func main(){
	fmt.Println("Main!")
 	




	//Funções funcionando - simulador_usuarios
 	users := criar_users(10)
	fmt.Println(users)

	random_atributes(&users)
	fmt.Println(users)

	calculadora()
	previdencia()

	
}