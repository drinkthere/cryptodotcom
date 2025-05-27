package rest

import (
	"encoding/json"
	"github.com/drinkthere/cryptodotcom"
	requests "github.com/drinkthere/cryptodotcom/requests/rest/account"
	responses "github.com/drinkthere/cryptodotcom/responses/account"
	"net/http"
)

type Account struct {
	client *ClientRest
}

func NewAccount(c *ClientRest) *Account {
	return &Account{c}
}

func (c *Account) GetAccounts(req requests.GetAccounts) (response responses.GetAccounts, err error) {
	p := "private/get-accounts"
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

func (c *Account) GetBalances() (response responses.GetBalances, err error) {
	p := "private/user-balance"
	m := cryptodotcom.S2M(map[string]string{})
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

func (c *Account) GetPositions(req requests.GetPositions) (response responses.GetPositions, err error) {
	p := "private/get-positions"
	m := cryptodotcom.S2M(req)
	res, err := c.client.Do(http.MethodGet, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}

/*

//// SetLeverage
//func (c *Account) SetLeverage(req requests.SetLeverage) (response responses.Leverage, err error) {
//	p := "/api/v5/account/set-leverage"
//	m := okx.S2M(req)
//	res, err := c.client.Do(http.MethodPost, p, true, m)
//	if err != nil {
//		return
//	}
//	defer res.Body.Close()
//	d := json.NewDecoder(res.Body)
//	err = d.Decode(&response)
//
//	return
//}

// SetAccountLevel
//
// https://www.okx.com/docs-v5/zh/#trading-account-rest-api-set-account-mode
func (c *Account) SetAccountLevel(req requests.SetAccountLevel) (response responses.SetAccountLevel, err error) {
	p := "/api/v5/account/set-account-level"
	m := okx.S2M(req)
	res, err := c.client.Do(http.MethodPost, p, true, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)

	return
}
*/
