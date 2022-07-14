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

### Some LOC stats

<!--
cloc gen/ partials/ templates/ schema/ seed/ seed/ templates/
cloc demo.cue
cloc out/
-->

#### cloc the generator

```text
$ cloc gen/ partials/ templates/ schema/ templates/
      15 text files.
      15 unique files.                              
       0 files ignored.

github.com/AlDanial/cloc v 1.86  T=0.02 s (600.7 files/s, 24586.6 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              11             75             33            345
Cue                              4             24             35            102
-------------------------------------------------------------------------------
SUM:                            15             99             68            447
-------------------------------------------------------------------------------
```

#### cloc the user input

```text
$ cloc demo.cue
       1 text file.
       1 unique file.
       0 files ignored.

github.com/AlDanial/cloc v 1.86  T=0.02 s (41.9 files/s, 3058.4 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Cue                              1             13             15             45
-------------------------------------------------------------------------------
```

#### cloc application code

```text
$ cloc out/ seed/
       8 text files.
       8 unique files.                              
       2 files ignored.

github.com/AlDanial/cloc v 1.86  T=0.02 s (288.1 files/s, 32841.9 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               4            111             47            473
JSON                             1              0              0             46
Markdown                         1              3              0              4
-------------------------------------------------------------------------------
SUM:                             6            114             47            523
-------------------------------------------------------------------------------
```

