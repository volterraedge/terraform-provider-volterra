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
	ves_io_schema_uztna_uztna_healthcheck "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/uztna/uztna_healthcheck"
)

// resourceVolterraUztnaHealthcheck is implementation of Volterra's UztnaHealthcheck resources
func resourceVolterraUztnaHealthcheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraUztnaHealthcheckCreate,
		Read:   resourceVolterraUztnaHealthcheckRead,
		Update: resourceVolterraUztnaHealthcheckUpdate,
		Delete: resourceVolterraUztnaHealthcheckDelete,

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

			"tcp_health_check": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"expected_response": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"send_payload": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"interval": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"timeout": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

// resourceVolterraUztnaHealthcheckCreate creates UztnaHealthcheck resource
func resourceVolterraUztnaHealthcheckCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_uztna_uztna_healthcheck.CreateSpecType{}
	createReq := &ves_io_schema_uztna_uztna_healthcheck.CreateRequest{
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

	//health_check

	healthCheckTypeFound := false

	if v, ok := d.GetOk("tcp_health_check"); ok && !isIntfNil(v) && !healthCheckTypeFound {

		healthCheckTypeFound = true
		healthCheckInt := &ves_io_schema_uztna_uztna_healthcheck.CreateSpecType_TcpHealthCheck{}
		healthCheckInt.TcpHealthCheck = &ves_io_schema_uztna_uztna_healthcheck.TcpHealthCheck{}
		createSpec.HealthCheck = healthCheckInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["expected_response"]; ok && !isIntfNil(v) {

					healthCheckInt.TcpHealthCheck.ExpectedResponse = v.(string)

				}

				if v, ok := cs["send_payload"]; ok && !isIntfNil(v) {

					healthCheckInt.TcpHealthCheck.SendPayload = v.(string)

				}

			}
		}

	}

	//interval
	if v, ok := d.GetOk("interval"); ok && !isIntfNil(v) {

		createSpec.Interval =
			uint32(v.(int))

	}

	//timeout
	if v, ok := d.GetOk("timeout"); ok && !isIntfNil(v) {

		createSpec.Timeout =
			uint32(v.(int))

	}

	log.Printf("[DEBUG] Creating Volterra UztnaHealthcheck object with struct: %+v", createReq)

	createUztnaHealthcheckResp, err := client.CreateObject(context.Background(), ves_io_schema_uztna_uztna_healthcheck.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating UztnaHealthcheck: %s", err)
	}
	d.SetId(createUztnaHealthcheckResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraUztnaHealthcheckRead(d, meta)
}

func resourceVolterraUztnaHealthcheckRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_uztna_uztna_healthcheck.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaHealthcheck %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaHealthcheck %q: %s", d.Id(), err)
	}
	return setUztnaHealthcheckFields(client, d, resp)
}

func setUztnaHealthcheckFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraUztnaHealthcheckUpdate updates UztnaHealthcheck resource
func resourceVolterraUztnaHealthcheckUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_uztna_uztna_healthcheck.ReplaceSpecType{}
	updateReq := &ves_io_schema_uztna_uztna_healthcheck.ReplaceRequest{
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

	healthCheckTypeFound := false

	if v, ok := d.GetOk("tcp_health_check"); ok && !isIntfNil(v) && !healthCheckTypeFound {

		healthCheckTypeFound = true
		healthCheckInt := &ves_io_schema_uztna_uztna_healthcheck.ReplaceSpecType_TcpHealthCheck{}
		healthCheckInt.TcpHealthCheck = &ves_io_schema_uztna_uztna_healthcheck.TcpHealthCheck{}
		updateSpec.HealthCheck = healthCheckInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["expected_response"]; ok && !isIntfNil(v) {

					healthCheckInt.TcpHealthCheck.ExpectedResponse = v.(string)

				}

				if v, ok := cs["send_payload"]; ok && !isIntfNil(v) {

					healthCheckInt.TcpHealthCheck.SendPayload = v.(string)

				}

			}
		}

	}

	if v, ok := d.GetOk("interval"); ok && !isIntfNil(v) {

		updateSpec.Interval =
			uint32(v.(int))

	}

	if v, ok := d.GetOk("timeout"); ok && !isIntfNil(v) {

		updateSpec.Timeout =
			uint32(v.(int))

	}

	log.Printf("[DEBUG] Updating Volterra UztnaHealthcheck obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_uztna_uztna_healthcheck.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating UztnaHealthcheck: %s", err)
	}

	return resourceVolterraUztnaHealthcheckRead(d, meta)
}

func resourceVolterraUztnaHealthcheckDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_uztna_uztna_healthcheck.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaHealthcheck %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaHealthcheck before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra UztnaHealthcheck obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_uztna_uztna_healthcheck.ObjectType, namespace, name, opts...)
}
