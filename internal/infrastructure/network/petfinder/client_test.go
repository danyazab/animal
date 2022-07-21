package petfinder

import (
	"context"
	"danyazab/animal/pkg/http/client"
	"danyazab/animal/pkg/http/clienttesting"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func Test_GetCatsBreeds(t *testing.T) {
	type testCase = struct {
		name    string
		breeds  []string
		mocHttp func()
	}

	cases := []testCase{
		{
			name:   "load 2 beards by http",
			breeds: []string{"Abyssinian", "American Bobtail"},
			mocHttp: func() {
				clienttesting.RegisterHttpResponse(
					http.MethodPost,
					"/v2/oauth2/token",
					`{"token_type":"Bearer","expires_in":3600,"access_token":"yyy.xxx.zzz"}`,
					http.StatusOK,
				)
				clienttesting.RegisterHttpResponse(
					http.MethodGet,
					"/v2/types/cat/breeds",
					`{"breeds":[
						{"name":"Abyssinian","_links":{"type":{"href":"\/v2\/types\/cat"}}},
						{"name":"American Bobtail","_links":{"type":{"href":"\/v2\/types\/cat"}}}
					]}`,
					http.StatusOK,
				)
			},
		},
		{
			name:   "http endpoint return null instead empty array",
			breeds: []string{},
			mocHttp: func() {
				clienttesting.RegisterHttpResponse(
					http.MethodPost,
					"/v2/oauth2/token",
					`{"token_type":"Bearer","expires_in":3600,"access_token":"yyy.xxx.zzz"}`,
					http.StatusOK,
				)
				clienttesting.RegisterHttpResponse(
					http.MethodGet,
					"/v2/types/cat/breeds",
					`{"breeds":null}`,
					http.StatusOK,
				)
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d case: %s", i, c.name), clienttesting.Inject(func(t *testing.T, rc *resty.Client) {
			c.mocHttp()

			cl := NewClient(
				client.NewTransport(rc),
			)

			breeds, err := cl.GetCatsBreeds(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, c.breeds, breeds)
		}))
	}
}
