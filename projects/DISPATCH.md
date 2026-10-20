# Dispatch: independent enterprise capstone

Build a small webhook delivery service using the knowledge from stages 1–6.
This is a specification and schema scaffold, **not an implemented reference
service**. Start with one process, SQLite, and operator-configured destinations.

## Scope and contract

- `POST /events`: verified tenant identity plus JSON `{topic, payload}` and an
  `Idempotency-Key`. Commit the event and its pending delivery together, then
  return `202` and the event ID.
- The same tenant/key/payload returns the original accepted result. The same
  tenant/key with different input returns `409`. Define whether payload equality
  means exact bytes or a canonical representation; test that policy.
- `GET /deliveries/{id}`: returns state, attempt count, and next scheduled attempt
  for the verified tenant. Unknown or unauthorized resources must not leak data.
- An operator can replay a dead delivery through an authorized, audited action.
- Destinations come from trusted configuration in version one. Do not accept
  arbitrary internet URLs from unauthenticated requests.

## Milestones

1. **Acceptance:** request size limits, validation, categorized errors, request
   IDs, and a store interface with only the operations the use case needs.
2. **Persistence:** schema migrations; event and work-item writes in one
   transaction; unique tenant/idempotency key; durable lookup after restart.
3. **Delivery:** fixed worker count, shared HTTP client, per-attempt deadline,
   response/body limits, persisted attempt state, and a stable delivery ID.
4. **Recovery:** atomic claims with expiring leases and claim tokens; stale
   workers cannot update a newer claim; crashes leave recoverable jobs.
5. **Retry:** explicit response classification, bounded backoff plus jitter,
   `Retry-After` where applicable, exhausted/dead state, and controlled replay.
6. **Access:** verified tenant identity and resource authorization in API,
   storage, caches, worker payloads, and operator tools.
7. **Operations:** request/delivery logs, queue age, retry counts, readiness,
   bounded shutdown, backup/restore instructions, and a release workflow.

Use [the schema scaffold](../examples/dispatch/schema.sql) as a starting point.
The scaffold deliberately does not implement the job-claim transaction: explain
and test the atomicity of the algorithm you choose for your database.

## Failure drills and acceptance criteria

| Drill | Evidence required |
|---|---|
| Crash after returning 202 | The event and pending work survive restart |
| Two identical requests arrive together | One logical event is accepted |
| Same key, different payload | A conflict is returned without altering the original |
| Destination returns 503 twice, then 204 | Attempts and retry schedule follow the policy |
| Destination accepts, response is lost | Repeated delivery uses the same ID; receiver deduplicates |
| Worker dies with a claimed job | The job becomes claimable after lease expiry |
| Stale worker reports success | Its old lease token cannot overwrite a newer claim |
| Tenant B guesses tenant A's ID | No record or private details are exposed |
| Traffic exceeds worker capacity | Concurrency stays bounded; pending work is observable |
| Shutdown during delivery | Work either completes within deadline or remains recoverable |
| Backup is restored to a fresh environment | Event and job history can be queried and resumed |

## Review packet

Deliver a README, API contract, migration procedure, test output, a short
architecture diagram, and a runbook. Explain at-least-once delivery, duplicate
handling, lease ownership, and failure recovery. State limits honestly, including
single-node assumptions and any checks not yet implemented. Ask a reviewer to
choose a new requirement and make a small change with you.

## Extensions after the first complete version

Move to PostgreSQL and verify concurrency semantics on PostgreSQL itself. Add
signed outbound payloads using a reviewed HMAC contract with key rotation. Add
an event broker only after you can explain which scaling or integration need it
solves. Compare a modular monolith with independently deployed services using
measured requirements rather than a preference for more infrastructure.
