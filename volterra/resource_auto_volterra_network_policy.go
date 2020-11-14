//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_network_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_policy"
)

// resourceVolterraNetworkPolicy is implementation of Volterra's NetworkPolicy resources
func resourceVolterraNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraNetworkPolicyCreate,
		Read:   resourceVolterraNetworkPolicyRead,
		Update: resourceVolterraNetworkPolicyUpdate,
		Delete: resourceVolterraNetworkPolicyDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			"egress_rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
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

			"ingress_rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
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

			"prefix": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"prefix": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"prefix_selector": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"expressions": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraNetworkPolicyCreate creates NetworkPolicy resource
func resourceVolterraNetworkPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_network_policy.CreateSpecType{}
	createReq := &ves_io_schema_network_policy.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("egress_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		egressRulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.EgressRules = egressRulesInt
		for i, ps := range sl {

			erMapToStrVal := ps.(map[string]interface{})
			egressRulesInt[i] = &ves_io_schema.ObjectRefType{}

			egressRulesInt[i].Kind = "network_policy_rule"

			if v, ok := erMapToStrVal["name"]; ok && !isIntfNil(v) {
				egressRulesInt[i].Name = v.(string)
			}

			if v, ok := erMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				egressRulesInt[i].Namespace = v.(string)
			}

			if v, ok := erMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				egressRulesInt[i].Tenant = v.(string)
			}

			if v, ok := erMapToStrVal["uid"]; ok && !isIntfNil(v) {
				egressRulesInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("ingress_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		ingressRulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.IngressRules = ingressRulesInt
		for i, ps := range sl {

			irMapToStrVal := ps.(map[string]interface{})
			ingressRulesInt[i] = &ves_io_schema.ObjectRefType{}

			ingressRulesInt[i].Kind = "network_policy_rule"

			if v, ok := irMapToStrVal["name"]; ok && !isIntfNil(v) {
				ingressRulesInt[i].Name = v.(string)
			}

			if v, ok := irMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				ingressRulesInt[i].Namespace = v.(string)
			}

			if v, ok := irMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				ingressRulesInt[i].Tenant = v.(string)
			}

			if v, ok := irMapToStrVal["uid"]; ok && !isIntfNil(v) {
				ingressRulesInt[i].Uid = v.(string)
			}

		}

	}

	localEndpointTypeFound := false

	if v, ok := d.GetOk("prefix"); ok && !localEndpointTypeFound {

		localEndpointTypeFound = true
		localEndpointInt := &ves_io_schema_network_policy.CreateSpecType_Prefix{}
		localEndpointInt.Prefix = &ves_io_schema.PrefixListType{}
		createSpec.LocalEndpoint = localEndpointInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				localEndpointInt.Prefix.Prefix = ls

			}

		}

	}

	if v, ok := d.GetOk("prefix_selector"); ok && !localEndpointTypeFound {

		localEndpointTypeFound = true
		localEndpointInt := &ves_io_schema_network_policy.CreateSpecType_PrefixSelector{}
		localEndpointInt.PrefixSelector = &ves_io_schema.LabelSelectorType{}
		createSpec.LocalEndpoint = localEndpointInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				localEndpointInt.PrefixSelector.Expressions = ls

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra NetworkPolicy object with struct: %+v", createReq)

	createNetworkPolicyResp, err := client.CreateObject(context.Background(), ves_io_schema_network_policy.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating NetworkPolicy: %s", err)
	}
	d.SetId(createNetworkPolicyResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraNetworkPolicyRead(d, meta)
}

func resourceVolterraNetworkPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_network_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkPolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkPolicy %q: %s", d.Id(), err)
	}
	return setNetworkPolicyFields(client, d, resp)
}

func setNetworkPolicyFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraNetworkPolicyUpdate updates NetworkPolicy resource
func resourceVolterraNetworkPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_network_policy.ReplaceSpecType{}
	updateReq := &ves_io_schema_network_policy.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("egress_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		egressRulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.EgressRules = egressRulesInt
		for i, ps := range sl {

			erMapToStrVal := ps.(map[string]interface{})
			egressRulesInt[i] = &ves_io_schema.ObjectRefType{}

			egressRulesInt[i].Kind = "network_policy_rule"

			if v, ok := erMapToStrVal["name"]; ok && !isIntfNil(v) {
				egressRulesInt[i].Name = v.(string)
			}

			if v, ok := erMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				egressRulesInt[i].Namespace = v.(string)
			}

			if v, ok := erMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				egressRulesInt[i].Tenant = v.(string)
			}

			if v, ok := erMapToStrVal["uid"]; ok && !isIntfNil(v) {
				egressRulesInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("ingress_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		ingressRulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.IngressRules = ingressRulesInt
		for i, ps := range sl {

			irMapToStrVal := ps.(map[string]interface{})
			ingressRulesInt[i] = &ves_io_schema.ObjectRefType{}

			ingressRulesInt[i].Kind = "network_policy_rule"

			if v, ok := irMapToStrVal["name"]; ok && !isIntfNil(v) {
				ingressRulesInt[i].Name = v.(string)
			}

			if v, ok := irMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				ingressRulesInt[i].Namespace = v.(string)
			}

			if v, ok := irMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				ingressRulesInt[i].Tenant = v.(string)
			}

			if v, ok := irMapToStrVal["uid"]; ok && !isIntfNil(v) {
				ingressRulesInt[i].Uid = v.(string)
			}

		}

	}

	localEndpointTypeFound := false

	if v, ok := d.GetOk("prefix"); ok && !localEndpointTypeFound {

		localEndpointTypeFound = true
		localEndpointInt := &ves_io_schema_network_policy.ReplaceSpecType_Prefix{}
		localEndpointInt.Prefix = &ves_io_schema.PrefixListType{}
		updateSpec.LocalEndpoint = localEndpointInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				localEndpointInt.Prefix.Prefix = ls

			}

		}

	}

	if v, ok := d.GetOk("prefix_selector"); ok && !localEndpointTypeFound {

		localEndpointTypeFound = true
		localEndpointInt := &ves_io_schema_network_policy.ReplaceSpecType_PrefixSelector{}
		localEndpointInt.PrefixSelector = &ves_io_schema.LabelSelectorType{}
		updateSpec.LocalEndpoint = localEndpointInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				localEndpointInt.PrefixSelector.Expressions = ls

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra NetworkPolicy obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_network_policy.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating NetworkPolicy: %s", err)
	}

	return resourceVolterraNetworkPolicyRead(d, meta)
}

func resourceVolterraNetworkPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_network_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkPolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkPolicy before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra NetworkPolicy obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_network_policy.ObjectType, namespace, name)
}
