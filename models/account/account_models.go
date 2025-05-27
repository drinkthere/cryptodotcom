package account

import "github.com/drinkthere/cryptodotcom"

type (
	Account struct {
		UUID              string                `json:"uuid"`
		MasterAccountUUID string                `json:"master_account_uuid"`
		MarginAccountUUID string                `json:"margin_account_uuid"`
		Label             string                `json:"label"`
		Enabled           bool                  `json:"enabled"`
		Tradable          bool                  `json:"tradable"`
		Name              string                `json:"name"`
		Email             string                `json:"email"`
		MobileNumber      string                `json:"mobile_number"`
		CountryCode       string                `json:"country_code"`
		Address           string                `json:"address"`
		MarginAccess      string                `json:"margin_access"`
		DerivativesAccess string                `json:"derivatives_access"`
		CreateTime        cryptodotcom.JSONTime `json:"create_time"`
		UpdateTime        cryptodotcom.JSONTime `json:"update_time"`
		TwoFaEnabled      bool                  `json:"two_fa_enabled"`
		KycLevel          string                `json:"kyc_level"`
		Suspended         bool                  `json:"suspended"`
		Terminated        bool                  `json:"terminated"`
	}

	Balance struct {
		InstrumentName            string            `json:"instrument_name"`
		TotalAvailableBalance     string            `json:"total_available_balance"`
		TotalMarginBalance        string            `json:"total_margin_balance"`
		TotalInitialMargin        string            `json:"total_initial_margin"`
		TotalPositionIm           string            `json:"total_position_im"`
		TotalHaircut              string            `json:"total_haircut"`
		TotalMaintenanceMargin    string            `json:"total_maintenance_margin"`
		TotalPositionCost         string            `json:"total_position_cost"`
		TotalCashBalance          string            `json:"total_cash_balance"`
		TotalCollateralValue      string            `json:"total_collateral_value"`
		TotalSessionUnrealizedPnl string            `json:"total_session_unrealized_pnl"`
		TotalSessionRealizedPnl   string            `json:"total_session_realized_pnl"`
		IsLiquidating             bool              `json:"is_liquidating"`
		TotalEffectiveLeverage    string            `json:"total_effective_leverage"`
		PositionLimit             string            `json:"position_limit"`
		UsedPositionLimit         string            `json:"used_position_limit"`
		PositionBalances          []*BalanceDetails `json:"position_balances"`
	}

	BalanceDetails struct {
		InstrumentName       string `json:"instrument_name"`
		Quantity             string `json:"quantity"`
		MarketValue          string `json:"market_value"`
		CollateralEligible   bool   `json:"collateral_eligible"`
		Haircut              string `json:"haircut"`
		CollateralAmount     string `json:"collateral_amount"`
		MaxWithdrawalBalance string `json:"max_withdrawal_balance"`
		ReservedQty          string `json:"reserved_qty"`
	}

	Position struct {
		InstrumentName    string                `json:"instrument_name"`
		Type              string                `json:"type"`
		Quantity          string                `json:"quantity"`
		Cost              string                `json:"cost"`
		OpenPositionPnl   string                `json:"open_position_pnl"`
		OpenPosCost       string                `json:"open_pos_cost"`
		SessionPnl        string                `json:"session_pnl"`
		UpdateTimestampMs cryptodotcom.JSONTime `json:"update_timestamp_ms"`
	}
	//BalanceAndPosition struct {
	//	EventType okx.EventType     `json:"eventType"`
	//	PTime     okx.JSONTime      `json:"pTime"`
	//	UTime     okx.JSONTime      `json:"uTime"`
	//	PosData   []*Position       `json:"posData"`
	//	BalData   []*BalanceDetails `json:"balData"`
	//}
	//PositionAndAccountRisk struct {
	//	AdjEq   okx.JSONFloat64                      `json:"adjEq,omitempty"`
	//	BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
	//	PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
	//	TS      okx.JSONTime                         `json:"ts"`
	//}
	//PositionAndAccountRiskBalanceData struct {
	//	Ccy   string          `json:"ccy"`
	//	Eq    okx.JSONFloat64 `json:"eq"`
	//	DisEq okx.JSONFloat64 `json:"disEq"`
	//}
	//PositionAndAccountRiskPositionData struct {
	//	InstID      string             `json:"instId"`
	//	PosCcy      string             `json:"posCcy,omitempty"`
	//	Ccy         string             `json:"ccy"`
	//	NotionalCcy okx.JSONFloat64    `json:"notionalCcy"`
	//	Pos         okx.JSONFloat64    `json:"pos"`
	//	NotionalUsd okx.JSONFloat64    `json:"notionalUsd"`
	//	PosSide     okx.PositionSide   `json:"posSide"`
	//	InstType    okx.InstrumentType `json:"instType"`
	//	MgnMode     okx.MarginMode     `json:"mgnMode"`
	//}
	//Bill struct {
	//	Ccy       string             `json:"ccy"`
	//	InstID    string             `json:"instId"`
	//	Notes     string             `json:"notes"`
	//	BillID    string             `json:"billId"`
	//	OrdID     string             `json:"ordId"`
	//	BalChg    okx.JSONFloat64    `json:"balChg"`
	//	PosBalChg okx.JSONFloat64    `json:"posBalChg"`
	//	Bal       okx.JSONFloat64    `json:"bal"`
	//	PosBal    okx.JSONFloat64    `json:"posBal"`
	//	Sz        okx.JSONFloat64    `json:"sz"`
	//	Pnl       okx.JSONFloat64    `json:"pnl"`
	//	Fee       okx.JSONFloat64    `json:"fee"`
	//	From      okx.AccountType    `json:"from,omitempty"`
	//	To        okx.AccountType    `json:"to,omitempty"`
	//	InstType  okx.InstrumentType `json:"instType"`
	//	MgnMode   okx.MarginMode     `json:"MgnMode"`
	//	Type      okx.BillType       `json:"type,string"`
	//	SubType   okx.BillSubType    `json:"subType,string"`
	//	TS        okx.JSONTime       `json:"ts"`
	//}
	//Config struct {
	//	Level       string           `json:"level"`
	//	LevelTmp    string           `json:"levelTmp"`
	//	AcctLv      string           `json:"acctLv"`
	//	AutoLoan    bool             `json:"autoLoan"`
	//	UID         string           `json:"uid"`
	//	MainUID     string           `json:"mainUid"`
	//	Label       string           `json:"label"`
	//	IP          string           `json:"ip"`
	//	Permissions string           `json:"perm"`
	//	GreeksType  okx.GreekType    `json:"greeksType"`
	//	PosMode     okx.PositionType `json:"posMode"`
	//}
	//PositionMode struct {
	//	PosMode okx.PositionType `json:"posMode"`
	//}
	//Leverage struct {
	//	InstID  string           `json:"instId"`
	//	Lever   okx.JSONFloat64  `json:"lever"`
	//	MgnMode okx.MarginMode   `json:"mgnMode"`
	//	PosSide okx.PositionSide `json:"posSide"`
	//}
	//MaxBuySellAmount struct {
	//	InstID  string          `json:"instId"`
	//	Ccy     string          `json:"ccy"`
	//	MaxBuy  okx.JSONFloat64 `json:"maxBuy"`
	//	MaxSell okx.JSONFloat64 `json:"maxSell"`
	//}
	//MaxAvailableTradeAmount struct {
	//	InstID    string          `json:"instId"`
	//	AvailBuy  okx.JSONFloat64 `json:"availBuy"`
	//	AvailSell okx.JSONFloat64 `json:"availSell"`
	//}
	//MarginBalanceAmount struct {
	//	InstID  string           `json:"instId"`
	//	Amt     okx.JSONFloat64  `json:"amt"`
	//	PosSide okx.PositionSide `json:"posSide,string"`
	//	Type    okx.CountAction  `json:"type,string"`
	//}
	//Loan struct {
	//	InstID  string          `json:"instId"`
	//	MgnCcy  string          `json:"mgnCcy"`
	//	Ccy     string          `json:"ccy"`
	//	MaxLoan okx.JSONFloat64 `json:"maxLoan"`
	//	MgnMode okx.MarginMode  `json:"mgnMode"`
	//	Side    okx.OrderSide   `json:"side"`
	//}
	//Fee struct {
	//	Level     string             `json:"level"`
	//	Taker     okx.JSONFloat64    `json:"taker"`
	//	Maker     okx.JSONFloat64    `json:"maker"`
	//	Delivery  okx.JSONFloat64    `json:"delivery,omitempty"`
	//	Exercise  okx.JSONFloat64    `json:"exercise,omitempty"`
	//	Category  okx.FeeCategory    `json:"category,string"`
	//	InstType  okx.InstrumentType `json:"instType"`
	//	TakerU    okx.JSONFloat64    `json:"takerU"`
	//	MakerU    okx.JSONFloat64    `json:"MakerU"`
	//	TakerUSDC okx.JSONFloat64    `json:"takerUSDC"`
	//	MakerUSDC okx.JSONFloat64    `json:"makerUSDC"`
	//	TS        okx.JSONTime       `json:"ts"`
	//}
	//InterestAccrued struct {
	//	InstID       string          `json:"instId"`
	//	Ccy          string          `json:"ccy"`
	//	Interest     okx.JSONFloat64 `json:"interest"`
	//	InterestRate okx.JSONFloat64 `json:"interestRate"`
	//	Liab         okx.JSONFloat64 `json:"liab"`
	//	MgnMode      okx.MarginMode  `json:"mgnMode"`
	//	TS           okx.JSONTime    `json:"ts"`
	//}
	//InterestRate struct {
	//	Ccy          string          `json:"ccy"`
	//	InterestRate okx.JSONFloat64 `json:"interestRate"`
	//}
	//Greek struct {
	//	GreeksType string `json:"greeksType"`
	//}
	//MaxWithdrawal struct {
	//	Ccy   string          `json:"ccy"`
	//	MaxWd okx.JSONFloat64 `json:"maxWd"`
	//}
	//AutoLoan struct {
	//	AutoLoan bool `json:"autoLoan"`
	//}
	//AcctLevel struct {
	//	AcctLv string `json:"acctLv"`
	//}
	//InterestLimitsRecordDetail struct {
	//	AllAcctRemainingQuota okx.JSONFloat64 `json:"allAcctRemainingQuota"`
	//	CurAcctRemainingQuota okx.JSONFloat64 `json:"curAcctRemainingQuota,omitempty"`
	//	PlatRemainingQuota    string          `json:"platRemainingQuota"`
	//}
	//InterestLimitsRecord struct {
	//	Ccy               string                     `json:"ccy"`
	//	Rate              okx.JSONFloat64            `json:"rate"`
	//	LoanQuota         okx.JSONFloat64            `json:"loanQuota"`
	//	SurplusLmt        okx.JSONFloat64            `json:"surplusLmt"`
	//	SurplusLmtDetails InterestLimitsRecordDetail `json:"surplusLmtDetails,omitempty"`
	//	UsedLmt           okx.JSONFloat64            `json:"usedLmt"`
	//	Interest          okx.JSONFloat64            `json:"interest,omitempty"`
	//	PosLoan           okx.JSONFloat64            `json:"posLoan,omitempty"`
	//	AvailLoan         okx.JSONFloat64            `json:"availLoan,omitempty"`
	//	UsedLoan          okx.JSONFloat64            `json:"usedLoan,omitempty"`
	//	AvgRate           okx.JSONFloat64            `json:"avgRate,omitempty"`
	//}
	//InterestLimits struct {
	//	Debt             okx.JSONFloat64         `json:"debt"`
	//	Interest         okx.JSONFloat64         `json:"interest,omitempty"`
	//	NextDiscountTime okx.JSONTime            `json:"nextDiscountTime"`
	//	NextInterestTime okx.JSONTime            `json:"nextInterestTime"`
	//	LoanAlloc        string                  `json:"loanAlloc"`
	//	Records          []*InterestLimitsRecord `json:"records"`
	//}
)
