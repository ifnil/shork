package runner

import (
	"context"
	"fmt"

	"github.com/ifnil/shork/internal/groups"
	"github.com/ifnil/shork/internal/host"
	"github.com/ifnil/shork/internal/sshx"
)

type Result struct {
	Output   string
	ExitCode string
	Err      error
}

type Runner struct {
	kr *sshx.Keyring
	hm *host.HostMap
	gm groups.GroupMap
}

func NewRunner(kr *sshx.Keyring, hm *host.HostMap, gm groups.GroupMap) *Runner {
	return &Runner{kr: kr, hm: hm, gm: gm}
}

// RunCmd executes a command on a host
func (r *Runner) RunCmd(ctx context.Context, hostname, cmd string) Result {
	h := r.hm.Get(hostname)
	cfg := sshx.CreateSSHConfig(r.kr, h)

	addr := fmt.Sprintf("%s:%s", h.HostName, h.Port)
	client, err := sshx.NewClient(ctx, addr, cfg)
	if err != nil {
		return Result{Err: err}
	}
	defer client.Close()

	sess, err := client.NewSession()
	if err != nil {
		return Result{Err: err}
	}
	defer sess.Close()

	out, err := sess.CombinedOutput(cmd)
	if err != nil {
		return Result{Output: string(out), Err: err}
	}

	return Result{Output: string(out)}
}

// RunGroupCmd runs a command for every host in a group
func (r *Runner) RunGroupCmd(ctx context.Context, group, cmd string) error {
	g, ok := r.gm[group]
	if !ok {
		return fmt.Errorf("group doesn't exist!")
	}

	ch := make(chan Result, len(g))
	for _, h := range g {
		go func() {
			ch <- r.RunCmd(ctx, h, cmd)
		}()
	}

	for range g {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case res := <-ch:
			if res.Err != nil {
				return res.Err
			}

			fmt.Println(res.Output)
		}
	}

	return nil
}

func (r *Runner) RunAdHoc(ctx context.Context, cmd string, hosts ...string) error {
	return nil
}
