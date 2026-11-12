// Package pass issues the boarding passes of the wallet.
package pass

import (
	"context"
	"encoding/json"
	"net/http"

	"github.example.com/opco-a/wallet-api/internal/flags"
	"github.example.com/opco-a/wallet-api/internal/flightstatus"
)

// BoardingPass is the pass that the wallet shows. Its JSON is the
// BoardingPass of openapi/wallet-api.yaml.
type BoardingPass struct {
	Booking   string `json:"booking"`
	Passenger string `json:"passenger"`
	Flight    string `json:"flight"`
	Seat      string `json:"seat"`
	Gate      string `json:"gate,omitempty"`
	Boarding  string `json:"boarding"`
	Group     string `json:"group"`
}

// Request asks for the pass of a passenger on a flight.
type Request struct {
	Booking   string `json:"booking"`
	Passenger string `json:"passenger"`
	Member    string `json:"member,omitempty"`
	Flight    string `json:"flight"`
	Day       string `json:"day"`
	Seat      string `json:"seat"`
	Group     string `json:"group"`
}

// Flights reads the status of a flight.
type Flights interface {
	Get(ctx context.Context, flight, day string) (flightstatus.Status, error)
}

// Signer signs a pass for the wallets.
type Signer interface {
	Sign(ctx context.Context, p BoardingPass) ([]byte, error)
}

// Options configures an Issuer.
type Options struct {
	Flights Flights
	Signer  Signer
	Flags   flags.Flags
}

// Issuer issues boarding passes.
type Issuer struct{ opts Options }

// NewIssuer returns an issuer.
func NewIssuer(opts Options) *Issuer { return &Issuer{opts: opts} }

// Issue builds the pass of a request.
func (i *Issuer) Issue(ctx context.Context, r Request) (BoardingPass, error) {
	p := BoardingPass{Booking: r.Booking, Passenger: r.Passenger, Flight: r.Flight, Seat: r.Seat, Group: r.Group}
	st, err := i.opts.Flights.Get(ctx, r.Flight, r.Day)
	if err != nil {
		return BoardingPass{}, err
	}
	p.Gate = st.Gate
	p.Boarding = st.Departure.Add(-40 * 60e9).Format("15:04")
	return p, nil
}

// Handler serves POST /v1/passes.
func (i *Issuer) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "the body is not a pass request", http.StatusBadRequest)
			return
		}
		p, err := i.Issue(r.Context(), req)
		if err != nil {
			http.Error(w, "the pass cannot be issued now", http.StatusBadGateway)
			return
		}
		signed, err := i.opts.Signer.Sign(r.Context(), p)
		if err != nil {
			http.Error(w, "the pass cannot be signed now", http.StatusBadGateway)
			return
		}
		w.Header().Set("Content-Type", "application/vnd.apple.pkpass")
		_, _ = w.Write(signed)
	})
}
