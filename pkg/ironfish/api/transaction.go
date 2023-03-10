package api

import (
	"encoding/hex"

	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/types"
	"lukechampine.com/blake3"
)

func (c *Client) GetAccountTransaction(req *types.GetAccountTransactionRequest) (*types.GetAccountTransactionResponse, error) {
	resp := &types.GetAccountTransactionResponse{}
	err := request(c, types.GetAccountTransactionPath, req, resp)
	return resp, err
}

func (c *Client) CreateTransaction(req *types.CreateTransactionRequest) (*types.CreateTransactionResponse, error) {
	resp := &types.CreateTransactionResponse{}
	err := request(c, types.CreateTransactionPath, req, resp)
	return resp, err
}

func (c *Client) PostTransaction(req *types.PostTransactionRequest) (*types.PostTransactionResponse, error) {
	resp := &types.PostTransactionResponse{}
	err := request(c, types.PostTransactionPath, req, resp)
	return resp, err
}
func (c *Client) AddTransaction(req *types.AddTransactionRequest) (*types.AddTransactionResponse, error) {
	resp := &types.AddTransactionResponse{}
	err := request(c, types.CreateTransactionPath, req, resp)

	if err != nil && len(resp.Accounts) > 0 {
		txData, err := hex.DecodeString(req.Transaction)
		if err != nil {
			return resp, err
		}
		_txHash := blake3.Sum256(txData)
		resp.Hash = hex.EncodeToString(_txHash[:])
	}

	return resp, err
}
