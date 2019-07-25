package common

import (
	"github.com/songtianyi/rrframework/logs"
)

var (
	// MapStringToLevel trans string to logs level enum
	MapStringToLevel = map[string]int{
		"EMERGENCY": logs.LevelEmergency,
		"ALERT":     logs.LevelAlert,
		"CRITICAL":  logs.LevelCritical,
		"ERROR":     logs.LevelError,
		"WARNING":   logs.LevelWarning,
		"NOTICE":    logs.LevelNotice,
		"INFO":      logs.LevelInformational,
		"DEBUG":     logs.LevelDebug,
	}
)
