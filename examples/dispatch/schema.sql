-- Learning schema for a single-node durable webhook service using SQLite.
PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS events (
    tenant_id TEXT NOT NULL,
    id TEXT NOT NULL,
    idempotency_key TEXT NOT NULL,
    body_hash TEXT NOT NULL,
    payload BLOB NOT NULL,
    created_at INTEGER NOT NULL,
    PRIMARY KEY (tenant_id, id),
    UNIQUE (tenant_id, idempotency_key)
);
CREATE TABLE IF NOT EXISTS deliveries (
    tenant_id TEXT NOT NULL,
    id TEXT NOT NULL,
    event_id TEXT NOT NULL,
    destination_id TEXT NOT NULL,
    state TEXT NOT NULL CHECK (state IN ('pending','leased','sent','dead')),
    attempts INTEGER NOT NULL DEFAULT 0,
    next_attempt INTEGER NOT NULL,
    lease_until INTEGER,
    lease_token TEXT,
    PRIMARY KEY (tenant_id, id),
    FOREIGN KEY (tenant_id, event_id) REFERENCES events(tenant_id, id)
);
CREATE INDEX IF NOT EXISTS deliveries_ready ON deliveries(state, next_attempt);
