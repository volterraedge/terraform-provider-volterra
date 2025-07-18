//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-tf-provider. DO NOT EDIT.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
	ves_io_schema_virtual_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_site"
)

// resourceVolterraVirtualSite is implementation of Volterra's VirtualSite resources
func resourceVolterraVirtualSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraVirtualSiteCreate,
		Read:   resourceVolterraVirtualSiteRead,
		Update: resourceVolterraVirtualSiteUpdate,
		Delete: resourceVolterraVirtualSiteDelete,

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
				ForceNew: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"site_selector": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
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

			"site_type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// resourceVolterraVirtualSiteCreate creates VirtualSite resource
func resourceVolterraVirtualSiteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_virtual_site.CreateSpecType{}
	createReq := &ves_io_schema_virtual_site.CreateRequest{
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

	//site_selector
	if v, ok := d.GetOk("site_selector"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		siteSelector := &ves_io_schema.LabelSelectorType{}
		createSpec.SiteSelector = siteSelector
		for _, set := range sl {
			if set != nil {
				siteSelectorMapStrToI := set.(map[string]interface{})

				if w, ok := siteSelectorMapStrToI["expressions"]; ok && !isIntfNil(w) {
					ls := make([]string, len(w.([]interface{})))
					for i, v := range w.([]interface{}) {
						if v == nil {
							return fmt.Errorf("please provide valid non-empty string value of field expressions")
						}
						if str, ok := v.(string); ok {
							ls[i] = str
						}
					}
					siteSelector.Expressions = ls
				}

			}
		}

	}

	//site_type
	if v, ok := d.GetOk("site_type"); ok && !isIntfNil(v) {

		createSpec.SiteType = ves_io_schema_site.SiteType(ves_io_schema_site.SiteType_value[v.(string)])

	}

	log.Printf("[DEBUG] Creating Volterra VirtualSite object with struct: %+v", createReq)

	createVirtualSiteResp, err := client.CreateObject(context.Background(), ves_io_schema_virtual_site.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating VirtualSite: %s", err)
	}
	d.SetId(createVirtualSiteResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraVirtualSiteRead(d, meta)
}

func resourceVolterraVirtualSiteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_virtual_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] VirtualSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra VirtualSite %q: %s", d.Id(), err)
	}
	return setVirtualSiteFields(client, d, resp)
}

func setVirtualSiteFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraVirtualSiteUpdate updates VirtualSite resource
func resourceVolterraVirtualSiteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_virtual_site.ReplaceSpecType{}
	updateReq := &ves_io_schema_virtual_site.ReplaceRequest{
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

	log.Printf("[DEBUG] Updating Volterra VirtualSite obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_virtual_site.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating VirtualSite: %s", err)
	}

	return resourceVolterraVirtualSiteRead(d, meta)
}

func resourceVolterraVirtualSiteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_virtual_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] VirtualSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra VirtualSite before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra VirtualSite obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_virtual_site.ObjectType, namespace, name, opts...)
}
