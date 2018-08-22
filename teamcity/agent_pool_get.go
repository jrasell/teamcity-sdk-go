package teamcity

import (
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) GetAgentPoolById(pool int) (*types.AgentPools, error) {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/agentPools/id:%d", c.version, pool)
	var agp *types.AgentPools

	err := c.doRetryRequest("GET", path, nil, &agp)
	if err != nil {
		return nil, err
	}

	return agp, nil
}

func (c *Client) GetAgentPoolByName(pool string) (*types.AgentPools, error) {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/agentPools/name:%s", c.version, pool)
	var agp *types.AgentPools

	err := c.doRetryRequest("GET", path, nil, &agp)
	if err != nil {
		return nil, err
	}

	return agp, nil
}
