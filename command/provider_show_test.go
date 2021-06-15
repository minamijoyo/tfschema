package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAccProviderShow(t *testing.T) {
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
			desc:       "format table",
			name:       "hashicups",
			version:    "0.3.1",
			args:       []string{"provider", "show", "hashicups"},
			exitStatus: 0,
			stdout: `+-----------+--------+----------+----------+----------+-----------+
| ATTRIBUTE | TYPE   | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+-----------+--------+----------+----------+----------+-----------+
| host      | string | false    | true     | false    | false     |
| password  | string | false    | true     | false    | true      |
| username  | string | false    | true     | false    | false     |
+-----------+--------+----------+----------+----------+-----------+

`,
		},
		{
			desc:       "format json",
			name:       "hashicups",
			version:    "0.3.1",
			args:       []string{"provider", "show", "-format=json", "hashicups"},
			exitStatus: 0,
			stdout: `{
    "attributes": [
        {
            "name": "host",
            "type": "string",
            "required": false,
            "optional": true,
            "computed": false,
            "sensitive": false
        },
        {
            "name": "password",
            "type": "string",
            "required": false,
            "optional": true,
            "computed": false,
            "sensitive": true
        },
        {
            "name": "username",
            "type": "string",
            "required": false,
            "optional": true,
            "computed": false,
            "sensitive": false
        }
    ],
    "block_types": []
}
`,
		},
		{
			desc:       "not found",
			name:       "hashicups",
			version:    "0.3.1",
			args:       []string{"provider", "show", "foo"},
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
