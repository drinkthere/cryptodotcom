package responses

import "github.com/drinkthere/cryptodotcom"

type (
	Basic struct {
		ID      cryptodotcom.JSONInt64 `json:"id,omitempty"`
		Code    cryptodotcom.JSONInt64 `json:"code"`
		Method  string                 `json:"method"`
		Msg     string                 `json:"msg,omitempty"`
		Message string                 `json:"message,omitempty"`
	}
)
