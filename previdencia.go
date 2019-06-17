package main

import  "fmt"

func  previdencia () {
	enca  :=  `
	
	*******************	
	    Previdencia
	*******************
	`

	fmt. Println (enca) // aqui o programa mostra as opções disponíveis ao usuário
	fmt. Println ( ` Nosso simulador possui as seguintes opções: 
		1-> Quero saber qual o valor que preciso contribuir	
		2-> Quero saber quanto que irei receber por mês		
		` )

	fmt. Println ( "Escolha uma das opções" ) // aqui o programa seleciona qual das duas áreas ele entrara
	var  ope  int
	fmt. Scanln (& ope)
	var  num  int
	var  num1  int

//opção 1 - valor da contribuição

	if ope == 1 { // escolha do usuario na opção 1 apresentado
		fmt. Println ( " Qual a sua idade? " ) // pergunta 1 ao usuario
		fmt. Scanln (& idade_op1) //le e armazena a resposta 1 do usuario
		fmt. Println ( " Quantos anos você quer se aposentar? " ) // pergunta 2 ao usuario
		fmt. Scanln (& anos_op1) //le e armazena a resposta 2 do usuario
		fmt. Println ("Quando você se aposentar, quanto quer receber?") // pergunta 3 ao usuario
		fmt. Scanln (& receber_op1) //le e armazena a resposta 3 do usuario
		
		// Subtração (para saber qual o tempo de contribuição em anos)
			tc_op1  :=  `
			************************
			*tempo_de_contribuicao*
			************************
			`  //ta usando as estrelinhas para atribuir o valor do que ta escrito dentro
			tempo_de_contribuicao_op1  := anos_op1 - idade_op1 // resultado da subtracao do tempo de contribuicao
			fmt. Println (tc_op1)
			fmt. Println ( "O seu tempo de contribuição é:" , tempo_de_contribuicao_op1) // ele printa na tela o tempo de contribuicao

		// Multiplicação (para saber qual o tempo de contribuição em meses)
			
				tcm_op1  :=  `
				********************************
				* tempo_de_contribuicao_meses *
				********************************
				`  //ta usando as estrelinhas para atribuir o valor do que ta escrito dentro
				tempo_de_contribuicao_meses_op1  := tc_op1 * 12 // resultado da multiplicação do tempo em meses
				fmt. Println (tcm_op1)
				fmt. Println ( "O tempo de contribuição em meses é de: " , tempo_de_contribuicao_meses_op1) // ele printa na tela o tempo de contribuicao


		// Divisão (para saber qual o valor ele precisará contribuir)
		vc_op1  :=  `
		***************************
		* valor_da_contribuicao *
		***************************
		`  //ta usando as estrelinhas para atribuir o valor do que ta escrito dentro
		valor_da_contribuicao_op1  := receber_op1 / tempo_de_contribuicao_meses_op1 //divisao do que quer receber pelo tempo de contribuicao em meses
		fmt. Println (vc_op1)
		fmt. Println ( "O valor que precisará contribuir para alcançar seu objetivo é: R$" , valor_da_contribuicao_op1) // ele printa na tela o valor necessário

}

//opção 2 - valor do salário mensal

	if ope == 2 { // escolha do usuario na opção 2 apresentado
		fmt. Println ( " Qual a sua idade? " ) // pergunta 1 ao usuario
		fmt. Scanln (& idade_op2) //le e armazena a resposta 1 do usuario
		fmt. Println ( " Quantos anos você quer se aposentar? " ) // pergunta 2 ao usuario
		fmt. Scanln (& anos_op2) //le e armazena a resposta 2 do usuario
		fmt. Println ("Quanto você quer contribuir?") // pergunta 3 ao usuario
		fmt. Scanln (& contribuir_op2) //le e armazena a resposta 3 do usuario
		fmt. Println ("Quando você se aposentar, até que idade deseja receber renda mensal?") // pergunta 4 ao usuario
		fmt. Scanln (& idosinho_op2) //le e armazena a resposta 4 do usuario
		
		// Subtração (para saber qual o tempo de contribuição em anos)
			tc_op2  :=  `
			************************
			*tempo_de_contribuicao*
			************************
			` //ta usando as estrelinhas para atribuir o valor do que ta escrito dentro
			tempo_de_contribuicao_op2  := anos_op2 - idade_op2 // resultado da subtracao do tempo de contribuicao
			fmt. Println (tc_op2)
			fmt. Println ( "O seu tempo de contribuição é:" , tempo_de_contribuicao_op2) // ele printa na tela o tempo de contribuicao

		// Multiplicação (para saber qual o tempo de contribuição em meses)
			
				tcm_op2  :=  `
				********************************
				* tempo_de_contribuicao_meses *
				********************************
				` //ta usando as estrelinhas para atribuir o valor do que ta escrito dentro
				tempo_de_contribuicao_meses_op2  := tc_op2 * 12 // resultado da multiplicação do tempo em meses
				fmt. Println (tcm_op2)
				fmt. Println ( "O tempo de contribuição em meses é de: " , tempo_de_contribuicao_meses_op2) // ele printa na tela o tempo de contribuicao

		// Multiplicação (para saber qual o valor total em meses)
			
				vtm_op2  :=  `
				********************************
				* valor_total_meses_op2 *
				********************************
				` //ta usando as estrelinhas para atribuir o valor do que ta escrito dentro
				valor_total_meses_op2  := tcm_op2 * contribuir_op2  // resultado da multiplicação valor total em meses para o contribuicao 
				fmt. Println (tcm_op2)
				fmt. Println ( "O valor total em meses é: " , valor_total_meses_op2) // ele printa na tela o valor total em meses

		// Conta para juros compostos

				vf_op2 := `
				********************************
				* valor_final_op2 *
				********************************
				` //ta usando as estrelinhas para atribuir o valor do que ta escrito dentro
				valor_final_op2  := vtm_op2 * (1+9)^tcm_op2  // resultado da multiplicação valor total em meses para o contribuicao 
				fmt. Println (vf_op2)
				fmt. Println ( "Você irá receber por mês " , valor_final_op2) // ele printa na tela o valor final que voce irá receber por mês
	}
	
	fmt. Println ( " Presione enter para sair \n " )
	var  espera  corda
	fmt. Scanln (& espera)
	fmt. Println ( " Tchau! " )


}