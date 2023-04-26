package types

type Account struct {
	Version         int    `json:"version"`
	Name            string `json:"name"`
	SpendingKey     string `json:"spendingKey"`
	ViewKey         string `json:"viewKey"`
	IncomingViewKey string `json:"incomingViewKey"`
	OutgoingViewKey string `json:"outgoingViewKey"`
	PublicAddress   string `json:"publicAddress"`
	CreatedAt       string `json:"createdAt"`
}

const GetBalancePath = "wallet/getBalance"

type GetBalanceRequest struct {
	Account       string `json:"account"`
	AssetId       string `json:"assetId"`
	Confirmations uint   `json:"confirmations"`
}

type GetBalanceResponse struct {
	Account          string `json:"account"`
	AssetId          string `json:"assetId"`
	Confirmed        string `json:"confirmed"`
	Unconfirmed      string `json:"unconfirmed"`
	UnconfirmedCount uint   `json:"unconfirmedCount"`
	Pending          string `json:"pending"`
	PendingCount     uint   `json:"pendingCount"`
	Available        string `json:"available"`
	Confirmations    uint   `json:"confirmations"`
	BlockHash        string `json:"blockHash"`
	Sequence         uint   `json:"sequence"`
}

const ExportAccountPath = "wallet/exportAccount"

type ExportAccountRequest struct {
	Account  string `json:"account"`
	ViewOnly bool   `json:"viewOnly"`
}
type ExportAccountResponse struct {
	Account Account `json:"account"`
}

const CreateAccountPath = "wallet/create"

type CreateAccountRequest struct {
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

type CreateAccountResponse struct {
	Name             string `json:"name"`
	PublicAddress    string `json:"publicAddress"`
	IsDefaultAccount bool   `json:"isDefaultAccount"`
}

const ImportAccountPath = "wallet/importAccount"

type ImportAccountRequest struct {
	Account Account `json:"account"`
	Rescan  bool    `json:"rescan"`
}
type ImportAccountResponse struct {
	Name             string `json:"name"`
	IsDefaultAccount bool   `json:"isDefaultAccount"`
}

const IsValidPublicAddressPath = "chain/isValidPublicAddress"

type IsValidPublicAddressRequest struct {
	Address string `json:"address"`
}
type IsValidPublicAddressResponse struct {
	Valid bool `json:"valid"`
}
