package gen

import (
	hof "github.com/hofstadter-io/hof/schema/gen"
	"github.com/hofstadter-io/demos/schema"
)

// This is your reusable generator module
#Generator: hof.#Generator & {

	//
	// user input fields
	//

	Name:      string
	Module:    string
	Datamodel: schema.#Datamodel
	Config:    schema.#Config

	//
	// Internal Fields
	//

	In: {
		// pass in full input
		"Name":      Name
		"Module":    Module
		"Datamodel": Datamodel
		"Config":    Config
	}

	// required for hof CUE modules to work
	// your users do not set or see this field
	PackageName: string | *"github.com/hofstadter-io/demos"

	// The final list of files for hof to generate
	Out: [...hof.#File] & [
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
		In: {
			Model: el
		}

		TemplatePath: "model.go"
		Filepath:     "pkg/{{ camel .Model.Name }}.go"
	}]

	// so others can build on this
	...
}
