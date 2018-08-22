package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientSetBuildConfigurationDescription(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TestClientSetBuildConfigurationDescription")
	require.NoError(t, err, "Expected no error")
	config := &types.BuildConfiguration{
		ProjectID:   "Empty",
		Name:        "TestClientSetBuildConfigurationDescription",
		Description: "Love is in the air",
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")
	assert.Equal(t, "Love is in the air", config.Description, "Expected create to return Description")

	err = client.SetBuildConfigurationDescription("Empty_TestClientSetBuildConfigurationDescription", "Nok")
	require.NoError(t, err, "Expected no error")

	config, err = client.GetBuildConfiguration("Empty_TestClientSetBuildConfigurationDescription")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, "Nok", config.Description, "Expected get to return Description")
}

func TestClientSetBuildConfigurationDescriptionReset(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TestClientSetBuildConfigurationDescriptionReset")
	require.NoError(t, err, "Expected no error")
	config := &types.BuildConfiguration{
		ProjectID:   "Empty",
		Name:        "TestClientSetBuildConfigurationDescriptionReset",
		Description: "Love is in the air",
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")
	assert.Equal(t, "Love is in the air", config.Description, "Expected create to return Description")

	err = client.SetBuildConfigurationDescription("Empty_TestClientSetBuildConfigurationDescriptionReset", "")
	require.NoError(t, err, "Expected no error")

	config, err = client.GetBuildConfiguration("Empty_TestClientSetBuildConfigurationDescriptionReset")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, "", config.Description, "Expected get to return Description")
}
