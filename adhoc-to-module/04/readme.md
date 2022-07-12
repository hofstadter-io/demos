# 04 - create a reusable module


## Convert to generator module

Our flags are starting to get a little verbose.
When you find this happening, you can easily
convert to a _generator module_ by adding the
`--as-module <name>' flag to any `hof gen` call.

```sh
# name should match repo
hof gen ... --as-module github.com/username/demo
hof gen -w
```

- module name
- templates dir


## let's make this run

### tidy up the code

### running the app

```sh
# hof deps & code gen
hof mod vendor cue
hof gen

# go deps
go mod init github.com/username/demo
go mod tidy

# build app
go build -o app ./out/cmd/app

# test it
./app serve
./app alive
```

