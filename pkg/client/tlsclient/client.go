package tlsclient

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	NetworkType  = "tcp"
	endChar      = '\x0c'
	readScanTime = time.Microsecond * 200
)

type reqMessage struct {
	MsgType string     `json:"type"`
	MsgData reqMsgData `json:"data"`
}

type reqMessageNoData struct {
	MsgType string           `json:"type"`
	MsgData reqMsgDataNoData `json:"data"`
}

type reqMsgData struct {
	Mid       uint            `json:"mid"`
	MsgType   string          `json:"type"`
	AuthToken string          `json:"auth"`
	Data      json.RawMessage `json:"data"`
}

type reqMsgDataNoData struct {
	Mid       uint   `json:"mid"`
	MsgType   string `json:"type"`
	AuthToken string `json:"auth"`
}

type respMessage struct {
	MsgType string      `json:"type"`
	MsgData respMsgData `json:"data"`
}

type respMsgData struct {
	Id     uint            `json:"id"`
	Status uint            `json:"status"`
	Data   json.RawMessage `json:"data"`
}

type respWrongMsg struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
}

type TlsClient struct {
	Address     string
	AuthToken   string
	msgCount    uint
	conn        *tls.Conn
	msgChannel  map[uint]chan respMsgData
	connChannel chan bool
}

func NewClient(addr string, authToken string) *TlsClient {
	return &TlsClient{
		Address:     addr,
		AuthToken:   authToken,
		msgCount:    1,
		msgChannel:  make(map[uint]chan respMsgData),
		connChannel: make(chan bool),
	}
}

func (tc *TlsClient) Connect(timeout time.Duration) error {
	if tc.conn != nil {
		return nil
	}
	// TODO:should open securty in config
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial(NetworkType, tc.Address, conf)
	if err != nil {
		return err
	}
	tc.conn = conn
	go tc.recv()
	tc.connChannel = make(chan bool)
	return nil
}

func (tc *TlsClient) Request(path string, data []byte, timeout time.Duration) ([]byte, error) {
	mid, err := tc.sendMsg(path, data)
	if err != nil {
		return nil, err
	}

	defer func() {
		delete(tc.msgChannel, mid)
	}()
	ticker := time.NewTicker(timeout)
	checkTicker := time.NewTicker(readScanTime)
	tc.msgChannel[mid] = make(chan respMsgData)
	for {
		select {
		case <-ticker.C:
			return nil, errors.New("request timeout")
		case <-checkTicker.C:
			if tc.conn == nil {
				return nil, errors.New("not connect to server")
			}
		case resp := <-tc.msgChannel[mid]:
			if resp.Status != 200 {
				wrongMsg := &respWrongMsg{}
				json.Unmarshal(resp.Data, wrongMsg)
				return nil, errors.New(wrongMsg.Message)
			}
			return resp.Data, nil
		}
	}
}

func (tc *TlsClient) sendMsg(path string, data []byte) (uint, error) {
	if tc.conn == nil {
		return 0, errors.New("not connect to server")
	}
	var msg any
	var mid = tc.msgCount
	if len(data) > 2 {
		msg = &reqMessage{
			MsgType: "message",
			MsgData: reqMsgData{
				Mid:       mid,
				MsgType:   path,
				AuthToken: tc.AuthToken,
				Data:      data,
			},
		}
	} else {
		msg = &reqMessageNoData{
			MsgType: "message",
			MsgData: reqMsgDataNoData{
				Mid:       mid,
				MsgType:   path,
				AuthToken: tc.AuthToken,
			},
		}
	}

	reqMsg, err := json.Marshal(msg)
	if err != nil {
		return 0, err
	}

	tc.msgCount++
	_, err = tc.conn.Write(append(reqMsg, endChar))
	if err != nil {
		return mid, err
	}
	return mid, err
}

func (tc *TlsClient) recv() {
	// start recv
	go func() {
		respMsg := &respMessage{}
		ticker := time.NewTicker(readScanTime)
		reader := bufio.NewReader(tc.conn)
		for {
			<-ticker.C
			recvData, err := reader.ReadBytes(endChar)
			if err != nil {
				fmt.Println(err)
				tc.Close()
			}
			recvData = recvData[:len(recvData)-1]
			err = json.Unmarshal(recvData, respMsg)
			if err != nil {
				fmt.Println(err)
			}
			if ok := tc.msgChannel[respMsg.MsgData.Id]; ok != nil {
				time.Sleep(time.Second * 2)
				tc.msgChannel[respMsg.MsgData.Id] <- respMsg.MsgData
			}
		}
	}()

	// wait for close
	<-tc.connChannel
	tc.conn.Close()
	tc.conn = nil
	tc.msgCount = 1
}

func (tc *TlsClient) Close() error {
	tc.connChannel <- false
	for k := range tc.msgChannel {
		delete(tc.msgChannel, k)
	}
	return nil
}
