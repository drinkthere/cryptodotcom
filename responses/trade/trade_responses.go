package trade

import (
	"github.com/drinkthere/cryptodotcom/models/trade"
	"github.com/drinkthere/cryptodotcom/responses"
)

type (
	CreateOrder struct {
		responses.Basic
		Result *trade.CreateOrderResult `json:"result"`
	}

	//CancelOrder struct {
	//	responses.Basic
	//	CancelOrders []*trade.CancelOrder `json:"data"`
	//}
	//
	//OrderList struct {
	//	responses.Basic
	//	Orders []*trade.Order `json:"data"`
	//}

)
