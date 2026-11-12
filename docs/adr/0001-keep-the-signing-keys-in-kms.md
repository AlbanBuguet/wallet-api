# ADR-0001: Keep the signing keys of the passes in AWS KMS

- Status: Accepted
- Date: 2025-02-18

## Context

Apple Wallet and Google Wallet trust a pass that the key of OpCo A signs. Whoever holds the key can issue a pass for any flight.

## Decision

- The signing keys live in AWS KMS, in the production account of the trip wallet, and never leave it.
- Only `wallet-pass-signer` signs, through `kms:Sign`, with its own IAM role.
- No key, certificate password or key material goes in code, configuration, CI variables or logs.

## Consequences

- `wallet-api` asks `wallet-pass-signer` to sign each pass.
- A new kind of pass, or a new field, changes `wallet-pass-signer` too.
