// Command wallet-api serves the API of the trip wallet.
package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.example.com/opco-a/wallet-api/internal/flags"
	"github.example.com/opco-a/wallet-api/internal/flightstatus"
	"github.example.com/opco-a/wallet-api/internal/pass"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	flights := flightstatus.New(os.Getenv("FLIGHT_STATUS_URL"), &http.Client{Timeout: 2 * time.Second})
	issuer := pass.NewIssuer(pass.Options{
		Flights: flights,
		Signer:  pass.NewSignerClient(os.Getenv("PASS_SIGNER_URL")),
		Flags:   flags.FromEnv(),
	})
	mux := http.NewServeMux()
	mux.Handle("POST /v1/passes", issuer.Handler())
	logger.Info("listening", "addr", ":8080")
	srv := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("server stopped", "err", err)
		os.Exit(1)
	}
}
