package teamcity

import (
	"testing"
	"time"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientCreateProjectMock(t *testing.T) {
	client := NewTestClient(newResponse(`{"id": "Empty_Plink", "parentProjectId":"Empty","name":"Plink"}`), nil)

	project := &types.Project{
		ParentProjectID: "Empty",
		Name:            "Plink",
	}

	err := client.CreateProject(project)
	require.NoError(t, err, "Expected no error")

	assert.Equal(t, "Empty_Plink", project.ID, "Expected create to return ID")
}

func TestClientCreateProjectMinimal(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteProject("Empty_Hello")
	require.NoError(t, err, "Expected no error")
	time.Sleep(5 * time.Second)

	project := &types.Project{
		ParentProjectID: "Empty",
		Name:            "Hello",
		Description:     "Look over there",
	}
	err = client.CreateProject(project)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, project, "Create to return project")

	assert.Equal(t, "Empty_Hello", project.ID, "Expected create to return ID")
	assert.Equal(t, types.ProjectId("Empty"), project.ParentProjectID, "Expected create to return ParentProjectID")
	assert.Equal(t, "Hello", project.Name, "Expected create to return Name")
	assert.Equal(t, "", project.Description, "Expected create to return Description")
	assert.Equal(t, types.Parameters{
		"env.MUH": types.Parameter{
			Value: client.VersionParameterValue(t, "env.MUH"),
			Spec: &types.ParameterSpec{
				Label:       "Muh value",
				Description: "The Muh value that does all the Muhing",
				Display:     types.Normal,
				Type:        types.PasswordType{},
			},
		},
	}, project.Parameters, "no parameters")

	project, err = client.GetProject("Empty_Hello")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, project, "Get to return project")
	assert.Equal(t, types.ProjectId("Empty"), project.ParentProjectID, "Expected get to return ParentProjectID")
	assert.Equal(t, "Hello", project.Name, "Expected get to return Name")
	assert.Equal(t, "", project.Description, "Expected Empty to return Description")
	assert.Equal(t, types.Parameters{
		"env.MUH": types.Parameter{
			Value: client.VersionParameterValue(t, "env.MUH"),
			Spec: &types.ParameterSpec{
				Label:       "Muh value",
				Description: "The Muh value that does all the Muhing",
				Display:     types.Normal,
				Type:        types.PasswordType{},
			},
		},
	}, project.Parameters, "no parameters")
}

func TestClientCreateProjectIgnoresParameters(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteProject("Empty_Full")
	require.NoError(t, err, "Expected no error")
	time.Sleep(5 * time.Second)

	project := &types.Project{
		ParentProjectID: "Empty",
		Name:            "Full",
		Parameters: types.Parameters{
			"env.AWW": types.Parameter{
				Value: "Parent",
			},
		},
	}
	err = client.CreateProject(project)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, project, "Create to return project")

	assert.Equal(t, "Empty_Full", project.ID, "Expected create to return ID")
	assert.Equal(t, types.Parameters{
		"env.MUH": types.Parameter{
			Value: client.VersionParameterValue(t, "env.MUH"),
			Spec: &types.ParameterSpec{
				Label:       "Muh value",
				Description: "The Muh value that does all the Muhing",
				Display:     types.Normal,
				Type:        types.PasswordType{},
			},
		},
	}, project.Parameters, "no parameters")
}

func TestClientCreateProjectUsedID(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	client.retries = 1

	project := &types.Project{
		ID:   "Single",
		Name: "Hej Med Dig",
	}

	err = client.CreateProject(project)
	assert.Error(t, err, "Expected error")
}

func TestClientCreateProjectUsedName(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	client.retries = 1

	project := &types.Project{
		Name: "Single",
	}

	err = client.CreateProject(project)
	assert.Error(t, err, "Expected error")
}

func TestClientCreateProjectUsedNameExplicitID(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	client.retries = 1

	project := &types.Project{
		ID:   "Single_Dubie",
		Name: "Single",
	}

	err = client.CreateProject(project)
	assert.Error(t, err, "Expected error")
}
