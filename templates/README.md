# {{.Name}}

TODO brief intro

## Why

TODO

## Testing

Run:

```
make check
```

To get coverage results from tests:

```
make coverage
```

Open generated coverage on a browser:

```
make coverage-show
```

To perform static analysis:

```
make analyze
```

## Releasing

Run:

```
make release version=<version>
```

It will create a git tag with the provided **<version>**
and build and publish a docker image.

## Vendoring

To update vendored dependencies run:

```
make vendor
```
