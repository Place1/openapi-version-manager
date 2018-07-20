package commands

import (
	"path/filepath"
	"fmt"
	"os"
)

func outputPath(version string) string {
	directory := filepath.Join(os.Getenv("LOCALAPPDATA"), "openapi-version-manager")
	os.MkdirAll(directory, os.ModePerm)
	return filepath.Join(directory, fmt.Sprintf("openapi-generator.%v.jar", version))
}

func getOpenapiExecutablePath() string {
	return filepath.Join("C:/", "Windows", "System32", "openapi-generator.bat")
}

func getOpenapiExecutableContent(openapiJar string) string {
	return fmt.Sprintf("java -jar \"%s\" %%*", openapiJar)
}
