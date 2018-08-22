package teamcity

import (
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) AttachBuildConfigurationVcsRoot(buildConfID string, vcsRoot *types.VcsRootEntry) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/vcs-root-entries", c.version, buildConfID)
	var vcsRootReturn *types.VcsRootEntry

	err := c.doRetryRequest("POST", path, vcsRoot, &vcsRootReturn)
	if err != nil {
		return err
	}

	if vcsRootReturn == nil {
		return errors.New("vcs root entry not created")
	}
	*vcsRoot = *vcsRootReturn

	return nil
}
