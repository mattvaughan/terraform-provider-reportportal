package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"reportportal"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:			schema.TypeString,
				Required:	true,
			},
			"api_url": {
				Type:			schema.TypeString,
				Required:	true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"reportportal_project": resourceProject(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	accessKey := d.Get("access_key").(string)
	apiUrl := d.Get("api_url").(string)
	return &reportportal.Client{
		ApiUrl: apiUrl,
		ApiKey: accessKey,
	}, nil
}
