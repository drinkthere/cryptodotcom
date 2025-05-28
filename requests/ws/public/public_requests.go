package public

import "github.com/drinkthere/cryptodotcom"

type (
	Tickers struct {
		InstrumentNames []string
	}

	OrderBooks struct {
		Instruments          []*Instrument
		BookSubscriptionType cryptodotcom.BookSubscriptionType `json:"book_subscription_type"`
		BookUpdateFrequency  cryptodotcom.BookUpdateFrequency  `json:"book_update_frequency"`
	}
	Instrument struct {
		InstrumentName string `json:"instrument_name"`
		Depth          string `json:"depth"`
	}
)
