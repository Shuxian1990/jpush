package common

import (
	"github.com/printfcoder/goutils/stringutils"
	"net/http"
)

// Client 基础客户端
type Client struct {
	InitParams
}

func (c *Client) AddAuthToHeader(header *http.Header) {
	header.Add("Content-Type", "application/json")
	sign := stringutils.ToBase64(c.AppKey + ":" + c.MasterSecret)
	header.Add("Authorization", "Basic "+sign)
}
