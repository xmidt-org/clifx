// SPDX-FileCopyrightText: 2025 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import "github.com/alecthomas/kong"

// SuppressExit sets a noop exit function so suppress calling os.Exit
// when the command line parsing fails. This is mostly useful for testing
// and examples, but can be useful when an application wants to handle
// parsing errors in a custom manner.
func SuppressExit() kong.Option {
	return kong.Exit(
		func(int) {},
	)
}
