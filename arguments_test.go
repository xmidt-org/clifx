// SPDX-FileCopyrightText: 2024 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
)

type ArgumentsTestSuite struct {
	suite.Suite
}

func (suite *ArgumentsTestSuite) TestGetArguments() {
	suite.Run("Empty", func() {
		suite.Equal(
			Arguments{},
			getArguments(Arguments{}),
		)
	})

	suite.Run("NotEmpty", func() {
		custom := Arguments{
			"-f", "somefile.txt",
		}

		suite.Equal(custom, getArguments(custom))
	})

	suite.Run("Empty", func() {
		suite.Equal(
			Arguments(os.Args[1:]),
			getArguments(nil),
		)
	})
}

func (suite *ArgumentsTestSuite) TestSupplyArguments() {
	suite.Run("Empty", func() {
		var args Arguments
		app := fx.New(
			fx.NopLogger,
			SupplyArguments(),
			fx.Populate(&args),
		)

		suite.NoError(app.Err())
		suite.NotNil(args)
		suite.Empty(args)
	})

	suite.Run("NotEmpty", func() {
		var args Arguments
		app := fx.New(
			fx.NopLogger,
			SupplyArguments(
				"-f", "somefile.yml",
			),
			fx.Populate(&args),
		)

		suite.NoError(app.Err())
		suite.NotNil(args)
		suite.Equal(
			Arguments{"-f", "somefile.yml"},
			args,
		)
	})
}

func TestArguments(t *testing.T) {
	suite.Run(t, new(ArgumentsTestSuite))
}
