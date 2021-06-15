package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAccDataShow(t *testing.T) {
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
			name:       "local",
			version:    "2.1.0",
			args:       []string{"data", "show", "local_file"},
			exitStatus: 0,
			stdout: `+----------------+--------+----------+----------+----------+-----------+
| ATTRIBUTE      | TYPE   | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+----------------+--------+----------+----------+----------+-----------+
| content        | string | false    | false    | true     | false     |
| content_base64 | string | false    | false    | true     | false     |
| filename       | string | true     | false    | false    | false     |
| id             | string | false    | true     | true     | false     |
+----------------+--------+----------+----------+----------+-----------+

`,
		},
		{
			desc:       "format json",
			name:       "local",
			version:    "2.1.0",
			args:       []string{"data", "show", "-format=json", "local_file"},
			exitStatus: 0,
			stdout: `{
    "attributes": [
        {
            "name": "content",
            "type": "string",
            "required": false,
            "optional": false,
            "computed": true,
            "sensitive": false
        },
        {
            "name": "content_base64",
            "type": "string",
            "required": false,
            "optional": false,
            "computed": true,
            "sensitive": false
        },
        {
            "name": "filename",
            "type": "string",
            "required": true,
            "optional": false,
            "computed": false,
            "sensitive": false
        },
        {
            "name": "id",
            "type": "string",
            "required": false,
            "optional": true,
            "computed": true,
            "sensitive": false
        }
    ],
    "block_types": []
}
`,
		},
		{
			desc:       "resource not found",
			name:       "local",
			version:    "2.1.0",
			args:       []string{"data", "show", "local_foo"},
			exitStatus: 1,
			stdout:     "",
		},
		{
			desc:       "provider not found",
			name:       "local",
			version:    "2.1.0",
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
