package public

import (
	"github.com/drinkthere/cryptodotcom"
)

type (
	Basic struct {
		ID      cryptodotcom.JSONInt64 `json:"id,omitempty"`
		Code    cryptodotcom.JSONInt64 `json:"code"`
		Method  string                 `json:"method"`
		Message string                 `json:"message,omitempty"`
	}

	Tickers struct {
		Basic
		Result TickerResult `json:"result"`
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

	OrderBooks struct {
		Basic
		Result OrderBookResult `json:"result"`
	}
	OrderBookResult struct {
		InstrumentName string                 `json:"instrument_name"`
		Subscription   string                 `json:"subscription"`
		Channel        string                 `json:"channel"`
		Depth          cryptodotcom.JSONInt64 `json:"depth"`
		Data           []*OrderBook           `json:"data"`
	}
	OrderBook struct {
		Asks              [][]string             `json:"asks"`
		Bids              [][]string             `json:"bids"`
		PublishTime       cryptodotcom.JSONTime  `json:"t"`
		LastUpdateTime    cryptodotcom.JSONTime  `json:"tt"`
		UpdateSequence    cryptodotcom.JSONInt64 `json:"u"`
		PreUpdateSequence cryptodotcom.JSONInt64 `json:"pu,omitempty"`
	}
)
