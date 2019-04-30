package navexplorer

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Address struct {
	Hash               string  `json:"hash"`
	Received           float64 `json:"received"`
	ReceivedCount      int     `json:"receivedCount"`
	Sent               float64 `json:"sent"`
	SentCount          int     `json:"sentCount"`
	Staked             float64 `json:"staked"`
	StakedCount        int     `json:"stakedCount"`
	StakedSent         float64 `json:"stakedSent"`
	StakedReceived     float64 `json:"stakedReceived"`
	ColdStaked         float64 `json:"coldStaked"`
	ColdStakedCount    int     `json:"coldStakedCount"`
	ColdStakedSent     float64 `json:"coldStakedSent"`
	ColdStakedReceived float64 `json:"coldStakedReceived"`
	ColdStakedBalance  float64 `json:"coldStakedBalance"`
	Balance            float64 `json:"balance"`
	BlockIndex         int     `json:"blockIndex"`
	RichListPosition   int64   `json:"richListPosition"`
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

type TransactionType string

const (
	TX_SEND         TransactionType = "SEND"
	TX_RECEIVE      TransactionType = "RECEIVE"
	TX_STAKING      TransactionType = "STAKING"
	TX_COLD_STAKING TransactionType = "COLD_STAKING"
)

func (e *ExplorerApi) GetAddresses(page int, size int) (addresses []Address, paginator Paginator, err error) {
	method := fmt.Sprintf("/api/address?page=%d&size=%d", page, size)

	response, paginator, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &addresses)
	return
}

func (e *ExplorerApi) GetAddress(hash string) (address Address, err error) {
	method := fmt.Sprintf("/api/address/%s", hash)

	response, _, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &address)
	return
}

func (e *ExplorerApi) GetAddressTransactions(hash string, filters []TransactionType, page int, size int) (transactions []Transaction, paginator Paginator, err error) {
	method := fmt.Sprintf("/api/address/%s/tx?page=%d&size=%d&filters=%s", hash, page, size, filtersToString(filters))

	response, paginator, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &transactions)
	return
}

func (e *ExplorerApi) GetAddressColdTransactions(hash string, filters []TransactionType, page int, size int) (transactions []Transaction, paginator Paginator, err error) {
	method := fmt.Sprintf("/api/address/%s/coldtx?page=%d&size=%d&filters=%s", hash, page, size, filtersToString(filters))

	response, paginator, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &transactions)
	return
}

func filtersToString(filters []TransactionType) string {
	filterString := make([]string, len(filters))
	for i := range filters {
		filterString[i] = string(filters[i])
	}

	return strings.Join(filterString, ",")
}
