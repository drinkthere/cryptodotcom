package rest

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/drinkthere/cryptodotcom"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sort"
	"strconv"
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
func (c *ClientRest) Do(httpMethod, method string, private bool, params map[string]string) (*http.Response, error) {
	u := fmt.Sprintf("%s%s", c.baseURL, method)
	var (
		r   *http.Request
		err error
	)
	reqID := GenerateRequestID()
	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)
	if httpMethod == http.MethodGet {
		r, err = http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			return nil, err
		}

		q := r.URL.Query()
		for k, v := range params {
			q.Add(k, strings.ReplaceAll(v, "\"", ""))
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
			sign := GenerateSignature(method, reqID, c.apiKey, nonce, c.secretKey, params)
			bodyArr["api_key"] = c.apiKey
			bodyArr["sig"] = sign
		}
		var j []byte
		j, err = json.Marshal(bodyArr)
		if err != nil {
			return nil, err
		}
		r, err = http.NewRequest(httpMethod, u, bytes.NewBuffer(j))
		if err != nil {
			return nil, err
		}
		r.Header.Add("Content-Type", "application/json")
	}
	if err != nil {
		return nil, err
	}

	return c.Client.Do(r)
}

//
//// DoBatch the private post request to the server with parameters of type slice
//func (c *ClientRest) DoBatch(path string, params interface{}) (*http.Response, error) {
//	httpMethod := "POST"
//	u := fmt.Sprintf("%s%s", c.baseURL, path)
//	var (
//		r    *http.Request
//		err  error
//		j    []byte
//		body string
//	)
//
//	j, err = json.Marshal(params)
//	if err != nil {
//		return nil, err
//	}
//	body = string(j)
//	if body == "{}" || body == "[]" {
//		body = ""
//	}
//	r, err = http.NewRequest(httpMethod, u, bytes.NewBuffer(j))
//	if err != nil {
//		return nil, err
//	}
//	r.Header.Add("Content-Type", "application/json")
//
//	if err != nil {
//		return nil, err
//	}
//
//	timestamp, sign := c.sign(httpMethod, path, body)
//	r.Header.Add("OK-ACCESS-KEY", c.apiKey)
//	r.Header.Add("OK-ACCESS-SIGN", sign)
//	r.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)
//
//	return c.Client.Do(r)
//}

func GenerateSignature(method, id, apiKey, nonce string, apiSecret []byte, params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var paramStrBuilder strings.Builder
	for _, k := range keys {
		paramStrBuilder.WriteString(k)
		paramStrBuilder.WriteString(params[k])
	}
	paramStr := paramStrBuilder.String()

	signStr := method + id + apiKey + paramStr + nonce

	mac := hmac.New(sha256.New, apiSecret)
	mac.Write([]byte(signStr))
	signature := mac.Sum(nil)

	// Step 5: 转成 hex 字符串
	return hex.EncodeToString(signature)
}

func GenerateRequestID() string {
	now := time.Now().UnixMilli()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randPart := r.Int63n(1000) // 0~999
	return strconv.FormatInt(now*1000+randPart, 10)
}
