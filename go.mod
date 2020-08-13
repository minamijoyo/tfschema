module github.com/minamijoyo/tfschema

go 1.14

require (
	github.com/goreleaser/goreleaser v0.106.0
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/go-plugin v1.3.0
	github.com/hashicorp/hcl2 v0.0.0-20190515223218-4b22149b7cef
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/terraform v0.13.0
	github.com/mitchellh/cli v1.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/panicwrap v1.0.0
	github.com/olekukonko/tablewriter v0.0.0-20180506121414-d4647c9c7a84
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/posener/complete v1.2.1
	github.com/zclconf/go-cty v1.5.1
	golang.org/x/lint v0.0.0-20190409202823-959b441ac422
)

// After updating to Terraform v0.13.0,
// we got an ambiguous import for github.com/Azure/go-autorest.
// It seems to require explicit replace as a workaround.
// https://github.com/hashicorp/terraform/commit/481b03c34ac01af17ca31a604e70b06764c312c8
// https://github.com/hashicorp/terraform/commit/23fb8f6d21ec3829a67d824936d87df8879c801e

replace (
	github.com/Azure/go-autorest => github.com/tombuildsstuff/go-autorest v14.0.1-0.20200416184303-d4e299a3c04a+incompatible
	github.com/Azure/go-autorest v11.1.2+incompatible => github.com/Azure/go-autorest v12.1.0+incompatible
	github.com/Azure/go-autorest/autorest => github.com/tombuildsstuff/go-autorest/autorest v0.10.1-0.20200416184303-d4e299a3c04a
	github.com/Azure/go-autorest/autorest/azure/auth => github.com/tombuildsstuff/go-autorest/autorest/azure/auth v0.4.3-0.20200416184303-d4e299a3c04a
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)
