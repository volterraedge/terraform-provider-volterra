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
	ves_io_schema_container_registry "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/container_registry"
)

// resourceVolterraContainerRegistry is implementation of Volterra's ContainerRegistry resources
func resourceVolterraContainerRegistry() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraContainerRegistryCreate,
		Read:   resourceVolterraContainerRegistryRead,
		Update: resourceVolterraContainerRegistryUpdate,
		Delete: resourceVolterraContainerRegistryDelete,

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

			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"password": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"blindfold_secret_info": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"decryption_provider": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"location": {
										Type:     schema.TypeString,
										Required: true,
									},

									"store_provider": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"clear_secret_info": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"provider": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"url": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"registry": {
				Type:     schema.TypeString,
				Required: true,
			},

			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// resourceVolterraContainerRegistryCreate creates ContainerRegistry resource
func resourceVolterraContainerRegistryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_container_registry.CreateSpecType{}
	createReq := &ves_io_schema_container_registry.CreateRequest{
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

	//email
	if v, ok := d.GetOk("email"); ok && !isIntfNil(v) {

		createSpec.Email =
			v.(string)

	}

	//password
	if v, ok := d.GetOk("password"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		password := &ves_io_schema.SecretType{}
		createSpec.Password = password
		for _, set := range sl {
			if set != nil {
				passwordMapStrToI := set.(map[string]interface{})

				secretInfoOneofTypeFound := false

				if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

					secretInfoOneofTypeFound = true
					secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
					secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
					password.SecretInfoOneof = secretInfoOneofInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["decryption_provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.BlindfoldSecretInfo.DecryptionProvider = v.(string)

							}

							if v, ok := cs["location"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.BlindfoldSecretInfo.Location = v.(string)

							}

							if v, ok := cs["store_provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.BlindfoldSecretInfo.StoreProvider = v.(string)

							}

						}
					}

				}

				if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

					secretInfoOneofTypeFound = true
					secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
					secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
					password.SecretInfoOneof = secretInfoOneofInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)

							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)

							}

						}
					}

				}

			}
		}

	}

	//registry
	if v, ok := d.GetOk("registry"); ok && !isIntfNil(v) {

		createSpec.Registry =
			v.(string)

	}

	//user_name
	if v, ok := d.GetOk("user_name"); ok && !isIntfNil(v) {

		createSpec.UserName =
			v.(string)

	}

	log.Printf("[DEBUG] Creating Volterra ContainerRegistry object with struct: %+v", createReq)

	createContainerRegistryResp, err := client.CreateObject(context.Background(), ves_io_schema_container_registry.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating ContainerRegistry: %s", err)
	}
	d.SetId(createContainerRegistryResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraContainerRegistryRead(d, meta)
}

func resourceVolterraContainerRegistryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_container_registry.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ContainerRegistry %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ContainerRegistry %q: %s", d.Id(), err)
	}
	return setContainerRegistryFields(client, d, resp)
}

func setContainerRegistryFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraContainerRegistryUpdate updates ContainerRegistry resource
func resourceVolterraContainerRegistryUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_container_registry.ReplaceSpecType{}
	updateReq := &ves_io_schema_container_registry.ReplaceRequest{
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

	if v, ok := d.GetOk("email"); ok && !isIntfNil(v) {

		updateSpec.Email =
			v.(string)

	}

	if v, ok := d.GetOk("password"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		password := &ves_io_schema.SecretType{}
		updateSpec.Password = password
		for _, set := range sl {
			if set != nil {
				passwordMapStrToI := set.(map[string]interface{})

				secretInfoOneofTypeFound := false

				if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

					secretInfoOneofTypeFound = true
					secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
					secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
					password.SecretInfoOneof = secretInfoOneofInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["decryption_provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.BlindfoldSecretInfo.DecryptionProvider = v.(string)

							}

							if v, ok := cs["location"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.BlindfoldSecretInfo.Location = v.(string)

							}

							if v, ok := cs["store_provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.BlindfoldSecretInfo.StoreProvider = v.(string)

							}

						}
					}

				}

				if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

					secretInfoOneofTypeFound = true
					secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
					secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
					password.SecretInfoOneof = secretInfoOneofInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)

							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)

							}

						}
					}

				}

			}
		}

	}

	if v, ok := d.GetOk("registry"); ok && !isIntfNil(v) {

		updateSpec.Registry =
			v.(string)

	}

	if v, ok := d.GetOk("user_name"); ok && !isIntfNil(v) {

		updateSpec.UserName =
			v.(string)

	}

	log.Printf("[DEBUG] Updating Volterra ContainerRegistry obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_container_registry.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating ContainerRegistry: %s", err)
	}

	return resourceVolterraContainerRegistryRead(d, meta)
}

func resourceVolterraContainerRegistryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_container_registry.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ContainerRegistry %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ContainerRegistry before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra ContainerRegistry obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_container_registry.ObjectType, namespace, name, opts...)
}
