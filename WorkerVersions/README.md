

# Worker Pool Versions


## Generic Worker

Total: `430`

Count by version:

| Version | Count |
| :--- | ---: |
| 100.0.1 | 4 |
| 100.5.0 | 2 |
| 107.0.0 | 229 |
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
| 99.2.0 | 126 |


Count by image:

| Version | Count |
| :--- | ---: |
| unknown | 142 |
|  | 31 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-gui-googlecompute-2024-09-18t05-46-31z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-arm64-gui-googlecompute-2024-09-18t19-02-58z | 2 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha | 6 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2025-01-13t22-33-40z | 2 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-alpha | 4 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-alpha-tc | 7 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-arm64-headless-googlecompute-2026-08-26 | 8 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-gui-googlecompute-2024-08-22t22-48-09z | 10 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-arm64-gui-googlecompute-2024-09-18t14-57-52z | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-gui-googlecompute-2026-08-26 | 5 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-alpha | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-headless-googlecompute-2026-08-26 | 128 |
| projects/fxci-production-level3-workers/global/images/gw-fxci-gcp-l3-2404-amd64-headless-googlecompute-2026-08-26 | 60 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-arm64-headless-googlecompute-2026-08-26 | 16 |


| Worker Pool | Implementation | Version | Engine | Revision | OS | Arch | GO | Total Workers | Total Capacity |
| --- | --- | --- | --- | --- | --- | --- | --- | ---: | ---: |
| **adhoc-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 10 | 10 |
| **adhoc-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 38 | 38 |
| **adhoc-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 5 | 5 |
| **adhoc-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **app-services-1/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 311 | 311 |
| **app-services-1/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **app-services-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 6 | 6 |
| **app-services-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **app-services-3/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 593 | 593 |
| **app-services-3/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **app-services-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 21 | 21 |
| **app-services-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-1/linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-1/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 18 | 18 |
| **code-analysis-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-3/linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **code-analysis-3/linux-gw-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **code-coverage/bot** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 374 | 374 |
| **code-review/bot** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 11061 | 11061 |
| **comm-1/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1662 | 1662 |
| **comm-1/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-1/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 11 | 11 |
| **comm-1/b-linux-docker-large-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **comm-1/b-linux-docker-xlarge-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 4 | 4 |
| **comm-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-1/b-linux-large** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-1/b-linux-xlarge** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **comm-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **comm-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 457 | 457 |
| **comm-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **comm-1/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-docker-large-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-docker-xlarge-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **comm-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-2/b-linux-large** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/b-linux-xlarge** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **comm-2/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-2/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-3/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 5425 | 5425 |
| **comm-3/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-3/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2250 | 2250 |
| **comm-3/b-linux-docker-large-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 9 | 9 |
| **comm-3/b-linux-docker-xlarge-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 6 | 6 |
| **comm-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **comm-3/b-linux-large** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **comm-3/b-linux-xlarge** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **comm-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 759 | 759 |
| **comm-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **comm-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 190 | 190 |
| **comm-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **comm-3/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-t/misc** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 127 | 127 |
| **comm-t/t-linux-arm64-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **comm-t/t-linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 14 | 14 |
| **comm-t/t-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 4382 | 4382 |
| **comm-t/t-linux-docker-noscratch** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 63 | 63 |
| **comm-t/t-linux-docker-noscratch-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 7054 | 7054 |
| **comm-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **comm-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 114 | 114 |
| **comm-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **comm-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 9292 | 9292 |
| **comm-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **comm-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 21 | 21 |
| **enterprise-1/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1030 | 1030 |
| **enterprise-1/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 19 | 19 |
| **enterprise-1/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 173 | 173 |
| **enterprise-1/b-linux-docker-large-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 9 | 9 |
| **enterprise-1/b-linux-docker-xlarge-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1306 | 1306 |
| **enterprise-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 101 | 101 |
| **enterprise-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **enterprise-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 120 | 120 |
| **enterprise-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 5 | 5 |
| **enterprise-1/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 3 | 3 |
| **enterprise-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **enterprise-3/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 4303 | 4303 |
| **enterprise-3/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **enterprise-3/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 768 | 768 |
| **enterprise-3/b-linux-docker-large-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 62 | 62 |
| **enterprise-3/b-linux-docker-xlarge-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 533 | 533 |
| **enterprise-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 610 | 610 |
| **enterprise-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **enterprise-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 62 | 62 |
| **enterprise-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 5 | 5 |
| **enterprise-3/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 3 | 3 |
| **enterprise-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **enterprise-t/misc** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 237 | 237 |
| **enterprise-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 863 | 863 |
| **enterprise-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/t-linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/t-linux-docker-16c32gb-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1078 | 1078 |
| **enterprise-t/t-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 4720 | 4720 |
| **enterprise-t/t-linux-docker-kvm** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **enterprise-t/t-linux-docker-noscratch-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 8105 | 8105 |
| **enterprise-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **enterprise-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 604 | 604 |
| **enterprise-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 7147 | 7147 |
| **enterprise-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 413 | 413 |
| **enterprise-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3040 | 3040 |
| **enterprise-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **enterprise-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 372 | 372 |
| **enterprise-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **enterprise-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 7750 | 7750 |
| **gecko-1/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 47 | 47 |
| **gecko-1/b-linux-docker-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-1/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 11169 | 11169 |
| **gecko-1/b-linux-docker-large-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1214 | 1214 |
| **gecko-1/b-linux-docker-updatebot-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-1/b-linux-docker-xlarge-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1424 | 1424 |
| **gecko-1/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-1/b-linux-kvm** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 166 | 166 |
| **gecko-1/b-linux-large** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-1/b-linux-medium** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 19586 | 19586 |
| **gecko-1/b-linux-xlarge** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1185 | 1185 |
| **gecko-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-1/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-1/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 11 | 11 |
| **gecko-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 273 | 273 |
| **gecko-1/b-win2025-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 177 | 177 |
| **gecko-1/b-win2025-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 28181 | 28181 |
| **gecko-1/b-win2025-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **gecko-1/b-win2025-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **gecko-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1842 | 1842 |
| **gecko-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 76 | 76 |
| **gecko-1/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 16 | 16 |
| **gecko-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 33 | 33 |
| **gecko-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **gecko-2/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 185 | 185 |
| **gecko-2/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **gecko-2/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 307 | 307 |
| **gecko-2/b-linux-docker-large-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 39 | 39 |
| **gecko-2/b-linux-docker-updatebot-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-2/b-linux-docker-xlarge-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 75 | 75 |
| **gecko-2/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-2/b-linux-kvm** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 15 | 15 |
| **gecko-2/b-linux-large** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-2/b-linux-medium** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 328 | 328 |
| **gecko-2/b-linux-xlarge** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-2/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 23 | 23 |
| **gecko-2/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-2/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **gecko-2/b-win2025-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **gecko-2/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 20 | 20 |
| **gecko-2/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-2/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **gecko-3/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 26804 | 26804 |
| **gecko-3/b-linux-2204-kvm-gcp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 45 | 45 |
| **gecko-3/b-linux-docker-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-3/b-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 28537 | 28537 |
| **gecko-3/b-linux-docker-large-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2850 | 2850 |
| **gecko-3/b-linux-docker-updatebot-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 46 | 46 |
| **gecko-3/b-linux-docker-xlarge-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3756 | 3756 |
| **gecko-3/b-linux-gcp-aarch64** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-3/b-linux-kvm** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 630 | 630 |
| **gecko-3/b-linux-large** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 45 | 45 |
| **gecko-3/b-linux-medium** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 32943 | 32943 |
| **gecko-3/b-linux-xlarge** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 192 | 192 |
| **gecko-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 8035 | 8035 |
| **gecko-3/b-win2022-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-3/b-win2022-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 12 | 12 |
| **gecko-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 246 | 246 |
| **gecko-3/b-win2025-headless** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **gecko-3/b-win2025-updatebot** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 221 | 221 |
| **gecko-3/b-win2025-xxlarge** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 220 | 220 |
| **gecko-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 4773 | 4773 |
| **gecko-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 70 | 70 |
| **gecko-3/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 10 | 10 |
| **gecko-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 236 | 236 |
| **gecko-t/misc** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1386 | 1386 |
| **gecko-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 11285 | 11285 |
| **gecko-t/t-linux-2204-wayland-arm64-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | arm64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-root-exp** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **gecko-t/t-linux-2204-wayland-snap** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 58 | 58 |
| **gecko-t/t-linux-2404-headless-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-arm64-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-headless-ssd-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-2404-wayland-snap** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 57 | 57 |
| **gecko-t/t-linux-2404-wayland-snap-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/t-linux-arm64-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 37 | 37 |
| **gecko-t/t-linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1344 | 1344 |
| **gecko-t/t-linux-docker-16c32gb-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2605 | 2605 |
| **gecko-t/t-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 45054 | 45054 |
| **gecko-t/t-linux-docker-amd-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **gecko-t/t-linux-docker-kvm** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 85557 | 85557 |
| **gecko-t/t-linux-docker-kvm-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 21 | 21 |
| **gecko-t/t-linux-docker-noscratch** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 14539 | 14539 |
| **gecko-t/t-linux-docker-noscratch-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 237214 | 237214 |
| **gecko-t/t-linux-docker-noscratch-amd-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 70 | 70 |
| **gecko-t/t-linux-xlarge-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 7581 | 7581 |
| **gecko-t/t-linux-xlarge-2404-wayland** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **gecko-t/win10-64-2009** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6499 | 6499 |
| **gecko-t/win10-64-2009-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 6 | 6 |
| **gecko-t/win10-64-2009-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 1147 | 1147 |
| **gecko-t/win10-64-2009-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 8 | 8 |
| **gecko-t/win10-64-2009-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 108 | 108 |
| **gecko-t/win10-64-2009-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 12 | 12 |
| **gecko-t/win10-64-2009-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win10-64-2009-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-2009** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3549 | 3549 |
| **gecko-t/win11-64-2009-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-2009-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 406 | 406 |
| **gecko-t/win11-64-2009-gpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-source** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 31 | 31 |
| **gecko-t/win11-64-2009-source-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-ssd-gpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 2 | 2 |
| **gecko-t/win11-64-2009-webgpu** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 4 | 4 |
| **gecko-t/win11-64-2009-webgpu-alpha** | generic-worker | 96.2.3 | multiuser | 7e335dd50c | windows | amd64 | 1.25.7 | 3 | 3 |
| **gecko-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4060 | 4060 |
| **gecko-t/win11-64-24h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 23 | 23 |
| **gecko-t/win11-64-24h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 359 | 359 |
| **gecko-t/win11-64-24h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-t/win11-64-24h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 5 | 5 |
| **gecko-t/win11-64-24h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 18 | 18 |
| **gecko-t/win11-64-24h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-24h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-24h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-24h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 119221 | 119221 |
| **gecko-t/win11-64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-amd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-gpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4 | 4 |
| **gecko-t/win11-64-25h2-gpu-perf-experiment** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-large** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 15172 | 15172 |
| **gecko-t/win11-64-25h2-large-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-privileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 439 | 439 |
| **gecko-t/win11-64-25h2-privileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-source** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2517 | 2517 |
| **gecko-t/win11-64-25h2-source-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-25h2-ssd** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-ssd-gpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **gecko-t/win11-64-25h2-unprivileged-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-64-25h2-webgpu** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 4725 | 4725 |
| **gecko-t/win11-64-25h2-webgpu-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **gecko-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 350 | 350 |
| **gecko-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 6 | 6 |
| **glean-1/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 40 | 40 |
| **glean-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **glean-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **glean-3/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 21 | 21 |
| **glean-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 7 | 7 |
| **glean-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **infra/build-decision-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mobile-1/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 198 | 198 |
| **mobile-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 11 | 11 |
| **mobile-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mobile-3/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 243 | 243 |
| **mobile-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 154 | 154 |
| **mobile-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-1/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 219 | 219 |
| **mozilla-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 4 | 4 |
| **mozilla-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 10 | 10 |
| **mozilla-3/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozilla-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 219 | 219 |
| **mozilla-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 456 | 456 |
| **mozilla-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 5 | 5 |
| **mozilla-t/pre-commit** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 11 | 11 |
| **mozilla-t/t-linux-2204-wayland** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 3 | 3 |
| **mozilla-t/t-linux-2204-wayland-relsre** | generic-worker | 64.3.0 | multiuser | b66b6614b9 | linux | amd64 | 1.22.2 | 2 | 2 |
| **mozilla-t/t-linux-2404-wayland** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 13 | 13 |
| **mozilla-t/t-linux-2404-wayland-alpha** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 857 | 857 |
| **mozilla-t/t-linux-arm64-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **mozilla-t/t-linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 44 | 44 |
| **mozilla-t/t-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozilla-t/t-linux-docker-noscratch** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 12 | 12 |
| **mozilla-t/t-linux-docker-noscratch-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **mozillavpn-1/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 292 | 292 |
| **mozillavpn-1/b-linux-large** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 244 | 244 |
| **mozillavpn-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 50 | 50 |
| **mozillavpn-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 219 | 219 |
| **mozillavpn-1/b-win2025-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **mozillavpn-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 6 | 6 |
| **mozillavpn-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 32 | 32 |
| **mozillavpn-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **mozillavpn-3/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 138 | 138 |
| **mozillavpn-3/b-linux-large** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 149 | 149 |
| **mozillavpn-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 27 | 27 |
| **mozillavpn-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 219 | 219 |
| **mozillavpn-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 17 | 17 |
| **mozillavpn-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 10 | 10 |
| **mozillavpn-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 3 | 3 |
| **nss-1/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 457 | 457 |
| **nss-1/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 486 | 486 |
| **nss-1/b-win2022-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **nss-1/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 219 | 219 |
| **nss-1/b-win2025-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **nss-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 44 | 44 |
| **nss-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **nss-1/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **nss-1/linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1334 | 1334 |
| **nss-1/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 5 | 5 |
| **nss-1/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-1/win11-a64-25h2-builder-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 9 | 9 |
| **nss-3/b-linux-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 549 | 549 |
| **nss-3/b-win2022** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 435 | 435 |
| **nss-3/b-win2025** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 219 | 219 |
| **nss-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 32 | 32 |
| **nss-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **nss-3/images-aarch64** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2 | 2 |
| **nss-3/linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 1707 | 1707 |
| **nss-3/win11-a64-25h2-builder** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 4 | 4 |
| **nss-t/t-linux-arm64-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | arm64 | 1.27.0 | 2959 | 2959 |
| **nss-t/t-linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3963 | 3963 |
| **nss-t/t-linux-docker-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **nss-t/win11-a64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 2 | 2 |
| **nss-t/win11-a64-25h2-alpha** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | arm64 | 1.26.2 | 17 | 17 |
| **proj-autophone/gecko-t-bitbar-gw-perf-a55** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-p6** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-autophone/gecko-t-bitbar-gw-perf-s24** | generic-worker | 36.0.0 | simple | b9cb11293f | linux | amd64 | 1.13.7 | 0 | 0 |
| **proj-fuzzing/bugmon-processor-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/ci-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **proj-fuzzing/grizzly-reduce-worker-windows-ngpu** | generic-worker | 100.0.1 | multiuser | fcf9d8ed58 | windows | amd64 | 1.26.2 | 2 | 2 |
| **releng-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 48 | 48 |
| **releng-1/linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 7 | 7 |
| **releng-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 27 | 27 |
| **releng-3/linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 26 | 26 |
| **releng-hardware/applicationservices-b-1-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/applicationservices-b-3-osx1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-1-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/enterprise-3-b-osx-arm64** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | arm64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
| **releng-hardware/gecko-1-b-osx-1015-staging** | generic-worker | 60.3.4 | multiuser | 943a6f2b0d | darwin | amd64 | 1.22.0 | 0 | 0 |
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
| **releng-hardware/win11-64-24h2-hw** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-perf-debug** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref-alpha** | generic-worker | 96.2.2 | multiuser | 773509947e | windows | amd64 | 1.25.7 | 0 | 0 |
| **releng-t/linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 101 | 101 |
| **relops-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **relops-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **relops-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 28 | 28 |
| **relops-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **reviewer-assignment/bot** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 7 | 7 |
| **scriptworker-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 6 | 6 |
| **scriptworker-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 32 | 32 |
| **scriptworker-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 8 | 8 |
| **scriptworker-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 52 | 52 |
| **taskgraph-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 10 | 10 |
| **taskgraph-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **taskgraph-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **taskgraph-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **taskgraph-t/linux-docker** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 12 | 12 |
| **taskgraph-t/win11-64-24h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 3 | 3 |
| **taskgraph-t/win11-64-25h2** | generic-worker | 99.2.0 | multiuser | bc2efe2827 | windows | amd64 | 1.26.2 | 2 | 2 |
| **translations-1/b-linux-large-gcp-1tb-32-256-d2g** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 9 | 9 |
| **translations-1/b-linux-large-gcp-1tb-32-256-std-d2g** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 17 | 17 |
| **translations-1/b-linux-large-gcp-1tb-64-512-d2g** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 5 | 5 |
| **translations-1/b-linux-large-gcp-1tb-64-512-std-d2g** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 97 | 97 |
| **translations-1/b-linux-large-gcp-d2g-1tb** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-1tb-standard** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-large-gcp-d2g-300gb** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 296 | 296 |
| **translations-1/b-linux-v100-gpu-d2g** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 241 | 241 |
| **translations-1/b-linux-v100-gpu-d2g-4** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 142 | 142 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 43 | 43 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-standard** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-1tb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 11 | 11 |
| **translations-1/b-linux-v100-gpu-d2g-4-2tb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 58 | 58 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-standard** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **translations-1/b-linux-v100-gpu-d2g-4-300gb-std-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-4-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/b-linux-v100-gpu-d2g-alpha** | generic-worker | 84.1.2 | multiuser | 4a139f395a | linux | amd64 | 1.24.4 | 2 | 2 |
| **translations-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 15 | 15 |
| **translations-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **translations-t/t-linux-gw-noscratch-amd** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |
| **xpi-1/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 105 | 105 |
| **xpi-1/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 36 | 36 |
| **xpi-1/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 10 | 10 |
| **xpi-3/b-linux** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 3 | 3 |
| **xpi-3/decision** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 4 | 4 |
| **xpi-3/images** | generic-worker | 107.0.0 | multiuser | 6b2f6f6818 | linux | amd64 | 1.27.0 | 2 | 2 |


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
| **infra/build-decision** | docker-worker | 38.0.5 | 5 | 40 |
| **proj-fuzzing/bugmon-pernosco** | docker-worker | 44.23.4 | 2 | 2 |


## Script Worker

Total: `44`



| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **scriptworker-k8s/adhoc-t-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-beetmover** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/app-services-3-shipit** | Scriptworker | <no value> | 0 | 0 |
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
| **scriptworker-k8s/mobile-3-github** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-pushapk** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-shipit** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-signing** | Scriptworker | <no value> | 0 | 0 |
| **scriptworker-k8s/mobile-3-tree** | Scriptworker | <no value> | 0 | 0 |
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
| **gecko-t/t-linux-vm-2204-wayland** |  | No artifacts found | 74 | 74 |
| **gecko-t/t-linux-vm-2204-wayland-snap** |  | No artifacts found | 2 | 2 |


## Version not determined [^2]

Total: `16`


Count by image:

| Version | Count |
| :--- | ---: |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-arm64-mqscbutmsbawtytlroxk | 1 |
| ami-02619e55246806e8d | 1 |
| projects/taskcluster-imaging/global/images/gw-fxci-gcp-l1-2404-amd64-googlecompute-alpha | 1 |
|  | 5 |
| projects/taskcluster-imaging/global/images/generic-worker-ubuntu-24-04-qioehfdraiiishpuvgvl | 7 |
| unknown | 1 |


| Worker Pool | Implementation | Version | Total Workers | Total Capacity |
| --- | --- | --- | ---: | ---: |
| **gecko-t/t-linux-2404-relsre** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **gecko-t/win11-64-25h2-gpu** |  | Version not determined; task not (yet) claimed | 24905 | 24905 |
| **nope/nope** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-autophone/gecko-t-lambda-perf-a55** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **proj-fuzzing/bugmon-monitor** |  | Version not determined; task not (yet) claimed | 89 | 89 |
| **proj-fuzzing/bugmon-pernosco-staging** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/bugmon-processor** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/ci** |  | Version not determined; task not (yet) claimed | 68 | 68 |
| **proj-fuzzing/ci-arm64** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/decision** |  | Version not determined; task not (yet) claimed | 224 | 224 |
| **proj-fuzzing/grizzly-reduce-worker** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/grizzly-reduce-worker-android** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **proj-fuzzing/nss-corpus-update-worker** |  | Version not determined; task not (yet) claimed | 3 | 3 |
| **releng-hardware/gecko-1-b-osx-arm64** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/nss-3-b-osx-1015** |  | Version not determined; task not (yet) claimed | 0 | 0 |
| **releng-hardware/win11-64-24h2-hw-ref** |  | Version not determined; task not (yet) claimed | 0 | 0 |



[^1]: Those are the pools whose tasks were claimed and resolved by a worker as expected, but the worker did not publish either artifact `public/logs/live_backing.log` nor `public/logs/chain_of_trust.log`, which is the source used to identify the worker implementation.

[^2]: Probing task remains pending after two hours. Those are the pools that were not able to start any worker to claim the task within two hours.
