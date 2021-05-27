package twilio

import (
	"context"
	"net/url"
)

const networkPathPart = "Networks"
const napPathPart = "NetworkAccessProfiles"

type NetworkService struct {
	client *Client
}

type Network struct {
	Sid          string        `json:"sid"`
	Url          string        `json:"url"`
	FriendlyName string        `json:"friendly_name"`
	IsoCountry   string        `json:"iso_country"`
	Identifiers  []interface{} `json:"identifiers"`
}

// NetworkPage represents a page of Networks.
type NetworkPage struct {
	Meta   Meta       `json:"meta"`
	Fleets []*Network `json:"networks"`
}

type networkPageIterator struct {
	p *PageIterator
}

// Get finds a single Network by its sid, or returns an error.
func (s *NetworkService) GetNetwork(ctx context.Context, sid string) (*Network, error) {
	network := new(Network)
	err := s.client.GetResource(ctx, networkPathPart, sid, network)
	return network, err
}

// GetPage returns a single Page of Networks, filtered by data.
//
// See https://www.twilio.com/docs/iot/supersim/api/network-resource?code-sample=code-fetch-a-network-resource.
func (s *NetworkService) GetNetworkPage(ctx context.Context, data url.Values) (*NetworkPage, error) {
	return s.GetNetworkPageIterator(data).Next(ctx)
}

// GetPageIterator returns a networkPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *NetworkService) GetNetworkPageIterator(data url.Values) *networkPageIterator {
	iter := NewPageIterator(s.client, data, networkPathPart)
	return &networkPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *networkPageIterator) Next(ctx context.Context) (*NetworkPage, error) {
	ap := new(NetworkPage)
	err := s.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(ap.Meta.NextPageURL)
	return ap, nil
}

type NetworkAccessProfile struct {
	Sid         string            `json:"sid"`
	UniqueName  string            `json:"unique_name"`
	AccountSid  string            `json:"account_sid"`
	DateCreated TwilioTime        `json:"date_created"`
	DateUpdated TwilioTime        `json:"date_updated"`
	Url         string            `json:"url"`
	Links       map[string]string `json:"links"`
}

// NAPPage represents a page of Network Access Profiles.
type NAPPage struct {
	Meta                  Meta                    `json:"meta"`
	NetworkAccessProfiles []*NetworkAccessProfile `json:"networks"`
}

type NAPPageIterator struct {
	p *PageIterator
}

// Create creates a new NetworkAccessProfile with the data provided, or returns an error.
func (s *NetworkService) CreateNetworkAccessProfile(ctx context.Context, data url.Values) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.CreateResource(ctx, napPathPart, data, networkAccessProfile)
	return networkAccessProfile, err
}

// Get finds a single NetworkAccessProfile by its sid, or returns an error.
func (s *NetworkService) GetNetworkAccessProfile(ctx context.Context, sid string) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.GetResource(ctx, napPathPart, sid, networkAccessProfile)
	return networkAccessProfile, err
}

// Update updates the specified NetworkAccessProfile with the data provided, or returns an error.
func (s *NetworkService) UpdateNetworkAccessProfile(ctx context.Context, sid string, data url.Values) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.UpdateResource(ctx, napPathPart, sid, data, networkAccessProfile)
	return networkAccessProfile, err
}

// GetPage returns a single Page of NetworkAccessProfiles, filtered by data.
//
// See https://www.twilio.com/docs/iot/supersim/api/networkaccessprofile-resource#read-multiple-networkaccessprofile-resources.
func (s *NetworkService) GetNetworkAccessProfilePage(ctx context.Context, data url.Values) (*NAPPage, error) {
	return s.GetNAPPageIterator(data).Next(ctx)
}

// GetPageIterator returns a NAPPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *NetworkService) GetNAPPageIterator(data url.Values) *NAPPageIterator {
	iter := NewPageIterator(s.client, data, networkPathPart)
	return &NAPPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *NAPPageIterator) Next(ctx context.Context) (*NAPPage, error) {
	ap := new(NAPPage)
	err := s.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(ap.Meta.NextPageURL)
	return ap, nil
}

type NAPNetwork struct {
	Sid          string        `json:"sid"`
	NAPSid       string        `json:"network_access_profile_sid"`
	FriendlyName string        `json:"friendly_name"`
	ISOCountry   string        `json:"iso_country"`
	Identifiers  []interface{} `json:"identifiers"`
	Url          string        `json:"url"`
}

// NAPNetworkPage represents a page of NAPNetworks.
type NAPNetworkPage struct {
	Meta        Meta          `json:"meta"`
	NAPNetworks []*NAPNetwork `json:"networks"`
}

type NAPNetworkPageIterator struct {
	p *PageIterator
}

// Create creates a new NAP Network associated with the NetworkAccessProfile provided, or returns an error.
func (s *NetworkService) CreateNAPNetwork(ctx context.Context, napSid string, data url.Values) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.CreateResource(ctx, napPathPart+"/"+napSid, data, networkAccessProfile)
	return networkAccessProfile, err
}

// Get finds a single NAP Network, or returns an error.
func (s *NetworkService) GetNAPNetwork(ctx context.Context, napSid, sid string) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.GetResource(ctx, napPathPart+"/"+napSid, sid, networkAccessProfile)
	return networkAccessProfile, err
}

// Update updates a single NAP Network with the data provided, or returns an error.
func (s *NetworkService) UpdateNAPNetwork(ctx context.Context, napSid, sid string, data url.Values) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.UpdateResource(ctx, napPathPart+"/"+napSid, sid, data, networkAccessProfile)
	return networkAccessProfile, err
}

// Delete deletes the provided NAP Network, or returns an error.
func (s *NetworkService) DeleteNAPNetwork(ctx context.Context, napSid, sid string, data url.Values) error {
	return s.client.DeleteResource(ctx, napPathPart+"/"+napSid, sid)
}

// GetPage returns a single Page of NAPNetworks, filtered by data.
//
// See https://www.twilio.com/docs/iot/supersim/api/networkaccessprofile-resource/network-resource#read-multiple-networkaccessprofile-network-resources.
func (s *NetworkService) GetNAPNetworkPage(ctx context.Context, napSid string, data url.Values) (*NAPNetworkPage, error) {
	return s.GetNAPNetworkPageIterator(napPathPart+"/"+napSid, data).Next(ctx)
}

// GetPageIterator returns a NAPNetworkPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *NetworkService) GetNAPNetworkPageIterator(napSid string, data url.Values) *NAPNetworkPageIterator {
	iter := NewPageIterator(s.client, data, napPathPart+"/"+napSid)
	return &NAPNetworkPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *NAPNetworkPageIterator) Next(ctx context.Context) (*NAPNetworkPage, error) {
	ap := new(NAPNetworkPage)
	err := s.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(ap.Meta.NextPageURL)
	return ap, nil
}
