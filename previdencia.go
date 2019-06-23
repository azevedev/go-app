package main
 

var es []float64

func varInit(){
	es = []float64{36.6, 35.7, 34.8, 34.0, 33.1, 32.2, 31.4, 30.5, 29.7, 28.8, 28.0, 27.2, 26.4, 25.6, 24.8, 24.0, 23.2, 22.4, 21.6, 20.9, 20.1, 19.4, 18.7, 18.0, 17.3, 16.6, 15.9, 15.2}
}

func calcula_salario_aposentadoria(idade int, tempC int, valorMensal float64) float64{
	a := 0.08 //taxa selic reduzida
	temp := tempC * 12
	total := 0.0
	for i := 1; i <= temp; i++ {
		total+= valorMensal
		if(i%12 == 0){
			total += total*a
		}
	}
	aliquota := 0.095 // os valores mais comuns são 11% e 8%. fiz uma média.
	ft := ((float64(tempC) * aliquota)/es[idade-43])*(1 +((float64(tempC) * aliquota)/100))
	//calculo do INSS para o Fator Previdenciário, baseado na espectativa de vida do aposentado.
	//o valor da alíquota foi reduzido, pois aqui considera-se um salário 'até morrer' para o aposentado
	total *= ft
	total /= float64(120-idade)
	return total

}

func calcula_idade_aposentadoria(){

}

func calculo_investimento_em_previdencia_mensal(){

}


//
// Previdencia() {
// 	Get_Renda()
// 	Get_Idade()
// 	Get_Investimento_em_previdencia_mensal()
// 	Get_Idade_aposentadoria()
// 	Get_Salario_aposentadoria()
// 	//Previsao_aposentadoria
// 	Calcula_Salario_aposentadoria()
// 	Calcula_Idade_aposentadoria() 
// 	Calculo_Investimento_em_previdencia_mensal()
// }
//
