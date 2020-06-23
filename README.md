# tfschema

A schema inspector for Terraform providers.

# Features

- Get resource type definitions dynamically from Terraform providers via go-plugin protocol.
- List available resource types.
- Autocomplete resource types in bash/zsh.
- Open official provider documents quickly by your system web browser.
- Terraform v0.13 support (minimum requirements: Terraform >= v0.11)

![demo](/images/tfschema-demo.gif)

# Getting Started

```
$ brew install minamijoyo/tfschema/tfschema

$ echo 'provider "aws" {}' > main.tf
$ terraform init
```

```
$ tfschema resource list aws | grep aws_iam_user
aws_iam_user
aws_iam_user_group_membership
aws_iam_user_login_profile
aws_iam_user_policy
aws_iam_user_policy_attachment
aws_iam_user_ssh_key
```

```
$ tfschema resource show aws_iam_user
+----------------------+-------------+----------+----------+----------+-----------+
| ATTRIBUTE            | TYPE        | REQUIRED | OPTIONAL | COMPUTED | SENSITIVE |
+----------------------+-------------+----------+----------+----------+-----------+
| arn                  | string      | false    | false    | true     | false     |
| force_destroy        | bool        | false    | true     | false    | false     |
| id                   | string      | false    | true     | true     | false     |
| name                 | string      | true     | false    | false    | false     |
| path                 | string      | false    | true     | false    | false     |
| permissions_boundary | string      | false    | true     | false    | false     |
| tags                 | map(string) | false    | true     | false    | false     |
| unique_id            | string      | false    | false    | true     | false     |
+----------------------+-------------+----------+----------+----------+-----------+
```

# Install

If you are Mac OSX user:

```
$ brew install minamijoyo/tfschema/tfschema
```

or

If you have Go 1.14+ development environment:

```bash
$ git clone https://github.com/minamijoyo/tfschema
$ cd tfschema
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
If your provider support Terraform >= v0.11, it should work.
Strictly speaking, the tfschema depends on the Terraform's GetSchema API.

The tfschema requires the provider's dependency library version to:

- hashicorp/terraform >= v0.10.8
- zclconf/go-cty >= 14e23b14828dd12cc7ae0956813c7e91a196e68f (2018/01/06)

# Rules of finding provider's binary
When `terraform init` command is executed, provider's binary is installed under the auto installed directory ( .terraform/plugins/`<OS>_<ARCH>` ) by default.
The tfschema can use the same provider's binary as terraform uses, so you can run `tfschema` command in the same directory where you run the `terraform` command.

The tfschema finds provider's binary under the following directories.

1. current directory
2. same directory as `tfschema` executable
3. user vendor directory ( terraform.d/plugins/`<OS>_<ARCH>` )
4. auto installed directory ( .terraform/plugins/`<SOURCE ADDRESS>`/`<VERSION>`/`<OS>_<ARCH>` )
5. legacy auto installed directory ( .terraform/plugins/`<OS>_<ARCH>` )
6. global plugin directory ( $HOME/.terraform.d/plugins )
7. global plugin directory with os and arch ( $HOME/.terraform.d/plugins/`<OS>_<ARCH>` )
8. gopath ( $GOPATH/bin )

If you are Mac OSX user, `<OS>_<ARCH>` is `darwin_amd64`.
The `<SOURCE ADDRESS>` is a fully qualified provider name in Terraform 0.13+. (e.g. `registry.terraform.io/hashicorp/aws`)

Note that it doesn't have exactly the same behavior of Terraform because of some reasons:

- Support multiple Terraform versions
- Can't import internal packages of Terraform and it's too complicated to support
- For debug

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
# Contributions
Any feedback and contributions are welcome. Feel free to open an issue and submit a pull request.

# Acknowledgments
The tfschema is built on Terraform and its providers. I'm sincerely grateful to those authors.

# License
MIT
