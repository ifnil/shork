package sshx

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/ifnil/shork/internal/host"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func GetAgentClient() (agent.ExtendedAgent, error) {
	socket := os.Getenv("SSH_AUTH_SOCK")
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, err
	}
	return agent.NewClient(conn), nil
}

func CreateSSHConfig(kr *Keyring, h host.Host) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: h.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(kr.agentSigners),
			ssh.PublicKeysCallback(func() (signers []ssh.Signer, err error) {
				return kr.signersFor(h.PrivateKey)
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func NewClient(ctx context.Context, addr string, cfg *ssh.ClientConfig) (*ssh.Client, error) {
	// dial with context
	d := net.Dialer{}
	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(30 * time.Second))
	c, chans, reqs, err := ssh.NewClientConn(conn, addr, cfg)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Time{})

	client := ssh.NewClient(c, chans, reqs)
	return client, nil
}
