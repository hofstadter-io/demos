package app

import (
	"github.com/hofstadter-io/demos/adhoc-to-module/04:demo"
)

// This is the core data model
// whish is augmented and extended
// by combining CUE and hof generators
App: demo.#Datamodels & {

	Name: "app"
	Module: "github.com/hofstadter-io/demos/full-stack-app"

	Models: {
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
}
