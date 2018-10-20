# Gouken

Simple wrapper around grpc-go, the scaffold to start building a grpc service.


## Development

Use `dep init -v` to install the dependencies.

Because gouken itself is only the collection of tools that helps you build grpc apps,
it is agnostic to what version of grpc/protobuf/other libs you project is using.
So, to avoid constraints that may affect your project, gouken does not
include `Gopkg.toml` or `Gopkg.lock` in project files.
