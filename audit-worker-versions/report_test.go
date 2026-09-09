package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestWorkerPoolURL(t *testing.T) {
	worker := WorkerInfo{WorkerPoolID: "releng-hardware/gecko-t-win7-32-hw"}
	want := "https://firefox-ci-tc.services.mozilla.com/provisioners/releng-hardware/worker-types/gecko-t-win7-32-hw?sortBy=Last%20Active&sortDirection=desc"
	if got := worker.WorkerPoolURL(); got != want {
		t.Fatalf("WorkerPoolURL() = %q, want %q", got, want)
	}
}

func TestSortedVersionCountsUsesNaturalOrder(t *testing.T) {
	versions := map[string]int{
		"100.0.1": 1,
		"108.0.0": 1,
		"9.10.0":  1,
		"9.2.0":   1,
		"99.2.0":  1,
	}
	want := []string{"9.2.0", "9.10.0", "99.2.0", "100.0.1", "108.0.0"}

	got := sortedVersionCounts(versions)
	for i := range want {
		if got[i].Key != want[i] {
			t.Fatalf("version %d = %q, want %q", i, got[i].Key, want[i])
		}
	}
}

func TestRenderReadmeIncludesLinksAndSubheadings(t *testing.T) {
	workers := []WorkerInfo{{
		WorkerPoolID:          "example/pool",
		Implementation:        "generic-worker",
		Version:               "1.2.3",
		Imageset:              "image-one",
		Details:               map[string]string{"revision": "1234567890"},
		ConfiguredMinCapacity: intPointer(1),
		ConfiguredMaxCapacity: intPointer(16),
		CapacityPerWorker:     intPointer(8),
		ConfiguredMinWorkers:  intPointer(1),
		ConfiguredMaxWorkers:  intPointer(2),
	}, {
		WorkerPoolID:   "example/other-pool",
		Implementation: "generic-worker",
		Version:        "2.0.0",
		Imageset:       "image-two",
		Details:        map[string]string{"revision": "1234567890"},
	}}

	got := renderReadmeAt(WorkerSnapshot{
		GeneratedAt:    time.Date(2026, time.September, 9, 15, 29, 53, 0, time.UTC),
		ProbeStartedAt: time.Date(2026, time.September, 9, 7, 58, 29, 0, time.UTC),
		Workers:        workers,
	}, time.Date(2026, time.September, 10, 1, 2, 3, 0, time.UTC))
	for _, want := range []string{
		"This report shows the latest detailed inventory of Firefox CI worker pools alongside historical trends from earlier snapshots.",
		"Probe run started: **2026-09-09 07:58 UTC** · Results collected: **2026-09-09 15:29 UTC** · Report generated: **2026-09-10 01:02 UTC**",
		"Total worker pools: `2`",
		"### Count by version",
		"intentionally malformed probe task",
		"expected to fail with a malformed-payload exception",
		"live Worker Manager launch configuration",
		"### Worker pools",
		"Configured Workers | Configured Capacity | Slots per Worker",
		"pool's autoscaling range in concurrent task slots",
		"| 1–2 | 1–16 | 8 |",
		"[**example/pool**](https://firefox-ci-tc.services.mozilla.com/provisioners/example/worker-types/pool?sortBy=Last%20Active&sortDirection=desc)",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("rendered README does not contain %q", want)
		}
	}
}

func TestRenderReadmeDistinguishesImageStatuses(t *testing.T) {
	details := map[string]string{"revision": "1234567890"}
	workers := []WorkerInfo{
		{Implementation: "generic-worker", Imageset: "unknown", ImageStatus: imageStatusNotApplicable, Details: details},
		{Implementation: "generic-worker", Imageset: "unknown", ImageStatus: imageStatusUnavailable, Details: details},
		{Implementation: "generic-worker", Imageset: "unknown", ImageStatus: imageStatusNotDetermined, Details: details},
		{Implementation: "generic-worker", Imageset: "azure/image", ImageStatus: imageStatusKnown, Details: details},
	}

	got := renderReadme(WorkerSnapshot{Workers: workers})

	for _, want := range []string{
		"| Image | Count |",
		"| Not applicable (standalone) | 1 |",
		"| Configuration unavailable | 1 |",
		"| Image not determined | 1 |",
		"| azure/image | 1 |",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("rendered README does not contain %q", want)
		}
	}
}

func TestCompactAzureImageReferences(t *testing.T) {
	tests := map[string]string{
		"gallery with redundant image name": "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/galleries/win2022/images/win2022/versions/1.0.0",
		"gallery with distinct image name":  "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/galleries/gallery/images/image/versions/2.0",
		"managed image":                     "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/imageset-worker-eastus",
		"non-Azure image":                   "projects/example/global/images/worker-image",
	}
	want := map[string]string{
		"gallery with redundant image name": "Azure gallery win2022@1.0.0",
		"gallery with distinct image name":  "Azure gallery gallery/image@2.0",
		"managed image":                     "Azure image imageset-worker-eastus",
		"non-Azure image":                   "projects/example/global/images/worker-image",
	}

	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			if got := compactAzureImageReference(input); got != want[name] {
				t.Fatalf("compactAzureImageReference() = %q, want %q", got, want[name])
			}
		})
	}
}

func TestCompactImageReferencesPreservesConfiguredSet(t *testing.T) {
	input := "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/one,/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/two"
	want := "Azure image one, Azure image two"
	if got := compactImageReferences(input); got != want {
		t.Fatalf("compactImageReferences() = %q, want %q", got, want)
	}
}

func TestCompactImageReferencesSummarizesRegionalAzureSet(t *testing.T) {
	input := strings.Join([]string{
		"/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/imageset-abcdefghijklmnopqrst-westus2-fuzzing",
		"/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/images/imageset-zyxwvutsrqponmlkjihg-eastus-fuzzing",
	}, ",")
	want := "Azure image set fuzzing (eastus, westus2)"
	if got := compactImageReferences(input); got != want {
		t.Fatalf("compactImageReferences() = %q, want %q", got, want)
	}
}

func TestImageHoverTitleSeparatesImageSetWithLineBreaks(t *testing.T) {
	if got := imageHoverTitle("/azure/one,/azure/two"); got != "/azure/one&#10;/azure/two" {
		t.Fatalf("imageHoverTitle() = %q, want encoded line break", got)
	}
}

func TestRenderReadmeShowsFullAzureImageOnHover(t *testing.T) {
	fullImage := "/subscriptions/sub/resourceGroups/rg/providers/Microsoft.Compute/galleries/win2022/images/win2022/versions/1.0.0"
	workers := []WorkerInfo{
		{Implementation: "generic-worker", Imageset: fullImage, ImageStatus: imageStatusKnown, Details: map[string]string{"revision": "1234567890"}},
		{Implementation: "generic-worker", Imageset: "projects/example/global/images/linux", ImageStatus: imageStatusKnown, Details: map[string]string{"revision": "1234567890"}},
	}

	got := renderReadme(WorkerSnapshot{Workers: workers})

	want := `<abbr title="` + fullImage + `">Azure gallery win2022@1.0.0</abbr>`
	if !strings.Contains(got, want) {
		t.Fatalf("rendered README does not contain Azure hover label %q", want)
	}
	if strings.Contains(got, `<abbr title="projects/example`) {
		t.Fatal("non-Azure image unexpectedly received a hover label")
	}
}

func TestRenderReadmeExplainsIncompleteProbesInline(t *testing.T) {
	workers := []WorkerInfo{
		{WorkerPoolID: "example/no-artifact", Details: map[string]string{"error": "No artifacts found"}, hasNoArtifacts: true},
		{WorkerPoolID: "example/pending", Details: map[string]string{"error": "Version not determined; task not (yet) claimed"}, isUnknown: true},
	}

	got := renderReadme(WorkerSnapshot{Workers: workers})
	for _, want := range []string{
		"## No artifacts found\n",
		"did not publish `public/logs/live_backing.log`",
		"## Version not determined\n",
		"did not claim the probe task within two hours",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("rendered README does not contain %q", want)
		}
	}
	if strings.Contains(got, "[^1]") || strings.Contains(got, "[^2]") {
		t.Error("rendered README contains obsolete footnote markers")
	}
}

func TestReadSnapshotRestoresRenderingState(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "workers.json")
	data := `[
		{"WorkerPoolID":"one/pool","Details":{"error":"No artifacts found"},"TotalWorkers":42,"TotalCapacity":84},
		{"WorkerPoolID":"two/pool","Details":{"error":"Version not determined; task not (yet) claimed"}}
	]`
	if err := os.WriteFile(filename, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	snapshot, err := readSnapshot(filename)
	if err != nil {
		t.Fatal(err)
	}
	workers := snapshot.Workers
	if !workers[0].hasNoArtifacts {
		t.Error("no-artifacts state was not restored")
	}
	if workers[0].LegacyTotalWorkers == nil || *workers[0].LegacyTotalWorkers != 42 ||
		workers[0].LegacyTotalCapacity == nil || *workers[0].LegacyTotalCapacity != 84 {
		t.Error("legacy totals were not recognized")
	}
	if !workers[1].isUnknown {
		t.Error("unknown-version state was not restored")
	}
}

func TestReadSnapshotReadsMetadata(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "workers.json")
	data := `{
		"generatedAt":"2026-09-09T15:29:53Z",
		"probeStartedAt":"2026-09-09T07:58:29.422Z",
		"taskGroupId":"AnhEjBL2SYuUedNBvjgsWA",
		"workers":[{"WorkerPoolID":"one/pool","Details":{}}]
	}`
	if err := os.WriteFile(filename, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	snapshot, err := readSnapshot(filename)
	if err != nil {
		t.Fatal(err)
	}
	if got := snapshot.GeneratedAt.Format(time.RFC3339); got != "2026-09-09T15:29:53Z" {
		t.Errorf("generated time = %q", got)
	}
	if got := snapshot.ProbeStartedAt.Format(time.RFC3339Nano); got != "2026-09-09T07:58:29.422Z" {
		t.Errorf("probe start time = %q", got)
	}
	if snapshot.TaskGroupID != "AnhEjBL2SYuUedNBvjgsWA" {
		t.Errorf("task group ID = %q", snapshot.TaskGroupID)
	}
	if len(snapshot.Workers) != 1 || snapshot.Workers[0].WorkerPoolID != "one/pool" {
		t.Errorf("workers = %#v", snapshot.Workers)
	}
}
