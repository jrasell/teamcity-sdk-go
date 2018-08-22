package teamcity

import (
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllBuildConfigurationParameters(buildConfID string, parameters *types.Parameters) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/parameters", c.version, buildConfID)
	var parametersReturn *types.Parameters

	err := c.doRetryRequest("PUT", path, parameters, &parametersReturn)
	if err != nil {
		return err
	}

	if parametersReturn == nil {
		return errors.New("build configuration not updated")
	}
	*parameters = *parametersReturn

	return nil
}
