package teamcity

import (
	"bytes"
	"fmt"
)

func (c *Client) ReplaceBuildConfigurationField(buildConfID, name string, value string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/%s", c.version, buildConfID, name)

	fmt.Printf("Replace build config field %s\n", value)
	body := bytes.NewBuffer([]byte(value))
	_, err := c.doNotJSONRequest("PUT", path, "text/plain", "text/plain", body)
	if err != nil {
		return err
	}
	return nil
}
