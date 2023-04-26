package main

import (
	"fmt"
	"time"

	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/api"
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

}
