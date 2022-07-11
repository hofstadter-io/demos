# full-stack app in 60s

This demo shows using the generator module
created in [adhoc-to-module](../adhoc-to-module/).
With the code published and tagged,
anyone can now use it to make a
full-stack app with their own data model.


### new app with different types

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
	"github.com/username/demo"
)

// This is using your generator
App: demo.#DemoGenerator & {
	@gen(app)

	// inputs to the generator
	Name: "app"
	Module: "github.com/username/app"

	// add our datamodel
	"Types": Types,

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
