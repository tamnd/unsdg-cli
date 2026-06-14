---
title: "Quick start"
description: "Run your first unsdg command."
weight: 30
---

Once `unsdg` is on your `PATH`:

```bash
unsdg goals               # list all 17 SDGs as a table
unsdg goals -o json       # full JSON
unsdg targets             # list all 169 targets
unsdg targets --goal 1    # targets for goal 1 only
unsdg targets -o jsonl | jq .title
```

Output auto-adapts: a table on a terminal, JSONL when piped. Pick a format
explicitly with `-o table|json|jsonl|csv|tsv`.

Keep only the fields you need:

```bash
unsdg goals --fields code,title
```

Limit the number of records:

```bash
unsdg goals -n 5
```
