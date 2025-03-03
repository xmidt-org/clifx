// SPDX-FileCopyrightText: 2025 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ArgumentsTestSuite struct {
	suite.Suite
}

func (suite *ArgumentsTestSuite) TestStandardArguments() {
	suite.Equal(
		os.Args[1:],
		[]string(StandardArguments()),
	)
}

func (suite *ArgumentsTestSuite) TestAsArguments() {
	testCases := [][]string{
		nil,
		{},
		{"foo"},
		{"foo", "bar"},
	}

	for i, testCase := range testCases {
		suite.Run(strconv.Itoa(i), func() {
			actual := AsArguments(testCase...)
			suite.Equal([]string(testCase), []string(actual))
		})
	}
}

func TestArguments(t *testing.T) {
	suite.Run(t, new(ArgumentsTestSuite))
}
