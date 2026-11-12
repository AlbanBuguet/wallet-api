package pass

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SignerClient calls wallet-pass-signer, which signs with a key of AWS KMS
// that never leaves it (ADR-0001).
type SignerClient struct {
	base string
	http *http.Client
}

// NewSignerClient returns a client of wallet-pass-signer at base.
func NewSignerClient(base string) *SignerClient {
	return &SignerClient{base: base, http: &http.Client{Timeout: 3 * time.Second}}
}

// Sign returns the signed pass.
func (c *SignerClient) Sign(ctx context.Context, p BoardingPass) ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/sign", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wallet-pass-signer: %s", resp.Status)
	}
	return io.ReadAll(resp.Body)
}
