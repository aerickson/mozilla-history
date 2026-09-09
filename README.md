# mozilla-history
Taskcluster history from https://firefox-ci-tc.services.mozilla.com deployment
for:

* [Clients](/Clients)
* [Hooks](/Hooks)
* [Roles](/Roles)
* [Worker Pools](/WorkerPools)

For https://community-tc.services.mozilla.com history, please see
[community-history](https://github.com/taskcluster/community-history).

Please note, until 9 November 2019 this repository stored the history for the
https://taskcluster.net deployment. On 9 November 2019 the
https://taskcluster.net deployment was decommissioned and the
https://firefox-ci-tc.services.mozilla.com was instated.

## Entity filenames

The files are named after the entities themselves, except for the following
character conversions:

  1. `*` -> `★`

This conversion avoids illegal filenames.

  2. `/` -> `⁄`

Rather than creating nested subdirectories, this conversion avoids directory
names colliding with entity filenames.

## Installing

```
go get github.com/taskcluster/mozilla-history@v1.0.0
```

## Running

```
unset TASKCLUSTER_CLIENT_ID TASKCLUSTER_ACCESS_TOKEN TASKCLUSTER_CERTIFICATE
export TASKCLUSTER_ROOT_URL='https://firefox-ci-tc.services.mozilla.com'
mozilla-history
```

This will populate subdirectories `Clients`, `Hooks`, `Roles` and `WorkerPools`
of the current directory.

## Refreshing All Local Data

For a complete local refresh, set Taskcluster credentials and run:

```sh
export TASKCLUSTER_CLIENT_ID='...'
export TASKCLUSTER_ACCESS_TOKEN='...'
./fetch_and_generate.py
```

The script builds temporary copies of the Go tools, refreshes `Clients`,
`Hooks`, `Roles`, and `WorkerPools`, schedules worker-version probes, polls their
sealed task group until every task reaches a terminal state, writes
`WorkerVersions`, and rebuilds `docs/history.json`. It does not pull, commit,
push, or deploy anything. Review the generated changes before committing them.

Use `--poll-interval SECONDS` to change the 60-second polling interval. If the
script is interrupted after scheduling probes, resume without creating another
group using `--task-group-id TASK_GROUP_ID`. `TASKCLUSTER_ROOT_URL`,
`REPORT_SCHEDULER_ID`, and `REPORT_PREFIX` can be overridden in the environment.

Worker versions require two phases because Worker Manager configuration does
not expose the implementation and version actually running in each pool. The
first phase schedules an intentionally malformed task on every pool so the
worker identifies itself in its task log. Once all tasks are resolved, the
second phase reads those logs and generates the snapshot. Polling replaces the
older fixed-delay assumption while preserving that probe-and-collect design.

## Rendering an Existing Worker Snapshot

The worker-version report can be regenerated from a saved snapshot without
Taskcluster credentials or probe tasks:

```sh
go run ./audit-worker-versions render \
  WorkerVersions/workers.json \
  WorkerVersions/README.md
```

Omit the output path to print the generated Markdown to standard output.

To preview the generated report in the website, serve the repository root:

```sh
python3 -m http.server 8000
```

Then open <http://localhost:8000/docs/index-local.html>. The local preview uses
`WorkerVersions/README.md` and `docs/history.json` from the checkout. The
published page continues to load the current report from GitHub.

## Production Automation

The existing `run-reports.sh` and `audit.sh` scripts implement the repository's
production publishing workflow. They include git and production-site behavior;
use `fetch_and_generate.py` for local refreshes.

### Prerequisites
- Valid Taskcluster credentials must be set in the environment variables:
  - `TASKCLUSTER_CLIENT_ID`
  - `TASKCLUSTER_ACCESS_TOKEN`

### How it works
1. `run-reports.sh` executes `audit.sh`
2. `audit.sh` schedules tasks for each worker pool to extract worker implementation details from logs
3. Results are stored in the `WorkerVersions` directory
4. `mozilla-history` stores Taskcluster configurations in their respective directories:
   - Hooks definitions in `Hooks/`
   - Roles definitions in `Roles/`
   - Worker Pool definitions in `WorkerPools/`
5. `build-docs-history.sh` collects every past revision of `WorkerVersions/workers.json` into
   `docs/history.json`, which feeds the two GitHub Pages views:
   - `docs/index.html` &mdash; the current report plus the full history tables and graphs
   - `docs/migration.html` &mdash; an animated month-by-month timeline of the docker-worker to
     generic-worker migration and of the generic-worker version rollout
