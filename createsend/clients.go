package createsend

import (
	"fmt"
	"strings"
)

// A Client represents a client of a Campaign Monitor account.
//
// See http://www.campaignmonitor.com/api/account/#getting_your_clients for more
// information.
type Client struct {
	ClientID string
	Name     string
}

// ListClients lists the clients associated with the authenticated account.
//
// See http://www.campaignmonitor.com/api/account/#getting_your_clients for more
// information.
func (c *APIClient) ListClients() ([]Client, error) {
	u := "clients.json"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	clients := new([]Client)
	err = c.Do(req, clients)
	if err != nil {
		// EOF is not a real error according to the Internet
		// See: https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
		if strings.Compare("EOF", err.Error()) == 0 {
			return *clients, nil
		} else {
			return nil, err
		}
	}
	return nil, err

	return *clients, err
}

// ListLists returns all of the subscriber lists that belong to a client.
//
// See http://www.campaignmonitor.com/api/clients/#subscriber_lists for more
// information.
func (c *APIClient) ListLists(clientID string) ([]*List, error) {
	u := fmt.Sprintf("clients/%s/lists.json", clientID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var lists []*List
	err = c.Do(req, &lists)
	if err != nil {
		// EOF is not a real error according to the Internet
		// See: https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
		if strings.Compare("EOF", err.Error()) == 0 {
			return lists, nil
		} else {
			return lists, err
		}
	}
	return lists, err
}

// ListForEmail represents a subscriber list *and* a specific email address's
// subscription to that list. The schema differs from that of List.
//
// See http://www.campaignmonitor.com/api/clients/#lists_for_email for more
// information.
type ListForEmail struct {
	ListID                 string
	ListName               string
	SubscriberState        string
	DateSubscriberAddedStr string `json:"DateSubscriberAdded"`
}

func (e *ListForEmail) IsSubscribed() bool {
	return e.SubscriberState == "Active"
}

func (e *ListForEmail) IsUnsubscribed() bool {
	return e.SubscriberState == "Unsubscribed"
}

// ListsForEmail returns all of the client's subscriber lists to which the email
// address is subscribed.
//
// See http://www.campaignmonitor.com/api/clients/#lists_for_email for more
// information.
func (c *APIClient) ListsForEmail(clientID string, email string) ([]*ListForEmail, error) {
	u := fmt.Sprintf("clients/%s/listsforemail.json?email=%s", clientID, email)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var lists []*ListForEmail
	err = c.Do(req, &lists)
	if err != nil {
		// EOF is not a real error according to the Internet
		// See: https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
		if strings.Compare("EOF", err.Error()) == 0 {
			return lists, nil
		} else {
			return lists, err
		}
	}

	return lists, err
}

type Campaign struct {
	FromName          string `json:"FromName"`
	FromEmail         string `json:"FromEmail"`
	ReplyTo           string `json:"ReplyTo"`
	WebVersionURL     string `json:"WebVersionURL"`
	WebVersionTextURL string `json:"WebVersionTextURL"`
	CampaignID        string `json:"CampaignID"`
	Subject           string `json:"Subject"`
	Name              string `json:"Name"`
	SentDate          string `json:"SentDate"`
	TotalRecipients   int64  `json:"TotalRecipients"`
}

// Campaigns return all the sent campaigns for a specific client
//
// See https://www.campaignmonitor.com/api/clients/#sent_campaigns for more
// information.
func (c *APIClient) Campaigns(clientID string) ([]*Campaign, error) {
	u := fmt.Sprintf("clients/%s/campaigns.json", clientID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var campaigns []*Campaign
	err = c.Do(req, &campaigns)
	if err != nil {
		// EOF is not a real error according to the Internet
		// See: https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
		if strings.Compare("EOF", err.Error()) == 0 {
			return campaigns, nil
		} else {
			return campaigns, err
		}
	}
	return campaigns, err
}

type Template struct {
	TemplateID    string `json:"TemplateID"`
	Name          string `json:"Name"`
	PreviewURL    string `json:"PreviewURL"`
	ScreenshotURL string `json:"ScreenshotURL"`
}

// Campaigns return all the templates for a specific client
func (c *APIClient) ListTemplates(clientID string) ([]*Template, error) {
	u := fmt.Sprintf("clients/%s/templates.json", clientID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var templates []*Template
	err = c.Do(req, &templates)
	if err != nil {
		// EOF is not a real error according to the Internet
		// See: https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
		if strings.Compare("EOF", err.Error()) == 0 {
			return templates, nil
		} else {
			return templates, err
		}
	}

	return templates, err
}
