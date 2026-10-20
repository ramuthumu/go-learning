// Optional UI check: install Playwright or set PLAYWRIGHT_MODULE to its module path.
const { chromium } = require(process.env.PLAYWRIGHT_MODULE || 'playwright');
const assert = require('node:assert/strict');
const path = require('node:path');
const { pathToFileURL } = require('node:url');
const url = process.env.SITE_URL || pathToFileURL(path.resolve(__dirname, '../index.html')).href;
(async () => {
  const browser = await chromium.launch({channel: 'chrome', headless: true});
  try {
    const context = await browser.newContext({viewport:{width:1440,height:1000}});
    const page = await context.newPage();
    const errors = []; page.on('pageerror', e => errors.push(e.message));
    await page.goto(url);
    await page.waitForSelector('body.enhanced');
    assert.equal(await page.locator('.lesson').count(),36);
    assert.equal(await page.locator('.view:visible').count(),1);
    await page.screenshot({path:'/tmp/go-learning-desktop.png',fullPage:true});
    const ids = await page.locator('.lesson').evaluateAll(xs=>xs.map(x=>x.id));
    for(const id of ids) {
      await page.locator(`[data-lesson="${id}"]`).click();
      await page.locator(`#${id}`).waitFor({state:'visible'});
      assert.equal(await page.locator('.view:visible').getAttribute('id'),id);
      assert.ok((await page.locator(`#${id} h2`).textContent()).length>5);
    }
    await page.goto(url+'#first-program');
    await page.locator('[data-note="first-program"]').fill('A module contains packages.');
    await page.locator('[data-check="first-program:0"]').check();
    await page.locator('[data-complete="first-program"]').click();
    await page.reload();
    assert.equal(await page.locator('[data-note="first-program"]').inputValue(),'A module contains packages.');
    assert.equal(await page.locator('[data-check="first-program:0"]').isChecked(),true);
    assert.match(await page.locator('#progress-label').textContent(),/^1 of 36/);
    await page.locator('#first-program .reference summary').click();
    assert.equal(await page.locator('#first-program .reference pre').isVisible(),true);
    await page.locator('#search').fill('idempotency');
    assert.ok(await page.locator('.lesson-link:visible').count()>0);
    await page.locator('#search').fill('zzzz_no_topic');
    assert.equal(await page.locator('#search-empty').isVisible(),true);
    await page.locator('#search').fill('');
    const downloadPromise=page.waitForEvent('download');await page.locator('#export').click();
    const download=await downloadPromise;assert.equal(download.suggestedFilename(),'go-learning-progress.json');
    await page.setViewportSize({width:390,height:844});
    await page.locator('#menu-toggle').click();
    assert.equal(await page.locator('#sidebar').isVisible(),true);
    await page.locator('[data-lesson="html-forms"]').click();
    await page.locator('#html-forms').waitFor({state:'visible'});
    assert.equal(await page.locator('#sidebar').isVisible(),false);
    assert.equal(await page.evaluate(()=>document.documentElement.scrollWidth<=innerWidth+1),true);
    await page.screenshot({path:'/tmp/go-learning-mobile.png',fullPage:true});
    await page.evaluate(()=>localStorage.setItem('go-in-practice-v1','{broken'));
    await page.reload();
    assert.equal(await page.locator('.view:visible').count(),1);
    await context.close();
    const noJS=await browser.newContext({javaScriptEnabled:false});
    const plain=await noJS.newPage();await plain.goto(url);
    assert.equal(await plain.locator('.lesson:visible').count(),36);
    await noJS.close();
    assert.deepEqual(errors,[]);
    console.log('PASS: 36 routes, notes/progress reload, checkpoints, references, search, export, mobile layout, corrupted storage, no-JS reading; no page errors.');
  } finally { await browser.close(); }
})().catch(e=>{console.error(e);process.exit(1)});
