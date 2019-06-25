package main

// Usuario struct que denfine o user
type Usuario struct {
    id             int
    renda              float64
    idade              int
    nome               string
    perfilInvestimento string
    expec_idade_aposent int
    expec_contribuicao float64
    expec_salario float64
}

//cria um novo usuário e o retorna
func newUser(id int, renda float64, idade int, nome string, pf string, EIA int, EC float64, ES float64) Usuario {
	user := Usuario{
		id:                 id,
		renda:              renda,
		idade:              idade,
		nome:               nome,
		perfilInvestimento: pf,
		expec_idade_aposent: EIA,
		expec_contribuicao: EC,
		expec_salario: ES,
	}
	return user
}

//cria um novo usuário com atributos nulos e o retorna
func userNew() Usuario {
	user := new(Usuario)
	return *user
}
