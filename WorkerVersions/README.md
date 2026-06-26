

# Worker Pool Versions


## Generic Worker

Total: `404`

Count by version:

| Version | Count |
| :--- | ---: |
| 100.0.1 | 4 |
| 100.4.0 | 12 |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 18 |
| 61.0.0 | 1 |
| 64.3.0 | 21 |
| 65.1.0 | 1 |
| 84.1.2 | 7 |
| 87.0.0 | 1 |
| 88.0.2 | 2 |
| 95.0.3 | 1 |
| 96.2.2 | 6 |
| 96.2.3 | 12 |
| 99.2.0 | 101 |
| 99.2.1 | 213 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-05-04 | 16 |
|  | 34 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| unknown | 117 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-05-04 | 59 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-05-04 | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-05-04 | 125 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-05-04 | 8 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **adhoc-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **adhoc-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **adhoc-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 415 | 415 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **app-services-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1108 | 1108 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 30 | 30 |
| **app-services-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 26 | 26 |
| **code-analysis-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 4 | 4 |
| **code-coverage/bot** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 267 | 267 |
| **code-review/bot** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 9832 | 9832 |
| **comm-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 936 | 936 |
| **comm-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 262 | 262 |
| **comm-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **comm-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 3 | 3 |
| **comm-2/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-2/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **comm-2/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3929 | 3929 |
| **comm-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 3 | 3 |
| **comm-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1887 | 1887 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **comm-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 520 | 520 |
| **comm-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 120 | 120 |
| **comm-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **comm-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 3 | 3 |
| **comm-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 47 | 47 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2305 | 2305 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3376 | 3376 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4457 | 4457 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **enterprise-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1237 | 1237 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 21 | 21 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 464 | 464 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 27 | 27 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1056 | 1056 |
| **enterprise-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 110 | 110 |
| **enterprise-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 81 | 81 |
| **enterprise-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 91 | 91 |
| **enterprise-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 16 | 16 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **enterprise-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1805 | 1805 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 25 | 25 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 674 | 674 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 73 | 73 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 366 | 366 |
| **enterprise-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 242 | 242 |
| **enterprise-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 41 | 41 |
| **enterprise-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 73 | 73 |
| **enterprise-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 13 | 13 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **enterprise-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 313 | 313 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 599 | 599 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1361 | 1361 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3898 | 3898 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5795 | 5795 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 276 | 276 |
| **enterprise-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5630 | 5630 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 264 | 264 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1967 | 1967 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 187 | 187 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 8289 | 8289 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 234 | 234 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 4 | 4 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10223 | 10223 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 108 | 108 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1068 | 1068 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 170 | 170 |
| **gecko-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 15878 | 15878 |
| **gecko-1/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **gecko-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 725 | 725 |
| **gecko-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 69 | 69 |
| **gecko-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 812 | 812 |
| **gecko-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 107 | 107 |
| **gecko-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 15 | 15 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 18 | 18 |
| **gecko-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 361 | 361 |
| **gecko-2/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 50 | 50 |
| **gecko-2/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 3 | 3 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 78 | 78 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 31 | 31 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23 | 23 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **gecko-2/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 129 | 129 |
| **gecko-2/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-2/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12 | 12 |
| **gecko-2/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23234 | 23234 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 19 | 19 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 24825 | 24825 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 646 | 646 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2746 | 2746 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 486 | 486 |
| **gecko-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **gecko-3/b-linux-medium** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 29290 | 29290 |
| **gecko-3/b-linux-xlarge** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 151 | 151 |
| **gecko-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5599 | 5599 |
| **gecko-3/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 12 | 12 |
| **gecko-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4469 | 4469 |
| **gecko-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 120 | 120 |
| **gecko-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 19 | 19 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 115 | 115 |
| **gecko-t/misc** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 716 | 716 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 7569 | 7569 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | arm64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 29 | 29 |
| **gecko-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 1139 | 1139 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3107 | 3107 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 36791 | 36791 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 3 | 3 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 65985 | 65985 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 21 | 21 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10024 | 10024 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 145288 | 145288 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 61 | 61 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 5550 | 5550 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3396 | 3396 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 34 | 34 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 437 | 437 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-t/win10-64-2009-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 96 | 96 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6 | 6 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10 | 10 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1714 | 1714 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 302 | 302 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 38 | 38 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2647 | 2647 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 113 | 113 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10 | 10 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 229 | 229 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6 | 6 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 20 | 20 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 15 | 15 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7 | 7 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 75244 | 75244 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 118 | 118 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 16240 | 16240 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 13 | 13 |
| **gecko-t/win11-64-25h2-gpu-perf-experiment** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10111 | 10111 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 30 | 30 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 257 | 257 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7 | 7 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1875 | 1875 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7 | 7 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3794 | 3794 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 262 | 262 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 16 | 16 |
| **glean-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 46 | 46 |
| **glean-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
| **glean-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 25 | 25 |
| **glean-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 8 | 8 |
| **glean-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 153 | 153 |
| **mobile-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **mobile-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 267 | 267 |
| **mobile-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 147 | 147 |
| **mobile-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **mozilla-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **mozilla-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 447 | 447 |
| **mozilla-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **mozilla-t/pre-commit** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 100.4.0 | multiuser | d87e9cc5ac | linux | amd64 | 1.26.4 | 605 | 605 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 23 | 23 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 392 | 392 |
| **mozillavpn-1/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 305 | 305 |
| **mozillavpn-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 83 | 83 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **mozillavpn-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 3 | 3 |
| **mozillavpn-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **mozillavpn-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 178 | 178 |
| **mozillavpn-3/b-linux-large** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 172 | 172 |
| **mozillavpn-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 29 | 29 |
| **mozillavpn-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 20 | 20 |
| **mozillavpn-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 18 | 18 |
| **nss-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 56 | 56 |
| **nss-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **nss-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **nss-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-1/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 251 | 251 |
| **nss-1/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 18 | 18 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 205 | 205 |
| **nss-3/b-linux-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 3 | 3 |
| **nss-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 102 | 102 |
| **nss-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **nss-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-3/images-aarch64** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 306 | 306 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | arm64 | 1.26.2 | 45 | 45 |
| **nss-t/t-linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 616 | 616 |
| **nss-t/t-linux-docker-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **nss-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 17 | 17 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** | generic-worker | 87.0.0 | insecure | 99a1fcafbb | linux | amd64 | 1.24.5 | 0 | 0 |
| **proj-fuzzing/bugmon-processor-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 178 | 178 |
| **proj-fuzzing/ci-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 111 | 111 |
| **proj-fuzzing/grizzly-reduce-worker-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 107 | 107 |
| **proj-fuzzing/grizzly-reduce-worker-windows-ngpu** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 111 | 111 |
| **releng-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 52 | 52 |
| **releng-1/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 35 | 35 |
| **releng-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 36 | 36 |
| **releng-3/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 50 | 50 |
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
| **releng-hardware/gecko-t-osx-1015-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1400-r8-staging** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m-vms** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4** | generic-worker | 95.0.3 | multiuser | a996b8269b | darwin | arm64 | 1.25.5 | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-ipv6** | generic-worker | 60.3.4 | simple | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-relops1213** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 269 | 269 |
| **relops-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **relops-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **relops-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 13 | 13 |
| **relops-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **scriptworker-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12 | 12 |
| **scriptworker-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 79 | 79 |
| **scriptworker-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **scriptworker-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 71 | 71 |
| **taskgraph-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 12 | 12 |
| **taskgraph-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **taskgraph-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **taskgraph-3/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 5 | 5 |
| **taskgraph-t/linux-docker** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 118 | 118 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7 | 7 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 6 | 6 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 29 | 29 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 48 | 48 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 184 | 184 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 43 | 43 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 59 | 59 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 11 | 11 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 4 | 4 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 14 | 14 |
| **translations-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 163 | 163 |
| **xpi-1/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 75 | 75 |
| **xpi-1/images** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 7 | 7 |
| **xpi-3/b-linux** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 24 | 24 |
| **xpi-3/decision** | generic-worker | 99.2.1 | multiuser | ddb9ce7efd | linux | amd64 | 1.26.2 | 10 | 10 |
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
| ami-03e4f8db63254ce7e,ami-0a6e926238859761c,ami-0b5dd0bbb670ec80e | 1 |
| projects/taskcluster-imaging/global/images/docker-firefoxci-gcp-l1-googlecompute-2025-06-13t18-31-38z | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **infra/build-decision** | docker-worker | 38.0.5 | 4 | 32 |
| **proj-fuzzing/bugmon-pernosco** | docker-worker | 44.23.4 | 2 | 2 |


## Script Worker

Total: `42`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
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
| **scriptworker-k8s/gecko-1-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushflatpak** | Scriptworker | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 77 | 77 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `12`


Count by image:

| Version | Count |
| :--- | ---: |
| ami-02619e55246806e8d | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-arm64-mqscbutmsbawtytlroxk | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-qioehfdraiiishpuvgvl | 7 |
|  | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/bugmon-monitor** |  | Version not determined; task not (yet) claimed | 87 | 87 |
| **proj-fuzzing/bugmon-pernosco-staging** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/bugmon-processor** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/ci** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-fuzzing/ci-arm64** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/decision** |  | Version not determined; task not (yet) claimed | 199 | 199 |
| **proj-fuzzing/grizzly-reduce-worker** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **proj-fuzzing/grizzly-reduce-worker-android** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **proj-fuzzing/nss-corpus-update-worker** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **releng-hardware/gecko-t-osx-1015-r8** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/gecko-t-osx-1500-m4-staging** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.
