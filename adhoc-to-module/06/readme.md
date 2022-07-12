# 06 - sharing your module

### Share it for others to use

All you have to do is push a git tag.
Modules require the `vX.Y.Z` version / tag format.
`hof mod` has a generic dependency management system based on Go mods.
We use that for CUE modules until `cue` impliments native support.
`hof` code generation modules are CUE modules as well.

In fact, `hof` datamodels are also CUE modules and they,
with the generator modules, can work side-by-side or in concert.
The [working-with-datamodels](../../working-with-datamodels/) section
covers this and more.

It's now time to use the module you just published
to create a [full-stack app in 60s](../../full-stack-app/)!

