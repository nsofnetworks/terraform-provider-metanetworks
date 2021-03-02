package metanetworks

import (
	"errors"
	"log"
)

const (
	easyLinksEndpoint string = "/v1/easylinks"
)

// EasyLinks ...
type EasyLinks struct {
	Name            string   `json:"name"`
	Description     string   `json:"description,omitempty"`
	Audit           bool     `json:"audit,omitempty"`
	Viewers         []string `json:"viewers,omitempty"`
	AccessFQDN      string   `json:"access_fqdn,omitempty"`
	AccessType      string   `json:"access_type" type:"bool"`
	CertificateID   string   `json:"certificate_id,omitempty"`
	DomainName      string   `json:"domain_name,omitempty" type:"bool"`
	EnableSNI       bool     `json:"enable_sni,omitempty"`
	Icon            string   `json:"icon,omitempty"`
	MappedElementID string   `json:"mapped_element_id,omitempty"`
	Protocol        string   `json:"protocol,omitempty"`
	Port            int      `json:"port,omitempty"`
	RoothPath       string   `json:"root_path,omitempty"`
	CreatedAt       string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID              string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt      string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID           string   `json:"org_id,omitempty" meta_api:"read_only"`
}

// GetEasyLinks ...
func (c *Client) GetEasyLinks(easyLinksID string) (*EasyLinks, error) {
	var easyLinks EasyLinks
	err := c.Read(easyLinksEndpoint+"/"+easyLinksID, &easyLinks)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Easylink Settings from Get: %s", easyLinks.ID)
	return &easyLinks, nil
}

// UpdateEasyLinks ...
func (c *Client) UpdateEasyLinks(easyLinksID string, easyLinks *EasyLinks) (*EasyLinks, error) {
	resp, err := c.Update(easyLinksEndpoint+"/"+easyLinksID, *easyLinks)
	if err != nil {
		return nil, err
	}
	updatedEasyLinks, _ := resp.(*EasyLinks)

	log.Printf("Returning Easylink Settings from Update: %s", updatedEasyLinks.ID)
	return updatedEasyLinks, nil
}

// CreateEasyLinks ...
func (c *Client) CreateEasyLinks(easyLinks *EasyLinks) (*EasyLinks, error) {
	resp, err := c.Create(easyLinksEndpoint, *easyLinks)
	if err != nil {
		return nil, err
	}

	createdEasyLinks, ok := resp.(*EasyLinks)
	if !ok {
		return nil, errors.New("Object returned from API was not an Easylink Pointer")
	}

	log.Printf("Returning Easylink Settings from Create: %s", createdEasyLinks.ID)
	return createdEasyLinks, nil
}

// DeleteEasyLinks ...
func (c *Client) DeleteEasyLinks(easyLinksID string) error {
	err := c.Delete(easyLinksEndpoint + "/" + easyLinksID)
	if err != nil {
		return err
	}

	return nil
}