package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

const (
	version    string = "0.1.0"
	repository string = "https://github.com/micheam/find-readme"
)

var ignoreDir = "node_modules"

func init() {
	cli.AppHelpTemplate += "\nISSUES:"
	cli.AppHelpTemplate += "\n    " + repository + "/issues"
	cli.AppHelpTemplate += "\n"
}

func main() {
	app := cli.NewApp()
	app.Name = "find-readme"
	app.Usage = "Find All README.md under path."
	app.Version = version
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Michito Maeda",
			Email: "https://twitter.com/michitomaeda",
		},
	}
	app.ArgsUsage = "path ..."
	app.Action = action

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	if !c.Args().Present() {
		_ = cli.ShowAppHelp(c)
		return errors.New("No path specified :(")
	}

	for _, target := range c.Args() {
		var fn filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				if info.Name() == ignoreDir {
					return filepath.SkipDir
				}
				return nil
			}

			name := info.Name()
			switch name {
			case "readme.md", "README.md", "readme.markdown", "README.markdown":
				fmt.Println(filepath.Join(path))
			}
			return nil
		}
		err := filepath.Walk(target, fn)
		if err != nil {
			return err
		}
	}
	return nil
}
