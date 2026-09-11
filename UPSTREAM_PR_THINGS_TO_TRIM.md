# Upstream PR preparation

Proposed scope cleanup for `20260908-aje_work-rebased-r2` before merging into
`taskcluster/mozilla-history`. These are recommendations, not completed changes.

## Remove or separate Quick preview infrastructure

- [ ] Exclude `deploy-quick-prod` and `deploy-quick-staging` unless upstream wants
  to maintain Quick deployments.
- [ ] Remove Quick-specific `/quick.js` loading and `#__quick_nav` styling from
  upstream-facing pages. Preserve useful local preview and redirect behavior.
- [ ] Revert the branch's change to `docs/worker-metrics.html`; it only adds Quick
  navigation styling and is unrelated to the report improvements.
- [ ] Review the root `index.html` together with the Pages workflow. Its redirect
  is needed by that proposed deployment layout; removing the entire file without
  adjusting the workflow would break the landing page.

## Separate production automation from report improvements

- [ ] Move `.github/workflows/reports.yml`, `.github/workflows/pages.yml`, and
  their production automation documentation in `README.md` to a separate PR.
- [ ] Review Taskcluster credentials and scopes, publishing permissions, the
  existing NAS job, and the Pages layout as part of that operational change.
- [ ] Keep scheduling opt-in through `ENABLE_SCHEDULED_REPORTS`. The current
  workflow does not start scheduled probes merely because it is merged.

Keep `.github/workflows/ci.yml` in the report PR. It runs Go, Python, JavaScript,
and Firefox layout tests on pull requests and pushes to `master`, with read-only
repository permissions and no Taskcluster credentials or deployment steps.

## Handle the generated report separately

- [ ] Prefer regenerating `WorkerVersions/README.md` upstream after the reporting
  code merges, or isolate its changes in a clearly identified generated-report
  commit if reviewers want to see the rendered result in this PR.
- [ ] Regenerate provenance in the upstream checkout; the current report links
  its source revision to the fork that generated it.
- [ ] Retain upstream's latest raw worker snapshot. The rebase discarded the
  superseded September 10 snapshot and preserved upstream's September 11 data;
  there is no older raw worker-data dump to include from this branch.

## Fix integration details before merging

- [ ] Add `/audit-worker-versions/audit-worker-versions` to `.gitignore`. The
  branch removed the previous broad ignore rule, leaving the executable produced
  by a build inside the auditor directory unignored.
- [ ] Reconcile the Docker Worker paragraph's migration link with the eventual
  Pages layout. It currently targets
  `https://taskcluster.github.io/mozilla-history/migration.html`, while the new
  Pages workflow places that page under `/mozilla-history/docs/migration.html`.
  Update the report template and generated report consistently if the layout
  changes; local and fork previews should also have a sensible destination.

## Keep in the core report PR

- Reporting logic, configured capacity fields, snapshot metadata, and offline
  report rendering, with their tests.
- Report layout, accessible navigation and permalinks, sorting, freshness
  notices, and History's Worker Migration subsection.
- History generation support for both legacy arrays and the new snapshot format.
- Local refresh and preview tooling, excluding Quick-specific deployment pieces.
- Unit tests, Firefox layout tests, and the new CI workflow.

This list records a scope and integration review, not a complete correctness
audit. Before opening the upstream PR, review the final diff against current
upstream `master` and rerun the relevant checks after trimming. This preparation
document can remain on the working branch and be omitted from the upstream PR.
