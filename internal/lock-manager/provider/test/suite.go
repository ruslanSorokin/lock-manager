package test

import (
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
	"github.com/stretchr/testify/suite"
)

type ProviderSuite struct {
	suite.TestingSuite

	Provider provider.LockProviderI
}

func NewProviderSuite(s suite.TestingSuite, p provider.LockProviderI) *ProviderSuite {
	return &ProviderSuite{
		TestingSuite: s,
		Provider:     p,
	}
}
