package public

import "github.com/drinkthere/cryptodotcom"

type (
	Tickers struct {
		InstrumentNames []string
	}

	OrderBooks struct {
		Instrument           []*OrderBook
		BookSubscriptionType cryptodotcom.BookSubscriptionType `json:"book_subscription_type"`
		BookUpdateFrequency  cryptodotcom.BookUpdateFrequency  `json:"book_update_frequency"`
	}
	OrderBook struct {
		InstID string `json:"instId"`
		Depth  string `json:"channel"`
	}
)
