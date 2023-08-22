//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/incu6us/goimports-reviser/v3"
	_ "github.com/vektra/mockery/v2"
	_ "mvdan.cc/gofumpt"
	_ "golang.org/x/tools/cmd/goimports"
)
