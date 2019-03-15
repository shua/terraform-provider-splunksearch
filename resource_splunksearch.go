package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSplunkSearch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSplunkSearchCreate,
		Read:   resourceSplunkSearchRead,
		Update: resourceSplunkSearchUpdate,
		Delete: resourceSplunkSearchDelete,

		Schema: map[string]*schema.Schema{
			"search": &schema.Schema{
				Type:	 schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSplunkSearchCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSplunkSearchRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSplunkSearchUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSplunkSearchDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
