//
// Copyright (c) 2023 F5 Inc. All rights reserved.
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
	ves_io_schema_public_ip "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/public_ip"
)

// resourceVolterraPublicIp is implementation of Volterra's PublicIp resources
func resourceVolterraPublicIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraPublicIpCreate,
		Read:   resourceVolterraPublicIpRead,
		Update: resourceVolterraPublicIpUpdate,
		Delete: resourceVolterraPublicIpDelete,

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

			"virtual_sites": {
				Type:     schema.TypeSet,
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

// resourceVolterraPublicIpCreate creates PublicIp resource
func resourceVolterraPublicIpCreate(d *schema.ResourceData, meta interface{}) error {
	if err := resourceVolterraPublicIpRead(d, meta); err != nil {
		return err
	}
	return resourceVolterraPublicIpUpdate(d, meta)
}

func resourceVolterraPublicIpRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_public_ip.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] PublicIp %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra PublicIp %q: %s", d.Id(), err)
	}
	return setPublicIpFields(client, d, resp)
}

func setPublicIpFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	d.SetId(metadata.GetUid())

	return nil
}

// resourceVolterraPublicIpUpdate updates PublicIp resource
func resourceVolterraPublicIpUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_public_ip.ReplaceSpecType{}
	updateReq := &ves_io_schema_public_ip.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {
		ms := map[string]string{}
		for k, v := range v.(map[string]interface{}) {
			ms[k] = v.(string)
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
			ms[k] = v.(string)
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

	if v, ok := d.GetOk("virtual_sites"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		virtualSitesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.VirtualSites = virtualSitesInt
		for i, ps := range sl {

			vsMapToStrVal := ps.(map[string]interface{})
			virtualSitesInt[i] = &ves_io_schema.ObjectRefType{}

			virtualSitesInt[i].Kind = "virtual_site"

			if v, ok := vsMapToStrVal["name"]; ok && !isIntfNil(v) {
				virtualSitesInt[i].Name = v.(string)
			}

			if v, ok := vsMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				virtualSitesInt[i].Namespace = v.(string)
			}

			if v, ok := vsMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				virtualSitesInt[i].Tenant = v.(string)
			}

			if v, ok := vsMapToStrVal["uid"]; ok && !isIntfNil(v) {
				virtualSitesInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra PublicIp obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_public_ip.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating PublicIp: %s", err)
	}

	return resourceVolterraPublicIpRead(d, meta)
}

func resourceVolterraPublicIpDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_public_ip.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] PublicIp %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra PublicIp before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra PublicIp obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_public_ip.ObjectType, namespace, name)
}
