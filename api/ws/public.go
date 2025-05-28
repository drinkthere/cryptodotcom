package ws

import (
	"encoding/json"
	"fmt"
	"github.com/drinkthere/cryptodotcom"
	"github.com/drinkthere/cryptodotcom/events"
	"github.com/drinkthere/cryptodotcom/events/public"
	requests "github.com/drinkthere/cryptodotcom/requests/ws/public"
	"strings"
)

type Public struct {
	*ClientWs
	tCh  chan *public.Tickers
	obCh chan *public.OrderBooks
}

// NewPublic returns a pointer to a fresh Public
func NewPublic(c *ClientWs) *Public {
	return &Public{ClientWs: c}
}

func (c *Public) Tickers(req requests.Tickers, ch chan *public.Tickers) error {
	c.tCh = ch
	channelNames := fillTickerChannel(req.InstrumentNames)
	return c.Subscribe(false, channelNames)
}

func (c *Public) OrderBooks(req requests.OrderBooks, ch chan *public.OrderBooks) error {
	c.obCh = ch
	channelNames := fillOrderBookChannel(req.Instruments)
	return c.Subscribe(false, channelNames)
}

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
		case "book":
			e := public.OrderBooks{}
			err := json.Unmarshal(data, &e)
			if err != nil {
				fmt.Println(err.Error())
				return false
			}
			if c.obCh != nil {
				c.obCh <- &e
			}
			return true
		default:
		}
	}
	return false
}

func fillTickerChannel(instrumentNames []string) []string {
	prefix := cryptodotcom.ChannelPrefixTicker
	result := make([]string, len(instrumentNames))
	for i := range instrumentNames {
		result[i] = strings.Join([]string{string(prefix), instrumentNames[i]}, ".")
	}
	return result
}

func fillOrderBookChannel(contents []*requests.Instrument) []string {
	prefix := cryptodotcom.ChannelPrefixBook
	result := make([]string, len(contents))
	for i := range contents {
		result[i] = strings.Join([]string{string(prefix), contents[i].InstrumentName, contents[i].Depth}, ".")
	}
	return result
}
