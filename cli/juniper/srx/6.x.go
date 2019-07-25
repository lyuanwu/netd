package srx

import (
	"fmt"
	"io"
	"regexp"

	"github.com/sky-cloud-tec/netd/cli"
	"golang.org/x/crypto/ssh"
)

func init() {
	// register srx 6.x
	cli.OperatorManagerInstance.Register(`(?i)juniper\.srx\.6\.[0-9]{1,}`, newSixZeroOperator())
}

type sixZeroOperator struct {
	lineBeak    string // \r\n \n
	transitions map[string][]string
	prompts     map[string][]*regexp.Regexp
	errs        []*regexp.Regexp
}

func newSixZeroOperator() cli.Operator {
	sixZeroLoginPrompt := regexp.MustCompile("^[[:alnum:]_]{1,}[.]{0,1}[[:alnum:]_-]{0,}@[[:alnum:]._-]+> $")
	sixZeroConfigPrompt := regexp.MustCompile("^[[:alnum:]_]{1,}[.]{0,1}[[:alnum:]_-]{0,}@[[:alnum:]._-]+# $")
	return &sixZeroOperator{
		// mode transition
		// login -> configure_private
		// login -> configure_exclusive
		// login -> configure
		transitions: map[string][]string{
			"login->configure_private":   {"configure private"},
			"configure_private->login":   {"exit"},
			"login->configure_exclusive": {"configure exclusive"},
			"configure_exclusive->login": {"exit"},
			"login->configure":           {"configure"},
			"configure->login":           {"exit"},
		},
		prompts: map[string][]*regexp.Regexp{
			"login":               {sixZeroLoginPrompt},
			"configure":           {sixZeroConfigPrompt},
			"configure_private":   {sixZeroConfigPrompt},
			"configure_exclusive": {sixZeroConfigPrompt},
		},
		errs: []*regexp.Regexp{
			regexp.MustCompile("^syntax error\\.$"),
			regexp.MustCompile("^unknown command\\.$"),
			regexp.MustCompile("^missing argument\\.$"),
			regexp.MustCompile("\\^$"),
			regexp.MustCompile("^error:"),
		},
		lineBeak: "\n",
	}
}

func (s *sixZeroOperator) GetPrompts(k string) []*regexp.Regexp {
	if v, ok := s.prompts[k]; ok {
		return v
	}
	// return empty slice
	// return make([]*regexp.Regexp, 0)
	return nil
}
func (s *sixZeroOperator) GetTransitions(c, t string) []string {
	k := c + "->" + t
	if v, ok := s.transitions[k]; ok {
		return v
	}
	return nil
}

func (s *sixZeroOperator) GetErrPatterns() []*regexp.Regexp {
	return s.errs
}

func (s *sixZeroOperator) GetLinebreak() string {
	return s.lineBeak
}

func (s *sixZeroOperator) GetSSHInitializer() cli.SSHInitializer {
	return func(c *ssh.Client) (io.Reader, io.WriteCloser, *ssh.Session, error) {
		var err error
		session, err := c.NewSession()
		if err != nil {
			return nil, nil, nil, fmt.Errorf("new ssh session failed, %s", err)
		}
		// get stdout and stdin channel
		r, err := session.StdoutPipe()
		if err != nil {
			session.Close()
			return nil, nil, nil, fmt.Errorf("create stdout pipe failed, %s", err)
		}
		w, err := session.StdinPipe()
		if err != nil {
			session.Close()
			return nil, nil, nil, fmt.Errorf("create stdin pipe failed, %s", err)
		}
		if err := session.Shell(); err != nil {
			session.Close()
			return nil, nil, nil, fmt.Errorf("create shell failed, %s", err)
		}
		return r, w, session, nil
	}
}
