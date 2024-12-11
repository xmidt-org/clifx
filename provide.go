// SPDX-FileCopyrightText: 2024 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"fmt"
	"reflect"

	"github.com/alecthomas/kong"
	"go.uber.org/fx"
)

// In defines the dependencies for bootstrapping a command line object. These are the
// dependencies for fx options returned from the ProvideXXX methods in this package.
type In struct {
	fx.In

	// Arguments are the parameters passed on the command line. If unset,
	// os.Args[1:] is used. A caller can use fx.Supply(Arguments{...}) to
	// supply a custom command line that is different from the one passed to
	// the program. This is useful for testing and for editing the process
	// command line prior to parsing it.
	//
	// Supplying an empty Arguments slice runs the command with no arguments.
	// This is different than a nil Arguments slice, which tells clifx to
	// use os.Args[1:]
	Arguments Arguments `optional:"true"`

	// KongOptions, if supplied, are used to build the internal kong parser.
	// This options are injected from the fx.App and, if there are duplicates,
	// are overridden by any external kong.Options passed to the ProvideXXX
	// functions.
	KongOptions []kong.Option `optional:"true"`
}

// Out defines the components that the various ProvideXXX functions in this
// package create.
type Out[C any] struct {
	fx.Out

	// CommandLine is the parsed command line.
	CommandLine *C

	// KongContext is the parse context returned by kong.Parse(...).
	// This component can be used to inspect the metadata around the command line,
	// such as which subcommand(s) were used.
	KongContext *kong.Context
}

// ProvideCustom establishes the struct used to hold command line information.
// The components from In are taken as dependencies.
//
// The type C must be a struct of a form usable by the kong parser. This function
// will emit a component of type *C, in addition to the kong context.
//
// The kong options passed to this function will override any that are injected
// via the application. The options used are essentially append(In.KongOptions, external...).
// These options allow a caller to supply external options from outside the fx.App.
func ProvideCustom[C any](external ...kong.Option) fx.Option {
	ctype := reflect.TypeOf((*C)(nil)).Elem()
	if ctype.Kind() != reflect.Struct {
		return fx.Error(
			fmt.Errorf("%s must be a struct (pointer to struct is not supported)", ctype),
		)
	}

	return fx.Provide(
		func(in In) (out Out[C], err error) {
			out.CommandLine = reflect.New(ctype).Interface().(*C)

			var k *kong.Kong
			kopts := make([]kong.Option, 0, len(in.KongOptions)+len(external))
			kopts = append(kopts, in.KongOptions...)
			kopts = append(kopts, external...)
			k, err = kong.New(out.CommandLine, kopts...)

			if err == nil {
				out.KongContext, err = k.Parse(getArguments(in.Arguments))
			}

			return
		},
	)
}

// Provide is the same as using ProvideCustom[clifx.CommandLine]. This function
// provides simple bootstrapping for applications that don't need custom command lines.
func Provide(external ...kong.Option) fx.Option {
	return ProvideCustom[CommandLine](external...)
}
