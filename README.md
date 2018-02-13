# tfschema

A schema inspector for Terraform providers.

![demo](/images/tfschema-demo.gif)

Note that tfschema is under development and its interface is unstable.

# Features

- Get resource type definitons dynamically from Terraform providers.
- List available resource types.
- Autocomplete resource types in bash/zsh.
- Open official provider documents quickly by your system browser.

# Example

```bash
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

```bash
$ go get -u github.com/minamijoyo/tfschema
```

The tfschema depends on the Terraform's GetSchema API, and currently does not work unless you patch the provider.

The tfschema requires the provider's dependency library version to:

- hashicorp/terraform >= v0.10.8
- zclconf/go-cty >= 14e23b14828dd12cc7ae0956813c7e91a196e68f

For example, to update the aws provider's go-cty version, execute the following command:

```bash
$ go get -u github.com/terraform-providers/terraform-provider-aws
$ go get -u github.com/kardianos/govendor
$ govendor fetch github.com/zclconf/go-cty/...
$ go install
```

This step will be unnecessary in the future if the provider's dependency is updated officially.

# FAQ
If you got errors like the following, this means your provider does not support GetSchema API correctly, you need to update the go-cty in the provider's dependency.

```bash
$ tfschema resource show aws_security_group
Failed to get schema from provider: reading body error decoding cty.Type: gob: name not registered for interface: "github.com/terraform-providers/terraform-provider-aws/vendor/github.com/zclconf/go-cty/cty.primitiveType"
```

# Autocomplete

To enable autocomplete, execute the following command:

```bash
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

```bash
$ tfschema --help
Usage: tfschema [--version] [--help] <command> [<args>]

Available commands are:
    data
    provider
    resource
```

```bash
$ tfschema resource --help
This command is accessed by using one of the subcommands below.

Subcommands:
    browse    Browse a documentation of resource
    list      List resource types
    show      Show a type definition of resource
```

```bash
$ tfschema resource show --help
Usage: tfschema resource show [options] RESOURCE_TYPE

Options:

  -format=type    Set output format to table or json (default: table)
```

# Contributions
Any feedback and contributions are welcome. Feel free to open an issue and submit a pull request.

# Acknowledgments
The tfschema is built on Terraform and its providers. I'm sincerely grateful to these authors.

# License
MIT
