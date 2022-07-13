# 01 - start using `hof gen`

data + text/template = _

1. show at cli, write to file
1. watch, update, hot reload
1. NewType func, default role
1. diff3 & custom code

```sh
#0 see the full datamodel spec
cue export models.cue
cue def models.cue

#1 the basics
hof gen models.cue -T models.go
hof gen models.cue -T models.go=out.go
hof gen models.cue -O out -T models.go

#2/3 watch, regen, develop
hof gen models.cue -O out -T models.go -w

#4 diff3 & custom code
hof gen models.cue -O out -T models.go -w --diff3
```
