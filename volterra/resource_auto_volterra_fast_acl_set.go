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
	ves_io_schema_fast_acl_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/fast_acl_set"
)

// resourceVolterraFastAclSet is implementation of Volterra's FastAclSet resources
func resourceVolterraFastAclSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraFastAclSetCreate,
		Read:   resourceVolterraFastAclSetRead,
		Update: resourceVolterraFastAclSetUpdate,
		Delete: resourceVolterraFastAclSetDelete,

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

			"acl_list": {

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

// resourceVolterraFastAclSetCreate creates FastAclSet resource
func resourceVolterraFastAclSetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_fast_acl_set.CreateSpecType{}
	createReq := &ves_io_schema_fast_acl_set.CreateRequest{
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

	if v, ok := d.GetOk("acl_list"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		aclListInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.AclList = aclListInt
		for i, ps := range sl {

			alMapToStrVal := ps.(map[string]interface{})
			aclListInt[i] = &ves_io_schema.ObjectRefType{}

			aclListInt[i].Kind = "fast_acl"

			if v, ok := alMapToStrVal["name"]; ok && !isIntfNil(v) {
				aclListInt[i].Name = v.(string)
			}

			if v, ok := alMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				aclListInt[i].Namespace = v.(string)
			}

			if v, ok := alMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				aclListInt[i].Tenant = v.(string)
			}

			if v, ok := alMapToStrVal["uid"]; ok && !isIntfNil(v) {
				aclListInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra FastAclSet object with struct: %+v", createReq)

	createFastAclSetResp, err := client.CreateObject(context.Background(), ves_io_schema_fast_acl_set.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating FastAclSet: %s", err)
	}
	d.SetId(createFastAclSetResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraFastAclSetRead(d, meta)
}

func resourceVolterraFastAclSetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_fast_acl_set.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] FastAclSet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra FastAclSet %q: %s", d.Id(), err)
	}
	return setFastAclSetFields(client, d, resp)
}

func setFastAclSetFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraFastAclSetUpdate updates FastAclSet resource
func resourceVolterraFastAclSetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_fast_acl_set.ReplaceSpecType{}
	updateReq := &ves_io_schema_fast_acl_set.ReplaceRequest{
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

	if v, ok := d.GetOk("acl_list"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		aclListInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.AclList = aclListInt
		for i, ps := range sl {

			alMapToStrVal := ps.(map[string]interface{})
			aclListInt[i] = &ves_io_schema.ObjectRefType{}

			aclListInt[i].Kind = "fast_acl"

			if v, ok := alMapToStrVal["name"]; ok && !isIntfNil(v) {
				aclListInt[i].Name = v.(string)
			}

			if v, ok := alMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				aclListInt[i].Namespace = v.(string)
			}

			if v, ok := alMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				aclListInt[i].Tenant = v.(string)
			}

			if v, ok := alMapToStrVal["uid"]; ok && !isIntfNil(v) {
				aclListInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra FastAclSet obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_fast_acl_set.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating FastAclSet: %s", err)
	}

	return resourceVolterraFastAclSetRead(d, meta)
}

func resourceVolterraFastAclSetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_fast_acl_set.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] FastAclSet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra FastAclSet before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra FastAclSet obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_fast_acl_set.ObjectType, namespace, name)
}
