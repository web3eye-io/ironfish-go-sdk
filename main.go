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
	sdk_on := api.NewClient("172.16.3.90:8020", "421fe347266eef69d83d3ec5d78e9872c961119bc95a2efe4a017c542ebb1071", false)

	err := sdk_on.Connect(ConnectTimeout)
	if err != nil {
		fmt.Println(err)
	}

	off_account := types.Account{
		Version: 1,
		// name: "tulip",
		Name:        "apple",
		SpendingKey: "5cb10c405be2b62cf036a0deec9ede36522347d5d5746a6e578bac88a9f1b082",
		// SpendingKey:     "",
		ViewKey:         "b45c9bdd3585934632ce148e79f1ef0f9b93a720ec46381bb99422010e279c0c9c80e6be969608e3317703adb3bd05fb3bd551209d1f329645e468e0a64bb5ea",
		IncomingViewKey: "2c75a15a40de0c2d823d46fc0db95b99a5f59b90c9517efa41348ed255e2c207",
		OutgoingViewKey: "6c75e38945673efd214465818be4a45d5886175c2397e56537e02d210c5cac55",
		PublicAddress:   "70ff227690e8799d28456aa9d5ea741ba9e5066a6bb8c61b5580c663274387d1",
	}

	createTXResp, err := sdk_on.CreateTransaction(&types.CreateTransactionRequest{
		Account: off_account.Name,
		Outputs: []types.Output{
			{
				PublicAddress: off_account.PublicAddress,
				Amount:        "1000",
				Memo:          "",
			},
		},
		Fee: "1",
	})
	fmt.Println(utils.PrettyStruct(createTXResp), err)

}
