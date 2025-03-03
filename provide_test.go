// SPDX-FileCopyrightText: 2025 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"testing"

	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

type TestCLI struct {
	Debug bool   `short:"d" default:"false"`
	Value string `short:"v" required:"true"`
}

type ProvideTestSuite struct {
	suite.Suite
}

func (suite *ProvideTestSuite) testNewNotAStruct() {
	_, _, err := New[int](AsArguments("does", "not", "matter"))
	suite.Error(err)
}

func (suite *ProvideTestSuite) TestNew() {
	suite.Run("NotAStruct", suite.testNewNotAStruct)
}

func (suite *ProvideTestSuite) testProvideSimpleUsage() {
	var cli TestCLI
	var kctx *kong.Context

	fxtest.New(
		suite.T(),
		fx.NopLogger,
		Provide[TestCLI](
			AsArguments("-v", "foo"),
			SuppressExit(),
		),
		fx.Populate(&cli, &kctx),
	)

	suite.Equal(
		TestCLI{
			Value: "foo",
		},
		cli,
	)

	suite.NotNil(kctx)
}

func (suite *ProvideTestSuite) testProvideSuppressExit() {
	var cli TestCLI
	var kctx *kong.Context

	app := fx.New(
		fx.NopLogger,
		Provide[TestCLI](
			AsArguments("--invalid", "not a valid value"),
			SuppressExit(),
		),
		fx.Populate(&cli, &kctx),
	)

	suite.Error(app.Err())

	var pe *kong.ParseError
	suite.ErrorAs(app.Err(), &pe)
}

func (suite *ProvideTestSuite) TestProvide() {
	suite.Run("SimpleUsage", suite.testProvideSimpleUsage)
	suite.Run("SuppressExit", suite.testProvideSuppressExit)
}

func TestProvide(t *testing.T) {
	suite.Run(t, new(ProvideTestSuite))
}
