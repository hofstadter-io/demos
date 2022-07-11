# 03
rotate our templates and generated code

- repeated templates (types -> type)
- partials
- more targets [cli,client/{go,js}(http)]

- main / pkg
- wire it up

```sh
hof gen types.cue -w --diff3 -O out \
	-T main.go='cmd/{{ .Name }}/main.go' \
	-T type.go:Datamodel.Models='[]pkg/{{ .name }}.go' \
	-T pkg.go=pkg/pkg.go \
	-P 'partials/*'
```
