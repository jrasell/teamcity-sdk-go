package teamcity

import (
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllBuildConfigurationSnapshotDependencies(buildConfID string, snapshotDependencies *types.BuildSnapshotDependencies) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/snapshot-dependencies", c.version, buildConfID)
	var buildSnapshotDependenciesReturn *types.BuildSnapshotDependencies

	err := c.doRetryRequest("PUT", path, snapshotDependencies, &buildSnapshotDependenciesReturn)
	if err != nil {
		return err
	}

	if buildSnapshotDependenciesReturn == nil {
		return errors.New("build configuration snapshot dependencies not updated")
	}
	*snapshotDependencies = *buildSnapshotDependenciesReturn

	return nil
}
