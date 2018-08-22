package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientSetProjectDescription(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteProject("Empty_TestClientSetProjectDescription")
	require.NoError(t, err, "Expected no error")
	project := &types.Project{
		ParentProjectID: "Empty",
		Name:            "TestClientSetProjectDescription",
		Description:     "Love is in the air",
	}
	err = client.CreateProject(project)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, project, "Create to return project")
	assert.Equal(t, "", project.Description, "Expected create to return Description")

	err = client.SetProjectDescription("Empty_TestClientSetProjectDescription", "Nok")
	require.NoError(t, err, "Expected no error")

	project, err = client.GetProject("Empty_TestClientSetProjectDescription")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, project, "Get to return project")
	assert.Equal(t, "Nok", project.Description, "Expected get to return Description")
}

func TestClientSetProjectDescriptionReset(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteProject("Empty_TestClientSetProjectDescriptionReset")
	require.NoError(t, err, "Expected no error")
	project := &types.Project{
		ParentProjectID: "Empty",
		Name:            "TestClientSetProjectDescriptionReset",
		Description:     "Love is in the air",
	}
	err = client.CreateProject(project)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, project, "Create to return project")
	assert.Equal(t, "", project.Description, "Expected create to return Description")

	err = client.SetProjectDescription("Empty_TestClientSetProjectDescriptionReset", "")
	require.NoError(t, err, "Expected no error")

	project, err = client.GetProject("Empty_TestClientSetProjectDescriptionReset")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, project, "Get to return project")
	assert.Equal(t, "", project.Description, "Expected get to return Description")
}
