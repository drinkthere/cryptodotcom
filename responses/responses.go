package responses

type (
	Basic struct {
		ID      int64  `json:"id,omitempty"`
		Code    int    `json:"code"`
		Method  string `json:"method,string"`
		Message string `json:"message,omitempty"`
	}
)
