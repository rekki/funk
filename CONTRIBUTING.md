# Contributing

Contributions are welcome.

## Development

Use one of the Go versions listed in the CI matrix, then run:

```bash
gofmt -w .
go test -race ./...
go vet ./...
go mod tidy
git diff --exit-code
```

Changes to exported APIs should include documentation, tests, and a clear
compatibility rationale. During `v0`, breaking changes should be intentional
and called out in release notes.

By contributing, you agree that your contribution may be distributed under
the repository's MIT license.
