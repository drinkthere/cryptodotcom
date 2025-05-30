package rest

import (
	"encoding/json"
	"github.com/drinkthere/cryptodotcom"
	requests "github.com/drinkthere/cryptodotcom/requests/rest/trade"
	responses "github.com/drinkthere/cryptodotcom/responses/trade"
	"net/http"
)

type Trade struct {
	client *ClientRest
}

func NewTrade(c *ClientRest) *Trade {
	return &Trade{c}
}

func (c *Trade) CreateOrder(req requests.CreateOrder) (response responses.CreateOrder, err error) {
	p := "private/create-order"
	m := cryptodotcom.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

func (c *Trade) GetOpenOrders(req requests.GetOpenOrders) (response responses.GetOpenOrders, err error) {
	p := "private/get-open-orders"
	m := cryptodotcom.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

func (c *Trade) CancelOrder(req requests.CancelOrder) (response responses.CancelOrder, err error) {
	var p string
	var res *http.Response

	p = "private/cancel-order"
	m := cryptodotcom.S2M(req)
	res, err = c.client.Do(http.MethodPost, p, true, m)

	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

func (c *Trade) CancelAllOrders(req requests.CancelAllOrders) (response responses.CancelAllOrders, err error) {
	var p string
	var res *http.Response

	p = "private/cancel-all-orders"
	m := cryptodotcom.S2M(req)
	res, err = c.client.Do(http.MethodPost, p, true, m)

	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

func (c *Trade) GetAccountSettings() (response responses.GetAccountSettings, err error) {
	var p string
	var res *http.Response

	p = "private/get-account-settings"
	res, err = c.client.Do(http.MethodPost, p, true, make(map[string]interface{}))

	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}

func (c *Trade) ChangeAccountLeverage(req requests.ChangeAccountLeverage) (response responses.ChangeAccountLeverage, err error) {
	var p string
	var res *http.Response

	p = "private/change-account-leverage"
	m := cryptodotcom.S2M(req)
	res, err = c.client.Do(http.MethodPost, p, true, m)

	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}
