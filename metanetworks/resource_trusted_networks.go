package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrustedNetworks() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"criteria": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_ip_config": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"apply_to_org": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"apply_to_entities": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"exempt_entities": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceVersionControlsCreate,
		Read:   resourceVersionControlsRead,
		Update: resourceVersionControlsUpdate,
		Delete: resourceVersionControlsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceTrustedNetworksCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	applyToOrg := d.Get("apply_to_org").(bool)
	applyToEntities := resourceTypeSetToStringSlice(d.Get("apply_to_entities").(*schema.Set))
	exemptEntities := resourceTypeSetToStringSlice(d.Get("exempt_entities").(*schema.Set))
	criteria := d.Get("criteria").([]Criteria)

	trustedNetworks := TrustedNetworks{
		Name:            name,
		Description:     description,
		Enabled:         enabled,
		ApplyToOrg:      applyToOrg,
		ApplyToEntities: applyToEntities,
		ExemptEntities:  exemptEntities,
		Criteria:        criteria,
	}

	var newTrustedNetworks *TrustedNetworks
	newTrustedNetworks, err := client.CreateTrustedNetworks(&trustedNetworks)
	if err != nil {
		return err
	}

	d.SetId(newTrustedNetworks.ID)

	err = trustedNetworksToResource(d, newTrustedNetworks)
	if err != nil {
		return err
	}

	return resourceTrustedNetworksRead(d, m)
}

func resourceTrustedNetworksRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	trustedNetworks, err := client.GetTrustedNetworks(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = trustedNetworksToResource(d, trustedNetworks)
	if err != nil {
		return err
	}

	return nil
}

func resourceTrustedNetworksUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	applyToOrg := d.Get("apply_to_org").(bool)
	applyToEntities := resourceTypeSetToStringSlice(d.Get("apply_to_entities").(*schema.Set))
	exemptEntities := resourceTypeSetToStringSlice(d.Get("exempt_entities").(*schema.Set))
	criteria := d.Get("criteria").([]Criteria)

	trustedNetworks := TrustedNetworks{
		Name:            name,
		Description:     description,
		Enabled:         enabled,
		ApplyToOrg:      applyToOrg,
		ApplyToEntities: applyToEntities,
		ExemptEntities:  exemptEntities,
		Criteria:        criteria,
	}

	var updatedTrustedNetworks *TrustedNetworks
	updatedTrustedNetworks, err := client.UpdateTrustedNetworks(d.Id(), &trustedNetworks)
	if err != nil {
		return err
	}

	d.SetId(updatedTrustedNetworks.ID)

	err = trustedNetworksToResource(d, updatedTrustedNetworks)
	if err != nil {
		return err
	}

	return resourceTrustedNetworksRead(d, m)
}

func resourceTrustedNetworksDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteTrustedNetworks(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func trustedNetworksToResource(d *schema.ResourceData, m *TrustedNetworks) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("enabled", m.Enabled)
	d.Set("apply_to_org", m.ApplyToOrg)
	d.Set("apply_to_entities", m.ApplyToEntities)
	d.Set("exempt_entities", m.ExemptEntities)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("criteria", m.CreatedAt)

	d.SetId(m.ID)

	return nil
}
