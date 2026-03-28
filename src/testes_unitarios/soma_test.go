// importante!!!

/*
Para testar com sucesso é necessário estar em um projeto
crie uma pasta ou use o nome de uma existente

mkdir <nome_pasta>
cd <nome_pasta>
go mod init example.com/<nome_projeto>
*/

package main

import "testing"

// padrão triple A - AAA
// A - Arrange (preparar)
// A - Act (Rodar)
// A - Assert (verificar as asserções)

func TestShouldSumCorrect(t *testing.T) {  // É importante que os nomes das funções de teste por convenção comecem com Uppercase e "Test"
	// arrange
	teste := soma(3, 2, 1)
	// act
	resultado := 6
	// assert
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}
}

func TestShouldMultiplyCorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}
}

func TestShouldMultiplyIncorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 256
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado", teste)
	}
}
