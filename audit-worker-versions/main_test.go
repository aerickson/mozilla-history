package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/taskcluster/taskcluster/v48/clients/client-go/tcworkermanager"
)

func intPointer(value int) *int {
	return &value
}

func TestEnrichWorkerInfoWithConfiguredCapacity(t *testing.T) {
	tests := []struct {
		name   string
		config string
		slots  int
	}{
		{
			name:   "current nested capacity",
			config: `{"minCapacity":1,"maxCapacity":16,"launchConfigs":[{"workerManager":{"capacityPerInstance":8}}]}`,
			slots:  8,
		},
		{
			name:   "legacy direct capacity",
			config: `{"minCapacity":1,"maxCapacity":16,"launchConfigs":[{"capacityPerInstance":8}]}`,
			slots:  8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			worker := WorkerInfo{}
			pool := tcworkermanager.WorkerPoolFullDefinition{
				Config: json.RawMessage(test.config),
			}

			enrichWorkerInfo(&worker, &pool)

			for name, result := range map[string]struct {
				got  *int
				want int
			}{
				"minimum capacity": {worker.ConfiguredMinCapacity, 1},
				"maximum capacity": {worker.ConfiguredMaxCapacity, 16},
				"slots per worker": {worker.CapacityPerWorker, test.slots},
				"minimum workers":  {worker.ConfiguredMinWorkers, 1},
				"maximum workers":  {worker.ConfiguredMaxWorkers, 2},
			} {
				if result.got == nil || *result.got != result.want {
					t.Errorf("%s = %v, want %d", name, result.got, result.want)
				}
			}

			encoded, err := json.Marshal(worker)
			if err != nil {
				t.Fatal(err)
			}
			for _, want := range []string{
				`"ConfiguredMinCapacity":1`,
				`"ConfiguredMaxCapacity":16`,
				`"CapacityPerWorker":8`,
				`"ConfiguredMinWorkers":1`,
				`"ConfiguredMaxWorkers":2`,
			} {
				if !strings.Contains(string(encoded), want) {
					t.Errorf("serialized worker does not contain %s", want)
				}
			}
			if strings.Contains(string(encoded), "TotalWorkers") || strings.Contains(string(encoded), "TotalCapacity") {
				t.Error("new worker snapshot contains legacy totals")
			}
		})
	}
}

func TestEnrichWorkerInfoDoesNotDeriveWorkersForMixedCapacities(t *testing.T) {
	worker := WorkerInfo{}
	pool := tcworkermanager.WorkerPoolFullDefinition{
		Config: json.RawMessage(`{
			"minCapacity": 0,
			"maxCapacity": 16,
			"launchConfigs": [
				{"workerManager":{"capacityPerInstance":2}},
				{"workerManager":{"capacityPerInstance":8}}
			]
		}`),
	}

	enrichWorkerInfo(&worker, &pool)

	if worker.ConfiguredMinCapacity == nil || worker.ConfiguredMaxCapacity == nil {
		t.Fatal("configured capacity bounds were not preserved")
	}
	if worker.CapacityPerWorker != nil || worker.ConfiguredMinWorkers != nil || worker.ConfiguredMaxWorkers != nil {
		t.Fatal("worker counts were derived for heterogeneous launch capacities")
	}
}
