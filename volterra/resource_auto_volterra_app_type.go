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
	ves_io_schema_app_type "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/app_type"
)

// resourceVolterraAppType is implementation of Volterra's AppType resources
func resourceVolterraAppType() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAppTypeCreate,
		Read:   resourceVolterraAppTypeRead,
		Update: resourceVolterraAppTypeUpdate,
		Delete: resourceVolterraAppTypeDelete,

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

			"business_logic_markup_setting": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"disable": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"features": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraAppTypeCreate creates AppType resource
func resourceVolterraAppTypeCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_app_type.CreateSpecType{}
	createReq := &ves_io_schema_app_type.CreateRequest{
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

	if v, ok := d.GetOk("business_logic_markup_setting"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		businessLogicMarkupSetting := &ves_io_schema_app_type.BusinessLogicMarkupSetting{}
		createSpec.BusinessLogicMarkupSetting = businessLogicMarkupSetting
		for _, set := range sl {

			businessLogicMarkupSettingMapStrToI := set.(map[string]interface{})

			learnFromRedirectTrafficTypeFound := false

			if v, ok := businessLogicMarkupSettingMapStrToI["disable"]; ok && !isIntfNil(v) && !learnFromRedirectTrafficTypeFound {

				learnFromRedirectTrafficTypeFound = true

				if v.(bool) {
					learnFromRedirectTrafficInt := &ves_io_schema_app_type.BusinessLogicMarkupSetting_Disable{}
					learnFromRedirectTrafficInt.Disable = &ves_io_schema.Empty{}
					businessLogicMarkupSetting.LearnFromRedirectTraffic = learnFromRedirectTrafficInt
				}

			}

			if v, ok := businessLogicMarkupSettingMapStrToI["enable"]; ok && !isIntfNil(v) && !learnFromRedirectTrafficTypeFound {

				learnFromRedirectTrafficTypeFound = true

				if v.(bool) {
					learnFromRedirectTrafficInt := &ves_io_schema_app_type.BusinessLogicMarkupSetting_Enable{}
					learnFromRedirectTrafficInt.Enable = &ves_io_schema.Empty{}
					businessLogicMarkupSetting.LearnFromRedirectTraffic = learnFromRedirectTrafficInt
				}

			}

		}

	}

	if v, ok := d.GetOk("features"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		features := make([]*ves_io_schema_app_type.Feature, len(sl))
		createSpec.Features = features
		for i, set := range sl {
			features[i] = &ves_io_schema_app_type.Feature{}

			featuresMapStrToI := set.(map[string]interface{})

			if v, ok := featuresMapStrToI["type"]; ok && !isIntfNil(v) {

				features[i].Type = ves_io_schema_app_type.FeatureType(ves_io_schema_app_type.FeatureType_value[v.(string)])

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra AppType object with struct: %+v", createReq)

	createAppTypeResp, err := client.CreateObject(context.Background(), ves_io_schema_app_type.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating AppType: %s", err)
	}
	d.SetId(createAppTypeResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraAppTypeRead(d, meta)
}

func resourceVolterraAppTypeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_app_type.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AppType %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AppType %q: %s", d.Id(), err)
	}
	return setAppTypeFields(client, d, resp)
}

func setAppTypeFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraAppTypeUpdate updates AppType resource
func resourceVolterraAppTypeUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_app_type.ReplaceSpecType{}
	updateReq := &ves_io_schema_app_type.ReplaceRequest{
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

	if v, ok := d.GetOk("business_logic_markup_setting"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		businessLogicMarkupSetting := &ves_io_schema_app_type.BusinessLogicMarkupSetting{}
		updateSpec.BusinessLogicMarkupSetting = businessLogicMarkupSetting
		for _, set := range sl {

			businessLogicMarkupSettingMapStrToI := set.(map[string]interface{})

			learnFromRedirectTrafficTypeFound := false

			if v, ok := businessLogicMarkupSettingMapStrToI["disable"]; ok && !isIntfNil(v) && !learnFromRedirectTrafficTypeFound {

				learnFromRedirectTrafficTypeFound = true

				if v.(bool) {
					learnFromRedirectTrafficInt := &ves_io_schema_app_type.BusinessLogicMarkupSetting_Disable{}
					learnFromRedirectTrafficInt.Disable = &ves_io_schema.Empty{}
					businessLogicMarkupSetting.LearnFromRedirectTraffic = learnFromRedirectTrafficInt
				}

			}

			if v, ok := businessLogicMarkupSettingMapStrToI["enable"]; ok && !isIntfNil(v) && !learnFromRedirectTrafficTypeFound {

				learnFromRedirectTrafficTypeFound = true

				if v.(bool) {
					learnFromRedirectTrafficInt := &ves_io_schema_app_type.BusinessLogicMarkupSetting_Enable{}
					learnFromRedirectTrafficInt.Enable = &ves_io_schema.Empty{}
					businessLogicMarkupSetting.LearnFromRedirectTraffic = learnFromRedirectTrafficInt
				}

			}

		}

	}

	if v, ok := d.GetOk("features"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		features := make([]*ves_io_schema_app_type.Feature, len(sl))
		updateSpec.Features = features
		for i, set := range sl {
			features[i] = &ves_io_schema_app_type.Feature{}

			featuresMapStrToI := set.(map[string]interface{})

			if v, ok := featuresMapStrToI["type"]; ok && !isIntfNil(v) {

				features[i].Type = ves_io_schema_app_type.FeatureType(ves_io_schema_app_type.FeatureType_value[v.(string)])

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra AppType obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_app_type.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating AppType: %s", err)
	}

	return resourceVolterraAppTypeRead(d, meta)
}

func resourceVolterraAppTypeDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_app_type.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AppType %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AppType before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra AppType obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_app_type.ObjectType, namespace, name)
}
