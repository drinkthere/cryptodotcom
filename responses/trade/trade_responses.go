package trade

import (
	"github.com/drinkthere/cryptodotcom/models/trade"
	"github.com/drinkthere/cryptodotcom/responses"
)

type (
	CreateOrder struct {
		responses.Basic
		PlaceOrders []*trade.CreateOrder `json:"data"`
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
