# ADR-0005: Version the API additively

- Status: Accepted
- Date: 2025-06-30

## Context

`mobile-bff` and `wallet-web` call `/v1`, and the apps of OpCo A stay on phones for months.

## Decision

- `/v1` only grows: a change adds optional fields or new operations.
- No field is renamed, removed or made required. No status code changes.
- A change that cannot be additive goes to `/v2`, and `/v1` stays for a year.
- `openapi/wallet-api.yaml` describes every change, in the same pull request.

## Consequences

- The consumers ignore the fields that they do not know.
