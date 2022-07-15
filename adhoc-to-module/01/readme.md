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

### models.cue

We will use this `Datamodel` throughout the numbered subsections.

<!--
cat adhoc-to-module/01/models.cue
-->

```cue
package demo

// Schema for a data model in CUE
#Datamodel: { ... // keep open so user can add own fields
	// metadata
	Name:   string
	Module: string

	// a set of models
	Models: [M=string]: { ...
		// name everything for the user
		Name: M
		
		// model fields
		Fields: [F=string]: { ...
			Name: F
			Type: string
		}

		// relations to other models
		Reln: [R=string]: { ...
			Name: R
			Type: "OwnedBy" | "HasOne" | "HasMany" | "ManyToMany"
		}
	}
}

// This is our demo data model
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
				Email: Type: "string"
				Email: unique: true
				Role: Type:  "string"
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

### out/models.go

The result

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
```

