package commands

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"gopkg.in/cheggaaa/pb.v1"
)

func Use(version string) error {
	output := outputPath(version)
	url := downloadPath(version)
	err := downloadFile(url, output)
	if err != nil {
		return err
	}
	err = setOpenapiBinFile(output)
	if err != nil {
		return err
	}
	fmt.Printf("using openapi-generator %s\n", version)
	return nil
}

func downloadPath(version string) string {
	return fmt.Sprintf("http://search.maven.org/remotecontent?filepath=org/openapitools/openapi-generator-cli/%s/openapi-generator-cli-%s.jar", version, version)
}

func setOpenapiBinFile(openapiJar string) error {
	filePath := getOpenapiExecutablePath()
	scriptContent := getOpenapiExecutableContent(openapiJar)
	fmt.Printf("updated %s\n", filePath)
	err := ioutil.WriteFile(filePath, []byte(scriptContent), os.ModePerm)
	if err != nil {
		fmt.Printf("failed to update openapi-generator executable\n")
	}
	return err
}

func downloadFile(url string, output string) error {
	_, err := os.Stat(output)
	if err == nil {
		fmt.Printf("using cached openapi-generator-cli.jar at: %s\n", output)
		return nil
	}

	out, err := os.Create(output)
	if err != nil {
		return errors.New("failed to create file on local filesystem")
	}
	defer out.Close()

	fmt.Printf("downloading %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("failed to request openapi-generator-cli.jar from maven")
	}
	defer resp.Body.Close()

	bar := pb.New(int(resp.ContentLength)).SetUnits(pb.U_BYTES)
	bar.Format("[=> ]")
	bar.Start()
	_, err = io.Copy(out, bar.NewProxyReader(resp.Body))
	if err != nil {
		return errors.New("failed to download openapi-generator-cli.jar from maven")
	}
	bar.Finish()
	return nil
}
