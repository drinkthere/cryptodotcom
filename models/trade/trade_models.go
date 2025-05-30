package trade

import "github.com/drinkthere/cryptodotcom"

type (
	HandleOrderResult struct {
		OrderID       string `json:"order_id"`
		ClientOrderID string `json:"client_oid"`
	}

	Order struct {
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
		CreateTime         cryptodotcom.JSONTime    `json:"create_time"`
		CreateTimeNS       string                   `json:"create_time_ns"`
		UpdateTime         cryptodotcom.JSONTime    `json:"update_time"`
		InstrumentName     string                   `json:"instrument_name"`
		FeeInstrumentName  string                   `json:"fee_instrument_name"`
	}

	AccountSettings struct {
		Leverage cryptodotcom.JSONInt64 `json:"leverage"`
		StpId    cryptodotcom.JSONInt64 `json:"stp_id"`
		StpScope string                 `json:"stp_scope"`
		StpInst  string                 `json:"stp_inst"`
	}

	FeeRate struct {
		SpotTier                   string `json:"spot_tier"`
		DerivTier                  string `json:"deriv_tier"`
		EffectiveSpotMakerRateBps  string `json:"effective_spot_maker_rate_bps"`
		EffectiveSpotTakerRateBps  string `json:"effective_spot_taker_rate_bps"`
		EffectiveDerivMakerRateBps string `json:"effective_deriv_maker_rate_bps"`
		EffectiveDerivTakerRateBps string `json:"effective_deriv_taker_rate_bps"`
	}
)
