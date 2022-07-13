package demo

import (
	"github.com/hofstadter-io/demos/schema"
)

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
