package teamcity

import (
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllVcsRootProperties(VcsRootId string, properties *types.Properties) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/vcs-roots/id:%s/properties", c.version, VcsRootId)
	var propertiesReturn *types.Properties

	err := c.doRetryRequest("PUT", path, properties, &propertiesReturn)
	if err != nil {
		return err
	}

	if propertiesReturn == nil {
		return errors.New("VCS Root configuration properties not updated")
	}
	*properties = *propertiesReturn

	return nil
}
