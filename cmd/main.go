package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

// version, gitCommit and buildDate will be populated by the build options.
var (
	version          = "v0.0.1-dev.0"
	buildDescription = "Development build"
	gitCommit        = ""
	buildDate        = ""
)

func main() {
	app := cli.NewApp()
	app.Name = "compose"
	app.Usage = "Run or convert any runtime spec anywhere."

	ver := []string{
		version,
		buildDescription,
	}
	if gitCommit != "" {
		ver = append(ver, gitCommit)
	}
	if buildDate != "" {
		ver = append(ver, buildDate)
	}
	app.Version = strings.Join(ver, " - ")

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("hello\n")
}
