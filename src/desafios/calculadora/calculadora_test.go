package main

import (
	"testing"
)

// padrão triple A - AAA
// A - Arrange (preparar)
// A - Act (Rodar)
// A - Assert (verificar as asserções)

func TestSoma(t *testing.T) {
	teste := soma(6, 7)
	resultado := 13
	if teste != resultado {
		t.Error("\n\nValor esperado: ", resultado, "\nValor Obtido: ", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := soma(6, 7)
	resultado := 12
	if teste != resultado {
		t.Error("\n\nValor esperado: ", resultado, "\nValor Obtido: ", teste)
	}
}

func TestSubtracao(t *testing.T) {
	teste := subtracao(171, 157)
	resultado := -328
	if teste != resultado {
		t.Error("\n\nValor esperado: ", resultado, "\nValor Obtido: ", teste)
	}
}

func TestSubtracao2(t *testing.T) {
	teste := subtracao(171, 157)
	resultado := -14
	if teste != resultado {
		t.Error("\n\nValor esperado: ", resultado, "\nValor Obtido: ", teste)
	}
}

func TestMultiplicacao(t *testing.T) {
	teste := multiplicacao(1, 2, 3)
	resultado := 6
	if teste != resultado {
		t.Error("\n\nValor esperado: ", resultado, "\nValor Obtido: ", teste)
	}
}

func TestMultiplicacao2(t *testing.T) {
	teste := multiplicacao(1, 2, 3)
	resultado := 7
	if teste != resultado {
		t.Error("\n\nValor esperado: ", resultado, "\nValor Obtido: ", teste)
	}
}

func TestDivisao(t *testing.T) {
	teste := divisao(20, 40)
	var resultado float64 = 2
	if teste != resultado {
		t.Error("\n\nValor esperado: ", resultado, "\nValor Obtido: ", teste)
	}
}

func TestDivisao2(t *testing.T) {
	teste := divisao(20, 40)
	var resultado float64 = 800
	if teste != resultado {
		t.Error("\n\nValor esperado: ", resultado, "\nValor Obtido: ", teste)
	}
}