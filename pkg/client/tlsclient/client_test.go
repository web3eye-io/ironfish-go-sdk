package tlsclient

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type getAccountReq struct {
	Default     bool `json:"default"`
	DisplayName bool `json:"displayName"`
}

func TestClient(t *testing.T) {
	reqMsg, err := json.Marshal(getAccountReq{Default: false})
	if err != nil {
		fmt.Println(err)
	}
	//  please give right addr and authToken
	var addr = "172.16.3.90:8020"
	var authToken = "421fe347266eef69d83d3ec5d78e9872c961119bc95a2efe4a017c542ebb1071"
	cli := NewClient(addr, authToken)
	err = cli.Connect(time.Second)
	assert.Nil(t, err)
	assert.NotNil(t, cli)
	assert.NotNil(t, cli.conn)

	resp, err := cli.Request("wallet/getAccounts", reqMsg, time.Second*3)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	resp, err = cli.Request("wallet/getAccounts", reqMsg, time.Second*3)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	cli.Close()
	resp, err = cli.Request("wallet/getAccounts", reqMsg, time.Second*3)
	assert.NotNil(t, err)
	assert.Nil(t, resp)

	resp, err = cli.Request("wallet/getAccounts", reqMsg, time.Second*3)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}
