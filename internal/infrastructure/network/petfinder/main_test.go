package petfinder

import (
	"github.com/danyazab/animal/pkg/http/clienttesting"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(clienttesting.Setup(m, false))
}
