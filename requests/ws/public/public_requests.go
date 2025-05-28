package public

type (
	Tickers struct {
		InstrumentNames []string
	}
	OrderBook struct {
		InstID  string `json:"instId"`
		Channel string `json:"channel"`
	}
)
