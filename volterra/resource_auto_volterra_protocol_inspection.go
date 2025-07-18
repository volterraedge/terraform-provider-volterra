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
	ves_io_schema_protocol_inspection "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/protocol_inspection"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraProtocolInspection is implementation of Volterra's ProtocolInspection resources
func resourceVolterraProtocolInspection() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraProtocolInspectionCreate,
		Read:   resourceVolterraProtocolInspectionRead,
		Update: resourceVolterraProtocolInspectionUpdate,
		Delete: resourceVolterraProtocolInspectionDelete,

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

			"action": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"enable_disable_compliance_checks": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"disable_compliance_checks": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_compliance_checks": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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
				},
			},

			"enable_disable_signatures": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"disable_signature": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_signature": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraProtocolInspectionCreate creates ProtocolInspection resource
func resourceVolterraProtocolInspectionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_protocol_inspection.CreateSpecType{}
	createReq := &ves_io_schema_protocol_inspection.CreateRequest{
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

	//action
	if v, ok := d.GetOk("action"); ok && !isIntfNil(v) {

		createSpec.Action = ves_io_schema_protocol_inspection.Action(ves_io_schema_protocol_inspection.Action_value[v.(string)])

	}

	//enable_disable_compliance_checks
	if v, ok := d.GetOk("enable_disable_compliance_checks"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		enableDisableComplianceChecks := &ves_io_schema_protocol_inspection.EnableDisableComplianceChecks{}
		createSpec.EnableDisableComplianceChecks = enableDisableComplianceChecks
		for _, set := range sl {
			if set != nil {
				enableDisableComplianceChecksMapStrToI := set.(map[string]interface{})

				complianceCheckChoiceTypeFound := false

				if v, ok := enableDisableComplianceChecksMapStrToI["disable_compliance_checks"]; ok && !isIntfNil(v) && !complianceCheckChoiceTypeFound {

					complianceCheckChoiceTypeFound = true

					if v.(bool) {
						complianceCheckChoiceInt := &ves_io_schema_protocol_inspection.EnableDisableComplianceChecks_DisableComplianceChecks{}
						complianceCheckChoiceInt.DisableComplianceChecks = &ves_io_schema.Empty{}
						enableDisableComplianceChecks.ComplianceCheckChoice = complianceCheckChoiceInt
					}

				}

				if v, ok := enableDisableComplianceChecksMapStrToI["enable_compliance_checks"]; ok && !isIntfNil(v) && !complianceCheckChoiceTypeFound {

					complianceCheckChoiceTypeFound = true
					complianceCheckChoiceInt := &ves_io_schema_protocol_inspection.EnableDisableComplianceChecks_EnableComplianceChecks{}
					complianceCheckChoiceInt.EnableComplianceChecks = &ves_io_schema_views.ObjectRefType{}
					enableDisableComplianceChecks.ComplianceCheckChoice = complianceCheckChoiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								complianceCheckChoiceInt.EnableComplianceChecks.Name = v.(string)

							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								complianceCheckChoiceInt.EnableComplianceChecks.Namespace = v.(string)

							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								complianceCheckChoiceInt.EnableComplianceChecks.Tenant = v.(string)

							}

						}
					}

				}

			}
		}

	}

	//enable_disable_signatures
	if v, ok := d.GetOk("enable_disable_signatures"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		enableDisableSignatures := &ves_io_schema_protocol_inspection.EnableDisableSignatures{}
		createSpec.EnableDisableSignatures = enableDisableSignatures
		for _, set := range sl {
			if set != nil {
				enableDisableSignaturesMapStrToI := set.(map[string]interface{})

				signatureChoiceTypeFound := false

				if v, ok := enableDisableSignaturesMapStrToI["disable_signature"]; ok && !isIntfNil(v) && !signatureChoiceTypeFound {

					signatureChoiceTypeFound = true

					if v.(bool) {
						signatureChoiceInt := &ves_io_schema_protocol_inspection.EnableDisableSignatures_DisableSignature{}
						signatureChoiceInt.DisableSignature = &ves_io_schema.Empty{}
						enableDisableSignatures.SignatureChoice = signatureChoiceInt
					}

				}

				if v, ok := enableDisableSignaturesMapStrToI["enable_signature"]; ok && !isIntfNil(v) && !signatureChoiceTypeFound {

					signatureChoiceTypeFound = true

					if v.(bool) {
						signatureChoiceInt := &ves_io_schema_protocol_inspection.EnableDisableSignatures_EnableSignature{}
						signatureChoiceInt.EnableSignature = &ves_io_schema.Empty{}
						enableDisableSignatures.SignatureChoice = signatureChoiceInt
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Creating Volterra ProtocolInspection object with struct: %+v", createReq)

	createProtocolInspectionResp, err := client.CreateObject(context.Background(), ves_io_schema_protocol_inspection.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating ProtocolInspection: %s", err)
	}
	d.SetId(createProtocolInspectionResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraProtocolInspectionRead(d, meta)
}

func resourceVolterraProtocolInspectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_protocol_inspection.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ProtocolInspection %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ProtocolInspection %q: %s", d.Id(), err)
	}
	return setProtocolInspectionFields(client, d, resp)
}

func setProtocolInspectionFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraProtocolInspectionUpdate updates ProtocolInspection resource
func resourceVolterraProtocolInspectionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_protocol_inspection.ReplaceSpecType{}
	updateReq := &ves_io_schema_protocol_inspection.ReplaceRequest{
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

	if v, ok := d.GetOk("action"); ok && !isIntfNil(v) {

		updateSpec.Action = ves_io_schema_protocol_inspection.Action(ves_io_schema_protocol_inspection.Action_value[v.(string)])

	}

	if v, ok := d.GetOk("enable_disable_compliance_checks"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		enableDisableComplianceChecks := &ves_io_schema_protocol_inspection.EnableDisableComplianceChecks{}
		updateSpec.EnableDisableComplianceChecks = enableDisableComplianceChecks
		for _, set := range sl {
			if set != nil {
				enableDisableComplianceChecksMapStrToI := set.(map[string]interface{})

				complianceCheckChoiceTypeFound := false

				if v, ok := enableDisableComplianceChecksMapStrToI["disable_compliance_checks"]; ok && !isIntfNil(v) && !complianceCheckChoiceTypeFound {

					complianceCheckChoiceTypeFound = true

					if v.(bool) {
						complianceCheckChoiceInt := &ves_io_schema_protocol_inspection.EnableDisableComplianceChecks_DisableComplianceChecks{}
						complianceCheckChoiceInt.DisableComplianceChecks = &ves_io_schema.Empty{}
						enableDisableComplianceChecks.ComplianceCheckChoice = complianceCheckChoiceInt
					}

				}

				if v, ok := enableDisableComplianceChecksMapStrToI["enable_compliance_checks"]; ok && !isIntfNil(v) && !complianceCheckChoiceTypeFound {

					complianceCheckChoiceTypeFound = true
					complianceCheckChoiceInt := &ves_io_schema_protocol_inspection.EnableDisableComplianceChecks_EnableComplianceChecks{}
					complianceCheckChoiceInt.EnableComplianceChecks = &ves_io_schema_views.ObjectRefType{}
					enableDisableComplianceChecks.ComplianceCheckChoice = complianceCheckChoiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								complianceCheckChoiceInt.EnableComplianceChecks.Name = v.(string)

							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								complianceCheckChoiceInt.EnableComplianceChecks.Namespace = v.(string)

							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								complianceCheckChoiceInt.EnableComplianceChecks.Tenant = v.(string)

							}

						}
					}

				}

			}
		}

	}

	if v, ok := d.GetOk("enable_disable_signatures"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		enableDisableSignatures := &ves_io_schema_protocol_inspection.EnableDisableSignatures{}
		updateSpec.EnableDisableSignatures = enableDisableSignatures
		for _, set := range sl {
			if set != nil {
				enableDisableSignaturesMapStrToI := set.(map[string]interface{})

				signatureChoiceTypeFound := false

				if v, ok := enableDisableSignaturesMapStrToI["disable_signature"]; ok && !isIntfNil(v) && !signatureChoiceTypeFound {

					signatureChoiceTypeFound = true

					if v.(bool) {
						signatureChoiceInt := &ves_io_schema_protocol_inspection.EnableDisableSignatures_DisableSignature{}
						signatureChoiceInt.DisableSignature = &ves_io_schema.Empty{}
						enableDisableSignatures.SignatureChoice = signatureChoiceInt
					}

				}

				if v, ok := enableDisableSignaturesMapStrToI["enable_signature"]; ok && !isIntfNil(v) && !signatureChoiceTypeFound {

					signatureChoiceTypeFound = true

					if v.(bool) {
						signatureChoiceInt := &ves_io_schema_protocol_inspection.EnableDisableSignatures_EnableSignature{}
						signatureChoiceInt.EnableSignature = &ves_io_schema.Empty{}
						enableDisableSignatures.SignatureChoice = signatureChoiceInt
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Updating Volterra ProtocolInspection obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_protocol_inspection.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating ProtocolInspection: %s", err)
	}

	return resourceVolterraProtocolInspectionRead(d, meta)
}

func resourceVolterraProtocolInspectionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_protocol_inspection.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ProtocolInspection %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ProtocolInspection before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra ProtocolInspection obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_protocol_inspection.ObjectType, namespace, name, opts...)
}
