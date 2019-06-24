package main

import "fmt"

func simuladorInvestimentos() {

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

func main() {
	analisePerfil()
}
