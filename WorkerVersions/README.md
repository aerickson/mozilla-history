

# Worker Pool Versions

This report shows the latest detailed inventory of Firefox CI worker pools alongside historical trends from earlier snapshots. Worker implementation and version are inferred from the failure log produced when each pool is given an intentionally malformed probe task; image and capacity metadata come from Worker Manager. Summary counts represent worker pools, not individual workers or tasks.


## Generic Worker

Total: `428`

Count by version:

_Source: version information parsed from the log artifact produced when each worker claims an intentionally malformed probe task. The task is expected to fail with a malformed-payload exception; known worker implementations and versions are identified from their distinct log output._

| Version | Count |
| :--- | ---: |
| 100.0.1 | 4 |
| 100.5.0 | 2 |
| 108.0.0 | 332 |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 15 |
| 61.0.0 | 1 |
| 64.3.0 | 21 |
| 65.1.0 | 1 |
| 67.1.0 | 1 |
| 84.1.2 | 7 |
| 88.0.2 | 3 |
| 91.0.2 | 1 |
| 96.2.2 | 4 |
| 96.2.3 | 12 |
| 99.2.0 | 20 |


Count by image:

_Source: image references in each pool's live Worker Manager launch configuration at report time. A value may represent multiple configured images; unknown means no supported image reference was found._

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-09-03 | 5 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-09-03 | 16 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-09-03 | 128 |
| unknown | 139 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-09-03 | 60 |
|  | 32 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-09-03 | 8 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 5 | 5 |
| **adhoc-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 19 | 19 |
| **adhoc-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 5 | 5 |
| **adhoc-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3 | 3 |
| **app-services-1/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 246 | 246 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3 | 3 |
| **app-services-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 5 | 5 |
| **app-services-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 556 | 556 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 22 | 22 |
| **app-services-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 20 | 20 |
| **code-analysis-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-coverage/bot** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 332 | 332 |
| **code-review/bot** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 11079 | 11079 |
| **comm-1/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1154 | 1154 |
| **comm-1/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 26 | 26 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 6 | 6 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3 | 3 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3 | 3 |
| **comm-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 5 | 5 |
| **comm-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 338 | 338 |
| **comm-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 10 | 10 |
| **comm-1/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 4 | 4 |
| **comm-2/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 5426 | 5426 |
| **comm-3/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3064 | 3064 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 10 | 10 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 8 | 8 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 4 | 4 |
| **comm-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 834 | 834 |
| **comm-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 186 | 186 |
| **comm-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 23 | 23 |
| **comm-3/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-t/misc** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 170 | 170 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 14 | 14 |
| **comm-t/t-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3194 | 3194 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 65 | 65 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 5807 | 5807 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 118 | 118 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **comm-t/win11-64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 7471 | 7471 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 29 | 29 |
| **enterprise-1/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1389 | 1389 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 34 | 34 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 602 | 602 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 48 | 48 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1776 | 1776 |
| **enterprise-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 68 | 68 |
| **enterprise-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 106 | 106 |
| **enterprise-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 39 | 39 |
| **enterprise-1/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 7 | 7 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 5927 | 5927 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 12 | 12 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1090 | 1090 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 75 | 75 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1428 | 1428 |
| **enterprise-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 803 | 803 |
| **enterprise-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 101 | 101 |
| **enterprise-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 14 | 14 |
| **enterprise-3/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 387 | 387 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 1072 | 1072 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 4 | 4 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1620 | 1620 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 6017 | 6017 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3 | 3 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 8395 | 8395 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/win10-64-2009** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 780 | 780 |
| **enterprise-t/win11-64-24h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/win11-64-24h2-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/win11-64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 6719 | 6719 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 596 | 596 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 4021 | 4021 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 514 | 514 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 8491 | 8491 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 83 | 83 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 11134 | 11134 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1428 | 1428 |
| **gecko-1/b-linux-docker-updatebot-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1517 | 1517 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 163 | 163 |
| **gecko-1/b-linux-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-1/b-linux-medium** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 17005 | 17005 |
| **gecko-1/b-linux-xlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 1110 | 1110 |
| **gecko-1/b-win2022-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 4 | 4 |
| **gecko-1/b-win2022-updatebot** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 13 | 13 |
| **gecko-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 120 | 120 |
| **gecko-1/b-win2025-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2025-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2025-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1862 | 1862 |
| **gecko-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 79 | 79 |
| **gecko-1/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 21 | 21 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 22 | 22 |
| **gecko-1/win11-a64-25h2-builder-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 2 | 2 |
| **gecko-2/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 106 | 106 |
| **gecko-2/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 150 | 150 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 27 | 27 |
| **gecko-2/b-linux-docker-updatebot-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 39 | 39 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 9 | 9 |
| **gecko-2/b-linux-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 164 | 164 |
| **gecko-2/b-linux-xlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 15 | 15 |
| **gecko-2/b-win2022-updatebot** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-2/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-win2025-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 13 | 13 |
| **gecko-2/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 21920 | 21920 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 34 | 34 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 26451 | 26451 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2661 | 2661 |
| **gecko-3/b-linux-docker-updatebot-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 46 | 46 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3677 | 3677 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 497 | 497 |
| **gecko-3/b-linux-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 29197 | 29197 |
| **gecko-3/b-linux-xlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 87 | 87 |
| **gecko-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 5972 | 5972 |
| **gecko-3/b-win2022-headless** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-3/b-win2022-updatebot** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 12 | 12 |
| **gecko-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 27 | 27 |
| **gecko-3/b-win2025-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2025-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2025-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 4282 | 4282 |
| **gecko-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 71 | 71 |
| **gecko-3/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 14 | 14 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 151 | 151 |
| **gecko-t/misc** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1667 | 1667 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 10235 | 10235 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 41 | 41 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 40 | 40 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 32 | 32 |
| **gecko-t/t-linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 806 | 806 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 3580 | 3580 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 42036 | 42036 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 4 | 4 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 74021 | 74021 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 21 | 21 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 5633 | 5633 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 227242 | 227242 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 91 | 91 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6339 | 6339 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 6458 | 6458 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 85 | 85 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 992 | 992 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 16 | 16 |
| **gecko-t/win10-64-2009-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 64 | 64 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 15 | 15 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 8 | 8 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1156 | 1156 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 13 | 13 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 31 | 31 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 1706 | 1706 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 162 | 162 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 5 | 5 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 18 | 18 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 155 | 155 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 6 | 6 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 6 | 6 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 13 | 13 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 9 | 9 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 5 | 5 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 132 | 132 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 22700 | 22700 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 18 | 18 |
| **gecko-t/win11-64-25h2-gpu-perf-experiment** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 12945 | 12945 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 41 | 41 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 224 | 224 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2572 | 2572 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 8 | 8 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 4476 | 4476 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 5 | 5 |
| **gecko-t/win11-a64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 302 | 302 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 13 | 13 |
| **glean-1/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 41 | 41 |
| **glean-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 4 | 4 |
| **glean-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 25 | 25 |
| **glean-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 8 | 8 |
| **glean-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 168 | 168 |
| **mobile-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 15 | 15 |
| **mobile-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 224 | 224 |
| **mobile-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 126 | 126 |
| **mobile-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 9 | 9 |
| **mozilla-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 28 | 28 |
| **mozilla-3/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 461 | 461 |
| **mozilla-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 6 | 6 |
| **mozilla-t/pre-commit** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 19 | 19 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 915 | 915 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 133 | 133 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 36 | 36 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 138 | 138 |
| **mozillavpn-1/b-linux-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 107 | 107 |
| **mozillavpn-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 33 | 33 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **mozillavpn-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 13 | 13 |
| **mozillavpn-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 32 | 32 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 5 | 5 |
| **mozillavpn-3/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 54 | 54 |
| **mozillavpn-3/b-linux-large** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 59 | 59 |
| **mozillavpn-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 15 | 15 |
| **mozillavpn-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 7 | 7 |
| **mozillavpn-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 13 | 13 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 3 | 3 |
| **nss-1/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 20 | 20 |
| **nss-1/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 95 | 95 |
| **nss-1/b-win2022-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 3 | 3 |
| **nss-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **nss-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 37 | 37 |
| **nss-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **nss-1/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 279 | 279 |
| **nss-1/win11-a64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 2 | 2 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 5 | 5 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 2 | 2 |
| **nss-1/win11-a64-25h2-builder-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 9 | 9 |
| **nss-3/b-linux-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 179 | 179 |
| **nss-3/b-win2022** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 159 | 159 |
| **nss-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **nss-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 14 | 14 |
| **nss-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **nss-3/images-aarch64** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 774 | 774 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 3 | 3 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | arm64 | 1.27.0 | 843 | 843 |
| **nss-t/t-linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 1241 | 1241 |
| **nss-t/t-linux-docker-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **nss-t/win11-a64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 2 | 2 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | arm64 | 1.27.0 | 19 | 19 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-fuzzing/bugmon-processor-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/ci-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows-ngpu** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **releng-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 54 | 54 |
| **releng-1/linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 25 | 25 |
| **releng-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 25 | 25 |
| **releng-3/linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 32 | 32 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
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
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 226 | 226 |
| **relops-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **relops-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **relops-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 22 | 22 |
| **relops-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **reviewer-assignment/bot** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 6 | 6 |
| **scriptworker-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 11 | 11 |
| **scriptworker-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 58 | 58 |
| **scriptworker-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 7 | 7 |
| **scriptworker-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 37 | 37 |
| **taskgraph-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 24 | 24 |
| **taskgraph-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 31 | 31 |
| **taskgraph-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 12 | 12 |
| **taskgraph-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 13 | 13 |
| **taskgraph-t/linux-docker** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 124 | 124 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 24 | 24 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 108.0.0 | multiuser | ae7697a544 | windows | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 15 | 15 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 10 | 10 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 5 | 5 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 44 | 44 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 183 | 183 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 205 | 205 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 97 | 97 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 54 | 54 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 8 | 8 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 12 | 12 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 25 | 25 |
| **translations-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 107 | 107 |
| **xpi-1/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 44 | 44 |
| **xpi-1/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 6 | 6 |
| **xpi-3/b-linux** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 21 | 21 |
| **xpi-3/decision** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 5 | 5 |
| **xpi-3/images** | generic-worker | 108.0.0 | multiuser | ae7697a544 | linux | amd64 | 1.27.0 | 6 | 6 |


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
| **infra/build-decision** | docker-worker | 38.0.5 | 7 | 56 |
| **proj-fuzzing/bugmon-pernosco** | docker-worker | 44.23.4 | 2 | 2 |


## Script Worker

Total: `43`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
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
| **scriptworker-k8s/gecko-3-pushmsix** | Scriptworker | <no value> | 0 | 0 |
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

Total: `3`


Count by image:

_Source: image references in each pool's live Worker Manager launch configuration at report time. A value may represent multiple configured images; unknown means no supported image reference was found._

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 1 |
|  | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 55 | 55 |


## Version not determined [^2]

Total: `18`


Count by image:

_Source: image references in each pool's live Worker Manager launch configuration at report time. A value may represent multiple configured images; unknown means no supported image reference was found._

| Version | Count |
| :--- | ---: |
|  | 3 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-qioehfdraiiishpuvgvl | 7 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 1 |
| unknown | 4 |
| ami-02619e55246806e8d | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-arm64-mqscbutmsbawtytlroxk | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-1/b-win2025-alpha** |  | Version not determined; task not (yet) claimed | 153 | 153 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | Version not determined; task not (yet) claimed | 2 | 2 |
| **gecko-t/win11-64-25h2** |  | Version not determined; task not (yet) claimed | 109109 | 109109 |
| **mozillavpn-1/b-win2025-alpha** |  | Version not determined; task not (yet) claimed | 152 | 152 |
| **nss-1/b-win2025-alpha** |  | Version not determined; task not (yet) claimed | 151 | 151 |
| **proj-autophone/gecko-t-lambda-perf-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-fuzzing/bugmon-monitor** |  | Version not determined; task not (yet) claimed | 97 | 97 |
| **proj-fuzzing/bugmon-pernosco-staging** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/bugmon-processor** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/ci** |  | Version not determined; task not (yet) claimed | 69 | 69 |
| **proj-fuzzing/ci-arm64** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/decision** |  | Version not determined; task not (yet) claimed | 225 | 225 |
| **proj-fuzzing/grizzly-reduce-worker** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/grizzly-reduce-worker-android** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/nss-corpus-update-worker** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **releng-hardware/enterprise-1-b-osx-arm64** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.
