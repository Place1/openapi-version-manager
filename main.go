package main

import (
	"openapi-version-manager/cli"
	"log"
	"os"
)

func main() {
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
