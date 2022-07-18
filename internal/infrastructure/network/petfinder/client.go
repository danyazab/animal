package petfinder

import (
	"context"
	"danyazab/animal/pkg/http/client"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	appId     string
	appSecret string

	transport client.Transport
}

func NewClient(transport client.Transport) Client {
	return Client{
		appId:     "NKfpgAFDY3u7n0ME0aWAX8HlulkG1bgx9Ffi03Lb83WOlmdNAL",
		appSecret: "MVllHQTuOTUei9Aq8JxR4wGa9f3DFA5kO1QNDVUL",
		transport: transport,
	}
}

func (c *Client) GetCatsBreeds(ctx context.Context) ([]string, error) {
	token, err := c.getAuthToken(ctx)
	if err != nil {
		return nil, err
	}

	return c.catsBreeds(ctx, token)
}

func (c *Client) getAuthToken(ctx context.Context) (string, error) {
	route := "/v2/oauth2/token"

	clientReq := c.transport.R()
	clientReq.
		SetContext(context.WithValue(ctx, client.RouteNameKey, route)).
		SetHeader(client.HdrContentTypeKey, "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     c.appId,
			"client_secret": c.appSecret,
		})

	result := struct {
		AccessToken string `json:"access_token"`
	}{}

	return result.AccessToken, c.transport.Execute(clientReq, resty.MethodPost, route, &result)
}

func (c *Client) catsBreeds(ctx context.Context, token string) ([]string, error) {
	route := "/v2/types/cat/breeds"
	result := struct {
		Breeds []struct {
			Name string `json:"name"`
		} `json:"breeds"`
	}{}

	clientReq := c.transport.R().WithAuth(token)
	clientReq.SetContext(context.WithValue(ctx, client.RouteNameKey, route))

	if err := c.transport.Execute(clientReq, resty.MethodGet, route, &result); err != nil {
		return nil, err
	}

	res := make([]string, len(result.Breeds))
	for i, v := range result.Breeds {
		res[i] = v.Name
	}

	return res, nil
}
