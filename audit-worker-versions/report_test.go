package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
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

	got := renderReadme(workers)
	for _, want := range []string{
		"This report shows the latest detailed inventory of Firefox CI worker pools alongside historical trends from earlier snapshots.",
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

func TestRenderReadmeExplainsIncompleteProbesInline(t *testing.T) {
	workers := []WorkerInfo{
		{WorkerPoolID: "example/no-artifact", Details: map[string]string{"error": "No artifacts found"}, hasNoArtifacts: true},
		{WorkerPoolID: "example/pending", Details: map[string]string{"error": "Version not determined; task not (yet) claimed"}, isUnknown: true},
	}

	got := renderReadme(workers)
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

	workers, err := readSnapshot(filename)
	if err != nil {
		t.Fatal(err)
	}
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
