# Fabric.io Go client
[![GoDoc](https://godoc.org/github.com/TommyO/fabric.io/client?status.svg)](http://godoc.org/github.com/TommyO/fabric.io/client)

A simple Go library that fetches mobile application statistics from [Fabric.io](http://fabric.io) using its ~~private~~ not publicly opened API.

This is based on the excellent [Fabricio project for Ruby](https://github.com/strongself/fabricio)

Specifically it's using the [Swagger spec](https://github.com/strongself/fabricio/blob/develop/docs/api_reference.md) from that project,
with minor tweaks. I'll do my best to keep this aligned with any changes that happen upstream.

There is a possibility that in some point of time it may break. Use at your own risk.

## Rebuilding

To rebuild, get the [swagger tool](https://github.com/go-swagger/go-swagger) and run:
```
swagger generate client -c fabric
```
