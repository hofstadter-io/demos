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

The output from conversion can be found under [./orig/demo.cue](./orig/demo.cue).
The final form is the [demo.cue](./demo.cue) found in this directory. 

## code changes

The code in the repo is the final code of this section.
These are the changes to make to get from 03 to here.

- set module name, move out of datamodel, update templates
- move templates to own dir, like partials
- remove a lot of generator config, use defaults
- add static/* without change

### running the app

```sh
# hof deps & code gen
hof mod vendor cue
hof gen

# cd to app code
cd out

# go deps
go mod init github.com/username/demo
go mod tidy

# build app
go build ./cmd/demo

# test it
./demo serve
./demo alive
```

### Final directory tree

```text
adhoc-to-module/04/
├── readme.md
├── demo.cue
├── models.cue
├── orig
│   └── demo.cue
├── partials
│   ├── handler.go
│   ├── lib.go
│   └── struct.go
├── static
│   ├── LICENSE
│   └── readme.md
└── templates
    ├── main.go
    ├── model.go
    └── pkg.go

4 directories, 12 files
```
