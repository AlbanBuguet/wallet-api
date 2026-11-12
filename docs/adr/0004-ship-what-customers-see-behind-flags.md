# ADR-0004: Ship what customers see behind feature flags

- Status: Accepted
- Date: 2025-05-12

## Context

A pass in a wallet stays on the phone of the customer. A wrong field shows until the pass updates.

## Decision

- Every change that a customer sees ships behind a flag of `internal/flags`, named `wallet.<feature>`, off by default.
- The team turns it on in staging, then for 5% of the passes, then for all.
- The code without the flag stays until the flag has been on for a month.

## Consequences

- Tests cover the change with the flag on and off.
