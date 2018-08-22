package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientAllReplaceBuildConfigurationParameter(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Single_TestClientAllReplaceBuildConfigurationParameter")
	require.NoError(t, err, "Expected no error")
	err = client.CreateBuildConfiguration(&types.BuildConfiguration{
		ID:        "Single_TestClientAllReplaceBuildConfigurationParameter",
		ProjectID: "Single",
		Name:      "TestClientAllReplaceBuildConfigurationParameter",
		Parameters: types.Parameters{
			"env.HELLO": types.Parameter{"Good job", nil},
			"env.MUH": types.Parameter{
				Value: "Hello",
				Spec: &types.ParameterSpec{
					Type: types.PasswordType{},
				},
			},
		},
	})
	require.NoError(t, err, "Expected no error")

	err = client.ReplaceBuildConfigurationParameter("Single_TestClientAllReplaceBuildConfigurationParameter",
		"env.MUH",
		&types.Parameter{
			Value: "Bad Job",
			Spec: &types.ParameterSpec{
				Type: types.TextType{"not_empty"},
			},
		})
	require.NoError(t, err, "Expected no error")

	expected := types.Parameters{
		"env.HELLO": types.Parameter{"Good job", nil},
		"env.MUH": types.Parameter{
			Value: "Bad Job",
			Spec: &types.ParameterSpec{
				Type: types.TextType{"not_empty"},
			},
		},
	}
	config, err := client.GetBuildConfiguration("Single_TestClientAllReplaceBuildConfigurationParameter")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, expected, config.Parameters)
}
