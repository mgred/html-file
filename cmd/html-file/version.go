package main

import "strings"

const defaultVersion = "dev"

func GetVersion() string {
	version := "$Format:%(describe:tags=true)$"
	if strings.HasPrefix(version, "$Format") {
		return defaultVersion
	}

	if strings.HasPrefix(version, "v") {
		return strings.TrimPrefix(version, "v")
	}

	return defaultVersion
}
