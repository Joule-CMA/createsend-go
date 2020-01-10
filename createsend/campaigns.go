package createsend

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// CampaignRecipientsOptions represents the URL parameters that may be used to
// fetch campaign recipients.
//
// See https://www.campaignmonitor.com/api/campaigns/#campaign_recipients for
// more information.
type CampaignRecipientsOptions struct {
	Page           int
	PageSize       int
	OrderField     string
	OrderDirection string
}

// CampaignRecipients lists all the recipients from a campaign.
// See https://www.campaignmonitor.com/api/campaigns/#campaign_recipients for
// more information.
type CampaignRecipients struct {
	Results              []*Recipient `json:"Results"`
	ResultsOrderedBy     string       `json:"ResultsOrderedBy"`
	OrderDirection       string       `json:"OrderDirection"`
	PageNumber           int          `json:"PageNumber"`
	PageSize             int          `json:"PageSize"`
	RecordsOnThisPage    int          `json:"RecordsOnThisPage"`
	TotalNumberOfRecords int          `json:"TotalNumberOfRecords"`
	NumberOfPages        int          `json:"NumberOfPages"`
}

type Recipient struct {
	EmailAddress string `json:"EmailAddress"`
	ListID       string `json:"ListID"`
}

func (c *APIClient) CampaignRecipients(campaignID string, opt *CampaignRecipientsOptions) (*CampaignRecipients, error) {

	u := fmt.Sprintf("campaigns/%s/recipients.json", campaignID)

	if opt != nil {
		v := url.Values{}
		if opt.Page > 0 {
			v.Set("page", strconv.Itoa(opt.Page))
		}
		if opt.PageSize > 0 {
			v.Set("pagesize", strconv.Itoa(opt.PageSize))
		}
		if opt.OrderField != "" {
			v.Set("orderfield", opt.OrderField)
		}
		if opt.OrderDirection != "" {
			v.Set("orderdirection", opt.OrderDirection)
		}

		q := v.Encode()
		if q != "" {
			u = fmt.Sprintf("%s?%s", u, q)
		}
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var results CampaignRecipients
	err = c.Do(req, &results)
	return &results, err
}

// Campaign struct to create a campaign
// See https://www.campaignmonitor.com/api/campaigns/ for
// more information.
type CreateCampaign struct {
	Name            string          `json:"Name"`
	Subject         string          `json:"Subject"`
	FromName        string          `json:"FromName"`
	FromEmail       string          `json:"FromEmail"`
	ReplyTo         string          `json:"ReplyTo"`
	ListIDs         []string        `json:"ListIDs"`
	SegmentIDs      []string        `json:"SegmentIDs"`
	TemplateID      string          `json:"TemplateID"`
	TemplateContent TemplateContent `json:"TemplateContent"`
}
type Singleline struct {
	Label   string `json:"Label,omitempty"`
	Content string `json:"Content"`
	Href    string `json:"Href,omitempty"`
}
type Multiline struct {
	Content string `json:"Content"`
}
type Image struct {
	Content string `json:"Content"`
	Alt     string `json:"Alt,omitempty"`
	Href    string `json:"Href,omitempty"`
}
type Item struct {
	Layout      string       `json:"Layout"`
	Singlelines []Singleline `json:"Singlelines"`
	Multilines  []Multiline  `json:"Multilines"`
	Images      []Image      `json:"Images"`
}
type Repeater struct {
	Items []Item `json:"Items"`
}
type TemplateContent struct {
	Singlelines []Singleline `json:"Singlelines,omitempty"`
	Multilines  []Multiline  `json:"Multilines,omitempty"`
	Images      []Image      `json:"Images,omitempty"`
	Repeaters   []Repeater   `json:"Repeaters,omitempty"`
}

func (c *APIClient) CreateCampaign(clientID string, campaign CreateCampaign) (string, error) {

	u := fmt.Sprintf("campaigns/%s.json", clientID)

	req, err := c.NewRequest("POST", u, campaign)
	if err != nil {
		return "", err
	}

	var results string
	err = c.Do(req, &results)
	return results, err
}

func (c *APIClient) CreateCampaignFromTemplate(clientID string, campaign CreateCampaign) (string, error) {

	u := fmt.Sprintf("campaigns/%s/fromTemplate.json", clientID)

	req, err := c.NewRequest("POST", u, campaign)
	if err != nil {
		return "", err
	}

	var results string
	err = c.Do(req, &results)
	return results, err
}

type ScheduleCampaign struct {
	ConfirmationEmail string `json:"ConfirmationEmail"`
	SendDate          string `json:"SendDate"`
}

func (c *APIClient) ScheduleCampaign(campaignID string, confirmationEmail string, sendDate time.Time) (bool, error) {
	u := fmt.Sprintf("campaigns/%s/send.json", campaignID)

	sendDateStr := sendDate.Format("2006-01-02 15:04")

	scheduleCampaign := ScheduleCampaign{ConfirmationEmail: confirmationEmail, SendDate: sendDateStr}
	req, err := c.NewRequest("POST", u, scheduleCampaign)
	if err != nil {
		return false, err
	}

	var results string
	err = c.Do(req, &results)
	if err != nil {
		// EOF is not a real error according to the Internet
		// See: https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
		if strings.Compare("EOF", err.Error()) == 0 {
			return true, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
