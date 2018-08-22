package teamcity

import (
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) CreateVcsRoot(vcs *types.VcsRoot) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/vcs-roots", c.version)
	var vcsReturn *types.VcsRoot

	err := c.doRetryRequest("POST", path, vcs, &vcsReturn)
	if err != nil {
		return err
	}

	if vcsReturn == nil {
		return errors.New("VCS Root not created")
	}
	*vcs = *vcsReturn

	return nil
}
