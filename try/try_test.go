package main

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FlagSetBit uint

const (
	FlagSetNone FlagSetBit = 1 << iota
	FlagSetHTTP
	FlagSetOutputField
	FlagSetOutputFormat
)

func (s *TrySuite) TestSyntax() {
	logrus.Infof("name: %q", "hello")
	s.EqualValues(1, FlagSetNone)                             // 1
	s.EqualValues(1<<1, FlagSetHTTP)                          // 2
	s.EqualValues(1<<2, FlagSetOutputField)                   // 4
	s.EqualValues(12, FlagSetOutputField|FlagSetOutputFormat) // 4 + 8
}

func TestTrySuite(t *testing.T) {
	suite.Run(t, &TrySuite{})
}

type TrySuite struct {
	suite.Suite
}

func (s *TrySuite) SetupTest() {
}

// The TearDownSuite method will be run after all tests have been run.
func (s *TrySuite) TearDownSuite() {
}
