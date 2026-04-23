package main

import (
	"os"

	"github.com/paladin-devops/waypoint/internal/cli"
)

func main() {
	cli.ExposeDocs = true
	os.Exit(cli.Main([]string{"waypoint", "cli-docs"}))
}
