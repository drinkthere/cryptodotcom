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

/*

// CancelOrder
// Cancel an incomplete order.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-cancel-order
//
// Cancel incomplete orders in batches. Maximum 20 orders can be canceled at a time. Request parameters should be passed in the form of an array.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-cancel-multiple-orders
func (c *Trade) CancelOrder(req []requests.CancelOrder) (response responses.CancelOrder, err error) {
	var p string
	var res *http.Response
	if len(req) > 1 {
		p = "/api/v5/trade/cancel-batch-orders"
		var m interface{}
		m = req
		res, err = c.client.DoBatch(p, m)
	} else {
		p = "/api/v5/trade/cancel-order"
		m := okx.S2M(req[0])
		res, err = c.client.Do(http.MethodPost, p, true, m)
	}
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}


// GetOrderList
// Retrieve all incomplete orders under the current account.
//
// https://www.okx.com/docs-v5/en/#rest-api-trade-get-order-list
func (c *Trade) GetOrderList(req requests.OrderList) (response responses.OrderList, err error) {
	p := "/api/v5/trade/orders-pending"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}
*/
