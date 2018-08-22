package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientAllReplaceBuildConfigurationParameters(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Single_TestClientUpdateBuildConfiguration")
	require.NoError(t, err, "Expected no error")
	err = client.CreateBuildConfiguration(&types.BuildConfiguration{
		ID:        "Single_TestClientUpdateBuildConfiguration",
		ProjectID: "Single",
		Name:      "Test Client Update Build Configuration",
		Parameters: types.Parameters{
			"env.MUH": types.Parameter{
				Value: "Hello",
				Spec: &types.ParameterSpec{
					Type: types.PasswordType{},
				},
			},
		},
	})
	require.NoError(t, err, "Expected no error")

	parameters := types.Parameters{
		"env.HELLO": types.Parameter{"Good job", nil},
		"aws.hush": types.Parameter{
			Value: "Bad Job",
			Spec: &types.ParameterSpec{
				Type: types.PasswordType{},
			},
		},
	}
	err = client.ReplaceAllBuildConfigurationParameters("Single_TestClientUpdateBuildConfiguration", &parameters)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, parameters, "Update to return parameters")

	expected := types.Parameters{
		"env.HELLO": types.Parameter{"Good job", nil},
		"aws.hush": types.Parameter{
			Value: "",
			Spec: &types.ParameterSpec{
				Type: types.PasswordType{},
			},
		},
	}
	assert.Equal(t, expected, parameters)

	config, err := client.GetBuildConfiguration("Single_TestClientUpdateBuildConfiguration")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, expected, config.Parameters)
}

func TestClientReplaceAllBuildConfigurationParametersInherited(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteProject("Empty2")
	require.NoError(t, err, "Expected no error")
	err = client.CreateProject(&types.Project{
		ParentProjectID: "",
		Name:            "Empty2",
	})
	require.NoError(t, err, "Expected no error")
	err = client.ReplaceAllProjectParameters("Empty2", &types.Parameters{
		"env.MUH": types.Parameter{
			Value: "Hush",
			Spec: &types.ParameterSpec{
				Label:       "Muh value",
				Description: "The Muh value that does all the Muhing",
				Type:        types.PasswordType{},
			},
		},
		"env.AWW": types.Parameter{
			Value: "Parent",
			Spec: &types.ParameterSpec{
				Label: "AWW",
				Type:  types.TextType{"any"},
			},
		},
		"env.DAMM": types.Parameter{
			Value: "Parent",
		},
	})
	err = client.DeleteBuildConfiguration("Empty2_TestClientUpdateBuildConfigurationInherited")
	require.NoError(t, err, "Expected no error")
	err = client.CreateBuildConfiguration(&types.BuildConfiguration{
		ID:        "Empty2_TestClientUpdateBuildConfigurationInherited",
		ProjectID: "Empty2",
		Name:      "Test Client Update Build Configuration",
		Parameters: types.Parameters{
			"config.remove": types.Parameter{
				Value: "Hello",
			},
			"config.replace": types.Parameter{
				Value: "Dink",
				Spec: &types.ParameterSpec{
					Label: "Buhhhhh",
					Type:  types.TextType{"any"},
				},
			},
			"env.AWW": types.Parameter{
				Value: "BuildConf",
			},
			"env.DAMM": types.Parameter{
				Value: "BuildConf",
			},
		},
	})
	require.NoError(t, err, "Expected no error")

	parameters := types.Parameters{
		"env.HELLO": types.Parameter{"Good job", nil},
		"config.replace": types.Parameter{
			Value: "Mink",
			Spec: &types.ParameterSpec{
				Label: "Minker",
				Type: types.CheckboxType{
					Checked: "Flunk",
				},
			},
		},
		"aws.hush": types.Parameter{
			Value: "Bad Job",
			Spec: &types.ParameterSpec{
				Type: types.PasswordType{},
			},
		},
		"env.AWW": types.Parameter{
			Value: "",
			Spec: &types.ParameterSpec{
				Type: types.CheckboxType{
					Checked: "BuildCD",
				},
			},
		},
		"env.MUH": types.Parameter{
			Value: "Hello",
			Spec: &types.ParameterSpec{
				Label: "Plunk",
				Type:  types.PasswordType{},
			},
		},
	}
	err = client.ReplaceAllBuildConfigurationParameters("Empty2_TestClientUpdateBuildConfigurationInherited", &parameters)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, parameters, "Update to return parameters")

	expected := types.Parameters{
		"env.HELLO": types.Parameter{"Good job", nil},
		"config.replace": types.Parameter{
			Value: "Mink",
			Spec: &types.ParameterSpec{
				Label: "Minker",
				Type: types.CheckboxType{
					Checked: "Flunk",
				},
			},
		},
		"aws.hush": types.Parameter{
			Value: "",
			Spec: &types.ParameterSpec{
				Type: types.PasswordType{},
			},
		},
		"env.AWW": types.Parameter{
			Value: "",
			Spec: &types.ParameterSpec{
				Label: "AWW",
				Type:  types.TextType{"any"},
			},
		},
		"env.MUH": types.Parameter{
			Value: "",
			Spec: &types.ParameterSpec{
				Label:       "Muh value",
				Description: "The Muh value that does all the Muhing",
				Type:        types.PasswordType{},
			},
		},
		"env.DAMM": types.Parameter{
			Value: "Parent",
		},
	}
	assert.Equal(t, expected, parameters)

	config, err := client.GetBuildConfiguration("Empty2_TestClientUpdateBuildConfigurationInherited")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, expected, config.Parameters)
}
