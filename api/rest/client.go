package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/drinkthere/cryptodotcom"
	"github.com/drinkthere/cryptodotcom/utils"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ClientRest is the rest api client
type ClientRest struct {
	Account     *Account
	Trade       *Trade
	PublicData  *PublicData
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
	c.Account = NewAccount(c)
	c.Trade = NewTrade(c)
	c.PublicData = NewPublicData(c)
	return c
}

// Do the http request to the server
func (c *ClientRest) Do(httpMethod, method string, private bool, params map[string]interface{}) (*http.Response, error) {
	u := fmt.Sprintf("%s%s", c.baseURL, method)
	var (
		r   *http.Request
		err error
	)
	reqID := utils.GenerateRequestID()
	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)

	if httpMethod == http.MethodGet {
		r, err = http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			return nil, err
		}

		q := r.URL.Query()
		for k, v := range params {
			// Handle different types of values
			switch val := v.(type) {
			case string:
				q.Add(k, strings.ReplaceAll(val, "\"", ""))
			case []string:
				for _, item := range val {
					q.Add(k, strings.ReplaceAll(item, "\"", ""))
				}
			default:
				// Convert other types to string
				q.Add(k, fmt.Sprintf("%v", val))
			}
		}
		r.URL.RawQuery = q.Encode()
		method += "?id=" + reqID + "&nonce=" + nonce
		if len(params) > 0 {
			method += "&" + r.URL.RawQuery
		}
	} else {
		bodyArr := map[string]interface{}{}
		bodyArr["id"] = reqID
		bodyArr["method"] = method
		bodyArr["params"] = params
		bodyArr["nonce"] = nonce

		if private {
			sign := utils.GenerateSignature(method, reqID, c.apiKey, nonce, c.secretKey, params)
			bodyArr["api_key"] = c.apiKey
			bodyArr["sig"] = sign
		}

		j, err := json.Marshal(bodyArr)
		if err != nil {
			return nil, err
		}

		r, err = http.NewRequest(httpMethod, u, bytes.NewBuffer(j))
		if err != nil {
			return nil, err
		}
		r.Header.Add("Content-Type", "application/json")
	}

	return c.Client.Do(r)
}
