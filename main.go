package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Hello, I am am Go program (version=%s commit=%s date=%s)\n", version, commit, date)

	if true {
		update, err := checkForUpdate("st3fan", "embedded-versions", version)
		if err != nil {
			log.Printf("Failed to do update check: %s\n", err)
		}

		if update != "" {
			log.Printf("A newer version <%s> is available\n", update)
		}
	}
}
