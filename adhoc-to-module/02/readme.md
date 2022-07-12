# 02 - scaffold with data + templates

build out around your types

- db (Gorm)
- lib (helpers/type functions)
- handlers (Echo)
- defaults, common db fields

```sh
hof gen types.cue -w --diff3 -O out \
  -T types.go -T db.go -T libs.go -T handlers.go
```

