
package demo

import (
	"github.com/hofstadter-io/hof/schema/gen"
)

// This is example usage of your generator
DemoExample: #DemoGenerator & {
	@gen(demo)

	// inputs to the generator
	"Name": Name,
	"Module": Module,
	"Types": Types,
	

	// other settings
	Diff3: true	
	Outdir: "./out/"

	// required by examples inside the same module
	// your users do not set or see this field
	PackageName: ""
}


// This is your reusable generator module
#DemoGenerator: gen.#Generator & {

	//
	// user input fields
	//

	// this is the interface for this generator module
	// typically you enforce schema(s) here
	Name: _
	Module: _
	Types: _
	

	//
	// Internal Fields
	//

	// This is the global input data the templates will see
	// You can reshape and transform the user inputs
	// While we put it under internal, you can expose In
	In: {
		// if you want to user your input data
		// add top-level fields from your
		// CUE entrypoints here, adjusting as needed
		// Since you made this a module for others,
		// it won't output until this field is filled

		"Name": Name
		"Module": Module
		"Types": Types
		

		...
	}

	// required for hof CUE modules to work
	// your users do not set or see this field
	PackageName: string | *"github.com/hofstadter-io/demos"

	// Templates: [gen.#Templates & {Globs: ["./templates/**/*"], TrimPrefix: "./templates/"}]
	// Templates: [ { Globs: [ "main.go", "type.go", "pkg.go",  ] } ]
	Templates:  [ { Globs: [ "templates/*",  ] } ]
	
	// Partials: [gen.#Templates & {Globs: ["./partials/**/*"], TrimPrefix: "./partials/"}]
	Partials:  [ { Globs: [ "partials/*",  ] } ]
	

	

	// The final list of files for hof to generate
	Out: [...gen.#File] & [
		t_0,
		for _, t in t_1 { t },
		t_2,
		
	]

	// These are the -T mappings
	t_0: {
		
		
		TemplatePath: "main.go"
		Filepath:     "cmd/{{ .Name }}/main.go"
	}
	t_1: [ for _,el in In.Types {
		In: el
		
		TemplatePath: "type.go"
		Filepath:     "pkg/{{ .name }}.go"
	}]
	t_2: {
		
		
		TemplatePath: "pkg.go"
		Filepath:     "pkg/pkg.go"
	}
	

	// so your users can build on this
	...
}
