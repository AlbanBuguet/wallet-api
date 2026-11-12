# ADR-0003: Call Loyalty through the API gateway

- Status: Accepted
- Date: 2025-09-22

## Context

The wallet will show what Loyalty knows of a member, such as their tier. Loyalty accounts is a Tier 1 system of another domain, with its own release cycle.

## Decision

- `wallet-api` calls Loyalty accounts only through the API gateway of OpCo A, at `https://gateway.opco-a.example/loyalty/v2`, never `accounts-api` directly.
- It signs in with OAuth2 client credentials, with the scope `loyalty.tier.read` and nothing more.
- It caches each answer for 5 minutes, in memory, and times out after 800 ms.
- When Loyalty fails or is slow, the wallet issues the pass without what Loyalty knows. A pass never waits for Loyalty.

## Consequences

- The gateway enforces the quotas of Loyalty.
- Loyalty sees one client, `wallet-api`, in its dashboards.
