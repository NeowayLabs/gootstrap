# Gootstrap

Gootstrap aims to help you bootstrap Go projects.

It provides just a set of scripts that you can copy
to your Go project to help you:

* Vendor things, like conan (no semver, you have integration tests right ?)
* Run tests recursively without including the vendor directory
* Run tests with coverage, coalescing all the packages reports in one
* Cool static analysis
* Embedding --version on binaries using git commit tag

Perhaps it may have more scripts on the future, but the idea is to
keep it as simple as possible while useful in our context.

## Installation

The only dependency of this project is [nash](https://github.com/NeowayLabs/nash).
Installing it is trivial. After installing it, just copy the scripts to
wherever you want them to be and use them.
