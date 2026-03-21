//go:build !cgo || windows

package tools

import (
	"errors"
	"os"
)

func openPTY() (*os.File, *os.File, error) {
	return nil, nil, errors.New("PTY is not supported in this build (requires cgo and non-Windows platform)")
}
