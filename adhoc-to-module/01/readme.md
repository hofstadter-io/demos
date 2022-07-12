# 01 - start using `hof gen`

data + text/template = _

1. show at cli, write to file
1. watch, update, hot reload
1. NewType func, default role
1. diff3 & custom code

```sh
#0 see the full type spec
cue export types.cue
cue def types.cue

#1 the basics
hof gen types.cue -T types.go
hof gen types.cue -T types.go=out.go
hof gen types.cue -O out -T types.go

#2/3 watch, regen, develop
hof gen types.cue -O out -T types.go -w

#4 diff3 & custom code
hof gen types.cue -O out -T types.go -w --diff3
```
