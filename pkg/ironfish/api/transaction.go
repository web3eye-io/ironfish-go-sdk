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

	if err == nil && len(resp.Transaction) > 0 {
		txHash, err := getTxHash(resp.Transaction)
		if err != nil {
			return nil, err
		}
		resp.Hash = txHash
	}
	return resp, err
}
func (c *Client) AddTransaction(req *types.AddTransactionRequest) (*types.AddTransactionResponse, error) {
	resp := &types.AddTransactionResponse{}
	err := request(c, types.AddTransactionPath, req, resp)

	if err == nil && len(resp.Accounts) > 0 {
		txHash, err := getTxHash(req.Transaction)
		if err != nil {
			return nil, err
		}
		resp.Hash = txHash
	}

	return resp, err
}

func (c *Client) SpendTransaction(req *types.SendTransactionRequest) (*types.SendTransactionResponse, error) {
	resp := &types.SendTransactionResponse{}
	err := request(c, types.SendTransactionPath, req, resp)
	return resp, err
}

func getTxHash(tx string) (string, error) {
	txData, err := hex.DecodeString(tx)
	if err != nil {
		return "", err
	}
	_txHash := blake3.Sum256(txData)
	return hex.EncodeToString(_txHash[:]), nil
}
