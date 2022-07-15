package demo

import (
	"github.com/hofstadter-io/hof/schema/gen"
)

// This is example usage of your generator
DemoExample: #DemoGenerator & {
	@gen(demo)

	// inputs to the generator
	Name:        "demo"
	Module:      "github.com/hofstadter-io/demos"
	"Datamodel": Datamodel

	// other settings
	Diff3:  true
	Outdir: "./out/"

	// watch settings
	WatchFull: ["*.cue"]

	// required by examples inside the same module
	// your users do not set or see this field
	PackageName: ""
}

// This is your reusable generator module
#DemoGenerator: gen.#Generator & {

	//
	// user input fields
	//

	// most generators will want these 3 fields
	Name:      string
	Module:    string
	Datamodel: #Datamodel

	//
	// Internal Fields
	//

	// this is passed to the templates
	In: {
		// pass in full input
		"Name":      Name
		"Module":    Module
		"Datamodel": Datamodel
	}

	// required for hof CUE modules to work
	// your users do not set or see this field
	PackageName: string | *"github.com/hofstadter-io/demos"

	// The final list of files for hof to generate
	Out: [...gen.#File] & [
		for _, t in _onceFiles {t},
		for _, t in _modelFiles {t},
	]

	// templates rendered once per code gen event
	_onceFiles: [{
		TemplatePath: "main.go"
		Filepath:     "cmd/{{ camel .Name }}/main.go"
	}, {
		TemplatePath: "pkg.go"
		Filepath:     "pkg/pkg.go"
	}]

	// templates rendered per elem, per code gen event
	_modelFiles: [ for _, el in In.Datamodel.Models {
		In: el

		TemplatePath: "model.go"
		Filepath:     "pkg/{{ camel .Name }}.go"
	}]

	// so others can build on this
	...
}
