package teamcity

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) ReplaceBuildConfigurationParameter(buildConfID, name string, parameter *types.Parameter) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/parameters/%s", c.version, buildConfID, name)
	var parameterReturn *types.NamedParameter
	actual := types.NamedParameter{
		Name:      name,
		Parameter: *parameter,
	}

	sd, _ := json.Marshal(&actual)
	fmt.Printf("Replace build config parameter %s\n", string(sd))
	err := c.doRetryRequest("PUT", path, actual, &parameterReturn)
	if err != nil {
		return err
	}

	if parameterReturn == nil {
		return errors.New("build configuration parameter not updated")
	}
	*parameter = parameterReturn.Parameter

	return nil
}
