# Optional Quick previews

These helpers deploy this checkout to Mozilla Quick for contributor previews.
They are separate from upstream's production reporting and GitHub Pages
publishing workflows.

## Setup

Install and configure Mozilla's Quick CLI for your account before using these
helpers. The `quick` command must be available on your `PATH`.

Site names are explicit and independent of the checkout directory or Quick's
inferred name:

- `<username>-mozilla-history`
- `<username>-mozilla-history-staging`

The username defaults to `id -un` (your local OS account). Set
`QUICK_PREVIEW_USER` to override it, for example
`QUICK_PREVIEW_USER=aerickson scripts/quick/deploy-quick`. Names are lowercased,
non-alphanumeric runs become hyphens, and names over 63 characters are rejected.

## Usage

From the repository root:

```sh
scripts/quick/deploy-quick
scripts/quick/deploy-quick-staging
```

Both helpers resolve the repository root from their own location, so they work
when invoked by absolute path from another directory. They package the static
files from `docs/` at the site root alongside `WorkerVersions/README.md` and
`WorkerVersions/workers.json` in a temporary directory. The report opens directly
at the preview URL and loads the packaged snapshot; no root redirect is needed.
Only the packaged landing page gets the Quick SDK script. The temporary package
is removed when deployment finishes or fails, and the checkout is not modified.

In an interactive terminal, each helper displays the target URL and asks for
approval, defaulting to No. For explicitly approved automation, use:

```sh
scripts/quick/deploy-quick --confirm
scripts/quick/deploy-quick-staging --confirm
```

Without a terminal, the helpers exit unless `--confirm` is supplied. Use
`--help` for a usage summary. These commands deploy the files already in the
checkout; they do not refresh worker data or create Taskcluster probes.
