package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientSetBuildConfigurationTemplate(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplate")
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplateTemplate")
	require.NoError(t, err, "Expected no error")
	config := &types.BuildConfiguration{
		ProjectID:    "Empty",
		Name:         "TestClientSetBuildConfigurationTemplateTemplate",
		TemplateFlag: true,
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")
	config = &types.BuildConfiguration{
		ProjectID: "Empty",
		Name:      "TestClientSetBuildConfigurationTemplate",
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")
	assert.Equal(t, "", string(config.TemplateID), "Expected create to return empty TemplateID")

	err = client.SetBuildConfigurationTemplate("Empty_TestClientSetBuildConfigurationTemplate", "Empty_TestClientSetBuildConfigurationTemplateTemplate")
	require.NoError(t, err, "Expected no error")

	config, err = client.GetBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplate")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, "Empty_TestClientSetBuildConfigurationTemplateTemplate", string(config.TemplateID), "Expected get to return templateID")
}

func TestClientSetBuildConfigurationTemplateReset(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplateReset")
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplateResetTemplate")
	require.NoError(t, err, "Expected no error")
	config := &types.BuildConfiguration{
		ProjectID:    "Empty",
		Name:         "TestClientSetBuildConfigurationTemplateResetTemplate",
		TemplateFlag: true,
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")
	config = &types.BuildConfiguration{
		ProjectID:  "Empty",
		Name:       "TestClientSetBuildConfigurationTemplateReset",
		TemplateID: "Empty_TestClientSetBuildConfigurationTemplateResetTemplate",
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")
	assert.Equal(t, "Empty_TestClientSetBuildConfigurationTemplateResetTemplate", string(config.TemplateID), "Expected create to return TemplateID")

	err = client.SetBuildConfigurationTemplate("Empty_TestClientSetBuildConfigurationTemplateReset", "")
	require.NoError(t, err, "Expected no error")

	config, err = client.GetBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplateReset")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, "", string(config.TemplateID), "Expected create to return TemplateID")
}

func TestClientSetBuildConfigurationTemplateChange(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplateChange")
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplateChangeTemplate")
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplateChangeTemplate2")
	require.NoError(t, err, "Expected no error")
	config := &types.BuildConfiguration{
		ProjectID:    "Empty",
		Name:         "TestClientSetBuildConfigurationTemplateChangeTemplate",
		TemplateFlag: true,
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")
	config = &types.BuildConfiguration{
		ProjectID:    "Empty",
		Name:         "TestClientSetBuildConfigurationTemplateChangeTemplate2",
		TemplateFlag: true,
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")
	config = &types.BuildConfiguration{
		ProjectID:  "Empty",
		Name:       "TestClientSetBuildConfigurationTemplateChange",
		TemplateID: "Empty_TestClientSetBuildConfigurationTemplateChangeTemplate",
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")
	assert.Equal(t, "Empty_TestClientSetBuildConfigurationTemplateChangeTemplate", string(config.TemplateID), "Expected create to return TemplateID")

	err = client.SetBuildConfigurationTemplate("Empty_TestClientSetBuildConfigurationTemplateChange", "Empty_TestClientSetBuildConfigurationTemplateChangeTemplate2")
	require.NoError(t, err, "Expected no error")

	config, err = client.GetBuildConfiguration("Empty_TestClientSetBuildConfigurationTemplateChange")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, "Empty_TestClientSetBuildConfigurationTemplateChangeTemplate2", string(config.TemplateID), "Expected create to return TemplateID")
}
