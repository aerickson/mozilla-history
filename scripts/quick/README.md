# Optional Quick previews

These helpers deploy this checkout to Mozilla Quick for contributor previews.
They are separate from upstream's production reporting and GitHub Pages
publishing workflows.

## Setup

Install and configure Mozilla's Quick CLI for your account before using these
helpers. The `quick` command must be available on your `PATH`. To check the
default preview destination, run `quick url` from the repository root.

Quick resolves the default site name from the checkout and your configuration.
Use a personal or team preview namespace when configuring Quick. The staging
helper appends `-staging` to that resolved site name; it rejects names longer
than Quick's 63-character limit.

## Usage

From the repository root:

```sh
scripts/quick/deploy-quick
scripts/quick/deploy-quick-staging
```

Both helpers resolve the repository root from their own location and deploy
the entire checkout, so you can also invoke them by absolute path from another
directory. The root landing page redirects to the report in `docs/`.

In an interactive terminal, each helper displays the target URL and asks for
approval, defaulting to No. For explicitly approved automation, use:

```sh
scripts/quick/deploy-quick --confirm
scripts/quick/deploy-quick-staging --confirm
```

Without a terminal, the helpers exit unless `--confirm` is supplied. Use
`--help` for a usage summary. These commands deploy the files already in the
checkout; they do not refresh worker data or create Taskcluster probes.
