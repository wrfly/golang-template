package main

import (
	"fmt"
	"strings"

	"gopkg.in/urfave/cli.v2"
)

// basic info
var (
	Version string
	Commit  string
	Branch  string
	Builder string
	BuildAt string
)

var versionInfoTemplate = `Version: %s
Branch:  %s
Commit:  %s
Builder: %s
BuildAt: %s`

var versionInfo = fmt.Sprintf(versionInfoTemplate,
	Version, Branch, Commit, Builder,
	strings.Replace(BuildAt, ".", " ", 1),
)

var simpleVersionInfo = fmt.Sprintf("%s @%s v%s",
	Commit, Branch, Version,
)

var author = []*cli.Author{
	&cli.Author{
		Name:  "wrfly",
		Email: "mr.wrfly@gmail.com",
	},
}

var helpTemplate = `NAME:
	{{.Name}} - {{.Usage}}

AUTHOR:
	{{range .Authors}}{{ . }}
	{{end}}
VERSION:
	{{.Version}}

OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}`
