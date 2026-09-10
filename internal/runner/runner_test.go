package runner_test

import (
	"testing"

	"github.com/ifnil/shork/internal/groups"
	"github.com/ifnil/shork/internal/host"
	"github.com/ifnil/shork/internal/runner"
	"github.com/ifnil/shork/internal/sshx"
)

func TestRunner_RunCmd(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		kr *sshx.Keyring
		hm *host.HostMap
		gm groups.GroupMap
		// Named input parameters for target function.
		hostname string
		cmd      string
		want     runner.Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := runner.NewRunner(tt.kr, tt.hm, tt.gm)
			got := r.RunCmd(t.Context(), tt.hostname, tt.cmd)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("RunCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
