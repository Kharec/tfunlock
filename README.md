# tfunlock

`tfunlock` is a simple Golang program to unlock your terraform cloud workspaces. 

## Disclaimer

In fact, it's my first program in Go. It is one of my goals to master this language before December 31, 2024.

So it may be unclean or mispackaged, be kind.

This program cannot unlock the terraform workspaces that were manually locked by an organization user.

## Prerequisites

You must set your [terraform cloud token](https://developer.hashicorp.com/terraform/cloud-docs/users-teams-organizations/api-tokens) in an environment variable named `TF_CLOUD_TOKEN` before using this program.

‼️ **Change** the line :

```
const (
	organization = "" // put your organization name here
)
```

To set up your terraform cloud organization.

Also be sure that `make` is installed on your system.

## Build

Once you fufilled the prerequisites, just run `make` to build the executable.

## Install

Move the executable to a directory that is in your path.
