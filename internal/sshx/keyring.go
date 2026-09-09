package sshx

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type entry struct {
	once    sync.Once
	signers []ssh.Signer
	err     error
}

type Keyring struct {
	agent   agent.ExtendedAgent
	mu      sync.Mutex
	entries map[string]*entry
}

func NewKeyring(agc agent.ExtendedAgent) *Keyring {
	return &Keyring{
		agent:   agc,
		entries: make(map[string]*entry),
	}
}

func (kr *Keyring) agentSigners() ([]ssh.Signer, error) {
	if kr.agent == nil {
		return nil, nil
	}

	return kr.agent.Signers()
}

func (kr *Keyring) signersFor(path string) ([]ssh.Signer, error) {
	if path == "" {
		return nil, nil
	}

	kr.mu.Lock()
	e, ok := kr.entries[path]
	if !ok {
		e = &entry{}
		kr.entries[path] = e
	}
	kr.mu.Unlock()

	e.once.Do(func() {
		e.signers, e.err = kr.load(path)
	})
	return e.signers, e.err
}

func (kr *Keyring) load(path string) ([]ssh.Signer, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(data)
	if err == nil {
		return []ssh.Signer{signer}, nil
	}

	if _, needsPass := errors.AsType[*ssh.PassphraseMissingError](err); !needsPass {
		return nil, err
	}

	// TODO: prompt for pass
	// ...
	return nil, fmt.Errorf("%s requires a passphrase: %w", path, err)
}
