module github.com/minamijoyo/tfschema

go 1.15

require (
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/go-plugin v1.3.0
	github.com/hashicorp/hcl2 v0.0.0-20190515223218-4b22149b7cef
	github.com/hashicorp/hil v0.0.0-20190212112733-ab17b08d6590 // indirect
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/terraform v0.14.0-rc1
	github.com/hashicorp/vault v0.10.4 // indirect
	github.com/keybase/go-crypto v0.0.0-20161004153544-93f5b35093ba // indirect
	github.com/mitchellh/cli v1.1.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/mitchellh/panicwrap v1.0.0
	github.com/olekukonko/tablewriter v0.0.0-20180506121414-d4647c9c7a84
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/posener/complete v1.2.1
	github.com/vmihailenco/msgpack v4.0.1+incompatible // indirect
	github.com/zclconf/go-cty v1.7.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
)

replace (
	google.golang.org/grpc v1.31.1 => google.golang.org/grpc v1.27.1
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)
