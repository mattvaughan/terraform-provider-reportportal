package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"reportportal"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	client := m.(*reportportal.Client)
	client.CreateProject(name)
	d.SetId(name)
	return resourceRead(d, m)
}

func resourceRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	client := m.(*reportportal.Client)
	project := client.GetProject(name)

	if project.ProjectId != "" {
		d.Set("name", name)
	}

	return nil
}

func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRead(d, m)
}

func resourceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
