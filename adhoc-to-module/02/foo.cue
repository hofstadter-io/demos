
package foo

import (
	"github.com/hofstadter-io/hof/schema/gen"
)

// This is example usage of your generator
FooExample: #FooGenerator & {
	@gen(foo)

	// inputs to the generator
	"Datamodel": Datamodel,
	

	// other settings
	Diff3: true
	
	Outdir: "out"
	
	
	
	

	// required by examples inside the same module
	// your users do not set or see this field
	PackageName: ""
}


// This is your reusable generator module
#FooGenerator: gen.#Generator & {

	//
	// user input fields
	//

	// this is the interface for this generator module
	// typically you enforce schema(s) here
	Datamodel: _
	

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

		"Datamodel": Datamodel
		

		...
	}

	// required for hof CUE modules to work
	// your users do not set or see this field
	PackageName: string | *"hof.io/foo"

	// Templates: [gen.#Templates & {Globs: ["./templates/**/*"], TrimPrefix: "./templates/"}]
	Templates: [ { Globs: [ "models.go", "db.go", "libs.go", "handlers.go",  ] } ]
	
	
	// Partials: [gen.#Templates & {Globs: ["./partials/**/*"], TrimPrefix: "./partials/"}]
	Partials: []
	

	

	// The final list of files for hof to generate
	Out: [...gen.#File] & [
		t_0,
		t_1,
		t_2,
		t_3,
		
	]

	// These are the -T mappings
	t_0: {
		
		
		TemplatePath: "models.go"
		Filepath:     ""
	}
	t_1: {
		
		
		TemplatePath: "db.go"
		Filepath:     ""
	}
	t_2: {
		
		
		TemplatePath: "libs.go"
		Filepath:     ""
	}
	t_3: {
		
		
		TemplatePath: "handlers.go"
		Filepath:     ""
	}
	

	// so your users can build on this
	...
}
