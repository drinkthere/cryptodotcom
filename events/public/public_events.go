package public

import (
	"github.com/drinkthere/cryptodotcom"
	"github.com/drinkthere/cryptodotcom/events"
)

type (
	Tickers struct {
		events.Basic
		result TickerResult `json:"data"`
	}
	TickerResult struct {
		InstrumentName string    `json:"instrument_name"`
		Subscription   string    `json:"subscription"`
		Channel        string    `json:"channel"`
		Data           []*Ticker `json:"data"`
	}

	Ticker struct {
		HighPx         string                `json:"h"`
		LowPx          string                `json:"l"`
		LastTradePx    string                `json:"a"`
		PxChange       string                `json:"c"`
		BidPx          string                `json:"b"`
		BidSz          string                `json:"bs"`
		AskPx          string                `json:"k"`
		AskSz          string                `json:"ks"`
		InstrumentName string                `json:"i"`
		Volume         string                `json:"v"`
		Notional       string                `json:"vv"`
		OpenInterest   string                `json:"oi"`
		UpdateTime     cryptodotcom.JSONTime `json:"t"`
	}

	//OrderBook struct {
	//	Arg    *events.Argument      `json:"arg"`
	//	Action string                `json:"action"`
	//	Books  []*market.OrderBookWs `json:"data"`
	//}
)
