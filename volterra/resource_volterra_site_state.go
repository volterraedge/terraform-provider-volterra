//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/google/uuid"
	"gopkg.volterra.us/stdlib/codec"

	ves_io_schema_site "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/site"
)

const (
	siteCredRPCFQN = "ves.io.schema.site.CustomStateAPI"
	siteURI        = "/public/namespaces/system/site"
)

// resourceVolterraSiteState is implementation of Volterra's Site State
func resourceVolterraSiteState() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSiteStateCreate,
		Read:   resourceVolterraSiteStateRead,
		Update: resourceVolterraSiteStateUpdate,
		Delete: resourceVolterraSiteStateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"REREGISTRATION",
					"DECOMMISSIONING",
					"DECOMMISSIONNING",
				}, false),
			},
			"when": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"create",
					"delete",
				}, false),
			},
			"retry": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"wait_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}

// resourceVolterraSiteStateCreate creates Site resource
func resourceVolterraSiteStateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	when := d.Get("when").(string)
	state := d.Get("state").(string)
	var retry, waitTime int32
	if v, ok := d.GetOk("retry"); ok {
		retry = int32(v.(int))
	}
	if v, ok := d.GetOk("wait_time"); ok {
		waitTime = int32(v.(int))
	}

	var err error
	if when == "create" {
		for i := 0; i < int(retry); i++ {
			err = siteStateVal(client, name, state)
			if err != nil {
				log.Printf("[INFO] Retry setting site %s state to %s after %d sec", name, state, waitTime)
				time.Sleep(time.Duration(rand.Int31n(waitTime)) * time.Second)
				continue
			}
			break
		}
	}
	if err != nil {
		return fmt.Errorf("Error setting site %s state to %s: %s", name, state, err)
	}
	d.SetId(uuid.New().String())
	return nil
}

func siteStateVal(client *APIClient, name, state string) error {
	siteStateVal := ves_io_schema_site.SiteState_value[state]
	siteStateReq := &ves_io_schema_site.SetStateReq{
		Name:      name,
		Namespace: "system",
		State:     ves_io_schema_site.SiteState(siteStateVal),
	}
	yamlReq, err := codec.ToYAML(siteStateReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	_, err = client.CustomAPI(context.Background(), http.MethodPost,
		fmt.Sprintf("%s/%s/state", siteURI, name), fmt.Sprintf("%s.%s", siteCredRPCFQN, "SetState"), yamlReq)
	return err
}

func resourceVolterraSiteStateRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceVolterraSiteStateUpdate updates Site resource
func resourceVolterraSiteStateUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceVolterraSiteStateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	when := d.Get("when").(string)
	state := d.Get("state").(string)
	if state == "DECOMMISSIONNING" {
		state = ves_io_schema_site.DECOMMISSIONING.String()
	}
	var retry, waitTime int32
	if v, ok := d.GetOk("retry"); ok {
		retry = int32(v.(int))
	}
	if v, ok := d.GetOk("wait_time"); ok {
		waitTime = int32(v.(int))
	}

	var err error
	if when == "delete" {
		for i := 0; i < int(retry); i++ {
			err = siteStateVal(client, name, state)
			if err != nil {
				if strings.Contains(err.Error(), "status code 404") {
					log.Printf("[INFO] Site %s doesnt exists", d.Id())
					d.SetId("")
					return nil
				}
				log.Printf("[INFO] Retry setting site %s state to %s after %d sec", name, state, waitTime)
				time.Sleep(time.Duration(rand.Int31n(waitTime)) * time.Second)
				continue
			}
			break
		}
	}
	if err != nil {
		return fmt.Errorf("Error setting site %s state to %s: %s", name, state, err)
	}
	d.SetId("")
	return nil
}
