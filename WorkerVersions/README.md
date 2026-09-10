

# Worker Pool Versions

This report shows the latest detailed inventory of Firefox CI worker pools alongside historical trends from earlier snapshots.

- **Implementation and version** are inferred from the failure log produced when each pool receives an intentionally malformed probe task.
- **Image and capacity metadata** come from Worker Manager.
- **Summary values** are counts of worker pools, not individual workers or tasks.


## Generic Worker

Total: `424`

Count by version:

_Source: version information parsed from the log artifact produced when each worker claims an intentionally malformed probe task. The task is expected to fail with a malformed-payload exception; known worker implementations and versions are identified from their distinct log output._

| Version | Count |
| :--- | ---: |
| 100.0.1 | 6 |
| 100.5.0 | 2 |
| 108.0.0 | 73 |
| 108.1.0 | 269 |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 16 |
| 61.0.0 | 1 |
| 64.3.0 | 21 |
| 65.1.0 | 1 |
| 67.1.0 | 1 |
| 84.1.2 | 7 |
| 87.0.0 | 2 |
| 88.0.2 | 3 |
| 91.0.2 | 2 |
| 96.2.2 | 4 |
| 96.2.3 | 12 |


Count by image:

_Source: image references in each pool's live Worker Manager launch configuration at report time. A value may represent multiple configured images; unknown means no supported image reference was found._

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-09-09 | 128 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-9ed5a812d3d844ac9e4e | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-09-09 | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
|  | 37 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-09-09 | 8 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-09-09 | 60 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| unknown | 126 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-09-09 | 16 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **adhoc-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **adhoc-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **adhoc-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **app-services-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 249 | 249 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **app-services-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 8 | 8 |
| **app-services-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **app-services-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1022 | 1022 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **app-services-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 29 | 29 |
| **app-services-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **code-analysis-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **code-analysis-1/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 38 | 38 |
| **code-analysis-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **code-analysis-3/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 26 | 26 |
| **code-coverage/bot** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 392 | 392 |
| **code-review/bot** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 10754 | 10754 |
| **comm-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 778 | 778 |
| **comm-1/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **comm-1/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1082 | 1082 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 9 | 9 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 7 | 7 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **comm-1/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-1/b-linux-xlarge** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **comm-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 11 | 11 |
| **comm-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 443 | 443 |
| **comm-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 29 | 29 |
| **comm-1/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **comm-2/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **comm-2/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **comm-2/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/b-linux-xlarge** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **comm-2/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-2/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **comm-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3494 | 3494 |
| **comm-3/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 9 | 9 |
| **comm-3/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2113 | 2113 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 45 | 45 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 27 | 27 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **comm-3/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **comm-3/b-linux-xlarge** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **comm-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 551 | 551 |
| **comm-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 164 | 164 |
| **comm-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 43 | 43 |
| **comm-3/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **comm-t/misc** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 326 | 326 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **comm-t/t-linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 14 | 14 |
| **comm-t/t-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4057 | 4057 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 43 | 43 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 5517 | 5517 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **comm-t/win11-64-24h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 74 | 74 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **comm-t/win11-64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 6503 | 6503 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **comm-t/win11-a64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 3 | 3 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | arm64 | 1.27.1 | 18 | 18 |
| **enterprise-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3642 | 3642 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 18 | 18 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1213 | 1213 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 59 | 59 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2362 | 2362 |
| **enterprise-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 223 | 223 |
| **enterprise-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 119 | 119 |
| **enterprise-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 75 | 75 |
| **enterprise-1/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 15 | 15 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 3 | 3 |
| **enterprise-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 8282 | 8282 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 25 | 25 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1619 | 1619 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 125 | 125 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1643 | 1643 |
| **enterprise-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 1151 | 1151 |
| **enterprise-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 124 | 124 |
| **enterprise-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 136 | 136 |
| **enterprise-3/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 24 | 24 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 4 | 4 |
| **enterprise-t/misc** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 608 | 608 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1180 | 1180 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **enterprise-t/t-linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1626 | 1626 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 8136 | 8136 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 11699 | 11699 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **enterprise-t/win10-64-2009** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 848 | 848 |
| **enterprise-t/win11-64-24h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 5 | 5 |
| **enterprise-t/win11-64-24h2-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 4 | 4 |
| **enterprise-t/win11-64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 8516 | 8516 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 3 | 3 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 513 | 513 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3597 | 3597 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 869 | 869 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 4 | 4 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 11105 | 11105 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **gecko-1/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 59 | 59 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 13208 | 13208 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1744 | 1744 |
| **gecko-1/b-linux-docker-updatebot-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1944 | 1944 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-1/b-linux-kvm** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 267 | 267 |
| **gecko-1/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-1/b-linux-medium** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 18902 | 18902 |
| **gecko-1/b-linux-xlarge** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 22 | 22 |
| **gecko-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 1433 | 1433 |
| **gecko-1/b-win2022-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-1/b-win2022-headless** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 4 | 4 |
| **gecko-1/b-win2022-updatebot** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 15 | 15 |
| **gecko-1/b-win2025-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 612 | 612 |
| **gecko-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2030 | 2030 |
| **gecko-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 314 | 314 |
| **gecko-1/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 68 | 68 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 51 | 51 |
| **gecko-1/win11-a64-25h2-builder-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | arm64 | 1.27.1 | 6 | 6 |
| **gecko-2/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 9 | 9 |
| **gecko-2/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 45 | 45 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 13 | 13 |
| **gecko-2/b-linux-docker-updatebot-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 14 | 14 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-2/b-linux-kvm** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 6 | 6 |
| **gecko-2/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-2/b-linux-medium** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 53 | 53 |
| **gecko-2/b-linux-xlarge** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-2/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 4 | 4 |
| **gecko-2/b-win2022-updatebot** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-2/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 12 | 12 |
| **gecko-2/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-2/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **gecko-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 29179 | 29179 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **gecko-3/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 43 | 43 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 30556 | 30556 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3318 | 3318 |
| **gecko-3/b-linux-docker-updatebot-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 42 | 42 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4396 | 4396 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-3/b-linux-kvm** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 579 | 579 |
| **gecko-3/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **gecko-3/b-linux-medium** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 28943 | 28943 |
| **gecko-3/b-linux-xlarge** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 102 | 102 |
| **gecko-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 7492 | 7492 |
| **gecko-3/b-win2022-headless** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-3/b-win2022-updatebot** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 13 | 13 |
| **gecko-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4797 | 4797 |
| **gecko-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 157 | 157 |
| **gecko-3/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 29 | 29 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 149 | 149 |
| **gecko-t/misc** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3263 | 3263 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 10283 | 10283 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 4 | 4 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 65 | 65 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 66 | 66 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 39 | 39 |
| **gecko-t/t-linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 796 | 796 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 2877 | 2877 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 47158 | 47158 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 7 | 7 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 86029 | 86029 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 41 | 41 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 8694 | 8694 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 226370 | 226370 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 190 | 190 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6678 | 6678 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **gecko-t/win10-64-2009** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 6436 | 6436 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 76 | 76 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 1098 | 1098 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 14 | 14 |
| **gecko-t/win10-64-2009-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 108 | 108 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 23 | 23 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 14 | 14 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1890 | 1890 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 25 | 25 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-24h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2001 | 2001 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 161 | 161 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 19 | 19 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 169 | 169 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 7 | 7 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 11 | 11 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 21 | 21 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 11 | 11 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 7 | 7 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 5 | 5 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 6 | 6 |
| **gecko-t/win11-64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 110067 | 110067 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 132 | 132 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 22446 | 22446 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 26 | 26 |
| **gecko-t/win11-64-25h2-gpu-perf-experiment** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 12998 | 12998 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 48 | 48 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 277 | 277 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 5 | 5 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 4349 | 4349 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 11 | 11 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 3 | 3 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 4686 | 4686 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 6 | 6 |
| **gecko-t/win11-a64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 317 | 317 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | arm64 | 1.27.1 | 22 | 22 |
| **glean-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 78 | 78 |
| **glean-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 14 | 14 |
| **glean-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **glean-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 50 | 50 |
| **glean-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 13 | 13 |
| **glean-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **infra/build-decision-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **mobile-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 192 | 192 |
| **mobile-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 13 | 13 |
| **mobile-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **mobile-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 301 | 301 |
| **mobile-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 155 | 155 |
| **mobile-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **mozilla-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **mozilla-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **mozilla-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **mozilla-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 15 | 15 |
| **mozilla-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **mozilla-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **mozilla-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 449 | 449 |
| **mozilla-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 6 | 6 |
| **mozilla-t/pre-commit** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 17 | 17 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 956 | 956 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **mozilla-t/t-linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 116 | 116 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 6 | 6 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **mozillavpn-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 188 | 188 |
| **mozillavpn-1/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 157 | 157 |
| **mozillavpn-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 52 | 52 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 5 | 5 |
| **mozillavpn-1/b-win2025-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 609 | 609 |
| **mozillavpn-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 5 | 5 |
| **mozillavpn-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 13 | 13 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 5 | 5 |
| **mozillavpn-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 115 | 115 |
| **mozillavpn-3/b-linux-large** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 135 | 135 |
| **mozillavpn-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 27 | 27 |
| **mozillavpn-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 18 | 18 |
| **mozillavpn-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 4 | 4 |
| **nss-1/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 23 | 23 |
| **nss-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 54 | 54 |
| **nss-1/b-win2022-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 4 | 4 |
| **nss-1/b-win2025-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 609 | 609 |
| **nss-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 17 | 17 |
| **nss-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **nss-1/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **nss-1/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 145 | 145 |
| **nss-1/win11-a64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 3 | 3 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | arm64 | 1.27.1 | 5 | 5 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 3 | 3 |
| **nss-1/win11-a64-25h2-builder-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | arm64 | 1.27.1 | 3 | 3 |
| **nss-3/b-linux-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 130 | 130 |
| **nss-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 139 | 139 |
| **nss-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 14 | 14 |
| **nss-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **nss-3/images-aarch64** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 3 | 3 |
| **nss-3/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 708 | 708 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 3 | 3 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | arm64 | 1.27.1 | 784 | 784 |
| **nss-t/t-linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1197 | 1197 |
| **nss-t/t-linux-docker-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **nss-t/win11-a64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 3 | 3 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | arm64 | 1.27.1 | 24 | 24 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-alpha-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **proj-fuzzing/bugmon-processor-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 3 | 3 |
| **proj-fuzzing/ci-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 3 | 3 |
| **proj-fuzzing/grizzly-reduce-worker-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 3 | 3 |
| **proj-fuzzing/grizzly-reduce-worker-windows-ngpu** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 3 | 3 |
| **proj-taskcluster/ci** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 15 | 15 |
| **proj-taskcluster/gw-ci-macos** | generic-worker | 108.0.0 | multiuser | ae7697a544 | darwin | arm64 | 1.27.0 | 2 | 2 |
| **proj-taskcluster/gw-ubuntu-24-04** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 238 | 238 |
| **proj-taskcluster/gw-ubuntu-24-04-gui** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 11 | 11 |
| **proj-taskcluster/gw-windows-2022** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 41 | 41 |
| **proj-taskcluster/gw-windows-2022-gui** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 41 | 41 |
| **proj-taskcluster/release** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 1 | 1 |
| **releng-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 52 | 52 |
| **releng-1/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 7 | 7 |
| **releng-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 31 | 31 |
| **releng-3/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 32 | 32 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** | generic-worker | 65.1.0 | insecure | 1a085daa37 | linux | amd64 | 1.22.3 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404-bug2055283** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 67.1.0 | insecure | 0e62d3bf79 | darwin | amd64 | 1.22.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 100.5.0 | insecure | f5fc37cc8f | darwin | amd64 | 1.26.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** | generic-worker | 100.5.0 | multiuser | f5fc37cc8f | darwin | arm64 | 1.26.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-2600-m4** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 108.1.0 | multiuser | f86624a8bf | windows | amd64 | 1.27.1 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 113 | 113 |
| **relops-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **relops-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **relops-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 22 | 22 |
| **relops-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **reviewer-assignment/bot** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 6 | 6 |
| **scriptworker-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 23 | 23 |
| **scriptworker-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 135 | 135 |
| **scriptworker-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 15 | 15 |
| **scriptworker-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 85 | 85 |
| **taskgraph-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 8 | 8 |
| **taskgraph-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 20 | 20 |
| **taskgraph-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **taskgraph-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **taskgraph-t/linux-docker** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 44 | 44 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 7 | 7 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 24 | 24 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 8 | 8 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 157 | 157 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 4 | 4 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 193 | 193 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 151 | 151 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 62 | 62 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 64 | 64 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 4 | 4 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 8 | 8 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 26 | 26 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 32 | 32 |
| **translations-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **xpi-1/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 82 | 82 |
| **xpi-1/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 46 | 46 |
| **xpi-1/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **xpi-3/b-linux** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **xpi-3/decision** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |
| **xpi-3/images** | generic-worker | 108.1.0 | multiuser | f86624a8bf | linux | amd64 | 1.27.1 | 3 | 3 |


## Docker Worker

Total: `2`

Count by version:

_Source: version information parsed from the log artifact produced when each worker claims an intentionally malformed probe task. The task is expected to fail with a malformed-payload exception; known worker implementations and versions are identified from their distinct log output._

| Version | Count |
| :--- | ---: |
| 38.0.5 | 1 |
| 44.23.4 | 1 |


Count by image:

_Source: image references in each pool's live Worker Manager launch configuration at report time. A value may represent multiple configured images; unknown means no supported image reference was found._

| Version | Count |
| :--- | ---: |
| ami-03e4f8db63254ce7e,ami-0a6e926238859761c,ami-0b5dd0bbb670ec80e | 1 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **infra/build-decision** | docker-worker | 38.0.5 | 19 | 152 |
| **proj-fuzzing/bugmon-pernosco** | docker-worker | 44.23.4 | 3 | 3 |


## Script Worker

Total: `43`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-addon** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/translations-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/adhoc-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/comm-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-adhoc-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-comm-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-enterprise-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-gecko-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-mozillavpn-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/enterprise-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/gecko-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-prov-v1/mozillavpn-signing-mac14m2** | Scriptworker | <no value> | 0 | 0 |


## No artifacts found [^1]

Total: `4`


Count by image:

_Source: image references in each pool's live Worker Manager launch configuration at report time. A value may represent multiple configured images; unknown means no supported image reference was found._

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 2 |
|  | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 58 | 58 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 4 | 4 |


## Version not determined [^2]

Total: `30`


Count by image:

_Source: image references in each pool's live Worker Manager launch configuration at report time. A value may represent multiple configured images; unknown means no supported image reference was found._

| Version | Count |
| :--- | ---: |
| unknown | 20 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-arm64-mqscbutmsbawtytlroxk | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-9ed5a812d3d844ac9e4e | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| ami-02619e55246806e8d | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **comm-1/b-win2025** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **comm-3/b-win2025** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **enterprise-1/b-win2025** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **enterprise-3/b-win2025** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-1/b-win2025** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-1/b-win2025-headless** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-1/b-win2025-updatebot** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-1/b-win2025-xxlarge** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-2/b-win2025** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-2/b-win2025-updatebot** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-3/b-win2025** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-3/b-win2025-headless** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-3/b-win2025-updatebot** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-3/b-win2025-xxlarge** |  | Version not determined; task not (yet) claimed | 94 | 94 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **mozilla-1/b-win2025** |  | Version not determined; task not (yet) claimed | 93 | 93 |
| **mozilla-3/b-win2025** |  | Version not determined; task not (yet) claimed | 92 | 92 |
| **mozillavpn-1/b-win2025** |  | Version not determined; task not (yet) claimed | 93 | 93 |
| **mozillavpn-3/b-win2025** |  | Version not determined; task not (yet) claimed | 92 | 92 |
| **nss-1/b-win2025** |  | Version not determined; task not (yet) claimed | 92 | 92 |
| **nss-3/b-win2025** |  | Version not determined; task not (yet) claimed | 92 | 92 |
| **proj-fuzzing/bugmon-monitor** |  | Version not determined; task not (yet) claimed | 91 | 91 |
| **proj-fuzzing/bugmon-pernosco-staging** |  | Version not determined; task not (yet) claimed | 5 | 5 |
| **proj-fuzzing/bugmon-processor** |  | Version not determined; task not (yet) claimed | 5 | 5 |
| **proj-fuzzing/ci** |  | Version not determined; task not (yet) claimed | 72 | 72 |
| **proj-fuzzing/ci-arm64** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **proj-fuzzing/decision** |  | Version not determined; task not (yet) claimed | 225 | 225 |
| **proj-fuzzing/grizzly-reduce-worker** |  | Version not determined; task not (yet) claimed | 5 | 5 |
| **proj-fuzzing/grizzly-reduce-worker-android** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **proj-fuzzing/nss-corpus-update-worker** |  | Version not determined; task not (yet) claimed | 6 | 6 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.
