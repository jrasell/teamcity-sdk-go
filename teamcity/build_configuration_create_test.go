package teamcity

import (
	"testing"
	"time"

	"github.com/jrasell/teamcity-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientCreateBuildConfigurationMock(t *testing.T) {
	client := NewTestClient(newResponse(`{"id": "Empty_Hello", "projectId":"Empty","templateFlag":false,"name":"Hello"}`), nil)

	config := &types.BuildConfiguration{
		ProjectID: "Empty",
		Name:      "Hello",
	}

	err := client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")

	assert.Equal(t, "Empty_Hello", config.ID, "Expected create to return ID")
}

func TestClientCreateBuildConfigurationMinimal(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_Hello")
	require.NoError(t, err, "Expected no error")

	config := &types.BuildConfiguration{
		ProjectID:   "Empty",
		Name:        "Hello",
		Description: "Love is in the air",
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")

	assert.Equal(t, "Empty_Hello", config.ID, "Expected create to return ID")
	assert.Equal(t, "Empty", config.ProjectID, "Expected create to return ProjectID")
	assert.Equal(t, "Hello", config.Name, "Expected create to return Name")
	assert.Equal(t, "Love is in the air", config.Description, "Expected create to return Description")
	assert.Equal(t, make(types.BuildSteps, 0), config.Steps, "no steps")

	config, err = client.GetBuildConfiguration("Empty_Hello")
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Get to return config")

	assert.Equal(t, "Empty", config.ProjectID, "Expected get to return ProjectID")
	assert.Equal(t, "Hello", config.Name, "Expected get to return Name")
	assert.Equal(t, "Love is in the air", config.Description, "Expected get to return Description")
	assert.Equal(t, make(types.BuildSteps, 0), config.Steps, "no steps")
}

func TestClientCreateTemplate(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Single_Templer")
	require.NoError(t, err, "Expected no error")

	config := &types.BuildConfiguration{
		ProjectID:    "Single",
		Name:         "Templer",
		TemplateFlag: true,
		VcsRootEntries: types.VcsRootEntries{
			types.VcsRootEntry{
				VcsRootID:     "Single_HttpsGithubComUmweltdkDockerNodeGit",
				CheckoutRules: "+:refs/heads/master\n+:refs/heads/trigger*",
			},
		},
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")

	assert.Equal(t, "Single_Templer", config.ID, "Expected create to return ID")
	assert.Equal(t, make(types.BuildSteps, 0), config.Steps, "no steps")
	assert.Equal(t, types.VcsRootEntries{
		types.VcsRootEntry{
			ID:            "Single_HttpsGithubComUmweltdkDockerNodeGit",
			VcsRootID:     "Single_HttpsGithubComUmweltdkDockerNodeGit",
			CheckoutRules: "+:refs/heads/master\n+:refs/heads/trigger*",
		},
	}, config.VcsRootEntries, "vcs root entries")

	loaded, err := client.GetBuildConfiguration("Single_Templer")
	assert.Equal(t, true, config.TemplateFlag, "Expected template")
	assert.Equal(t, loaded.VcsRootEntries, config.VcsRootEntries, "vcs root entries")
}

func TestClientCreateBuildConfigurationTemplateFull(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TemplateFull")
	require.NoError(t, err, "Expected no error")
	time.Sleep(5 * time.Second)

	config := &types.BuildConfiguration{
		ProjectID:  "Empty",
		Name:       "Template Full",
		TemplateID: "Tempy",
		Steps: types.BuildSteps{
			types.BuildStep{
				Name: "Muh",
				Type: "simpleRunner",
				Properties: types.Properties{
					"script.content":     "env",
					"teamcity.step.mode": "default",
					"use.custom.script":  "true",
				},
			},
			types.BuildStep{
				Name: "Env",
				Type: "simpleRunner",
				Properties: types.Properties{
					"script.content":     "env",
					"teamcity.step.mode": "default",
					"use.custom.script":  "true",
				},
			},
		},
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")

	var id1 string
	var id2 string
	if len(config.Steps) >= 3 {
		id1 = config.Steps[1].ID
		id2 = config.Steps[2].ID
	}

	assert.Equal(t, "Empty_TemplateFull", config.ID, "Supplied ID")
	assert.Equal(t, types.BuildSteps{
		types.BuildStep{
			ID:   "RUNNER_3",
			Name: "Env",
			Type: "simpleRunner",
			Properties: types.Properties{
				"script.content":     "env",
				"teamcity.step.mode": "default",
				"use.custom.script":  "true",
			},
		},
		types.BuildStep{
			ID:   id1,
			Name: "Muh",
			Type: "simpleRunner",
			Properties: types.Properties{
				"script.content":     "env",
				"teamcity.step.mode": "default",
				"use.custom.script":  "true",
			},
		},
		types.BuildStep{
			ID:   id2,
			Name: "Env (1)",
			Type: "simpleRunner",
			Properties: types.Properties{
				"script.content":     "env",
				"teamcity.step.mode": "default",
				"use.custom.script":  "true",
			},
		},
	}, config.Steps, "Build steps")
}

func TestClientCreateBuildConfigurationTemplateReorder(t *testing.T) {
	client, err := NewRealTestClient(t)
	client.SkipOlder(t, 10, 0)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TemplateReorder")
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_TemplateReorder")
	require.NoError(t, err, "Expected no error")
	time.Sleep(10 * time.Second)

	config := &types.BuildConfiguration{
		ProjectID:  "Empty",
		Name:       "Template Reorder",
		TemplateID: "Tempy",
		Steps: types.BuildSteps{
			types.BuildStep{
				Name: "Muh",
				Type: "simpleRunner",
				Properties: types.Properties{
					"script.content":     "env",
					"teamcity.step.mode": "default",
					"use.custom.script":  "true",
				},
			},
			types.BuildStep{
				ID:   "RUNNER_3",
				Name: "Env",
				Type: "simpleRunner",
				Properties: types.Properties{
					"script.content":     "env",
					"teamcity.step.mode": "default",
					"use.custom.script":  "true",
				},
			},
		},
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")

	var id1 string
	if len(config.Steps) >= 1 {
		id1 = config.Steps[0].ID
	}

	assert.Equal(t, "Empty_TemplateReorder", config.ID, "Supplied ID")
	assert.Equal(t, types.BuildSteps{
		types.BuildStep{
			ID:   id1,
			Name: "Muh",
			Type: "simpleRunner",
			Properties: types.Properties{
				"script.content":     "env",
				"teamcity.step.mode": "default",
				"use.custom.script":  "true",
			},
		},
		types.BuildStep{
			ID:   "RUNNER_3",
			Name: "Env",
			Type: "simpleRunner",
			Properties: types.Properties{
				"script.content":     "env",
				"teamcity.step.mode": "default",
				"use.custom.script":  "true",
			},
		},
	}, config.Steps, "Build steps")
}

func TestClientCreateBuildConfigurationFull(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	err = client.DeleteBuildConfiguration("Empty_Daws")
	require.NoError(t, err, "Expected no error")

	config := &types.BuildConfiguration{
		ID:          "Empty_Daws",
		ProjectID:   "Empty",
		Name:        "Maws",
		Description: "Maws Description",
		Steps: types.BuildSteps{
			types.BuildStep{
				Name: "Muh",
				Type: "simpleRunner",
				Properties: types.Properties{
					"script.content":     "env",
					"teamcity.step.mode": "default",
					"use.custom.script":  "true",
				},
			},
		},
		Features: types.BuildFeatures{
			types.BuildFeature{
				Type: "commit-status-publisher",
				Properties: types.Properties{
					"github_authentication_type": "token",
					"github_host":                "https://api.github.com",
					"publisherId":                "githubStatusPublisher",
				},
			},
		},
		Triggers: types.BuildTriggers{
			types.BuildTrigger{
				Type: "vcsTrigger",
				Properties: types.Properties{
					"groupCheckinsByCommitter":   "true",
					"perCheckinTriggering":       "true",
					"quietPeriodMode":            "DO_NOT_USE",
					"watchChangesInDependencies": "true",
				},
			},
		},

		SnapshotDependencies: types.BuildSnapshotDependencies{
			types.BuildSnapshotDependency{
				Type: "snapshot_dependency",
				Properties: types.Properties{
					"run-build-if-dependency-failed":          "RUN_ADD_PROBLEM",
					"run-build-if-dependency-failed-to-start": "MAKE_FAILED_TO_START",
					"run-build-on-the-same-agent":             "true",
					"take-started-build-with-same-revisions":  "true",
					"take-successful-builds-only":             "true",
				},
				SourceBuildType: types.BuildType{
					ID: "Single_Normal",
				},
			},
		},
		ArtifactDependencies: types.BuildArtifactDependencies{
			types.BuildArtifactDependency{
				Type: "artifact_dependency",
				Properties: types.Properties{
					"cleanDestinationDirectory": "false",
					"pathRules":                 "+:dist/release.tar.gz => dist/",
					"revisionName":              "sameChainOrLastFinished",
					"revisionValue":             "latest.sameChainOrLastFinished",
				},
				SourceBuildType: types.BuildType{
					ID: "Single_Normal",
				},
			},
		},
		AgentRequirements: types.BuildAgentRequirements{
			types.BuildAgentRequirement{
				Type: "more-than",
				Properties: types.Properties{
					"property-name":  "env.CONFIG_FILE",
					"property-value": "1",
				},
			},
		},
	}
	err = client.CreateBuildConfiguration(config)
	require.NoError(t, err, "Expected no error")
	require.NotNil(t, config, "Create to return config")

	var id1 string
	if len(config.Steps) >= 1 {
		id1 = config.Steps[0].ID
	}
	assert.Equal(t, "Empty_Daws", config.ID, "Supplied ID")
	assert.Equal(t, types.BuildSteps{
		types.BuildStep{
			ID:   id1,
			Name: "Muh",
			Type: "simpleRunner",
			Properties: types.Properties{
				"script.content":     "env",
				"teamcity.step.mode": "default",
				"use.custom.script":  "true",
			},
		},
	}, config.Steps, "Build steps")

	assert.Equal(t, types.BuildFeatures{
		types.BuildFeature{
			ID:   config.Features[0].ID,
			Type: "commit-status-publisher",
			Properties: types.Properties{
				"github_authentication_type": "token",
				"github_host":                "https://api.github.com",
				"publisherId":                "githubStatusPublisher",
			},
		},
	}, config.Features, "Build Features")

	assert.Equal(t, types.BuildTriggers{
		types.BuildTrigger{
			ID:   config.Triggers[0].ID,
			Type: "vcsTrigger",
			Properties: types.Properties{
				"groupCheckinsByCommitter":   "true",
				"perCheckinTriggering":       "true",
				"quietPeriodMode":            "DO_NOT_USE",
				"watchChangesInDependencies": "true",
			},
		},
	}, config.Triggers, "Build Triggers")

	assert.Equal(t, types.BuildSnapshotDependencies{
		types.BuildSnapshotDependency{
			ID:   config.SnapshotDependencies[0].ID,
			Type: "snapshot_dependency",
			Properties: types.Properties{
				"run-build-if-dependency-failed":          "RUN_ADD_PROBLEM",
				"run-build-if-dependency-failed-to-start": "MAKE_FAILED_TO_START",
				"run-build-on-the-same-agent":             "true",
				"take-started-build-with-same-revisions":  "true",
				"take-successful-builds-only":             "true",
			},
			SourceBuildType: types.BuildType{
				ID:          "Single_Normal",
				Name:        config.SnapshotDependencies[0].SourceBuildType.Name,
				ProjectName: config.SnapshotDependencies[0].SourceBuildType.ProjectName,
				ProjectID:   config.SnapshotDependencies[0].SourceBuildType.ProjectID,
				Href:        config.SnapshotDependencies[0].SourceBuildType.Href,
			},
		},
	}, config.SnapshotDependencies, "Build Snapshot Dependencies")

	assert.Equal(t, types.BuildArtifactDependencies{
		types.BuildArtifactDependency{
			ID:   config.ArtifactDependencies[0].ID,
			Type: "artifact_dependency",
			Properties: types.Properties{
				"cleanDestinationDirectory": "false",
				"pathRules":                 "+:dist/release.tar.gz => dist/",
				"revisionName":              "sameChainOrLastFinished",
				"revisionValue":             "latest.sameChainOrLastFinished",
			},
			SourceBuildType: types.BuildType{
				ID:          "Single_Normal",
				Name:        config.ArtifactDependencies[0].SourceBuildType.Name,
				ProjectName: config.ArtifactDependencies[0].SourceBuildType.ProjectName,
				ProjectID:   config.ArtifactDependencies[0].SourceBuildType.ProjectID,
				Href:        config.ArtifactDependencies[0].SourceBuildType.Href,
			},
		},
	}, config.ArtifactDependencies, "Build Artifact Dependencies")

	assert.Equal(t, types.BuildAgentRequirements{
		types.BuildAgentRequirement{
			ID:   config.AgentRequirements[0].ID,
			Type: "more-than",
			Properties: types.Properties{
				"property-name":  "env.CONFIG_FILE",
				"property-value": "1",
			},
		},
	}, config.AgentRequirements, "Build Agent Requirements")

}

func TestClientCreateBuildConfigurationUsedID(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	client.retries = 1

	config := &types.BuildConfiguration{
		ID:        "Single_Normal",
		ProjectID: "Single",
		Name:      "Hej Med Dig",
	}

	err = client.CreateBuildConfiguration(config)
	assert.Error(t, err, "Expected error")
}

func TestClientCreateBuildConfigurationUsedName(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	client.retries = 1

	config := &types.BuildConfiguration{
		ProjectID: "Single",
		Name:      "Normal",
	}

	err = client.CreateBuildConfiguration(config)
	assert.Error(t, err, "Expected error")
}

func TestClientCreateBuildConfigurationUsedNameExplicitID(t *testing.T) {
	client, err := NewRealTestClient(t)
	require.NoError(t, err, "Expected no error")
	client.retries = 1

	config := &types.BuildConfiguration{
		ID:        "Single_Dubie",
		ProjectID: "Single",
		Name:      "Normal",
	}

	err = client.CreateBuildConfiguration(config)
	assert.Error(t, err, "Expected error")
}
