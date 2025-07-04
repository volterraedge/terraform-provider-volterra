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
	ves_io_schema_k8s_cluster_role_binding "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/k8s_cluster_role_binding"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraK8SClusterRoleBinding is implementation of Volterra's K8SClusterRoleBinding resources
func resourceVolterraK8SClusterRoleBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraK8SClusterRoleBindingCreate,
		Read:   resourceVolterraK8SClusterRoleBindingRead,
		Update: resourceVolterraK8SClusterRoleBindingUpdate,
		Delete: resourceVolterraK8SClusterRoleBindingDelete,

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

			"k8s_cluster_role": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
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

			"subjects": {

				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"group": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"service_account": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"namespace": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"user": {

							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraK8SClusterRoleBindingCreate creates K8SClusterRoleBinding resource
func resourceVolterraK8SClusterRoleBindingCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_k8s_cluster_role_binding.CreateSpecType{}
	createReq := &ves_io_schema_k8s_cluster_role_binding.CreateRequest{
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

	//k8s_cluster_role
	if v, ok := d.GetOk("k8s_cluster_role"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		k8SClusterRoleInt := &ves_io_schema_views.ObjectRefType{}
		createSpec.K8SClusterRole = k8SClusterRoleInt

		for _, set := range sl {
			if set != nil {
				kcrMapToStrVal := set.(map[string]interface{})
				if val, ok := kcrMapToStrVal["name"]; ok && !isIntfNil(v) {
					k8SClusterRoleInt.Name = val.(string)
				}
				if val, ok := kcrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
					k8SClusterRoleInt.Namespace = val.(string)
				}

				if val, ok := kcrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
					k8SClusterRoleInt.Tenant = val.(string)
				}
			}
		}

	}

	//subjects
	if v, ok := d.GetOk("subjects"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		subjects := make([]*ves_io_schema_k8s_cluster_role_binding.SubjectType, len(sl))
		createSpec.Subjects = subjects
		for i, set := range sl {
			if set != nil {
				subjects[i] = &ves_io_schema_k8s_cluster_role_binding.SubjectType{}
				subjectsMapStrToI := set.(map[string]interface{})

				subjectChoiceTypeFound := false

				if v, ok := subjectsMapStrToI["group"]; ok && !isIntfNil(v) && !subjectChoiceTypeFound {

					subjectChoiceTypeFound = true
					subjectChoiceInt := &ves_io_schema_k8s_cluster_role_binding.SubjectType_Group{}

					subjects[i].SubjectChoice = subjectChoiceInt

					subjectChoiceInt.Group = v.(string)

				}

				if v, ok := subjectsMapStrToI["service_account"]; ok && !isIntfNil(v) && !subjectChoiceTypeFound {

					subjectChoiceTypeFound = true
					subjectChoiceInt := &ves_io_schema_k8s_cluster_role_binding.SubjectType_ServiceAccount{}
					subjectChoiceInt.ServiceAccount = &ves_io_schema_k8s_cluster_role_binding.ServiceAccountType{}
					subjects[i].SubjectChoice = subjectChoiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								subjectChoiceInt.ServiceAccount.Name = v.(string)

							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								subjectChoiceInt.ServiceAccount.Namespace = v.(string)

							}

						}
					}

				}

				if v, ok := subjectsMapStrToI["user"]; ok && !isIntfNil(v) && !subjectChoiceTypeFound {

					subjectChoiceTypeFound = true
					subjectChoiceInt := &ves_io_schema_k8s_cluster_role_binding.SubjectType_User{}

					subjects[i].SubjectChoice = subjectChoiceInt

					subjectChoiceInt.User = v.(string)

				}

			}
		}

	}

	log.Printf("[DEBUG] Creating Volterra K8SClusterRoleBinding object with struct: %+v", createReq)

	createK8SClusterRoleBindingResp, err := client.CreateObject(context.Background(), ves_io_schema_k8s_cluster_role_binding.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating K8SClusterRoleBinding: %s", err)
	}
	d.SetId(createK8SClusterRoleBindingResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraK8SClusterRoleBindingRead(d, meta)
}

func resourceVolterraK8SClusterRoleBindingRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_k8s_cluster_role_binding.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] K8SClusterRoleBinding %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra K8SClusterRoleBinding %q: %s", d.Id(), err)
	}
	return setK8SClusterRoleBindingFields(client, d, resp)
}

func setK8SClusterRoleBindingFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraK8SClusterRoleBindingUpdate updates K8SClusterRoleBinding resource
func resourceVolterraK8SClusterRoleBindingUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_k8s_cluster_role_binding.ReplaceSpecType{}
	updateReq := &ves_io_schema_k8s_cluster_role_binding.ReplaceRequest{
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

	if v, ok := d.GetOk("k8s_cluster_role"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		k8SClusterRoleInt := &ves_io_schema_views.ObjectRefType{}
		updateSpec.K8SClusterRole = k8SClusterRoleInt

		for _, set := range sl {
			if set != nil {
				kcrMapToStrVal := set.(map[string]interface{})
				if val, ok := kcrMapToStrVal["name"]; ok && !isIntfNil(v) {
					k8SClusterRoleInt.Name = val.(string)
				}
				if val, ok := kcrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
					k8SClusterRoleInt.Namespace = val.(string)
				}

				if val, ok := kcrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
					k8SClusterRoleInt.Tenant = val.(string)
				}
			}
		}

	}

	if v, ok := d.GetOk("subjects"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		subjects := make([]*ves_io_schema_k8s_cluster_role_binding.SubjectType, len(sl))
		updateSpec.Subjects = subjects
		for i, set := range sl {
			if set != nil {
				subjects[i] = &ves_io_schema_k8s_cluster_role_binding.SubjectType{}
				subjectsMapStrToI := set.(map[string]interface{})

				subjectChoiceTypeFound := false

				if v, ok := subjectsMapStrToI["group"]; ok && !isIntfNil(v) && !subjectChoiceTypeFound {

					subjectChoiceTypeFound = true
					subjectChoiceInt := &ves_io_schema_k8s_cluster_role_binding.SubjectType_Group{}

					subjects[i].SubjectChoice = subjectChoiceInt

					subjectChoiceInt.Group = v.(string)

				}

				if v, ok := subjectsMapStrToI["service_account"]; ok && !isIntfNil(v) && !subjectChoiceTypeFound {

					subjectChoiceTypeFound = true
					subjectChoiceInt := &ves_io_schema_k8s_cluster_role_binding.SubjectType_ServiceAccount{}
					subjectChoiceInt.ServiceAccount = &ves_io_schema_k8s_cluster_role_binding.ServiceAccountType{}
					subjects[i].SubjectChoice = subjectChoiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								subjectChoiceInt.ServiceAccount.Name = v.(string)

							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								subjectChoiceInt.ServiceAccount.Namespace = v.(string)

							}

						}
					}

				}

				if v, ok := subjectsMapStrToI["user"]; ok && !isIntfNil(v) && !subjectChoiceTypeFound {

					subjectChoiceTypeFound = true
					subjectChoiceInt := &ves_io_schema_k8s_cluster_role_binding.SubjectType_User{}

					subjects[i].SubjectChoice = subjectChoiceInt

					subjectChoiceInt.User = v.(string)

				}

			}
		}

	}

	log.Printf("[DEBUG] Updating Volterra K8SClusterRoleBinding obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_k8s_cluster_role_binding.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating K8SClusterRoleBinding: %s", err)
	}

	return resourceVolterraK8SClusterRoleBindingRead(d, meta)
}

func resourceVolterraK8SClusterRoleBindingDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_k8s_cluster_role_binding.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] K8SClusterRoleBinding %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra K8SClusterRoleBinding before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra K8SClusterRoleBinding obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_k8s_cluster_role_binding.ObjectType, namespace, name, opts...)
}
