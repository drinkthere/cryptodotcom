package public

import "github.com/drinkthere/cryptodotcom"

type (
	Tickers struct {
		InstrumentNames []string
	}

	OrderBooks struct {
		Instrument           []*Instrument
		BookSubscriptionType cryptodotcom.BookSubscriptionType `json:"book_subscription_type"`
		BookUpdateFrequency  cryptodotcom.BookUpdateFrequency  `json:"book_update_frequency"`
	}
	Instrument struct {
		InstID string `json:"instId"`
		Depth  string `json:"channel"`
	}
)
