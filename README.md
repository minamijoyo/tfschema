# tfschema

A schema inspector for Terraform providers.

NOTICE: The master branch may be broken due to development for Terraform v0.12 support.

# Features

- Get resource type definitions dynamically from Terraform providers via go-plugin protocol.
- List available resource types.
- Autocomplete resource types in bash/zsh.
- Open official provider documents quickly by your system web browser.
- Terraform v0.11/v0.12 support

![demo](/images/tfschema-demo.gif)

# Getting Started

```
$ brew install minamijoyo/tfschema/tfschema

$ echo 'provider "aws" {}' > main.tf
$ terraform init
```

```
$ tfschema resource list aws | grep aws_security
aws_security_group
aws_security_group_rule
```

```
$ tfschema resource show aws_security_group
+------------------------+-------------+----------+----------+----------+-----------+
| ATTRIBUTE              | TYPE        | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+------------------------+-------------+----------+----------+----------+-----------+
| description            | String      | false    | true     | false    | false     |
| name                   | String      | false    | true     | true     | false     |
| name_prefix            | String      | false    | true     | false    | false     |
| owner_id               | String      | false    | false    | true     | false     |
| revoke_rules_on_delete | Bool        | false    | true     | false    | false     |
| tags                   | Map(String) | false    | true     | false    | false     |
| vpc_id                 | String      | false    | true     | true     | false     |
+------------------------+-------------+----------+----------+----------+-----------+

block_type: egress, nesting: NestingSet, min_items: 0, max_items: 0
+------------------+--------------+----------+----------+----------+-----------+
| ATTRIBUTE        | TYPE         | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+------------------+--------------+----------+----------+----------+-----------+
| cidr_blocks      | List(String) | false    | true     | false    | false     |
| description      | String       | false    | true     | false    | false     |
| from_port        | Number       | true     | false    | false    | false     |
| ipv6_cidr_blocks | List(String) | false    | true     | false    | false     |
| prefix_list_ids  | List(String) | false    | true     | false    | false     |
| protocol         | String       | true     | false    | false    | false     |
| security_groups  | Set(String)  | false    | true     | false    | false     |
| self             | Bool         | false    | true     | false    | false     |
| to_port          | Number       | true     | false    | false    | false     |
+------------------+--------------+----------+----------+----------+-----------+

block_type: ingress, nesting: NestingSet, min_items: 0, max_items: 0
+------------------+--------------+----------+----------+----------+-----------+
| ATTRIBUTE        | TYPE         | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+------------------+--------------+----------+----------+----------+-----------+
| cidr_blocks      | List(String) | false    | true     | false    | false     |
| description      | String       | false    | true     | false    | false     |
| from_port        | Number       | true     | false    | false    | false     |
| ipv6_cidr_blocks | List(String) | false    | true     | false    | false     |
| protocol         | String       | true     | false    | false    | false     |
| security_groups  | Set(String)  | false    | true     | false    | false     |
| self             | Bool         | false    | true     | false    | false     |
| to_port          | Number       | true     | false    | false    | false     |
+------------------+--------------+----------+----------+----------+-----------+
```

# Install

If you are Mac OSX user:

```
$ brew install minamijoyo/tfschema/tfschema
```

or

If you have Go 1.12+ development environment:

```bash
$ git clone https://github.com/minamijoyo/tfschema
$ cd tfschema
$ export GO111MODULE=on
$ go install
```

or

Download the latest compiled binaries and put it anywhere in your executable path.

https://github.com/minamijoyo/tfschema/releases

# Requirements for Terraform providers

- terraform-provider-aws >= v1.11.0
- terraform-provider-google >= v1.5.0
- terraform-provider-azurerm >= v1.3.0

## Other providers
Your provider may or may not support a required API.
The tfschema depends on the Terraform's GetSchema API, and currently some providers do not work unless you patch the provider.

The tfschema requires the provider's dependency library version to:

- hashicorp/terraform >= v0.10.8
- zclconf/go-cty >= 14e23b14828dd12cc7ae0956813c7e91a196e68f (2018/01/06)

# Rules of finding provider's binary
When `terraform init` command is executed, provider's binary is installed under the auto installed directory ( .terraform/plugins/`<OS>_<ARCH>` ) by default.
The tfschema can use the same provider's binary as terraform uses, so you can run `tfschema` command in the same directory where you run the `terraform` command.

If your code base does not use the supported version of the provider yet, you can also use tfschema with other provider's binary installed in different directories.

The tfschema finds provider's binary under the following directories.
This is almost the same as terraform, but not exactly the same.

1. current directory
2. same directory as `tfschema` executable
3. user vendor directory ( terraform.d/plugins/`<OS>_<ARCH>` )
4. auto installed directory ( .terraform/plugins/`<OS>_<ARCH>` )
5. global plugin directory ( $HOME/.terraform.d/plugins )
6. global plugin directory with os and arch ( $HOME/.terraform.d/plugins/`<OS>_<ARCH>` )
7. gopath ( $GOPATH/bin )

Note: if you are Mac OSX user, `<OS>_<ARCH>` is `darwin_amd64`.

# Autocomplete

To enable autocomplete, execute the following command:

```
$ tfschema -install-autocomplete
```

The above command adds the following line to your ~/.bashrc and ~/.zshrc:

.bashrc

```bash
complete -C </path/to/tfschema> tfschema
```

.zshrc

```bash
autoload -U +X bashcompinit && bashcompinit
complete -o nospace -C </path/to/tfschema> tfschema
```

Check your .bashrc and/or .zshrc and reload it.

# Usage

```
$ tfschema --help
Usage: tfschema [--version] [--help] <command> [<args>]

Available commands are:
    data
    provider
    resource
```

```
$ tfschema resource --help
This command is accessed by using one of the subcommands below.

Subcommands:
    browse    Browse a documentation of resource
    list      List resource types
    show      Show a type definition of resource
```

```
$ tfschema resource show --help
Usage: tfschema resource show [options] RESOURCE_TYPE

Options:

  -format=type    Set output format to table or json (default: table)
```

# Known Issues
## Decoding cty.Type
If you got errors like the following, this means your provider does not support GetSchema API correctly, you need to update the `go-cty` in the provider's dependency.

```
$ tfschema resource show aws_security_group
Failed to get schema from provider: reading body error decoding cty.Type: gob: name not registered for interface: "github.com/terraform-providers/terraform-provider-aws/vendor/github.com/zclconf/go-cty/cty.primitiveType"
```

Note that this error log is output from the old `aws` provider, but the latest `aws` provider does not need a patch already.

For example, to update your provider's go-cty version, execute the following command:

```
$ go get -u github.com/terraform-providers/terraform-provider-<NAME>
$ go get -u github.com/kardianos/govendor
$ govendor fetch github.com/zclconf/go-cty/...
$ go install
```

This step will be unnecessary in the future if your provider's dependency is updated officially.

## Invalid schema

If you got errors like the following, this means some resource types have invalid schema information.

```
$ tfschema resource show aws_glue_catalog_database
Failed to get schema from provider: unexpected EOF
The child panicked:

panic: invalid Schema.Elem 4; need *Schema or *Resource
```

There is a conflict between the Terraform core error checking and the provider implementation.

See: https://github.com/hashicorp/terraform/pull/17097

Currently the following resources have this problem for major providers.

### terraform-provider-aws
- data/aws_ecs_container_definition
- data/aws_vpc_peering_connection
- resource/aws_api_gateway_deployment
- resource/aws_api_gateway_gateway_response
- resource/aws_api_gateway_integration
- resource/aws_api_gateway_integration_response
- resource/aws_api_gateway_method
- resource/aws_api_gateway_method_response
- resource/aws_batch_job_definition
- resource/aws_glue_catalog_database
- resource/aws_lambda_function

refs: https://github.com/terraform-providers/terraform-provider-aws/pull/3429

### terraform-provider-google
- data/source_storage_object_signed_url
- resource/bigquery_dataset
- resource/bigquery_table
- resource/compute_instance
- resource/compute_project_metadata
- resource/dataproc_cluster
- resource/pubsub_subscription

refs: https://github.com/terraform-providers/terraform-provider-google/pull/950

## Can't find method Plugin.GetSchema

If you got errors like the following, this means your provider depends on too old terraform core to support GetSchema API, you need to update the `terraform` and the `go-cty` in the provider's dependency.

```
$ tfschema resource show azurerm_app_service
Failed to get schema from provider: rpc: can't find method Plugin.GetSchema
```

Note that this error log is output from the old `azurerm` provider, but the latest `azurerm` provider does not need a patch already.

# Contributions
Any feedback and contributions are welcome. Feel free to open an issue and submit a pull request.

# Acknowledgments
The tfschema is built on Terraform and its providers. I'm sincerely grateful to those authors.

# License
MIT
