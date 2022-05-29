//go:build tools
// +build tools

// following https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package build

import (
	_ "github.com/google/wire/cmd/wire"
)
