// SPDX-FileCopyrightText: 2024 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

// CommandLine is the parsed, standard xmidt-org command line. Kong (github.com/alecthomas/kong)
// is used to parse the command line.
//
// To extend this command line and add application-specific arguments, simply nest this struct
// and use kong's "embed" struct tag.
type CommandLine struct {
	// ConfigFile overrides any search path and sets the configuration file the application
	// should use. If unset, the application is responsible for locating its configuration file.
	ConfigFile string `short:"f" optional:"" help:"sets the location of the application's configuration file, overriding the search path"`

	// NoHealth indicates whether the application should start a health monitoring endpoint.
	// If false, the application should start an http.Server to monitor its health.
	NoHealth bool `optional:"" help:"shuts off the health endpoint"`

	// Pprof turns on the application's pprof endpoint.  If true, the application should start
	// a pprof endpoint for debugging.
	Pprof bool `help:"turns on the pprof endpoint"`
}
