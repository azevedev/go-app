package main

// Usuario struct que denfine o user
type Usuario struct {
	id                 int
	renda              float64
	idade              int
	nome               string
	perfilInvestimento string
}

//cria um novo usuário e o retorna
func newUser(id int, renda float64, idade int, nome string, pf string) Usuario {
	user := Usuario{
		id:                 id,
		renda:              renda,
		idade:              idade,
		nome:               nome,
		perfilInvestimento: pf,
	}
	return user
}

//cria um novo usuário com atributos nulos e o retorna
func userNew() Usuario {
	user := new(Usuario)
	return *user
}
