package cryptodotcom

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type (
	BaseURL        string
	InstrumentType string
	OrderType      string
	ExecInst       string
	TimeInForce    string
	OrderMode      string

	ContractType string
	PositionType string
	PositionSide string
	ActualSide   string
	TradeMode    string
	CountAction  string
	OrderSide    string
	GreekType    string
	BarSize      string
	TradeSide    string
	ChannelName  string
	Operation    string
	EventType    string

	AlgoOrderType        string
	QuantityType         string
	OrderFlowType        string
	OrderState           string
	ActionType           string
	APIKeyAccess         string
	OptionType           string
	AliasType            string
	InstrumentState      string
	DeliveryExerciseType string
	CandleStickWsBarSize string

	Destination           int
	FeeCategory           uint16
	TransferType          uint16
	AccountType           uint16
	DepositState          uint16
	WithdrawalDestination uint16
	WithdrawalState       int8

	ConvertType uint16

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

	ExecInstPostOnly = ExecInst("POST_ONLY")

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

	ContractLinearType  = ContractType("linear")
	ContractInverseType = ContractType("inverse")

	PositionLongShortMode = PositionType("long_short_mode")
	PositionNetMode       = PositionType("net_mode")

	PositionLongSide  = PositionSide("long")
	PositionShortSide = PositionSide("short")
	PositionNetSide   = PositionSide("net")

	TpSide = ActualSide("tp")
	SlSide = ActualSide("sl")

	TradeCrossMode    = TradeMode("cross")
	TradeIsolatedMode = TradeMode("isolated")
	TradeCashMode     = TradeMode("cash")

	CountIncrease = CountAction("add")
	CountDecrease = CountAction("reduce")

	LoginOperation            = Operation("login")
	SubscribeOperation        = Operation("subscribe")
	UnsubscribeOperation      = Operation("unsubscribe")
	OrderOperation            = Operation("order")
	BatchOrderOperation       = Operation("batch-orders")
	CancelOrderOperation      = Operation("cancel-order")
	BatchCancelOrderOperation = Operation("batch-cancel-orders")
	AmendOrderOperation       = Operation("amend-order")
	BatchAmendOrderOperation  = Operation("batch-amend-orders")

	// OrderFOK is Fill-or-kill order
	OrderFOK = OrderType("fok")
	// OrderIOC is Immediate-or-cancel order
	OrderIOC = OrderType("ioc")
	// OrderOptimalLimitIoc is Market order with immediate-or-cancel order
	OrderOptimalLimitIoc = OrderType("optimal_limit_ioc")
	// OrderMMP is Market Maker Protection (only applicable to Option in Portfolio Margin mode)
	OrderMMP = OrderType("mmp")
	// OrderMMPPostOnly is Market Maker Protection and Post-only order
	// (only applicable to Option in Portfolio Margin mode)
	OrderMMPPostOnly = OrderType("mmp_and_post_only")
	// OrderOPFOK is Simple options (fok)
	OrderOPFOK = OrderType("op_fok")

	// AlgoOrderConditional is One-way stop order
	AlgoOrderConditional = AlgoOrderType("conditional")
	// AlgoOrderOCO is One-cancels-the-other order
	AlgoOrderOCO      = AlgoOrderType("oco")
	AlgoOrderTrigger  = AlgoOrderType("trigger")
	AlgoOrderIceberg  = AlgoOrderType("iceberg")
	AlgoOrderTwap     = AlgoOrderType("twap")
	AlgoOrderTrailing = AlgoOrderType("move_order_stop")

	QuantityBaseCcy  = QuantityType("base_ccy")
	QuantityQuoteCcy = QuantityType("quote_ccy")

	OrderTakerFlow = OrderFlowType("T")
	OrderMakerFlow = OrderFlowType("M")

	ClassA = FeeCategory(1)
	ClassB = FeeCategory(2)
	ClassC = FeeCategory(3)
	ClassD = FeeCategory(4)

	OrderCancel          = OrderState("canceled")
	OrderPause           = OrderState("pause")
	OrderLive            = OrderState("live")
	OrderPartiallyFilled = OrderState("partially_filled")
	OrderFilled          = OrderState("filled")
	OrderUnfilled        = OrderState("unfilled")
	OrderEffective       = OrderState("effective")
	OrderFailed          = OrderState("order_failed")

	TransferWithinAccount     = TransferType(0)
	MasterAccountToSubAccount = TransferType(1)
	MasterSubAccountToAccount = TransferType(2)

	SpotAccount    = AccountType(1)
	FuturesAccount = AccountType(3)
	MarginAccount  = AccountType(5)
	FundingAccount = AccountType(6)
	SwapAccount    = AccountType(9)
	OptionsAccount = AccountType(12)
	UnifiedAccount = AccountType(18)

	WaitingForConfirmation     = DepositState(0)
	DepositCredited            = DepositState(1)
	DepositSuccessful          = DepositState(2)
	DepositTemporarySuspension = DepositState(8)

	WithdrawalokxDestination            = WithdrawalDestination(3)
	WithdrawalDigitalAddressDestination = WithdrawalDestination(4)

	WithdrawalPendingCancel              = WithdrawalState(-3)
	WithdrawalCanceled                   = WithdrawalState(-2)
	WithdrawalFailed                     = WithdrawalState(-1)
	WithdrawalPending                    = WithdrawalState(0)
	WithdrawalSending                    = WithdrawalState(1)
	WithdrawalSent                       = WithdrawalState(2)
	WithdrawalAwaitingEmailVerification  = WithdrawalState(3)
	WithdrawalAwaitingManualVerification = WithdrawalState(4)
	WithdrawalIdentityManualVerification = WithdrawalState(5)

	ActionPurchase = ActionType("purchase")
	ActionRedempt  = ActionType("redempt")

	APIKeyReadOnly = APIKeyAccess("read_only")
	APIKeyTrade    = APIKeyAccess("trade")
	APIKeyWithdraw = APIKeyAccess("withdraw")

	OptionCall = OptionType("C")
	OptionPut  = OptionType("P")

	AliasThisWeek    = AliasType("this_week")
	AliasNextWeek    = AliasType("next_week")
	AliasQuarter     = AliasType("quarter")
	AliasNextQuarter = AliasType("next_quarter")

	InstrumentLive    = InstrumentState("live")
	InstrumentSuspend = InstrumentState("suspend")
	InstrumentPreOpen = InstrumentState("preopen")

	Delivery   = DeliveryExerciseType("delivery")
	Exercise   = DeliveryExerciseType("exercised")
	ExpiredOtm = DeliveryExerciseType("expired_otm")

	CandleStick1Y  = CandleStickWsBarSize("candle1Y")
	CandleStick6M  = CandleStickWsBarSize("candle6M")
	CandleStick3M  = CandleStickWsBarSize("candle3M")
	CandleStick1M  = CandleStickWsBarSize("candle1M")
	CandleStick5D  = CandleStickWsBarSize("candle5D")
	CandleStick3D  = CandleStickWsBarSize("candle3D")
	CandleStick2D  = CandleStickWsBarSize("candle2D")
	CandleStick1D  = CandleStickWsBarSize("candle1D")
	CandleStick12H = CandleStickWsBarSize("candle12H")
	CandleStick6H  = CandleStickWsBarSize("candle6H")
	CandleStick4H  = CandleStickWsBarSize("candle4H")
	CandleStick2H  = CandleStickWsBarSize("candle2H")
	CandleStick1H  = CandleStickWsBarSize("candle1H")
	CandleStick30m = CandleStickWsBarSize("candle30m")
	CandleStick15m = CandleStickWsBarSize("candle15m")
	CandleStick5m  = CandleStickWsBarSize("candle5m")
	CandleStick3m  = CandleStickWsBarSize("candle3m")
	CandleStick1m  = CandleStickWsBarSize("candle1m")

	ConvertTypeContract = ConvertType(1)
	ConvertTypeCurrency = ConvertType(2)
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
func (t *WithdrawalState) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseInt(r, 10, 8)
	if err != nil {
		return err
	}
	*(*int8)(t) = int8(q)
	return
}

func (t *FeeCategory) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseUint(r, 10, 16)
	if err != nil {
		return err
	}
	*(*uint16)(t) = uint16(q)
	return
}
func (t *AccountType) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseUint(r, 10, 16)
	if err != nil {
		return err
	}
	*(*uint16)(t) = uint16(q)
	return
}
func (t *DepositState) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseUint(r, 10, 16)
	if err != nil {
		return err
	}
	*(*uint16)(t) = uint16(q)
	return
}

func S2M(i interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(i)
	_ = json.Unmarshal(j, &m)

	return m
}
