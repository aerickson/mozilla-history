// This works for both the Mozilla and community history GitHub Pages sites.
let basePath = String(window.location.pathname)
if (!basePath.includes('-history')) basePath = '/mozilla-history/'

const localPreview = new URLSearchParams(window.location.search).has('local')
const readmeUrl = localPreview
  ? new URL('../WorkerVersions/README.md', window.location.href)
  : `https://raw.githubusercontent.com/taskcluster${basePath}master/WorkerVersions/README.md`
const historyUrl = new URL('history.json', window.location.href)

async function loadReadme() {
  const response = await fetch(readmeUrl)
  const markdown = await response.text()
  showdown.setFlavor('github')
  document.getElementById('content').innerHTML = new showdown.Converter({
    ghCompatibleHeaderId: true,
    tables: true,
  }).makeHtml(markdown)
}

function renderHistoryTables(data) {
  const dates = Object.keys(data).sort((a, b) => b.localeCompare(a))

  const tableByDate = (title, property, description = '') => {
    let columns = []
    dates.forEach(date => {
      const values = data[date][property]
      Object.keys(values).forEach(key => {
        if (!columns.includes(key)) columns.push(key)
      })
    })
    columns = columns.sort((a, b) => a.localeCompare(b))

    let table = `<h3>${title}</h3>${description}`
    table += `<table><thead><tr><th>Date</th>${columns.map(key => `<th>${key || 'unknown'}</th>`).join('')}</thead><tbody>`
    dates.forEach(date => {
      table += `<tr><td>${date}</td>
        ${columns.map(column => data[date][property][column] || '-').map(value => `<td>${value}</td>`).join('')}
      </tr>`
    })
    table += '</tbody></table>'
    return table
  }

  let html = '<h2>History</h2>'
  html += tableByDate(
    'Worker implementations',
    'implementations',
    '<p class="info">Worker implementation is inferred from distinctive content in the log artifact produced by the intentionally malformed probe task. <strong>Docker Worker identifies the legacy docker-worker implementation; it does not indicate that a generic-worker pool is configured to accept docker-worker-style payloads.</strong></p>',
  )
  html += tableByDate('Worker versions', 'versions')
  document.getElementById('history-content').innerHTML = html
}

function renderGraphs(data) {
  const dates = Object.keys(data).sort()
  const implementationTraces = [
    ['Generic worker', 'generic-worker'],
    ['Docker worker', 'docker-worker'],
    ['Unknown', ''],
  ].map(([name, implementation]) => ({
    x: dates,
    y: dates.map(date => data[date].implementations[implementation] || 0),
    name,
    type: 'line',
  }))

  const allVersions = [...new Set(dates.map(date => Object.keys(data[date].versions)).flat())].sort()
  const versionTraces = allVersions.map(version => ({
    x: dates,
    y: dates.map(date => data[date].versions[version] || 0),
    name: version,
    type: 'line',
  }))

  const layout = { barmode: 'group' }
  Plotly.newPlot('graph-implementations', implementationTraces, {
    ...layout,
    title: 'Worker Implementations',
  })
  Plotly.newPlot('graph-versions', versionTraces, {
    ...layout,
    title: 'Worker Versions',
  })
}

async function loadHistory() {
  const response = await fetch(historyUrl)
  const data = await response.json()
  renderHistoryTables(data)
  renderGraphs(data)
}

function slugify(value) {
  return value
    .toLowerCase()
    .trim()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-|-$/g, '')
}

function buildTableOfContents() {
  const headings = document.querySelectorAll([
    '#content h2',
    '#content h3',
    '#history-content h2',
    '#history-content h3',
    '#graphs-content h2',
    '#graphs-content h3',
  ].join(', '))
  const toc = document.getElementById('toc-list')
  const usedIds = new Set()
  let sectionId = ''
  let sectionItem = null
  let subsectionList = null

  headings.forEach(heading => {
    const baseId = slugify(heading.textContent)
    const candidateId = heading.tagName === 'H3' && sectionId
      ? `${sectionId}-${baseId}`
      : baseId
    let id = candidateId
    let suffix = 2
    while (usedIds.has(id)) {
      id = `${candidateId}-${suffix}`
      suffix += 1
    }
    usedIds.add(id)
    heading.id = id

    const item = document.createElement('li')
    const link = document.createElement('a')
    link.href = `#${id}`
    link.textContent = heading.textContent
    item.appendChild(link)

    if (heading.tagName === 'H2') {
      sectionId = id
      sectionItem = item
      subsectionList = null
      toc.appendChild(item)
      return
    }

    if (!subsectionList) {
      subsectionList = document.createElement('ul')
      sectionItem.appendChild(subsectionList)
    }
    subsectionList.appendChild(item)
  })
}

async function init() {
  await Promise.all([loadReadme(), loadHistory()])
  buildTableOfContents()
  document.querySelectorAll('table').forEach(table => table.classList.add('sortable'))
}

init()
