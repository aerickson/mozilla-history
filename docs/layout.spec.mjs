import { expect, test } from '@playwright/test'

async function loadReport(page) {
  await page.goto('/docs/index.html?local')
  await expect(page.locator('#content h1')).toBeVisible({ timeout: 30_000 })
  await expect(page.locator('#toc-list li').first()).toBeVisible()
}

test('wide sidebar collapses, persists, and never overlaps content', async ({ page }, testInfo) => {
  await page.setViewportSize({ width: 1440, height: 900 })
  await loadReport(page)

  const toc = page.locator('#toc')
  const content = page.locator('.content-column')
  const toggle = page.locator('#toc-toggle')
  const expandedToc = await toc.boundingBox()
  const expandedContent = await content.boundingBox()

  expect(expandedToc.x + expandedToc.width).toBeLessThanOrEqual(expandedContent.x)
  await page.screenshot({ path: testInfo.outputPath('wide-expanded.png') })

  await toggle.click()
  await expect(toggle).toHaveAttribute('aria-expanded', 'false')
  await expect(page.locator('#toc-list')).toBeHidden()
  const collapsedContent = await content.boundingBox()
  expect(collapsedContent.x).toBeLessThan(expandedContent.x)
  await page.screenshot({ path: testInfo.outputPath('wide-collapsed.png') })

  await page.reload()
  await expect(toggle).toHaveAttribute('aria-expanded', 'false')
  await expect(page.locator('#toc-list')).toBeHidden()
})

test('narrow sidebar stays in flow above the report', async ({ page }, testInfo) => {
  await page.setViewportSize({ width: 720, height: 900 })
  await loadReport(page)

  const toc = page.locator('#toc')
  const content = page.locator('.content-column')
  const tocBox = await toc.boundingBox()
  const contentBox = await content.boundingBox()

  expect(await toc.evaluate(element => getComputedStyle(element).position)).toBe('static')
  expect(tocBox.y + tocBox.height).toBeLessThanOrEqual(contentBox.y)
  await page.screenshot({ path: testInfo.outputPath('narrow.png') })
})

test('wide tables retain page scrolling and sticky cells', async ({ page }, testInfo) => {
  await page.setViewportSize({ width: 1280, height: 800 })
  await loadReport(page)
  await page.locator('#toc-toggle').click()

  const table = page.locator('#history-content table').last()
  await table.scrollIntoViewIfNeeded()
  expect(await table.evaluate(element => getComputedStyle(element).overflow)).toBe('visible')
  expect(await page.evaluate(() => document.documentElement.scrollWidth > innerWidth)).toBe(true)

  const tableBox = await table.boundingBox()
  const tableTop = await page.evaluate(y => scrollY + y, tableBox.y)
  await page.evaluate(({ x, y }) => scrollTo(x, y), {
    x: 500,
    y: tableTop + 100,
  })

  const firstHeader = table.locator('thead th').first()
  const firstCell = table.locator('tbody td').first()
  expect((await firstHeader.boundingBox()).y).toBeCloseTo(0, 0)
  expect((await firstCell.boundingBox()).x).toBeCloseTo(0, 0)
  await page.screenshot({ path: testInfo.outputPath('sticky-table.png') })
})
