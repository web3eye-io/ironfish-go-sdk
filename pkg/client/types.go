package client

import (
	"encoding/json"
	"time"
)

const (
	NetworkType  = "tcp"
	EndChar      = '\x0c'
	ReadScanTime = time.Microsecond * 50
)

type ReqMessage struct {
	MsgType string     `json:"type"`
	MsgData ReqMsgData `json:"data"`
}

type ReqMessageNoData struct {
	MsgType string           `json:"type"`
	MsgData ReqMsgDataNoData `json:"data"`
}

type ReqMsgData struct {
	Mid       uint            `json:"mid"`
	MsgType   string          `json:"type"`
	AuthToken string          `json:"auth"`
	Data      json.RawMessage `json:"data"`
}

type ReqMsgDataNoData struct {
	Mid       uint   `json:"mid"`
	MsgType   string `json:"type"`
	AuthToken string `json:"auth"`
}

type RespMessage struct {
	MsgType string      `json:"type"`
	MsgData RespMsgData `json:"data"`
}

type RespMsgData struct {
	Id     uint            `json:"id"`
	Status uint            `json:"status"`
	Data   json.RawMessage `json:"data"`
}

type RespWrongMsg struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
}
type IronfishClient interface {
	Connect(timeout time.Duration) error
	Request(path string, data []byte, timeout time.Duration) ([]byte, error)
	Close() error
}
