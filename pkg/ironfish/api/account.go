package api

import (
	"github.com/web3eye-io/ironfish-go-sdk/pkg/ironfish/types"
)

func (c *Client) GetBalance(req *types.GetBalanceRequest) (*types.GetBalanceResponse, error) {
	resp := &types.GetBalanceResponse{}
	err := request(c, types.GetBalancePath, req, resp)
	return resp, err
}

func (c *Client) ImportAccount(req *types.ImportAccountRequest) (*types.ImportAccountResponse, error) {
	resp := &types.ImportAccountResponse{}
	err := request(c, types.ImportAccountPath, req, resp)
	return resp, err
}

func (c *Client) ExportAccount(req *types.ExportAccountRequest) (*types.ExportAccountResponse, error) {
	resp := &types.ExportAccountResponse{}
	err := request(c, types.ExportAccountPath, req, resp)
	return resp, err
}

func (c *Client) CreateAccount(req *types.CreateAccountRequest) (*types.CreateAccountResponse, error) {
	resp := &types.CreateAccountResponse{}
	err := request(c, types.CreateAccountPath, req, resp)
	return resp, err
}

func (c *Client) IsValidPublicAddress(req *types.IsValidPublicAddressRequest) (*types.IsValidPublicAddressResponse, error) {
	resp := &types.IsValidPublicAddressResponse{}
	err := request(c, types.IsValidPublicAddressPath, req, resp)
	return resp, err
}
