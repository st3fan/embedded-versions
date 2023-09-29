package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func checkForUpdate(owner string, repo string, currentVersion string) (string, error) {
	type GitHubRelease struct {
		TagName string `json:"tag_name"`
	}

	client := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var releases []GitHubRelease

	if err := json.Unmarshal(body, &releases); err != nil {
		return "", err
	}

	if releases[0].TagName[1:] != version {
		return releases[0].TagName[1:], nil
	}

	return "", nil
}
