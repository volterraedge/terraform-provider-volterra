//
// Copyright (c) 2023 F5 Inc. All rights reserved.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"tenant_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeInt,
				Optional: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(int)
					if value != 0 && value != 1 {
						errors = append(errors, fmt.Errorf("%q must be either 0 or 1", k))
					}
					return
				},
			},
			"site_name": {
				Type:     schema.TypeString,
				Optional: true,
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

	createMeta := &ves_io_schema.ObjectCreateMetaType{
		Name:      name,
		Namespace: namespace,
	}
	createSpec := &ves_io_schema_token.CreateSpecType{}
	createReq := &ves_io_schema_token.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("type"); ok {
		createSpec.Type = ves_io_schema_token.Type(int32(v.(int)))
	}

	if v, ok := d.GetOk("site_name"); ok && !isIntfNil(v) {
		createSpec.SiteName = v.(string)
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

	getResp, _ := client.GetObject(context.Background(), ves_io_schema_token.ObjectType, namespace, name)
	if getResp != nil {
		gcSpec := getResp.GetObjSpec().(*ves_io_schema_token.SpecType).GetGcSpec()
		specType := gcSpec.GetType()
		if specType == createSpec.Type {
			if specType == ves_io_schema_token.JWT {
				d.SetId(gcSpec.GetContent())
			} else {
				d.SetId(getResp.GetObjSystemMetadata().GetUid())
			}
		} else {
			if createSpec.Type == ves_io_schema_token.JWT {
				return fmt.Errorf("A NORMAL token with the same name already exists. Please delete the existing token if it is no longer in use, or create a new JWT token with a different name.")
			} else {
				return fmt.Errorf("A JWT token with the same name already exists. Please delete the existing token if it is no longer in use, or create a new NORMAL token with a different name.")
			}
		}
		d.Set("tenant_name", getResp.GetObjSystemMetadata().GetTenant())
		return resourceVolterraTokenRead(d, meta)
	}
	log.Printf("[DEBUG] Creating Volterra Token object with struct: %+v", createReq)

	createTokenResp, err := client.CreateObject(context.Background(), ves_io_schema_token.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Token: %s", err)
	}
	if createSpec.Type == ves_io_schema_token.JWT {
		d.SetId(createTokenResp.GetObjSpec().(*ves_io_schema_token.SpecType).GetGcSpec().GetContent())
	} else {
		d.SetId(createTokenResp.GetObjSystemMetadata().GetUid())
	}
	d.Set("tenant_name", createTokenResp.GetObjSystemMetadata().GetTenant())
	return resourceVolterraTokenRead(d, meta)
}

func resourceVolterraTokenRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
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

func setTokenFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {

	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraTokenUpdate updates Token resource
func resourceVolterraTokenUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
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

func resourceVolterraTokenDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	getResp, err := client.GetObject(context.Background(), ves_io_schema_token.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Token %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Token before deleting %q: %s", d.Id(), err)
	}

	creationTimeString, err := convertTimestampToString(getResp.GetObjSystemMetadata().GetCreationTimestamp())
	if err != nil {
		return fmt.Errorf("error converting expiration_timestamp : %s", err)
	}

	creationTime, err := time.Parse(time.RFC3339, creationTimeString)
	if err != nil {
		return fmt.Errorf("error parsing creation timestamp: : %s", err)
	}

	cutoffTime, err := time.Parse(time.RFC3339, "2025-01-01T00:00:00Z")
	if err != nil {
		return fmt.Errorf("error parsing cutoff date: : %s", err)
	}

	if creationTime.After(cutoffTime) {
		log.Printf("[DEBUG] Deleting Volterra Token obj with name %+v in namespace %+v", name, namespace)
		return client.DeleteObject(context.Background(), ves_io_schema_token.ObjectType, namespace, name)
	}
	return nil
}
