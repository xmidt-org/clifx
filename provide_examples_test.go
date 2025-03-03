// SPDX-FileCopyrightText: 2025 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"fmt"

	"github.com/alecthomas/kong"
	"go.uber.org/fx"
)

// ExampleProvide_basic shows how to bootstrap a basic command line using clifx.
func ExampleProvide_basic() {
	type cli struct {
		Address string `short:"a"`
	}

	var c cli
	var kctx *kong.Context

	fx.New(
		fx.NopLogger,
		Provide[cli](
			AsArguments("-a", ":8080"), // can use StandardArguments here to pass the process command-line arguments
			SuppressExit(),             // in case of an error, prevent this example from calling os.Exit
		),
		fx.Populate(
			&c,
			&kctx,
		),
	)

	fmt.Println(c.Address)
	fmt.Println(kctx.Args)

	// Output:
	// :8080
	// [-a :8080]
}

// ExampleProvide_run shows how to execute a CLI as part of app.Run().
func ExampleProvide_run() {
	type cli struct {
		Address string `short:"a"`
	}

	fx.New(
		fx.NopLogger,
		fx.Supply(123), // just to illustrate another component
		Provide[cli](
			AsArguments("-a", ":8080"),
			SuppressExit(),
		),
		fx.Invoke(
			func(kctx *kong.Context, sh fx.Shutdowner, value int) error {
				defer sh.Shutdown() //nolint: errcheck
				fmt.Println(kctx.Args)
				fmt.Println(value)

				// see the Kong documentation for how to implement this:
				return kctx.Run(value)
			},
		),
	)

	// Output:
	// [-a :8080]
	// 123
}
