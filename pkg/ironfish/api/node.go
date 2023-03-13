package api

import (
	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/types"
)

func (c *Client) GetNodeStatus() (*types.GetNodeStatusResponse, error) {
	resp := &types.GetNodeStatusResponse{}
	err := request(c, types.GetNodeStatusPath, &types.GetNodeStatusRequest{Stream: false}, resp)
	return resp, err
}
