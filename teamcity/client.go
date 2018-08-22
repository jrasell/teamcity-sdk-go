package teamcity

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/jrasell/teamcity-sdk-go/types"
)

// Client to access a TeamCity API
type Client struct {
	HTTPClient *http.Client
	username   string
	password   string
	host       string
	version    string
	retries    int
}

func New(host, username, password string, version string) *Client {
	if version == "" {
		version = "latest"
	}
	return &Client{
		HTTPClient: http.DefaultClient,
		username:   username,
		password:   password,
		host:       host,
		version:    version,
		retries:    8,
	}
}

func (c *Client) Server() (*types.Server, error) {
	var server *types.Server
	err := c.doRequest("GET", fmt.Sprintf("/httpAuth/app/rest/%s/server", c.version), nil, &server)
	return server, err
}

func (c *Client) QueueBuild(buildTypeID string, branchName string, properties types.Properties) (*types.Build, error) {
	jsonQuery := struct {
		BuildTypeID string           `json:"buildTypeId,omitempty"`
		Properties  types.Properties `json:"properties"`
		BranchName  string           `json:"branchName,omitempty"`
	}{}
	jsonQuery.BuildTypeID = buildTypeID
	if branchName != "" {
		jsonQuery.BranchName = fmt.Sprintf("refs/heads/%s", branchName)
	}
	jsonQuery.Properties = properties

	build := &types.Build{}

	err := withRetry(c.retries, func() error {
		return c.doRequest("POST", fmt.Sprintf("/httpAuth/app/rest/%s/buildQueue", c.version), jsonQuery, &build)
	})
	if err != nil {
		return nil, err
	}

	return build, nil
}

func (c *Client) SearchBuild(locator string) ([]*types.Build, error) {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/builds/?locator=%s&fields=count,build(*,tags(tag),triggered(*),properties(property),problemOccurrences(*,problemOccurrence(*)),testOccurrences(*,testOccurrence(*)),changes(*,change(*)))", c.version, locator)

	respStruct := struct {
		Count int
		Build []*types.Build
	}{}
	err := withRetry(c.retries, func() error {
		return c.doRequest("GET", path, nil, &respStruct)
	})
	if err != nil {
		return nil, err
	}

	return respStruct.Build, nil
}

func (c *Client) GetBuild(buildID string) (*types.Build, error) {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/builds/id:%s?fields=*,triggered(*),properties(property),problemOccurrences(*,problemOccurrence(*)),testOccurrences(*,testOccurrence(*)),changes(*,change(*))", c.version, buildID)
	var build *types.Build

	err := withRetry(c.retries, func() error {
		return c.doRequest("GET", path, nil, &build)
	})

	if err != nil {
		return nil, err
	}

	if build == nil {
		return nil, errors.New("build not found")
	}

	return build, nil
}

func (c *Client) GetBuildID(buildTypeID, branchName, buildNumber string) (string, error) {
	type builds struct {
		Count    int
		Href     string
		NextHref string
		Build    []types.Build
	}

	path := fmt.Sprintf("/httpAuth/app/rest/%s/buildTypes/id:%s/builds?locator=branch:%s,number:%s,count:1", c.version, buildTypeID, branchName, buildNumber)

	var build *builds
	err := withRetry(c.retries, func() error {
		return c.doRequest("GET", path, nil, &build)
	})
	if err != nil {
		return "ID not found", err
	}

	if build == nil {
		return "ID not found", errors.New("build not found")
	}

	return fmt.Sprintf("%d", build.Build[0].ID), nil
}

func (c *Client) GetBuildProperties(buildID string) (types.Properties, error) {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/builds/id:%s/resulting-properties", c.version, buildID)

	var response types.Properties

	err := withRetry(c.retries, func() error {
		return c.doRequest("GET", path, nil, &response)
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetChanges(path string) ([]types.Change, error) {
	var changes struct {
		Change []types.Change
	}

	path += ",count:99999"
	err := c.doRequest("GET", path, nil, &changes)
	if err != nil {
		return nil, err
	}

	if changes.Change == nil {
		return nil, errors.New("changes not found")
	}

	return changes.Change, nil
}

func (c *Client) GetProblems(path string, count int64) ([]types.ProblemOccurrence, error) {
	var problems struct {
		Count             int64
		Default           bool
		ProblemOccurrence []types.ProblemOccurrence
	}

	path += fmt.Sprintf(",count:%v&fields=*,problemOccurrence(*,details)", count)
	err := c.doRequest("GET", path, nil, &problems)
	if err != nil {
		return nil, err
	}

	if problems.ProblemOccurrence == nil {
		return nil, errors.New("problemOccurrence list not found")
	}

	return problems.ProblemOccurrence, nil
}

func (c *Client) GetTests(path string, count int64, failingOnly bool, ignoreMuted bool) ([]types.TestOccurrence, error) {
	var tests struct {
		Count          int64
		HREF           string
		TestOccurrence []types.TestOccurrence
	}

	if ignoreMuted {
		path += ",currentlyMuted:false"
	}
	if failingOnly {
		path += ",status:FAILURE"
	}
	path += fmt.Sprintf(",count:%v", count)
	err := c.doRequest("GET", path, nil, &tests)
	if err != nil {
		return nil, err
	}

	return tests.TestOccurrence, nil
}

func (c *Client) CancelBuild(buildID int64, comment string) error {
	body := map[string]interface{}{
		"buildCancelRequest": map[string]interface{}{
			"comment":       comment,
			"readIntoQueue": true,
		},
	}
	return c.doRequest("POST", fmt.Sprintf("/httpAuth/app/rest/id:%d", buildID), body, nil)
}

func (c *Client) GetBuildLog(buildID string) (string, error) {
	cnt, err := c.doNotJSONRequest("GET", fmt.Sprintf("/httpAuth/downloadBuildLog.html?buildId=%s", buildID), "application/json", "", nil)
	buf := bytes.NewBuffer(cnt)
	return buf.String(), err
}

func (c *Client) doRetryRequest(method string, path string, data interface{}, v interface{}) error {
	var err error
	if c.retries > 1 {
		err = withRetry(c.retries, func() error {
			return c.doRequest(method, path, data, v)
		})
	} else {
		err = c.doRequest(method, path, data, v)
	}
	return err
}

func (c *Client) doRequest(method string, path string, data interface{}, v interface{}) error {
	var body io.Reader
	if data != nil {
		jsonReq, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("marshaling data: %s", err)
		}

		log.Printf("Request body %s\n", string(jsonReq))
		body = bytes.NewBuffer(jsonReq)
	}

	jsonCnt, err := c.doNotJSONRequest(method, path, "application/json", "application/json", body)
	if err != nil {
		return err
	}
	if jsonCnt == nil {
		return nil
	}

	if v != nil {
		err = json.Unmarshal(jsonCnt, &v)
		if err != nil {
			return fmt.Errorf("json unmarshal: %s (%q)", err, truncate(string(jsonCnt), 1000))
		}
	}

	return nil
}

func (c *Client) doNotJSONRequest(method string, path string, accept string, mime string, body io.Reader) ([]byte, error) {
	//Perform some validation on host. Allow them to specify http vs https
	//if desired and remove trailing slash if present
	host := c.host
	if strings.HasSuffix(host, "/") {
		host = strings.TrimSuffix(host, "/")
	}
	prefix := "https://"
	if strings.HasPrefix(strings.ToLower(host), "http") {
		prefix = ""
	}
	authURL := fmt.Sprintf("%s%s%s", prefix, host, path)

	log.Printf("%s %s\n", method, authURL)
	fmt.Printf("%s %s\n", method, authURL)

	req, _ := http.NewRequest(method, authURL, body)
	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Accept", accept)

	if body != nil {
		req.Header.Add("Content-Type", mime)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 && resp.Header["Content-Type"][0] == "text/plain" {
		return nil, errors.New(string(respBody))
	}

	return respBody, err
}

func truncate(s string, l int) string {
	if len(s) > l {
		return s[:l]
	}
	return s
}

func withRetry(retries int, f func() error) (err error) {
	for i := 0; i < retries; i++ {
		err = f()
		if err != nil {
			log.Printf("Retry: %v / %v, error: %v\n", i, retries, err)
		} else {
			return
		}
	}
	return
}
