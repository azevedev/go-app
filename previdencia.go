package main

import (
	"math"
)

//Calcula o salario recebido na aposentadoria
func  CalculadoraAposentadoria (Usr Usuario) float64{
	TaxaAnualPrevPrivadaEstimada := 3.74
	expec_vida_brasileiro_contribuinte := 80
	TaxaMensal := ((math.Pow(TaxaAnualPrevPrivadaEstimada, 1.0/12))-1.0)/100.0
	Meses_contribuindo := (Usr.expec_idade_aposent - Usr.idade)*12
	Montante := 0.0
	for i := 0; i < Meses_contribuindo; i++ {
		Montante = Montante + (Montante * TaxaMensal)
		Montante = Montante + (Usr.expec_contribuicao)

	}
	Meses_aposentado := (expec_vida_brasileiro_contribuinte - Usr.expec_idade_aposent) * 12
	aposentadoria := Montante/((1.0-(1.0/math.Pow((1+TaxaMensal), float64(Meses_aposentado))))/TaxaMensal)
	return aposentadoria
}

//Calcula quanto precisa contribuir para quanhar o salario desejado
func  CalculadoraContribuicao (Usr Usuario) float64{
	TaxaAnualPrevPrivadaEstimada := 3.74
	expec_vida_brasileiro_contribuinte := 80
	TaxaMensal := ((math.Pow(TaxaAnualPrevPrivadaEstimada, 1.0/12))-1.0)/100.0
	Meses_contribuindo := (Usr.expec_idade_aposent - Usr.idade)*12
	Meses_aposentado := (expec_vida_brasileiro_contribuinte - Usr.expec_idade_aposent) * 12
	x := ((1.0-(1.0/math.Pow((1+TaxaMensal), float64(Meses_aposentado))))/TaxaMensal)
	Montante := x * Usr.expec_salario
	y := (1+TaxaMensal)*((math.Pow(1+TaxaMensal, float64(Meses_contribuindo))-1)/TaxaMensal)
	contribuicao := Montante/y
	return contribuicao
}

