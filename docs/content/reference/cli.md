---
title: "CLI"
description: "Every command and subcommand, with the flags that matter."
weight: 10
---

```
unsdg <command> [flags]
```

Run `unsdg <command> --help` for the full flag list on any command.

## Commands

| Command | What it does |
|---|---|
| `goals` | List all 17 UN Sustainable Development Goals |
| `targets` | List UN SDG targets (all 169, or filtered by `--goal`) |
| `version` | Print the version and exit |

## Global flags

These apply to every command:

| Flag | Default | Meaning |
|---|---|---|
| `-o`, `--output` | `auto` | Output format: `table`, `json`, `jsonl`, `csv`, `tsv`, `url`, `raw` |
| `--fields` | all | Comma-separated columns to include |
| `--no-header` | false | Omit header row in `table`, `csv`, `tsv` |
| `--template` | | Go `text/template` applied to each record |
| `-n`, `--limit` | 0 (all) | Maximum records to return |
| `-q`, `--quiet` | false | Suppress progress on stderr |
| `--delay` | 300ms | Minimum spacing between requests |
| `--timeout` | 30s | Per-request timeout |
| `--retries` | 3 | Retry attempts on 429/5xx |
| `--user-agent` | browser UA | User-Agent header |

## goals

```bash
unsdg goals [flags]
```

Fetches all 17 SDGs from the UN SDG API and prints them. No required arguments.

## targets

```bash
unsdg targets [--goal N] [flags]
```

Fetches SDG targets. Without `--goal`, returns all ~169 targets. With
`--goal 1`, returns only the targets for goal 1.

| Flag | Meaning |
|---|---|
| `--goal` | Filter by goal number (1-17) |
