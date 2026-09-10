import assert from 'node:assert/strict'
import { readFile } from 'node:fs/promises'
import test from 'node:test'

const html = await readFile(new URL('./index.html', import.meta.url), 'utf8')

test('the report page has one syntactically valid inline module', () => {
  const modules = [...html.matchAll(/<script type="module">([\s\S]*?)<\/script>/g)]
  assert.equal(modules.length, 1)

  const sourceWithoutImports = modules[0][1].replace(/^\s*import .*$/gm, '')
  assert.doesNotThrow(() => new Function(sourceWithoutImports))
})

test('Quick and local previews use the bundled report', () => {
  assert.match(html, /query\.has\('local'\)/)
  assert.match(html, /'127\.0\.0\.1', 'localhost', '\[::1\]'/)
  assert.match(html, /loopbackHost \|\| window\.location\.hostname/)
  assert.match(html, /hostname\.endsWith\('\.quick\.mozilla\.cloud'\)/)
  assert.match(html, /new URL\('\.\.\/WorkerVersions\/', baseUrl\)/)
})

test('report metadata is inserted at the end of the report introduction', () => {
  assert.doesNotMatch(html, /<div class="content-column">\s*<p class="source"/)
  assert.match(html, /const firstSection = content\.querySelector\('h2'\)/)
  assert.match(html, /if \(firstSection\) firstSection\.before\(context\)/)
  assert.match(html, /Source code: <a href="\$\{repoUrl\}">GitHub<\/a>/)
  assert.match(html, /Current worker-pool report \(Markdown\)/)
  assert.match(html, /Current worker-pool data \(JSON\)/)
  assert.match(html, /Historical worker-pool data \(JSON\)/)
  assert.match(html, /startsWith\('Probe run started:'\)/)
  assert.match(html, /if \(reportTiming\) reportTiming\.append\(coverage\)/)
  assert.match(html, /includes \$\{summary\}/)
  assert.match(html, /\$\{dates\.length\} snapshots from \$\{dates\[0\]\} to \$\{dates\[dates\.length - 1\]\}/)
  assert.doesNotMatch(html, /\.report-context \{[\s\S]*?font-size:/)
  assert.doesNotMatch(html, /Disclaimer:/)
})

test('report artifacts are revalidated at their stable URLs', () => {
  assert.match(html, /fetch\(url, \{ cache: 'no-cache' \}\)/)
  assert.doesNotMatch(html, /cache: 'no-store'/)
})

test('the merged page retains report and upstream table enhancements', () => {
  assert.match(html, /import \{ decorateSortableTables \} from '\.\/table-sort\.js'/)
  assert.match(html, /import \{ enhanceTables \} from '\.\/table-utils\.js'/)
  assert.match(html, /id="toc-list"/)
  assert.match(html, /function buildTableOfContents\(\)/)
  assert.doesNotMatch(html, /\.\/app\.js/)
})
