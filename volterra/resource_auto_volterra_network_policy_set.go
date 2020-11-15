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
	ves_io_schema_network_policy_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_policy_set"
)

// resourceVolterraNetworkPolicySet is implementation of Volterra's NetworkPolicySet resources
func resourceVolterraNetworkPolicySet() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraNetworkPolicySetCreate,
		Read:   resourceVolterraNetworkPolicySetRead,
		Update: resourceVolterraNetworkPolicySetUpdate,
		Delete: resourceVolterraNetworkPolicySetDelete,

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

			"policies": {

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
		},
	}
}

// resourceVolterraNetworkPolicySetCreate creates NetworkPolicySet resource
func resourceVolterraNetworkPolicySetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_network_policy_set.CreateSpecType{}
	createReq := &ves_io_schema_network_policy_set.CreateRequest{
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

	if v, ok := d.GetOk("policies"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		policiesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.Policies = policiesInt
		for i, ps := range sl {

			pMapToStrVal := ps.(map[string]interface{})
			policiesInt[i] = &ves_io_schema.ObjectRefType{}

			policiesInt[i].Kind = "network_policy"

			if v, ok := pMapToStrVal["name"]; ok && !isIntfNil(v) {
				policiesInt[i].Name = v.(string)
			}

			if v, ok := pMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				policiesInt[i].Namespace = v.(string)
			}

			if v, ok := pMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				policiesInt[i].Tenant = v.(string)
			}

			if v, ok := pMapToStrVal["uid"]; ok && !isIntfNil(v) {
				policiesInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra NetworkPolicySet object with struct: %+v", createReq)

	createNetworkPolicySetResp, err := client.CreateObject(context.Background(), ves_io_schema_network_policy_set.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating NetworkPolicySet: %s", err)
	}
	d.SetId(createNetworkPolicySetResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraNetworkPolicySetRead(d, meta)
}

func resourceVolterraNetworkPolicySetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_network_policy_set.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkPolicySet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkPolicySet %q: %s", d.Id(), err)
	}
	return setNetworkPolicySetFields(client, d, resp)
}

func setNetworkPolicySetFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraNetworkPolicySetUpdate updates NetworkPolicySet resource
func resourceVolterraNetworkPolicySetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_network_policy_set.ReplaceSpecType{}
	updateReq := &ves_io_schema_network_policy_set.ReplaceRequest{
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

	if v, ok := d.GetOk("policies"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		policiesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.Policies = policiesInt
		for i, ps := range sl {

			pMapToStrVal := ps.(map[string]interface{})
			policiesInt[i] = &ves_io_schema.ObjectRefType{}

			policiesInt[i].Kind = "network_policy"

			if v, ok := pMapToStrVal["name"]; ok && !isIntfNil(v) {
				policiesInt[i].Name = v.(string)
			}

			if v, ok := pMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				policiesInt[i].Namespace = v.(string)
			}

			if v, ok := pMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				policiesInt[i].Tenant = v.(string)
			}

			if v, ok := pMapToStrVal["uid"]; ok && !isIntfNil(v) {
				policiesInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra NetworkPolicySet obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_network_policy_set.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating NetworkPolicySet: %s", err)
	}

	return resourceVolterraNetworkPolicySetRead(d, meta)
}

func resourceVolterraNetworkPolicySetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_network_policy_set.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkPolicySet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkPolicySet before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra NetworkPolicySet obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_network_policy_set.ObjectType, namespace, name)
}