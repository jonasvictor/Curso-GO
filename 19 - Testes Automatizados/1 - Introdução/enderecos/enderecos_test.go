// TESTE DE UNIDADE
package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// Começando obrigatóriamente com a palavra Test e com letra maiúscula
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		//{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA UFPB", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		//{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido, cenario.retornoEsperado)
		}
	}

	//if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	//	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
	//		tipoDeEnderecoEsperado,
	//		tipoDeEnderecoRecebido)
	//}

}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
