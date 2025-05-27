package account

type (
	GetAccounts struct {
		PageSize string `json:"page_size,omitempty,string"`
		Page     string `json:"page,omitempty,string"`
	}

	GetPositions struct {
		InstrumentName string `json:"instrument_name,omitempty"`
	}

	//SetAccountLevel struct {
	//	AcctLv string `json:"acctLv"`
	//}
)
