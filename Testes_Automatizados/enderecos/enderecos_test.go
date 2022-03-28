package enderecos

import "testing"

// TestTipoDeEndereco : Testes unitários
func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Erasmo Braga"

	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TiposDeEnderecos(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e rebebeu %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido,
		)
	}
}
