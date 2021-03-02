package metanetworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEasyLinks() *schema.Resource {
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
			"audit": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"viewers": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mapped_element_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"access_fqdn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"access_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_sni": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"icon": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"root_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enterprise_access": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},
						"hosts": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"http_host_header": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"navigator_compatibility": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"rewrite_content_types": {
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
						"rewrite_hosts": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},
						"rewrite_hosts_client": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},
						"rewrite_hosts_service_worker": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},
						"rewrite_http": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},
						"shared_cookies": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},
						"strip_origin_header": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},
						"strip_user_agent_header": {
							Type:     schema.TypeBool,
							Default:  false,
							Optional: true,
						},
					},
				},
			},
			"rdp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remote_app": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"server_keyboard_layout": {
							Type:     schema.TypeString,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
						},
					},
				},
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
		Create: resourceEasyLinksCreate,
		Read:   resourceEasyLinksRead,
		Update: resourceEasyLinksUpdate,
		Delete: resourceEasyLinksDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceEasyLinksCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	audit := d.Get("audit").(bool)
	viewers := resourceTypeSetToStringSlice(d.Get("viewers").(*schema.Set))
	protocol := d.Get("protocol").(string)
	port := d.Get("port").(int)
	domainname := d.Get("domain_name").(string)
	mappedElementID := d.Get("mapped_element_id").(string)
	accessfqdn := d.Get("access_fqdn").(string)
	accesstype := d.Get("access_type").(string)
	certificateid := d.Get("certificate_id").(string)
	enablesni := d.Get("enable_sni").(bool)
	icon := d.Get("icon").(string)
	rootpath := d.Get("root_path").(string)
	proxy := d.Get("proxy").([]Proxy)
	rdp := d.Get("rdp").([]RDP)

	easyLinks := EasyLinks{
		Name:            name,
		Description:     description,
		Audit:           audit,
		Protocol:        protocol,
		Port:            port,
		Viewers:         viewers,
		DomainName:      domainname,
		MappedElementID: mappedElementID,
		AccessFQDN:      accessfqdn,
		AccessType:      accesstype,
		CertificateID:   certificateid,
		EnableSNI:       enablesni,
		Icon:            icon,
		RoothPath:       rootpath,
		Proxy:           proxy,
		RDP:             rdp,
	}

	var newEasyLinks *EasyLinks
	newEasyLinks, err := client.CreateEasyLinks(&easyLinks)
	if err != nil {
		return err
	}

	d.SetId(newEasyLinks.ID)
	err = easyLinksToResource(d, newEasyLinks)
	if err != nil {
		return err
	}
	return resourceEasyLinksRead(d, m)
}

func resourceEasyLinksRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	easyLinks, err := client.GetEasyLinks(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	err = easyLinksToResource(d, easyLinks)
	if err != nil {
		return err
	}

	return nil
}

func resourceEasyLinksUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	audit := d.Get("audit").(bool)
	viewers := resourceTypeSetToStringSlice(d.Get("viewers").(*schema.Set))
	protocol := d.Get("protocol").(string)
	port := d.Get("port").(int)
	domainname := d.Get("domain_name").(string)
	mappedElementID := d.Get("mapped_element_id").(string)
	accessfqdn := d.Get("access_fqdn").(string)
	accesstype := d.Get("access_type").(string)
	certificateid := d.Get("certificate_id").(string)
	enablesni := d.Get("enable_sni").(bool)
	icon := d.Get("icon").(string)
	rootpath := d.Get("root_path").(string)
	proxy := d.Get("proxy").([]Proxy)
	rdp := d.Get("rdp").([]RDP)

	easyLinks := EasyLinks{
		Name:            name,
		Description:     description,
		Audit:           audit,
		Protocol:        protocol,
		Port:            port,
		Viewers:         viewers,
		DomainName:      domainname,
		MappedElementID: mappedElementID,
		AccessFQDN:      accessfqdn,
		AccessType:      accesstype,
		CertificateID:   certificateid,
		EnableSNI:       enablesni,
		Icon:            icon,
		RoothPath:       rootpath,
		Proxy:           proxy,
		RDP:             rdp,
	}

	var updatedEasyLinks *EasyLinks
	updatedEasyLinks, err := client.UpdateEasyLinks(d.Id(), &easyLinks)
	if err != nil {
		return err
	}
	err = easyLinksToResource(d, updatedEasyLinks)
	if err != nil {
		return err
	}

	return resourceEasyLinksRead(d, m)
}

func resourceEasyLinksDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*Client)

	err := client.DeleteEasyLinks(d.Id())
	if err != nil {
		return err
	}

	return nil
}

func easyLinksToResource(d *schema.ResourceData, m *EasyLinks) error {
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("audit", m.Audit)
	d.Set("viewers", m.Viewers)
	d.Set("domain_name", m.DomainName)
	d.Set("mapped_element_id", m.MappedElementID)
	d.Set("protocol", m.Protocol)
	d.Set("port", m.Port)
	d.Set("access_fqdn", m.AccessFQDN)
	d.Set("access_type", m.AccessType)
	d.Set("enable_sni", m.EnableSNI)
	d.Set("icon", m.Icon)
	d.Set("root_path", m.RoothPath)
	d.Set("proxy", m.Proxy)
	d.Set("rdp", m.RDP)
	d.Set("created_at", m.CreatedAt)
	d.Set("modified_at", m.ModifiedAt)
	d.Set("org_id", m.OrgID)

	d.SetId(m.ID)

	return nil
}
