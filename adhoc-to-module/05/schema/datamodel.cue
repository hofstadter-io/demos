package schema

// Add schema or augment the data model for easier templating
// Can be used to validate, enrich, and extend user input,
// or otherwise transform the input in interesting ways
#Datamodel: {
	// metadata
	Name:   string

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