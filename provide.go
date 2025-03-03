// SPDX-FileCopyrightText: 2025 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"fmt"
	"reflect"

	"github.com/alecthomas/kong"
	"go.uber.org/fx"
)

// New parses the command-line given by args.  The optional set of Kong options
// is used to build the parser. This is the most flexible way to bootstrap Kong
// within an fx application, as it allows a caller to control how all the various
// objects are seen by the fx.App.
//
// The type C must be a struct. Otherwise, this function returns an error.
func New[C any](args Arguments, options ...kong.Option) (cli C, kctx *kong.Context, err error) {
	ctype := reflect.TypeOf((*C)(nil)).Elem()
	if ctype.Kind() != reflect.Struct {
		err = fmt.Errorf("%s must be a struct (pointer to struct is not supported)", ctype)
	}

	var k *kong.Kong
	if err == nil {
		k, err = kong.New(&cli, options...)
	}

	if err == nil {
		kctx, err = k.Parse(args)
	}

	return
}

// NewConstructor returns a closure that uses New to build the configuration object
// and context. This function is a useful alternative to New when a caller
// does not wish to inject the arguments and options.
func NewConstructor[C any](args Arguments, options ...kong.Option) func() (C, *kong.Context, error) {
	return func() (C, *kong.Context, error) {
		return New[C](args, options...)
	}
}

// Provide implements the simplest use case:  Global components for the command-line and the
// kong context are emitted based the supplied arguments and options.
func Provide[C any](args Arguments, options ...kong.Option) fx.Option {
	return fx.Provide(
		NewConstructor[C](args, options...),
	)
}
