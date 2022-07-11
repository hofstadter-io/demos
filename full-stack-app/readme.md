# full-stack app in 60s

This demo shows using the generator module
created in [adhoc-to-module](../adhoc-to-module/).
With the code published and tagged,
anyone can now use it to make a
full-stack app with their own data model.


### new app with different types

```sh
# working dir
mkdir app && cd app

go mod init github.com/username/app
hof mod init cue github.com/username/app
```

`cue.mods`

```go
module github.com/username/app

cue v0.4.3

require (
	github.com/hofstadter-io/hof v0.6.3
)
```

`app.cue`

```cue
package app

import (
	"github.com/username/demo"
)

// This is using your generator
App: #DemoGenerator & {
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

- albums / artists

`model.cue`

```cue
Types: {...}
```

### running the app

```sh
# hof deps, code gen
hof mod vendor cue
hof gen

# go deps, server run
go mod tidy
go run

# test it out
curl localhost:4242/internal/alive
```
