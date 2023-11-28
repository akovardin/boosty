package boosty

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BoostyTestSuite struct {
	suite.Suite
}

func (s *BoostyTestSuite) SetupTest() {
	//
}

func TestBoostyTestSuite(t *testing.T) {
	suite.Run(t, new(BoostyTestSuite))
}
