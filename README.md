# pointer

[![tag](https://img.shields.io/github/tag/peczenyj/pointer.svg)](https://github.com/peczenyj/pointer/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://pkg.go.dev/badge/github.com/peczenyj/pointer)](http://pkg.go.dev/github.com/peczenyj/pointer)
[![Go](https://github.com/peczenyj/pointer/actions/workflows/go.yml/badge.svg)](https://github.com/peczenyj/pointer/actions/workflows/go.yml)
[![Lint](https://github.com/peczenyj/pointer/actions/workflows/lint.yml/badge.svg)](https://github.com/peczenyj/pointer/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/peczenyj/pointer/graph/badge.svg?token=9y6f3vGgpr)](https://codecov.io/gh/peczenyj/pointer)
[![Report card](https://goreportcard.com/badge/github.com/peczenyj/pointer)](https://goreportcard.com/report/github.com/peczenyj/pointer)
[![CodeQL](https://github.com/peczenyj/pointer/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/peczenyj/pointer/actions/workflows/github-code-scanning/codeql)
[![Dependency Review](https://github.com/peczenyj/pointer/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/peczenyj/pointer/actions/workflows/dependency-review.yml)
[![License](https://img.shields.io/github/license/peczenyj/pointer)](./LICENSE)
[![Latest release](https://img.shields.io/github/release/peczenyj/pointer.svg)](https://github.com/peczenyj/pointer/releases/latest)
[![GitHub Release Date](https://img.shields.io/github/release-date/peczenyj/pointer.svg)](https://github.com/peczenyj/pointer/releases/latest)
[![Last commit](https://img.shields.io/github/last-commit/peczenyj/pointer.svg)](https://github.com/peczenyj/pointer/commit/HEAD)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/peczenyj/pointer/blob/main/CONTRIBUTING.md#pull-request-process)

yet another pointer helper in golang

## usage

```go
    p := pointer.To(64) // assume type is int

    fmt.Printf("return a pointer to %T with value %v", *p, *p)

    //Outputs: "return a pointer to int with value 64"
```
