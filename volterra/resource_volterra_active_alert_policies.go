// Copyright (c) 2023 F5 Inc. All rights reserved.

package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/codec"

	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

const (
	setActiveAPsRPCFQN = "ves.io.schema.namespace.NamespaceCustomAPI.SetActiveAlertPolicies"
	activeAPsURI       = "/public/namespaces/%s/active_alert_policies"
	getActiveAPsRPCFQN = "ves.io.schema.namespace.NamespaceCustomAPI.GetActiveAlertPolicies"
)

// resourceVolterraActiveAlertPolicies is implementation of setting active alert policies for a given namespace
func resourceVolterraActiveAlertPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraActiveAlertPoliciesCreate,
		Read:   resourceVolterraActiveAlertPoliciesRead,
		Update: resourceVolterraActiveAlertPoliciesUpdate,
		Delete: resourceVolterraActiveAlertPoliciesDelete,

		Schema: map[string]*schema.Schema{

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"policies": {

				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraActiveAlertPoliciesCreate sets active alert policies for a namespace
func resourceVolterraActiveAlertPoliciesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req := &ves_io_schema_ns.SetActiveAlertPoliciesRequest{}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		req.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("policies"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		policies := make([]*ves_io_schema_views.ObjectRefType, len(sl))
		req.AlertPolicies = policies
		for i, ps := range sl {

			policyMapData := ps.(map[string]interface{})
			policies[i] = &ves_io_schema_views.ObjectRefType{}

			if v, ok := policyMapData["name"]; ok && !isIntfNil(v) {
				policies[i].Name = v.(string)
			}

			if v, ok := policyMapData["namespace"]; ok && !isIntfNil(v) {
				policies[i].Namespace = v.(string)
			}

			if v, ok := policyMapData["tenant"]; ok && !isIntfNil(v) {
				policies[i].Tenant = v.(string)
			}

		}

	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Setting active alert policies for a namespace, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(activeAPsURI, req.Namespace), setActiveAPsRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting active alert policies for a namespace: %s", err)
	}

	d.SetId(uuid.New().String())

	return resourceVolterraActiveAlertPoliciesRead(d, meta)
}

func resourceVolterraActiveAlertPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	namespace := d.Get("namespace").(string)

	req := &ves_io_schema_ns.GetActiveAlertPoliciesRequest{
		Namespace: namespace,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Get active alert policies for a namespace %s", namespace)

	_, err = client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(activeAPsURI, req.Namespace), getActiveAPsRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Active Alert Policies are not set in namespace: %s", namespace)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Active Alert Policies in namespace %s: %s", namespace, err)
	}
	return nil
}

// resourceVolterraActiveAlertPoliciesUpdate update active alert policies for a namespace
func resourceVolterraActiveAlertPoliciesUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceVolterraActiveAlertPoliciesCreate(d, meta)
}

// resourceVolterraActiveAlertPoliciesDelete remove active alert policies from a namespace
func resourceVolterraActiveAlertPoliciesDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	req := &ves_io_schema_ns.SetActiveAlertPoliciesRequest{}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		req.Namespace =
			v.(string)
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Setting active alert policies for a namespace, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(activeAPsURI, req.Namespace), setActiveAPsRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Active alert Policies are not set in namespace %s", req.Namespace)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Active Alert Policies %q: %s", d.Id(), err)
	}
	return nil
}
