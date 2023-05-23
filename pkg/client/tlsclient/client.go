package tlsclient

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/web3eye-io/ironfish-go-sdk/pkg/client"
)

type TlsClient struct {
	Address     string
	AuthToken   string
	tlsOn       bool
	msgCount    uint
	conn        net.Conn
	lk          sync.Mutex
	msgMap      sync.Map
	connChannel chan bool
}

func NewClient(addr string, authToken string, tlsOn bool) *TlsClient {
	return &TlsClient{
		Address:   addr,
		AuthToken: authToken,
		tlsOn:     tlsOn,
		lk:        sync.Mutex{},
		msgCount:  1,
		msgMap:    sync.Map{},
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

	tc.connChannel = make(chan bool)
	go tc.recv()

	return nil
}

func (tc *TlsClient) Request(path string, data []byte) ([]byte, error) {
	traceID := uuid.NewString()
	mid, err := tc.sendMsg(path, data, traceID)
	if err != nil {
		return nil, err
	}

	msgChan := make(chan *client.RespMsgData)
	tc.msgMap.Store(mid, msgChan)
	defer func() {
		tc.msgMap.Delete(mid)
	}()

	resp := <-msgChan
	if resp == nil {
		log.Printf("recv msg, traceID: %s, connection is closed, time: %s", traceID, time.Now().String())
		return nil, errors.New("connection is closed by ironfish node")
	}

	log.Printf("recv msg, traceID: %s, recv msg: {mid: %d, status: %d, data:%s}, time: %s", traceID, resp.Id, resp.Status, string(resp.Data), time.Now().String())
	if resp.Status != 200 {
		wrongMsg := &client.RespWrongMsg{}
		json.Unmarshal(resp.Data, wrongMsg)
		return nil, errors.New(wrongMsg.Message)
	}
	return resp.Data, nil
}

func (tc *TlsClient) sendMsg(path string, data []byte, traceID string) (uint, error) {
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

	log.Printf("send start, traceID: %s, time: %s", traceID, time.Now().String())
	tc.lk.Lock()
	tc.msgCount++
	_, err = tc.conn.Write(append(reqMsg, client.EndChar))
	tc.lk.Unlock()
	log.Printf("send end, traceID: %s, send msg: %s, req err: %v, time: %s", traceID, string(reqMsg), err, time.Now().String())

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
			log.Printf("recv msg: %v, err: %v", recvData, err)
			if err != nil {
				tc.Close()
				break
			}
			if len(recvData) < 2 {
				continue
			}
			recvData = recvData[:len(recvData)-1]
			err = json.Unmarshal(recvData, respMsg)
			if err != nil {
				fmt.Println(err)
			}
			if ch, ok := tc.msgMap.Load(respMsg.MsgData.Id); ok {
				ch.(chan *client.RespMsgData) <- &respMsg.MsgData
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
	tc.msgMap.Range(func(key, value interface{}) bool {
		if ch, ok := tc.msgMap.Load(key); ok {
			close(ch.(chan *client.RespMsgData))
		}

		tc.msgMap.Delete(key)
		return true
	})
	tc.lk.Lock()
	if tc.connChannel != nil {
		close(tc.connChannel)
		tc.connChannel = nil
	}
	tc.lk.Unlock()
	return nil
}
