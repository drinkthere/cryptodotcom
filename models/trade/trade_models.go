package trade

type (
	CreateOrderResult struct {
		OrderID       string `json:"order_id"`
		ClientOrderID string `json:"client_oid"`
	}
	//CancelOrder struct {
	//	OrdID   string          `json:"ordId"`
	//	ClOrdID string          `json:"clOrdId"`
	//	SMsg    string          `json:"sMsg"`
	//	SCode   okx.JSONFloat64 `json:"sCode"`
	//}
	//
	//Order struct {
	//	InstID       string             `json:"instId"`
	//	Ccy          string             `json:"ccy"`
	//	OrdID        string             `json:"ordId"`
	//	AlgoID       string             `json:"algoId"`
	//	ClOrdID      string             `json:"clOrdId"`
	//	AlgoClOrdID  string             `json:"algoClOrdId"`
	//	TradeID      string             `json:"tradeId"`
	//	Tag          string             `json:"tag"`
	//	Category     string             `json:"category"`
	//	FeeCcy       string             `json:"feeCcy"`
	//	RebateCcy    string             `json:"rebateCcy"`
	//	QuickMgnType string             `json:"quickMgnType"`
	//	ReduceOnly   string             `json:"reduceOnly"`
	//	Px           okx.JSONFloat64    `json:"px"`
	//	Sz           okx.JSONFloat64    `json:"sz"`
	//	Pnl          okx.JSONFloat64    `json:"pnl"`
	//	AccFillSz    okx.JSONFloat64    `json:"accFillSz"`
	//	FillPx       okx.JSONFloat64    `json:"fillPx"`
	//	FillSz       okx.JSONFloat64    `json:"fillSz"`
	//	FillTime     okx.JSONFloat64    `json:"fillTime"`
	//	AvgPx        okx.JSONFloat64    `json:"avgPx"`
	//	Lever        okx.JSONFloat64    `json:"lever"`
	//	TpTriggerPx  okx.JSONFloat64    `json:"tpTriggerPx"`
	//	TpOrdPx      okx.JSONFloat64    `json:"tpOrdPx"`
	//	SlTriggerPx  okx.JSONFloat64    `json:"slTriggerPx"`
	//	SlOrdPx      okx.JSONFloat64    `json:"slOrdPx"`
	//	Fee          okx.JSONFloat64    `json:"fee"`
	//	Rebate       okx.JSONFloat64    `json:"rebate"`
	//	State        okx.OrderState     `json:"state"`
	//	TdMode       okx.TradeMode      `json:"tdMode"`
	//	PosSide      okx.PositionSide   `json:"posSide"`
	//	Side         okx.OrderSide      `json:"side"`
	//	OrdType      okx.OrderType      `json:"ordType"`
	//	InstType     okx.InstrumentType `json:"instType"`
	//	TgtCcy       okx.QuantityType   `json:"tgtCcy"`
	//	UTime        okx.JSONTime       `json:"uTime"`
	//	CTime        okx.JSONTime       `json:"cTime"`
	//}
)
