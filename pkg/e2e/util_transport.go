package e2e

import (
	"context"
	"github.com/danyazab/animal/pkg/http/client"
	"net/http"
)

type Transport interface {
	Get(ctx context.Context, url string, res interface{}) error
	Post(ctx context.Context, url string, body interface{}, res interface{}) error
}

type transport struct {
	rc client.Transport
}

func (t *transport) Get(ctx context.Context, url string, res interface{}) error {
	clientReq := t.rc.R()
	clientReq.
		SetContext(ctx).
		SetHeader(client.HdrContentTypeKey, "application/json").
		SetHeader(http.CanonicalHeaderKey("Accept"), "application/json")

	return t.rc.Execute(clientReq, http.MethodGet, url, &res)
}

func (t *transport) Post(ctx context.Context, url string, body interface{}, res interface{}) error {
	clientReq := t.rc.R()
	clientReq.
		SetContext(ctx).
		SetBody(body).
		SetHeader(client.HdrContentTypeKey, "application/json").
		SetHeader(http.CanonicalHeaderKey("Accept"), "application/json")

	return t.rc.Execute(clientReq, http.MethodPost, url, &res)
}
