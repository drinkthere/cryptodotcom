package trade

import "github.com/drinkthere/cryptodotcom"

type (
	CreateOrder struct {
		InstrumentName    string                   `json:"instrument_name"`
		Side              cryptodotcom.OrderSide   `json:"side"`
		Type              cryptodotcom.OrderType   `json:"type"`
		Price             string                   `json:"price"`
		Quantity          string                   `json:"quantity"`
		Notional          string                   `json:"notional,omitempty"`
		ClientOrderID     string                   `json:"client_oid,omitempty"`
		ExecInst          []cryptodotcom.ExecInst  `json:"exec_inst,omitempty"`
		TimeInForce       cryptodotcom.TimeInForce `json:"time_in_force,omitempty"`
		SpotMargin        cryptodotcom.OrderMode   `json:"spot_margin,omitempty"`
		FeeInstrumentName string                   `json:"fee_instrument_name,omitempty"`
	}
	//CancelOrder struct {
	//	ID      string `json:"-"`
	//	InstID  string `json:"instId"`
	//	OrdID   string `json:"ordId,omitempty"`
	//	ClOrdID string `json:"clOrdId,omitempty"`
	//}
	//
	//OrderList struct {
	//	Uly      string             `json:"uly,omitempty"`
	//	InstID   string             `json:"instId,omitempty"`
	//	After    float64            `json:"after,omitempty,string"`
	//	Before   float64            `json:"before,omitempty,string"`
	//	Limit    float64            `json:"limit,omitempty,string"`
	//	InstType okx.InstrumentType `json:"instType,omitempty"`
	//	OrdType  okx.OrderType      `json:"ordType,omitempty"`
	//	State    okx.OrderState     `json:"state,omitempty"`
	//}
)
