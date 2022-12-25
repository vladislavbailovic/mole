# mole

File watcher and task runner


## Table of Contents

- [Quick Start](#markdown-header-quick-start)
	- [Building](#markdown-header-building)
	- [Running](#markdown-header-running)
	- [Testing](#markdown-header-testing)


## Quick Start


### Building

```console
$ go build -o ./ mole/cmd/...
```


### Running

```console
$ go run mole/cmd/mole
```


### Testing

```console
$ go test ./...
```


## Usage

```console
$ mole -path './**/*.go' -command 'go test ./...' -target none
```

With no arguments, it will try to load `molerc.json` file which is a file with list of configurations to run. See [the one in this repo](molerc.json) for an example.
