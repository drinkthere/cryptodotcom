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

	Order struct {
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
)
