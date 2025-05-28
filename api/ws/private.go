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
	bCh chan *private.Balance
	//pCh   chan *private.Position
	//bnpCh chan *private.BalanceAndPosition
	//tCh   chan *private.Trade
}

// NewPrivate returns a pointer to a fresh Private
func NewPrivate(c *ClientWs) *Private {
	return &Private{ClientWs: c}
}

func (c *Private) Order(req requests.Order, ch chan *private.Order) error {
	c.oCh = ch
	channelNames := fillOrderChannel(req.InstrumentNames)
	return c.Subscribe(true, channelNames)
}

func (c *Private) Balance(ch chan *private.Balance) error {
	c.bCh = ch
	channelNames := []string{"user.balance"}
	return c.Subscribe(true, channelNames)
}

func (c *Private) Process(data []byte, e *events.Basic) bool {
	if e.Code == 0 {
		if e.Result != nil {
			ch := e.Result.Channel
			if ch == "user.balance" {
				e := private.Balance{}
				err := json.Unmarshal(data, &e)
				if err != nil {
					return false
				}
				if len(e.Result.Data) > 0 {
					go func() {
						if c.bCh != nil {
							c.bCh <- &e
						}
					}()
				}
				return true
			} else if strings.HasPrefix(ch, "user.order") {
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
