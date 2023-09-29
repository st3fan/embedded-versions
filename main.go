package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Hello, I am am Go program (version=%s commit=%s tag=%s date=%s)\n", version, commit, tag, date)

	// For release builds, check if a newer version is available
	if version != "" {
		newestVersion, err := fetchLatestReleaseVersion("st3fan", "embedded-versions")
		if err != nil {
			log.Printf("Failed to do update check: %s\n", err)
		}

		if newestVersion != version {
			log.Printf("A newer version <%s> is available\n", newestVersion)
		}
	}
}
