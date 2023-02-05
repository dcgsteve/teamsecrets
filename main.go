package main

import (
	"github.com/alecthomas/kong"
)

type Context struct {
	CurrentProject string
}

// ProjectCmd menu items
type ProjectCmd struct {
	New struct {
		ProjectName string `arg:"" help:"project name"`
	} `cmd:"" help:"Create new project"`
	Rm struct {
		ProjectName string `arg:"" help:"project name"`
	} `cmd:"" help:"Remove existing project"`
	List struct {
	} `cmd:"" help:"List existing project(s)"`
}

type SetCmd struct {
	SecretName string `arg:"" help:"secret name"`
	SecretData string `arg:"" help:"secret information"`
}
type GetCmd struct {
	SecretName string `arg:"" help:"secret name"`
}

// main menu structure
var cli struct {
	Get     GetCmd     `cmd:"" help:"Get a secret"`
	Set     SetCmd     `cmd:"" help:"Set a secret"`
	Project ProjectCmd `cmd:"" help:"Manage projects"`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{CurrentProject: "test"})
	ctx.FatalIfErrorf(err)
}
