package metanetworks

import (
	"errors"
	"log"
)

const (
	trustedNetworksEndpoint string = "/v1/trusted_networks"
)

// TrustedNetworks ...
type TrustedNetworks struct {
	Name            string     `json:"name"`
	Description     string     `json:"description,omitempty"`
	Enabled         bool       `json:"enabled" type:"bool"`
	ApplyToOrg      bool       `json:"apply_to_org"`
	ExemptEntities  []string   `json:"exempt_entities,omitempty"`
	ApplyToEntities []string   `json:"apply_to_entities,omitempty"`
	Criteria        []Criteria `json:"criteria,omitempty"`
	CreatedAt       string     `json:"created_at,omitempty" meta_api:"read_only"`
	ID              string     `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt      string     `json:"modified_at,omitempty" meta_api:"read_only"`
}

// Criteria ...
type Criteria struct {
	ExternalIPConfig      []ExternalIPConfig
	ResolvedAddressConfig []ResolvedAddressConfig
}

// ExternalIPConfig ...
type ExternalIPConfig struct {
	AddressesRanges []string `json:"addresses_ranges,"`
}

// ResolvedAddressConfig ...
type ResolvedAddressConfig struct {
	AddressesRanges []string `json:"addresses_ranges,"`
	Hostname        string   `json:"hostname,"`
}

// GetTrustedNetworks ...
func (c *Client) GetTrustedNetworks(trustedNetworksID string) (*TrustedNetworks, error) {
	var trustedNetworks TrustedNetworks
	err := c.Read(trustedNetworksEndpoint+"/"+trustedNetworksID, &trustedNetworks)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning Trusted Networks Settings from Get: %s", trustedNetworks.ID)
	return &trustedNetworks, nil
}

// UpdateTrustedNetworks ...
func (c *Client) UpdateTrustedNetworks(trustedNetworksID string, trustedNetworks *TrustedNetworks) (*TrustedNetworks, error) {
	resp, err := c.Update(trustedNetworksEndpoint+"/"+trustedNetworksID, *trustedNetworks)
	if err != nil {
		return nil, err
	}
	updatedTrustedNetworks, _ := resp.(*TrustedNetworks)

	log.Printf("Returning TrustedNetworks  Settings from Update: %s", updatedTrustedNetworks.ID)
	return updatedTrustedNetworks, nil
}

// CreateTrustedNetworks ...
func (c *Client) CreateTrustedNetworks(trustedNetworks *TrustedNetworks) (*TrustedNetworks, error) {
	resp, err := c.Create(trustedNetworksEndpoint, *trustedNetworks)
	if err != nil {
		return nil, err
	}

	createdTrustedNetworks, ok := resp.(*TrustedNetworks)
	if !ok {
		return nil, errors.New("Object returned from API was not a TrustedNetworks Pointer")
	}

	log.Printf("Returning TrustedNetworks Settings from Create: %s", createdTrustedNetworks.ID)
	return createdTrustedNetworks, nil
}

// DeleteTrustedNetworks ...
func (c *Client) DeleteTrustedNetworks(trustedNetworksID string) error {
	err := c.Delete(trustedNetworksEndpoint + "/" + trustedNetworksID)
	if err != nil {
		return err
	}

	return nil
}
