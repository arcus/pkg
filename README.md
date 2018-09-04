# pkg

Collection of packages containing boilerplate utilities or shared abstractions for Go service implementations.

This repository should **not** contain any packages implementing core domain logic, client libraries for services, etc.

## Packages

- [log](./log) - Simple log interface and implementation that outputs structured (JSON) logs in a type safe way.
- [reader](./reader) - Provides a `UniversalReader` type that handles carriage returns and the BOM, originally developed for properly reading CSV data.
- [status](./status) - Defines a set of standard protocol-agnostic status codes with an HTTP mapping.
