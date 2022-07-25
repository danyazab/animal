package e2e

import (
	"github.com/danyazab/animal/pkg/http/client"
	"github.com/go-resty/resty/v2"
	"testing"
)

var state = transport{}

func Setup(m *testing.M, host string, debug bool) int {
	rc := resty.New()

	rc.SetDebug(debug)
	rc.SetBaseURL(host)

	state.rc = client.NewTransport(rc)

	return m.Run()
}

// Inject provides a transaction within test container
func Inject(f func(*testing.T, Transport)) func(*testing.T) {
	return func(t *testing.T) {
		// mark this function as a helper function for stack introspection
		t.Helper()

		f(t, &state)
	}
}
