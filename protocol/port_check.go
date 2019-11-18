package protocol

import (
	"time"
)

// PortCheckRequest struct
type PortCheckRequest struct {
	IP        string        `json:"ip"`
	Port      string        `json:"port"`
	Proto     string        `json:"proto"`
	Timeout   time.Duration `json:"timeout"`   // req timeout setting
	LogPrefix string        `json:"logPrefix"` // log prefix
	EnablePwd string        `json:"enablePwd"` // enable password for cisco devices
	Session   string        `json:"session"`   // session uuid
}

// PortCheckResponse struct
type PortCheckResponse struct {
	Retcode int
	Message string
}
