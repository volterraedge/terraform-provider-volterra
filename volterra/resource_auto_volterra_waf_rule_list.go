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
	ves_io_schema_waf_rule_list "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf_rule_list"
)

// resourceVolterraWafRuleList is implementation of Volterra's WafRuleList resources
func resourceVolterraWafRuleList() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraWafRuleListCreate,
		Read:   resourceVolterraWafRuleListRead,
		Update: resourceVolterraWafRuleListUpdate,
		Delete: resourceVolterraWafRuleListDelete,

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

			"rule_ids": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// resourceVolterraWafRuleListCreate creates WafRuleList resource
func resourceVolterraWafRuleListCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_waf_rule_list.CreateSpecType{}
	createReq := &ves_io_schema_waf_rule_list.CreateRequest{
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

	if v, ok := d.GetOk("rule_ids"); ok && !isIntfNil(v) {

		rule_idsList := []ves_io_schema_waf_rule_list.WafRuleID{}
		for _, j := range v.([]interface{}) {
			rule_idsList = append(rule_idsList, ves_io_schema_waf_rule_list.WafRuleID(ves_io_schema_waf_rule_list.WafRuleID_value[j.(string)]))
		}
		createSpec.RuleIds = rule_idsList

	}

	log.Printf("[DEBUG] Creating Volterra WafRuleList object with struct: %+v", createReq)

	createWafRuleListResp, err := client.CreateObject(context.Background(), ves_io_schema_waf_rule_list.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating WafRuleList: %s", err)
	}
	d.SetId(createWafRuleListResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraWafRuleListRead(d, meta)
}

func resourceVolterraWafRuleListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_waf_rule_list.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] WafRuleList %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra WafRuleList %q: %s", d.Id(), err)
	}
	return setWafRuleListFields(client, d, resp)
}

func setWafRuleListFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraWafRuleListUpdate updates WafRuleList resource
func resourceVolterraWafRuleListUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_waf_rule_list.ReplaceSpecType{}
	updateReq := &ves_io_schema_waf_rule_list.ReplaceRequest{
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

	if v, ok := d.GetOk("rule_ids"); ok && !isIntfNil(v) {

		rule_idsList := []ves_io_schema_waf_rule_list.WafRuleID{}
		for _, j := range v.([]interface{}) {
			rule_idsList = append(rule_idsList, ves_io_schema_waf_rule_list.WafRuleID(ves_io_schema_waf_rule_list.WafRuleID_value[j.(string)]))
		}
		updateSpec.RuleIds = rule_idsList

	}

	log.Printf("[DEBUG] Updating Volterra WafRuleList obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_waf_rule_list.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating WafRuleList: %s", err)
	}

	return resourceVolterraWafRuleListRead(d, meta)
}

func resourceVolterraWafRuleListDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_waf_rule_list.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] WafRuleList %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra WafRuleList before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra WafRuleList obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_waf_rule_list.ObjectType, namespace, name)
}
