package navexplorer

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ValidateAddress struct {
	Valid           bool   `json:"isValid"`
	Address         string `json:"address"`
	StakingAddress  string `json:"stakingAddress"`
	SpendingAddress string `json:"spendingAddress"`
	ColdStaking     bool   `json:"isColdStaking"`
}

type Address struct {
	Hash               string `json:"hash"`
	Received           int64  `json:"received"`
	ReceivedCount      uint   `json:"receivedCount"`
	Sent               int64  `json:"sent"`
	SentCount          uint   `json:"sentCount"`
	Staked             int64  `json:"staked"`
	StakedCount        uint   `json:"stakedCount"`
	ColdStaked         int64  `json:"coldStaked"`
	ColdStakedCount    uint   `json:"coldStakedCount"`
	ColdStakedSent     int64  `json:"coldSent"`
	ColdStakedReceived int64  `json:"coldStakedReceived"`
	ColdStakedBalance  int64  `json:"coldStakedBalance"`
	Balance            int64  `json:"balance"`
	Position           int64  `json:"position"`
}

type Transaction struct {
	Time                time.Time `json:"time"`
	Address             string    `json:"address"`
	Type                string    `json:"type"`
	Transaction         string    `json:"transaction"`
	Height              int       `json:"height"`
	Balance             float64   `json:"balance"`
	Sent                float64   `json:"sent"`
	Received            float64   `json:"received"`
	ColdStaking         bool      `json:"coldStaking"`
	ColdStakingBalance  float64   `json:"coldStakingBalance"`
	ColdStakingSent     float64   `json:"coldStakingSent"`
	ColdStakingReceived float64   `json:"coldStakingReceived"`
}

type Balance struct {
	Address           string  `json:"address"`
	Balance           float64 `json:"balance"`
	ColdStakedBalance float64 `json:"coldStakedBalance"`
}

type TransactionType string

const (
	TX_SEND         TransactionType = "SEND"
	TX_RECEIVE      TransactionType = "RECEIVE"
	TX_STAKING      TransactionType = "STAKING"
	TX_COLD_STAKING TransactionType = "COLD_STAKING"
)

func (e *ExplorerApi) GetAddresses(page int, size int) (addresses []Address, paginator Paginator, err error) {
	method := fmt.Sprintf("/address?page=%d&size=%d", page, size)

	response, paginator, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &addresses)
	return
}

func (e *ExplorerApi) GetAddress(hash string) (address Address, err error) {
	method := fmt.Sprintf("/address/%s", hash)

	response, _, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &address)
	return
}

func (e *ExplorerApi) ValidateAddress(hash string) (validateAddress ValidateAddress, err error) {
	method := fmt.Sprintf("/address/%s/validate", hash)

	response, _, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &validateAddress)
	return
}

func (e *ExplorerApi) GetAddressTransactions(hash string, filters []TransactionType, page int, size int) (transactions []Transaction, paginator Paginator, err error) {
	method := fmt.Sprintf("/address/%s/tx?page=%d&size=%d&filters=%s", hash, page, size, filtersToString(filters))

	response, paginator, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &transactions)
	return
}

func (e *ExplorerApi) GetBalances(addresses []string) (balances []Balance, err error) {
	method := fmt.Sprintf("/balance?addresses=%s", strings.Join(addresses, ","))

	response, _, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &balances)
	return
}

func filtersToString(filters []TransactionType) string {
	filterString := make([]string, len(filters))
	for i := range filters {
		filterString[i] = string(filters[i])
	}

	return strings.Join(filterString, ",")
}
