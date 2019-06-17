package main

import "fmt"

func  calculadora () {
	enca  :=`
	*******************	
	    CALCULADORA
	*******************
	`

	fmt. Println (enca)
	fmt. Println ( `
		1-> Soma		
		2-> Subtração
		3-> multiplicação
		4-> Divisão
		
		` )
	fmt. Println ( " Qual operaçao deseja realizar? " )
	var  ope  int
	fmt. Scanln (& ope)
	var  num  int
	var  num1  int

	// Soma

	if ope == 1 {
		fmt.Println ("Coloque  o primeiro número para a soma: " )
		fmt.Scanln (&num)
		fmt.Println ("Coloque o segundo número para a soma: " )
		fmt.Scanln (&num1)
		resul  :=  `
		***********************
		* resultado da soma *
		***********************	
		`
		resultado  := num + num1
		fmt. Println (resul)
		fmt. Println ( "O resultado da Soma é: " , resultado)

	}
	// Subtração
	if ope == 2 {
		fmt. Println ( " Coloque  o primeiro número para a Subtração: " )
		fmt. Scanln (& num)
		fmt. Println ( " Coloque  o primeiro número para a Subtração: " )
		fmt. Scanln (& num1)
		resul  :=  `
		************************
		* resultado de Subtração *
		************************
		`
		resultado  := num - num1
		fmt. Println (resul)
		fmt. Println ( "O resultado da Subtração: " , resultado)

	}
	if ope == 3 {
		fmt. Println ( " Coloque  o primeiro número para Multiplicação: " )
		fmt. Scanln (& num)
		fmt. Println ( " Coloque  o primeiro número para Multiplicación: " )
		fmt. Scanln (& num1)
		resul  :=  `
		********************************
		* resultado da multiplicação *
		********************************
		`
		resultado  := num * num1
		fmt. Println (resul)
		fmt. Println ( "O resultado da multiplicação: " , resultado)

	}
	if ope == 4 {
		fmt. Println ( " Coloque  o primeiro número para Divisão: " )
		fmt. Scanln (& num)
		fmt. Println ( " Coloque  o primeiro número para Divisão: " )
		fmt. Scanln (& num1)
		resul  :=  `
		***************************
		* resultado da Divisão *
		***************************
		`
		resultado  := num / num1
		fmt. Println (resul)
		fmt. Println ( "O resultado da Divisão es: " , resultado)

	}
	// sair
	fmt. Println ( " Presione enter para sair \n " )
	fmt. Scanln ()
	fmt. Println ( " Tchau " )
}