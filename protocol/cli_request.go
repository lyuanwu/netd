// NetD makes network device operations easy.
// Copyright (C) 2019  sky-cloud.net
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
