//go:build !windows
// +build !windows

package netutil

func isPacketTooBig(err error) bool {
	return false
}
