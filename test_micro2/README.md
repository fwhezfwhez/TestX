# Test_micro2 Service

This is the Test_micro2 service

Generated with

```
micro new test_X/test_micro2 --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.test_micro2
- Type: srv
- Alias: test_micro2

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./test_micro2-srv
```

Build a docker image
```
make docker
```