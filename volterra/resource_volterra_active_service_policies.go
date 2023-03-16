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
	setActiveSPsRPCFQN = "ves.io.schema.namespace.NamespaceCustomAPI.SetActiveServicePolicies"
	activeSPsURI       = "/public/namespaces/%s/active_service_policies"
	getActiveSPsRPCFQN = "ves.io.schema.namespace.NamespaceCustomAPI.GetActiveServicePolicies"
)

// resourceVolterraActiveServicePolicies is implementation of setting active service policies for a given namespace
func resourceVolterraActiveServicePolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraActiveServicePoliciesCreate,
		Read:   resourceVolterraActiveServicePoliciesRead,
		Update: resourceVolterraActiveServicePoliciesUpdate,
		Delete: resourceVolterraActiveServicePoliciesDelete,

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

// resourceVolterraActiveServicePoliciesCreate sets active service policies for a namespace
func resourceVolterraActiveServicePoliciesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req := &ves_io_schema_ns.SetActiveServicePoliciesRequest{}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		req.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("policies"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		policies := make([]*ves_io_schema_views.ObjectRefType, len(sl))
		req.ServicePolicies = policies
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

	log.Printf("[DEBUG] Setting active service policies for a namespace, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(activeSPsURI, req.Namespace), setActiveSPsRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting active service policies for a namespace: %s", err)
	}

	d.SetId(uuid.New().String())

	return resourceVolterraActiveServicePoliciesRead(d, meta)
}

func resourceVolterraActiveServicePoliciesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	namespace := d.Get("namespace").(string)

	req := &ves_io_schema_ns.GetActiveServicePoliciesRequest{
		Namespace: namespace,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Get active service policies for a namespace %s", namespace)

	_, err = client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(activeSPsURI, req.Namespace), getActiveSPsRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Active Service Policies are not set in namespace: %s", namespace)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Active Service Policies in namespace %s: %s", namespace, err)
	}
	return nil
}

// resourceVolterraActiveServicePoliciesUpdate update active service policies for a namespace
func resourceVolterraActiveServicePoliciesUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceVolterraActiveServicePoliciesCreate(d, meta)
}

// resourceVolterraActiveServicePoliciesDelete remove active service policies from a namespace
func resourceVolterraActiveServicePoliciesDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	req := &ves_io_schema_ns.SetActiveServicePoliciesRequest{}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		req.Namespace =
			v.(string)
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Setting active service policies for a namespace, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(activeSPsURI, req.Namespace), setActiveSPsRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Active Service Policies are not set in namespace %s", req.Namespace)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Active Service Policies %q: %s", d.Id(), err)
	}
	return nil
}
