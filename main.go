package main

import (
	"fmt"
	"time"

	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/api"
	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/types"
)

const ConnectTimeout = time.Second * 3

func main() {
	sdk_on := api.NewClient("172.16.3.90:8020", "f067c8c3dce2f165d595f1c82e3d3b61dd502130157f5b503b49fcf4784afc7f", true)

	err := sdk_on.Connect(ConnectTimeout)
	if err != nil {
		fmt.Println(err)
	}

	// efrResp, _ := sdk_on.EstimateFeeRates()
	// feeRate, _ := strconv.ParseUint(efrResp.Average, 10, 64)

	// resp, _ := sdk_on.CreateTransaction(&types.CreateTransactionRequest{
	// 	Account: "default",
	// 	Outputs: []types.Output{{PublicAddress: "c180d3eaa2e5285a109cf1c830f2d68d59a5d8ab69c36f0bc48db42a6e7a3813", Amount: "100000000000", Memo: ""}},
	// 	FeeRate: "10",
	// })
	// fmt.Println(hex.EncodeToString(binary.LittleEndian.AppendUint64([]byte{}, 1010598585)))

	// hexStr, _ := hex.DecodeString(resp.Transaction)
	// fee := binary.LittleEndian.Uint64(hexStr)

	// fmt.Println(fee, "sss", feeRate)

	sdk_on.GetAccountTransaction(&types.GetAccountTransactionRequest{
		Hash:    "e33b678e41f6f3308b82504d699ce3ac4ee44d5aa844da7c1946686aebf6df80",
		Account: "206c09e0af23a2824383fbe64299918afc5f1bebe389de3e251cd59fa00e2453",
	})
}
