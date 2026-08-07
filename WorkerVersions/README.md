

# Worker Pool Versions


## Generic Worker

Total: `406`

Count by version:

| Version | Count |
| :--- | ---: |
| 100.0.1 | 4 |
| 100.5.0 | 2 |
| 103.0.1 | 12 |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 14 |
| 61.0.0 | 1 |
| 64.3.0 | 21 |
| 65.1.0 | 1 |
| 67.1.0 | 1 |
| 84.1.2 | 7 |
| 88.0.2 | 3 |
| 96.2.2 | 5 |
| 96.2.3 | 12 |
| 99.2.0 | 103 |
| 99.2.1 | 216 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| unknown | 119 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-05-04 | 16 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-05-04 | 5 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-05-04 | 8 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-05-04 | 127 |
|  | 31 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-05-04 | 60 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **adhoc-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **adhoc-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **adhoc-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 290 | 290 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **app-services-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 898 | 898 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 27 | 27 |
| **app-services-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 40 | 40 |
| **code-analysis-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 33 | 33 |
| **code-coverage/bot** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 262 | 262 |
| **code-review/bot** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13415 | 13415 |
| **comm-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1641 | 1641 |
| **comm-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 3 | 3 |
| **comm-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 8 | 8 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 298 | 298 |
| **comm-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **comm-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 3 | 3 |
| **comm-2/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **comm-2/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3885 | 3885 |
| **comm-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2273 | 2273 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 8 | 8 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 707 | 707 |
| **comm-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 138 | 138 |
| **comm-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 20 | 20 |
| **comm-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 81 | 81 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2914 | 2914 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6338 | 6338 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10721 | 10721 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 11 | 11 |
| **enterprise-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2626 | 2626 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 6 | 6 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 354 | 354 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1514 | 1514 |
| **enterprise-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 297 | 297 |
| **enterprise-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 90 | 90 |
| **enterprise-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 39 | 39 |
| **enterprise-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 8 | 8 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 12 | 12 |
| **enterprise-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3048 | 3048 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 17 | 17 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 845 | 845 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 67 | 67 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 530 | 530 |
| **enterprise-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 386 | 386 |
| **enterprise-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 57 | 57 |
| **enterprise-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 35 | 35 |
| **enterprise-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 7 | 7 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **enterprise-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 202 | 202 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 939 | 939 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1689 | 1689 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4188 | 4188 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **enterprise-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 564 | 564 |
| **enterprise-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 453 | 453 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3098 | 3098 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 196 | 196 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6255 | 6255 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 52 | 52 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 3 | 3 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 8975 | 8975 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 765 | 765 |
| **gecko-1/b-linux-docker-updatebot-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1010 | 1010 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-1/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 147 | 147 |
| **gecko-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **gecko-1/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12462 | 12462 |
| **gecko-1/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 689 | 689 |
| **gecko-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-1/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1283 | 1283 |
| **gecko-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **gecko-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 19 | 19 |
| **gecko-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **gecko-2/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 133 | 133 |
| **gecko-2/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 298 | 298 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 37 | 37 |
| **gecko-2/b-linux-docker-updatebot-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 61 | 61 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 4 | 4 |
| **gecko-2/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 19 | 19 |
| **gecko-2/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 236 | 236 |
| **gecko-2/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **gecko-2/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10 | 10 |
| **gecko-2/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-2/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 21 | 21 |
| **gecko-2/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 26196 | 26196 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **gecko-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 24736 | 24736 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2258 | 2258 |
| **gecko-3/b-linux-docker-updatebot-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 37 | 37 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3625 | 3625 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-3/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 652 | 652 |
| **gecko-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 28455 | 28455 |
| **gecko-3/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 198 | 198 |
| **gecko-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7436 | 7436 |
| **gecko-3/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 19 | 19 |
| **gecko-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6820 | 6820 |
| **gecko-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 41 | 41 |
| **gecko-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 8 | 8 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 257 | 257 |
| **gecko-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 982 | 982 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 9235 | 9235 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 91 | 91 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | arm64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 97 | 97 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 4 | 4 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 30 | 30 |
| **gecko-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1130 | 1130 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4092 | 4092 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 45110 | 45110 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 5 | 5 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 68346 | 68346 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 60 | 60 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 16939 | 16939 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 183243 | 183243 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 179 | 179 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 6296 | 6296 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5387 | 5387 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6 | 6 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 807 | 807 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 13 | 13 |
| **gecko-t/win10-64-2009-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 64 | 64 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6 | 6 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2893 | 2893 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 600 | 600 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 19 | 19 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 5 | 5 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6 | 6 |
| **gecko-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5512 | 5512 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 21 | 21 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 476 | 476 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 22 | 22 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 12 | 12 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 124673 | 124673 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2025 | 2025 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 25336 | 25336 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 259 | 259 |
| **gecko-t/win11-64-25h2-gpu-perf-experiment** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 13635 | 13635 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 462 | 462 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 371 | 371 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 14 | 14 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2043 | 2043 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 137 | 137 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4307 | 4307 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 385 | 385 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **glean-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 66 | 66 |
| **glean-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **glean-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 26 | 26 |
| **glean-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **glean-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 139 | 139 |
| **mobile-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15 | 15 |
| **mobile-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 283 | 283 |
| **mobile-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 170 | 170 |
| **mobile-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **mozilla-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 70 | 70 |
| **mozilla-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 442 | 442 |
| **mozilla-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/pre-commit** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 19 | 19 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 103.0.1 | multiuser | 85c75bc86b | linux | amd64 | 1.26.4 | 510 | 510 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 73 | 73 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 48 | 48 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 191 | 191 |
| **mozillavpn-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 177 | 177 |
| **mozillavpn-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 40 | 40 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **mozillavpn-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **mozillavpn-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23 | 23 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **mozillavpn-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 101 | 101 |
| **mozillavpn-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 104 | 104 |
| **mozillavpn-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 27 | 27 |
| **mozillavpn-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15 | 15 |
| **mozillavpn-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 60 | 60 |
| **nss-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 114 | 114 |
| **nss-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **nss-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **nss-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 260 | 260 |
| **nss-1/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 9 | 9 |
| **nss-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 122 | 122 |
| **nss-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 174 | 174 |
| **nss-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **nss-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 570 | 570 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 756 | 756 |
| **nss-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1129 | 1129 |
| **nss-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 7 | 7 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-fuzzing/bugmon-processor-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/ci-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows-ngpu** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **releng-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 71 | 71 |
| **releng-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 54 | 54 |
| **releng-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23 | 23 |
| **releng-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 49 | 49 |
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
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 67.1.0 | insecure | 0e62d3bf79 | darwin | amd64 | 1.22.4 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 100.5.0 | insecure | f5fc37cc8f | darwin | amd64 | 1.26.4 | 0 | 0 |
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
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 245 | 245 |
| **relops-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **relops-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **relops-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12 | 12 |
| **relops-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **reviewer-assignment/bot** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **scriptworker-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 16 | 16 |
| **scriptworker-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 104 | 104 |
| **scriptworker-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **scriptworker-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 99 | 99 |
| **taskgraph-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 33 | 33 |
| **taskgraph-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 45 | 45 |
| **taskgraph-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **taskgraph-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 14 | 14 |
| **taskgraph-t/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 210 | 210 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 27 | 27 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15 | 15 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 82 | 82 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 156 | 156 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 111 | 111 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 84 | 84 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 131 | 131 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9 | 9 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23 | 23 |
| **translations-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 101 | 101 |
| **xpi-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 51 | 51 |
| **xpi-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **xpi-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **xpi-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |


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
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 1 |
| ami-03e4f8db63254ce7e,ami-0a6e926238859761c,ami-0b5dd0bbb670ec80e | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **infra/build-decision** | docker-worker | 38.0.5 | 3 | 24 |
| **proj-fuzzing/bugmon-pernosco** | docker-worker | 44.23.4 | 2 | 2 |


## Script Worker

Total: `37`



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
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing-dev** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 98 | 98 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `18`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 6 |
| unknown | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-arm64-mqscbutmsbawtytlroxk | 1 |
| ami-02619e55246806e8d | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-05-04 | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-qioehfdraiiishpuvgvl | 7 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **enterprise-t/t-linux-docker-noscratch-amd** |  | Version not determined; task not (yet) claimed | 8480 | 8480 |
| **enterprise-t/win11-64-25h2** |  | Version not determined; task not (yet) claimed | 6948 | 6948 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **proj-autophone/gecko-t-lambda-perf-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-fuzzing/bugmon-monitor** |  | Version not determined; task not (yet) claimed | 97 | 97 |
| **proj-fuzzing/bugmon-pernosco-staging** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/bugmon-processor** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/ci** |  | Version not determined; task not (yet) claimed | 66 | 66 |
| **proj-fuzzing/ci-arm64** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/decision** |  | Version not determined; task not (yet) claimed | 212 | 212 |
| **proj-fuzzing/grizzly-reduce-worker** |  | Version not determined; task not (yet) claimed | 5 | 5 |
| **proj-fuzzing/grizzly-reduce-worker-android** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **proj-fuzzing/nss-corpus-update-worker** |  | Version not determined; task not (yet) claimed | 9 | 9 |
| **releng-hardware/gecko-t-osx-1015-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-ipv6** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.
