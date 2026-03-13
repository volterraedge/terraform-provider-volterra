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
	"gopkg.volterra.us/stdlib/codec"
	ves_io_schema_known_label "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/known_label"
	ves_io_schema_known_label_key "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/known_label_key"
)

const (
	setKnownLableRPCFQN = "ves.io.schema.known_label.CustomAPI"
	uriLabel            = "/public/namespaces/%s/known_label%s"
)

// resourceVolterraSetKnownLabel is implementation of Volterra's known_label resources
func resourceVolterraSetKnownLabel() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSetKnownLabelCreate,
		Read:   resourceVolterraSetKnownLabelRead,
		Update: resourceVolterraSetKnownLabelUpdate,
		Delete: resourceVolterraSetKnownLabelDelete,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceVolterraSetKnownLabelCreate creates known_label resource
func resourceVolterraSetKnownLabelCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var key, namespace, value, description string
	if v, ok := d.GetOk("key"); ok {
		key = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}
	if v, ok := d.GetOk("value"); ok {
		value = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		description = v.(string)
	}
	req := &ves_io_schema_known_label.CreateRequest{
		Key:         key,
		Namespace:   namespace,
		Description: description,
		Value:       value,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] setting volterra known label create request: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(uriLabel, namespace, "/create"), fmt.Sprintf("%s.%s", setKnownLableRPCFQN, "Create"), yamlReq)
	if err != nil {
		return fmt.Errorf("Error encountered while creating a label. Error: %s", err)
	}
	d.SetId(uuid.New().String())
	return resourceVolterraSetKnownLabelRead(d, meta)
}

// resourceVolterraSetKnownLabelRead read known_label resource
func resourceVolterraSetKnownLabelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var key, namespace, value string
	if v, ok := d.GetOk("key"); ok {
		key = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}
	if v, ok := d.GetOk("value"); ok {
		value = v.(string)
	}
	req := &ves_io_schema_known_label.GetRequest{
		Key:       key,
		Namespace: namespace,
		Query:     ves_io_schema_known_label.QUERY_EXACT_LABEL,
		Value:     value,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] setting volterra known label get request value: %+v", req)
	rsp, err := client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(uriLabel, namespace, "s"), fmt.Sprintf("%s.%s", setKnownLableRPCFQN, "Get"), yamlReq)
	if err != nil {
		fmt.Printf("Error encountered while fetching a label info. Error: %s", err)
	}
	if rsp != nil {
		rspVal := rsp.(*ves_io_schema_known_label.GetResponse).Label
		if len(rspVal) > 0 {
			d.Set("key", rspVal[0].Key)
			d.Set("value", rspVal[0].Value)
		}
	}
	return nil
}

// resourceVolterraSetKnownLabelUpdate updates known_label resource
func resourceVolterraSetKnownLabelUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	rawState := d.GetRawState()
	statefile_key := rawState.GetAttr("key").AsString()
	statefile_value := rawState.GetAttr("value").AsString()
	statefile_description := rawState.GetAttr("description").AsString()
	var key, namespace, value, description string
	if v, ok := d.GetOk("key"); ok {
		key = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}
	if v, ok := d.GetOk("value"); ok {
		value = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		description = v.(string)
	}
	if statefile_value != value || statefile_description != description || statefile_key != key {
		// Get Request for known_labe_key
		getReq := &ves_io_schema_known_label_key.GetRequest{
			Key:       key,
			Namespace: namespace,
			Query:     ves_io_schema_known_label_key.QUERY_EXACT_LABEL_KEY,
		}
		yamlReq, err := codec.ToYAML(getReq)
		if err != nil {
			return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] setting volterra known label key get request value: %+v", getReq)
		response, err := client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(uri, namespace, "s"), fmt.Sprintf("%s.%s", setKnownLableKeyRPCFQN, "Get"), yamlReq)
		if err != nil {
			fmt.Printf("Update: Error encountered while fetching a label info. Error: : %s", err)
		}
		rs := response.(*ves_io_schema_known_label_key.GetResponse).LabelKey
		if len(rs) == 0 {
			d.Set("description", statefile_description)
			d.Set("value", statefile_value)
			d.Set("key", statefile_key)
			return fmt.Errorf("key %s is not present on F5XC console.", key)
		}

		// Delete Request
		reqDel := &ves_io_schema_known_label.DeleteRequest{
			Key:       statefile_key,
			Namespace: namespace,
			Value:     statefile_value,
		}
		yamlReqDel, err := codec.ToYAML(reqDel)
		if err != nil {
			log.Printf("Ignored: Error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] setting volterra known label delete request value: %+v", reqDel)
		_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(uriLabel, namespace, "/delete"), fmt.Sprintf("%s.%s", setKnownLableRPCFQN, "Delete"), yamlReqDel)
		if err != nil {
			log.Printf("Update: Ignored: Error encountered while deleting a label. Error: : %s", err)
		}

		// Create Request
		reqCre := &ves_io_schema_known_label.CreateRequest{
			Key:         key,
			Namespace:   namespace,
			Description: description,
			Value:       value,
		}
		yamlReqCre, err := codec.ToYAML(reqCre)
		if err != nil {
			return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] setting volterra known label create request value: %+v", reqCre)
		resCre, err := client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(uriLabel, namespace, "/create"), fmt.Sprintf("%s.%s", setKnownLableRPCFQN, "Create"), yamlReqCre)
		if err != nil {
			return fmt.Errorf("Update: Error encountered while creating a label. Error: : %s", err)
		}
		rsp := resCre.(*ves_io_schema_known_label.CreateResponse).Label
		d.Set("description", rsp.Description)
		d.Set("value", rsp.Value)
		d.Set("key", rsp.Key)
	}
	return nil
}

// resourceVolterraSetKnownLabelDelete delete known_label resource
func resourceVolterraSetKnownLabelDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var key, namespace, value string
	if v, ok := d.GetOk("key"); ok {
		key = v.(string)
	}
	if v, ok := d.GetOk("namespace"); ok {
		namespace = v.(string)
	}
	if v, ok := d.GetOk("value"); ok {
		value = v.(string)
	}
	// Get Request
	getReq := &ves_io_schema_known_label.GetRequest{
		Key:       key,
		Namespace: namespace,
		Query:     ves_io_schema_known_label.QUERY_KEY_PREFIX_LABELS,
		Value:     value,
	}
	yamlReq, err := codec.ToYAML(getReq)
	if err != nil {
		return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] setting volterra known label get request value: %+v", getReq)
	getRsp, err := client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(uriLabel, namespace, "s"), fmt.Sprintf("%s.%s", setKnownLableRPCFQN, "Get"), yamlReq)
	if err != nil {
		return fmt.Errorf("Delete: Error encountered while fetching all label. Error: : %s", err)
	}
	response := getRsp.(*ves_io_schema_known_label.GetResponse).Label
	for _, label := range response {
		req := &ves_io_schema_known_label.DeleteRequest{
			Key:       label.Key,
			Namespace: namespace,
			Value:     label.Value,
		}
		yamlReq, err := codec.ToYAML(req)
		if err != nil {
			log.Printf("Ignored: Error marshalling rpc response to yaml: %s", err)
		}
		log.Printf("[DEBUG] setting volterra known label delete request value: %+v", req)
		_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(uriLabel, namespace, "/delete"), fmt.Sprintf("%s.%s", setKnownLableRPCFQN, "Delete"), yamlReq)
		if err != nil {
			log.Printf("Error encountered while deleting a label. Error: : %s", err)
		}
	}
	return nil
}
