package matematica

import "testing"

const erroPadrao = "Valor esperado: %v, Valor recebido: %v"

// padrão Teste + nome da função e passando parâmetro t que é um ponteiro para o objeto testing.T
func TestMedia(t *testing.T) {

	resultado := Media(7.2, 9.9, 6.1, 5.9)
	esperado := 7.28

	if resultado != esperado {
		t.Errorf(erroPadrao, esperado, resultado)
	}
}
