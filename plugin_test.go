//go:build !linux || (linux && !arm64)

package nexp_plugin

import (
	_ "nexp_plugin"
	"testing"
)

func TestPlugin(t *testing.T) {
	// This test makes sure that executable that imports plugin
	// package can actually run. See issue #28789 for details.
}
