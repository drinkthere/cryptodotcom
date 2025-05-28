package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/drinkthere/cryptodotcom"
	"github.com/drinkthere/cryptodotcom/events"
	"net"
	"net/http"
	"strconv"
	"strings"

	//"github.com/drinkthere/cryptodotcom/events"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type ClientWs struct {
	url           map[bool]cryptodotcom.BaseURL
	apiKey        string
	secretKey     []byte
	conn          map[bool]*websocket.Conn
	mu            map[bool]*sync.RWMutex
	closed        map[bool]bool
	ctx           context.Context
	Cancel        context.CancelFunc
	DoneChan      chan interface{}
	ErrChan       chan *events.Basic
	LoginChan     chan *events.Login
	SuccessChan   chan *events.Basic
	sendChan      map[bool]chan []byte
	lastTransmit  sync.Map
	AuthRequested *time.Time
	Authorized    bool
	//Private       *Private
	Public *Public
	//Trade         *Trade
	LocalIP string
}

const (
	redialTick = 2 * time.Second
	writeWait  = 3 * time.Second
	pingWait   = 35 * time.Second
	PingPeriod = 15 * time.Second
)

func NewClient(ctx context.Context, apiKey, secretKey string, url map[bool]cryptodotcom.BaseURL, ip string) *ClientWs {
	ctx, cancel := context.WithCancel(ctx)
	c := &ClientWs{
		url:       url,
		apiKey:    apiKey,
		secretKey: []byte(secretKey),
		conn:      make(map[bool]*websocket.Conn),
		closed:    make(map[bool]bool),
		mu:        map[bool]*sync.RWMutex{true: {}, false: {}},
		ctx:       ctx,
		Cancel:    cancel,
		sendChan:  map[bool]chan []byte{true: make(chan []byte, 3), false: make(chan []byte, 3)},
		DoneChan:  make(chan interface{}, 32),
		LocalIP:   ip,
	}

	c.Public = NewPublic(c)
	//c.Private = NewPrivate(c)
	//c.Trade = NewTrade(c)
	now := time.Now()
	c.lastTransmit.Store(true, &now)
	c.lastTransmit.Store(false, &now)
	return c
}

func (c *ClientWs) Connect(p bool) error {
	if c.CheckConnect(p) {
		return nil
	}

	err := c.dial(p)
	if err == nil {
		return nil
	}

	ticker := time.NewTicker(redialTick)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err = c.dial(p)
			if err == nil {
				return nil
			}
		case <-c.ctx.Done():
			return c.handleCancel("connect")
		}
	}
}

// CheckConnect into the server
func (c *ClientWs) CheckConnect(p bool) bool {
	c.mu[p].RLock()
	defer c.mu[p].RUnlock()
	if c.conn[p] != nil && !c.closed[p] {
		return true
	}
	return false
}

//
//// Login
////
//// https://www.okx.com/docs-v5/en/#websocket-api-login
//func (c *ClientWs) Login() error {
//	if c.Authorized {
//		return nil
//	}
//
//	if c.AuthRequested != nil && time.Since(*c.AuthRequested).Seconds() < 30 {
//		return nil
//	}
//
//	now := time.Now()
//	c.AuthRequested = &now
//	method := http.MethodGet
//	path := "/users/self/verify"
//	ts, sign := c.sign(method, path)
//	args := []map[string]string{
//		{
//			"apiKey":     c.apiKey,
//			"passphrase": c.passphrase,
//			"timestamp":  ts,
//			"sign":       sign,
//		},
//	}
//
//	return c.Send(true, okx.LoginOperation, args)
//}

func (c *ClientWs) Subscribe(p bool, ch []string) error {
	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)
	return c.Send(p, nonce, cryptodotcom.SubscribeOperation, ch)

}

func (c *ClientWs) Unsubscribe(p bool, ch []string) error {
	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)
	return c.Send(p, nonce, cryptodotcom.UnsubscribeOperation, ch)
}

func (c *ClientWs) Pong(p bool, nonce string) error {
	return c.Send(p, nonce, cryptodotcom.PongOperation, []string{})
}

// Send message through either connections
func (c *ClientWs) Send(p bool, nonce string, method cryptodotcom.Operation, channelNames []string) error {
	if method != cryptodotcom.LoginOperation {
		err := c.Connect(p)
		if err == nil {
			if p {
				//err = c.WaitForAuthorization()
				//if err != nil {
				//	return err
				//}
			}
		} else {
			return err
		}
	}

	var data map[string]interface{}
	if method == cryptodotcom.PongOperation {
		// 单独处理pong消息
		data = map[string]interface{}{
			"id":     nonce,
			"method": "public/respond-heartbeat",
		}
	} else {
		data = map[string]interface{}{
			"id":     nonce,
			"method": method,
			"params": map[string][]string{
				"channels": channelNames,
			},
			"nonce": nonce,
		}
	}

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	c.mu[p].RLock()
	c.sendChan[p] <- j
	c.mu[p].RUnlock()
	return nil
}

// SetChannels to receive certain events on separate channel
func (c *ClientWs) SetChannels(lCh chan *events.Login, errCh, sCh chan *events.Basic) {
	c.LoginChan = lCh
	c.ErrChan = errCh
	c.SuccessChan = sCh
}

// SetErrChannel set error channel
func (c *ClientWs) SetErrChannel(errCh chan *events.Basic) {
	c.ErrChan = errCh
}

// SetLoginChannel set error channel
func (c *ClientWs) SetLoginChannel(lCh chan *events.Login) {
	c.LoginChan = lCh
}

//
//// WaitForAuthorization waits for the auth responses and try to log in if it was needed
//func (c *ClientWs) WaitForAuthorization() error {
//	if c.Authorized {
//		return nil
//	}
//
//	if err := c.Login(); err != nil {
//		return err
//	}
//
//	ticker := time.NewTicker(time.Millisecond * 300)
//	defer ticker.Stop()
//
//	for range ticker.C {
//		if c.Authorized {
//			return nil
//		}
//	}
//
//	return nil
//}

func (c *ClientWs) dial(p bool) error {
	c.mu[p].Lock()
	var dialer websocket.Dialer
	if c.LocalIP != "" {
		dialer = websocket.Dialer{
			NetDial: func(network, addr string) (net.Conn, error) {
				localAddr, err := net.ResolveTCPAddr("tcp", c.LocalIP+":0") // 替换为您的出口IP地址
				if err != nil {
					return nil, err
				}
				d := net.Dialer{
					LocalAddr: localAddr,
				}
				return d.Dial(network, addr)
			},
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		}
	} else {
		dialer = websocket.Dialer{
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		}
	}
	conn, res, err := dialer.Dial(string(c.url[p]), nil)
	if err != nil {
		var statusCode int
		if res != nil {
			statusCode = res.StatusCode
		}

		c.mu[p].Unlock()

		return fmt.Errorf("error %d: %w", statusCode, err)
	}
	defer res.Body.Close()

	go func() {
		defer func() {
			// Cleaning the connection with ws
			c.Cancel()
			c.mu[p].Lock()
			c.conn[p].Close()
			c.closed[p] = true
			fmt.Printf("receiver connection closed\n")
			c.mu[p].Unlock()
		}()
		err := c.receiver(p)
		if err != nil {
			if !strings.Contains(err.Error(), "operation cancelled: receiver") {
				c.ErrChan <- &events.Basic{
					Method:  "error",
					Message: err.Error(),
				}
			}
			fmt.Printf("receiver error: %v\n", err)
		}
	}()

	go func() {
		defer func() {
			// Cleaning the connection with ws
			c.Cancel()
			c.mu[p].Lock()
			c.conn[p].Close()
			c.closed[p] = true
			fmt.Printf("sender connection closed\n")
			c.mu[p].Unlock()
		}()
		err := c.sender(p)
		if err != nil {
			if !strings.Contains(err.Error(), "operation cancelled: sender") {
				c.ErrChan <- &events.Basic{
					Method:  "error",
					Message: err.Error(),
				}
			}
			fmt.Printf("sender error: %v\n", err)
			c.Authorized = false
		}
	}()

	c.conn[p] = conn
	c.closed[p] = false
	c.mu[p].Unlock()

	return nil
}

func (c *ClientWs) sender(p bool) error {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()

	for {
		c.mu[p].RLock()
		dataChan := c.sendChan[p]
		c.mu[p].RUnlock()

		select {
		case data := <-dataChan:
			c.mu[p].RLock()
			err := c.conn[p].SetWriteDeadline(time.Now().Add(writeWait))
			if err != nil {
				c.mu[p].RUnlock()
				return fmt.Errorf("failed to set write deadline for ws connection, error: %w", err)
			}

			w, err := c.conn[p].NextWriter(websocket.TextMessage)
			if err != nil {
				c.mu[p].RUnlock()
				return fmt.Errorf("failed to get next writer for ws connection, error: %w", err)
			}

			if _, err = w.Write(data); err != nil {
				c.mu[p].RUnlock()
				return fmt.Errorf("failed to write data via ws connection, error: %w", err)
			}

			c.mu[p].RUnlock()

			if err := w.Close(); err != nil {
				return fmt.Errorf("failed to close ws connection, error: %w", err)
			}
		case <-ticker.C:
			lastTransmitInterface, _ := c.lastTransmit.Load(p)
			lastTransmit := lastTransmitInterface.(*time.Time)
			c.mu[p].RLock()
			if c.conn[p] != nil && (lastTransmit == nil || (lastTransmit != nil && time.Since(*lastTransmit) > PingPeriod)) {
				go func() {
					c.mu[p].RLock()
					c.sendChan[p] <- []byte("ping")
					c.mu[p].RUnlock()
				}()
			}

			c.mu[p].RUnlock()
		case <-c.ctx.Done():
			return c.handleCancel("sender")
		}
	}
}
func (c *ClientWs) receiver(p bool) error {
	for {
		select {
		case <-c.ctx.Done():
			return c.handleCancel("receiver")
		default:
			c.mu[p].RLock()
			err := c.conn[p].SetReadDeadline(time.Now().Add(pingWait))
			if err != nil {
				c.mu[p].RUnlock()
				return fmt.Errorf("failed to set read deadline for ws connection, error: %w", err)
			}

			mt, data, err := c.conn[p].ReadMessage()
			if err != nil {
				c.mu[p].RUnlock()
				// 判断是否为读超时错误
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					return fmt.Errorf("server inactive for %v: no ping received", pingWait)
				}
				return fmt.Errorf("failed to read message from ws connection, error: %v\n", err)
			}
			c.mu[p].RUnlock()

			now := time.Now()
			c.lastTransmit.Store(p, &now)

			if mt == websocket.TextMessage {
				e := &events.Basic{}
				if err := json.Unmarshal(data, e); err != nil {
					return fmt.Errorf("failed to unmarshall message from ws, error: %w", err)
				}
				go c.process(p, data, e)
			}
		}
	}
}

func (c *ClientWs) sign(method, path string) (string, string) {
	t := time.Now().UTC().Unix()
	ts := fmt.Sprint(t)
	s := ts + method + path
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)

	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *ClientWs) handleCancel(msg string) error {
	go func() {
		c.DoneChan <- msg
	}()

	return fmt.Errorf("operation cancelled: %s", msg)
}

// TODO: break each case into a separate function
func (c *ClientWs) process(p bool, data []byte, e *events.Basic) bool {
	switch e.Method {
	case "error":
		go func() {
			if c.ErrChan != nil {
				c.ErrChan <- e
			}
		}()
		return true

	case "public/heartbeat":
		c.Pong(p, strconv.FormatInt(int64(e.ID), 10))
		return true

	case "subscribe":
		//if c.Private.Process(data, e) {
		//	return true
		//}

		if c.Public.Process(data, e) {
			return true
		}

		return true

		//case "login":
		//	if time.Since(*c.AuthRequested).Seconds() > 30 {
		//		c.AuthRequested = nil
		//		_ = c.Login()
		//		break
		//	}
		//
		//	c.Authorized = true
		//
		//	e := events.Login{}
		//	_ = json.Unmarshal(data, &e)
		//	go func() {
		//		if c.LoginChan != nil {
		//			c.LoginChan <- &e
		//		}
		//	}()
		//
		//	return true
	}

	if e.Code != 0 {
		ee := *e
		ee.Method = "error"

		return c.process(p, data, &ee)
	}

	if c.SuccessChan != nil {
		c.SuccessChan <- e
	}

	return true
}
