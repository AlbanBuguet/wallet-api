# wallet-api

The API of the trip wallet of OpCo A. It issues the boarding passes that customers add to Apple Wallet and Google Wallet, keeps their trips, and updates the passes when a flight changes.

## What it talks to

- `wallet-pass-signer` signs each pass. The keys never leave AWS KMS.
- `flight-status-api` gives the gate, the times and the delays of each flight.
- `accounts-api` of Loyalty accounts gives the tier of a member, through the API gateway.
- `wallet-web` and `mobile-bff` call its `/v1`.

## Run it

```sh
go run ./cmd/wallet-api
```

## Test it

```sh
go test ./...
```
