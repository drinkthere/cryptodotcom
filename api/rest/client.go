package rest

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/drinkthere/cryptodotcom"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// ClientRest is the rest api client
type ClientRest struct {
	//Account     *Account
	//Trade       *Trade
	//Market      *Market
	PublicData *PublicData
	//TradeData   *TradeData
	apiKey      string
	secretKey   []byte
	destination cryptodotcom.Destination
	baseURL     cryptodotcom.BaseURL
	Client      *http.Client
}

// NewClient returns a pointer to a fresh ClientRest
func NewClient(apiKey, secretKey string, baseURL cryptodotcom.BaseURL, destination cryptodotcom.Destination, ip string) *ClientRest {
	httpClient := http.DefaultClient
	if ip != "" {
		parsedIP := net.ParseIP(ip)
		if parsedIP == nil {
			log.Fatalf("NewClient ip=%s is invalid", ip)
		}

		dialer := &net.Dialer{
			LocalAddr: &net.TCPAddr{
				IP:   parsedIP, // 设置本地出口 IP 地址
				Port: 0,        // 0 表示随机端口
			},
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}

		transport := &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.DialContext(ctx, network, addr)
			},
		}

		httpClient = &http.Client{
			Transport: transport,
		}
	}

	c := &ClientRest{
		apiKey:      apiKey,
		secretKey:   []byte(secretKey),
		baseURL:     baseURL,
		destination: destination,
		Client:      httpClient,
	}
	//c.Account = NewAccount(c)
	//c.Trade = NewTrade(c)
	//c.Market = NewMarket(c)
	c.PublicData = NewPublicData(c)
	//c.TradeData = NewTradeData(c)
	return c
}

// Do the http request to the server
func (c *ClientRest) Do(method, path string, private bool, params ...map[string]string) (*http.Response, error) {
	u := fmt.Sprintf("%s%s", c.baseURL, path)
	var (
		r    *http.Request
		err  error
		j    []byte
		body string
	)
	if method == http.MethodGet {
		r, err = http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			return nil, err
		}

		if len(params) > 0 {
			q := r.URL.Query()
			for k, v := range params[0] {
				q.Add(k, strings.ReplaceAll(v, "\"", ""))
			}
			r.URL.RawQuery = q.Encode()
			if len(params[0]) > 0 {
				path += "?" + r.URL.RawQuery
			}
		}
	} else {
		j, err = json.Marshal(params[0])
		if err != nil {
			return nil, err
		}
		body = string(j)
		if body == "{}" {
			body = ""
		}
		r, err = http.NewRequest(method, u, bytes.NewBuffer(j))
		if err != nil {
			return nil, err
		}
		r.Header.Add("Content-Type", "application/json")
	}
	if err != nil {
		return nil, err
	}
	if private {
		timestamp, sign := c.sign(method, path, body)
		r.Header.Add("OK-ACCESS-KEY", c.apiKey)
		r.Header.Add("OK-ACCESS-SIGN", sign)
		r.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)
	}
	return c.Client.Do(r)
}

// DoBatch the private post request to the server with parameters of type slice
func (c *ClientRest) DoBatch(path string, params interface{}) (*http.Response, error) {
	method := "POST"
	u := fmt.Sprintf("%s%s", c.baseURL, path)
	var (
		r    *http.Request
		err  error
		j    []byte
		body string
	)

	j, err = json.Marshal(params)
	if err != nil {
		return nil, err
	}
	body = string(j)
	if body == "{}" || body == "[]" {
		body = ""
	}
	r, err = http.NewRequest(method, u, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	timestamp, sign := c.sign(method, path, body)
	r.Header.Add("OK-ACCESS-KEY", c.apiKey)
	r.Header.Add("OK-ACCESS-SIGN", sign)
	r.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)

	return c.Client.Do(r)
}

func (c *ClientRest) sign(method, path, body string) (string, string) {
	format := "2006-01-02T15:04:05.999Z07:00"
	t := time.Now().UTC().Format(format)
	ts := fmt.Sprint(t)
	s := ts + method + path + body
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}
