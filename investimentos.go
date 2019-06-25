package main

import "fmt"

func poupanca(depositoInicial float64, meses float64) {
	var poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
	}
	poupanca = poupanca - depositoInicial
	fmt.Printf("Na poupança você renderia R$ %.2f\n", poupanca)
}

func tesouroPrefixado(depositoInicial float64, meses float64) {
	var tesouroPrefixado, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		tesouroPrefixado = (poupanca - depositoInicial) * 1.4538
	}
	fmt.Printf("No Tesouro Pré Fixado você renderá R$ %.2f\n", tesouroPrefixado)
}

func tesouroSelic(depositoInicial float64, meses float64) {
	var tesouroSelic, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		tesouroSelic = (poupanca - depositoInicial) * 1.0315
	}
	fmt.Printf("No Tesouro Selic você renderá R$ %.2f\n", tesouroSelic)
}

func tesouroIPCA(depositoInicial float64, meses float64) {
	var tesouroIPCA, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		tesouroIPCA = (poupanca - depositoInicial) * 1.2333
	}
	fmt.Printf("No Tesouro IPCA+ você renderá R$ %.2f\n", tesouroIPCA)
}

func CDBeLC(depositoInicial float64, meses float64) {
	var CDBeLC, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		CDBeLC = (poupanca - depositoInicial) * 1.278
	}
	fmt.Printf("No CDB e LC você renderá R$ %.2f\n", CDBeLC)
}

func LCIeLCA(depositoInicial float64, meses float64) {
	var LCIeLCA, poupanca float64
	var periodo int = int(meses)
	poupanca = depositoInicial
	for i := 1; i <= periodo; i++ {
		poupanca = poupanca * 1.003715
		LCIeLCA = (poupanca - depositoInicial) * 1.246
	}
	fmt.Printf("No LCI e LCA você renderá R$ %.2f\n", LCIeLCA)
}

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
