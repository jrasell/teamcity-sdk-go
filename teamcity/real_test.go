package teamcity

import (
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/jrasell/teamcity-sdk-go/types"
)

var host = flag.String("host", "localhost", "hostname to test against")

func (c *Client) WaitForReady() error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/projects", c.version)
	var projects struct {
		Count    int
		Href     string
		NextHref string
		Project  []types.Project
	}

	err := withRetry(200, func() error {
		err := c.doRequest("GET", path, nil, &projects)
		if err != nil {
			time.Sleep(20 * time.Second)
		}
		return err
	})

	return err
}

func IsRealTestSkip(t *testing.T) {
	if *host == "" {
		t.Skip("Skipping real test because no hostname specified")
	}
}

func (c *Client) SkipOlder(t *testing.T, major uint, minor uint) {
	server, err := c.Server()
	if err != nil || server == nil {
		t.Fatal("Unexpected error getting server info", err)
	}
	if server.VersionMajor < major || (server.VersionMajor == major && server.VersionMinor < minor) {
		t.Skipf("Version to old for test %d.%d < %d.%d", server.VersionMajor, server.VersionMinor,
			major, minor)
	}
}

func (c *Client) VersionParameterValue(t *testing.T, parameter string) string {
	server, err := c.Server()
	if err != nil || server == nil {
		t.Fatal("Unexpected error getting server info", err)
	}
	if server.VersionMajor < 10 {
		fmt.Printf("Found major version %d.%d", server.VersionMajor, server.VersionMinor)
		return ""
	}
	fmt.Printf("Secure major version %d.%d", server.VersionMajor, server.VersionMinor)
	return fmt.Sprintf("%%secure:teamcity.password.%s%%", parameter)
}

func NewRealTestClient(t *testing.T) (*Client, error) {
	IsRealTestSkip(t)
	client := New(fmt.Sprintf("http://%s:8112", *host), "admin", "admin", "latest")
	err := client.WaitForReady()
	if err != nil {
		return nil, err
	}
	return client, nil
}
