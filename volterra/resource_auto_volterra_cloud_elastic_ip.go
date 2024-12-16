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
	ves_io_schema_cloud_elastic_ip "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cloud_elastic_ip"
)

// resourceVolterraCloudElasticIp is implementation of Volterra's CloudElasticIp resources
func resourceVolterraCloudElasticIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraCloudElasticIpCreate,
		Read:   resourceVolterraCloudElasticIpRead,
		Update: resourceVolterraCloudElasticIpUpdate,
		Delete: resourceVolterraCloudElasticIpDelete,

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

			"count_cloud_elastic_ip": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"site_ref": {

				Type:     schema.TypeList,
				Required: true,
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

// resourceVolterraCloudElasticIpCreate creates CloudElasticIp resource
func resourceVolterraCloudElasticIpCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_cloud_elastic_ip.CreateSpecType{}
	createReq := &ves_io_schema_cloud_elastic_ip.CreateRequest{
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

	//count_cloud_elastic_ip
	if v, ok := d.GetOk("count_cloud_elastic_ip"); ok && !isIntfNil(v) {

		createSpec.Count =
			uint32(v.(int))

	}

	//site_ref
	if v, ok := d.GetOk("site_ref"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		siteRefInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.SiteRef = siteRefInt
		for i, ps := range sl {

			srMapToStrVal := ps.(map[string]interface{})
			siteRefInt[i] = &ves_io_schema.ObjectRefType{}

			siteRefInt[i].Kind = "site"

			if v, ok := srMapToStrVal["name"]; ok && !isIntfNil(v) {
				siteRefInt[i].Name = v.(string)
			}

			if v, ok := srMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				siteRefInt[i].Namespace = v.(string)
			}

			if v, ok := srMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				siteRefInt[i].Tenant = v.(string)
			}

			if v, ok := srMapToStrVal["uid"]; ok && !isIntfNil(v) {
				siteRefInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra CloudElasticIp object with struct: %+v", createReq)

	createCloudElasticIpResp, err := client.CreateObject(context.Background(), ves_io_schema_cloud_elastic_ip.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating CloudElasticIp: %s", err)
	}
	d.SetId(createCloudElasticIpResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraCloudElasticIpRead(d, meta)
}

func resourceVolterraCloudElasticIpRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_cloud_elastic_ip.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] CloudElasticIp %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra CloudElasticIp %q: %s", d.Id(), err)
	}
	return setCloudElasticIpFields(client, d, resp)
}

func setCloudElasticIpFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraCloudElasticIpUpdate updates CloudElasticIp resource
func resourceVolterraCloudElasticIpUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_cloud_elastic_ip.ReplaceSpecType{}
	updateReq := &ves_io_schema_cloud_elastic_ip.ReplaceRequest{
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

	if v, ok := d.GetOk("count_cloud_elastic_ip"); ok && !isIntfNil(v) {

		updateSpec.Count =
			uint32(v.(int))

	}

	if v, ok := d.GetOk("site_ref"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		siteRefInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.SiteRef = siteRefInt
		for i, ps := range sl {

			srMapToStrVal := ps.(map[string]interface{})
			siteRefInt[i] = &ves_io_schema.ObjectRefType{}

			siteRefInt[i].Kind = "site"

			if v, ok := srMapToStrVal["name"]; ok && !isIntfNil(v) {
				siteRefInt[i].Name = v.(string)
			}

			if v, ok := srMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				siteRefInt[i].Namespace = v.(string)
			}

			if v, ok := srMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				siteRefInt[i].Tenant = v.(string)
			}

			if v, ok := srMapToStrVal["uid"]; ok && !isIntfNil(v) {
				siteRefInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra CloudElasticIp obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_cloud_elastic_ip.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating CloudElasticIp: %s", err)
	}

	return resourceVolterraCloudElasticIpRead(d, meta)
}

func resourceVolterraCloudElasticIpDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_cloud_elastic_ip.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] CloudElasticIp %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra CloudElasticIp before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra CloudElasticIp obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_cloud_elastic_ip.ObjectType, namespace, name)
}
