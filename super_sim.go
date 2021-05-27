package twilio

import (
	"context"
	"net/url"
)

const superSimPathPart = "Sims"
const superSimFleetPathPart = "Fleets"
const superSimNetworkPathPart = "Networks"
const superSimNAPPathPart = "NetworkAccessProfiles"
const superSimUsageRecordPathPart = "UsageRecords"
const superSimSMSCommandsPathPart = "SMSCommands"
const superSimCommandsPathPart = "Commands"

type SuperSimService struct {
	client *Client
}

type SuperSim struct {
	Sid         string     `json:"sid"`
	Status      string     `json:"status"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
	AccountSid  string     `json:"account_sid"`
	UniqueName string `json:"unique_name"`
	Iccid string `json:"iccid"`
	FleetSid string `json:"fleet_sid"`
	Url string `json:"url"`

}

// SuperSimPage represents a page of Alerts.
type SuperSimPage struct {
	Meta   Meta     `json:"meta"`
	SuperSims []*SuperSim `json:"sims"`
}

type superSimPageIterator struct {
	p *PageIterator
}

// Register registers a new SIM with the provided account.
//
// See https://www.twilio.com/docs/api/monitor/alerts#list-get-filters.
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

// Get finds a single SuperSim resource by its sid or unique name, or returns an error.
func (s *SuperSimService) Update(ctx context.Context, sidOrUniqueName string, data url.Values) (*SuperSim, error) {
	superSim := new(SuperSim)
	err := s.client.UpdateResource(ctx, superSimPathPart, sidOrUniqueName, data, superSim)
	return superSim, err
}

// GetPage returns a single Page of resources, filtered by data.
//
// See https://www.twilio.com/docs/api/monitor/alerts#list-get-filters.
func (s *SuperSimService) GetPage(ctx context.Context, data url.Values) (*SuperSimPage, error) {
	return s.GetPageIterator(data).Next(ctx)
}

// GetPageIterator returns a AlertPageIterator with the given page
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
	Sid string `json:"sid"`
	Url string `json:"url"`
	AccountSid string `json:"account_sid"`
	UniqueName string `json:"unique_name""`
	DataEnabled bool `json:"data_enabled"`
	DataLimit int `json:"data_limit"`
	DataMetering string `json:"data_metering"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
	CommandsEnabled bool `json:"commands_enabled"`
	CommandsUrl url.URL `json:"commands_url"`
	CommandsMethod string `json:"commands_method"`
	SmsCommandsEnabled bool `json:"sms_commands_enabled"`
	SmsCommandsMethod string `json:"sms_commands_method"`
	IPCommandsMethod string `json:"ip_commands_method"`
	IPCommandsUrl string `json:"ip_commands_url"`
	NetworkAccessProfileSid string `json:"network_access_profile_sid"`
}

// FleetPage represents a page of Fleet.
type FleetPage struct {
	Meta   Meta     `json:"meta"`
	Fleets []*Fleet `json:"fleets"`
}

type fleetPageIterator struct {
	p *PageIterator
}

// Get finds a single SuperSim Fleet by its sid or unique name, or returns an error.
func (s *SuperSimService) CreateFleet(ctx context.Context, data url.Values) (*Fleet, error) {
	fleet := new(Fleet)
	err := s.client.CreateResource(ctx, superSimFleetPathPart, data, fleet)
	return fleet, err
}

// Get finds a single SuperSim Fleet by its sid or unique name, or returns an error.
func (s *SuperSimService) GetFleet(ctx context.Context, sidOrUniqueName string) (*Fleet, error) {
	fleet := new(Fleet)
	err := s.client.GetResource(ctx, superSimFleetPathPart, sidOrUniqueName, fleet)
	return fleet, err
}

// GetPage returns a single Page of resources, filtered by data.
//
// See https://www.twilio.com/docs/api/monitor/alerts#list-get-filters.
func (s *SuperSimService) GetFleetPage(ctx context.Context, data url.Values) (*FleetPage, error) {
	return s.GetFleetPageIterator(data).Next(ctx)
}

// GetPageIterator returns a AlertPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetFleetPageIterator(data url.Values) *fleetPageIterator {
	iter := NewPageIterator(s.client, data, superSimFleetPathPart)
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

type Network struct {
	Sid string `json:"sid"`
	Url string `json:"url"`
	FriendlyName string `json:"friendly_name"`
	IsoCountry string `json:iso_country`
	Identifiers []interface{} `json:identifiers`
}

// FleetPage represents a page of Fleet.
type NetworkPage struct {
	Meta   Meta     `json:"meta"`
	Fleets []*Network `json:"networks"`
}

type networkPageIterator struct {
	p *PageIterator
}

// Get finds a single SuperSim Fleet by its sid or unique name, or returns an error.
func (s *SuperSimService) GetNetwork(ctx context.Context, sid string) (*Network, error) {
	network := new(Network)
	err := s.client.GetResource(ctx, superSimNetworkPathPart, sid, network)
	return network, err
}

// GetPage returns a single Page of resources, filtered by data.
//
// See https://www.twilio.com/docs/api/monitor/alerts#list-get-filters.
func (s *SuperSimService) GetNetworkPage(ctx context.Context, data url.Values) (*NetworkPage, error) {
	return s.GetNetworkPageIterator(data).Next(ctx)
}

// GetPageIterator returns a AlertPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetNetworkPageIterator(data url.Values) *networkPageIterator {
	iter := NewPageIterator(s.client, data, superSimNetworkPathPart)
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
	Sid string `json:"sid"`
	UniqueName string `json:"unique_name"`
	AccountSid string `json:"account_sid"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
	Url string `json:"url"`
	Links map[string]string `json:"links"`
}

// NAPPage represents a page of Network Access Profiles.
type NAPPage struct {
	Meta   Meta     `json:"meta"`
	NetworkAccessProfiles []*NetworkAccessProfile `json:"networks"`
}

type NAPPageIterator struct {
	p *PageIterator
}

// Get finds a single SuperSim Fleet by its sid or unique name, or returns an error.
func (s *SuperSimService) CreateNetworkAccessProfile(ctx context.Context, data url.Values) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.CreateResource(ctx, superSimNAPPathPart, data, networkAccessProfile)
	return networkAccessProfile, err
}

// Get finds a single SuperSim Fleet by its sid or unique name, or returns an error.
func (s *SuperSimService) GetNetworkAccessProfile(ctx context.Context, sid string) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.GetResource(ctx, superSimNAPPathPart, sid, networkAccessProfile)
	return networkAccessProfile, err
}

// Get finds a single SuperSim Fleet by its sid or unique name, or returns an error.
func (s *SuperSimService) UpdateNetworkAccessProfile(ctx context.Context, sid string, data url.Values) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.UpdateResource(ctx, superSimNAPPathPart, sid, data, networkAccessProfile)
	return networkAccessProfile, err
}

// GetPage returns a single Page of resources, filtered by data.
//
// See https://www.twilio.com/docs/api/monitor/alerts#list-get-filters.
func (s *SuperSimService) GetNetworkAccessProfilePage(ctx context.Context, data url.Values) (*NAPPage, error) {
	return s.GetNAPPageIterator(data).Next(ctx)
}

// GetPageIterator returns a AlertPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetNAPPageIterator(data url.Values) *NAPPageIterator {
	iter := NewPageIterator(s.client, data, superSimNetworkPathPart)
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
	Sid string `json:"sid"`
	NAPSid string `json:"network_access_profile_sid"`
	FriendlyName string `json:"friendly_name"`
	ISOCountry string `json:"iso_country"`
	Identifiers []interface{} `json:identifiers`
	Url string `json:"url"`
}

// NAPNetworkPage represents a page of Network Access Profiles.
type NAPNetworkPage struct {
	Meta   Meta     `json:"meta"`
	NAPNetworks []*NAPNetwork `json:"networks"`
}

type NAPNetworkPageIterator struct {
	p *PageIterator
}

// Get finds a single SuperSim NAP Network by its sid or unique name, or returns an error.
func (s *SuperSimService) CreateNAPNetwork(ctx context.Context, napSid string, data url.Values) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.CreateResource(ctx, superSimNAPPathPart + "/" + napSid, data, networkAccessProfile)
	return networkAccessProfile, err
}

// Get finds a single SuperSim NAP Network by its sid or unique name, or returns an error.
func (s *SuperSimService) GetNAPNetwork(ctx context.Context, napSid, sid string) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.GetResource(ctx, superSimNAPPathPart + "/" + napSid, sid, networkAccessProfile)
	return networkAccessProfile, err
}

// Get finds a single SuperSim NAP Network by its sid or unique name, or returns an error.
func (s *SuperSimService) UpdateNAPNetwork(ctx context.Context, napSid, sid string, data url.Values) (*NetworkAccessProfile, error) {
	networkAccessProfile := new(NetworkAccessProfile)
	err := s.client.UpdateResource(ctx, superSimNAPPathPart + "/" + napSid, sid, data, networkAccessProfile)
	return networkAccessProfile, err
}

// Get finds a single SuperSim NAP Network by its sid or unique name, or returns an error.
func (s *SuperSimService) DeleteNAPNetwork(ctx context.Context, napSid, sid string, data url.Values) error {
	return s.client.DeleteResource(ctx, superSimNAPPathPart + "/" + napSid, sid)
}

// GetPage returns a single Page of resources, filtered by data.
//
// See https://www.twilio.com/docs/api/monitor/alerts#list-get-filters.
func (s *SuperSimService) GetNAPNetworkPage(ctx context.Context, napSid string, data url.Values) (*NAPNetworkPage, error) {
	return s.GetNAPNetworkPageIterator(superSimNAPPathPart + "/" + napSid, data).Next(ctx)
}

// GetPageIterator returns a AlertPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetNAPNetworkPageIterator(napSid string, data url.Values) *NAPNetworkPageIterator {
	iter := NewPageIterator(s.client, data, superSimNAPPathPart + "/" + napSid)
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

type UsageRecord struct {
	AccountSid string        `json:"account_sid"`
	SimSid     string        `json:"sim_sid"`
	Commands   CommandsUsage `json:"commands"`
	Data       AllDataUsage  `json:"data"`
	Period     UsagePeriod   `json:"period"`
}

// NAPNetworkPage represents a page of Network Access Profiles.
type UsageRecordPage struct {
	Meta   Meta     `json:"meta"`
	UsageRecords []*UsageRecord `json:"usage_records"`
}

type UsageRecordPageIterator struct {
	p *PageIterator
}

// Get finds a single SuperSim UsageRecords by its sid or unique name, or returns an error.
func (s *SuperSimService) GetUsageRecord(ctx context.Context, sid string) (*UsageRecord, error) {
	usageRecord := new(UsageRecord)
	err := s.client.GetResource(ctx, superSimUsageRecordPathPart, sid, usageRecord)
	return usageRecord, err
}

// GetPage returns a single Page of resources, filtered by data.
//
// See https://www.twilio.com/docs/api/monitor/alerts#list-get-filters.
func (s *SuperSimService) GetUsageRecordPage(ctx context.Context, napSid string, data url.Values) (*UsageRecordPage, error) {
	return s.GetUsageRecordPageIterator(superSimUsageRecordPathPart, data).Next(ctx)
}

// GetPageIterator returns a AlertPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetUsageRecordPageIterator(napSid string, data url.Values) *UsageRecordPageIterator {
	iter := NewPageIterator(s.client, data, superSimUsageRecordPathPart)
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

type SMSCommand struct {
	AccountSid string `json:"account_sid"`
	Payload string `json:"payload"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
	SimSid string `json:"sim_sid"`
	Status string `json:"status"`
	Sid string `json:"sid"`
	Direction string `json:"direction"`
	Url string `json:"url"`
}

// NAPNetworkPage represents a page of Network Access Profiles.
type SMSCommandPage struct {
	Meta   Meta     `json:"meta"`
	SMSCommands []*SMSCommand `json:"usage_records"`
}

type SMSCommandPageIterator struct {
	p *PageIterator
}

// Get finds a single SuperSim NAP Network by its sid or unique name, or returns an error.
func (s *SuperSimService) CreateSMSCommand(ctx context.Context, napSid string, data url.Values) (*SMSCommand, error) {
	smsCommand := new(SMSCommand)
	err := s.client.CreateResource(ctx, superSimCommandsPathPart, data, smsCommand)
	return smsCommand, err
}

// Get finds a single SuperSim NAP Network by its sid or unique name, or returns an error.
func (s *SuperSimService) GetSMSCommand(ctx context.Context, sid string) (*SMSCommand, error) {
	smsCommand := new(SMSCommand)
	err := s.client.GetResource(ctx, superSimCommandsPathPart, sid, smsCommand)
	return smsCommand, err
}

// GetPage returns a single Page of resources, filtered by data.
//
// See https://www.twilio.com/docs/api/monitor/alerts#list-get-filters.
func (s *SuperSimService) GetSMSCommandPage(ctx context.Context, napSid string, data url.Values) (*SMSCommandPage, error) {
	return s.GetSMSCommandPageIterator(superSimUsageRecordPathPart, data).Next(ctx)
}

// GetPageIterator returns a AlertPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetSMSCommandPageIterator(napSid string, data url.Values) *SMSCommandPageIterator {
	iter := NewPageIterator(s.client, data, superSimUsageRecordPathPart)
	return &SMSCommandPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *SMSCommandPageIterator) Next(ctx context.Context) (*SMSCommandPage, error) {
	ap := new(SMSCommandPage)
	err := s.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(ap.Meta.NextPageURL)
	return ap, nil
}

type SuperSimCommand struct {
	Command   string `json:"command"`
	AccountSid string        `json:"account_sid"`
	SimSid     string        `json:"sim_sid"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
	CommandMode   string `json:"command_mode"`
	DeliveryReceiptRequested bool `json:"delivery_receipt_requested"`
	Direction string `json:"direction"`
	Status string `json:"status"`
	Transport string `json:"transport"`
	Url string `json:"url"`
}

// NAPNetworkPage represents a page of Network Access Profiles.
type SuperSimCommandPage struct {
	Meta   Meta     `json:"meta"`
	SuperSimCommands []*SuperSimCommand `json:"commands"`
}

type SuperSimCommandPageIterator struct {
	p *PageIterator
}

// Get finds a single SuperSim NAP Network by its sid or unique name, or returns an error.
func (s *SuperSimService) CreateCommand(ctx context.Context, napSid string, data url.Values) (*SuperSimCommand, error) {
	superSimCommand := new(SuperSimCommand)
	err := s.client.CreateResource(ctx, superSimCommandsPathPart, data, superSimCommand)
	return superSimCommand, err
}

// Get finds a single SuperSim NAP Network by its sid or unique name, or returns an error.
func (s *SuperSimService) GetCommand(ctx context.Context, sid string) (*SuperSimCommand, error) {
	superSimCommand := new(SuperSimCommand)
	err := s.client.GetResource(ctx, superSimCommandsPathPart, sid, superSimCommand)
	return superSimCommand, err
}

// GetPage returns a single Page of resources, filtered by data.
//
// See https://www.twilio.com/docs/api/monitor/alerts#list-get-filters.
func (s *SuperSimService) GetCommandPage(ctx context.Context, napSid string, data url.Values) (*CommandPage, error) {
	return s.GetCommandPageIterator(superSimUsageRecordPathPart, data).Next(ctx)
}

// GetPageIterator returns a AlertPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperSimService) GetCommandPageIterator(napSid string, data url.Values) *SuperSimCommandPageIterator {
	iter := NewPageIterator(s.client, data, superSimUsageRecordPathPart)
	return &SuperSimCommandPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *SuperSimCommandPageIterator) Next(ctx context.Context) (*CommandPage, error) {
	ap := new(CommandPage)
	err := s.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(ap.Meta.NextPageURL)
	return ap, nil
}