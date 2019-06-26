/*  Programa: investimentos.go
    Autor: Alexander Matheus de Melo Lima
    Versão: 1.0
    Concluída em: 24/06/2019

	Breve descrição:
		Funções que calculam o rendimento da poupança e o rendimento dos investimentos usando
		como parâmetros o depósito inicial e a quantidade de meses que o dinheiro vai ficar
		aplicado e levando em conta a porcentagem da rentabilidade de cada investimento ou poupança.

	Considerações utilizadas na simulação:
		1. Data da última atualização: 04/04/2018
		2. de rentabilidade da Poupança: 0,3715% a.m.
		3. Tesouro Selic: 6,50% a.a.
		4. Tesouro Prefixado: 8,05%
		5. Tesouro IPCA+: IPCA + 5,00% (Inflação 3% a.a.)

		* Valores referentes a data da última atualização e podem sofrer alterações de acordo com o mercado.*/

package main

import "fmt"

// Rendimento da poupança
func poupanca(depositoInicial float64, meses float64) float64 {
	var poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
	}
	poupanca = poupanca - depositoInicial
	return poupanca
}

// Rendimento do Tesouro Pré Fixado com base na poupança
func tesouroPrefixado(depositoInicial float64, meses float64) float64 {
	var tesouroPrefixado, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		tesouroPrefixado = (poupanca - depositoInicial) * 1.4538
	}
	return tesouroPrefixado
}

// Rendimento do Tesouro Selic com base na poupança
func tesouroSelic(depositoInicial float64, meses float64) float64 {
	var tesouroSelic, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		tesouroSelic = (poupanca - depositoInicial) * 1.0315
	}
	return tesouroSelic
}

// Rendimento do Tesouro IPCA+ com base na poupança
func tesouroIPCA(depositoInicial float64, meses float64) float64 {
	var tesouroIPCA, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		tesouroIPCA = (poupanca - depositoInicial) * 1.2333
	}
	return tesouroIPCA
}

// Rendimento de CDB e LC com base na poupança
func CDBeLC(depositoInicial float64, meses float64) float64 {
	var CDBeLC, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		CDBeLC = (poupanca - depositoInicial) * 1.278
	}
	return CDBeLC
}

// Rendimento de LCI e LCA com base na poupança
func LCIeLCA(depositoInicial float64, meses float64) float64 {
	var LCIeLCA, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		LCIeLCA = (poupanca - depositoInicial) * 1.246
	}
	return LCIeLCA
}

// A análise de perfil analisa o perfil do investidor com base nas respostas do usuário
//somando a pontuação de acordo com os níveis das respostas, resultando no perfil
func analisePerfil() {

	var resposta, perfil int

	fmt.Println("1. Você acha fácil investir? Sim [1] Não [0]")
	fmt.Scan(&resposta)
	perfil = perfil + resposta
	fmt.Println("2. Você possui alguma experiência ou formação na área financeira? Sim [1] Não [0]")
	fmt.Scan(&resposta)
	perfil = perfil + resposta
	fmt.Println("3. E antes de tomar qualquer decisão sobre investir, o que você procura fazer?")
	fmt.Println("a) Converso com amigos ou familiares [0]")
	fmt.Println("b) Consulto um especialista [1]")
	fmt.Println("c) Leio noticias sobre o mercado [2]")
	fmt.Scan(&resposta)
	perfil = perfil + resposta
	fmt.Println("4. E o que você mais valoriza ao investir?")
	fmt.Println("a) Ganhar sempre mais [0]")
	fmt.Println("b) Evitar perder capital [1]")
	fmt.Scan(&resposta)
	perfil = perfil + resposta
	fmt.Println("5. Em quanto tempo você pretende atingir esse objetivo?")
	fmt.Println("a) Até 1 ano [0]")
	fmt.Println("b) Entre 1 e 5 anos [1]")
	fmt.Println("c) Entre 5 e 10 anos [2]")
	fmt.Println("d) Acima de 10 anos [3]")
	fmt.Scan(&resposta)
	perfil = perfil + resposta
	fmt.Println("6. Qual é a sua renda mensal atualmente?")
	fmt.Println("a) Até R$ 4.000,00 [0]")
	fmt.Println("b) Entre R$ 4.000,00 e R$ 10.000,00 [1]")
	fmt.Println("c) Acima de 10.000,00 [2]")
	fmt.Scan(&resposta)
	perfil = perfil + resposta

	if perfil <= 4 {
		fmt.Println("Perfil: Conservador")
	} else if perfil >= 5 && perfil <= 8 {
		fmt.Println("Perfil: Moderado")
	} else if perfil >= 9 {
		fmt.Println("Perfil: Arrojado")
	}

}

//Testes
/*func main() {
	var depositoInicial, meses float64
	fmt.Println("Digite o depósito inicial em R$:")
	fmt.Scan(&depositoInicial)
	fmt.Println("Digite o tempo de investimento em meses:")
	fmt.Scan(&meses)
	poupanca(depositoInicial, meses)
	tesouroPrefixado(depositoInicial, meses)
	tesouroSelic(depositoInicial, meses)
	tesouroIPCA(depositoInicial, meses)
	CDBeLC(depositoInicial, meses)
	LCIeLCA(depositoInicial, meses)
}*/
