# 02 - scaffold with data + templates

Shows how use several templates
with a single datamodel input
to create the implementation around
the language types and through the call stack.

Add the following files:

- db.go (Gorm)
- lib (helpers/type functions)
- handlers (Echo)
- defaults, common db fields

```sh
hof gen models.cue -w --diff3 -O out \
  -T models.go -T db.go -T libs.go -T handlers.go
```


### Final directory tree

```text
adhoc-to-module/02/
├── db.go
├── handlers.go
├── libs.go
├── models.cue
├── models.go
├── out
│   ├── db.go
│   ├── handlers.go
│   ├── libs.go
│   └── models.go
└── readme.md

1 directory, 10 files
```


### Code Excerpts

Snippets from some of the templates.

#### db.go

```go
package {{ .Datamodel.Name }}

import (
	"gorm.io/gorm"
	"gorm.io/gorm/sqlite"
)

var db gorm.DB

func InitDB() (err error) {
	// Setup the database connection
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema for each Model
	{{ range .Datamodel.Models -}}
	db.AutoMigrate(&{{ .Name }}{})
	{{ end }}

	return nil
}
```


#### handlers.go

```go
// setup routes
func setupRoutes(e *echo.Echo) {
{{ range .Datamodel.Models }}
{{ $ModelName := camelT .Name }}
	e.POST("/{{ kebab .Name }}", handleCreate{{ $ModelName }})
	e.GET("/{{ kebab .Name }}", handleList{{ $ModelName }})
	e.GET("/{{ kebab .Name }}/:id", handleGet{{ $ModelName }})
	e.PUT("/{{ kebab .Name }}/:id", handleUpdate{{ $ModelName }})
	e.DELETE("/{{ kebab .Name }}/:id", handleDelete{{ $ModelName }})
{{ end }}
}

{{ range .Datamodel.Models }}
{{ $ModelName := camelT .Name }}
// {{ $ModelName }} handlers
func handleCreate{{ $ModelName }}(c echo.Context) error {
	in := new({{ .Name }})
	if err = c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ret, err := create{{ $ModelName }}(in)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

// ..
{{ end }}
```


### libs.go

```go
package {{ .Datamodel.Name }}

import (
	"gorm.io/gorm"
)

{{ range .Datamodel.Models }}
{{ $ModelName := camelT .Name }}
// helper functions for {{ $ModelName }}
func create{{ $ModelName }}(in *{{ .Name }}) error {
	res := db.Create(&{{ .Name }}{
		{{ range .Fields }}{{ .Name }}: in.{{ .Name }},
		{{ end }}
	})
	return res.Error
}

func list{{ $ModelName }}() ([]*{{ .Name }}, error) {
	out := make([]*{{ .Name }})
	res := db.Find(&out)
	return out, res.Error
}

// ...
{{ end }}
```

### models.go

We've updated this file for requirements from Gorm.

<!--
cat adhoc-to-module/02/models.go
-->

```go
package {{ .Datamodel.Name }}

import (
	"gorm.io/gorm"
)

{{ range .Datamodel.Models }}
type {{ .Name }} struct {
	// gorm fields
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`

	// defined fields
	{{ range .Fields -}}
	{{ .Name }} {{ .Type }}
	{{ end }}

	// relations
	{{ range .Reln -}}
	{{ if eq .Type "OwnedBy" }}{{ .Name }}ID uint{{end}}
	{{ if eq .Type "OwnedBy" "HasOne" -}}
	{{ .Name }} {{ .Name }}
	{{ end -}}
	{{ if eq .Type "HasMany" "ManyToMany" -}}
	{{ .Name }} []{{ .Name }}
	{{ end -}}
	{{ end }}
}
{{ end }}
```
