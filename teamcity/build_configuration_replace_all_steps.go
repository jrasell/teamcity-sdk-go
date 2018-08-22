package teamcity

import (
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllBuildConfigurationSteps(buildConfID string, steps *types.BuildSteps) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/steps", c.version, buildConfID)
	var buildstepsReturn *types.BuildSteps

	err := c.doRetryRequest("PUT", path, steps, &buildstepsReturn)
	if err != nil {
		return err
	}

	if buildstepsReturn == nil {
		return errors.New("build configuration steps not updated")
	}
	*steps = *buildstepsReturn

	return nil
}
