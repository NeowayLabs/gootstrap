# Gootstrap

Gootstrap aims to help you bootstrap Go projects.

It provides just a set of scripts that you can copy
to your Go project to help you:

* Vendor things, like conan (no semver, you have integration tests right ?)
* Run tests recursively without including the vendor directory
* Run tests with coverage, coalescing all the packages reports in one
* Cool static analysis
* Embedding --version on binaries using git commit tag

All commands happens on top of docker since it
is the basis of our development environments and production deployment.

It has much less usefulness in Go than other languages (since Go is very
simple, specially with vendoring) but at least avoids differences on Go
versions between developers and CI servers, also promotes uniformity on how
we work with other languages.

## Installation

Run:

```
go get github.com/NeowayLabs/gootstrap
```

And that is it =)

## Usage

The help of the project should be enough, just be
aware that besides running everything inside docker
it is important that the project is inside the
GOPATH of your host.

Some details on how the code is mapped to the containers
depend on this, also if you do so you will be able
to build and run tests directly on your host too if
you want (autocomplete and code navigation will also
work properly in vendored dependencies).
