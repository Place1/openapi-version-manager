// +build linux darwin netbsd openbsd freebsd

package commands

import (
	"fmt"
	"strings"
	"path/filepath"
	"os"
)

func outputPath(version string) string {
	directory := "/opt/openapi-version-manager"
	os.MkdirAll(directory, os.ModePerm)
	return filepath.Join(directory, fmt.Sprintf("openapi-generator.%v.jar", version))
}

func getOpenapiExecutablePath() string {
	return "/usr/local/bin/openapi-generator"
}

func getOpenapiExecutableContent(openapiJar string) string {
	return fmt.Sprintf(
		trimLines(`
			#!/bin/sh
			set -e
			java -jar "%s" $@;
		`),
		openapiJar,
	)
}

func trimLines(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := make([]string, len(lines))
	for index, line := range lines {
		output[index] = strings.TrimSpace(line)
	}
	return strings.Join(output, "\n")
}
