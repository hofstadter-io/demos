# 01 - start using `hof gen`

data + text/template = _

The below shows various ways to invoke
`hof gen` on some data and templates.

1. show at cli, write to file
1. watch, update, hot reload
1. Show new fields / types being added (comments, many-2-many support)
1. diff3 & custom code

```sh
#0 see the full datamodel spec
cue export models.cue
cue def models.cue

#1 the basics
hof gen models.cue -T models.go
hof gen models.cue -T models.go=out.go
hof gen models.cue -O out -T models.go

#2/3 watch, regen, develop
hof gen models.cue -O out -T models.go -w

#4 diff3 & custom code
hof gen models.cue -O out -T models.go -w --diff3
```

### models.go

<!--
cat adhoc-to-module/01/models.go
-->

```go
package {{ .Datamodel.Name }}

{{ range .Datamodel.Models }}
type {{ .Name }} struct {
	// defined fields
	{{ range .Fields -}}
	{{ .Name }} {{ .Type }}
	{{ end }}

	// relations
	{{ range .Reln -}}
	{{ .Name }} {{ .GoType }}
	{{ end }}
}
{{ end }}
// comment at the edge of code gen loop
// help prevent git like merge conflicts
// with custom code in diff3 mode
```

### out/models.go

<!--
cat adhoc-to-module/01/out/models.go
-->

```go
package demo

type Post struct {
	// defined fields
	Content string
	Draft   bool
	Title   string

	// relations
	User User
}

type User struct {
	// defined fields
	Email string
	Name  string
	Role  string

	// relations
	Post []Post
}

// comment at the edge of code gen loop
// help prevent git like merge conflicts
// with custom code in diff3 mode
```

### models.cue

We will use this `Datamodel` throughout the numbered subsections.

<!--
cat adhoc-to-module/01/models.cue
-->

```cue
package demo

// This is the core data model
// whish is augmented and extended
// by combining CUE and hof generators
Datamodel: #Datamodel & {
	// What's in a name?
	Name:   "demo"
	Module: "github.com/hofstadter-io/demos"

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

// Add schema or augment the data model for easier templating
// Can be used to validate, enrich, and extend user input,
// or otherwise transform the input in interesting ways
#Datamodel: {
	// metadata
	Name:   string
	Module: string

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
```
