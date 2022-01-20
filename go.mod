module github.com/minamijoyo/tfschema

go 1.16

require (
	github.com/hashicorp/go-hclog v0.15.0
	github.com/hashicorp/go-plugin v1.4.0
	github.com/hashicorp/hcl/v2 v2.9.1
	github.com/hashicorp/hcl2 v0.0.0-20190515223218-4b22149b7cef
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/terraform v0.15.0
	github.com/mitchellh/cli v1.1.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/panicwrap v1.0.0
	github.com/olekukonko/tablewriter v0.0.0-20180506121414-d4647c9c7a84
	github.com/pkg/browser v0.0.0-20201207095918-0426ae3fba23
	github.com/posener/complete v1.2.1
	github.com/zclconf/go-cty v1.8.1
)

replace (
	google.golang.org/grpc v1.31.1 => google.golang.org/grpc v1.27.1
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)
