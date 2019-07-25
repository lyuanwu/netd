package common

import "time"

const (
	// SSHConn ssh connection
	SSHConn = iota
	// TELNETConn telnet connection
	TELNETConn
)

const (
	// DefaultTimeout default max cli execution time
	DefaultTimeout = 5 * time.Second // seconds
)
