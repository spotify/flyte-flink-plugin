# flyte-flink-plugin

[![test](https://github.com/spotify/flyte-flink-plugin/actions/workflows/test.yml/badge.svg)](https://github.com/spotify/flyte-flink-plugin/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/spotify/flyte-flink-plugin)](https://goreportcard.com/report/github.com/spotify/flyte-flink-plugin)
[![GoDoc](https://pkg.go.dev/badge/github.com/spotify/flyte-flink-plugin?status.svg)](https://pkg.go.dev/github.com/spotify/flyte-flink-plugin?tab=doc)
[![Lifecycle](https://img.shields.io/badge/lifecycle-alpha-a0c3d2.svg)](https://img.shields.io/badge/lifecycle-alpha-a0c3d2.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/spotify/flyte-flink-plugin)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/spotify/flyte-flink-plugin)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Flyte Flink k8s plugin.

# Contents

- [About](#About)
- [Dependencies](#Dependencies)
- [Building](#Building)
- [Testing](#Testing)
- [License](#License)

# About

Current development status:

- MVP features are developed
- Missing user documentation
- Project being tested, and collecting feedback
- No guarantees of API stability

To learn more about Flyte refer to:

 - [Flyte homepage](https://flyte.org)
 - [Flyte master repository](https://github.com/lyft/flyte)

# Dependencies

To install the project dependencies, run:

```shell
go mod install
```

# Building

To build ginary, run:

```
make build
```

# Testing

To run tests, run:

```shell
make test
```

To run tests with coverage, run:

```shell
go test -v -covermode=count  ./...
```

# License

This project is released under the [Apache License 2.0](./LICENSE).
