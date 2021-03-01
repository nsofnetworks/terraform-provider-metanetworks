package metanetworks

import (
	"errors"
	"log"
)

const (
	swgPacFilesEndpoint string = "/v1/pac_files"
)

// SwgPacFiles ...
type SwgPacFiles struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Enabled       bool     `json:"enabled,omitempty" type:"bool"`
	ApplyToOrg    bool     `json:"apply_to_org,omitempty"`
	Sources       []string `json:"sources,omitempty"`
	ExemptSources []string `json:"exempt_sources,omitempty"`
	Priority      int      `json:"priority"`
	CreatedAt     string   `json:"created_at,omitempty" meta_api:"read_only"`
	ID            string   `json:"id,omitempty" meta_api:"read_only"`
	ModifiedAt    string   `json:"modified_at,omitempty" meta_api:"read_only"`
	OrgID         string   `json:"org_id,omitempty" meta_api:"read_only"`
}

// GetSwgPacFiles ...
func (c *Client) GetSwgPacFiles(swgPacFilesID string) (*SwgPacFiles, error) {
	var swgPacFiles SwgPacFiles
	err := c.Read(swgPacFilesEndpoint+"/"+swgPacFilesID, &swgPacFiles)
	if err != nil {
		return nil, err
	}

	log.Printf("Returning SwgPacFiles from Get: %s", swgPacFiles.ID)
	return &swgPacFiles, nil
}

// UpdateSwgPacFiles ...
func (c *Client) UpdateSwgPacFiles(swgPacFilesID string, swgPacFiles *SwgPacFiles) (*SwgPacFiles, error) {
	resp, err := c.Update(swgPacFilesEndpoint+"/"+swgPacFilesID, *swgPacFiles)
	if err != nil {
		return nil, err
	}
	updatedSwgPacFiles, _ := resp.(*SwgPacFiles)

	log.Printf("Returning SwgPacFiles from Update: %s", updatedSwgPacFiles.ID)
	return updatedSwgPacFiles, nil
}

// CreateSwgPacFiles ...
func (c *Client) CreateSwgPacFiles(swgPacFiles *SwgPacFiles) (*SwgPacFiles, error) {
	resp, err := c.Create(swgPacFilesEndpoint, *swgPacFiles)
	if err != nil {
		return nil, err
	}

	createdSwgPacFiles, ok := resp.(*SwgPacFiles)
	if !ok {
		return nil, errors.New("Object returned from API was not a SwgPacFiles Pointer")
	}

	log.Printf("Returning SwgPacFiles from Create: %s", createdSwgPacFiles.ID)
	return createdSwgPacFiles, nil
}

//DeleteSwgPacFiles ...
func (c *Client) DeleteSwgPacFiles(swgPacFilesID string) error {
	err := c.Delete(swgPacFilesEndpoint + "/" + swgPacFilesID)
	if err != nil {
		return err
	}

	return nil
}
