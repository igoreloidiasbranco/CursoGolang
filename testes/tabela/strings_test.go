//Dataset - técnica para verificar vários cenários de teste

package strings

import (
	"strings"
	"testing"
)

// mensagem de erro padrão para o teste de index
const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {

	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{ //Dataset de testes
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"igor", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
