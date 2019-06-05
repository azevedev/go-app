package main	

//struct do usuario
type Usuario struct{
	id int
	renda float64
	idade int
	nome string
	perfil_investimento string
	

}

//cria um novo usuário e o retorna
func new_user(id int, renda float64, idade int, nome string, p_f string) Usuario{
	user := Usuario{
		id: id,
		renda: renda,
		idade: idade,
		nome: nome,
		perfil_investimento: p_f,
	}
	return user
}

//cria um novo usuário com atributos nulos e o retorna
func user_new() Usuario{
	user := new(Usuario)
	return *user
}