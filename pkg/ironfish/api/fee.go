package api

import (
	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/types"
)

func (c *Client) EstimateFeeRates() (*types.EstimateFeeRatesResponse, error) {
	resp := &types.EstimateFeeRatesResponse{}
	err := request(c, types.EstimateFeeRatesPath, struct{}{}, resp)
	return resp, err
}
