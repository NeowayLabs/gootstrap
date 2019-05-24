package template

const GitHookPreCommit = `#!/bin/bash

set -o errexit
set -o nounset

make fmt
make modtidy
make static-analysis
make check
`
