import assert from 'node:assert/strict'
import test from 'node:test'
import { staleProbeAge } from './report-freshness.js'

const timing = 'Probe run started: 2026-09-09 07:58 UTC · Report generated: 2026-09-11 07:58 UTC'
const started = Date.parse('2026-09-09T07:58:00Z')
const hour = 60 * 60 * 1000

test('fresh and future probes do not show a notice, including the threshold', () => {
  for (const age of [-hour, 0, hour, 2 * hour]) {
    assert.equal(staleProbeAge(timing, started + age), null)
  }
})

test('staleness uses the probe time and displays hours or whole days', () => {
  assert.equal(staleProbeAge(timing, started + 2 * hour + 1), '2 hours ago')
  assert.equal(staleProbeAge(timing, started + 24 * hour), '1 day ago')
  assert.equal(staleProbeAge(timing, started + 49 * hour), '2 days ago')
})

test('missing and invalid probe dates do not show a notice', () => {
  for (const value of ['', 'Report generated: 2026-09-09 07:58 UTC',
    'Probe run started: 2026-02-30 07:58 UTC', 'Probe run started: 2026-09-09 25:58 UTC']) {
    assert.equal(staleProbeAge(value, started + 49 * hour), null)
  }
})
