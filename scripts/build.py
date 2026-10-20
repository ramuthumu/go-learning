#!/usr/bin/env python3
"""Build a dependency-free, progressively enhanced course from structured content."""
import html
import json
import re
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
course = json.loads((ROOT / 'content/course.json').read_text())
lessons = course['lessons']
stages = course['stages']
ids = [lesson['id'] for lesson in lessons]
assert len(ids) == len(set(ids)), 'Duplicate lesson IDs'
assert len(lessons) == 36, 'Expected 36 complete lessons'
assert all(lesson['stage'] in {s['id'] for s in stages} for lesson in lessons)

def esc(value):
    return html.escape(str(value), quote=True)

def prose(value):
    return re.sub(r'`([^`]+)`', r'<code>\1</code>', esc(value))

def codeblock(code, label='Code'):
    return f'<div class="code"><div class="code-top"><span>{esc(label)}</span><button class="copy" type="button" aria-label="Copy {esc(label)}">Copy</button></div><pre><code>{esc(code.rstrip())}</code></pre></div>'

nav = []
cards = []
for i, stage in enumerate(stages, 1):
    group = [l for l in lessons if l['stage'] == stage['id']]
    links = ''.join(f'<a class="lesson-link" data-lesson="{l["id"]}" href="#{l["id"]}"><span class="status-dot" aria-hidden="true"></span><span>{esc(l["title"])}</span></a>' for l in group)
    nav.append(f'<details class="stage-nav" open data-stage="{stage["id"]}"><summary>{esc(stage["title"])} <span class="stage-count" data-stage-count="{stage["id"]}">0/{len(group)}</span></summary>{links}</details>')
    cards.append(f'<a class="stage-card" href="#{group[0]["id"]}"><span class="eyebrow">Stage {i} · {esc(stage["pace"])}</span><h3>{esc(stage["title"].split(". ",1)[1])}</h3><p>{esc(stage["outcome"])}</p><span class="project-name">{esc(stage["project"])}</span></a>')

articles = []
book = ['# Go: from first program to professional practice\n', '36 guides • 6 stages • practical projects\n', 'Use this guide with the runnable examples in the repository. All shell commands assume the repository root unless a lesson explicitly creates a separate module.\n']
for i, l in enumerate(lessons):
    stage = next(s for s in stages if s['id'] == l['stage'])
    learned = ''.join(f'<h3>{esc(p["title"])}</h3><p>{prose(p["text"])}</p>' for p in l['learn'])
    steps = ''.join(f'<li>{prose(step)}</li>' for step in l['steps'])
    checks = ''.join(f'<label class="check"><input type="checkbox" data-check="{l["id"]}:{j}"><span>{prose(check)}</span></label>' for j,check in enumerate(l['checks']))
    refs = ''.join(f'<a href="{esc(r["url"])}">{esc(r["title"])}</a>' for r in l['references'])
    commands = '<h3>Run and observe</h3>' + codeblock(l['commands'], 'Terminal · repository root') if l['commands'] else ''
    example_text = l.get('snippet')
    source_link = ''
    if l.get('example'):
        source = ROOT / l['example']
        assert source.is_file(), source
        example_text = source.read_text()
        source_link = f'<p><a href="{esc(l["example"])}">Open source: {esc(l["example"])}</a></p>'
    label = 'Working reference' if l.get('example') and not l['example'].endswith('.sql') else 'Design excerpt'
    solution = f'<details class="reference"><summary>{label} · inspect after trying</summary>{source_link}{codeblock(example_text, l.get("example") or "Illustrative excerpt")}</details>' if example_text else ''
    prev = '#overview' if i == 0 else '#' + lessons[i-1]['id']
    next_link = '#' + lessons[i+1]['id'] if i+1 < len(lessons) else '#overview'
    next_text = 'Next guide →' if i+1 < len(lessons) else 'Back to learning path'
    articles.append(f'''<article class="view lesson" id="{l['id']}" data-stage="{l['stage']}" tabindex="-1">
<div class="eyebrow">{esc(stage['title'])} · Guide {i+1} of {len(lessons)}</div>
<h2>{esc(l['title'])}</h2><p class="lead">{prose(l['scenario'])}</p>
{learned}
<section class="practice"><h3>Build it yourself</h3><ol>{steps}</ol></section>
{commands}{solution}
<aside class="pitfall"><strong>Watch for this</strong><p>{prose(l['pitfall'])}</p></aside>
<section class="checkpoint"><h3>Before you move on</h3>{checks}</section>
<details class="recall"><summary>Recall: {esc(l['question'])}</summary><p>{prose(l['answer'])}</p></details>
<div class="sources"><span>Go deeper</span>{refs}</div>
<label class="notes-label" for="notes-{l['id']}">Your notes <span>What changed? What is still unclear?</span></label>
<textarea class="notes" id="notes-{l['id']}" data-note="{l['id']}" rows="3" placeholder="Keep a short explanation in your own words."></textarea>
<div class="lesson-actions"><a href="{prev}">← Previous</a><button class="complete" data-complete="{l['id']}" type="button">Mark complete</button><a href="{next_link}">{next_text}</a></div>
</article>''')
    book += [f'\n## {i+1}. {l["title"]}\n', l['scenario']+'\n']
    for p in l['learn']: book += [f'### {p["title"]}\n',p['text']+'\n']
    book += ['### Build it yourself\n'] + [f'{j+1}. {s}' for j,s in enumerate(l['steps'])]
    if l['commands']: book += ['\n```sh\n'+l['commands']+'\n```\n']
    if example_text: book += ['\n<details><summary>Reference / excerpt</summary>\n', '```\n'+example_text.rstrip()+'\n```\n','</details>\n']
    book += ['### Checkpoints\n'] + ['- '+c for c in l['checks']]
    book += ['\n**Watch for:** '+l['pitfall']+'\n', '**Recall:** '+l['question']+'\n', l['answer']+'\n', '\nFurther reading: '+ ' · '.join(f'[{r["title"]}]({r["url"]})' for r in l['references'])+'\n']

page = f'''<!doctype html>
<html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><meta name="theme-color" content="#f6f7f3"><meta name="description" content="A practical Go learning path: 36 guides from your first program to CLIs, web apps, concurrency, enterprise systems, and professional delivery."><title>Go, in practice — learn by building</title><link rel="stylesheet" href="assets/style.css"><script src="assets/app.js" defer></script></head>
<body><a class="skip" href="#main">Skip to content</a>
<header class="top"><a class="brand" href="#overview"><span class="go-mark">go</span> Go, in practice</a><div class="top-actions"><button id="menu-toggle" type="button" aria-expanded="false" aria-controls="sidebar">Guides</button><a href="https://github.com/ramuthumu/go-learning">GitHub ↗</a></div></header>
<div class="layout"><aside id="sidebar"><a class="overview-link" href="#overview">Learning path</a><label class="search-label" for="search">Find a topic</label><input type="search" id="search" placeholder="Try “errors” or “HTTP”" autocomplete="off"><p id="search-empty" hidden>No matching guides.</p><nav aria-label="Course chapters">{''.join(nav)}</nav><div class="sidebar-footer"><div id="progress-label">0 of {len(lessons)} complete</div><progress id="progress" max="{len(lessons)}" value="0" aria-label="Course progress"></progress><small id="storage-status">Progress and notes stay in this browser.</small><div class="backup"><button id="export" type="button">Back up</button><button id="import" type="button">Restore</button><input id="import-file" type="file" accept="application/json" hidden></div></div></aside>
<main id="main" tabindex="-1"><section class="view overview" id="overview" tabindex="-1">
<div class="eyebrow">A practical path · 36 guides · 6 stages</div><h1>Learn Go.<br>Build things that work.</h1><p class="lead">From your first line of code to useful CLIs, web applications, and reliable business systems. Learn the idea, build a feature, test a failure, then explain your decision.</p><div class="start-actions"><a class="primary" id="resume" href="#first-program">Start learning →</a><a href="https://github.com/ramuthumu/go-learning/archive/refs/heads/main.zip">Download the examples</a></div>
<div class="intro-note"><strong>Start where you are.</strong> New to programming? Work through stage 1 slowly. Know another language? Use its checkpoints to identify gaps. The first three project references run locally; Dispatch is an independent capstone with a detailed specification.</div>
<h2>Your learning path</h2><div class="stage-grid">{''.join(cards)}</div>
<h2>A routine you can sustain</h2><p>Use 45–60 minute sessions: recall yesterday’s idea, read one explanation, spend most of the session coding, then write a short note. Some guides need several sessions; the capstones need more. Return to a topic after a few days and implement a variation without the reference.</p>
<h3>Set up once</h3><p>Install a supported Go release (1.25 or newer), Git, and an editor. Clone or download this repository. Commands on the lesson pages run from its root. Go code runs on your computer; the website is the guide.</p>{codeblock('git clone https://github.com/ramuthumu/go-learning.git\ncd go-learning\ngo version\ngo run ./examples/01-hello','Terminal')}
<h3>Four projects, increasing responsibility</h3><ul><li><strong>LogLens:</strong> a tested CLI for real log files and shell pipelines.</li><li><strong>TeamDesk:</strong> an HTML web app and JSON API with SQLite persistence.</li><li><strong>Sitecheck:</strong> a bounded, cancellable concurrent CLI.</li><li><strong>Dispatch:</strong> a durable webhook service you design and implement from a <a href="projects/DISPATCH.md">project brief</a>.</li></ul>
<h3>What professional progress looks like</h3><p>You can implement unfamiliar requirements, explain Go’s types and ownership, handle failures, write meaningful tests, protect data, and help another developer run and maintain your work. A completion badge is a study aid; repeated independent delivery and review build professional judgment.</p>
<p class="resource-line"><a href="GUIDE.md">Read the complete guide</a> · <a href="CHEATSHEET.md">Quick reference</a> · <button id="print" type="button">Print / save as PDF</button></p>
</section>{''.join(articles)}
<footer>Original lessons and examples · Further reading links point to official documentation and standards. <span id="notice" role="status" aria-live="polite"></span></footer></main></div>
<noscript><p class="noscript">JavaScript is off. Every lesson is still available below; navigation jumps to each section. Progress tracking needs JavaScript.</p></noscript></body></html>'''
(ROOT/'index.html').write_text(page)
(ROOT/'GUIDE.md').write_text('\n'.join(book))
print(f'Built {len(lessons)} lessons and complete Markdown guide')
