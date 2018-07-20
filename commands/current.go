package commands

import (
	"fmt"
	"os/exec"
)

func Current() error {
	cmd := exec.Command("openapi-generator", "version")
	output, err := cmd.Output()
	if err != nil {
		// TODO: handle when openapi-generator doesn't exist on $PATH
		return err
	}
	fmt.Println(string(output))
	return nil
}
