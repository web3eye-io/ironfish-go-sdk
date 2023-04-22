package api

import (
	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/types"
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
	err := request(c, types.AddTransactionPath, req, resp)

	return resp, err
}

func (c *Client) SpendTransaction(req *types.SendTransactionRequest) (*types.SendTransactionResponse, error) {
	resp := &types.SendTransactionResponse{}
	err := request(c, types.SendTransactionPath, req, resp)
	return resp, err
}
