//
// Copyright (c) 2023 F5 Inc. All rights reserved.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ves_io_schema_known_label_key "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/known_label_key"
	"gopkg.volterra.us/stdlib/codec"
)

const (
	setKnownLableKeyRPCFQN = "ves.io.schema.known_label_key.CustomAPI"
	uri                    = "/public/namespaces/%s/known_label_key%s"
)

// resourceVolterraSetKnownLabelKey is implementation of Volterra's known_label_key resources
func resourceVolterraSetKnownLabelKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSetKnownLabelKeyCreate,
		Read:   resourceVolterraSetKnownLabelKeyRead,
		Update: resourceVolterraSetKnownLabelKeyUpdate,
		Delete: resourceVolterraSetKnownLabelKeyDelete,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceVolterraSetKnownLabelKeyCreate creates known_label_key resource
func resourceVolterraSetKnownLabelKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var key, namespace, description string
	if v, ok := d.GetOk("key"); ok {
		key = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		description = v.(string)
	}
	req := &ves_io_schema_known_label_key.CreateRequest{
		Key:         key,
		Namespace:   namespace,
		Description: description,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] Setting Volterra known label key create request value: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(uri, namespace, "/create"), fmt.Sprintf("%s.%s", setKnownLableKeyRPCFQN, "Create"), yamlReq)
	if err != nil {
		return fmt.Errorf("Error encountered while creating a key. Error: %s", err)
	}
	d.SetId(uuid.New().String())
	return resourceVolterraSetKnownLabelKeyRead(d, meta)
}

// resourceVolterraSetKnownLabelKeyRead read known_label_key resource
func resourceVolterraSetKnownLabelKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var key, namespace string
	if v, ok := d.GetOk("key"); ok {
		key = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}
	req := &ves_io_schema_known_label_key.GetRequest{
		Key:       key,
		Namespace: namespace,
		Query:     ves_io_schema_known_label_key.QUERY_EXACT_LABEL_KEY,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] setting volterra known label key get request value: %+v", req)
	rsp, err := client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(uri, namespace, "s"), fmt.Sprintf("%s.%s", setKnownLableKeyRPCFQN, "Get"), yamlReq)
	if err != nil {
		fmt.Printf("Error encountered while fetching a key info. Error: %s", err)
	}
	if rsp != nil {
		rspKey := rsp.(*ves_io_schema_known_label_key.GetResponse).LabelKey
		if len(rspKey) > 0 {
			d.Set("key", rspKey[0].Key)
		}
	}
	return nil
}

// resourceVolterraSetKnownLabelKeyUpdate updates known_label_key resource
func resourceVolterraSetKnownLabelKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	rawState := d.GetRawState()
	statefile_key := rawState.GetAttr("key").AsString()
	statefile_description := rawState.GetAttr("description").AsString()
	var key, description string
	if v, ok := d.GetOk("key"); ok {
		key = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		description = v.(string)
	}
	if statefile_key != key || statefile_description != description {
		d.Set("key", statefile_key)
		d.Set("description", statefile_description)
		return fmt.Errorf("Modifying the key name or description is not supported")
	}
	return nil
}

// resourceVolterraSetKnownLabelKeyDelete delete known_label_key resource
func resourceVolterraSetKnownLabelKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var key, namespace string
	if v, ok := d.GetOk("key"); ok {
		key = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}
	req := &ves_io_schema_known_label_key.DeleteRequest{
		Key:       key,
		Namespace: namespace,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		log.Printf("Ignored: Error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] setting volterra known label key delete request value: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(uri, namespace, "/delete"), fmt.Sprintf("%s.%s", setKnownLableKeyRPCFQN, "Delete"), yamlReq)
	if err != nil {
		return fmt.Errorf("Error encountered while deleting a key. Error: %s", err)
	}
	return nil
}
