# unsdg

Browse UN Sustainable Development Goals from the command line.

`unsdg` is a single pure-Go binary. It speaks to the UN SDG API over plain
HTTPS, shapes the responses into clean records, and pipes into the rest of your
tools. No API key, nothing to run alongside it.

## Install

```bash
go install github.com/tamnd/unsdg-cli/cmd/unsdg@latest
```

Or grab a prebuilt binary from the [releases](https://github.com/tamnd/unsdg-cli/releases), or run
the container image:

```bash
docker run --rm ghcr.io/tamnd/unsdg:latest --help
```

## Usage

```bash
unsdg goals               # list all 17 SDGs
unsdg targets             # list all 169 targets
unsdg targets --goal 1    # targets for goal 1 only
unsdg goals -o json       # JSON output
unsdg goals -o jsonl | jq .title
```

## Development

```
cmd/unsdg/   thin main, wires cli.Root into fang
cli/         the cobra command tree
unsdg/       the library: HTTP client and data models
docs/        tago documentation site
```

```bash
make build      # ./bin/unsdg
make test       # go test ./...
make vet        # go vet ./...
```

## Releasing

Push a version tag and GitHub Actions runs GoReleaser, which builds the
archives, Linux packages, the multi-arch GHCR image, checksums, SBOMs, and a
cosign signature:

```bash
git tag v0.1.0
git push --tags
```

The Homebrew and Scoop steps self-disable until their tokens exist, so the first
release works with no extra secrets.

## License

Apache-2.0. See [LICENSE](LICENSE).
