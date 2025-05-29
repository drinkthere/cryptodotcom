package private

import (
	"github.com/drinkthere/cryptodotcom"
)

type (
	Basic struct {
		ID     cryptodotcom.JSONInt64 `json:"id,omitempty"`
		Code   cryptodotcom.JSONInt64 `json:"code"`
		Method string                 `json:"method"`
	}

	Orders struct {
		Basic
		Result OrderResult `json:"result"`
	}
	OrderResult struct {
		InstrumentName string         `json:"instrument_name"`
		Subscription   string         `json:"subscription"`
		Channel        string         `json:"channel"`
		Data           []*OrderDetail `json:"data"`
	}
	OrderDetail struct {
		AccountID          string                   `json:"account_id"`
		OrderID            string                   `json:"order_id"`
		ClientOrderID      string                   `json:"client_oid"`
		OrderType          cryptodotcom.OrderType   `json:"order_type"`
		TimeInForce        cryptodotcom.TimeInForce `json:"time_in_force"`
		Side               cryptodotcom.OrderSide   `json:"side"`
		ExecInst           []cryptodotcom.ExecInst  `json:"exec_inst"`
		Quantity           string                   `json:"quantity"`
		LimitPrice         string                   `json:"limit_price"`
		OrderValue         string                   `json:"order_value"`
		MakerFeeRate       string                   `json:"maker_fee_rate"`
		TakerFeeRate       string                   `json:"taker_fee_rate"`
		AvgPrice           string                   `json:"avg_price"`
		CumulativeQuantity string                   `json:"cumulative_quantity"`
		CumulativeValue    string                   `json:"cumulative_value"`
		CumulativeFee      string                   `json:"cumulative_fee"`
		Status             cryptodotcom.OrderStatus `json:"status"`
		UpdateUserID       string                   `json:"update_user_id"`
		OrderDate          string                   `json:"order_date"`
		InstrumentName     string                   `json:"instrument_name"`
		FeeInstrumentName  string                   `json:"fee_instrument_name"`
		CreateTime         cryptodotcom.JSONTime    `json:"create_time"`
		CreateTimeNs       string                   `json:"create_time_ns"`
		UpdateTime         cryptodotcom.JSONTime    `json:"update_time"`
	}

	Balances struct {
		Basic
		Result BalanceResult `json:"result"`
	}
	BalanceResult struct {
		Subscription string           `json:"subscription"`
		Channel      string           `json:"channel"`
		Data         []*BalanceDetail `json:"data"`
	}
	BalanceDetail struct {
		TotalAvailableBalance     string             `json:"total_available_balance"`
		TotalMarginBalance        string             `json:"total_margin_balance"`
		TotalInitialMargin        string             `json:"total_initial_margin"`
		TotalPositionIm           string             `json:"total_position_im"`
		TotalHaircut              string             `json:"total_haircut"`
		TotalMaintenanceMargin    string             `json:"total_maintenance_margin"`
		TotalPositionCost         string             `json:"total_position_cost"`
		TotalCashBalance          string             `json:"total_cash_balance"`
		TotalCollateralValue      string             `json:"total_collateral_value"`
		TotalSessionUnrealizedPnl string             `json:"total_session_unrealized_pnl"`
		InstrumentName            string             `json:"instrument_name"`
		TotalSessionRealizedPnl   string             `json:"total_session_realized_pnl"`
		IsLiquidating             bool               `json:"is_liquidating"`
		TotalEffectiveLeverage    string             `json:"total_effective_leverage"`
		PositionLimit             string             `json:"position_limit"`
		UsedPositionLimit         string             `json:"used_position_limit"`
		PositionBalances          []*PositionBalance `json:"position_balances"`
	}
	PositionBalance struct {
		InstrumentName       string `json:"instrument_name"`
		Quantity             string `json:"quantity"`
		MarketValue          string `json:"market_value"`
		CollateralEligible   bool   `json:"collateral_eligible"`
		Haircut              string `json:"haircut"`
		CollateralAmount     string `json:"collateral_amount"`
		MaxWithdrawalBalance string `json:"max_withdrawal_balance"`
		ReservedQty          string `json:"reserved_qty"`
	}

	Positions struct {
		Basic
		Result PositionResult `json:"result"`
	}
	PositionResult struct {
		Subscription string            `json:"subscription"`
		Channel      string            `json:"channel"`
		Data         []*PositionDetail `json:"data"`
	}
	PositionDetail struct {
		AccountId            string                 `json:"account_id"`
		Quantity             string                 `json:"quantity"`
		SessionUnrealizedPnl string                 `json:"session_unrealized_pnl"`
		Cost                 string                 `json:"cost"`
		OpenPositionPnl      string                 `json:"open_position_pnl"`
		OpenPosCost          string                 `json:"open_pos_cost"`
		SessionPnl           string                 `json:"session_pnl"`
		PosInitialMargin     string                 `json:"pos_initial_margin"`
		PosMaintenanceMargin string                 `json:"pos_maintenance_margin"`
		MarketValue          string                 `json:"market_value"`
		MarkPrice            string                 `json:"mark_price"`
		TargetLeverage       string                 `json:"target_leverage"`
		UpdateTimestampMs    cryptodotcom.JSONInt64 `json:"update_timestamp_ms"`
		InstrumentName       string                 `json:"instrument_name"`
		Type                 string                 `json:"type"`
	}

	PositionAndBalance struct {
		Basic
		Result PositionAndBalanceResult `json:"result"`
	}

	PositionAndBalanceResult struct {
		Subscription string                      `json:"subscription"`
		Channel      string                      `json:"channel"`
		Data         []*PositionAndBalanceDetail `json:"data"`
	}

	PositionAndBalanceDetail struct {
		Balances  []*Bal `json:"balances"`
		Positions []*Pos `json:"positions"`
	}

	Bal struct {
		InstrumentName string `json:"instrument_name"`
		Quantity       string `json:"quantity"`
	}

	Pos struct {
		AccountId         string                 `json:"account_id"`
		InstrumentName    string                 `json:"instrument_name"`
		Type              string                 `json:"type"`
		Quantity          string                 `json:"quantity"`
		Cost              string                 `json:"cost"`
		OpenPositionPnl   string                 `json:"open_position_pnl"`
		SessionPnl        string                 `json:"session_pnl"`
		UpdateTimestampMs cryptodotcom.JSONInt64 `json:"update_timestamp_ms"`
		OpenPosCost       string                 `json:"open_pos_cost"`
	}
)
