package events

import (
	"github.com/drinkthere/cryptodotcom"
)

type (
	Basic struct {
		ID      cryptodotcom.JSONInt64 `json:"id,omitempty"`
		Code    cryptodotcom.JSONInt64 `json:"code"`
		Method  string                 `json:"method"`
		Message string                 `json:"message,omitempty"`
		Result  *Result                `json:"result,omitempty"`
	}

	Result struct {
		Subscription string `json:"subscription"`
		Channel      string `json:"channel"`
	}

	Login struct {
		ID      cryptodotcom.JSONInt64 `json:"id,omitempty"`
		Code    cryptodotcom.JSONInt64 `json:"code"`
		Method  string                 `json:"method"`
		Message string                 `json:"message,omitempty"`
	}

	HandleOrderResult struct {
		OrderID       string `json:"order_id"`
		ClientOrderID string `json:"client_oid"`
	}
)

//
//func (a *Argument) Get(k string) (interface{}, bool) {
//	v, ok := a.arg[k]
//	return v, ok
//}
//
//func (a *Argument) Set(k string, v interface{}) {
//	a.arg[k] = v
//}
//
//func (a *Argument) UnmarshalJSON(buf []byte) error {
//	a.arg = make(map[string]interface{})
//
//	if json.Unmarshal(buf, &a.arg) != nil {
//		return json.Unmarshal(buf, &a.untypedArg)
//	}
//
//	return nil
//}
