# 03 - rotated the template structures

Rather than having a single template
that loops over the models internally,
we can have a template who's context
is a single model from the datamodel.

We are rotating how we connect our datamodel
to the templates and generated code.

### Make the rotation

First, make the rotation in the type templates
by moving repeated code to the partial templates.

- move `models.go` -> `model.go`
- move inner loop content to `partials/{struct,lib,handler}.go`
- use the partials in `model.go`

#### model.go

With repeated templates and code,
we can render the template for each
item in an iterable (list or map).
This can simplify both the template
and the partials the act as the index for.

```go
package pkg

import (
	"net/http"
	"time"
	
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

{{ template "partials/struct.go" . }}

{{ template "partials/lib.go" . }}

{{ template "partials/handler.go" . }}
```

The other files were exact copies of
the sections found in `model(s).go`
before we started the template rotation.


### Add main.go and make it run

We need to add `main.go` and `pkg.go`
and then wire up the model functions
to make working programs.
We do this by having templates we still
loop over types in, but only to call a function.

#### pkg.go

This is an example of some type looping code.
We AutoMigrate per Model, or Go struct in this case.

```go
func InitDB() (err error) {
	// Create db connection
	db, err = gorm.Open(sqlite.Open("{{ .Datamodel.Name }}.db"), &gorm.Config{})
  if err != nil {
		return fmt.Errorf("failed to connect database:\n%s", err)
  }

  // Migrate the schema per Model
	{{ range .Datamodel.Models -}}
  err = db.AutoMigrate(&{{ .Name }}{})
  if err != nil {
		return fmt.Errorf("failed to migrate database:\n%s", err)
  }
  {{ end }}

	return nil
}
```


### Updated `hof gen` command

With the new files in place, we run the following.

Note the repeated template generation syntax `[]`

```sh
hof gen models.cue -w --diff3 -O out \
	-T main.go='cmd/{{ .Datamodel.Name }}/main.go' \
	-T model.go:Datamodel.Models='[]pkg/{{ kebab .Name }}.go' \
	-T pkg.go=pkg/pkg.go \
	-P 'partials/*'
```

### Final directory tree

```text
adhoc-to-module/03/
├── main.go
├── model.go
├── models.cue
├── out
│   ├── cmd
│   │   └── demo
│   │       └── main.go
│   └── pkg
│       ├── pkg.go
│       ├── post.go
│       └── user.go
├── partials
│   ├── handler.go
│   ├── lib.go
│   └── struct.go
├── pkg.go
└── readme.md

5 directories, 12 files
```

