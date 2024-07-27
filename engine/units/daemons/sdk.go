package daemons

// Daemon represents the daemon unit application
type Daemon interface {
	Start() bool
	Stop() bool
}
