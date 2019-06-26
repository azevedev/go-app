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

func Rendimentos(dep float64, meses float64) (string, string, string, string, string, string){
	a:= fmt.Sprintf("poupanca:%.2f R$", poupanca(dep, meses))
	b:=fmt.Sprintf("tesouroPrefixado:%.2f R$", tesouroPrefixado(dep, meses))
	c:=fmt.Sprintf("tesouroSelic:%.2f R$", tesouroSelic(dep, meses))
	d:=fmt.Sprintf("tesouroIPCA:%.2f R$", tesouroIPCA(dep, meses))
	e:=fmt.Sprintf("CDBeLC:%.2f R$", CDBeLC(dep, meses))
	f:=fmt.Sprintf("LCIeLCA:%.2f R$", LCIeLCA(dep, meses))
	return a, b, c, d, e, f
}

// A análise de perfil analisa o perfil do investidor com base nas respostas do usuário
//somando a pontuação de acordo com os níveis das respostas, resultando no perfil
func analisePerfil(perfil int) string{
	if perfil <= 4 {
		return "Conservador"
	} else if perfil >= 5 && perfil <= 8 {
		return "Moderado"
	} else if perfil >= 9 {
		return "Arrojado"
	}
	return ""

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
