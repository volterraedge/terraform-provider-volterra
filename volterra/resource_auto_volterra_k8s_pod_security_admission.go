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
	ves_io_schema_k8s_pod_security_admission "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/k8s_pod_security_admission"
)

// resourceVolterraK8SPodSecurityAdmission is implementation of Volterra's K8SPodSecurityAdmission resources
func resourceVolterraK8SPodSecurityAdmission() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraK8SPodSecurityAdmissionCreate,
		Read:   resourceVolterraK8SPodSecurityAdmissionRead,
		Update: resourceVolterraK8SPodSecurityAdmissionUpdate,
		Delete: resourceVolterraK8SPodSecurityAdmissionDelete,

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

			"pod_security_admission_specs": {

				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"audit": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enforce": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"warn": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"baseline": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"privileged": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"restricted": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraK8SPodSecurityAdmissionCreate creates K8SPodSecurityAdmission resource
func resourceVolterraK8SPodSecurityAdmissionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_k8s_pod_security_admission.CreateSpecType{}
	createReq := &ves_io_schema_k8s_pod_security_admission.CreateRequest{
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

	//pod_security_admission_specs
	if v, ok := d.GetOk("pod_security_admission_specs"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		podSecurityAdmissionSpecs := make([]*ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec, len(sl))
		createSpec.PodSecurityAdmissionSpecs = podSecurityAdmissionSpecs
		for i, set := range sl {
			if set != nil {
				podSecurityAdmissionSpecs[i] = &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec{}
				podSecurityAdmissionSpecsMapStrToI := set.(map[string]interface{})

				admissionModeChoiceTypeFound := false

				if v, ok := podSecurityAdmissionSpecsMapStrToI["audit"]; ok && !isIntfNil(v) && !admissionModeChoiceTypeFound {

					admissionModeChoiceTypeFound = true

					if v.(bool) {
						admissionModeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Audit{}
						admissionModeChoiceInt.Audit = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].AdmissionModeChoice = admissionModeChoiceInt
					}

				}

				if v, ok := podSecurityAdmissionSpecsMapStrToI["enforce"]; ok && !isIntfNil(v) && !admissionModeChoiceTypeFound {

					admissionModeChoiceTypeFound = true

					if v.(bool) {
						admissionModeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Enforce{}
						admissionModeChoiceInt.Enforce = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].AdmissionModeChoice = admissionModeChoiceInt
					}

				}

				if v, ok := podSecurityAdmissionSpecsMapStrToI["warn"]; ok && !isIntfNil(v) && !admissionModeChoiceTypeFound {

					admissionModeChoiceTypeFound = true

					if v.(bool) {
						admissionModeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Warn{}
						admissionModeChoiceInt.Warn = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].AdmissionModeChoice = admissionModeChoiceInt
					}

				}

				policyTypeChoiceTypeFound := false

				if v, ok := podSecurityAdmissionSpecsMapStrToI["baseline"]; ok && !isIntfNil(v) && !policyTypeChoiceTypeFound {

					policyTypeChoiceTypeFound = true

					if v.(bool) {
						policyTypeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Baseline{}
						policyTypeChoiceInt.Baseline = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].PolicyTypeChoice = policyTypeChoiceInt
					}

				}

				if v, ok := podSecurityAdmissionSpecsMapStrToI["privileged"]; ok && !isIntfNil(v) && !policyTypeChoiceTypeFound {

					policyTypeChoiceTypeFound = true

					if v.(bool) {
						policyTypeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Privileged{}
						policyTypeChoiceInt.Privileged = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].PolicyTypeChoice = policyTypeChoiceInt
					}

				}

				if v, ok := podSecurityAdmissionSpecsMapStrToI["restricted"]; ok && !isIntfNil(v) && !policyTypeChoiceTypeFound {

					policyTypeChoiceTypeFound = true

					if v.(bool) {
						policyTypeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Restricted{}
						policyTypeChoiceInt.Restricted = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].PolicyTypeChoice = policyTypeChoiceInt
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Creating Volterra K8SPodSecurityAdmission object with struct: %+v", createReq)

	createK8SPodSecurityAdmissionResp, err := client.CreateObject(context.Background(), ves_io_schema_k8s_pod_security_admission.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating K8SPodSecurityAdmission: %s", err)
	}
	d.SetId(createK8SPodSecurityAdmissionResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraK8SPodSecurityAdmissionRead(d, meta)
}

func resourceVolterraK8SPodSecurityAdmissionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_k8s_pod_security_admission.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] K8SPodSecurityAdmission %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra K8SPodSecurityAdmission %q: %s", d.Id(), err)
	}
	return setK8SPodSecurityAdmissionFields(client, d, resp)
}

func setK8SPodSecurityAdmissionFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraK8SPodSecurityAdmissionUpdate updates K8SPodSecurityAdmission resource
func resourceVolterraK8SPodSecurityAdmissionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_k8s_pod_security_admission.ReplaceSpecType{}
	updateReq := &ves_io_schema_k8s_pod_security_admission.ReplaceRequest{
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

	if v, ok := d.GetOk("pod_security_admission_specs"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		podSecurityAdmissionSpecs := make([]*ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec, len(sl))
		updateSpec.PodSecurityAdmissionSpecs = podSecurityAdmissionSpecs
		for i, set := range sl {
			if set != nil {
				podSecurityAdmissionSpecs[i] = &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec{}
				podSecurityAdmissionSpecsMapStrToI := set.(map[string]interface{})

				admissionModeChoiceTypeFound := false

				if v, ok := podSecurityAdmissionSpecsMapStrToI["audit"]; ok && !isIntfNil(v) && !admissionModeChoiceTypeFound {

					admissionModeChoiceTypeFound = true

					if v.(bool) {
						admissionModeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Audit{}
						admissionModeChoiceInt.Audit = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].AdmissionModeChoice = admissionModeChoiceInt
					}

				}

				if v, ok := podSecurityAdmissionSpecsMapStrToI["enforce"]; ok && !isIntfNil(v) && !admissionModeChoiceTypeFound {

					admissionModeChoiceTypeFound = true

					if v.(bool) {
						admissionModeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Enforce{}
						admissionModeChoiceInt.Enforce = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].AdmissionModeChoice = admissionModeChoiceInt
					}

				}

				if v, ok := podSecurityAdmissionSpecsMapStrToI["warn"]; ok && !isIntfNil(v) && !admissionModeChoiceTypeFound {

					admissionModeChoiceTypeFound = true

					if v.(bool) {
						admissionModeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Warn{}
						admissionModeChoiceInt.Warn = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].AdmissionModeChoice = admissionModeChoiceInt
					}

				}

				policyTypeChoiceTypeFound := false

				if v, ok := podSecurityAdmissionSpecsMapStrToI["baseline"]; ok && !isIntfNil(v) && !policyTypeChoiceTypeFound {

					policyTypeChoiceTypeFound = true

					if v.(bool) {
						policyTypeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Baseline{}
						policyTypeChoiceInt.Baseline = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].PolicyTypeChoice = policyTypeChoiceInt
					}

				}

				if v, ok := podSecurityAdmissionSpecsMapStrToI["privileged"]; ok && !isIntfNil(v) && !policyTypeChoiceTypeFound {

					policyTypeChoiceTypeFound = true

					if v.(bool) {
						policyTypeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Privileged{}
						policyTypeChoiceInt.Privileged = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].PolicyTypeChoice = policyTypeChoiceInt
					}

				}

				if v, ok := podSecurityAdmissionSpecsMapStrToI["restricted"]; ok && !isIntfNil(v) && !policyTypeChoiceTypeFound {

					policyTypeChoiceTypeFound = true

					if v.(bool) {
						policyTypeChoiceInt := &ves_io_schema_k8s_pod_security_admission.PodSecurityAdmissionSpec_Restricted{}
						policyTypeChoiceInt.Restricted = &ves_io_schema.Empty{}
						podSecurityAdmissionSpecs[i].PolicyTypeChoice = policyTypeChoiceInt
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Updating Volterra K8SPodSecurityAdmission obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_k8s_pod_security_admission.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating K8SPodSecurityAdmission: %s", err)
	}

	return resourceVolterraK8SPodSecurityAdmissionRead(d, meta)
}

func resourceVolterraK8SPodSecurityAdmissionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_k8s_pod_security_admission.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] K8SPodSecurityAdmission %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra K8SPodSecurityAdmission before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra K8SPodSecurityAdmission obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_k8s_pod_security_admission.ObjectType, namespace, name, opts...)
}
