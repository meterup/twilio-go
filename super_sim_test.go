package twilio

import (
	"context"
	"net/url"
	"testing"
	"time"
)

func assertSuperSimExpected(t *testing.T, sim *SuperSim) {
	if sim == nil {
		t.Fatal("SuperSim unexpectedly nil")
	}
	if sim.Sid == "" {
		t.Error("expected Sid to be populated")
	}
	if sim.Status == "" {
		t.Error("expected Status to be populated")
	}
	if !sim.DateCreated.Valid {
		t.Error("expected DateCreated to be valid")
	}
	if !sim.DateUpdated.Valid {
		t.Error("expected DateUpdated to be valid")
	}
	if sim.AccountSid == "" {
		t.Error("expected AccountSid to be populated")
	}
	if sim.Iccid == "" {
		t.Error("expected Iccid to be populated")
	}
	if sim.Url == "" {
		t.Error("expected Url to be populated")
	}
}

func assertPageExpected(t *testing.T, meta Meta) {
	if meta.FirstPageURL == "" {
		t.Error("expected FirstPageURL to be populated")
	}
	if meta.Key == "" {
		t.Error("expected Key to be populated")
	}
	if meta.PageSize == 0 {
		t.Error("expected default PageSize, got 0")
	}
}

func TestGetSuperSim(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	sid := "HS08d349f2f43fe4ac045905cbfd4e04e6"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	superSim, err := envClient.SuperSim.SuperSims.Get(ctx, sid)
	if err != nil {
		t.Fatal(err)
	}
	assertSuperSimExpected(t, superSim)
	if superSim.Sid != sid {
		t.Errorf("expected Sid to equal %s, got %s", sid, superSim.Sid)
	}
}

func TestGetSuperSimPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	page, err := envClient.SuperSim.SuperSims.GetPage(ctx, url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	assertPageExpected(t, page.Meta)
}

func TestGetFleetPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	page, err := envClient.SuperSim.SuperSims.GetFleetPage(ctx, url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	assertPageExpected(t, page.Meta)
}

func TestGetUsageRecordPage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping HTTP request in short mode")
	}
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	page, err := envClient.SuperSim.SuperSims.GetUsageRecordPage(ctx, url.Values{})
	if err != nil {
		t.Fatal(err)
	}
	assertPageExpected(t, page.Meta)
}
