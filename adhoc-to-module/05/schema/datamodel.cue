package schema

// Schema for a data model in CUE
#Datamodel: { ... // keep open so user can add own fields
	// metadata
	Name:   string

	// a set of models
	Models: [M=string]: { ...
		// name everything for the user
		Name: M
		PluralName: string | *"\(Name)s"
		
		// model fields
		Fields: [F=string]: { ...
			Name: F
			PluralName: string | *"\(Name)s"
			Type: string
			unique: bool | *false
		}

		// relations to other models
		Reln: [R=string]: { ...
			Name: R
			PluralName: string | *"\(Name)s"
			Type: "OwnedBy" | "HasOne" | "HasMany" | "ManyToMany"
		}

		// calculated fields
		isOwned: bool | *false
		for _, R in Reln if R.Type == "OwnedBy" {
			isOwnded: true
		}
	}
}
