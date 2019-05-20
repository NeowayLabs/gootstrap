package template

const Readme = `# {{.Project}}

TODO brief intro

## Why

TODO

## Testing

Run:` +

	"\n\n```\nmake check\n```\n" +

	"Open generated coverage on a browser:\n\n" +

	"```\nmake coverage\n```\n" +

	"To perform static analysis:\n\n```\nmake analyze\n```\n" +

	`
## Releasing

Run:` +

	"\n\n```\nmake release version=<version>\n```\n" +

	`
It will create a git tag with the provided **<version>**
and build and publish a docker image.
`
