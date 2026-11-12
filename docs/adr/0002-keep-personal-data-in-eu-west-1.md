# ADR-0002: Keep the personal data of the wallet in eu-west-1

- Status: Accepted
- Date: 2025-03-04

## Context

The passes and the trips hold the names, the bookings and the member numbers of the customers.

## Decision

- Every store of personal data of the trip wallet runs in eu-west-1: the database, the queues and the buckets.
- Caches of personal data live in memory, for 5 minutes at most.
- The logs hold no name, member number or booking reference.

## Consequences

- A new store, or a new region, needs a review of the data protection officer of OpCo A.
