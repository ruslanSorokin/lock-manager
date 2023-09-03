package test

import (
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
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
