package twilio

import (
	"context"
	"net/url"
	"testing"
	"time"
)

func TestGetCommandPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	page, err := envClient.SuperSim.SuperCommands.GetCommandPage(ctx, url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	assertPageExpected(t, page.Meta)
}

func TestGetSMSCommandPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	page, err := envClient.SuperSim.SuperCommands.GetSMSCommandPage(ctx, url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	assertPageExpected(t, page.Meta)
}