// SPDX-FileCopyrightText: 2025 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import "os"

// Arguments represent the raw command-line arguments passed to the program.
// This type exists mainly to allow unambiguous dependency injection.
type Arguments []string

// StandardArguments returns the command-line arguments passed to this process,
// i.e. os.Args[1:].
func StandardArguments() Arguments {
	return Arguments(os.Args[1:])
}

// AsArguments provides some syntactic sugar around converting a slice of
// strings into an Arguments.
func AsArguments(args ...string) Arguments {
	return Arguments(args)
}
