package cli

import (
	"io"
	"log"
	"regexp"

	"github.com/songtianyi/rrframework/logs"
	"golang.org/x/crypto/ssh"
)

// Operator interface
type Operator interface {
	GetTransitions(c, t string) []string
	GetPrompts(m string) []*regexp.Regexp
	GetErrPatterns() []*regexp.Regexp
	GetSSHInitializer() SSHInitializer
	GetLinebreak() string
}

var (
	// OperatorManagerInstance is OperatorManager instance
	OperatorManagerInstance *OperatorManager
)

// SSHInitializer ssh init func generator
type SSHInitializer func(*ssh.Client) (io.Reader, io.WriteCloser, *ssh.Session, error)

func init() {
	OperatorManagerInstance = &OperatorManager{
		operatorMap: make(map[string]Operator, 0),
	}
}

// OperatorManager manager cli operators
type OperatorManager struct {
	operatorMap map[string]Operator // operatorMap mapping vendor.type.version to operator
}

// Get method return Operator instance by string
func (s *OperatorManager) Get(t string) Operator {
	for k, v := range s.operatorMap {
		logs.Debug("[ matching ]", k, t, v)
		if regexp.MustCompile(k).MatchString(t) {
			logs.Debug("[ matched ]", k, t, v)
			return v
		}
	}
	return nil
}

// Register do operator registration
func (s *OperatorManager) Register(pattern string, o Operator) {
	logs.Info("Registering op", pattern, o)
	if _, ok := s.operatorMap[pattern]; ok {
		log.Fatal("pattern", pattern, "registered")
	}
	s.operatorMap[pattern] = o
}
