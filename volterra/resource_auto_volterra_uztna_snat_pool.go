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
	ves_io_schema_uztna_application_uztna_snat_pool "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/uztna/application/uztna_snat_pool"
)

// resourceVolterraUztnaSnatPool is implementation of Volterra's UztnaSnatPool resources
func resourceVolterraUztnaSnatPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraUztnaSnatPoolCreate,
		Read:   resourceVolterraUztnaSnatPoolRead,
		Update: resourceVolterraUztnaSnatPoolUpdate,
		Delete: resourceVolterraUztnaSnatPoolDelete,

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

			"ip_addresses": {

				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ipv4": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"addr": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"ipv6": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"addr": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraUztnaSnatPoolCreate creates UztnaSnatPool resource
func resourceVolterraUztnaSnatPoolCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_uztna_application_uztna_snat_pool.CreateSpecType{}
	createReq := &ves_io_schema_uztna_application_uztna_snat_pool.CreateRequest{
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

	//ip_addresses
	if v, ok := d.GetOk("ip_addresses"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		ipAddresses := make([]*ves_io_schema.IpAddressType, len(sl))
		createSpec.IpAddresses = ipAddresses
		for i, set := range sl {
			if set != nil {
				ipAddresses[i] = &ves_io_schema.IpAddressType{}
				ipAddressesMapStrToI := set.(map[string]interface{})

				verTypeFound := false

				if v, ok := ipAddressesMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

					verTypeFound = true
					verInt := &ves_io_schema.IpAddressType_Ipv4{}
					verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
					ipAddresses[i].Ver = verInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv4.Addr = v.(string)

							}

						}
					}

				}

				if v, ok := ipAddressesMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

					verTypeFound = true
					verInt := &ves_io_schema.IpAddressType_Ipv6{}
					verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
					ipAddresses[i].Ver = verInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv6.Addr = v.(string)

							}

						}
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Creating Volterra UztnaSnatPool object with struct: %+v", createReq)

	createUztnaSnatPoolResp, err := client.CreateObject(context.Background(), ves_io_schema_uztna_application_uztna_snat_pool.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating UztnaSnatPool: %s", err)
	}
	d.SetId(createUztnaSnatPoolResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraUztnaSnatPoolRead(d, meta)
}

func resourceVolterraUztnaSnatPoolRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_uztna_application_uztna_snat_pool.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaSnatPool %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaSnatPool %q: %s", d.Id(), err)
	}
	return setUztnaSnatPoolFields(client, d, resp)
}

func setUztnaSnatPoolFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraUztnaSnatPoolUpdate updates UztnaSnatPool resource
func resourceVolterraUztnaSnatPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_uztna_application_uztna_snat_pool.ReplaceSpecType{}
	updateReq := &ves_io_schema_uztna_application_uztna_snat_pool.ReplaceRequest{
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

	if v, ok := d.GetOk("ip_addresses"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		ipAddresses := make([]*ves_io_schema.IpAddressType, len(sl))
		updateSpec.IpAddresses = ipAddresses
		for i, set := range sl {
			if set != nil {
				ipAddresses[i] = &ves_io_schema.IpAddressType{}
				ipAddressesMapStrToI := set.(map[string]interface{})

				verTypeFound := false

				if v, ok := ipAddressesMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

					verTypeFound = true
					verInt := &ves_io_schema.IpAddressType_Ipv4{}
					verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
					ipAddresses[i].Ver = verInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv4.Addr = v.(string)

							}

						}
					}

				}

				if v, ok := ipAddressesMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

					verTypeFound = true
					verInt := &ves_io_schema.IpAddressType_Ipv6{}
					verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
					ipAddresses[i].Ver = verInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv6.Addr = v.(string)

							}

						}
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Updating Volterra UztnaSnatPool obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_uztna_application_uztna_snat_pool.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating UztnaSnatPool: %s", err)
	}

	return resourceVolterraUztnaSnatPoolRead(d, meta)
}

func resourceVolterraUztnaSnatPoolDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_uztna_application_uztna_snat_pool.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaSnatPool %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaSnatPool before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra UztnaSnatPool obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_uztna_application_uztna_snat_pool.ObjectType, namespace, name, opts...)
}
