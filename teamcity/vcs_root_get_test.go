package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientGetVcsRoot(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")

	vcs, err := client.GetVcsRoot("Empty_Rlink")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, vcs, "Create to return config")

	assert.Equal(t, "Empty_Rlink", vcs.ID, "Expected create to return ID")
	assert.Equal(t, "Rlink", vcs.Name, "Expected create to return Name")
	assert.Equal(t, "jetbrains.git", vcs.VcsName, "Expected create to return VcsName")
	assert.Equal(t, types.Properties{
		"url":                   "https://github.com/cardfree/teamcity-sdk-go",
		"usernameStyle":         "USERID",
		"agentCleanFilesPolicy": "ALL_UNTRACKED",
		"agentCleanPolicy":      "ON_BRANCH_CHANGE",
		"authMethod":            "ANONYMOUS",
		"branch":                "refs/heads/master",
		"submoduleCheckout":     "IGNORE",
	}, vcs.Properties, "no properties")

}

func TestClientGetVcsRootMissing(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")

	config, err := client.GetVcsRoot("Empty_")
	require.NoError(t, err, "Expected no error")
	require.Nil(t, config, "Expected no config")
}
