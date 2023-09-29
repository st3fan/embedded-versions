package main

// These variables are setup during the build process by gorelease.

var (
	version = "" // Semver from the tag
	tag     = "" // Git tag
	commit  = "" // Git SHA
	date    = "" // Build date
)
