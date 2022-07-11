package demo

// This is the core data model
// whish is augmented and extended
// by combining CUE and hof generators
Datamodel: #Datamodel & {
	// What's in a name?
	Name: "demo"
	Module: "github.com/hofstadter-io/demos"

	// Models in our datamodel
	Models: {
		// represents a blogger
		User: {
			Fields: {
				name: type:  "string"
				role: type:  "string"
				email: type: "string"
			}

			Reln: {
				Post: type: "HasMany"
			}
		}

		// represents a post
		Post: {
			Fields: {
				title: type: "string"
				content: type: "string"
				draft: type: "bool"
			}

			Reln: {
				User: type: "OwnedBy"
			}
		}
	}
}

// Add schema or augment the data model for easier templating
// Can be used to validate, enrich, and extend user input,
// or otherwise transform the input in interesting ways
#Datamodel: {
	// metadata
	Name: string
	Module: string

	// the actual datamodel
	Models: {
		[M=string]: {
			// give everything names
			name: M
			Fields: {
				[F=string]: { name: F, ... }
			}
			Reln: [R=string]: { name: R, ... }
			// map reln type to go type
			Reln: [R=string]: {
				// restrict reln types
				type: "OwnedBy" | "HasOne" | "HasMany" | "ManyToMany"
				// goType from type with faux switch statement
				goType: [
					if type == "OwnedBy" { R }
					if type == "HasOne"  { R }
					if type == "HasMany"  { "[]\(R)" }
					if type == "ManyToMany"  { "[]\(R)" }
					"panic, unknown Reln.type for \(R)"
				][0] // this is the key to the faux switch
			}
			...
		}
	}
}

