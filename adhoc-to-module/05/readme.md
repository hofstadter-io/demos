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

- unique email
- graceful shutdown
- client/{go} for API calls as pkg
- cli tool (cobra)
- apikeys
- seed data

### Unique Email

There will be code we want to generate that
requires more context than the current Datamodel provides.
Fortunately, it is highly extensible and you can add
any extra schema, config, and defaults you want.

1. add extra metadata / config to Field in Datamodel, as a layer on top of Datamodel
1. set email to be unique
1. update struct.go partial to add go struct tag that Gorm understands

### Graceful shutdown

Just add some code to one partial

### Go API Client

Just need to fill in the partial template.
Will test using the CLI we add next...


### CLI and Commands

- add `"github.com/spf13/cobra"` as import in `type.go` and `pkg.go` templates


We can now see how the same types are used and updated
on both the server and client.

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

## Regen, Rebuild, Retest

```sh
# regen if you aren't -w'n
hof gen

# run the server
./app serve
./app alive

# test the go client
./app seed data.cue
./app get user --email tony@hof.io
```


