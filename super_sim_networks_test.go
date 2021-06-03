package twilio

import (
	"context"
	"net/url"
	"testing"
)

func assertNetworkExpected(t *testing.T, network *Network) {
	if network == nil {
		t.Fatal("unexpected nil network")
	}
	if network.FriendlyName == "" {
		t.Error("expected FriendlyName to be populated")
	}
	if network.IsoCountry == "" {
		t.Error("expected IsoCountry to be populated")
	}
	if network.Identifiers == nil {
		t.Error("expected Identifiers to be populated")
	}
	if network.Sid == "" {
		t.Error("expected Sid to be populated")
	}
	if network.Url == "" {
		t.Error("expected Url to be populated")
	}
}

func assertNetworkAccessProfileExpected(t *testing.T, networkAccessProfile *NetworkAccessProfile) {
	if networkAccessProfile == nil {
		t.Fatal("unexpected nil network")
	}
	if networkAccessProfile.Sid == "" {
		t.Error("expected Sid to be populated")
	}
	if networkAccessProfile.UniqueName == "" {
		t.Error("expected UniqueName to be populated")
	}
	if networkAccessProfile.AccountSid == "" {
		t.Error("expected AccountSid to be populated")
	}
	if !networkAccessProfile.DateCreated.Valid {
		t.Error("expected DateCreated to be valid")
	}
	if !networkAccessProfile.DateUpdated.Valid {
		t.Error("expected DateUpdated to be populated")
	}
	if networkAccessProfile.Url == "" {
		t.Error("expected Url to be populated")
	}
	if networkAccessProfile.Links == nil {
		t.Error("expected Links to be populated")
	}
}

func assertNAPNetworkExpected(t *testing.T, napNetwork *NAPNetwork) {
	if napNetwork == nil {
		t.Fatal("unexpected nil network")
	}
	if napNetwork.Sid == "" {
		t.Error("expected Sid to be populated")
	}
	if napNetwork.FriendlyName == "" {
		t.Error("expected FriendlyName to be populated")
	}
	if napNetwork.NetworkAccessProfileSid == "" {
		t.Error("expected NetworkAccessProfileSid to be populated")
	}
	if napNetwork.ISOCountry == "" {
		t.Error("expected ISOCountry to be populated")
	}
	if napNetwork.Url == "" {
		t.Error("expected Url to be populated")
	}
	if napNetwork.Identifiers == nil {
		t.Error("expected Links to be populated")
	}
}

func TestNetwork(t *testing.T) {
	t.Parallel()

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(networkResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		network, err := client.SuperSim.Networks.GetNetwork(ctx, "89883070000123456789")
		if err != nil {
			t.Fatal(err)
		}
		assertNetworkExpected(t, network)
	})

	t.Run("GetPage", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(networkPageResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		resp, err := client.SuperSim.Networks.GetNetworkPage(ctx, url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.Networks) != 2 {
			t.Errorf("expected 2 networks, found %d", len(resp.Networks))
		}
		for _, network := range resp.Networks {
			assertNetworkExpected(t, network)
		}
		assertPageExpected(t, resp.Meta)
	})
}

func TestNetworkAccessProfile(t *testing.T) {
	t.Parallel()

	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(networkAccessProfileResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		networkAccessProfile, err := client.SuperSim.Networks.CreateNetworkAccessProfile(ctx, url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		assertNetworkAccessProfileExpected(t, networkAccessProfile)
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(networkAccessProfileResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		networkAccessProfile, err := client.SuperSim.Networks.GetNetworkAccessProfile(ctx, "89883070000123456789")
		if err != nil {
			t.Fatal(err)
		}
		assertNetworkAccessProfileExpected(t, networkAccessProfile)
	})

	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(networkAccessProfileResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		networkAccessProfile, err := client.SuperSim.Networks.UpdateNetworkAccessProfile(
			ctx, "89883070000123456789", url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		assertNetworkAccessProfileExpected(t, networkAccessProfile)
	})

	t.Run("GetPage", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(networkAccessProfilePageResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		resp, err := client.SuperSim.Networks.GetNetworkAccessProfilePage(ctx, url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.NetworkAccessProfiles) != 2 {
			t.Errorf("expected 2 networks, found %d", len(resp.NetworkAccessProfiles))
		}
		for _, networkAccessProfile := range resp.NetworkAccessProfiles {
			assertNetworkAccessProfileExpected(t, networkAccessProfile)
		}
		assertPageExpected(t, resp.Meta)
	})
}

func TestNAPNetwork(t *testing.T) {
	t.Parallel()

	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(napNetworkResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		napNetwork, err := client.SuperSim.Networks.CreateNAPNetwork(
			ctx, "HAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		assertNAPNetworkExpected(t, napNetwork)
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(napNetworkResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		napNetwork, err := client.SuperSim.Networks.GetNAPNetwork(
			ctx, "HAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "HWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		if err != nil {
			t.Fatal(err)
		}
		assertNAPNetworkExpected(t, napNetwork)
	})

	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(napNetworkResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		napNetwork, err := client.SuperSim.Networks.UpdateNAPNetwork(
			ctx, "HAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "HWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		assertNAPNetworkExpected(t, napNetwork)
	})

	t.Run("GetPage", func(t *testing.T) {
		t.Parallel()
		client, s := getServer(napNetworkPageResponse)
		defer s.Close()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		resp, err := client.SuperSim.Networks.GetNAPNetworkPage(ctx, "HWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", url.Values{})
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.NAPNetworks) != 2 {
			t.Errorf("expected 2 networks, found %d", len(resp.NAPNetworks))
		}
		for _, napNetwork := range resp.NAPNetworks {
			assertNAPNetworkExpected(t, napNetwork)
		}
		assertPageExpected(t, resp.Meta)
	})
}
