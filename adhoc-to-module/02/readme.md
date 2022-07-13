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

```


#### handlers.go

```go

```


### libs.go

```go

```

### models.go

```go

```

