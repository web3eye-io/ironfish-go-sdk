package main

import (
	"fmt"
	"time"

	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/api"
	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/types"
	"github.com/web3eye-io/ironfish-go-sdk/pkg/utils"
)

const ConnectTimeout = time.Second * 3

func main() {
	sdk_on := api.NewTlsClient("172.16.3.90:8020", "")
	sdk_off := api.NewTlsClient("127.0.0.1:8020", "")

	err := sdk_on.Connect(ConnectTimeout)
	if err != nil {
		fmt.Println(err)
	}

	err = sdk_off.Connect(ConnectTimeout)
	if err != nil {
		fmt.Println(err)
	}

	off_account := types.Account{
		Version: 1,
		// name: "tulip",
		Name:        "apple",
		SpendingKey: "",
		// SpendingKey:     "",
		ViewKey:         "",
		IncomingViewKey: "",
		OutgoingViewKey: "",
		PublicAddress:   "",
	}

	esFRResp, err := sdk_on.EstimateFeeRates()
	fmt.Println(utils.PrettyStruct(esFRResp), err)

	createTXResp, err := sdk_on.CreateTransaction(&types.CreateTransactionRequest{
		Account: off_account.Name,
		Outputs: []types.Output{
			{
				PublicAddress: off_account.PublicAddress,
				Amount:        "100",
				Memo:          "kuku",
			},
		},
		Fee:     "1",
		FeeRate: esFRResp.Average,
	})
	fmt.Println(utils.PrettyStruct(createTXResp), err)

	importAResp, err := sdk_off.ImportAccount(&types.ImportAccountRequest{Account: off_account})
	fmt.Println(utils.PrettyStruct(importAResp), err)

	postTXResp, err := sdk_off.PostTransaction(&types.PostTransactionRequest{Account: "apple", Transaction: createTXResp.Transaction})
	fmt.Println(utils.PrettyStruct(postTXResp), err)

	addTXResp, err := sdk_on.AddTransaction(&types.AddTransactionRequest{Transaction: postTXResp.Transaction, Broadcast: true})
	fmt.Println(utils.PrettyStruct(addTXResp), err)

	getATResp, err := sdk_on.GetAccountTransaction(&types.GetAccountTransactionRequest{Account: "apple", Hash: "3ae8ff87241c20a4045134fc2ced3bf18d3ae0a7211920098bec3549599b2162"})
	fmt.Println(utils.PrettyStruct(getATResp), err)
}
