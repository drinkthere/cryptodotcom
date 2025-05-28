package ws

import (
	"encoding/json"
	"github.com/drinkthere/cryptodotcom"
	"github.com/drinkthere/cryptodotcom/events"
	"github.com/drinkthere/cryptodotcom/events/public"
	requests "github.com/drinkthere/cryptodotcom/requests/ws/public"
)

type Public struct {
	*ClientWs
	tCh chan *public.Tickers
	//obCh chan *public.OrderBook
}

// NewPublic returns a pointer to a fresh Public
func NewPublic(c *ClientWs) *Public {
	return &Public{ClientWs: c}
}

func (c *Public) Tickers(req requests.Tickers) error {
	channelNames := cryptodotcom.FillChannelNames(cryptodotcom.ChannelPrefixTicker, req.InstrumentNames)
	return c.Subscribe(false, channelNames)
}

//
//func (c *Public) OrderBook(req requests.OrderBook, ch ...chan *public.OrderBook) error {
//	m := okx.S2M(req)
//	if len(ch) > 0 {
//		c.obCh = ch[0]
//	}
//	return c.Subscribe(false, []okx.ChannelName{}, m)
//}

func (c *Public) Process(data []byte, e *events.Basic) bool {
	if e.Code == 0 && e.Result != nil && len(data) > 0 {

		ch := e.Result.Channel
		switch ch {
		case "ticker":
			e := public.Tickers{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				return false
			}
			go func() {
				if c.tCh != nil {
					c.tCh <- &e
				}
			}()
			return true
		//case "book":
		//	e := public.OrderBook{}
		//	err := json.Unmarshal(data, &e)
		//	if err != nil {
		//		fmt.Println(err.Error())
		//		return false
		//	}
		//	if c.obCh != nil {
		//		c.obCh <- &e
		//	}
		//	return true
		default:
		}
	}
	return false
}
