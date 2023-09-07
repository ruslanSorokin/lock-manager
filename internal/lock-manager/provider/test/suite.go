package providertest

import (
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
)

type PSuite struct {
	suite.TestingSuite

	Provider provider.LockProviderI
}

func NewSuite(s suite.TestingSuite, p provider.LockProviderI) *PSuite {
	return &PSuite{
		TestingSuite: s,
		Provider:     p,
	}
}
