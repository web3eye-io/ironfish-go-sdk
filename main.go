package main

import (
	"fmt"
	"time"

	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/api"
	"github.com/web3eye-io/ironfish-go-sdk/pkg/utils"
)

const ConnectTimeout = time.Second * 3

func main() {
	sdk_on := api.NewClient("172.16.3.90:8020", "421fe347266eef69d83d3ec5d78e9872c961119bc95a2efe4a017c542ebb1071", true)

	err := sdk_on.Connect(ConnectTimeout)
	if err != nil {
		fmt.Println(err)
	}

	getNSReq, err := sdk_on.GetNodeStatus()
	fmt.Println(utils.PrettyStruct(getNSReq), err)
}
