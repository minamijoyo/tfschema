package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAccDataList(t *testing.T) {
	skipUnlessAcceptanceTestEnabled(t)

	cases := []struct {
		desc       string
		name       string
		version    string
		args       []string
		exitStatus int
		stdout     string
	}{
		{
			desc:       "simple",
			name:       "dns",
			version:    "3.2.0",
			args:       []string{"data", "list", "dns"},
			exitStatus: 0,
			stdout: `dns_a_record_set
dns_aaaa_record_set
dns_cname_record_set
dns_mx_record_set
dns_ns_record_set
dns_ptr_record_set
dns_srv_record_set
dns_txt_record_set
`,
		},
		{
			desc:       "not found",
			name:       "dns",
			version:    "3.2.0",
			args:       []string{"data", "list", "foo"},
			exitStatus: 1,
			stdout:     "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			setupTestAcc(t, tc.name, tc.version)

			ui := cli.NewMockUi()
			c := newMockCLI(ui, tc.args)

			exitStatus, err := c.Run()
			if tc.exitStatus == 0 && err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if exitStatus != tc.exitStatus {
				t.Fatalf("unexpected exitStatus. got = %d, want = %d, stderr = %s", exitStatus, tc.exitStatus, ui.ErrorWriter.String())
			}

			if tc.exitStatus == 0 {
				stdout := ui.OutputWriter.String()
				if stdout != tc.stdout {
					t.Errorf("got = %s, want = %s", stdout, tc.stdout)
				}
			}
		})
	}
}
