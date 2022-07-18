package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type routeName string

const RouteNameKey = routeName("routeName")

var (
	HdrContentTypeKey = http.CanonicalHeaderKey("Content-Type")
	hdrUserAgentKey   = http.CanonicalHeaderKey("User-Agent")
)

type Request struct {
	*resty.Request
}

func (r *Request) WithAuth(token string) *Request {
	r.SetAuthToken(token)

	return r
}

type Transport interface {
	R() *Request
	// Execute - execute raw request
	Execute(req *Request, method, url string, target interface{}) error
}

func NewTransport(httpClient *resty.Client) Transport {
	return &transport{
		httpClient: httpClient,
		userAgent: fmt.Sprintf(
			"%s/%s (http://localhost:8000)",
			"animals",
			"0.0.1",
		),
	}
}

type transport struct {
	httpClient *resty.Client
	userAgent  string
}

func (t *transport) R() *Request {
	r := t.httpClient.R()
	r.SetHeader(hdrUserAgentKey, t.userAgent)

	return &Request{r}
}

func (t transport) Execute(req *Request, method, url string, target interface{}) error {
	resp, err := req.Execute(method, url)
	if err != nil {
		return err
	}

	return t.handleExecResponse(resp, url, target)
}

func (t transport) handleExecResponse(resp *resty.Response, url string, target interface{}) error {
	if resp.IsError() {
		return NewHTTPClientError(resp)
	}

	// Sometimes resty.Response{}.StatusCode() may return 200,
	// but resty.RawResponse{}.StatusCode not
	if resp.RawResponse.StatusCode > 399 {
		return NewHTTPClientError(resp)
	}

	if target == nil {
		return nil
	}

	if resp.Size() == 0 {
		return nil
	}

	b := resp.Body()
	if err := json.Unmarshal(b, target); err != nil {
		return fmt.Errorf("%w JSON decode failed on %s: %s", err, url, resp.String())
	}

	return nil
}
