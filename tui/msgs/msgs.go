package msgs

import "github.com/ifnil/shork/internal/host"

type HostInfo struct{ Addr string }

type HostResult struct {
	Host   host.Host
	Result string
}

type SpawnModal struct{}

type ModalResult struct{ OK bool }
