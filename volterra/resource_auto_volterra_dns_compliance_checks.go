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
	ves_io_schema_dns_compliance_checks "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dns_compliance_checks"
)

// resourceVolterraDnsComplianceChecks is implementation of Volterra's DnsComplianceChecks resources
func resourceVolterraDnsComplianceChecks() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraDnsComplianceChecksCreate,
		Read:   resourceVolterraDnsComplianceChecksRead,
		Update: resourceVolterraDnsComplianceChecksUpdate,
		Delete: resourceVolterraDnsComplianceChecksDelete,

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

			"disallowed_query_type_list": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"disallowed_resource_record_type_list": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"domain_denylist": {

				Type: schema.TypeList,

				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// resourceVolterraDnsComplianceChecksCreate creates DnsComplianceChecks resource
func resourceVolterraDnsComplianceChecksCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_dns_compliance_checks.CreateSpecType{}
	createReq := &ves_io_schema_dns_compliance_checks.CreateRequest{
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

	//disallowed_query_type_list
	if v, ok := d.GetOk("disallowed_query_type_list"); ok && !isIntfNil(v) {

		disallowed_query_type_listList := []ves_io_schema_dns_compliance_checks.DisallowedQueryType{}
		for _, j := range v.([]interface{}) {
			if j == nil {
				return fmt.Errorf("please provide valid non-empty enum value of field disallowed_query_type_list")
			}
			disallowed_query_type_listList = append(disallowed_query_type_listList, ves_io_schema_dns_compliance_checks.DisallowedQueryType(ves_io_schema_dns_compliance_checks.DisallowedQueryType_value[j.(string)]))
		}
		createSpec.DisallowedQueryTypeList = disallowed_query_type_listList

	}

	//disallowed_resource_record_type_list
	if v, ok := d.GetOk("disallowed_resource_record_type_list"); ok && !isIntfNil(v) {

		disallowed_resource_record_type_listList := []ves_io_schema_dns_compliance_checks.DisallowedResourceRecordType{}
		for _, j := range v.([]interface{}) {
			if j == nil {
				return fmt.Errorf("please provide valid non-empty enum value of field disallowed_resource_record_type_list")
			}
			disallowed_resource_record_type_listList = append(disallowed_resource_record_type_listList, ves_io_schema_dns_compliance_checks.DisallowedResourceRecordType(ves_io_schema_dns_compliance_checks.DisallowedResourceRecordType_value[j.(string)]))
		}
		createSpec.DisallowedResourceRecordTypeList = disallowed_resource_record_type_listList

	}

	//domain_denylist
	if v, ok := d.GetOk("domain_denylist"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			if v == nil {
				return fmt.Errorf("please provide valid non-empty string value of field domain_denylist")
			}
			if str, ok := v.(string); ok {
				ls[i] = str
			}
		}
		createSpec.DomainDenylist = ls

	}

	log.Printf("[DEBUG] Creating Volterra DnsComplianceChecks object with struct: %+v", createReq)

	createDnsComplianceChecksResp, err := client.CreateObject(context.Background(), ves_io_schema_dns_compliance_checks.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating DnsComplianceChecks: %s", err)
	}
	d.SetId(createDnsComplianceChecksResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraDnsComplianceChecksRead(d, meta)
}

func resourceVolterraDnsComplianceChecksRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_dns_compliance_checks.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] DnsComplianceChecks %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra DnsComplianceChecks %q: %s", d.Id(), err)
	}
	return setDnsComplianceChecksFields(client, d, resp)
}

func setDnsComplianceChecksFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraDnsComplianceChecksUpdate updates DnsComplianceChecks resource
func resourceVolterraDnsComplianceChecksUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_dns_compliance_checks.ReplaceSpecType{}
	updateReq := &ves_io_schema_dns_compliance_checks.ReplaceRequest{
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

	if v, ok := d.GetOk("disallowed_query_type_list"); ok && !isIntfNil(v) {

		disallowed_query_type_listList := []ves_io_schema_dns_compliance_checks.DisallowedQueryType{}
		for _, j := range v.([]interface{}) {
			if j == nil {
				return fmt.Errorf("please provide valid non-empty enum value of field disallowed_query_type_list")
			}
			disallowed_query_type_listList = append(disallowed_query_type_listList, ves_io_schema_dns_compliance_checks.DisallowedQueryType(ves_io_schema_dns_compliance_checks.DisallowedQueryType_value[j.(string)]))
		}
		updateSpec.DisallowedQueryTypeList = disallowed_query_type_listList

	}

	if v, ok := d.GetOk("disallowed_resource_record_type_list"); ok && !isIntfNil(v) {

		disallowed_resource_record_type_listList := []ves_io_schema_dns_compliance_checks.DisallowedResourceRecordType{}
		for _, j := range v.([]interface{}) {
			if j == nil {
				return fmt.Errorf("please provide valid non-empty enum value of field disallowed_resource_record_type_list")
			}
			disallowed_resource_record_type_listList = append(disallowed_resource_record_type_listList, ves_io_schema_dns_compliance_checks.DisallowedResourceRecordType(ves_io_schema_dns_compliance_checks.DisallowedResourceRecordType_value[j.(string)]))
		}
		updateSpec.DisallowedResourceRecordTypeList = disallowed_resource_record_type_listList

	}

	if v, ok := d.GetOk("domain_denylist"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			if v == nil {
				return fmt.Errorf("please provide valid non-empty string value of field domain_denylist")
			}
			if str, ok := v.(string); ok {
				ls[i] = str
			}
		}
		updateSpec.DomainDenylist = ls

	}

	log.Printf("[DEBUG] Updating Volterra DnsComplianceChecks obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_dns_compliance_checks.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating DnsComplianceChecks: %s", err)
	}

	return resourceVolterraDnsComplianceChecksRead(d, meta)
}

func resourceVolterraDnsComplianceChecksDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_dns_compliance_checks.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] DnsComplianceChecks %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra DnsComplianceChecks before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra DnsComplianceChecks obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_dns_compliance_checks.ObjectType, namespace, name, opts...)
}
