package host

import (
	"os"
	"sync"

	"github.com/ifnil/shork/util"
	"github.com/kevinburke/ssh_config"
)

type HostMap struct {
	mu sync.Mutex
	m  map[string]*Host
}

func NewHostMap() *HostMap {
	return &HostMap{
		mu: sync.Mutex{},
		m:  make(map[string]*Host),
	}
}

func (hm *HostMap) LoadSSHConfig(cfgFile string) error {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	cfgPath, err := util.ExpandPath(cfgFile)
	if err != nil {
		return err
	}

	// open config file
	f, err := os.Open(cfgPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// decode it
	cfg, err := ssh_config.Decode(f)
	if err != nil {
		return err
	}

	// translate
	for _, host := range cfg.Hosts {
		for _, pat := range host.Patterns {
			name := pat.String()
			if name == "" || name == "*" || name == "#" {
				continue // skip catchall, blank, and comment
			}

			hostname, _ := cfg.Get(name, "HostName")
			if hostname == "" {
				continue
			}

			port, _ := cfg.Get(name, "Port")
			if port == "" {
				port = "22"
			}

			user, _ := cfg.Get(name, "User")
			if user == "" {
				user = "root"
			}

			pkey, _ := cfg.Get(name, "IdentityFile")
			if pkey == "" {
				continue
			}

			pkPath, _ := util.ExpandPath(pkey)
			hm.m[name] = &Host{
				HostName:   hostname,
				User:       user,
				Port:       port,
				PrivateKey: pkPath,
			}
		}
	}

	return nil
}

func (hm *HostMap) Get(name string) Host {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	h, ok := hm.m[name]
	if !ok {
		return Host{}
	}

	return *h
}

func (hm *HostMap) Set(name string, h Host) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	hm.m[name] = &Host{
		User:       h.User,
		HostName:   h.HostName,
		PrivateKey: h.PrivateKey,
		Port:       h.Port,
	}
}

func (hm *HostMap) Hosts() []string {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	var hosts []string
	for k := range hm.m {
		hosts = append(hosts, k)
	}

	return hosts
}
