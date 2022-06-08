# vSwarm-proto

vSwarm-proto is a utility for [vSwarm](https://github.com/ease-lab/vSwarm/) containing all the proto files and client applications necessary to run benchmarks using [relay](https://github.com/ease-lab/vSwarm/tree/main/tools/relay).

## CLI command
The `cmd` subfolder contains a command line interface for invoking any kind of function that implement one of the protocols in this repository.

## Releasing new versions
In order to release a new version do the following:
1. Increment the version number in `cmd/version.go`.
2. Build the release with `make release`
3. Push the new version with `git push <VERSION TAG>`

