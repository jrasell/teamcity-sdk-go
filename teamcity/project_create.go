package teamcity

import (
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) CreateProject(project *types.Project) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/projects", c.version)
	var projectReturn *types.Project

	err := c.doRetryRequest("POST", path, project, &projectReturn)
	if err != nil {
		return err
	}

	if projectReturn == nil {
		return errors.New("project not created")
	}
	*project = *projectReturn

	return nil
}
