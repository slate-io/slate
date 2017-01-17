package transport

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/crypto/ssh"
)

type SshTransport struct {
	*ssh.Client
}

type Command struct {
	args []string
	*SshTransport
}

func NewSshTransport(uri, password string) (*SshTransport, error) {
	// https://tools.ietf.org/html/draft-ietf-secsh-scp-sftp-ssh-uri-04#section-3
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if u.User == nil || u.Host == "" || password == "" {
		return nil, fmt.Errorf("")
	}
	config := &ssh.ClientConfig{
		User: u.User.Username(),
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	client, err := ssh.Dial("tcp", u.Host, config)
	if err != nil {
		return nil, err
	}
	return &SshTransport{client}, nil
}

func (s *SshTransport) CreateCommand(args ...string) *Command {
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
	}
	return &Command{args, s}
}

func (c *Command) String() string {
	return strings.Join(c.args, " ")
}

func (c *Command) execute() (string, error) {
	session, err := c.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	stdout, err := session.Output(c.String())
	return string(stdout), nil
}
