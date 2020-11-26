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
	ves_io_schema_token "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/token"
)

// resourceVolterraToken is implementation of Volterra's Token resources
func resourceVolterraToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraTokenCreate,
		Read:   resourceVolterraTokenRead,
		Update: resourceVolterraTokenUpdate,
		Delete: resourceVolterraTokenDelete,

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
			"created_by_tf": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// resourceVolterraTokenCreate creates Token resource
func resourceVolterraTokenCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var name, namespace string
	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		name = v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		namespace = v.(string)
	}

	tokenListResp, err := client.ListObjects(context.Background(), ves_io_schema_token.ObjectType, namespace)
	if err != nil || len(tokenListResp) == 0 {
		createMeta := &ves_io_schema.ObjectCreateMetaType{
			Name:      name,
			Namespace: namespace,
		}
		createSpec := &ves_io_schema_token.CreateSpecType{}
		createReq := &ves_io_schema_token.CreateRequest{
			Metadata: createMeta,
			Spec:     createSpec,
		}

		if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {
			ms := map[string]string{}
			for k, v := range v.(map[string]interface{}) {
				ms[k] = v.(string)
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
				ms[k] = v.(string)
			}
			createMeta.Labels = ms
		}

		log.Printf("[DEBUG] Creating Volterra Token object with struct: %+v", createReq)

		createTokenResp, err := client.CreateObject(context.Background(), ves_io_schema_token.ObjectType, createReq)
		if err != nil {
			return fmt.Errorf("error creating Token: %s", err)
		}
		d.SetId(createTokenResp.GetObjMetadata().GetUid())
		d.Set("created_by_tf", true)
		d.Set("tenant_name", createTokenResp.GetObjSystemMetadata().GetTenant())
		return resourceVolterraTokenRead(d, meta)
	}
	token := tokenListResp[0].(*ves_io_schema_token.ListResponseItem)
	d.SetId(token.GetUid())
	d.Set("tenant_name", token.GetTenant())
	d.Set("created_by_tf", false)
	return nil
}

func resourceVolterraTokenRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createdHere := d.Get("created_by_tf").(bool)

	if createdHere {
		name := d.Get("name").(string)
		namespace := d.Get("namespace").(string)

		resp, err := client.GetObject(context.Background(), ves_io_schema_token.ObjectType, namespace, name)
		if err != nil {
			if strings.Contains(err.Error(), "status code 404") {
				log.Printf("[INFO] Token %s no longer exists", d.Id())
				d.SetId("")
				return nil
			}
			return fmt.Errorf("Error finding Volterra Token %q: %s", d.Id(), err)
		}
		return setTokenFields(client, d, resp)
	}
	// hack to get the tenant_name, during update
	namespace := d.Get("namespace").(string)
	tokenListResp, err := client.ListObjects(context.Background(), ves_io_schema_token.ObjectType, namespace)
	if err != nil {
		return fmt.Errorf("Error getting list of Volterra Tokens during read %q: %s", d.Id(), err)
	}
	token := tokenListResp[0].(*ves_io_schema_token.ListResponseItem)
	d.Set("tenant_name", token.GetTenant())
	return nil
}

func setTokenFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {

	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	d.Set("tenant_name", resp.GetObjSystemMetadata().GetTenant())

	d.SetId(metadata.GetUid())

	return nil
}

// resourceVolterraTokenUpdate updates Token resource
func resourceVolterraTokenUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createdHere := d.Get("created_by_tf").(bool)
	if createdHere {
		updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
		updateSpec := &ves_io_schema_token.ReplaceSpecType{}
		updateReq := &ves_io_schema_token.ReplaceRequest{
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

		log.Printf("[DEBUG] Updating Volterra Token obj with struct: %+v", updateReq)

		err := client.ReplaceObject(context.Background(), ves_io_schema_token.ObjectType, updateReq)
		if err != nil {
			return fmt.Errorf("error updating Token: %s", err)
		}

		return resourceVolterraTokenRead(d, meta)
	}

	return nil

}

func resourceVolterraTokenDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_token.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Token %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Token before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Token obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_token.ObjectType, namespace, name)
}
