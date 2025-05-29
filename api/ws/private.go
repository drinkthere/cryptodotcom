package ws

import (
	"encoding/json"
	"fmt"
	"github.com/drinkthere/cryptodotcom"
	"github.com/drinkthere/cryptodotcom/events"
	"github.com/drinkthere/cryptodotcom/events/private"
	requests "github.com/drinkthere/cryptodotcom/requests/ws/private"
	"github.com/drinkthere/cryptodotcom/utils"
	"strings"
)

// Private
//
// https://www.okx.com/docs-v5/en/#websocket-api-private-channel
type Private struct {
	*ClientWs
	oCh   chan *private.Orders
	bCh   chan *private.Balances
	pCh   chan *private.Positions
	pnbCh chan *private.PositionAndBalance
}

// NewPrivate returns a pointer to a fresh Private
func NewPrivate(c *ClientWs) *Private {
	return &Private{ClientWs: c}
}

func (c *Private) Orders(req requests.Orders, ch chan *private.Orders) error {
	c.oCh = ch
	channelNames := fillOrderChannel(req.InstrumentNames)
	return c.Subscribe(true, channelNames)
}

func (c *Private) Balances(ch chan *private.Balances) error {
	c.bCh = ch
	channelNames := []string{"user.balance"}
	return c.Subscribe(true, channelNames)
}

func (c *Private) Positions(ch chan *private.Positions) error {
	c.pCh = ch
	channelNames := []string{"user.positions"}
	return c.Subscribe(true, channelNames)
}

func (c *Private) PositionAndBalance(ch chan *private.PositionAndBalance) error {
	c.pnbCh = ch
	channelNames := []string{"user.position_balance"}
	return c.Subscribe(true, channelNames)
}

func (c *Private) CreateOrder(order requests.CreateOrder) error {
	if c.TradeChan == nil {
		return fmt.Errorf("trade channel has not been initialized")
	}
	args := map[string]interface{}{
		"createOrder": order,
	}
	return c.Send(true, utils.GenerateRequestID(), cryptodotcom.CreateOrderOperation, args)
}

func (c *Private) Process(data []byte, e *events.Basic) bool {
	if e.Code == 0 {
		if e.Result != nil {
			ch := e.Result.Channel
			// why not use switch, because order channel includes "user.order and use.order.INSRUMENT_NAME" which need to compare the prefix
			if ch == "user.balance" {
				e := private.Balances{}
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
			} else if ch == "user.positions" {
				e := private.Positions{}
				err := json.Unmarshal(data, &e)
				if err != nil {
					return false
				}
				if len(e.Result.Data) > 0 {
					go func() {
						if c.pCh != nil {
							c.pCh <- &e
						}
					}()
				}
				return true
			} else if ch == "user.position_balance" {
				e := private.PositionAndBalance{}
				err := json.Unmarshal(data, &e)
				if err != nil {
					return false
				}
				if len(e.Result.Data) > 0 {
					go func() {
						if c.pnbCh != nil {
							c.pnbCh <- &e
						}
					}()
				}
				return true
			} else if strings.HasPrefix(ch, "user.order") {
				e := private.Orders{}
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
