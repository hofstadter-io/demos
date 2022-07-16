package demo

import (
	"github.com/hofstadter-io/demos/gen"
	"github.com/hofstadter-io/demos/schema"
)

watch: {
	@flow(watch)
	
	// run code gen with watch enabled there
	gen: {
		@task(os.Exec)
		cmd: ["hof", "gen", "-w"]
		exitcode: _
	}

	// run build and rerun
	watch: {
		@task(os.Watch)
		globs: ["out/*"]
		handler: {
			event?: _
			beg: {
				n:   string @task(gen.Now)
				s:   "\(n)"
			}
			start: {
				@task(os.Stdout)
				dep:  beg.s
				text: "rebuilding demo \(beg.s)\n"
				done: _
			}
			build: {
				@task(os.Exec)
				dep: start.done
				cmd: ["bash", "-c", "cd out && go build -o ../demo ./cmd/demo"]
				exitcode: _
			}
			end: {
				dep: build.exitcode
				n:   string @task(gen.Now)
				s:   "\(n) (\(dep))"
			}
			alert: {
				@task(os.Stdout)
				dep:  end.s
				text: "demo rebuilt \(end.s)\n"
			}
		}
	}
}

// This is example usage of your generator
DemoExample: gen.#Generator & {
	@gen(demo)

	// inputs to the generator
	Name:        "demo"
	Module:      "github.com/hofstadter-io/demos"

	// these could be in separate files
	"Config":    Config
	"Datamodel": Datamodel

	// other settings
	Diff3:  true
	Outdir: "./out/"

	// watch settings
	WatchFull: ["*.cue", "gen/*", "schema/*"]

	// required by examples inside the same module
	// your users do not set or see this field
	PackageName: ""
}

Config: schema.#Config & {
	// What's in a name?
	About: "a demo server"
	Help:  About // lazy dev...
}

// This is the core data model
// whish is augmented and extended
// by combining CUE and hof generators
Datamodel: schema.#Datamodel & {
	// What's in a name?
	Name:   "demo"

	// Models in our datamodel
	Models: {
		// represents a blogger
		User: {
			Fields: {
				Name: Type:  "string"
				Role: Type:  "string"
				Email: {
					Type: "string"
					required: true
				}
			}

			Reln: {
				Post: Type: "HasMany"
				Comment: Type: "HasMany"
			}
		}

		// represents a post
		Post: {
			Fields: {
				Title: Type:   "string"
				Content: Type: "string"
				Draft: Type:   "bool"
			}

			Reln: {
				User: Type: "OwnedBy"
				Comment: Type: "HasMany"
			}
		}

		Comment: {
			Fields: {
				Content: Type: "string"
			}
			Reln: {
				Post: Type: "OwnedBy"
				User: Type: "OwnedBy"
			}
		}
	}
}
