package demo

import (
	"github.com/hofstadter-io/demos/gen"
)

// This is example usage of your generator
DemoExample: gen.#Generator & {
	@gen(demo)

	// inputs to the generator
	Name:        "demo"
	Module:      "github.com/hofstadter-io/demos"
	"Datamodel": Datamodel
	"Config":    Config

	// other settings
	Diff3:  true
	Outdir: "./out/"

	// watch settings
	WatchFull: ["*.cue"]

	// required by examples inside the same module
	// your users do not set or see this field
	PackageName: ""
}
