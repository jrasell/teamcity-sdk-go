package teamcity

import (
	"testing"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientGetBuildConfigurationMock(t *testing.T) {
	client := NewTestClient(newResponse(`{"parameters": {"property":[{"name": "build.counter", "value": "12"}]}}`), nil)

	config, err := client.GetBuildConfiguration("999999")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")

	assert.Equal(t, types.Parameters{
		"build.counter": types.Parameter{"12", nil},
	}, config.Parameters, "Parameters")
}

func TestClientGetBuildConfigurationMock1(t *testing.T) {
	client := NewTestClient(newResponse(`{"id":"Single_Normal",
		"name":"Normal",
		"projectName":"Single",
		"projectId":"Single",
		"href":"/httpAuth/app/rest/buildTypes/id:Single_Normal",
		"webUrl":"http://teamcity:8111/viewType.html?buildTypeId=Single_Normal",
		"project":{
			"id":"Single",
			"name":"Single",
			"parentProjectId":"_Root",
			"href":"/httpAuth/app/rest/projects/id:Single",
			"webUrl":"http://teamcity:8111/project.html?projectId=Single"
		},
		"vcs-root-entries":{
			"count":1,
			"vcs-root-entry":[{
				"id":"Single_HttpsGithubComUmweltdkDockerNodeGit",
				"vcs-root":{
					"id":"Single_HttpsGithubComUmweltdkDockerNodeGit",
					"name":"https://github.com/umweltdk/docker-node.git",
					"href":"/httpAuth/app/rest/vcs-roots/id:Single_HttpsGithubComUmweltdkDockerNodeGit"
				},
				"checkout-rules":""
			}]
		},
		"settings":{
			"count":17,
			"property":[{
				"name":"allowExternalStatus",
				"value":"false"
			},{
				"name":"allowPersonalBuildTriggering",
				"value":"true"
			},{
				"name":"artifactRules",
				"value":""
			},{
				"name":"buildNumberCounter",
				"value":"1"
			},{
				"name":"buildNumberPattern",
				"value":"%build.counter%"
			},{
				"name":"checkoutDirectory"
			},{
				"name":"checkoutMode",
				"value":"ON_SERVER"
			},{
				"name":"cleanBuild",
				"value":"false"
			},{
				"name":"enableHangingBuildsDetection",
				"value":"true"
			},{
				"name":"executionTimeoutMin",
				"value":"0"
			},{
				"name":"maximumNumberOfBuilds",
				"value":"0"
			},{
				"name":"shouldFailBuildIfTestsFailed",
				"value":"true"
			},{
				"name":"shouldFailBuildOnAnyErrorMessage",
				"value":"false"
			},{
				"name":"shouldFailBuildOnBadExitCode",
				"value":"true"
			},{
				"name":"shouldFailBuildOnOOMEOrCrash",
				"value":"true"
			},{
				"name":"showDependenciesChanges",
				"value":"false"
			},{
				"name":"vcsLabelingBranchFilter",
				"value":"+:<default>"
			}]
		},
		"parameters":{
			"count":0,
			"href":"/app/rest/buildTypes/id:Single_Normal/parameters",
			"property":[]
		},
		"steps":{
			"count":1,
			"step":[{
				"id":"RUNNER_1",
				"name":"Echo",
				"type":"simpleRunner",
				"properties":{
					"count":3,
					"property":[{
						"name":"script.content",
						"value":"env"
					},{
						"name":"teamcity.step.mode",
						"value":"default"
					},{
						"name":"use.custom.script",
						"value":"true"
					}]
				}
			}]
		},
		"features":{
			"count":0
		},
		"triggers":{
			"count":0
		},
		"snapshot-dependencies":{
			"count":0
		},
		"artifact-dependencies":{
			"count":0
		},
		"agent-requirements":{
			"count":0
		},
		"builds":{
			"href":"/httpAuth/app/rest/buildTypes/id:Single_Normal/builds/"
		}}`), nil)
	config, err := client.GetBuildConfiguration("999999")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")

	assert.Equal(t, 17, len(config.Settings), "Settings")
	assert.Equal(t, types.BuildSteps{
		types.BuildStep{
			ID:   "RUNNER_1",
			Name: "Echo",
			Type: "simpleRunner",
			Properties: types.Properties{
				"script.content":     "env",
				"teamcity.step.mode": "default",
				"use.custom.script":  "true",
			},
		},
	}, config.Steps, "Build steps")
}

func TestClientGetBuildConfiguration(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")

	config, err := client.GetBuildConfiguration("Single_Normal")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")

	assert.Equal(t, types.BuildSteps{
		types.BuildStep{
			ID:   "RUNNER_1",
			Name: "Echo",
			Type: "simpleRunner",
			Properties: types.Properties{
				"script.content":     "env",
				"teamcity.step.mode": "default",
				"use.custom.script":  "true",
			},
		},
	}, config.Steps, "Build steps")
}

func TestClientGetBuildConfigurationMissing(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")

	config, err := client.GetBuildConfiguration("Single_Norm")
	require.NoError(t, err, "Expected no error")
	require.Nil(t, config, "Expected no config")
}
