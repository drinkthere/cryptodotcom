package public_data

import (
	"github.com/drinkthere/cryptodotcom/models/publicdata"
	"github.com/drinkthere/cryptodotcom/responses"
)

type (
	GetInstrumentsResult struct {
		Data []*publicdata.Instrument `json:"data"`
	}
	GetInstruments struct {
		responses.Basic
		Result GetInstrumentsResult `json:"result"`
	}
)
