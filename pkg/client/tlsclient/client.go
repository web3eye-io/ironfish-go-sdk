package tlsclient

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/web3eye-io/ironfish-go-sdk/pkg/client"
)

type TlsClient struct {
	Address     string
	AuthToken   string
	tlsOn       bool
	msgCount    uint
	conn        net.Conn
	msgChannel  map[uint]chan client.RespMsgData
	connChannel chan bool
}

func NewClient(addr string, authToken string, tlsOn bool) *TlsClient {
	return &TlsClient{
		Address:     addr,
		AuthToken:   authToken,
		tlsOn:       tlsOn,
		msgCount:    1,
		msgChannel:  make(map[uint]chan client.RespMsgData),
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
	if tc.tlsOn {
		conn, err := tls.Dial(client.NetworkType, tc.Address, conf)
		if err != nil {
			return err
		}
		tc.conn = conn
	} else {
		conn, err := net.Dial(client.NetworkType, tc.Address)
		if err != nil {
			return err
		}
		tc.conn = conn
	}

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
	checkTicker := time.NewTicker(client.ReadScanTime)
	tc.msgChannel[mid] = make(chan client.RespMsgData)
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
				wrongMsg := &client.RespWrongMsg{}
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
		msg = &client.ReqMessage{
			MsgType: "message",
			MsgData: client.ReqMsgData{
				Mid:       mid,
				MsgType:   path,
				AuthToken: tc.AuthToken,
				Data:      data,
			},
		}
	} else {
		msg = &client.ReqMessageNoData{
			MsgType: "message",
			MsgData: client.ReqMsgDataNoData{
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
	_, err = tc.conn.Write(append(reqMsg, client.EndChar))
	if err != nil {
		return mid, err
	}

	return mid, err
}

func (tc *TlsClient) recv() {
	// start recv
	go func() {
		respMsg := &client.RespMessage{}
		ticker := time.NewTicker(client.ReadScanTime)
		reader := bufio.NewReader(tc.conn)
		for {
			<-ticker.C
			recvData, err := reader.ReadBytes(client.EndChar)
			if err != nil {
				fmt.Println(err)
				tc.Close()
			}
			if len(recvData) < 2 {
				continue
			}
			recvData = recvData[:len(recvData)-1]
			err = json.Unmarshal(recvData, respMsg)
			if err != nil {
				fmt.Println(err)
			}
			if ok := tc.msgChannel[respMsg.MsgData.Id]; ok != nil {
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
