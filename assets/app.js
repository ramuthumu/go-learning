(() => {
  'use strict';
  const lessons = [...document.querySelectorAll('.lesson')];
  const ids = new Set(lessons.map(x => x.id));
  const key = 'go-in-practice-v1';
  const notice = document.querySelector('#notice');
  const empty = () => ({version: 1, current: 'overview', done: [], checks: {}, notes: {}});
  const validChecks = new Set([...document.querySelectorAll('[data-check]')].map(x => x.dataset.check));
  function clean(input) {
    if (!input || typeof input !== 'object' || Array.isArray(input)) return empty();
    const result = empty();
    result.current = ids.has(input.current) ? input.current : 'overview';
    result.done = [...new Set(Array.isArray(input.done) ? input.done.filter(id => ids.has(id)) : [])];
    if (input.notes && typeof input.notes === 'object') {
      for (const id of ids) if (typeof input.notes[id] === 'string') result.notes[id] = input.notes[id].slice(0, 40000);
    }
    if (input.checks && typeof input.checks === 'object') {
      for (const id of validChecks) if (input.checks[id] === true) result.checks[id] = true;
    }
    return result;
  }
  let state;
  let storageWorks = true;
  try { state = clean(JSON.parse(localStorage.getItem(key))); }
  catch { state = empty(); storageWorks = false; }
  function announce(message) { notice.textContent = message; }
  function save() {
    try { localStorage.setItem(key, JSON.stringify(state)); }
    catch { storageWorks = false; document.querySelector('#storage-status').textContent = 'Browser storage unavailable. Use Back up before closing.'; }
  }
  function updateProgress() {
    const completed = new Set(state.done);
    document.querySelector('#progress').value = completed.size;
    document.querySelector('#progress-label').textContent = `${completed.size} of ${lessons.length} complete`;
    document.querySelectorAll('[data-lesson]').forEach(link => {
      const done = completed.has(link.dataset.lesson);
      link.classList.toggle('done', done);
      link.title = done ? 'Completed' : '';
    });
    document.querySelectorAll('[data-stage-count]').forEach(span => {
      const group = lessons.filter(x => x.dataset.stage === span.dataset.stageCount);
      span.textContent = `${group.filter(x => completed.has(x.id)).length}/${group.length}`;
    });
    document.querySelectorAll('[data-complete]').forEach(button => {
      const done = completed.has(button.dataset.complete);
      button.textContent = done ? 'Completed ✓ · undo' : 'Mark complete';
      button.setAttribute('aria-pressed', String(done));
    });
    const resume = document.querySelector('#resume');
    const next = ids.has(state.current) ? state.current : lessons.find(x => !completed.has(x.id))?.id || lessons[0].id;
    resume.href = '#' + next;
    resume.textContent = ids.has(state.current) || completed.size ? 'Continue learning →' : 'Start learning →';
  }
  function restoreFields() {
    document.querySelectorAll('[data-note]').forEach(field => { field.value = state.notes[field.dataset.note] || ''; });
    document.querySelectorAll('[data-check]').forEach(field => { field.checked = !!state.checks[field.dataset.check]; });
  }
  function showCurrent(focus = false) {
    const hash = location.hash.slice(1);
    if (hash === 'main' && document.body.classList.contains('enhanced')) {
      document.querySelector('#main').focus();
      return;
    }
    const id = ids.has(hash) ? hash : 'overview';
    document.querySelectorAll('.view').forEach(view => { view.hidden = view.id !== id; });
    document.querySelectorAll('[data-lesson]').forEach(link => {
      if (link.dataset.lesson === id) { link.setAttribute('aria-current', 'page'); link.closest('details').open = true; }
      else link.removeAttribute('aria-current');
    });
    if (ids.has(id)) { state.current = id; save(); }
    updateProgress();
    document.querySelector('#sidebar').classList.remove('open');
    document.querySelector('#menu-toggle').setAttribute('aria-expanded', 'false');
    if (focus) {
      document.getElementById(id).focus({preventScroll: true});
      window.scrollTo(0, 0);
    }
  }
  document.querySelector('#menu-toggle').addEventListener('click', event => {
    const opened = document.querySelector('#sidebar').classList.toggle('open');
    event.currentTarget.setAttribute('aria-expanded', String(opened));
  });
  document.querySelector('#search').addEventListener('input', event => {
    const query = event.target.value.trim().toLowerCase();
    let matches = 0;
    document.querySelectorAll('[data-lesson]').forEach(link => {
      const article = document.getElementById(link.dataset.lesson);
      const match = !query || article.textContent.toLowerCase().includes(query);
      link.hidden = !match;
      if (match) matches++;
    });
    document.querySelectorAll('.stage-nav').forEach(group => {
      group.hidden = ![...group.querySelectorAll('[data-lesson]')].some(x => !x.hidden);
      if (query && !group.hidden) group.open = true;
    });
    document.querySelector('#search-empty').hidden = matches > 0;
  });
  document.querySelectorAll('[data-complete]').forEach(button => button.addEventListener('click', () => {
    const id = button.dataset.complete;
    state.done = state.done.includes(id) ? state.done.filter(x => x !== id) : [...state.done, id];
    save(); updateProgress();
    announce('Progress saved in this browser.');
  }));
  document.querySelectorAll('[data-check]').forEach(field => field.addEventListener('change', () => {
    state.checks[field.dataset.check] = field.checked; save();
  }));
  document.querySelectorAll('[data-note]').forEach(field => field.addEventListener('input', () => {
    state.notes[field.dataset.note] = field.value; save();
  }));
  document.querySelectorAll('.copy').forEach(button => button.addEventListener('click', async () => {
    const text = button.closest('.code').querySelector('pre').textContent;
    try {
      if (navigator.clipboard && window.isSecureContext) await navigator.clipboard.writeText(text);
      else {
        const field = document.createElement('textarea'); field.value = text; field.style.position = 'fixed'; field.style.opacity = '0';
        document.body.append(field); field.select(); const copied = document.execCommand('copy'); field.remove();
        if (!copied) throw new Error('copy unavailable');
      }
      button.textContent = 'Copied'; setTimeout(() => {button.textContent = 'Copy';}, 1800);
    } catch { button.textContent = 'Select code to copy'; }
  }));
  document.querySelector('#export').addEventListener('click', () => {
    const blob = new Blob([JSON.stringify(state, null, 2)], {type:'application/json'});
    const url = URL.createObjectURL(blob); const a = document.createElement('a');
    a.href = url; a.download = 'go-learning-progress.json'; a.click(); setTimeout(() => URL.revokeObjectURL(url), 1000);
  });
  document.querySelector('#import').addEventListener('click', () => document.querySelector('#import-file').click());
  document.querySelector('#import-file').addEventListener('change', async event => {
    const file = event.target.files[0]; if (!file) return;
    try {
      if (file.size > 2_000_000) throw new Error('backup too large');
      const input = JSON.parse(await file.text());
      if (input?.version !== 1 || !Array.isArray(input.done)) throw new Error('not a course backup');
      const restored = clean(input);
      state.done = [...new Set([...state.done, ...restored.done])];
      state.checks = {...state.checks, ...restored.checks};
      state.notes = {...restored.notes, ...state.notes};
      save(); restoreFields(); updateProgress(); announce('Backup merged. Existing notes were preserved.');
    } catch { announce('Could not restore this file. Choose a course progress JSON backup.'); }
    event.target.value = '';
  });
  let printDetails = [];
  window.addEventListener('beforeprint', () => { printDetails = [...document.querySelectorAll('main details')].filter(x => !x.open); printDetails.forEach(x => {x.open = true;}); });
  window.addEventListener('afterprint', () => {printDetails.forEach(x => {x.open = false;});});
  document.querySelector('#print').addEventListener('click', () => window.print());
  window.addEventListener('hashchange', () => showCurrent(true));
  restoreFields(); showCurrent();
  document.body.classList.add('enhanced');
  if (!storageWorks) document.querySelector('#storage-status').textContent = 'Browser storage may be unavailable. Use Back up to keep progress.';
})();
