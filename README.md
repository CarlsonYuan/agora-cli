## Compile yourself
```shell
$ git clone git@github.com:CarlsonYuan/agora-cli.git
$ cd agora-cli
$ go build ./cmd/agora-cli
$ ./agora-cli --version
agora-cli version 0.0.1
```

## Supported commands
* config
```
Manage app configurations

Usage:
  agora-cli config [command]

Available Commands:
  default     Set an application as the default
  list        List all applications
  new         Add a new application
  remove      Remove one or more application.
```
* chat
```
Allows you to interact with your Chat applications

Usage:
  agora-cli chat [command]

Available Commands:
  query-user  Query user
```
