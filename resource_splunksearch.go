package main

import (
	"os"
	"net/http"
	"github.com/hashicorp/terraform/helper/schema"
	spk "github.com/shua/splunksearch"
)

func resourceSplunkSearch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSearchCreate,
		Read:   resourceSearchRead,
		Update: resourceSearchUpdate,
		Delete: resourceSearchDelete,

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
			"disabled": &schema.Schema{
				Type:	 schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func client() spk.Client {
	return spk.Client{
		Endpoint: os.Getenv("SPLUNK_ENDPOINT"),
		Username: os.Getenv("SPLUNK_USERNAME"),
		Password: os.Getenv("SPLUNK_PASSWORD"),
		ApiPath:  os.Getenv("SPLUNK_APIPATH"), // eg "/servicesNS/<user>/<index>"
		Client: &http.Client{},
	}
}

func search(d *schema.ResourceData) spk.Search {
	ss := spk.Search{
		"name": spk.SType{Str: d.Get("name").(string)},
		"search": spk.SType{Str: d.Get("search").(string)},
		"description": spk.SType{Str: d.Get("description").(string)},
	}

	if (d.Get("disabled").(bool)) {
		ss["disabled"] = spk.SType{Str: "1"}
	} else {
		ss["disabled"] = spk.SType{Str: "0"}
	}

	return ss
}

func resourceSearchCreate(d *schema.ResourceData, m interface{}) error {
	sc := client()
	ss := search(d)

	ns, err := sc.NewSearch(ss)
	d.SetId(ns["name"].Str)
	return err
}

func resourceSearchRead(d *schema.ResourceData, m interface{}) error {
	sc := client()
	ss, err := sc.GetSearch(d.Get("name").(string))
	if err != nil {
		serr, ok := err.(spk.SError)
		if !ok {
			return err
		}
		if serr.StatusCode == 404 {
			d.SetId("")
			return nil
		}
	}

	d.Set("search", ss["search"].Str)
	d.Set("description", ss["description"].Str)
	d.Set("disabled", ss["disabled"].Str == "1")

	return nil
}

func resourceSearchUpdate(d *schema.ResourceData, m interface{}) error {
	sc := client()
	// I could create a search with *only* the updated fields, but it's 1 call either way
	ss := search(d)
	_, err := sc.UpdateSearch(ss)
	return err
}

func resourceSearchDelete(d *schema.ResourceData, m interface{}) error {
	sc := client()
	_, err := sc.DeleteSearch(d.Get("name").(string))
	return err
}
