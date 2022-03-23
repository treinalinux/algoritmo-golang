package main

import "fmt"

func menuSoftware(opcao int) string {
	switch opcao {
	case 1:
		return " Gerenciar DNS"
	case 2:
		return " Gerenciar DHCP"
	case 3:
		return " Gerenciar NIS"
	case 4:
		return " Gerenciar NTP"
	case 5:
		return "Sair"
	default:
		return " Opcao invalida!"
	}
}

func main() {
	menu := menuSoftware(1)
	fmt.Println(menu)
}
