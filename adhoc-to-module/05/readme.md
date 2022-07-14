# 05 - enhancing the api server

We want to add more features and advanced capabilites
to our applications without needing to code it every time.
This is what we will see in this section.

1. Reorg some code to setup for easy additions
1. Reorg the generator defintion and usage
2. Make the new features and update our demo app



## Reorganize partial templates

### 1. add new type partials for new capabilites

We want to generate a client for our API.
We will also build a CLI tool that uses
the client library to make requests.
(_in a later section, we will do this with JS and a web client_)

Create `client.go` and `command.go` partial template files.
Then add the following lines at the bottome of `type.go` regular template.

```
{{ template "client.go" . }}
{{ template "command.go" . }}
```

### 2. Make `pkg.go` as easy to use as `type.go` just was.

Frist, move the existing type partials into `partials/type/`
and update `type.go` to add the `type/` dir to the template imports.

Then,

1. create another nested subdir `partials/pkg/`
1. Add `cli.go`, `db.go`, `server.go` to the pkg partials dir
1. Move existing code from pkg.go to appropriate partials
1. Update `pkg.go` imports (consolidated) and use the partials like `type.go`

### 3. You dir should now look like this

```
// new files
partials/
  pkg/
	  cli.go
		db.go
		server.go
	type/
		client.go
	  command.go

partials/
	type/
		// moved here from partials root (../)
		handler.go
		lib.go
		struct.go

// existing files
demo.cue
types.go

templates/
  main.go
	pkg.go
	model.go
```

## Reorganize the generator CUE and usage

Schema:

- add `schema` dir, move #Datamodel
- update imports

Generator:

- move defs to `gen/demo.cue`
- add #DemoConfig
- add config to usage

## Making enhancemnts easily, for the last time

Two main ideas to gain from this section

#### 1. Easily add new capabilities

By (a) adding new file to partials
and (b) adding a single line to template

#### 2. For the last time

Every app built using your generator can get enhancements
by updating the versions and regnerating the application.
It's good practive to expose a feature toggle, and the
underlying configuation of the feature to the user
for most reusability and flexibility.

With this idea, and the extra feature configuration,
the input to code gen is more than a data model, it's a dm++


## Enhancements to this demo

- calculated fields
- unique email
- graceful shutdown
- client/{go} for API calls as pkg
- cli tool (cobra)
- seed data

### Calculated fields

- Model.isOwned
- PluralName on {Model, Reln}

in `schema/datamodel.cue`

```cue
#Datamodel: {
	Models: [M=string]: {
		// plural versions of Name
		Name: string
		PluralName: string | *"\(Name)s"
		Reln: [R=string]: {
			Name: string
			PluralName: string | *"\(Name)s"
		}

		// calculated fields
		isOwned: bool | *false
		for _, R in Reln if R.Type == "OwnedBy" {
			isOwnded: true
		}
		...
	}
}
```

We can then use `{{ .PluralName }}` when creating
field names in our output which will hold collections.

```go
type User struct {
	// ...
	Posts []Post
}
```

### Unique Email

There will be code we want to generate that
requires more context than the current Datamodel provides.
Fortunately, it is highly extensible and you can add
any extra schema, config, and defaults you want.

1. add extra metadata / config to Field in Datamodel, as a layer on top of Datamodel
1. set email to be unique
1. update struct.go partial to add go struct tag that Gorm understands

Update Datamodel Schema:

### Graceful shutdown

Add code to `partials/pkg/server.go` based on:
https://echo.labstack.com/cookbook/graceful-shutdown/

replacing this line:

```go
func runServer() error {
	// ...

	// Start server
	return e.Start(port)
}
```

with:

```go
	// Start server
	go func() {
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds. 
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
```

Now all of our servers have graceful shutdown.
We could expose the timeout
through configuration at runtime
or a command line flag.


### Go API Client

Just need to fill in the partial template.
Will test using the CLI we add next...

We use `github.com/imroc/req/v3`

Snippet from `partials/model/client.go`

```go
func {{ $ModelName }}Update(id string, input map[string]any) (*{{ .Name }}, error ) {
	data := new({{.Name}})
	url := host + "/{{ kebab .Name }}/{id}"
	
	client := req.C()
	_, err := client.R().
		SetPathParam("id", id).
		SetBody(input).
		SetResult(&data).
		Put(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
```

### CLI and Commands

- add `"github.com/spf13/cobra"` as import in `type.go` and `pkg.go` templates
- handle create/update data and model id as args

```
./demo user create Name=bob Email=bob@hof.io Role=admin
./demo user update 3 Role=usr
./demo user get 3
```

The same types are used and updated
on both the server / client divide.
They are now unified.

### Apikey

1. Add new field to user
1. Add middleware to Echo
1. Check API key and add calling User to echo.Context
1. Ensure users only modify thier own data
1. Add admin permission to CRUD other users

Want to show how we update across the stack
when adding new features...

### Add some other fields

...but not when modifying the datamodel

### Seed data

See `partials/pkg/db.go` for additions

See `seed/*`

## Regen, Rebuild, Retest

```sh
# regen if you aren't -w'n
hof gen
cd ./out
go build ./cmd/demo/

# seed the database
cue export ../seed/data.cue -o seed.json
./demo seed seed.json

# run the server
./demo serve

# test the go client
./demo get user --email tony@hof.io
```

## directory structue after updates

<!--
tree adhoc-to-module/05/ -I cue.mod
-->

```text
adhoc-to-module/05/
├── cue.mods
├── cue.sums
├── demo.cue
├── gen
│   └── demo.cue
├── out
│   ├── cmd
│   │   └── demo
│   │       └── main.go
│   ├── demo.db
│   ├── go.mod
│   ├── go.sum
│   ├── pkg
│   │   ├── pkg.go
│   │   ├── post.go
│   │   └── user.go
│   └── readme.md
├── partials
│   ├── model
│   │   ├── client.go
│   │   ├── command.go
│   │   ├── handler.go
│   │   ├── lib.go
│   │   └── struct.go
│   └── pkg
│       ├── cli.go
│       ├── db.go
│       └── server.go
├── readme.md
├── schema
│   ├── config.cue
│   └── datamodel.cue
├── seed
│   └── data.cue
├── static
│   └── readme.md
├── templates
│   ├── main.go
│   ├── model.go
│   └── pkg.go
└── types.cue

12 directories, 29 files
```


## Using hof/flow to watch code gen and perform post actions

You can use `hof/flow` to run
`hof gen` and `go build` in parallel watches.
When a change happens it triggers a
and cascading regen, rebuild sequence.
You can run anything post code gen like this.

Adding the following to `demo.cue`
will build a `./demo` binary on change.

Run `hof flow @watch`

```cue
watch: {
	// this is what @watch matches
	@flow(watch)

	// (1) and (2) run in parallel, you could have more
	
	// (1) run code gen with watch enabled there
	gen: {
		@task(os.Exec)
		cmd: ["hof", "gen", "-w"]
		exitcode: _
	}

	// (2) run build and tell user when done
	builder: {
		@task(os.Watch)
		globs: ["out/*"]

		// this is a hof/flow run on output change
		handler: {
			event?: _
			build: {
				@task(os.Exec)
				cmd: ["bash", "-c", "cd out && go build -o ../demo ./cmd/demo"]
				exitcode: _
			}
			now: {
				dep: build.exitcode
				n:   string @task(gen.Now)
				s:   "\(n) (\(dep))"
			}
			alert: {
				@task(os.Stdout)
				dep:  now.s
				text: "demo rebuilt \(now.s)\n"
			}
		}
	}
}
```

