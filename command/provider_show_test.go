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
			name:       "dns",
			version:    "3.2.0",
			args:       []string{"provider", "show", "dns"},
			exitStatus: 0,
			stdout: `+-----------+------+----------+----------+----------+-----------+
| ATTRIBUTE | TYPE | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+-----------+------+----------+----------+----------+-----------+
+-----------+------+----------+----------+----------+-----------+

block_type: update, nesting: NestingList, min_items: 0, max_items: 1
+---------------+--------+----------+----------+----------+-----------+
| ATTRIBUTE     | TYPE   | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+---------------+--------+----------+----------+----------+-----------+
| key_algorithm | string | false    | true     | false    | false     |
| key_name      | string | false    | true     | false    | false     |
| key_secret    | string | false    | true     | false    | false     |
| port          | number | false    | true     | false    | false     |
| retries       | number | false    | true     | false    | false     |
| server        | string | true     | false    | false    | false     |
| timeout       | string | false    | true     | false    | false     |
| transport     | string | false    | true     | false    | false     |
+---------------+--------+----------+----------+----------+-----------+

block_type: gssapi, nesting: NestingList, min_items: 0, max_items: 1
+-----------+--------+----------+----------+----------+-----------+
| ATTRIBUTE | TYPE   | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+-----------+--------+----------+----------+----------+-----------+
| keytab    | string | false    | true     | false    | false     |
| password  | string | false    | true     | false    | true      |
| realm     | string | true     | false    | false    | false     |
| username  | string | false    | true     | false    | false     |
+-----------+--------+----------+----------+----------+-----------+

`,
		},
		{
			desc:       "format json",
			name:       "dns",
			version:    "3.2.0",
			args:       []string{"provider", "show", "-format=json", "dns"},
			exitStatus: 0,
			stdout: `{
    "attributes": [],
    "block_types": [
        {
            "type_name": "update",
            "attributes": [
                {
                    "name": "key_algorithm",
                    "type": "string",
                    "required": false,
                    "optional": true,
                    "computed": false,
                    "sensitive": false
                },
                {
                    "name": "key_name",
                    "type": "string",
                    "required": false,
                    "optional": true,
                    "computed": false,
                    "sensitive": false
                },
                {
                    "name": "key_secret",
                    "type": "string",
                    "required": false,
                    "optional": true,
                    "computed": false,
                    "sensitive": false
                },
                {
                    "name": "port",
                    "type": "number",
                    "required": false,
                    "optional": true,
                    "computed": false,
                    "sensitive": false
                },
                {
                    "name": "retries",
                    "type": "number",
                    "required": false,
                    "optional": true,
                    "computed": false,
                    "sensitive": false
                },
                {
                    "name": "server",
                    "type": "string",
                    "required": true,
                    "optional": false,
                    "computed": false,
                    "sensitive": false
                },
                {
                    "name": "timeout",
                    "type": "string",
                    "required": false,
                    "optional": true,
                    "computed": false,
                    "sensitive": false
                },
                {
                    "name": "transport",
                    "type": "string",
                    "required": false,
                    "optional": true,
                    "computed": false,
                    "sensitive": false
                }
            ],
            "block_types": [
                {
                    "type_name": "gssapi",
                    "attributes": [
                        {
                            "name": "keytab",
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
                            "name": "realm",
                            "type": "string",
                            "required": true,
                            "optional": false,
                            "computed": false,
                            "sensitive": false
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
                    "block_types": [],
                    "nesting": 3,
                    "min_items": 0,
                    "max_items": 1
                }
            ],
            "nesting": 3,
            "min_items": 0,
            "max_items": 1
        }
    ]
}
`,
		},
		{
			desc:       "not found",
			name:       "dns",
			version:    "3.2.0",
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
