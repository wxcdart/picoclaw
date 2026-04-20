//go:build !windows && ((!darwin && !freebsd) || !purego)

package main

import _ "embed"

//go:embed icon.png
var iconPNG []byte

func getIcon() []byte {
	return iconPNG
}
