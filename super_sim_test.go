package twilio

import (
	"context"
	"net/url"
	"testing"
)

func assertSuperSimExpected(t *testing.T, sim *SuperSim) {
	if sim == nil {
		t.Fatal("SuperSim unexpectedly nil")
	}
	if sim.Sid == "" {
		t.Error("expected Sid to be populated")
	}
	if sim.FleetSid == "" {
		t.Error("expected FleetSid to be populated")
	}
	if sim.UniqueName == "" {
		t.Error("expected UniqueName to be populated")
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

func assertFleetExpected(t *testing.T, fleet *Fleet) {
	if fleet.AccountSid == "" {
		t.Error("expected AccountSid to be populated")
	}
	if fleet.UniqueName == "" {
		t.Error("expected AccountSid to be populated")
	}
	if !fleet.DataEnabled {
		t.Error("expected DataEnabled to be true")
	}
	if fleet.DataLimit == 0 {
		t.Error("expected DataLimit to be populated")
	}
	if fleet.DataMetering == "" {
		t.Error("expected DataMetering to be populated")
	}
	if !fleet.DateCreated.Valid {
		t.Error("expected DateCreated to be valid")
	}
	if !fleet.DateUpdated.Valid {
		t.Error("expected DateUpdated to be valid")
	}
	if !fleet.CommandsEnabled {
		t.Error("expected CommandsEnabled to be true")
	}
	if fleet.CommandsMethod == "" {
		t.Error("expected CommandsMethod to be present")
	}
	if fleet.CommandsUrl == "" {
		t.Error("expected CommandsUrl to be present")
	}
	if !fleet.SmsCommandsEnabled {
		t.Error("expected SmsCommandsEnabled to be true")
	}
	if fleet.SmsCommandsMethod == "" {
		t.Error("expected SmsCommandsMethod to be present")
	}
	if fleet.IPCommandsUrl == "" {
		t.Error("expected IPCommandsUrl to be present")
	}
	if fleet.IPCommandsMethod == "" {
		t.Error("expected IPCommandsMethod to be present")
	}
	if fleet.NetworkAccessProfileSid == "" {
		t.Error("expected NetworkAccessProfileSid to be present")
	}
	if fleet.Sid == "" {
		t.Error("expected Sid to be present")
	}
	if fleet.Url == "" {
		t.Error("expected Sid to be present")
	}
}

func assertUsageRecordExpected(t *testing.T, usageRecord *UsageRecord) {
	if usageRecord == nil {
		t.Fatal("SuperSim unexpectedly nil")
	}
	if usageRecord.SimSid == "" {
		t.Error("expected SimSid to be populated")
	}
	if usageRecord.AccountSid == "" {
		t.Error("expected AccountSid to be populated")
	}
	if usageRecord.FleetSid == "" {
		t.Error("expected FleetSid to be populated")
	}
	if usageRecord.NetworkSid == "" {
		t.Error("expected NetworkSid to be populated")
	}
	if usageRecord.IsoCountry == "" {
		t.Error("expected IsoCountry to be populated")
	}
	if usageRecord.FleetSid == "" {
		t.Error("expected FleetSid to be populated")
	}
	if !usageRecord.Period.Start.Valid {
		t.Error("expected Period Start time to be populated")
	}
	if !usageRecord.Period.End.Valid {
		t.Error("expected Period End time to be populated")
	}
	if usageRecord.DataDownload == 0 {
		t.Error("expected DataDownload to be populated")
	}
	if usageRecord.DataUpload == 0 {
		t.Error("expected DataUpload to be populated")
	}
	if usageRecord.DataTotal == 0 {
		t.Error("expected DataTotal to be populated")
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

func TestSuperSim(t *testing.T) {
	t.Parallel()

	t.Run("Register", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(superSimResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		superSim, err := client.SuperSim.SuperSims.Register(ctx, "89883070000123456789", "code")
		if err != nil {
			t.Fatal(err)
		}
		assertSuperSimExpected(t, superSim)
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(superSimResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		superSim, err := client.SuperSim.SuperSims.Get(ctx, "HFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		if err != nil {
			t.Fatal(err)
		}
		assertSuperSimExpected(t, superSim)
	})

	t.Run("Activate", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(superSimResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		superSim, err := client.SuperSim.SuperSims.Activate(ctx, "HFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		if err != nil {
			t.Fatal(err)
		}
		assertSuperSimExpected(t, superSim)
	})

	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(superSimResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		data := url.Values{}
		data.Set("Status", "inactive")
		superSim, err := client.SuperSim.SuperSims.Update(ctx, "HFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", data)
		if err != nil {
			t.Fatal(err)
		}
		assertSuperSimExpected(t, superSim)
	})

	t.Run("GetPage", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(superSimPageResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		data := url.Values{}
		data.Set("Status", "active")
		resp, err := client.SuperSim.SuperSims.GetPage(ctx, url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.SuperSims) != 5 {
			t.Error("expected 5 SuperSims to be returned")
		}
		assertPageExpected(t, resp.Meta)
		for _, sim := range resp.SuperSims {
			assertSuperSimExpected(t, sim)
		}
	})
}

func TestFleet(t *testing.T) {
	t.Parallel()

	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(fleetResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		fleet, err := client.SuperSim.SuperSims.CreateFleet(ctx, url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		assertFleetExpected(t, fleet)
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(fleetResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		fleet, err := client.SuperSim.SuperSims.GetFleet(ctx, "HFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		if err != nil {
			t.Fatal(err)
		}
		assertFleetExpected(t, fleet)
	})

	t.Run("GetPage", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(fleetPageResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		resp, err := client.SuperSim.SuperSims.GetFleetPage(ctx, url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.Fleets) != 5 {
			t.Errorf("expected 5 fleets, get %d", len(resp.Fleets))
		}
		for _, fleet := range resp.Fleets {
			assertFleetExpected(t, fleet)
		}
		assertPageExpected(t, resp.Meta)
	})
}

func TestUsageRecords(t *testing.T) {
	t.Parallel()
	t.Run("GetPage", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(usageRecordResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		resp, err := client.SuperSim.SuperSims.GetUsageRecordPage(ctx, url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.UsageRecords) != 2 {
			t.Errorf("expected 2 usage records, got %d", len(resp.UsageRecords))
		}
		for _, usageRecord := range resp.UsageRecords {
			assertUsageRecordExpected(t, usageRecord)
		}
		assertPageExpected(t, resp.Meta)
	})
}
