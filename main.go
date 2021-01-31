package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	ui := &cli.ColoredUi{
		InfoColor:  cli.UiColorGreen,
		ErrorColor: cli.UiColorRed,
		WarnColor:  cli.UiColorYellow,
		Ui: &cli.BasicUi{
			Reader:      bufio.NewReader(os.Stdin),
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
	}

	c := cli.NewCLI("hello-hc-cli", "0.0.1")
	c.Args = os.Args[1:]
	c.AutocompleteInstall = "install-autocomplete"
	c.AutocompleteUninstall = "uninstall-autocomplete"
	c.Commands = map[string]cli.CommandFactory{
		// default
		"": func() (cli.Command, error) {
			return &defaultCommand{ui}, nil
		},
		"bar": func() (cli.Command, error) {
			return &barCommand{ui}, nil
		},
		"foo": func() (cli.Command, error) {
			return &fooCommand{}, nil
		},
		"foo f1": func() (cli.Command, error) {
			return &f1Command{}, nil
		},
		"foo f2": func() (cli.Command, error) {
			return &f2Command{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		logrus.Error(err)
	}

	os.Exit(exitStatus)
}

type defaultCommand struct {
	UI cli.Ui
}

func (c *defaultCommand) Synopsis() string {
	return "default"
}

func (c *defaultCommand) Help() string {
	return `default help`
}

func (c *defaultCommand) Run(args []string) int {
	c.UI.Info(fmt.Sprintf("default run with %+v", args))
	return 0
}

type barCommand struct {
	UI cli.Ui
}

func (c *barCommand) Synopsis() string {
	return "bar"
}

func (c *barCommand) Help() string {
	return `bar help`
}

func (c *barCommand) Run(args []string) int {
	logrus.Infof("bar run with %+v", args)
	name, err := c.UI.Ask("your name:")
	if err != nil {
		c.UI.Error(err.Error())
	}
	c.UI.Warn(fmt.Sprintf("name: %s", name))

	pass, err := c.UI.AskSecret("login password:")
	if err != nil {
		c.UI.Error(err.Error())
	}
	c.UI.Error(fmt.Sprintf("password: %s", pass))
	c.UI.Info("name and pass got")
	return 0
}

type fooCommand struct {
}

func (c *fooCommand) Synopsis() string {
	return "foo"
}

func (c *fooCommand) Help() string {
	return `foo help`
}

func (c *fooCommand) Run(args []string) int {
	logrus.Infof("run foo with %+v", args)
	return cli.RunResultHelp
}

type f1Command struct {
}

func (c *f1Command) Synopsis() string {
	return "foo f1"
}

func (c *f1Command) Help() string {
	return `foo f1 help`
}

func (c *f1Command) Run(args []string) int {
	logrus.Infof("foo f1 run with %+v", args)
	return 0 // cli.RunResultHelp
}

type f2Command struct {
}

func (c *f2Command) Synopsis() string {
	return "foo f2"
}

func (c *f2Command) Help() string {
	return `foo f2 help`
}

func (c *f2Command) Run(args []string) int {
	logrus.Infof("foo f2 run with %+v", args)
	return 0 // cli.RunResultHelp
}
