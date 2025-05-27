package api

import (
	"context"
	"github.com/drinkthere/cryptodotcom"
	"github.com/drinkthere/cryptodotcom/api/rest"
	"github.com/drinkthere/cryptodotcom/api/ws"
)

// Client is the main api wrapper of okx
type Client struct {
	Rest *rest.ClientRest
	Ws   *ws.ClientWs
	ctx  context.Context
}

func NewClient(ctx context.Context, apiKey, secretKey string, destination cryptodotcom.Destination, ip string) (*Client, error) {
	restURL := cryptodotcom.RestURL
	wsMktURL := cryptodotcom.MarketWsURL
	wsUserURL := cryptodotcom.UserWsURL
	switch destination {
	case cryptodotcom.ColoServer:
		restURL = cryptodotcom.ColoRestURL
		wsMktURL = cryptodotcom.ColoMarketWsURL
		wsUserURL = cryptodotcom.ColoUserWsURL
	}

	r := rest.NewClient(apiKey, secretKey, restURL, destination, ip)
	c := ws.NewClient(ctx, apiKey, secretKey, map[bool]cryptodotcom.BaseURL{true: wsUserURL, false: wsMktURL}, ip)

	return &Client{r, c, ctx}, nil
}
