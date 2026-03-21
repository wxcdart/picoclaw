//go:build cgo && !windows

package tools

import (
	"os"

	"github.com/creack/pty"
)

func openPTY() (*os.File, *os.File, error) {
	return pty.Open()
}
