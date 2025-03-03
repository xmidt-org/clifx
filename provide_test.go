// SPDX-FileCopyrightText: 2025 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package clifx

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

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

func TestProvide(t *testing.T) {
	suite.Run(t, new(ProvideTestSuite))
}
