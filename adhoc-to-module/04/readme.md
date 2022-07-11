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
```

