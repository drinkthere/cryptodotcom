package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GenerateSignature(method, id, apiKey, nonce string, apiSecret []byte, params map[string]interface{}) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var paramStrBuilder strings.Builder
	for _, k := range keys {
		paramStrBuilder.WriteString(k)

		// Handle different parameter types
		switch v := params[k].(type) {
		case string:
			paramStrBuilder.WriteString(v)
		case []string:
			for _, val := range v {
				paramStrBuilder.WriteString(val)
			}
		case []interface{}:
			for _, item := range v {
				paramStrBuilder.WriteString(item.(string))
			}
		default:
			// Fallback for other types
			paramStrBuilder.WriteString(fmt.Sprintf("%v", v))
		}
	}
	paramStr := paramStrBuilder.String()
	signStr := method + id + apiKey + paramStr + nonce

	mac := hmac.New(sha256.New, apiSecret)
	mac.Write([]byte(signStr))
	signature := mac.Sum(nil)

	return hex.EncodeToString(signature)
}

func GenerateRequestID() string {
	now := time.Now().UnixMilli()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randPart := r.Int63n(1000) // 0~999
	return strconv.FormatInt(now*1000+randPart, 10)
}
