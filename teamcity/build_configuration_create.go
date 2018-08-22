package teamcity

import (
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) CreateBuildConfiguration(buildConfig *types.BuildConfiguration) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes", c.version)
	var buildConfigReturn *types.BuildConfiguration

	err := c.doRetryRequest("POST", path, buildConfig, &buildConfigReturn)
	if err != nil {
		return err
	}

	if buildConfigReturn == nil {
		return errors.New("build configuration not created")
	}
	*buildConfig = *buildConfigReturn

	return nil
}
