package pinger

import "fmt"

// Version information
const (
	VersionMajor = 0
	VersionMinor = 0
	VersionPatch = 3
)

// VersionString returns the version as a string
func VersionString() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}

// VersionInfo returns detailed version information
func VersionInfo() string {
	return fmt.Sprintf("go-pinger v%s", VersionString())
}
