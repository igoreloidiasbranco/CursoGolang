package arquitetura

// podemos realizar testes em determinadas arquiteturas selecionadas
import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}

	// ...
	//t.Fail()
}
