---
title: "Configuration"
description: "Environment variables and global flags."
weight: 20
---

`unsdg` needs almost no configuration. The global flags cover the cases that come up.

## Global flags

| Flag | Default | Meaning |
|---|---|---|
| `--delay` | 300ms | Minimum gap between requests |
| `--timeout` | 30s | Per-request timeout |
| `--retries` | 3 | Retry attempts on 429/5xx |
| `--user-agent` | browser UA | User-Agent header |

## Environment variables

None. All settings are passed via flags.
