package types

const (
	CONFIRMED   TransactionStatus = "confirmed"
	EXPIRED     TransactionStatus = "expired"
	PENDING     TransactionStatus = "pending"
	UNCONFIRMED TransactionStatus = "unconfirmed"
	UNKNOWN     TransactionStatus = "unknown"
)

type TransactionStatus string

type Output struct {
	PublicAddress string `json:"publicAddress"`
	Amount        string `json:"amount"`
	Memo          string `json:"memo"`
	AssetId       string `json:"assetId"`
}

const CreateTransactionPath = "wallet/createTransaction"

type CreateTransactionRequest struct {
	Account         string   `json:"account"`
	Outputs         []Output `json:"outputs"`
	Fee             string   `json:"fee"`
	FeeRate         string   `json:"feeRate"`
	Expiration      uint     `json:"expiration"`
	ExpirationDelta uint     `json:"expirationDelta"`
	Confirmations   uint     `json:"confirmations"`
}
type CreateTransactionResponse struct {
	Transaction string `json:"transaction"`
}

const PostTransactionPath = "wallet/postTransaction"

type PostTransactionRequest struct {
	Account     string `json:"account"`
	Transaction string `json:"transaction"`
	Broadcast   bool   `json:"broadcast"`
}
type PostTransactionResponse struct {
	Hash        string `json:"hash"`
	Transaction string `json:"transaction"`
}

const AddTransactionPath = "wallet/addTransaction"

type AddTransactionRequest struct {
	Transaction string `json:"transaction"`
	Broadcast   bool   `json:"broadcast"`
}
type AddTransactionResponse struct {
	Hash     string   `json:"hash"`
	Accounts []string `json:"accounts"`
}

const SendTransactionPath = "wallet/sendTransaction"

type SendTransactionRequest struct {
	Account         string   `json:"account"`
	Outputs         []Output `json:"outputs"`
	Fee             string   `json:"fee"`
	Expiration      uint     `json:"expiration"`
	ExpirationDelta uint     `json:"expirationDelta"`
	Confirmations   uint     `json:"confirmations"`
}
type SendTransactionResponse struct {
	Account     string `json:"account"`
	Hash        string `json:"hash"`
	Transaction string `json:"transaction"`
}

const GetAccountTransactionPath = "wallet/getAccountTransaction"

type AccountDecryptedNote struct {
	IsOwner   bool   `json:"isOwner"`
	Value     string `json:"value"`
	AssetId   string `json:"assetId"`
	AssetName string `json:"assetName"`
	Memo      string `json:"memo"`
	Sender    string `json:"sender"`
	Owner     string `json:"owner"`
	Spent     bool   `json:"spent"`
}

type GetAccountTransactionRequest struct {
	Hash          string `json:"hash"`
	Account       string `json:"account"`
	Confirmations uint   `json:"confirmations"`
}

type GetAccountTransactionResponse struct {
	Account     string `json:"account"`
	Transaction struct {
		Hash               string                 `json:"hash"`
		Status             TransactionStatus      `json:"status"`
		Type               string                 `json:"type"`
		Fee                string                 `json:"fee"`
		BlockHash          string                 `json:"blockHash"`
		BlockSequence      uint                   `json:"blockSequence"`
		NotesCount         uint                   `json:"notesCount"`
		SpendsCount        uint                   `json:"spendsCount"`
		MintsCount         uint                   `json:"mintsCount"`
		BurnsCount         uint                   `json:"burnsCount"`
		Timestamp          uint                   `json:"timestamp"`
		Notes              []AccountDecryptedNote `json:"notes"`
		AssetBalanceDeltas []struct {
			AssetId   string `json:"assetId"`
			AssetName string `json:"assetName"`
			Delta     string `json:"delta"`
		}
	} `json:"transaction"`
}
