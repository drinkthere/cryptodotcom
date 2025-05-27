package public_data

import (
	"github.com/drinkthere/cryptodotcom/models/publicdata"
	"github.com/drinkthere/cryptodotcom/responses"
)

type (
	GetInstruments struct {
		responses.Basic
		Data []*publicdata.Instrument `json:"result"`
	}
)
