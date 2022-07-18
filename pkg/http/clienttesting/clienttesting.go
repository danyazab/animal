package clienttesting

import (
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

const host = "https://testing.com"

var state = struct {
	rc *resty.Client
}{}

func Setup(m *testing.M, d bool) int {
	c := &http.Client{}
	httpmock.ActivateNonDefault(c)
	defer httpmock.DeactivateAndReset()

	rc := resty.NewWithClient(c)
	rc.SetDebug(d)
	rc.SetBaseURL(host)

	state.rc = rc

	return m.Run()
}

func RegisterHttpResponse(method, url, body string, code int) {
	httpmock.RegisterResponder(method, host+url, httpmock.NewStringResponder(code, body))
}

// Inject provides a transaction within test container
func Inject(f func(*testing.T, *resty.Client)) func(*testing.T) {
	return func(t *testing.T) {
		// mark this function as a helper function for stack introspection
		t.Helper()

		// remove any mocks
		httpmock.Reset()

		f(t, state.rc)
	}
}
