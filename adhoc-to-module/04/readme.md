# 04

let's make this run


### Convert to generator module

Our flags are starting to get a little verbose.
When you find this happening, you can easily
convert to a _generator module_ by adding the
`--as-module <name>' flag to any `hof gen` call.

```sh
# name should match repo
hof gen ... --as-module github.com/username/demo
hof gen -w -G demo
```

- module name
- templates dir
- seed data
- extras
  - cli tool
  - client/{go}
	- html frontend
- test it out (scripts?)
- html+js

// maybe split frontend into next section

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
./app create user tony admin tony@hof.io
```

### Share it for others to use

All you have to do is push a git tag.
Modules require the `vX.Y.Z` version / tag format.

The next section [full-stack app in 60s](../../full-stack-app/)
will use the module you pushlished.

