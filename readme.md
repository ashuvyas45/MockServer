# Mock Server Service

## Dependency 
- postgresql 10
- GNU make

## Setup

### Install go

- On OSX run brew install go.
- To install in other OSes follow the [link](https://golang.org/doc/install)
- Add the following in your .zshrc: (where `working_dir` is the directory in which you'll checkout your code)
```bash
export GOPATH=<working_dir>
export PATH="${PATH}:${GOPATH}/bin" 
```

### Dependencies

## Build

1. Clone the repo to $GOPATH/src/source.golabs.io/ops-tech directory.
2. Run `make setup` This install prerequisites for the build like linter.
3. Run `make build` This install the project dependencies.
4. Run `make copy config` This will create **test.yml** and **application.yml** fill the values accordingly make sure the db host is set correctly.
5. Run `make build  && make db.setup` This will run the migration.
6. Run `make test` to execute all tests.
7. Run `make build` This compiles and generate a binary.

## Start

- To start the server run: `out/mock-server start`
