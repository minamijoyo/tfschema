package command

import "testing"

func TestDetectProviderName(t *testing.T) {
	cases := []struct {
		desc string
		name string
		want string
		ok   bool
	}{
		{
			desc: "simple",
			name: "foo_bar",
			want: "foo",
			ok:   true,
		},
		{
			desc: "no underscore",
			name: "foo",
			want: "",
			ok:   false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := detectProviderName(tc.name)
			if tc.ok && err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if !tc.ok && err == nil {
				t.Fatalf("expected an error, but no error: %s", got)
			}

			if got != tc.want {
				t.Errorf("got = %s, want = %s", got, tc.want)
			}
		})
	}
}
