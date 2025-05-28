package cryptodotcom

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type (
	Destination          int
	BaseURL              string
	InstrumentType       string
	OrderType            string
	OrderSide            string
	ExecInst             string
	TimeInForce          string
	OrderMode            string
	OrderStatus          string
	ChannelPrefix        string
	Operation            string
	BookSubscriptionType string
	BookUpdateFrequency  string

	QuantityType string

	JSONFloat64 float64
	JSONInt64   int64
	JSONTime    time.Time

	ClientError error
)

const (
	RestURL     = BaseURL("https://api.crypto.com/exchange/v1/")
	MarketWsURL = BaseURL("wss://stream.crypto.com/exchange/v1/market")
	UserWsURL   = BaseURL("wss://stream.crypto.com/exchange/v1/user")

	ColoRestURL     = BaseURL("https://api.crypto.com/exchange/v1/")
	ColoMarketWsURL = BaseURL("wss://stream.crypto.com/exchange/v1/market")
	ColoUserWsURL   = BaseURL("wss://stream.crypto.com/exchange/v1/user")

	NormalServer = Destination(iota + 1)
	ColoServer   = NormalServer + 1

	OrderBuy  = OrderSide("BUY")
	OrderSell = OrderSide("SELL")

	OrderMarket = OrderType("MARKET")
	OrderLimit  = OrderType("LIMIT")

	ExecInstPostOnly    = ExecInst("POST_ONLY")
	ExecInstLiquidation = ExecInst("LIQUIDATION")

	TimeInForceGTC = TimeInForce("GOOD_TILL_CANCEL")
	TimeInForceFOK = TimeInForce("FILL_OR_KILL")
	TimeInForceIOC = TimeInForce("IMMEDIATE_OR_CANCEL")

	SpotInstrument    = InstrumentType("SPOT")
	MarginInstrument  = InstrumentType("MARGIN")
	SwapInstrument    = InstrumentType("PERPETUAL_SWAP")
	FuturesInstrument = InstrumentType("FUTURES")
	OptionsInstrument = InstrumentType("OPTION")
	WarrantInstrument = InstrumentType("WARRANT")
	IndexInstrument   = InstrumentType("INDEX")

	SpotOrderMode   = OrderMode("SPOT")
	MarginOrderMode = OrderMode("MARGIN")

	ChannelPrefixTicker = ChannelPrefix("ticker")
	ChannelPrefixBook   = ChannelPrefix("book")

	PongOperation             = Operation("pong")
	LoginOperation            = Operation("login")
	SubscribeOperation        = Operation("subscribe")
	UnsubscribeOperation      = Operation("unsubscribe")
	OrderOperation            = Operation("order")
	BatchOrderOperation       = Operation("batch-orders")
	CancelOrderOperation      = Operation("cancel-order")
	BatchCancelOrderOperation = Operation("batch-cancel-orders")
	AmendOrderOperation       = Operation("amend-order")
	BatchAmendOrderOperation  = Operation("batch-amend-orders")

	QuantityBaseCcy  = QuantityType("base_ccy")
	QuantityQuoteCcy = QuantityType("quote_ccy")

	OrderNew     = OrderStatus("NEW")
	OrderPending = OrderStatus("PENDING")
	OrderActive  = OrderStatus("ACTIVE")

	BookSubscriptionTypeSNU = BookSubscriptionType("SNAPSHOT_AND_UPDATE")
	BookSubscriptionTypeS   = BookSubscriptionType("SNAPSHOT")

	BookUpdateFrequency10  = BookUpdateFrequency("10")
	BookUpdateFrequency100 = BookUpdateFrequency("100")
	BookUpdateFrequency500 = BookUpdateFrequency("500")

	//OrderCancel = OrderState("canceled")
	//OrderPause  = OrderState("pause")
	//OrderPartiallyFilled = OrderState("partially_filled")
	//OrderFilled          = OrderState("filled")
	//OrderUnfilled        = OrderState("unfilled")
	//OrderEffective       = OrderState("effective")
	//OrderFailed          = OrderState("order_failed")
)

func (t JSONTime) String() string { return time.Time(t).String() }

func (t *JSONTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.UnixMilli(q)
	return
}
func (t *JSONFloat64) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseFloat(r, 64)
	if err != nil {
		return err
	}
	*(*float64)(t) = q
	return
}
func (t *JSONInt64) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*int64)(t) = q
	return
}

func S2M(i interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(i)
	_ = json.Unmarshal(j, &m)

	return m
}
