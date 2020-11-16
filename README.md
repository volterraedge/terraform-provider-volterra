# Terraform Provider for Volterra

![](https://www.volterra.io/img/main/logo.svg)

Terraform provider for managing your Volterra resources.

## Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) = 0.12.x
-	[Go](https://golang.org/doc/install) >= 1.14 (to build the provider plugin)

## Building The Provider
------------------------

Checkout the repository using below command:

```sh
$ git clone https://github.com/volterraedge/terraform-provider-volterra.git
```

Enter the provider directory and build the provider:

```sh
$ cd terraform-provider-volterra
$ cd cmd/terraform-provider-volterra; go build .
```

## Using the provider
---------------------

### Released provider version

Run [`terraform init`](https://www.terraform.io/docs/commands/init.html) and terraform will automatically install the provider.

To specify a particular provider version when installing released providers, see the [Terraform documentation on provider versioning](https://www.terraform.io/docs/configuration/providers.html#version-provider-versions).

### Custom built provider in your Terraform environment

Run below command

```
make install-tf-volterra
```

## Examples

Check out the [examples located](examples)


## Running Unit Tests
---------------------

Execute below command to run all unit tests

```sh
$ go test -v ./...
```

To execute unit test, use `-run` flag

```sh
$ go test -v -run=TestAccNamespaceBasic ./...
```