package teamcity

import (
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) GetProject(projectID string) (*types.Project, error) {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/projects/id:%s", c.version, projectID)
	var project *types.Project

	err := c.doRetryRequest("GET", path, nil, &project)
	if err != nil {
		return nil, err
	}

	return project, nil
}
