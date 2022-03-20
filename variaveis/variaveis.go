package main

import "fmt"

func main() {
	// https://www.w3schools.com/go/go_variable_naming_rules.php

	// ------------------ Declaracao de var

	var logs, erros int = 0, 0
	// Camel Case
	var analistaDeSistemas string = "Alan"

	var (
		versao           float32 = 0.1
		autoCarregamento bool    = true
		backupAutomatico bool    = false
		geradorDeLogs    bool    = true
	)

	const criador string = "Alan da Silva Alves"
	const dataCriacao string = "20/03/2022"

	mensagemFim := "Fim do progrma com inferencia de var"

	// ------------------ Fim da declaracao de var

	fmt.Printf("Logs: %v\nErros: %v\n", logs, erros)
	fmt.Println(analistaDeSistemas)

	fmt.Printf("Sou analista de sistemas %v.", analistaDeSistemas)
	fmt.Printf(" Sou uma var do tipo %T\n", analistaDeSistemas)

	fmt.Printf("- Criado por: %v\n- Criado em: %v\n", criador, dataCriacao)

	fmt.Printf(
		"SystemInfo:\n Verão: %v\n Load: %v\n Backup: %v\n Logs: %v\n",
		versao,
		autoCarregamento,
		backupAutomatico,
		geradorDeLogs)

	fmt.Printf("\n..... %v.....\n", mensagemFim)
}
