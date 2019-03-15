package main

import (
	"os"
	"net/http"
	"github.com/hashicorp/terraform/helper/schema"
	spk "github.com/shua/splunksearch"
)

func resourceSplunkSearch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSplunkSearchCreate,
		Read:   resourceSplunkSearchRead,
		Update: resourceSplunkSearchUpdate,
		Delete: resourceSplunkSearchDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:	 schema.TypeString,
				Required: true,
			},
			"search": &schema.Schema{
				Type:	 schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:	 schema.TypeString,
				Optional: true,
			},
		},
	}
}

func client() spk.SplunkClient {
	return spk.SplunkClient{
		Endpoint: os.Getenv("SPLUNK_ENDPOINT"),
		Username: os.Getenv("SPLUNK_USERNAME"),
		Password: os.Getenv("SPLUNK_PASSWORD"),
		ApiPath:  os.Getenv("SPLUNK_APIPATH"), // eg "/servicesNS/<user>/<index>"
		Client: &http.Client{},
	}
}

func resourceSplunkSearchCreate(d *schema.ResourceData, m interface{}) error {
	sc := client()
	ss := spk.SplunkSearch{
		"name": spk.SType{Str: d.Get("name").(string)},
		"search": spk.SType{Str: d.Get("search").(string)},
		"description": spk.SType{Str: d.Get("description").(string)},
	}

	_, err := sc.NewSearch(ss)
	return err
}

func resourceSplunkSearchRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSplunkSearchUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSplunkSearchDelete(d *schema.ResourceData, m interface{}) error {
	sc := client()
	_, err := sc.DeleteSearch(d.Get("name").(string))
	return err
}
