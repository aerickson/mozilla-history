package main

import (
	"os"
	"path/filepath"
	"testing"
)

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
