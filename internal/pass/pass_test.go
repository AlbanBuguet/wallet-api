package pass

import (
	"context"
	"testing"
	"time"

	"github.example.com/opco-a/wallet-api/internal/flags"
	"github.example.com/opco-a/wallet-api/internal/flightstatus"
)

type fakeFlights struct{}

func (fakeFlights) Get(context.Context, string, string) (flightstatus.Status, error) {
	return flightstatus.Status{Flight: "OA123", Gate: "B12", Departure: time.Date(2026, 11, 20, 9, 30, 0, 0, time.UTC)}, nil
}

func TestIssue(t *testing.T) {
	i := NewIssuer(Options{Flights: fakeFlights{}, Flags: flags.Flags{}})
	p, err := i.Issue(context.Background(), Request{Booking: "ABC123", Passenger: "A. Traveller", Flight: "OA123", Day: "2026-11-20", Seat: "12A", Group: "2"})
	if err != nil {
		t.Fatal(err)
	}
	if p.Gate != "B12" || p.Boarding != "08:50" {
		t.Fatalf("pass = %+v", p)
	}
}
