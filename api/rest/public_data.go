package rest

import (
	"encoding/json"
	"github.com/drinkthere/cryptodotcom"
	responses "github.com/drinkthere/cryptodotcom/responses/public_data"
	"net/http"
)

type PublicData struct {
	client *ClientRest
}

// NewPublicData returns a pointer to a fresh PublicData
func NewPublicData(c *ClientRest) *PublicData {
	return &PublicData{c}
}

// GetInstruments
// Retrieve a list of instruments with open contracts.
func (c *PublicData) GetInstruments() (response responses.GetInstruments, err error) {
	p := "/public/get-instruments"
	m := cryptodotcom.S2M(map[string]string{})
	res, err := c.client.Do(http.MethodGet, p, false, m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)
	err = d.Decode(&response)
	return
}
