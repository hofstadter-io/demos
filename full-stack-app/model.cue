package app

// This is the core data model
// whish is augmented and extended
// by combining CUE and hof generators
Types: #Type & {

	// represents a band
	Band: {
		Fields: {
			name: type: "string"
			genre: type: "string"
		}

		Reln: {
			Album: type: "HasMany"
			Date: type: "HasMany"
		}
	}

	// represents an album
	Album: {
		Fields: {
			title: type: "string"
			year:  type: "string"
		}

		Reln: {
			Band: type: "OwnedBy"
		}
	}

	// represents a tour date
	Date: {
		Fields: {
			location: type: "string"
			eventTime: type: "datetime"
		}

		Reln: {
			Band: type: "OwnedBy"
		}
	}
}
