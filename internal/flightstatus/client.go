// Package flightstatus reads the status of a flight from flight-status-api.
package flightstatus

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Status is what the wallet shows of a flight.
type Status struct {
	Flight    string    `json:"flight"`
	Gate      string    `json:"gate"`
	Departure time.Time `json:"departure"`
	Delay     int       `json:"delayMinutes"`
}

// Client calls flight-status-api.
type Client struct {
	base string
	http *http.Client
}

// New returns a client of flight-status-api at base.
func New(base string, c *http.Client) *Client { return &Client{base: base, http: c} }

// Get reads the status of a flight on a day.
func (c *Client) Get(ctx context.Context, flight, day string) (Status, error) {
	u := c.base + "/v2/flights/" + url.PathEscape(flight) + "?date=" + url.QueryEscape(day)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return Status{}, err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return Status{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Status{}, fmt.Errorf("flight-status-api: %s", resp.Status)
	}
	var s Status
	return s, json.NewDecoder(resp.Body).Decode(&s)
}
