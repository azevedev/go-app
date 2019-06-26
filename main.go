package main

import "fmt"

func test() int {
	return 10
}
func main() {

	//Funções funcionando

	//simulador_usuarios
	users := criarUsers(10)
	fmt.Println(users)

	randomAtributes(&users)
	fmt.Println(users)

	//simuladorInvestimentos()
	//analisePerfil()

	//Funções em desenvolvimento
	// calculadora()
	// previdencia()

	servidor()

}
