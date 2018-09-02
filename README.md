# pkg

Collection of packages containing boilerplate utilities or shared abstractions for Go service implementations.

This repository should **not** contain any packages implementing core domain logic, client libraries for services, etc.

## Packages

- [log](./log) - Log package which outputs structured (JSON) logs in a type safe way.
- [reader](./reader) - Package currently containing a `UniversalReader` that handles carriage returns and the BOM.
