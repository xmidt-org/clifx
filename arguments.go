// SPDX-FileCopyrightText: 2024 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"os"

	"go.uber.org/fx"
)

// Arguments represent the raw command-line arguments passed to the program.
// This type exists mainly to allow unambiguous dependency injection.
type Arguments []string

// getArguments returns the command line arguments that should be parsed.
//
// If v is non-nil, including if it is empty, this function returns v as is.
//
// If v is nil, this function returns os.Args[1:].
func getArguments(v Arguments) Arguments {
	if v != nil {
		return v
	}

	return os.Args[1:]
}

// SupplyArguments sets the command line arguments to parse. This overrides what is passed
// to the process. This function always supplies a non-nil Arguments. If no argument values
// are passed to this function, this function emits a non-nil, empty Arguments.
func SupplyArguments(args ...string) fx.Option {
	if len(args) == 0 {
		// covers both nil and empty
		return fx.Supply(Arguments{})
	} else {
		return fx.Supply(Arguments(args))
	}
}
