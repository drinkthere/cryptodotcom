package private

import "github.com/drinkthere/cryptodotcom"

type (
	Orders struct {
		InstrumentNames []string
	}

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

	CancelOrder struct {
		OrderID       string `json:"order_id,omitempty"`
		ClientOrderID string `json:"client_oid,omitempty"`
	}

	CancelAllOrders struct {
		InstrumentName string `json:"instrument_name,omitempty"` // BTCUSD-PERP
		Type           string `json:"type,omitempty"`            // LIMIT, TRIGGER, ALL
	}
)
