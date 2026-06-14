---
title: "Troubleshooting"
description: "The handful of things that trip people up, and how to fix each one."
weight: 40
---

## Requests start failing or returning 429

The UN SDG API is a public REST API. `unsdg` already paces requests and retries
transient failures, but a hard rate limit still means backing off. Raise the
delay with `--delay 1s` and retry later.

## Nothing returned from targets

The UN SDG API returns an empty array for invalid goal numbers. Check that the
goal number is between 1 and 17. `unsdg targets --goal 99` will exit with
code 3 (no data).

## The binary is not on your PATH

`go install` puts the binary in `$(go env GOPATH)/bin` (usually `~/go/bin`), and
a release archive leaves it wherever you unpacked it. If your shell cannot find
`unsdg`, add that directory to your `PATH`. See
[installation](/getting-started/installation/).
