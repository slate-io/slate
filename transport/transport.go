package transport

import (
	"net/url"

	"github.com/rightlag/slate/command"
	"golang.org/x/crypto/ssh"
)

type SshTransport struct {
	*ssh.Client
}

func NewSshTransport(uri, password string) (*SshTransport, error) {
	// https://tools.ietf.org/html/draft-ietf-secsh-scp-sftp-ssh-uri-04#section-3
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	config := &ssh.ClientConfig{
		User: u.User.Username(),
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	if err != nil {
		return nil, err
	}
	client, err := ssh.Dial("tcp", u.Host, config)
	if err != nil {
		return nil, err
	}
	return &SshTransport{client}, nil
}

func (s SshTransport) CreateCommand() *command.Command {
	return nil
}
