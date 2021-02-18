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
	setActiveNPsRPCFQN = "ves.io.schema.namespace.NamespaceCustomAPI.SetActiveNetworkPolicies"
	activeNPsURI       = "/public/namespaces/%s/active_network_policies"
	getActiveNPsRPCFQN = "ves.io.schema.namespace.NamespaceCustomAPI.GetActiveNetworkPolicies"
)

// resourceVolterraActiveNetworkPolicies is implementation of setting active network policies for a given namespace
func resourceVolterraActiveNetworkPolicies() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraActiveNetworkPoliciesCreate,
		Read:   resourceVolterraActiveNetworkPoliciesRead,
		Update: resourceVolterraActiveNetworkPoliciesUpdate,
		Delete: resourceVolterraActiveNetworkPoliciesDelete,

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

// resourceVolterraActiveNetworkPoliciesCreate sets active network policies for a namespace
func resourceVolterraActiveNetworkPoliciesCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req := &ves_io_schema_ns.SetActiveNetworkPoliciesRequest{}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		req.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("policies"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		policies := make([]*ves_io_schema_views.ObjectRefType, len(sl))
		req.NetworkPolicies = policies
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

	log.Printf("[DEBUG] Setting active network policies for a namespace, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(activeNPsURI, req.Namespace), setActiveNPsRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Setting active network policies for a namespace: %s", err)
	}

	d.SetId(uuid.New().String())

	return resourceVolterraActiveNetworkPoliciesRead(d, meta)
}

func resourceVolterraActiveNetworkPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	namespace := d.Get("namespace").(string)

	req := &ves_io_schema_ns.GetActiveNetworkPoliciesRequest{
		Namespace: namespace,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Setting active network policies for a namespace, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(activeNPsURI, req.Namespace), getActiveNPsRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Active network policies are not set in namespace: %s", namespace)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Active network policies in namespace %s: %s", namespace, err)
	}
	return nil
}

// resourceVolterraActiveNetworkPoliciesUpdate update active network policies for a namespace
func resourceVolterraActiveNetworkPoliciesUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceVolterraActiveNetworkPoliciesCreate(d, meta)
}

// resourceVolterraActiveNetworkPoliciesDelete remove active network policies from a namespace
func resourceVolterraActiveNetworkPoliciesDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	req := &ves_io_schema_ns.SetActiveNetworkPoliciesRequest{}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		req.Namespace =
			v.(string)
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Setting active network policies for a namespace, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(activeNPsURI, req.Namespace), setActiveNPsRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Active network policies are not set in namespace %s", req.Namespace)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Active network policies %q: %s", d.Id(), err)
	}
	return nil
}
