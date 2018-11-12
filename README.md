# pkg

[![Documentation](https://godoc.org/github.com/arcus/pkg?status.svg)](http://godoc.org/github.com/arcus/pkg)

Collection of packages containing boilerplate utilities or shared abstractions for Go service implementations.

This repository should **not** contain any packages implementing core domain logic, client libraries for services, etc.

## Packages

- [log](./log) - Simple log interface and implementation that outputs structured (JSON) logs in a type safe way.
- [reader](./reader) - Provides a `UniversalReader` type that handles carriage returns and the BOM, originally developed for properly reading CSV data.
- [status](./status) - Defines a set of standard protocol-agnostic status codes with an HTTP mapping.
- [config](./config) - Function to initialize configuration for main.
- [service](./service) - Interfaces for handlers and middleware, as well as a custom context implementation and related errors and common simple handlers.
- [transport/http](./transport/http) - HTTP server, endpoint group, and route implementations, as well as an adaptor interface and default adaptor creator for wrapping service handlers.
