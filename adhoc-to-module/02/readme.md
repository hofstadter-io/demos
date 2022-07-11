# 02

build out around your types

- db (Gorm)
- lib (helpers/type functions)
- handlers (Echo)
- defaults, common db fields

```
hof gen types.cue --diff3 -W types.cue -X '*.go' -O out \
  -T types.go -T db.go -T libs.go -T handlers.go
```

