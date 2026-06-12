

# Worker Pool Versions


## Generic Worker

Total: `384`

Count by version:

| Version | Count |
| :--- | ---: |
| 100.3.0 | 12 |
| 45.0.0 | 1 |
| 60.3.4 | 17 |
| 61.0.0 | 1 |
| 64.3.0 | 21 |
| 84.1.2 | 7 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 91.0.2 | 1 |
| 95.0.3 | 1 |
| 96.2.2 | 6 |
| 96.2.3 | 12 |
| 99.2.0 | 89 |
| 99.2.1 | 213 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-05-04 | 125 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-05-04 | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
|  | 30 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| unknown | 101 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-05-04 | 16 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-05-04 | 8 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-05-04 | 59 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **adhoc-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **adhoc-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **adhoc-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **app-services-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 814 | 814 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 31 | 31 |
| **app-services-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 871 | 871 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 26 | 26 |
| **app-services-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 30 | 30 |
| **code-analysis-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-coverage/bot** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 279 | 279 |
| **code-review/bot** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9118 | 9118 |
| **comm-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1210 | 1210 |
| **comm-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 4 | 4 |
| **comm-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 58 | 58 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-1/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **comm-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 50 | 50 |
| **comm-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 278 | 278 |
| **comm-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23 | 23 |
| **comm-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 3 | 3 |
| **comm-2/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **comm-2/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3705 | 3705 |
| **comm-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 4 | 4 |
| **comm-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2015 | 2015 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **comm-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 668 | 668 |
| **comm-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 111 | 111 |
| **comm-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 34 | 34 |
| **comm-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 90 | 90 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **comm-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2355 | 2355 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 66 | 66 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5415 | 5415 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 107 | 107 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5864 | 5864 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 20 | 20 |
| **enterprise-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1135 | 1135 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 15 | 15 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 192 | 192 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 37 | 37 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1429 | 1429 |
| **enterprise-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 109 | 109 |
| **enterprise-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 63 | 63 |
| **enterprise-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 61 | 61 |
| **enterprise-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 6 | 6 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 6 | 6 |
| **enterprise-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1599 | 1599 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 11 | 11 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 749 | 749 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 87 | 87 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 386 | 386 |
| **enterprise-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 228 | 228 |
| **enterprise-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 42 | 42 |
| **enterprise-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 80 | 80 |
| **enterprise-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 7 | 7 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 220 | 220 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 564 | 564 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1593 | 1593 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4695 | 4695 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7723 | 7723 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **enterprise-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 248 | 248 |
| **enterprise-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 32 | 32 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4977 | 4977 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 263 | 263 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1914 | 1914 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 18 | 18 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 8058 | 8058 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 13 | 13 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11919 | 11919 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 285 | 285 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1040 | 1040 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 128 | 128 |
| **gecko-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **gecko-1/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15906 | 15906 |
| **gecko-1/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9 | 9 |
| **gecko-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 983 | 983 |
| **gecko-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 13 | 13 |
| **gecko-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1168 | 1168 |
| **gecko-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12 | 12 |
| **gecko-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 10 | 10 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 20 | 20 |
| **gecko-2/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-2/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **gecko-2/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **gecko-2/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 25720 | 25720 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 15 | 15 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 27679 | 27679 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 746 | 746 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4594 | 4594 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 3 | 3 |
| **gecko-3/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 964 | 964 |
| **gecko-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 17 | 17 |
| **gecko-3/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 27152 | 27152 |
| **gecko-3/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 159 | 159 |
| **gecko-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6331 | 6331 |
| **gecko-3/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 11 | 11 |
| **gecko-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5136 | 5136 |
| **gecko-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 140 | 140 |
| **gecko-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 27 | 27 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 116 | 116 |
| **gecko-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 771 | 771 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8220 | 8220 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 59 | 59 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | arm64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 69 | 69 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 3 | 3 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 30 | 30 |
| **gecko-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1289 | 1289 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3520 | 3520 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 35483 | 35483 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 3 | 3 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 61636 | 61636 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 21 | 21 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6802 | 6802 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 174428 | 174428 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 72 | 72 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 5414 | 5414 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3482 | 3482 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 68 | 68 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 339 | 339 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6 | 6 |
| **gecko-t/win10-64-2009-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 78 | 78 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 15 | 15 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7 | 7 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1504 | 1504 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 315 | 315 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 44 | 44 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2045 | 2045 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 120 | 120 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1139 | 1139 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10 | 10 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 169 | 169 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10 | 10 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 8 | 8 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 16 | 16 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 8 | 8 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 69082 | 69082 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 15756 | 15756 |
| **gecko-t/win11-64-25h2-gpu-perf-experiment** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 53 | 53 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7034 | 7034 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 296 | 296 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1759 | 1759 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3344 | 3344 |
| **gecko-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 251 | 251 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 23 | 23 |
| **glean-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 92 | 92 |
| **glean-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 8 | 8 |
| **glean-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 46 | 46 |
| **glean-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 14 | 14 |
| **glean-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 189 | 189 |
| **mobile-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **mobile-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 282 | 282 |
| **mobile-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 161 | 161 |
| **mobile-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **mozilla-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 25 | 25 |
| **mozilla-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 447 | 447 |
| **mozilla-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **mozilla-t/pre-commit** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 100.3.0 | multiuser | fe314d9231 | linux | amd64 | 1.26.4 | 533 | 533 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 21 | 21 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12 | 12 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 392 | 392 |
| **mozillavpn-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 299 | 299 |
| **mozillavpn-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 83 | 83 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15 | 15 |
| **mozillavpn-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 40 | 40 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 5 | 5 |
| **mozillavpn-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 206 | 206 |
| **mozillavpn-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 240 | 240 |
| **mozillavpn-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 43 | 43 |
| **mozillavpn-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 27 | 27 |
| **mozillavpn-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15 | 15 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **nss-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 20 | 20 |
| **nss-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 175 | 175 |
| **nss-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **nss-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 26 | 26 |
| **nss-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 5 | 5 |
| **nss-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 689 | 689 |
| **nss-1/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 6 | 6 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 1 | 1 |
| **nss-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 257 | 257 |
| **nss-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 18 | 18 |
| **nss-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 1 | 1 |
| **nss-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1246 | 1246 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 12 | 12 |
| **nss-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1694 | 1694 |
| **nss-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 17 | 17 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 35 | 35 |
| **releng-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 27 | 27 |
| **releng-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 25 | 25 |
| **releng-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 40 | 40 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64-vms** | generic-worker | 95.0.3 | multiuser | a996b8269b | darwin | arm64 | 1.25.5 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-1804** | generic-worker | 61.0.0 | simple | 3bd4419b4b | linux | amd64 | 1.22.1 | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404** | generic-worker | 88.0.2 | insecure | fcaf4a25fc | linux | amd64 | 1.24.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4** | generic-worker | 91.0.2 | multiuser | 06628b3721 | darwin | arm64 | 1.25.3 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-ipv6** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 223 | 223 |
| **relops-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **relops-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **relops-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9 | 9 |
| **relops-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **scriptworker-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 61 | 61 |
| **scriptworker-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9 | 9 |
| **scriptworker-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 53 | 53 |
| **taskgraph-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 22 | 22 |
| **taskgraph-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 14 | 14 |
| **taskgraph-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15 | 15 |
| **taskgraph-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15 | 15 |
| **taskgraph-t/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 244 | 244 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 26 | 26 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12 | 12 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 17 | 17 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 77 | 77 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 376 | 376 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 191 | 191 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 123 | 123 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 50 | 50 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12 | 12 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 37 | 37 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 24 | 24 |
| **translations-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 46 | 46 |
| **xpi-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23 | 23 |
| **xpi-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **xpi-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **xpi-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |


## Docker Worker

Total: `1`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **infra/build-decision** | docker-worker | 38.0.5 | 10 | 80 |


## Script Worker

Total: `40`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/comm-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-tree** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-lando** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-addon** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-bitrise** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/mozillavpn-3-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/translations-1-beetmover** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-k8s/xpi-t-signing** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/adhoc-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/comm-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-adhoc-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-comm-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-enterprise-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-gecko-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/dep-mozillavpn-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/enterprise-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/gecko-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |
| **scriptworker-prov-v1/mozillavpn-signing-mac14m2** | Scriptworker Chain of Trust | <no value> | 0 | 0 |


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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 59 | 59 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `21`


Count by image:

| Version | Count |
| :--- | ---: |
|  | 8 |
| unknown | 12 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **enterprise-t/win11-64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 110 | 110 |
| **enterprise-t/win11-64-25h2-source-alpha** |  | Version not determined; task not (yet) claimed | 110 | 110 |
| **gecko-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 545 | 545 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **gecko-t/win11-64-25h2-alpha** |  | Version not determined; task not (yet) claimed | 110 | 110 |
| **gecko-t/win11-64-25h2-amd** |  | Version not determined; task not (yet) claimed | 109 | 109 |
| **gecko-t/win11-64-25h2-gpu-alpha** |  | Version not determined; task not (yet) claimed | 109 | 109 |
| **gecko-t/win11-64-25h2-large-alpha** |  | Version not determined; task not (yet) claimed | 108 | 108 |
| **gecko-t/win11-64-25h2-privileged-alpha** |  | Version not determined; task not (yet) claimed | 109 | 109 |
| **gecko-t/win11-64-25h2-source-alpha** |  | Version not determined; task not (yet) claimed | 109 | 109 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** |  | Version not determined; task not (yet) claimed | 114 | 114 |
| **gecko-t/win11-64-25h2-webgpu-alpha** |  | Version not determined; task not (yet) claimed | 109 | 109 |
| **nss-1/win11-a64-25h2-builder-alpha** |  | Version not determined; task not (yet) claimed | 380 | 380 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-arm64** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-linux-netperf-1804** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-linux-talos-2404-bug2031822** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.
