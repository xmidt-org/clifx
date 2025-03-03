# clifx

[![Build Status](https://github.com/xmidt-org/clifx/actions/workflows/ci.yml/badge.svg)](https://github.com/xmidt-org/clifx/actions/workflows/ci.yml)
[![codecov.io](http://codecov.io/github/xmidt-org/clifx/coverage.svg?branch=main)](http://codecov.io/github/xmidt-org/clifx?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/xmidt-org/clifx)](https://goreportcard.com/report/github.com/xmidt-org/clifx)
[![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/xmidt-org/clifx/blob/main/LICENSE)
[![GitHub Release](https://img.shields.io/github/release/xmidt-org/clifx.svg)](CHANGELOG.md)
[![GoDoc](https://pkg.go.dev/badge/github.com/xmidt-org/clifx)](https://pkg.go.dev/github.com/xmidt-org/clifx)

## Summary

Provides basic bootstrapping for a parsed command line into an [fx.App](https://pkg.go.dev/go.uber.org/fx#App). [Kong](https://pkg.go.dev/github.com/alecthomas/kong) is used to parse a command line into an arbitrary `golang` **struct**.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Installation](#installation)
- [Usage](#usage)
- [Custom Options](#custom-options)
- [Suppressing os.Exit](#suppressing-osexit)
- [Custom Arguments](#custom-arguments)
- [Contributing](#contributing)

## Code of Conduct

This project and everyone participating in it are governed by the [XMiDT Code Of Conduct](https://xmidt.io/code_of_conduct/).
By participating, you agree to this Code.

## Installation

```shell
go get github.com/xmidt-org/clifx@latest
```

## Usage

### Basic Usage

```go
import github.com/xmidt-org/clifx

type MyCLI struct {
  Debug bool
  Files []string
}

func main() {
  app := fx.New(
    clifx.Provide[MyCLI](
      clifx.StandardArguments(),
    ),
  
    // a component of type MyCLI will now
    // be available for dependency injection.
    // For example:
  
    fx.Invoke(
      func(cli MyCLI) error {
        // do things
        return nil
      },
    ),

    // the kong.Context can be used to run the CLI
    fx.Invoke(
      func(kctx *kong.Context, sh fx.Shutdowner) error {
        defer sh.Shutdown() // optional: this ensures the App exits from Run when the CLI is finished
        return kctx.Run() // you could pass dependencies to Run
      },
    )
  )
}
```

### Custom options

You can supply custom any number of `kong` options to `Provide`.

```go
clifx.Provide[MyCLI](
 clifx.StandardArguments(),
 kong.UsageOnError(),
 kong.Description("here is a custom tool"),
)
```

### Suppressing os.Exit

By default, `kong` will invoke `os.Exit(1)` anytime a parse fails. You can suppress this easily by providing a noop **Exit** function.  `clifx` provides a `kong` option for this purpose:

```go
import github.com/xmidt-org/clifx

type MyCLI struct {
  Debug bool
  Files []string
}

func main() {
  app := fx.New(
    clifx.Provide[MyCLI](
      clifx.StandardArguments(),
      clifx.SuppressExit(),
    ),
  
    fx.Invoke(
      func(cli MyCLI) error {
        return nil
      },
    ),
  )

  // since we didn't exit the process, we can test app.Err()
  var pe *kong.ParseError
  if errors.As(app.Err(), &pe) {
    // custom behavior in reaction to a bad command-line
  }
}
```

### Custom arguments

`clifx.StandardArguments` returns the command-line passed to the process. You can supply an arguments you like, which is useful for testing or interactive use:

```go
clifx.Provide[MyCLI](
  clifx.AsArguments("--bind", ":8080", "-v"),
)
```

## Contributing

Refer to [CONTRIBUTING.md](CONTRIBUTING.md).
