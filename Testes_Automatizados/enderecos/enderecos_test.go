package enderecos_test

// package enderecos

// import "testing"

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// TestTipoDeEndereco : Testes unitários
func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"Estrada ABC", "Estrada"},
		{"Praca ABC", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {

		retornoDeEnderecoRecebido := TiposDeEnderecos(cenario.enderecoInserido)

		if retornoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e rebebeu %s",
				retornoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
