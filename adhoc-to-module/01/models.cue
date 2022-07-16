package demo

// Schema for a data model in CUE
#Datamodel: { ... // keep open so user can add own fields
	// metadata
	Name:   string
	Module: string

	// a set of models
	Models: [M=string]: { ...
		// name everything for the user
		Name: M
		
		// model fields
		Fields: [F=string]: { ...
			Name: F
			Type: string
		}

		// relations to other models
		Reln: [R=string]: { ...
			Name: R
			Type: "OwnedBy" | "HasOne" | "HasMany" | "ManyToMany"
		}
	}
}

// This is our demo data model
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
				Email: Type: "string"
				Role: Type:  "string"
				Lang: Type:  "string"
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
