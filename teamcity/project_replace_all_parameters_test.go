package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientReplaceAllProjectParameters(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteProject("TestClientReplaceAllProjectParameters")
	require.NoError(t, err, "Expected no error")
	err = client.CreateProject(&types.Project{
		Name: "TestClientReplaceAllProjectParameters",
	})
	require.NoError(t, err, "Expected no error")
	err = client.ReplaceAllProjectParameters("TestClientReplaceAllProjectParameters", &types.Parameters{
		"env.MUH": types.Parameter{
			Value: "Hello",
			Spec: &types.ParameterSpec{
				Type: types.PasswordType{},
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
	err = client.ReplaceAllProjectParameters("TestClientReplaceAllProjectParameters", &parameters)
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

	config, err := client.GetProject("TestClientReplaceAllProjectParameters")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, expected, config.Parameters)
}

func TestClientReplaceProjectParametersInherited(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteProject("Empty3")
	require.NoError(t, err, "Expected no error")
	err = client.CreateProject(&types.Project{
		Name: "Empty3",
	})
	require.NoError(t, err, "Expected no error")
	err = client.ReplaceAllProjectParameters("Empty3", &types.Parameters{
		"env.MUH": types.Parameter{
			Value: "Hush",
			Spec: &types.ParameterSpec{
				Label:       "Muh value",
				Description: "The Muh value that does all the Muhing",
				Type:        types.PasswordType{},
			},
		},
		"config.inherited": types.Parameter{
			Value: "Parent",
			Spec: &types.ParameterSpec{
				Label: "AWW",
				Type:  types.CheckboxType{"Hello", "Copperhead"},
			},
		},
		"config.inherited2": types.Parameter{
			Value: "Parent",
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
	err = client.DeleteProject("Empty3_TestClientReplaceProjectParametersInherited")
	require.NoError(t, err, "Expected no error")
	err = client.CreateProject(&types.Project{
		ID:              "Empty3_TestClientReplaceProjectParametersInherited",
		ParentProjectID: "Empty3",
		Name:            "TestClientReplaceProjectParametersInherited",
	})
	require.NoError(t, err, "Expected no error")
	err = client.ReplaceAllProjectParameters("Empty3_TestClientReplaceProjectParametersInherited", &types.Parameters{
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
		"config.inherited": types.Parameter{
			Value: "Dink",
			Spec: &types.ParameterSpec{
				Label: "Buhhhhh",
				Type:  types.TextType{"any"},
			},
		},
		"config.inherited2": types.Parameter{
			Value: "Dink",
		},
		"env.AWW": types.Parameter{
			Value: "BuildConf",
		},
		"env.DAMM": types.Parameter{
			Value: "BuildConf",
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
		"config.inherited": types.Parameter{
			Value: "Dink",
			Spec: &types.ParameterSpec{
				Label: "Buhhhhh",
				Type:  types.TextType{"any"},
			},
		},
		"config.inherited2": types.Parameter{
			Value: "Dink",
			Spec: &types.ParameterSpec{
				Label: "Buhhhhh",
				Type:  types.TextType{"any"},
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
	err = client.ReplaceAllProjectParameters("Empty3_TestClientReplaceProjectParametersInherited", &parameters)
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
		"config.inherited": types.Parameter{
			Value: "Dink",
			Spec: &types.ParameterSpec{
				Label: "AWW",
				Type:  types.CheckboxType{"Hello", "Copperhead"},
			},
		},
		"config.inherited2": types.Parameter{
			Value: "Dink",
			Spec: &types.ParameterSpec{
				Label: "Buhhhhh",
				Type:  types.TextType{"any"},
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

	config, err := client.GetProject("Empty3_TestClientReplaceProjectParametersInherited")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")
	assert.Equal(t, expected, config.Parameters)
}
