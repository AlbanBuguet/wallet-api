// Package flags holds the feature flags of wallet-api. Every change that a
// customer sees ships behind one, off by default (ADR-0004).
package flags

import (
	"os"
	"strings"
)

// The flags.
const (
	// LiveGateUpdates pushes gate changes to the passes.
	LiveGateUpdates = "wallet.live-gate-updates"
)

// Flags tells which flags are on.
type Flags map[string]bool

// FromEnv reads the flags that WALLET_FLAGS turns on, separated by commas.
func FromEnv() Flags {
	f := Flags{}
	for _, name := range strings.Split(os.Getenv("WALLET_FLAGS"), ",") {
		if name = strings.TrimSpace(name); name != "" {
			f[name] = true
		}
	}
	return f
}

// On reports whether a flag is on.
func (f Flags) On(name string) bool { return f[name] }
