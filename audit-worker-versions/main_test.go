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

func TestRenderReadmeIncludesLinksAndSubheadings(t *testing.T) {
	workers := []WorkerInfo{{
		WorkerPoolID:   "example/pool",
		Implementation: "generic-worker",
		Version:        "1.2.3",
		Details:        map[string]string{"revision": "1234567890"},
	}, {
		WorkerPoolID:   "example/other-pool",
		Implementation: "generic-worker",
		Version:        "2.0.0",
		Details:        map[string]string{"revision": "1234567890"},
	}}

	got := renderReadme(workers)
	for _, want := range []string{
		"### Count by version",
		"[**example/pool**](https://firefox-ci-tc.services.mozilla.com/provisioners/example/worker-types/pool?sortBy=Last%20Active&sortDirection=desc)",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("rendered README does not contain %q", want)
		}
	}
}

func TestReadSnapshotRestoresRenderingState(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "workers.json")
	data := `[
		{"WorkerPoolID":"one/pool","Details":{"error":"No artifacts found"}},
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
	if !workers[1].isUnknown {
		t.Error("unknown-version state was not restored")
	}
}
