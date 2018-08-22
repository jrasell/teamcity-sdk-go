package teamcity

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) ReplaceProjectParameter(projectID, name string, parameter *types.Parameter) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/projects/id:%s/parameters/%s", c.version, projectID, name)
	var parameterReturn *types.NamedParameter
	actual := types.NamedParameter{
		Name:      name,
		Parameter: *parameter,
	}

	sd, _ := json.Marshal(&actual)
	fmt.Printf("Replace project parameter %s\n", string(sd))
	err := c.doRetryRequest("PUT", path, actual, &parameterReturn)
	if err != nil {
		return err
	}

	if parameterReturn == nil {
		return errors.New("project parameter not updated")
	}
	*parameter = parameterReturn.Parameter

	return nil
}
