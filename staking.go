package navexplorer

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Reward struct {
	Address string         `json:"address"`
	Periods []RewardPeriod `json:"periods"`
}

type RewardPeriod struct {
	Period  string `json:"period"`
	Stakes  int64  `json:"stakes"`
	Balance int64  `json:"balance"`
}

func (e *ExplorerApi) GetStakingRewardsForAddresses(addresses []string) (rewards []Reward, err error) {
	method := fmt.Sprintf("/api/staking/rewards?addresses=%s", strings.Join(addresses, ","))
	log.Print(method)

	response, _, err := e.client.call(method)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &rewards)
	return
}
