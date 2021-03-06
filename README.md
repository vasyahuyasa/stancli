# Stancli

Utility for send arbitrary data to nats streaming server from command-line.

#### Packages

[![stancli](https://img.shields.io/aur/version/stancli.svg?label=stancli)](https://aur.archlinux.org/packages/stancli/)
[![license](https://img.shields.io/github/license/vasyahuyasa/stancli.svg)](https://github.com/vasyahuyasa/stancli/blob/master/LICENSE)
[![Scc Count Badge](https://sloc.xyz/github/vasyahuyasa/stancli/)](https://github.com/vasyahuyasa/stancli/)

## Installation

```shell
$ go get -u github.com/vasyahuyasa/stancli
```

## Command-Line Options

```
Usage: stancli [OPTIONS] [FILE]

Send data to Nats streaming server.
With no FILE, or when FILE is -, read standard input.

Options:
  -client_id string
        Client ID (default "stan_cli")
  -cluster string
        Cluster ID (default "nats")
  -help
        This help text
  -subject string
        Subject (default "default")
  -url string
        NATS server url (default "nats://client:123456@localhost:4222")
```

## Example usage

```shell
$ cat example/changeEmployee.txt | stancli -subject employees
```
