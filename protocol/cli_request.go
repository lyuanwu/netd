package protocol

import "time"

// CliRequest structure
type CliRequest struct {
	Vendor    string        `json:"Vendor"`    // device vendor
	Type      string        `json:"Type"`      // device type
	Version   string        `json:"Version"`   // device os version
	Device    string        `json:"Device"`    // device identity, uuid, hostname, etc.
	Mode      string        `json:"Mode"`      // target mode
	Protocol  string        `json:"Protocol"`  // telnet or ssh
	Auth      Auth          `json:"Auth"`      // username and password
	Address   string        `json:"Address"`   // host:port eg. 192.168.1.101:22
	Commands  []string      `json:"Commands"`  // cli commands
	Timeout   time.Duration `json:"Timeout"`   // req timeout setting
	LogPrefix string        `json:"LogPrefix"` // log prefix
}

// Auth struct
type Auth struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}
