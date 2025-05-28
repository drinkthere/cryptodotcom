package ws

import (
	"encoding/json"
	"github.com/drinkthere/cryptodotcom"
	"github.com/drinkthere/cryptodotcom/events"
	"github.com/drinkthere/cryptodotcom/events/private"
	requests "github.com/drinkthere/cryptodotcom/requests/ws/private"
	"strings"
)

// Private
//
// https://www.okx.com/docs-v5/en/#websocket-api-private-channel
type Private struct {
	*ClientWs
	oCh chan *private.Order
	//aCh   chan *private.Balance
	//pCh   chan *private.Position
	//bnpCh chan *private.BalanceAndPosition
	//tCh   chan *private.Trade
}

// NewPrivate returns a pointer to a fresh Private
func NewPrivate(c *ClientWs) *Private {
	return &Private{ClientWs: c}
}

//
//// Account
//// Retrieve account information. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
////
//// https://www.okx.com/docs-v5/en/#websocket-api-private-channel-account-channel
//func (c *Private) Account(req requests.Account, ch ...chan *private.Account) error {
//	m := okx.S2M(req)
//	if len(ch) > 0 {
//		c.aCh = ch[0]
//	}
//	return c.Subscribe(true, []okx.ChannelName{"account"}, m)
//}
//
//// UAccount
////
//// https://www.okx.com/docs-v5/en/#websocket-api-private-channel-account-channel
//func (c *Private) UAccount(req requests.Account, rCh ...bool) error {
//	m := okx.S2M(req)
//	if len(rCh) > 0 && rCh[0] {
//		c.aCh = nil
//	}
//	return c.Unsubscribe(true, []okx.ChannelName{"account"}, m)
//}
//
//// Position
//// Retrieve position information. Initial snapshot will be pushed according to subscription granularity. Data will be pushed when triggered by events such as placing/canceling order, and will also be pushed in regular interval according to subscription granularity.
////
//// https://www.okx.com/docs-v5/en/#websocket-api-private-channel-positions-channel
//func (c *Private) Position(req requests.Position, ch ...chan *private.Position) error {
//	m := okx.S2M(req)
//	if len(ch) > 0 {
//		c.pCh = ch[0]
//	}
//	return c.Subscribe(true, []okx.ChannelName{"positions"}, m)
//}
//
//// UPosition
////
//// https://www.okx.com/docs-v5/en/#websocket-api-private-channel-positions-channel
//func (c *Private) UPosition(req requests.Position, rCh ...bool) error {
//	m := okx.S2M(req)
//	if len(rCh) > 0 && rCh[0] {
//		c.pCh = nil
//	}
//	return c.Unsubscribe(true, []okx.ChannelName{"positions"}, m)
//}
//
//// BalanceAndPosition
//// Retrieve account balance and position information. Data will be pushed when triggered by events such as filled order, funding transfer.
////
//// https://www.okx.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
//func (c *Private) BalanceAndPosition(ch ...chan *private.BalanceAndPosition) error {
//	m := make(map[string]string)
//	if len(ch) > 0 {
//		c.bnpCh = ch[0]
//	}
//	return c.Subscribe(true, []okx.ChannelName{"balance_and_position"}, m)
//}
//
//// UBalanceAndPosition unsubscribes a position channel
////
//// https://www.okx.com/docs-v5/en/#websocket-api-private-channel-balance-and-position-channel
//func (c *Private) UBalanceAndPosition(rCh ...bool) error {
//	m := make(map[string]string)
//	if len(rCh) > 0 && rCh[0] {
//		c.bnpCh = nil
//	}
//	return c.Unsubscribe(true, []okx.ChannelName{"balance_and_position"}, m)
//}

func (c *Private) Order(req requests.Order, ch chan *private.Order) error {
	c.oCh = ch
	channelNames := fillOrderChannel(req.InstrumentNames)
	return c.Subscribe(true, channelNames)
}

func (c *Private) Process(data []byte, e *events.Basic) bool {
	if e.Code == 0 {
		if e.Result != nil {
			ch := e.Result.Channel
			if strings.HasPrefix(ch, "user.order") {
				e := private.Order{}
				err := json.Unmarshal(data, &e)
				if err != nil {
					return false
				}
				if len(e.Result.Data) > 0 {
					go func() {
						if c.oCh != nil {
							c.oCh <- &e
						}
					}()
				}
				return true
			}

			//case "account":
			//	e := private.Account{}
			//	err := json.Unmarshal(data, &e)
			//	if err != nil {
			//		return false
			//	}
			//	go func() {
			//		if c.aCh != nil {
			//			c.aCh <- &e
			//		}
			//	}()
			//	return true
			//case "positions":
			//	e := private.Position{}
			//	err := json.Unmarshal(data, &e)
			//	if err != nil {
			//		return false
			//	}
			//	go func() {
			//		if c.pCh != nil {
			//			c.pCh <- &e
			//		}
			//	}()
			//	return true
			//case "balance_and_position":
			//	e := private.BalanceAndPosition{}
			//	err := json.Unmarshal(data, &e)
			//	if err != nil {
			//		return false
			//	}
			//	go func() {
			//		if c.bnpCh != nil {
			//			c.bnpCh <- &e
			//		}
			//	}()
			//	return true
		}
	}
	return false
}

func fillOrderChannel(instrumentNames []string) []string {
	prefix := cryptodotcom.ChannelPrefixOrder

	instLen := len(instrumentNames)
	if instLen > 0 {
		result := make([]string, len(instrumentNames))

		for i := range instrumentNames {
			result[i] = strings.Join([]string{string(prefix), instrumentNames[i]}, ".")
		}
		return result
	} else {
		return []string{string(prefix)}
	}
}
