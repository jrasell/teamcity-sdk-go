package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientGetProject(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")

	config, err := client.GetProject("Empty")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")

	assert.Equal(t, types.Parameters{
		"env.MUH": types.Parameter{
			Value: client.VersionParameterValue(t, "env.MUH"),
			Spec: &types.ParameterSpec{
				Label:       "Muh value",
				Description: "The Muh value that does all the Muhing",
				Type:        types.PasswordType{},
			},
		},
	}, config.Parameters, "Parameters")
}

func TestClientGetProjectMissing(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")

	config, err := client.GetProject("Empt")
	require.NoError(t, err, "Expected no error")
	require.Nil(t, config, "Expected no config")
}
