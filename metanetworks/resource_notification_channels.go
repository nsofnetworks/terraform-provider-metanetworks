package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNotificationChannels() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
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
			"email_config": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"recipients": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"apply_to_org": &schema.Schema{
				Type:     schema.TypeBool,
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
		Create: resourceNotificationChannelsCreate,
		Read:   resourceNotificationChannelsRead,
		Update: resourceNotificationChannelsUpdate,
		Delete: resourceNotificationChannelsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceNotificationChannelsCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)

	notificationChannels := NotificationChannels{
		Name:        name,
		Description: description,
		Enabled:     enabled,
	}

	var newNotificationChannels *NotificationChannels
	newNotificationChannels, err := client.CreateNotificationChannels(&notificationChannels)
	if err != nil {
		return err
	}

	d.SetId(newNotificationChannels.ID)

	err = notificationChannelsToResource(d, newNotificationChannels)
	if err != nil {
		return err
	}

	return resourceNotificationChannelsRead(d, m)
}

func resourceNotificationChannelsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	notificationChannels, err := client.GetNotificationChannels(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = notificationChannelsToResource(d, notificationChannels)
	if err != nil {
		return err
	}

	return nil
}

func resourceNotificationChannelsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	enabled := d.Get("enabled").(bool)

	notificationChannels := NotificationChannels{
		Name:        name,
		Description: description,
		Enabled:     enabled,
	}

	var updatedNotificationChannels *NotificationChannels
	updatedNotificationChannels, err := client.UpdateNotificationChannels(d.Id(), &notificationChannels)
	if err != nil {
		return err
	}

	d.SetId(updatedNotificationChannels.ID)

	err = notificationChannelsToResource(d, updatedNotificationChannels)
	if err != nil {
		return err
	}

	return resourceNotificationChannelsRead(d, m)
}

func resourceNotificationChannelsDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteNotificationChannels(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func notificationChannelsToResource(d *schema.ResourceData, m *NotificationChannels) error {
	d.Set("description", m.Description)
	d.Set("name", m.Name)
	d.Set("enabled", m.Enabled)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)

	d.SetId(m.ID)

	return nil
}
