package twilio

import (
	"context"
	"net/url"
)

const smsCommandsPathPart = "SmsCommands"
const commandsPathPart = "Commands"

// SuperCommandService handles API requests for all SuperSim Commands
type SuperCommandService struct {
	client *Client
}

type SMSCommand struct {
	AccountSid  string     `json:"account_sid"`
	Payload     string     `json:"payload"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
	SimSid      string     `json:"sim_sid"`
	Status      string     `json:"status"`
	Sid         string     `json:"sid"`
	Direction   string     `json:"direction"`
	Url         string     `json:"url"`
}

// SMSCommandPage represents a page of SMSCommands.
type SMSCommandPage struct {
	Meta        Meta          `json:"meta"`
	SMSCommands []*SMSCommand `json:"usage_records"`
}

type SMSCommandPageIterator struct {
	p *PageIterator
}

// Create creates a new SMS command with the data provided, or returns an error.
func (s *SuperCommandService) CreateSMSCommand(ctx context.Context, data url.Values) (*SMSCommand, error) {
	smsCommand := new(SMSCommand)
	err := s.client.CreateResource(ctx, smsCommandsPathPart, data, smsCommand)
	return smsCommand, err
}

// Get finds a single SMSCommand by its sid, or returns an error.
func (s *SuperCommandService) GetSMSCommand(ctx context.Context, sid string) (*SMSCommand, error) {
	smsCommand := new(SMSCommand)
	err := s.client.GetResource(ctx, smsCommandsPathPart, sid, smsCommand)
	return smsCommand, err
}

// GetPage returns a single Page of SMSCommands, filtered by data.
//
// See https://www.twilio.com/docs/iot/supersim/api/smscommand-resource#read-multiple-smscommand-resources.
func (s *SuperCommandService) GetSMSCommandPage(ctx context.Context, data url.Values) (*SMSCommandPage, error) {
	return s.GetSMSCommandPageIterator(data).Next(ctx)
}

// GetPageIterator returns a SMSCommandPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperCommandService) GetSMSCommandPageIterator(data url.Values) *SMSCommandPageIterator {
	iter := NewPageIterator(s.client, data, smsCommandsPathPart)
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
	Command                  string     `json:"command"`
	AccountSid               string     `json:"account_sid"`
	SimSid                   string     `json:"sim_sid"`
	DateCreated              TwilioTime `json:"date_created"`
	DateUpdated              TwilioTime `json:"date_updated"`
	CommandMode              string     `json:"command_mode"`
	DeliveryReceiptRequested bool       `json:"delivery_receipt_requested"`
	Direction                string     `json:"direction"`
	Status                   string     `json:"status"`
	Transport                string     `json:"transport"`
	Url                      string     `json:"url"`
}

// NAPNetworkPage represents a page of Network Access Profiles.
type SuperSimCommandPage struct {
	Meta             Meta               `json:"meta"`
	SuperSimCommands []*SuperSimCommand `json:"commands"`
}

type SuperSimCommandPageIterator struct {
	p *PageIterator
}

// Create creates a new SuperSimCommand with the data provided, or returns an error.
func (s *SuperCommandService) CreateCommand(ctx context.Context, data url.Values) (*SuperSimCommand, error) {
	superSimCommand := new(SuperSimCommand)
	err := s.client.CreateResource(ctx, commandsPathPart, data, superSimCommand)
	return superSimCommand, err
}

// Get finds a single SuperSimCommand by its sid, or returns an error.
func (s *SuperCommandService) GetCommand(ctx context.Context, sid string) (*SuperSimCommand, error) {
	superSimCommand := new(SuperSimCommand)
	err := s.client.GetResource(ctx, commandsPathPart, sid, superSimCommand)
	return superSimCommand, err
}

// GetPage returns a single Page of SuperSimCommands, filtered by data.
//
// See https://www.twilio.com/docs/iot/supersim/api/command-resource#read-multiple-command-resources.
func (s *SuperCommandService) GetCommandPage(ctx context.Context, data url.Values) (*SuperSimCommandPage, error) {
	return s.GetCommandPageIterator(data).Next(ctx)
}

// GetPageIterator returns a SuperSimCommandPageIterator with the given page
// filters. Call iterator.Next() to get the first page of resources (and again
// to retrieve subsequent pages).
func (s *SuperCommandService) GetCommandPageIterator(data url.Values) *SuperSimCommandPageIterator {
	iter := NewPageIterator(s.client, data, commandsPathPart)
	return &SuperSimCommandPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (s *SuperSimCommandPageIterator) Next(ctx context.Context) (*SuperSimCommandPage, error) {
	ap := new(SuperSimCommandPage)
	err := s.p.Next(ctx, ap)
	if err != nil {
		return nil, err
	}
	s.p.SetNextPageURI(ap.Meta.NextPageURL)
	return ap, nil
}
