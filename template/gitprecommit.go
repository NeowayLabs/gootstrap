package template

const GitHookPreCommit = `#!/bin/bash

set -o errexit
set -o nounset

make fmt
make modtidy
make analyze
make check
`
