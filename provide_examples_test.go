// SPDX-FileCopyrightText: 2024 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"fmt"

	"github.com/alecthomas/kong"
	"go.uber.org/fx"
)

// ExampleProvide_basic shows how to bootstrap a basic command line using clifx.
func ExampleProvide_basic() {
	var cli *CommandLine
	var kctx *kong.Context // this is completely optional
	fx.New(
		fx.NopLogger,

		// if we don't supply any arguments, clifx will use os.Args[1:]
		fx.Supply(
			Arguments{
				"-f", "/path/to/configfile.yml",
			},
		),

		Provide(
			SuppressExit(), // in case of an error, prevent this example from calling os.Exit
		),
		fx.Populate(
			&cli,
			&kctx,
		),
	)

	fmt.Println(cli.ConfigFile)
	fmt.Println(kctx.Args)

	// Output:
	// /path/to/configfile.yml
	// [-f /path/to/configfile.yml]
}

// ExampleProvide_custom shows how to use a custom command line object.
func ExampleProvide_custom() {
	type CustomCommandLine struct {
		TurnSomethingOn bool `name:"on" optional:"" help:"this is just an example"`
	}

	var cli *CustomCommandLine

	fx.New(
		fx.NopLogger,
		SupplyArguments(
			"--on",
		),
		ProvideCustom[CustomCommandLine](
			SuppressExit(),
		),
		fx.Populate(
			&cli,
		),
	)

	fmt.Println(cli.TurnSomethingOn)

	// Output:
	// true
}
