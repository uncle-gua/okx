package api

import (
	"context"

	"github.com/uncle-gua/okx"
	okex "github.com/uncle-gua/okx"
	"github.com/uncle-gua/okx/api/rest"
	"github.com/uncle-gua/okx/api/ws"
)

// Client is the main api wrapper of okex
type Client struct {
	Rest *rest.ClientRest
	Ws   *ws.ClientWs
	ctx  context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination okx.Destination) (*Client, error) {
	restURL := okx.RestURL
	wsPubURL := okx.PublicWsURL
	wsPriURL := okx.PrivateWsURL
	switch destination {
	case okx.AwsServer:
		restURL = okx.AwsRestURL
		wsPubURL = okx.AwsPublicWsURL
		wsPriURL = okx.AwsPrivateWsURL
	case okx.DemoServer:
		restURL = okx.DemoRestURL
		wsPubURL = okx.DemoPublicWsURL
		wsPriURL = okx.DemoPrivateWsURL
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okex.BaseURL{true: wsPriURL, false: wsPubURL})

	return &Client{r, c, ctx}, nil
}
