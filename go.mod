module github.com/minamijoyo/tfschema

go 1.15

require (
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/go-plugin v1.3.0
	github.com/hashicorp/hcl/v2 v2.7.2
	github.com/hashicorp/hcl2 v0.0.0-20190515223218-4b22149b7cef
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/terraform v0.14.0
	github.com/mitchellh/cli v1.1.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/panicwrap v1.0.0
	github.com/olekukonko/tablewriter v0.0.0-20180506121414-d4647c9c7a84
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/posener/complete v1.2.1
	github.com/zclconf/go-cty v1.7.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
)

replace (
	google.golang.org/grpc v1.31.1 => google.golang.org/grpc v1.27.1
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)
