package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSwgPacFiles() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"apply_to_org": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"sources": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"exempt_sources": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceSwgPacFilesCreate,
		Read:   resourceSwgPacFilesRead,
		Update: resourceSwgPacFilesUpdate,
		Delete: resourceSwgPacFilesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceSwgPacFilesCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	priority := d.Get("priority").(int)
	applyToOrg := d.Get("apply_to_org").(bool)
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))

	swgPacFiles := SwgPacFiles{
		Name:          name,
		Description:   description,
		Enabled:       enabled,
		Priority:      priority,
		ApplyToOrg:    applyToOrg,
		ExemptSources: exemptSources,
		Sources:       sources,
	}

	var newSwgPacFiles *SwgPacFiles
	newSwgPacFiles, err := client.CreateSwgPacFiles(&swgPacFiles)
	if err != nil {
		return err
	}

	d.SetId(newSwgPacFiles.ID)

	err = swgPacFilesToResource(d, newSwgPacFiles)
	if err != nil {
		return err
	}

	return resourceSwgPacFilesRead(d, m)
}

func resourceSwgPacFilesRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	var swgPacFiles *SwgPacFiles
	swgPacFiles, err := client.GetSwgPacFiles(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = swgPacFilesToResource(d, swgPacFiles)
	if err != nil {
		return err
	}

	return nil
}

func resourceSwgPacFilesUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)
	priority := d.Get("priority").(int)
	applyToOrg := d.Get("apply_to_org").(bool)
	exemptSources := resourceTypeSetToStringSlice(d.Get("exempt_sources").(*schema.Set))
	sources := resourceTypeSetToStringSlice(d.Get("sources").(*schema.Set))

	swgPacFiles := SwgPacFiles{
		Name:          name,
		Description:   description,
		Enabled:       enabled,
		Priority:      priority,
		ApplyToOrg:    applyToOrg,
		ExemptSources: exemptSources,
		Sources:       sources,
	}

	var updatedSwgPacFiles *SwgPacFiles
	updatedSwgPacFiles, err := client.UpdateSwgPacFiles(d.Id(), &swgPacFiles)
	if err != nil {
		return err
	}

	d.SetId(updatedSwgPacFiles.ID)

	err = swgPacFilesToResource(d, updatedSwgPacFiles)
	if err != nil {
		return err
	}

	return resourceSwgPacFilesRead(d, m)
}

func resourceSwgPacFilesDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteSwgPacFiles(d.Id())
	return err
}

func swgPacFilesToResource(d *schema.ResourceData, m *SwgPacFiles) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("enabled", m.Enabled)
	d.Set("priority", m.Priority)
	d.Set("apply_to_org", m.ExemptSources)
	d.Set("sources", m.Sources)
	d.Set("exempt_sources", m.ExemptSources)

	d.SetId(m.ID)

	return nil
}
