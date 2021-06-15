package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAccResourceShow(t *testing.T) {
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
			name:       "random",
			version:    "3.1.0",
			args:       []string{"resource", "show", "random_id"},
			exitStatus: 0,
			stdout: `+-------------+-------------+----------+----------+----------+-----------+
| ATTRIBUTE   | TYPE        | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+-------------+-------------+----------+----------+----------+-----------+
| b64_std     | string      | false    | false    | true     | false     |
| b64_url     | string      | false    | false    | true     | false     |
| byte_length | number      | true     | false    | false    | false     |
| dec         | string      | false    | false    | true     | false     |
| hex         | string      | false    | false    | true     | false     |
| id          | string      | false    | false    | true     | false     |
| keepers     | map(string) | false    | true     | false    | false     |
| prefix      | string      | false    | true     | false    | false     |
+-------------+-------------+----------+----------+----------+-----------+

`,
		},
		{
			desc:       "format json",
			name:       "random",
			version:    "3.1.0",
			args:       []string{"resource", "show", "-format=json", "random_id"},
			exitStatus: 0,
			stdout: `{
    "attributes": [
        {
            "name": "b64_std",
            "type": "string",
            "required": false,
            "optional": false,
            "computed": true,
            "sensitive": false
        },
        {
            "name": "b64_url",
            "type": "string",
            "required": false,
            "optional": false,
            "computed": true,
            "sensitive": false
        },
        {
            "name": "byte_length",
            "type": "number",
            "required": true,
            "optional": false,
            "computed": false,
            "sensitive": false
        },
        {
            "name": "dec",
            "type": "string",
            "required": false,
            "optional": false,
            "computed": true,
            "sensitive": false
        },
        {
            "name": "hex",
            "type": "string",
            "required": false,
            "optional": false,
            "computed": true,
            "sensitive": false
        },
        {
            "name": "id",
            "type": "string",
            "required": false,
            "optional": false,
            "computed": true,
            "sensitive": false
        },
        {
            "name": "keepers",
            "type": "map(string)",
            "required": false,
            "optional": true,
            "computed": false,
            "sensitive": false
        },
        {
            "name": "prefix",
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
			desc:       "resource not found",
			name:       "random",
			version:    "3.1.0",
			args:       []string{"resource", "show", "random_foo"},
			exitStatus: 1,
			stdout:     "",
		},
		{
			desc:       "provider not found",
			name:       "random",
			version:    "3.1.0",
			args:       []string{"resource", "show", "foo_bar"},
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
