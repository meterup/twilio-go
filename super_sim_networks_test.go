package twilio

import (
	"context"
	"net/url"
	"testing"
	"time"
)

func TestGetNetworkPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	page, err := envClient.SuperSim.Networks.GetNetworkPage(ctx, url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	assertPageExpected(t, page.Meta)
}

func TestGetNAPPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	page, err := envClient.SuperSim.Networks.GetNetworkAccessProfilePage(ctx, url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	assertPageExpected(t, page.Meta)
}
