

# Worker Pool Versions


## Generic Worker

Total: `411`

Count by version:

| Version | Count |
| :--- | ---: |
| 100.0.1 | 4 |
| 100.4.0 | 12 |
| 100.5.0 | 2 |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 16 |
| 61.0.0 | 1 |
| 64.3.0 | 21 |
| 65.1.0 | 1 |
| 67.1.0 | 1 |
| 84.1.2 | 7 |
| 87.0.0 | 2 |
| 88.0.2 | 2 |
| 91.0.2 | 1 |
| 96.2.2 | 5 |
| 96.2.3 | 12 |
| 99.2.0 | 104 |
| 99.2.1 | 216 |


Count by image:

| Version | Count |
| :--- | ---: |
| unknown | 120 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-05-04 | 60 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-05-04 | 127 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-05-04 | 16 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
|  | 35 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-05-04 | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-05-04 | 8 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **adhoc-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **adhoc-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **adhoc-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **app-services-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 216 | 216 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **app-services-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 869 | 869 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 21 | 21 |
| **app-services-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 33 | 33 |
| **code-analysis-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 9 | 9 |
| **code-coverage/bot** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 323 | 323 |
| **code-review/bot** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13238 | 13238 |
| **comm-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1528 | 1528 |
| **comm-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 180 | 180 |
| **comm-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **comm-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 4 | 4 |
| **comm-2/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-2/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6139 | 6139 |
| **comm-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2220 | 2220 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9 | 9 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 5 | 5 |
| **comm-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 610 | 610 |
| **comm-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 170 | 170 |
| **comm-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 96 | 96 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2203 | 2203 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4626 | 4626 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7550 | 7550 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 10 | 10 |
| **enterprise-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1276 | 1276 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 3 | 3 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 212 | 212 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 16 | 16 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 986 | 986 |
| **enterprise-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 54 | 54 |
| **enterprise-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 72 | 72 |
| **enterprise-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **enterprise-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2741 | 2741 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 621 | 621 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 60 | 60 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 378 | 378 |
| **enterprise-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 284 | 284 |
| **enterprise-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 50 | 50 |
| **enterprise-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **enterprise-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **enterprise-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 141 | 141 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 637 | 637 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1142 | 1142 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3413 | 3413 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5155 | 5155 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 322 | 322 |
| **enterprise-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4723 | 4723 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 295 | 295 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1977 | 1977 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 191 | 191 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 8066 | 8066 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 89 | 89 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 3 | 3 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10350 | 10350 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1034 | 1034 |
| **gecko-1/b-linux-docker-updatebot-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1 | 1 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1126 | 1126 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-1/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 201 | 201 |
| **gecko-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **gecko-1/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 18640 | 18640 |
| **gecko-1/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **gecko-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 855 | 855 |
| **gecko-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 8 | 8 |
| **gecko-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 979 | 979 |
| **gecko-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 88 | 88 |
| **gecko-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 18 | 18 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 30 | 30 |
| **gecko-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **gecko-2/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 300 | 300 |
| **gecko-2/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 139 | 139 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 229 | 229 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 28 | 28 |
| **gecko-2/b-linux-docker-updatebot-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 51 | 51 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-2/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **gecko-2/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 380 | 380 |
| **gecko-2/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **gecko-2/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 11 | 11 |
| **gecko-2/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-2/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 19 | 19 |
| **gecko-2/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 30620 | 30620 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 27 | 27 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23733 | 23733 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2194 | 2194 |
| **gecko-3/b-linux-docker-updatebot-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9 | 9 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2813 | 2813 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 577 | 577 |
| **gecko-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **gecko-3/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 29278 | 29278 |
| **gecko-3/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 260 | 260 |
| **gecko-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5864 | 5864 |
| **gecko-3/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 12 | 12 |
| **gecko-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5359 | 5359 |
| **gecko-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 40 | 40 |
| **gecko-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 12 | 12 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 192 | 192 |
| **gecko-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 622 | 622 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8722 | 8722 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 73 | 73 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | arm64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 89 | 89 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 33 | 33 |
| **gecko-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1625 | 1625 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2766 | 2766 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 43364 | 43364 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 66520 | 66520 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23603 | 23603 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 161121 | 161121 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6052 | 6052 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5012 | 5012 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 737 | 737 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 8 | 8 |
| **gecko-t/win10-64-2009-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 98 | 98 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 11 | 11 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4540 | 4540 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 986 | 986 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 26 | 26 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6 | 6 |
| **gecko-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 8682 | 8682 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 9 | 9 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 733 | 733 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6 | 6 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 46 | 46 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 14 | 14 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 103862 | 103862 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 21331 | 21331 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-gpu-perf-experiment** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 11535 | 11535 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 657 | 657 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1857 | 1857 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 14 | 14 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4178 | 4178 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 303 | 303 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **glean-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 36 | 36 |
| **glean-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 8 | 8 |
| **glean-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12 | 12 |
| **glean-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **glean-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 184 | 184 |
| **mobile-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **mobile-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 254 | 254 |
| **mobile-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 146 | 146 |
| **mobile-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1 | 1 |
| **mozilla-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 40 | 40 |
| **mozilla-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 442 | 442 |
| **mozilla-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/pre-commit** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 598 | 598 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 25 | 25 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 310 | 310 |
| **mozillavpn-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 266 | 266 |
| **mozillavpn-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 86 | 86 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **mozillavpn-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **mozillavpn-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 40 | 40 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 5 | 5 |
| **mozillavpn-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 173 | 173 |
| **mozillavpn-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 170 | 170 |
| **mozillavpn-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 35 | 35 |
| **mozillavpn-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 22 | 22 |
| **mozillavpn-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 16 | 16 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **nss-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 137 | 137 |
| **nss-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 173 | 173 |
| **nss-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **nss-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 18 | 18 |
| **nss-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 594 | 594 |
| **nss-1/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 10 | 10 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **nss-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 124 | 124 |
| **nss-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 169 | 169 |
| **nss-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **nss-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 771 | 771 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 1044 | 1044 |
| **nss-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1598 | 1598 |
| **nss-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 13 | 13 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-test-1** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **proj-fuzzing/bugmon-processor-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 3 | 3 |
| **proj-fuzzing/ci-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows-ngpu** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **releng-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 69 | 69 |
| **releng-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 110 | 110 |
| **releng-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 41 | 41 |
| **releng-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 83 | 83 |
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
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 67.1.0 | insecure | 0e62d3bf79 | darwin | amd64 | 1.22.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 100.5.0 | insecure | f5fc37cc8f | darwin | amd64 | 1.26.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-ipv6** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** | generic-worker | 100.5.0 | multiuser | f5fc37cc8f | darwin | arm64 | 1.26.4 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 305 | 305 |
| **relops-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **relops-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **relops-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **relops-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **scriptworker-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 91 | 91 |
| **scriptworker-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15 | 15 |
| **scriptworker-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 127 | 127 |
| **taskgraph-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **taskgraph-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **taskgraph-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **taskgraph-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9 | 9 |
| **taskgraph-t/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 112 | 112 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10 | 10 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 44 | 44 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 72 | 72 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 22 | 22 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9 | 9 |
| **translations-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 86 | 86 |
| **xpi-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 52 | 52 |
| **xpi-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 26 | 26 |
| **xpi-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **xpi-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |


## Docker Worker

Total: `2`

Count by version:

| Version | Count |
| :--- | ---: |
| 38.0.5 | 1 |
| 44.23.4 | 1 |


Count by image:

| Version | Count |
| :--- | ---: |
| ami-03e4f8db63254ce7e,ami-0a6e926238859761c,ami-0b5dd0bbb670ec80e | 1 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **infra/build-decision** | docker-worker | 38.0.5 | 5 | 40 |
| **proj-fuzzing/bugmon-pernosco** | docker-worker | 44.23.4 | 2 | 2 |


## Script Worker

Total: `45`



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
| **scriptworker-k8s/comm-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-addon** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-lando-dev** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-pushflatpak** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-pushmsix** | Scriptworker | <no value> | 0 | 0 |
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

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 2 |
|  | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 171 | 171 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `12`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 2 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-qioehfdraiiishpuvgvl | 7 |
| ami-02619e55246806e8d | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-arm64-mqscbutmsbawtytlroxk | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/bugmon-monitor** |  | Version not determined; task not (yet) claimed | 90 | 90 |
| **proj-fuzzing/bugmon-pernosco-staging** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **proj-fuzzing/bugmon-processor** |  | Version not determined; task not (yet) claimed | 6 | 6 |
| **proj-fuzzing/ci** |  | Version not determined; task not (yet) claimed | 51 | 51 |
| **proj-fuzzing/ci-arm64** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/decision** |  | Version not determined; task not (yet) claimed | 213 | 213 |
| **proj-fuzzing/grizzly-reduce-worker** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/grizzly-reduce-worker-android** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/nss-corpus-update-worker** |  | Version not determined; task not (yet) claimed | 8 | 8 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.
