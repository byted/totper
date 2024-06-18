# totper

A simple TOTP manager CLI written in Go.

# Install

```
> go install github.com/byted/totper
> totper help
A tool to manage TOTP secrets. Secrets are stored in your systems secrets storage.

Usage:
  totper [command]

Available Commands:
  add         Add a new TOTP account
  completion  Generate the autocompletion script for the specified shell
  get-totp    Gets the current TOTP for the account
  help        Help about any command
  list        List all TOTP accounts
  remove      Remove a TOTP account

Flags:
      --config string   config file (default is $HOME/.config/totper.yaml)
  -h, --help            help for totper

Use "totper [command] --help" for more information about a command.
`
