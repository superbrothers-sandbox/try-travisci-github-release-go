package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"runtime"
)

var (
	gitVersion = "v0.0.0"
	gitCommit  = "$Format:%H$"

	printVersion = false
)

const versionTemplate = `Git version:    {{.GitVersion}}
Git commit:     {{.GitCommit}}
Go version:     {{.GoVersion}}
Compiler:       {{.Compiler}}
Platform:       {{.Platform}}
`

func main() {
	flag.BoolVar(&printVersion, "version", printVersion, "Print version and exit")
	flag.Parse()

	if printVersion {
		tmpl := template.Must(template.New("").Parse(versionTemplate))
		if err := tmpl.Execute(os.Stderr, struct {
			GitVersion string
			GitCommit  string
			GoVersion  string
			Compiler   string
			Platform   string
		}{
			GitVersion: gitVersion,
			GitCommit:  gitCommit,
			GoVersion:  runtime.Version(),
			Compiler:   runtime.Compiler,
			Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		}); err != nil {
			panic(err)
		}

		os.Exit(1)
	}

	arg := "World"
	if flag.Arg(0) != "" {
		arg = flag.Arg(0)
	}

	fmt.Printf("Hello %s\n", arg)
}
