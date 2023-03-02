/*
Copyright © 2022-2023 Oliver Götz <developer@geekgasm.eu>
*/
package build

// filled by go build -ldflags="-X build.version=1.0 ..." or goreleaser

var Version string

var Commit string

var Date string
