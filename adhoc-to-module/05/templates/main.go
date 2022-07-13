package main

import "{{ .Module }}/pkg"

func main() {
	{{ if .Config.Cli.enabled }}
	pkg.RunCLI()
	{{ else }}
	pkg.RunServer()
	{{ end }}
}
