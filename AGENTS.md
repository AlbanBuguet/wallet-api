# wallet-api

How coding agents work in wallet-api.

- Build: `go build ./...`
- Test: `go test ./...`
- Lint: `golangci-lint run`
- Read the ADRs in `docs/adr` before a change. They bind every change.
- Every change that a customer sees ships behind a flag of `internal/flags`, off by default.
- `/v1` only grows: add optional fields, never rename or remove one. `mobile-bff` and `wallet-web` depend on it.
- Never log a member number, a name or a booking reference.
- Keep each change small, with its tests.
