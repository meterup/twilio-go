package twilio

import "context"

const superSimPathPart = "Sim"

type SuperSimService struct {
	client *Client
}

type SuperSim struct {
	Sid         string     `json:"sid"`
	Status      string     `json:"status"`
	DateCreated TwilioTime `json:"date_created"`
	DateUpdated TwilioTime `json:"date_updated"`
	AccountSid  string     `json:"account_sid"`
}

// Get finds a single Room resource by its sid or unique name, or returns an error.
func (r *SuperSimService) Get(ctx context.Context, sidOrUniqueName string) (*SuperSim, error) {
	superSim := new(SuperSim)
	err := r.client.GetResource(ctx, superSimPathPart, sidOrUniqueName, superSim)
	return superSim, err
}