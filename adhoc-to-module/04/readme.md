# 04

let's make this run


### Convert to generator module

Our flags are starting to get a little verbose.
When you find this happening, you can easily
convert to a _generator module_ by adding the
`--as-module <name>' flag to any `hof gen` call.

```
hof gen ... --as-module demo
hof gen -w -G demo
```
- extras
  - cli tool
  - client/{js,go}
	- html frontend
- seed data
- test it out
- html+js

- pipe API response back into a new `hof gen`
