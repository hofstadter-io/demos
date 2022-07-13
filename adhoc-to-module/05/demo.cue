
package demo

import (
	"github.com/hofstadter-io/hof/schema/gen"
)

// This is example usage of your generator
DemoExample: #DemoGenerator & {
	@gen(demo)

	// inputs to the generator
	"Datamodel": Datamodel,

	// other settings
	Diff3: true	
	Outdir: "./out/"

	// watch settings
	WatchFull: ["types.cue"]

	// required by examples inside the same module
	// your users do not set or see this field
	PackageName: ""
}


// This is your reusable generator module
#DemoGenerator: gen.#Generator & {

	//
	// user input fields
	//

	Datamodel: #Datamodel

	//
	// Internal Fields
	//

	In: {
		// pass in full input
		"Datamodel": Datamodel,

		// lift fields to top-level
		Name: Datamodel.Name,
		Module: Datamodel.Module,
	}

	// required for hof CUE modules to work
	// your users do not set or see this field
	PackageName: string | *"github.com/hofstadter-io/demos"

	// The final list of files for hof to generate
	Out: [...gen.#File] & [
		for _, t in _onceFiles { t },
		for _, t in _typeFiles { t },
	]

	// templates rendered once per code gen event
	_onceFiles: [{
		TemplatePath: "main.go"
		Filepath:     "cmd/{{ .Name }}/main.go"
	}, {
		TemplatePath: "pkg.go"
		Filepath:     "pkg/pkg.go"
	}]

	// templates rendered per elem, per code gen event
	_typeFiles: [ for _,el in In.Datamodel.Models {
		In: el
		
		TemplatePath: "type.go"
		Filepath:     "pkg/{{ .name }}.go"
	}]

	// so others can build on this
	...
}
