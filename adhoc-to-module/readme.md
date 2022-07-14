# Adhoc to Module

This section introduces `hof gen` code generation
by starting from simple files and building up to
a generator module for full stack apps.

In the next section, [full-stack app](../full-stack-app/),
we will use the module we create here
to build another app for a different purpose.

### Getting started

You'll want to pick a name and repo so you can have the format

`github.com/<username>/<name>`

#### The walkthrough uses `demo` as the app name

```sh
mkdir demo && cd demo
touch models.{cue,go}
```

- [01 - start using `hof gen`](./01/)
- [02 - scaffold with data + templates](./02/)
- [03 - rotate the template structures](./03/)
- [04 - create a reusable module](./04/)
- [05 - enhancing the api server](./05/)
- [06 - sharing your module](./06/)
- [07 - even more in an encore](./07/)
