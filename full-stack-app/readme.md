# full-stack app in 60s

This demo shows using the generator module
created in [adhoc-to-module](../adhoc-to-module/).
With the code published and tagged,
anyone can now use it to make a
full-stack app with their own data model.

---
dev notes...

1. show using section 01
2. skip frontend walkthrough here, just do second import & regen
3. build frontend in later section here, after the reader uses it here without the details

This should show both end-user cases


### new app with different types

Let's make an app called `jamr`
which tracks bands and their
albums and tour dates.

#### setup a working dir

```sh
# working dir
mkdir app && cd app

go mod init github.com/username/app
hof mod init cue github.com/username/app
```

#### cue.mods

Our CUE dependencies, in this case
just your generator module.

```go
module github.com/username/app

cue v0.4.3

require (
	github.com/username/demo v0.0.1
)
```

#### app.cue

This is where we put the generator usage

```cue
package app

import (
	"github.com/username/demo/gen"
)

// This is using your generator
App: gen.#DemoGenerator & {
	@gen(app)

	// inputs to the generator
	Name: "app"
	Module: "github.com/username/app"

	// add our datamodel
	"Datamodel": Datamodel,

	// Set Config inline
	Config: {
		About: "a music band app"
		Help: About
	}

	// other settings
	Outdir: "./out/"
}
```

### the data model

Our new app is a little cooler.
It tracks bands, albums, and tours.

#### model.cue

```cue
package app

// This is the core data model
// whish is augmented and extended
// by combining CUE and hof generators
Datamodel: #Type & {

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
```

### running the app

```sh
# hof deps, code gen
hof mod vendor cue
hof gen

# go deps, build app
go mod tidy
go build -o app ./out/cmd/app

# test it out
./app serve
./app alive
./app create user admin tony tony@hof.io
```
