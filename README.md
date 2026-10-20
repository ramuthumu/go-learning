# Go, in practice

A practical learning path from your first Go program to professional CLI,
web-app, and enterprise-service development.

**[Open the course](https://ramuthumu.github.io/go-learning/)** ·
[Complete guide](GUIDE.md) · [Quick reference](CHEATSHEET.md)

## What is included

36 lessons across six stages: language foundations, CLI development, web apps,
concurrency, enterprise engineering, and professional operation. Each lesson
contains explanations, a realistic scenario, exercises, checks, recall questions,
and links to primary documentation.

Runnable references include **LogLens** (CLI), **TeamDesk** (HTML app, JSON API,
SQLite), and **Sitecheck** (bounded concurrent CLI). **Dispatch** is an independent
enterprise capstone with a [detailed brief](projects/DISPATCH.md) and schema
scaffold, not a completed service.

## Run the examples

Install a supported Go version, **1.25 or newer**. From this repository:

```sh
go run ./examples/01-hello
go run ./examples/loglens -json examples/loglens/sample.log
go run ./examples/teamdesk -db /tmp/teamdesk-learning.db
# In another terminal, while TeamDesk is running:
go run ./examples/sitecheck http://localhost:8080/health
go test -race ./...
```

TeamDesk binds to localhost and is a single-user teaching reference. It has no
login. The security chapter explains the identity and authorization work required
before using it as a shared application. Code snippets labelled “design excerpt”
are illustrative pieces, not standalone programs.

## Website and course source

Open `index.html` directly or serve the directory with
`python3 -m http.server 8000 --bind 127.0.0.1`. The site has no JavaScript framework,
remote fonts, tracking, or server runtime. All lessons remain readable without
JavaScript. Progress, checkpoints, and notes live in browser storage; Back up and
Restore let you move them between browsers. They are not synced to GitHub.

Edit `content/course.json` and the reference files, then regenerate:

```sh
python3 scripts/build.py
node --check assets/app.js
```

The generated `index.html` and `GUIDE.md` are committed. GitHub Pages publishes the
repository root on `main`; `.nojekyll` keeps this a plain static site. GitHub hosts
the learning material, not the Go example servers. CI checks that generated
content is current and that the examples pass their tests.

Estimated stage pacing is a planning aid, not a promise of expertise. Work at
your own pace, rebuild without the solution, and seek code review on capstones.
Original educational material and code are under the MIT license. External
resources remain under their respective owners' terms.
