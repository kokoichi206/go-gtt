package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "0.0.2"

var (
	revision = "HEAD"
)

func main() {
	app := newApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newApp() *cli.App {
	return &cli.App{
		Name:    "gtt",
		Usage:   "generate terraform template",
		Version: fmt.Sprintf("%s (rev:%s)", version, revision),
	}
}
