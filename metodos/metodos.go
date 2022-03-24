package main

import "fmt"

// Usuario ...
type Usuario struct {
	nome  string
	idade uint8
}

// Metodo
func (u Usuario) salvar() {
	fmt.Println("Dentro do metodo salvar")
	fmt.Printf("Salvando no banco o user: %s\n", u.nome)
}

// Metodo
func (u Usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Metodo
func (u *Usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuarioUm := Usuario{"Alan", 43}
	fmt.Println(usuarioUm)
	usuarioUm.salvar()
	usuarioUm.fazerAniversario()
	fmt.Println(usuarioUm.idade)

	usuarioDois := Usuario{"Carla Fe.", 25}
	fmt.Println(usuarioDois)
	usuarioDois.salvar()
	maiorDeIdade := usuarioDois.maiorDeIdade()
	fmt.Println(maiorDeIdade)
}
