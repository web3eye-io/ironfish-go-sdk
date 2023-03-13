package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/types"
	"github.com/web3eye-io/ironfish-go-sdk/pkg/utils"
)

const ConnectTimeout = time.Second * 3

func TestBaseApi(t *testing.T) {
	// any key need be replace
	sdk_on := NewTlsClient("172.16.3.90:8020", "421fxxxxxxxxxxxxxxxxxxxxxxxxxxxx542ebb1071")
	sdk_off := NewTlsClient("127.0.0.1:8020", "8ae867d9xxxxxxxxxxxxxxxxxxxxxxxxxx0a8fa1f2394c79")

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
		Name: "apple",
		// spendingKey: "5cb10c40xxxxxxxxxxxxxxxxxxxxxx578bac88a9f1b082",
		SpendingKey:     "",
		ViewKey:         "b45c9bdd35859ssssssssssssssssssssssssssssc0c9c80e6be969608e3317703adb3bd05fb3bd551209d1f329645e468e0a64bb5ea",
		IncomingViewKey: "2c75a15a40de0c2dxxxxxxxxxxxxxxxx0c9517efa41348ed255e2c207",
		OutgoingViewKey: "6c75e38945673exxxxxxxxxxxxxxxxxx537e02d210c5cac55",
		PublicAddress:   "70ff227xxxxxxxxxxxxxxxxxxxxxxxxx61b5580c663274387d1",
	}

	createAResp, err := sdk_on.CreateAccount(&types.CreateAccountRequest{Name: "banana"})
	fmt.Println(utils.PrettyStruct(createAResp), err)

	importAResp, err := sdk_on.ImportAccount(&types.ImportAccountRequest{Account: types.Account{
		Version:     1,
		Name:        "beaf",
		SpendingKey: "",
		// SpendingKey:     "d09b1148e3b942xxxxxxxxxxxxx5b9539c827c59ed561f9",
		ViewKey:         "2d20afe4f094c21812axxxxxxxxxxxxxe55f3e93d7d1faa3f7021e23fe5b6e74b27ae3fdc0a0694ad1fcf478b46f7d3de63085173b",
		IncomingViewKey: "221efa45052e84xxxxxxxxxxx62274b444fc437b0f081b0db161806",
		OutgoingViewKey: "40aa9c0f5b79xxxxxxxxxxxxxcebb66e9af06092ce44a26a2ea7618dfc789",
		PublicAddress:   "f3ccf3b024be2xxxxxxxxxxxxxxxx11a1ca8e3ca8c64cba76ec341580fc9b57",
	}})
	fmt.Println(utils.PrettyStruct(importAResp), err)

	getBResp, err := sdk_on.GetBalance(&types.GetBalanceRequest{Account: off_account.Name})
	fmt.Println(utils.PrettyStruct(getBResp), err)

	sendTXResp, err := sdk_on.SpendTransaction(&types.SendTransactionRequest{
		Account: "baby",
		Outputs: []types.Output{
			{
				PublicAddress: off_account.PublicAddress,
				Amount:        "100",
			},
		},
		Fee: "1",
	})
	fmt.Println(utils.PrettyStruct(sendTXResp), err)

	getATResp, err := sdk_on.GetAccountTransaction(&types.GetAccountTransactionRequest{Account: "baby", Hash: "cab361a9b68e8ef1b05485575afd7d508eda0f5d0621289007f2c17be302b0d6"})
	fmt.Println(utils.PrettyStruct(getATResp), err)
}

func TestTransaactionApi(t *testing.T) {
	sdk_on := NewTlsClient("172.16.3.90:8020", "421fe347266exxxxxxxxxxxxxxxxxxxxx7c542ebb1071")
	sdk_off := NewTlsClient("127.0.0.1:8020", "8ae867d926eaxxxxxxxxxxxxxxxxxxxxxxxae0a8fa1f2394c79")

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
		SpendingKey: "5cb10c405be2b6xxxxxxxxxxxxxxxxxxxxxx5746a6e578bac88a9f1b082",
		// SpendingKey:     "",
		ViewKey:         "b45c9bdd358xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx6be969608e3317703adb3bd05fb3bd551209d1f329645e468e0a64bb5ea",
		IncomingViewKey: "2c75xxxxxxxxxxxxxxxxxxxxxxxxxx48ed255e2c207",
		OutgoingViewKey: "6c75exxxxxxxxxxxxxxxxxxxxxxxxxxxxx7e02d210c5cac55",
		PublicAddress:   "70ff2xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx74387d1",
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

func TestNodeStatus(t *testing.T) {
	sdk_on := NewTlsClient("172.16.3.90:8020", "421fe34xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx71")
	sdk_off := NewTlsClient("127.0.0.1:8020", "8ae8xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx2394c79")

	err := sdk_on.Connect(ConnectTimeout)
	if err != nil {
		fmt.Println(err)
	}

	err = sdk_off.Connect(ConnectTimeout)
	if err != nil {
		fmt.Println(err)
	}

	getNSReq, err := sdk_on.GetNodeStatus()
	fmt.Println(utils.PrettyStruct(getNSReq), err)
	fmt.Println(utils.PrettyStruct(getNSReq), err)
}
