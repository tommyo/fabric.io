# Fabric.io Go client
[![GoDoc](https://godoc.org/github.com/TommyO/fabric.io/client?status.svg)](http://godoc.org/github.com/TommyO/fabric.io/client)

A simple Go library that fetches mobile application statistics from [Fabric.io](http://fabric.io) using its ~~private~~ not publicly opened API.

This is influenced by the excellent [Fabricio project for Ruby](https://github.com/strongself/fabricio)

There is a possibility that in some point of time it may break. Use at your own risk.

## Rebuilding

To rebuild, get the [go-swagger tool](https://github.com/go-swagger/go-swagger) and run:
```
swagger generate client
```

## The future

I *may* run against the [oag tool](https://github.com/jbowes/oag) in the near future.
