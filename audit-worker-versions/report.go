package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type count struct {
	Key   string
	Value int
}

type reportSection struct {
	Title           string
	Description     string
	Count           int
	Versions        []count
	Images          []count
	Filtered        []WorkerInfo
	FullColumns     bool
	HasLegacyTotals bool
}

type reportData struct {
	GeneratedAt    string
	ProbeStartedAt string
	Sections       [5]reportSection
}

func (w WorkerInfo) WorkerPoolURL() string {
	parts := strings.SplitN(w.WorkerPoolID, "/", 2)
	if len(parts) != 2 {
		return ""
	}
	return "https://firefox-ci-tc.services.mozilla.com/provisioners/" +
		url.PathEscape(parts[0]) + "/worker-types/" + url.PathEscape(parts[1]) +
		"?sortBy=Last%20Active&sortDirection=desc"
}

func formatConfiguredRange(minimum, maximum *int) string {
	if minimum == nil || maximum == nil {
		return "—"
	}
	if *minimum == *maximum {
		return fmt.Sprintf("%d", *minimum)
	}
	return fmt.Sprintf("%d–%d", *minimum, *maximum)
}

func (w WorkerInfo) ConfiguredWorkerRange() string {
	return formatConfiguredRange(w.ConfiguredMinWorkers, w.ConfiguredMaxWorkers)
}

func (w WorkerInfo) ConfiguredCapacityRange() string {
	return formatConfiguredRange(w.ConfiguredMinCapacity, w.ConfiguredMaxCapacity)
}

func (w WorkerInfo) CapacityPerWorkerValue() string {
	if w.CapacityPerWorker == nil {
		return "—"
	}
	return fmt.Sprintf("%d", *w.CapacityPerWorker)
}

const readmeTpl = `
{{- define "row" -}}
## {{ .Title }}
{{ if .Description }}
{{ .Description }}
{{ end }}

Total worker pools: ` + "`" + `{{ .Count }}` + "`" + `
{{ if gt (len .Versions) 1 }}
### Count by version

_Source: version information parsed from the log artifact produced when each worker claims an intentionally malformed probe task. The task is expected to fail with a malformed-payload exception; known worker implementations and versions are identified from their distinct log output._

| Version | Count |
| :--- | ---: |
{{ range .Versions -}}
| {{ .Key }} | {{ .Value }} |
{{ end }}
{{- end }}
{{ if gt (len .Images) 1 }}
### Count by image

_Source: image references in each pool's live Worker Manager launch configuration at report time. A value may represent multiple configured images; unknown means no supported image reference was found._

| Version | Count |
| :--- | ---: |
{{ range .Images -}}
| {{ .Key }} | {{ .Value }} |
{{ end }}
{{- end }}
{{if .Count }}
### Worker pools

_Configured capacity is the pool's autoscaling range in concurrent task slots. The configured worker range is derived from those bounds and the slots per worker, rounding up. An em dash means the configuration is unavailable or a worker count cannot be derived._
{{ if .HasLegacyTotals }}
_Configured values were not collected in this snapshot. Legacy totals included stopped worker records and are intentionally not displayed._
{{ end }}

| Worker Pool | Implementation | Version {{ if .FullColumns }}| Engine | Revision | OS | Arch | GO {{ end }}| Configured Workers | Configured Capacity | Slots per Worker |
| --- | --- | --- {{ if .FullColumns }}| --- | --- | --- | --- | --- {{ end }}| ---: | ---: | ---: |
{{ range .Filtered -}}
| [**{{ .WorkerPoolID }}**]({{ .WorkerPoolURL }}) | {{ .Implementation }} | {{ or .Version .Details.error }} {{ if $.FullColumns }}| {{ or .Details.engine "-" }} | {{ or (slice .Details.revision 0 10) "-" }} | {{ or .Details.os "-" }} | {{ or .Details.arch "-" }} | {{ or .Details.go "-" }} {{ end }}| {{ .ConfiguredWorkerRange }} | {{ .ConfiguredCapacityRange }} | {{ .CapacityPerWorkerValue }} |
{{end}}
{{- end -}}
{{end}}

# Worker Pool Versions

This report shows the latest detailed inventory of Firefox CI worker pools alongside historical trends from earlier snapshots. Worker implementation and version are inferred from the failure log produced when each pool is given an intentionally malformed probe task; image and capacity metadata come from Worker Manager. Summary counts represent worker pools, not individual workers or tasks.

{{ if .ProbeStartedAt }}Probe run started: **{{ .ProbeStartedAt }}**{{ if .GeneratedAt }} · Results collected: **{{ .GeneratedAt }}**{{ end }}
{{ else if .GeneratedAt }}Results collected: **{{ .GeneratedAt }}**
{{ end }}

{{ range .Sections }}
{{ template "row" . }}
{{ end }}
`

func renderTemplate(data interface{}) string {
	t := template.Must(template.New("").Parse(readmeTpl))
	var content bytes.Buffer
	if err := t.Execute(&content, data); err != nil {
		panic(err)
	}
	return strings.TrimSpace(content.String()) + "\n"
}

func sortedCounts(values map[string]int) []count {
	counts := make([]count, 0, len(values))
	for key, value := range values {
		counts = append(counts, count{Key: key, Value: value})
	}
	sort.Slice(counts, func(i, j int) bool {
		return strings.Compare(counts[i].Key, counts[j].Key) < 0
	})
	return counts
}

func naturalLess(left, right string) bool {
	for leftIndex, rightIndex := 0, 0; leftIndex < len(left) && rightIndex < len(right); {
		leftDigit := left[leftIndex] >= '0' && left[leftIndex] <= '9'
		rightDigit := right[rightIndex] >= '0' && right[rightIndex] <= '9'
		if leftDigit && rightDigit {
			leftEnd, rightEnd := leftIndex, rightIndex
			for leftEnd < len(left) && left[leftEnd] >= '0' && left[leftEnd] <= '9' {
				leftEnd++
			}
			for rightEnd < len(right) && right[rightEnd] >= '0' && right[rightEnd] <= '9' {
				rightEnd++
			}

			leftSignificant, rightSignificant := leftIndex, rightIndex
			for leftSignificant < leftEnd-1 && left[leftSignificant] == '0' {
				leftSignificant++
			}
			for rightSignificant < rightEnd-1 && right[rightSignificant] == '0' {
				rightSignificant++
			}

			leftLength := leftEnd - leftSignificant
			rightLength := rightEnd - rightSignificant
			if leftLength != rightLength {
				return leftLength < rightLength
			}
			if leftNumber, rightNumber := left[leftSignificant:leftEnd], right[rightSignificant:rightEnd]; leftNumber != rightNumber {
				return leftNumber < rightNumber
			}
			if leftRunLength, rightRunLength := leftEnd-leftIndex, rightEnd-rightIndex; leftRunLength != rightRunLength {
				return leftRunLength < rightRunLength
			}

			leftIndex, rightIndex = leftEnd, rightEnd
			continue
		}

		if left[leftIndex] != right[rightIndex] {
			return left[leftIndex] < right[rightIndex]
		}
		leftIndex++
		rightIndex++
	}

	return len(left) < len(right)
}

func sortedVersionCounts(values map[string]int) []count {
	counts := sortedCounts(values)
	sort.Slice(counts, func(i, j int) bool {
		return naturalLess(counts[i].Key, counts[j].Key)
	})
	return counts
}

func generateReadmeSection(title, description string, workers []WorkerInfo, filter func(WorkerInfo) bool) reportSection {
	filtered := make([]WorkerInfo, 0)
	versions := make(map[string]int)
	imagesets := make(map[string]int)
	hasLegacyTotals := false

	for _, worker := range workers {
		if filter(worker) {
			filtered = append(filtered, worker)
			versions[worker.Version]++
			imagesets[worker.Imageset]++
			hasLegacyTotals = hasLegacyTotals || worker.LegacyTotalWorkers != nil || worker.LegacyTotalCapacity != nil
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return strings.Compare(filtered[i].WorkerPoolID, filtered[j].WorkerPoolID) < 0
	})

	return reportSection{
		Title:           title,
		Description:     description,
		Count:           len(filtered),
		Versions:        sortedVersionCounts(versions),
		Images:          sortedCounts(imagesets),
		Filtered:        filtered,
		FullColumns:     title == "Generic Worker",
		HasLegacyTotals: hasLegacyTotals,
	}
}

func writeReadme(snapshot WorkerSnapshot) {
	filename := filepath.Join(outputDir, "README.md")
	WriteFile(filename, []byte(renderReadme(snapshot)))
}

func renderReadme(snapshot WorkerSnapshot) string {
	workers := snapshot.Workers
	sections := [5]reportSection{
		generateReadmeSection("Generic Worker", "", workers, func(w WorkerInfo) bool { return w.Implementation == "generic-worker" }),
		generateReadmeSection("Docker Worker", "", workers, func(w WorkerInfo) bool { return w.Implementation == "docker-worker" }),
		generateReadmeSection("Script Worker", "", workers, func(w WorkerInfo) bool { return strings.Contains(w.Implementation, "Scriptworker") }),
		generateReadmeSection("No artifacts found", "These pools claimed and resolved the probe task, but did not publish `public/logs/live_backing.log` or `public/logs/chain_of_trust.log`, which are used to identify the worker implementation.", workers, func(w WorkerInfo) bool { return w.hasNoArtifacts }),
		generateReadmeSection("Version not determined", "These pools did not claim the probe task within two hours, so their worker implementation and version could not be determined.", workers, func(w WorkerInfo) bool { return w.isUnknown }),
	}

	const timestampFormat = "2006-01-02 15:04 UTC"
	data := reportData{Sections: sections}
	if !snapshot.GeneratedAt.IsZero() {
		data.GeneratedAt = snapshot.GeneratedAt.UTC().Format(timestampFormat)
	}
	if !snapshot.ProbeStartedAt.IsZero() {
		data.ProbeStartedAt = snapshot.ProbeStartedAt.UTC().Format(timestampFormat)
	}

	return renderTemplate(data)
}

func readSnapshot(filename string) (WorkerSnapshot, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return WorkerSnapshot{}, err
	}

	var snapshot WorkerSnapshot
	if strings.HasPrefix(strings.TrimSpace(string(contents)), "[") {
		if err := json.Unmarshal(contents, &snapshot.Workers); err != nil {
			return WorkerSnapshot{}, err
		}
	} else if err := json.Unmarshal(contents, &snapshot); err != nil {
		return WorkerSnapshot{}, err
	}

	// These flags are internal rendering state and are not serialized in the
	// snapshot. Restore them from the persisted error value for offline renders.
	for i := range snapshot.Workers {
		switch snapshot.Workers[i].Details["error"] {
		case "No artifacts found":
			snapshot.Workers[i].hasNoArtifacts = true
		case "Version not determined; task not (yet) claimed":
			snapshot.Workers[i].isUnknown = true
		}
	}

	return snapshot, nil
}
