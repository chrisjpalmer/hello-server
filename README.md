# Hello Server

A lovely server that says hello to you powered by gotempl and htmx!

## Prerequisites

1. install `asdf` plugin manager: https://asdf-vm.com/
2. install asdf plugins:

```sh
asdf plugin add golang
asdf plugin add dagger
```

3. install tools

```sh
asdf install
```

## VSCode Plugins

- go: https://marketplace.visualstudio.com/items?itemName=golang.Go
- go-templ: https://marketplace.visualstudio.com/items?itemName=a-h.templ
- htmx-tags: https://marketplace.visualstudio.com/items?itemName=otovo-oss.htmx-tags

## Running

```sh
go run .
```

Browse to localhost:8080 on your computer

## Dagger (build system)

Dagger is a workflow engine that runs in containers. Some functions are implemented in dagger to do common tasks locally and in CI.

1. Run from a container: `dagger call run up`
2. Generate any code: `dagger generate`
3. Build the source: `dagger call build`
4. Run any checks: `dagger checks`



