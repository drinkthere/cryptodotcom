package trade

import (
	"github.com/drinkthere/cryptodotcom/models/trade"
	"github.com/drinkthere/cryptodotcom/responses"
)

type (
	CreateOrder struct {
		responses.Basic
		Result *trade.HandleOrderResult `json:"result"`
	}

	GetOpenOrdersResult struct {
		Data []*trade.Order `json:"data"`
	}
	GetOpenOrders struct {
		responses.Basic
		Result GetOpenOrdersResult `json:"result"`
	}

	CancelOrder struct {
		responses.Basic
		Result *trade.HandleOrderResult `json:"result"`
	}

	CancelAllOrders struct {
		responses.Basic
	}

	GetAccountSettings struct {
		responses.Basic
		Result []*trade.AccountSettings `json:"result"`
	}

	ChangeAccountLeverage struct {
		responses.Basic
	}
)
