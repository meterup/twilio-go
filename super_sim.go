package twilio

import (
	"context"
	"net/url"
)

const superSimPathPart = "Sims"
const fleetPathPart = "Fleets"
const usageRecordPathPart = "UsageRecords"

type SuperSimService struct {
	client *Client
}

type SuperSim struct {
	Sid         string     `json:"sid"`
	Status      string     `json:"status"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
	AccountSid  string     `json:"account_sid"`
	UniqueName  string     `json:"unique_name"`
	Iccid       string     `json:"iccid"`
	FleetSid    string     `json:"fleet_sid"`
	Url         string     `json:"url"`
}

// SuperSimPage represents a page of SuperSims.
type SuperSimPage struct {
	Meta      Meta        `json:"meta"`
	SuperSims []*SuperSim `json:"sims"`
}

type superSimPageIterator struct {
	p *PageIterator
}

// Register registers a new SIM with the provided account.
//
// See https://www.twilio.com/docs/iot/supersim/api/sim-resource#add-a-super-sim-to-your-account
func (s *SuperSimService) Register(ctx context.Context, iccid, registrationCode string) (*SuperSim, error) {
	superSim := new(SuperSim)
	data := url.Values{}
	data.Set("Iccid", iccid)
	data.Set("RegistrationCode", registrationCode)
	err := s.client.CreateResource(ctx, superSimPathPart, data, superSim)
	return superSim, err
}

// Get finds a single SuperSim resource by its sid or unique name, or returns an error.
func (s *SuperSimService) Get(ctx context.Context, sidOrUniqueName string) (*SuperSim, error) {
	superSim := new(SuperSim)
	err := s.client.GetResource(ctx, superSimPathPart, sidOrUniqueName, superSim)
	return superSim, err
}

// Activate sets the status of the SIM provided to "active".
//
// See https://www.twilio.com/docs/iot/supersim/api/sim-resource#update-a-sim-resource
func (s *SuperSimService) Activate(ctx context.Context, sid string) (*SuperSim, error) {
	superSim := new(SuperSim)
	data := url.Values{}
	data.Set("Status", "active")
	err := s.client.UpdateResource(ctx, superSimPathPart, sid, data, superSim)
	return superSim, err
}

// Update updates the specified SuperSim resource with the data provided, or returns an error.
func (s *SuperSimService) Update(ctx context.Context, sid string, data url.Values) (*SuperSim, error) {
	superSim := new(SuperSim)
	err := s.client.UpdateResource(ctx, superSimPathPart, sid, data, superSim)
	return superSim, err
}

// GetPage returns a single Page of SuperSims, filtered by data.
//
// See https://www.twilio.com/docs/iot/supersim/api/sim-resource#read-multiple-sim-resources.
func (s *SuperSimService) GetPage(ctx context.Context, data url.Values) (*SuperSimPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

// GetPageIterator returns a superSimPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetPageIterator(data url.Values) *superSimPageIterator {
	iter := NewPageIterator(s.client, data, superSimPathPart)
	return &superSimPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *superSimPageIterator) Next(ctx context.Context) (*SuperSimPage, error) {
	ap := new(SuperSimPage)
	err := s.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(ap.Meta.NextPageURL)
	return ap, nil
}

type Fleet struct {
	Sid                     string     `json:"sid"`
	Url                     string     `json:"url"`
	AccountSid              string     `json:"account_sid"`
	UniqueName              string     `json:"unique_name""`
	DataEnabled             bool       `json:"data_enabled"`
	DataLimit               int64      `json:"data_limit"`
	DataMetering            string     `json:"data_metering"`
	DateCreated             TwilioTime `json:"date_created"`
	DateUpdated             TwilioTime `json:"date_updated"`
	CommandsEnabled         bool       `json:"commands_enabled"`
	CommandsUrl             string     `json:"commands_url"`
	CommandsMethod          string     `json:"commands_method"`
	SmsCommandsEnabled      bool       `json:"sms_commands_enabled"`
	SmsCommandsMethod       string     `json:"sms_commands_method"`
	IPCommandsMethod        string     `json:"ip_commands_method"`
	IPCommandsUrl           string     `json:"ip_commands_url"`
	NetworkAccessProfileSid string     `json:"network_access_profile_sid"`
}

// FleetPage represents a page of Fleets.
type FleetPage struct {
	Meta   Meta     `json:"meta"`
	Fleets []*Fleet `json:"fleets"`
}

type fleetPageIterator struct {
	p *PageIterator
}

// Create creates a new SuperSim Fleet with the data provided, or returns an error.
func (s *SuperSimService) CreateFleet(ctx context.Context, data url.Values) (*Fleet, error) {
	fleet := new(Fleet)
	err := s.client.CreateResource(ctx, fleetPathPart, data, fleet)
	return fleet, err
}

// Get finds a single SuperSim Fleet by its sid, or returns an error.
func (s *SuperSimService) GetFleet(ctx context.Context, sid string) (*Fleet, error) {
	fleet := new(Fleet)
	err := s.client.GetResource(ctx, fleetPathPart, sid, fleet)
	return fleet, err
}

// GetPage returns a single Page of fleets, filtered by data.
//
// See https://www.twilio.com/docs/iot/supersim/api/fleet-resource#read-multiple-fleet-resources.
func (s *SuperSimService) GetFleetPage(ctx context.Context, data url.Values) (*FleetPage, error) {
	return s.GetFleetPageIterator(data).Next(ctx)
}

// GetFleetPageIterator returns a fleetPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetFleetPageIterator(data url.Values) *fleetPageIterator {
	iter := NewPageIterator(s.client, data, fleetPathPart)
	return &fleetPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *fleetPageIterator) Next(ctx context.Context) (*FleetPage, error) {
	ap := new(FleetPage)
	err := s.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(ap.Meta.NextPageURL)
	return ap, nil
}

type UsageRecord struct {
	AccountSid string `json:"account_sid"`
	SimSid     string `json:"sim_sid"`
	FleetSid   string `json:"fleet_sid"`
	NetworkSid string `json:"network_sid"`
	// Total data uploaded in bytes, aggregated by the query parameters
	DataUpload int64 `json:"data_upload"`
	// Total data downloaded in bytes, aggregated by the query parameters
	DataDownload int64 `json:"data_download"`
	// Total of data_upload and data_download (in bytes).
	DataTotal int64 `json:"data_total"`
	// Alpha-2 ISO Country Code
	IsoCountry string              `json:"iso_country"`
	Period     SuperSimUsagePeriod `json:"period"`
}

type SuperSimUsagePeriod struct {
	Start TwilioTime `json:"start_time"`
	End   TwilioTime `json:"end_time"`
}

// UsageRecordPage represents a page of UsageRecords.
type UsageRecordPage struct {
	Meta         Meta           `json:"meta"`
	UsageRecords []*UsageRecord `json:"usage_records"`
}

type UsageRecordPageIterator struct {
	p *PageIterator
}

// GetPage returns a single Page of UsageRecords, filtered by data.
//
// See https://www.twilio.com/docs/iot/supersim/api/usage-record-resource#read-usagerecord-resources.
func (s *SuperSimService) GetUsageRecordPage(ctx context.Context, data url.Values) (*UsageRecordPage, error) {
	return s.GetUsageRecordPageIterator(data).Next(ctx)
}

// GetPageIterator returns a UsageRecordPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetUsageRecordPageIterator(data url.Values) *UsageRecordPageIterator {
	iter := NewPageIterator(s.client, data, usageRecordPathPart)
	return &UsageRecordPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *UsageRecordPageIterator) Next(ctx context.Context) (*UsageRecordPage, error) {
	ap := new(UsageRecordPage)
	err := s.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(ap.Meta.NextPageURL)
	return ap, nil
}
