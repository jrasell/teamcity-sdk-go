package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientDetachBuildConfigurationVcsRoot(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Single_TestClientDetachBuildConfigurationVcsRoot")
	require.NoError(t, err, "Expected no error")
	err = client.CreateBuildConfiguration(&types.BuildConfiguration{
		ID:        "Single_TestClientDetachBuildConfigurationVcsRoot",
		ProjectID: "Single",
		Name:      "TestClientDetachBuildConfigurationVcsRoot",
		VcsRootEntries: types.VcsRootEntries{
			types.VcsRootEntry{
				ID:            "Flflfl",
				VcsRootID:     "Single_HttpsGithubComUmweltdkDockerNodeGit",
				CheckoutRules: "+:refs/heads/master\n+:refs/heads/trigger*",
			},
		},
	})
	require.NoError(t, err, "Expected no error")

	err = client.DetachBuildConfigurationVcsRoot("Single_TestClientDetachBuildConfigurationVcsRoot", "Single_HttpsGithubComUmweltdkDockerNodeGit")
	require.NoError(t, err, "Expected no error")

	config, err := client.GetBuildConfiguration("Single_TestClientDetachBuildConfigurationVcsRoot")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, make(types.VcsRootEntries, 0), config.VcsRootEntries)
}

func TestClientDetachBuildConfigurationVcsRootMissing(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Single_TestClientDetachBuildConfigurationVcsRootMissing")
	require.NoError(t, err, "Expected no error")
	err = client.CreateBuildConfiguration(&types.BuildConfiguration{
		ID:        "Single_TestClientDetachBuildConfigurationVcsRootMissing",
		ProjectID: "Single",
		Name:      "TestClientDetachBuildConfigurationVcsRootMissing",
	})
	require.NoError(t, err, "Expected no error")

	err = client.DetachBuildConfigurationVcsRoot("Single_TestClientDetachBuildConfigurationVcsRootMissing", "Single_HttpsGithubComUmweltdkDockerNodeGit")
	require.NoError(t, err, "Expected no error")

	config, err := client.GetBuildConfiguration("Single_TestClientDetachBuildConfigurationVcsRootMissing")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, make(types.VcsRootEntries, 0), config.VcsRootEntries)
}
