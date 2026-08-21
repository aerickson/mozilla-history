

# Worker Pool Versions


## Generic Worker

Total: `410`

Count by version:

| Version | Count |
| :--- | ---: |
| 100.0.1 | 4 |
| 100.5.0 | 2 |
| 105.0.0 | 229 |
| 36.0.0 | 3 |
| 45.0.0 | 1 |
| 60.3.4 | 14 |
| 61.0.0 | 1 |
| 64.3.0 | 21 |
| 65.1.0 | 1 |
| 67.1.0 | 1 |
| 84.1.2 | 7 |
| 88.0.2 | 3 |
| 91.0.2 | 1 |
| 96.2.2 | 4 |
| 96.2.3 | 12 |
| 99.2.0 | 106 |


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| unknown | 122 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-08-18 | 8 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-08-18 | 128 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-08-18 | 16 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
|  | 31 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-08-18 | 60 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-08-18 | 5 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 4 | 4 |
| **adhoc-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 14 | 14 |
| **adhoc-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 5 | 5 |
| **adhoc-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 4 | 4 |
| **app-services-1/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 338 | 338 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **app-services-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 8 | 8 |
| **app-services-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **app-services-3/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1004 | 1004 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 30 | 30 |
| **app-services-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 22 | 22 |
| **code-analysis-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-coverage/bot** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 287 | 287 |
| **code-review/bot** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 12707 | 12707 |
| **comm-1/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1341 | 1341 |
| **comm-1/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 6 | 6 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **comm-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 376 | 376 |
| **comm-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **comm-1/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-2/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-2/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 6013 | 6013 |
| **comm-3/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2636 | 2636 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 8 | 8 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 4 | 4 |
| **comm-3/b-linux-large** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **comm-3/b-linux-xlarge** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **comm-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 839 | 839 |
| **comm-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 201 | 201 |
| **comm-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **comm-3/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **comm-t/misc** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 16 | 16 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 9 | 9 |
| **comm-t/t-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3603 | 3603 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 68 | 68 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 5180 | 5180 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 126 | 126 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **comm-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7246 | 7246 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 5 | 5 |
| **enterprise-1/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1459 | 1459 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 116 | 116 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 181 | 181 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 14 | 14 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1199 | 1199 |
| **enterprise-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 107 | 107 |
| **enterprise-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 84 | 84 |
| **enterprise-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **enterprise-1/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 4445 | 4445 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 12 | 12 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 779 | 779 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 50 | 50 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 485 | 485 |
| **enterprise-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 607 | 607 |
| **enterprise-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 59 | 59 |
| **enterprise-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **enterprise-3/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **enterprise-t/misc** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 225 | 225 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 757 | 757 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1324 | 1324 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3843 | 3843 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 6371 | 6371 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **enterprise-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 556 | 556 |
| **enterprise-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **enterprise-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5526 | 5526 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 413 | 413 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2779 | 2779 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 255 | 255 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 7433 | 7433 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 49 | 49 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 10698 | 10698 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1113 | 1113 |
| **gecko-1/b-linux-docker-updatebot-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1088 | 1088 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 135 | 135 |
| **gecko-1/b-linux-large** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-1/b-linux-medium** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 19785 | 19785 |
| **gecko-1/b-linux-xlarge** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 896 | 896 |
| **gecko-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 18 | 18 |
| **gecko-1/b-win2025-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 148 | 148 |
| **gecko-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1326 | 1326 |
| **gecko-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 16 | 16 |
| **gecko-1/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 13 | 13 |
| **gecko-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 5 | 5 |
| **gecko-2/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 57 | 57 |
| **gecko-2/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 78 | 78 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 22 | 22 |
| **gecko-2/b-linux-docker-updatebot-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 22 | 22 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 9 | 9 |
| **gecko-2/b-linux-large** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 102 | 102 |
| **gecko-2/b-linux-xlarge** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7 | 7 |
| **gecko-2/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-2/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 16 | 16 |
| **gecko-2/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 24436 | 24436 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 23 | 23 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 25439 | 25439 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2538 | 2538 |
| **gecko-3/b-linux-docker-updatebot-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 38 | 38 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3748 | 3748 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 617 | 617 |
| **gecko-3/b-linux-large** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 11 | 11 |
| **gecko-3/b-linux-medium** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 29017 | 29017 |
| **gecko-3/b-linux-xlarge** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 161 | 161 |
| **gecko-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6013 | 6013 |
| **gecko-3/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 13 | 13 |
| **gecko-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 6314 | 6314 |
| **gecko-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 27 | 27 |
| **gecko-3/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 6 | 6 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 181 | 181 |
| **gecko-t/misc** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 949 | 949 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 8496 | 8496 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 35 | 35 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 35 | 35 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 29 | 29 |
| **gecko-t/t-linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1226 | 1226 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3501 | 3501 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 48361 | 48361 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 85143 | 85143 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 21 | 21 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 11451 | 11451 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 184174 | 184174 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 61 | 61 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 5808 | 5808 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6277 | 6277 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 34 | 34 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 790 | 790 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 82 | 82 |
| **gecko-t/win10-64-2009-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 50 | 50 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10 | 10 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6 | 6 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 1924 | 1924 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 397 | 397 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 14 | 14 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 6 | 6 |
| **gecko-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3556 | 3556 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 311 | 311 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 10 | 10 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 17 | 17 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 129899 | 129899 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1024 | 1024 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 23134 | 23134 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6 | 6 |
| **gecko-t/win11-64-25h2-gpu-perf-experiment** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 14855 | 14855 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 249 | 249 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2534 | 2534 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 326 | 326 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 16 | 16 |
| **glean-1/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 78 | 78 |
| **glean-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 5 | 5 |
| **glean-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 65 | 65 |
| **glean-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 18 | 18 |
| **glean-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 276 | 276 |
| **mobile-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 10 | 10 |
| **mobile-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 308 | 308 |
| **mobile-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 181 | 181 |
| **mobile-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 16 | 16 |
| **mozilla-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 38 | 38 |
| **mozilla-3/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 451 | 451 |
| **mozilla-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 14 | 14 |
| **mozilla-t/pre-commit** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 9 | 9 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 28 | 28 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 564 | 564 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 141 | 141 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 14 | 14 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 198 | 198 |
| **mozillavpn-1/b-linux-large** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 170 | 170 |
| **mozillavpn-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 52 | 52 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-1/b-win2025-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1 | 1 |
| **mozillavpn-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 4 | 4 |
| **mozillavpn-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 24 | 24 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **mozillavpn-3/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 160 | 160 |
| **mozillavpn-3/b-linux-large** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 163 | 163 |
| **mozillavpn-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 36 | 36 |
| **mozillavpn-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 20 | 20 |
| **mozillavpn-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 4 | 4 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **nss-1/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 165 | 165 |
| **nss-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 264 | 264 |
| **nss-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **nss-1/b-win2025-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1 | 1 |
| **nss-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 29 | 29 |
| **nss-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **nss-1/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 676 | 676 |
| **nss-1/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 8 | 8 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-3/b-linux-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 323 | 323 |
| **nss-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 280 | 280 |
| **nss-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 24 | 24 |
| **nss-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **nss-3/images-aarch64** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 1233 | 1233 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | arm64 | 1.26.4 | 1744 | 1744 |
| **nss-t/t-linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2393 | 2393 |
| **nss-t/t-linux-docker-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **nss-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 9 | 9 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-fuzzing/bugmon-processor-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/ci-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows-ngpu** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **releng-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 32 | 32 |
| **releng-1/linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 9 | 9 |
| **releng-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 20 | 20 |
| **releng-3/linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 32 | 32 |
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
| **releng-hardware/gecko-t-win7-32-hw** | generic-worker | 45.0.0 | multiuser | 988e8100b3 | windows | 386 | 1.19.3 | 0 | 0 |
| **releng-hardware/mozillavpn-b-1-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/mozillavpn-b-3-osx** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 71 | 71 |
| **relops-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **relops-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **relops-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 23 | 23 |
| **relops-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **reviewer-assignment/bot** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **scriptworker-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 4 | 4 |
| **scriptworker-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 16 | 16 |
| **scriptworker-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 7 | 7 |
| **scriptworker-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 45 | 45 |
| **taskgraph-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 7 | 7 |
| **taskgraph-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 7 | 7 |
| **taskgraph-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **taskgraph-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **taskgraph-t/linux-docker** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 33 | 33 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 21 | 21 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 28 | 28 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 4 | 4 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 139 | 139 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 5 | 5 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 503 | 503 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 331 | 331 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 153 | 153 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 70 | 70 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 43 | 43 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 80 | 80 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 22 | 22 |
| **translations-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 8 | 8 |
| **xpi-1/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 47 | 47 |
| **xpi-1/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 24 | 24 |
| **xpi-1/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |
| **xpi-3/b-linux** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 4 | 4 |
| **xpi-3/decision** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 6 | 6 |
| **xpi-3/images** | generic-worker | 105.0.0 | multiuser | 66ac7d3ec5 | linux | amd64 | 1.26.4 | 2 | 2 |


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
| **infra/build-decision** | docker-worker | 38.0.5 | 11 | 88 |
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
| **scriptworker-k8s/comm-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/comm-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/enterprise-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-1-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-balrog** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-bouncer** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-lando** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/gecko-t-signing-dev** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/glean-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/glean-3-signing** | Scriptworker | <no value> | 0 | 0 |
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
|  | 2 |
| projects/taskcluster-imaging/global/images/generic-2204-wayland-vm-gcp-googlecompute-2023-09-22t17-39-37z | 2 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **built-in/fail** |  | No artifacts found | 0 | 0 |
| **built-in/succeed** |  | No artifacts found | 0 | 0 |
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 99 | 99 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `36`


Count by image:

| Version | Count |
| :--- | ---: |
| unknown | 21 |
|  | 5 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-qioehfdraiiishpuvgvl | 7 |
| ami-02619e55246806e8d | 1 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-arm64-mqscbutmsbawtytlroxk | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **codex-empty-auth-20260817/probe-a** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **comm-1/b-win2025** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **comm-3/b-win2025** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **enterprise-1/b-win2025** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **enterprise-3/b-win2025** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-1/b-win2025** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-1/b-win2025-headless** |  | Version not determined; task not (yet) claimed | 20755 | 20755 |
| **gecko-1/b-win2025-updatebot** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-1/b-win2025-xxlarge** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-2/b-win2025** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-2/b-win2025-updatebot** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-3/b-win2025** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-3/b-win2025-headless** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-3/b-win2025-updatebot** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-3/b-win2025-xxlarge** |  | Version not determined; task not (yet) claimed | 105 | 105 |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **gecko-t/win11-64-25h2-webgpu** |  | Version not determined; task not (yet) claimed | 4728 | 4728 |
| **mozilla-1/b-win2025** |  | Version not determined; task not (yet) claimed | 104 | 104 |
| **mozilla-3/b-win2025** |  | Version not determined; task not (yet) claimed | 104 | 104 |
| **mozillavpn-1/b-win2025** |  | Version not determined; task not (yet) claimed | 104 | 104 |
| **mozillavpn-3/b-win2025** |  | Version not determined; task not (yet) claimed | 104 | 104 |
| **nss-1/b-win2025** |  | Version not determined; task not (yet) claimed | 104 | 104 |
| **nss-3/b-win2025** |  | Version not determined; task not (yet) claimed | 104 | 104 |
| **proj-autophone/gecko-t-lambda-perf-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-fuzzing/bugmon-monitor** |  | Version not determined; task not (yet) claimed | 90 | 90 |
| **proj-fuzzing/bugmon-pernosco-staging** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/bugmon-processor** |  | Version not determined; task not (yet) claimed | 4 | 4 |
| **proj-fuzzing/ci** |  | Version not determined; task not (yet) claimed | 65 | 65 |
| **proj-fuzzing/ci-arm64** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/decision** |  | Version not determined; task not (yet) claimed | 213 | 213 |
| **proj-fuzzing/grizzly-reduce-worker** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/grizzly-reduce-worker-android** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/nss-corpus-update-worker** |  | Version not determined; task not (yet) claimed | 7 | 7 |
| **releng-hardware/nss-1-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.
