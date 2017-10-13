package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jigtools/compose/detect"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
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

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug log output",
		},
	}
	app.Before = runBefore
	app.After = runAfter
	app.Setup()

	app.Commands = []cli.Command{
		detectTypeCommand,
		convertFileCommand,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		log.Errorf("What did you do? %v", err)
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "NO ERROR\n")
}

func runBefore(context *cli.Context) error {
	fmt.Fprintf(os.Stderr, "BEFORE\n")

	return nil
}

func runAfter(context *cli.Context) error {
	fmt.Fprintf(os.Stderr, "AFTER\n")

	return nil
}

var (
	detectTypeCommand = cli.Command{
		Name:  "detect",
		Usage: "detect file type",
		ArgsUsage: `<filename>
		Where <filename> can be any service composition description file type.`,
		Action: func(context *cli.Context) error {
			return nil
		},
	}
	convertFileCommand = cli.Command{
		Name:  "convert",
		Usage: "convert input composition file to type/(s) specified by flag",
		ArgsUsage: `<filename>
		Where <filename> can be any service composition description file type.`,
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "composev1", Usage: ""},
			cli.BoolFlag{Name: "composev2", Usage: ""},
			cli.BoolFlag{Name: "composev3", Usage: ""},
			cli.BoolFlag{Name: "stack", Usage: "Docker Swarm stackfile"},
			cli.BoolFlag{Name: "runc-spec", Usage: "Opencontainers runtime spec"},
			cli.BoolFlag{Name: "rancheros-service", Usage: "RancherOS service file"},
			cli.BoolFlag{Name: "rancheros-cloudconfig", Usage: "RancherOS cloud-config"},
		},
		Action: func(context *cli.Context) error {
			argsRequired := 1
			if context.NArg() != argsRequired {
				cli.ShowCommandHelp(context, context.Command.Name)
				return fmt.Errorf(
					"Wrong number of arguments (need %d, got %d)",
					argsRequired,
					context.NArg())
			}
			filename := context.Args().Get(0)
			filetype, err := detect.Type(filename)
			if err != nil {
				return err
			}
			log.Infof("File type of %s is %s", filename, filetype)
			return err
		},
	}
)
