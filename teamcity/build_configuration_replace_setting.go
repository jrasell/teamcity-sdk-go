package teamcity

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) ReplaceBuildConfigurationSetting(buildConfID, name string, value string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/settings/%s", c.version, buildConfID, name)
	var settingReturn *types.BuildSetting
	actual := types.BuildSetting{
		Name:  name,
		Value: value,
	}

	sd, _ := json.Marshal(&actual)
	fmt.Printf("Replace build config setting %s\n", string(sd))
	err := c.doRetryRequest("PUT", path, actual, &settingReturn)
	if err != nil {
		return err
	}

	if settingReturn == nil {
		return errors.New("build configuration setting not updated")
	}

	return nil
}
