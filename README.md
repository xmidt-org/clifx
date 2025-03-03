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
  - [Basic](#basic)
  - [Lifecycle](#lifecycle)
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

### Basic

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

    // you can bind dependencies to the kong.Context.  if desired,
    // you must do this before calling kong.Context.Run:
    fx.Supply("dependency"),
    fx.Invoke(
      func(kctx *kong.Context, dependency string) {
        kctx.Bind(dependency) // this is now available in CLI methods
      }
    ),

    // the kong.Context can be used to run the CLI.
    // This will cause fx.New to run the command:
    fx.Invoke(
      func(kctx *kong.Context, sh fx.Shutdowner) error {
        defer sh.Shutdown() // optional: this ensures the app exits from app.New
        return kctx.Run() // you could pass dependencies to Run
      },
    )
  )
}
```

### Lifecycle

You can bind a CLI to the `fx.Lifecycle` in the same way as any other component. For example,
it's common to want to run the CLI when `fx.App.Run` is called, then shutdown the app when finished:

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

    fx.Invoke(
      func(kctx *kong.Context, l fx.Lifecycle, sh fx.Shutdowner /* any other dependencies from the enclosing app */) {
        l.Append(fx.Hook{
          OnStart: func(_ context.Context) error {
            // optional:  this just exits from app.Run when the CLI is done.
            // without this, app.Run will not return until explicitly stopped, such as
            // by hitting ctrl+C at a console.
            defer sh.Shutdown()

            // don't forget:  you can pass dependencies from the enclosing app here
            return kctx.Run()
          },
        })
      },
    ),
  )

  // this now causes the CLI to be executed.  Any error that is returned will be
  // from the CLI tool.
  app.Run()
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
