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

func init() {
	cli.AppHelpTemplate += "\nEXAMPLE:"
	cli.AppHelpTemplate += "\n    $ find-readme -i skip-me -i skip-me-too $HOME/go"
	cli.AppHelpTemplate += "\n"
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
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "ignore, i",
			Usage: "Directory `name` to ignore.",
		},
	}
	app.Action = action

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type ignoreDirs []string

func (i ignoreDirs) contains(name string) bool {
	for _, e := range i {
		if e == name {
			return true
		}
	}
	return false
}

func action(c *cli.Context) error {

	if !c.Args().Present() {
		_ = cli.ShowAppHelp(c)
		return errors.New("No path specified :(")
	}

	var igs ignoreDirs = c.StringSlice("ignore")
	igs = append(igs, "node_modules")

	for _, target := range c.Args() {
		var fn filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				if igs.contains(info.Name()) {
					return filepath.SkipDir
				}
				return nil
			}

			name := info.Name()
			if name == "readme.md" || name == "README.md" {
				fmt.Println(filepath.Join(path))
				return nil
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
