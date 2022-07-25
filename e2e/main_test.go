//-go:build e2e

package e2e

import (
	"github.com/danyazab/animal/pkg/e2e"
	"os"
)

import (
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(e2e.Setup(m, "http://localhost:8000", false))
}
