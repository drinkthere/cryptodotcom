package account

import (
	"github.com/drinkthere/cryptodotcom/models/account"
	"github.com/drinkthere/cryptodotcom/responses"
)

type (
	GetAccountsResult struct {
		MasterAccount  *account.Account   `json:"master_account"`
		SubAccountList []*account.Account `json:"sub_account_list"`
	}
	GetAccounts struct {
		responses.Basic
		Result GetAccountsResult `json:"result"`
	}

	GetBalancesResult struct {
		Data []*account.Balance `json:"data"`
	}
	GetBalances struct {
		responses.Basic
		Result GetBalancesResult `json:"result"`
	}

	GetPositionsResult struct {
		Data []*account.Position `json:"data"`
	}
	GetPositions struct {
		responses.Basic
		Result GetPositionsResult `json:"result"`
	}
)
