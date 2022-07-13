package demo

// This is the core data model
// whish is augmented and extended
// by combining CUE and hof generators
Datamodel: #Datamodel & {
	// What's in a name?
	Name:   "demo"
	Module: "github.com/hofstadter-io/demos"

	// Models in our datamodel
	Models: {
		// represents a blogger
		User: {
			Fields: {
				Name: Type:  "string"
				Role: Type:  "string"
				Email: Type: "string"
			}

			Reln: {
				Post: Type: "HasMany"
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
			}
		}
	}
}

// Add schema or augment the data model for easier templating
// Can be used to validate, enrich, and extend user input,
// or otherwise transform the input in interesting ways
#Datamodel: {
	// metadata
	Name:   string
	Module: string

	// the actual datamodel
	Models: {
		[M=string]: {
			// give everything names
			Name: M
			Fields: {
				[F=string]: {Name: F, ...}
			}
			Reln: [R=string]: {Name: R, ...}
			// map reln type to go type
			Reln: [R=string]: {
				// restrict reln types
				Type: "OwnedBy" | "HasOne" | "HasMany" | "ManyToMany"
				// goType from type with faux switch statement
				GoType: [
					if Type == "OwnedBy" {R},
					if Type == "HasOne" {R},
					if Type == "HasMany" {"[]\(R)"},
					if Type == "ManyToMany" {"[]\(R)"},
					"panic, unknown Reln.Type for \(R)",
				][0]    // this is the key to the faux switch
			}
			...
		}
	}
}
