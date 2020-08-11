package dns

import (
	"context"
	"errors"
	"os"

	"github.com/dnsimple/dnsimple-go/dnsimple"
)

// GetDNSimpleClient returns a token authenticated dnsimple client
func GetDNSimpleClient() (*dnsimple.Client, error) {
	token := os.Getenv("DNSIMPLE_API_TOKEN")
	if token == "" {
		return nil, errors.New("NOT TOKEN AVAILABLE")
	}
	tc := dnsimple.StaticTokenHTTPClient(context.Background(), token)
	client := dnsimple.NewClient(tc)
	return client, nil
}
