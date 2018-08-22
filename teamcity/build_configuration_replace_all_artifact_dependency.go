package teamcity

import (
	"errors"
	"fmt"

	"github.com/jrasell/teamcity-sdk-go/types"
)

func (c *Client) ReplaceAllBuildConfigurationArtifactDependencies(buildConfID string, artifactDependencies *types.BuildArtifactDependencies) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/artifact-dependencies", c.version, buildConfID)
	var buildArtifactDependenciesReturn *types.BuildArtifactDependencies

	err := c.doRetryRequest("PUT", path, artifactDependencies, &buildArtifactDependenciesReturn)
	if err != nil {
		return err
	}

	if buildArtifactDependenciesReturn == nil {
		return errors.New("build configuration artifact dependencies not updated")
	}
	*artifactDependencies = *buildArtifactDependenciesReturn

	return nil
}
