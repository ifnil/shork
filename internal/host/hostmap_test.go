package host_test

import (
	"testing"

	"github.com/ifnil/shork/internal/host"
)

func TestHostMap_LoadSSHConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cfgFile string
		wantErr bool
	}{
		{name: "LoadSSHConfig()", cfgFile: "~/.ssh/config", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := host.NewHostMap()
			gotErr := hm.LoadSSHConfig(tt.cfgFile)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadSSHConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadSSHConfig() succeeded unexpectedly")
			}
		})
	}
}
