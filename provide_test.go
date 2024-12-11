// SPDX-FileCopyrightText: 2024 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"testing"

	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
)

type ProvideTestSuite struct {
	suite.Suite
}

func (suite *ProvideTestSuite) testProvideSuccessUseDefaults() {
	var cli *CommandLine
	var kctx *kong.Context

	app := fx.New(
		fx.NopLogger,
		SupplyArguments(),
		Provide(),
		fx.Populate(
			&cli,
			&kctx,
		),
	)

	suite.NoError(app.Err())
	suite.Equal(
		CommandLine{},
		*cli,
	)
}

func (suite *ProvideTestSuite) testProvideSuccessWithArguments() {
	var cli *CommandLine
	var kctx *kong.Context

	app := fx.New(
		fx.NopLogger,
		SupplyArguments(
			"-f", "filename.yml",
			"--no-health",
			"--pprof",
		),
		Provide(),
		fx.Populate(
			&cli,
			&kctx,
		),
	)

	suite.NoError(app.Err())
	suite.Equal(
		CommandLine{
			ConfigFile: "filename.yml",
			NoHealth:   true,
			Pprof:      true,
		},
		*cli,
	)
}

func (suite *ProvideTestSuite) TestProvide() {
	suite.Run("Success/UseDefaults", suite.testProvideSuccessUseDefaults)
	suite.Run("Success/WithArguments", suite.testProvideSuccessWithArguments)
}

func (suite *ProvideTestSuite) testProvideCustomInvalidType() {
	app := fx.New(
		fx.NopLogger,
		SupplyArguments(),
		ProvideCustom[int](), // must be a struct
	)

	suite.Error(app.Err())
}

func (suite *ProvideTestSuite) TestProvideCustom() {
	suite.Run("InvalidType", suite.testProvideCustomInvalidType)
}

func TestProvide(t *testing.T) {
	suite.Run(t, new(ProvideTestSuite))
}
