package api

import (
	"encoding/json"
	"time"

	"github.com/web3eye-io/ironfish-go-sdk/pkg/client"
	"github.com/web3eye-io/ironfish-go-sdk/pkg/client/tlsclient"
)

type Client struct {
	client.IronfishClient
}

func NewClient(addr string, authToken string, tlsOn bool) *Client {
	tlsCli := tlsclient.NewClient(addr, authToken, tlsOn)
	return &Client{tlsCli}
}

const RequestTimeout = time.Second * 12

func request[REQ any, RESP any](c client.IronfishClient, path string, req REQ, resp RESP) error {
	reqData, err := json.Marshal(req)
	if err != nil {
		return err
	}
	respData, err := c.Request(path, reqData, RequestTimeout)
	if err != nil {
		return err
	}
	err = json.Unmarshal(respData, resp)
	if err != nil {
		return err
	}
	return err
}
