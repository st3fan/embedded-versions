package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// fetchLatestReleaseVersion returns the tag_name of the release in
// the given repository that is marked as latest.
func fetchLatestReleaseVersion(owner string, repo string) (string, error) {
	type GitHubRelease struct {
		TagName string `json:"tag_name"`
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response code while fetching <%s>: %d", url, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var release GitHubRelease

	if err := json.Unmarshal(body, &release); err != nil {
		return "", err
	}

	return release.TagName, nil
}
