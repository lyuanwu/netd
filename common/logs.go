package common

// LogConfig is the config struct for rrframework/logs
type LogConfig struct {
	Adaptor  string `json:"adaptor"`
	Filepath string `json:"filepath"`
	Level    string `json:"level"`
	MaxSize  int    `json:"maxsize"`
}
