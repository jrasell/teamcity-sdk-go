package teamcity

import (
	//"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) GetBuildConfiguration(buildConfID string) (*types.BuildConfiguration, error) {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s", c.version, buildConfID)
	var buildConfig *types.BuildConfiguration

	err := c.doRetryRequest("GET", path, nil, &buildConfig)
	if err != nil {
		return nil, err
	}

	/*if buildConfig == nil {
		return nil, errors.New("build configuration not found")
	}*/

	return buildConfig, nil
}
