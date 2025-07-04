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
	ves_io_schema_uztna_uztna_policy_template "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/uztna/uztna_policy_template"
)

// resourceVolterraUztnaPolicyTemplate is implementation of Volterra's UztnaPolicyTemplate resources
func resourceVolterraUztnaPolicyTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraUztnaPolicyTemplateCreate,
		Read:   resourceVolterraUztnaPolicyTemplateRead,
		Update: resourceVolterraUztnaPolicyTemplateUpdate,
		Delete: resourceVolterraUztnaPolicyTemplateDelete,

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

			"on_start_flow": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"simple": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"flows": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"geolocation_match": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"geomatch": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"kind": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},

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

												"saml_federation": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"saml": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"kind": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},

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
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraUztnaPolicyTemplateCreate creates UztnaPolicyTemplate resource
func resourceVolterraUztnaPolicyTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_uztna_uztna_policy_template.CreateSpecType{}
	createReq := &ves_io_schema_uztna_uztna_policy_template.CreateRequest{
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

	//on_start_flow
	if v, ok := d.GetOk("on_start_flow"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		onStartFlow := &ves_io_schema_uztna_uztna_policy_template.TemplateType{}
		createSpec.OnStartFlow = onStartFlow
		for _, set := range sl {
			if set != nil {
				onStartFlowMapStrToI := set.(map[string]interface{})

				templateTypeTypeFound := false

				if v, ok := onStartFlowMapStrToI["simple"]; ok && !isIntfNil(v) && !templateTypeTypeFound {

					templateTypeTypeFound = true
					templateTypeInt := &ves_io_schema_uztna_uztna_policy_template.TemplateType_Simple{}
					templateTypeInt.Simple = &ves_io_schema_uztna_uztna_policy_template.SimpleTemplate{}
					onStartFlow.TemplateType = templateTypeInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["flows"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								flows := make([]*ves_io_schema_uztna_uztna_policy_template.Flows, len(sl))
								templateTypeInt.Simple.Flows = flows
								for i, set := range sl {
									if set != nil {
										flows[i] = &ves_io_schema_uztna_uztna_policy_template.Flows{}
										flowsMapStrToI := set.(map[string]interface{})

										flowTypeTypeFound := false

										if v, ok := flowsMapStrToI["geolocation_match"]; ok && !isIntfNil(v) && !flowTypeTypeFound {

											flowTypeTypeFound = true
											flowTypeInt := &ves_io_schema_uztna_uztna_policy_template.Flows_GeolocationMatch{}
											flowTypeInt.GeolocationMatch = &ves_io_schema_uztna_uztna_policy_template.GeoLocationMatch{}
											flows[i].FlowType = flowTypeInt

											sl := v.([]interface{})
											for _, set := range sl {
												if set != nil {
													cs := set.(map[string]interface{})

													if v, ok := cs["geomatch"]; ok && !isIntfNil(v) {

														sl := v.([]interface{})
														geomatchInt := make([]*ves_io_schema.ObjectRefType, len(sl))
														flowTypeInt.GeolocationMatch.Geomatch = geomatchInt
														for i, ps := range sl {

															gMapToStrVal := ps.(map[string]interface{})
															geomatchInt[i] = &ves_io_schema.ObjectRefType{}

															geomatchInt[i].Kind = "uztna_flow"

															if v, ok := gMapToStrVal["name"]; ok && !isIntfNil(v) {
																geomatchInt[i].Name = v.(string)
															}

															if v, ok := gMapToStrVal["namespace"]; ok && !isIntfNil(v) {
																geomatchInt[i].Namespace = v.(string)
															}

															if v, ok := gMapToStrVal["tenant"]; ok && !isIntfNil(v) {
																geomatchInt[i].Tenant = v.(string)
															}

															if v, ok := gMapToStrVal["uid"]; ok && !isIntfNil(v) {
																geomatchInt[i].Uid = v.(string)
															}

														}

													}

												}
											}

										}

										if v, ok := flowsMapStrToI["saml_federation"]; ok && !isIntfNil(v) && !flowTypeTypeFound {

											flowTypeTypeFound = true
											flowTypeInt := &ves_io_schema_uztna_uztna_policy_template.Flows_SamlFederation{}
											flowTypeInt.SamlFederation = &ves_io_schema_uztna_uztna_policy_template.SAMLFederation{}
											flows[i].FlowType = flowTypeInt

											sl := v.([]interface{})
											for _, set := range sl {
												if set != nil {
													cs := set.(map[string]interface{})

													if v, ok := cs["saml"]; ok && !isIntfNil(v) {

														sl := v.([]interface{})
														samlInt := make([]*ves_io_schema.ObjectRefType, len(sl))
														flowTypeInt.SamlFederation.Saml = samlInt
														for i, ps := range sl {

															sMapToStrVal := ps.(map[string]interface{})
															samlInt[i] = &ves_io_schema.ObjectRefType{}

															samlInt[i].Kind = "uztna_flow"

															if v, ok := sMapToStrVal["name"]; ok && !isIntfNil(v) {
																samlInt[i].Name = v.(string)
															}

															if v, ok := sMapToStrVal["namespace"]; ok && !isIntfNil(v) {
																samlInt[i].Namespace = v.(string)
															}

															if v, ok := sMapToStrVal["tenant"]; ok && !isIntfNil(v) {
																samlInt[i].Tenant = v.(string)
															}

															if v, ok := sMapToStrVal["uid"]; ok && !isIntfNil(v) {
																samlInt[i].Uid = v.(string)
															}

														}

													}

												}
											}

										}

									}
								}

							}

						}
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Creating Volterra UztnaPolicyTemplate object with struct: %+v", createReq)

	createUztnaPolicyTemplateResp, err := client.CreateObject(context.Background(), ves_io_schema_uztna_uztna_policy_template.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating UztnaPolicyTemplate: %s", err)
	}
	d.SetId(createUztnaPolicyTemplateResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraUztnaPolicyTemplateRead(d, meta)
}

func resourceVolterraUztnaPolicyTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_uztna_uztna_policy_template.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaPolicyTemplate %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaPolicyTemplate %q: %s", d.Id(), err)
	}
	return setUztnaPolicyTemplateFields(client, d, resp)
}

func setUztnaPolicyTemplateFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraUztnaPolicyTemplateUpdate updates UztnaPolicyTemplate resource
func resourceVolterraUztnaPolicyTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_uztna_uztna_policy_template.ReplaceSpecType{}
	updateReq := &ves_io_schema_uztna_uztna_policy_template.ReplaceRequest{
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

	if v, ok := d.GetOk("on_start_flow"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		onStartFlow := &ves_io_schema_uztna_uztna_policy_template.TemplateType{}
		updateSpec.OnStartFlow = onStartFlow
		for _, set := range sl {
			if set != nil {
				onStartFlowMapStrToI := set.(map[string]interface{})

				templateTypeTypeFound := false

				if v, ok := onStartFlowMapStrToI["simple"]; ok && !isIntfNil(v) && !templateTypeTypeFound {

					templateTypeTypeFound = true
					templateTypeInt := &ves_io_schema_uztna_uztna_policy_template.TemplateType_Simple{}
					templateTypeInt.Simple = &ves_io_schema_uztna_uztna_policy_template.SimpleTemplate{}
					onStartFlow.TemplateType = templateTypeInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["flows"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								flows := make([]*ves_io_schema_uztna_uztna_policy_template.Flows, len(sl))
								templateTypeInt.Simple.Flows = flows
								for i, set := range sl {
									if set != nil {
										flows[i] = &ves_io_schema_uztna_uztna_policy_template.Flows{}
										flowsMapStrToI := set.(map[string]interface{})

										flowTypeTypeFound := false

										if v, ok := flowsMapStrToI["geolocation_match"]; ok && !isIntfNil(v) && !flowTypeTypeFound {

											flowTypeTypeFound = true
											flowTypeInt := &ves_io_schema_uztna_uztna_policy_template.Flows_GeolocationMatch{}
											flowTypeInt.GeolocationMatch = &ves_io_schema_uztna_uztna_policy_template.GeoLocationMatch{}
											flows[i].FlowType = flowTypeInt

											sl := v.([]interface{})
											for _, set := range sl {
												if set != nil {
													cs := set.(map[string]interface{})

													if v, ok := cs["geomatch"]; ok && !isIntfNil(v) {

														sl := v.([]interface{})
														geomatchInt := make([]*ves_io_schema.ObjectRefType, len(sl))
														flowTypeInt.GeolocationMatch.Geomatch = geomatchInt
														for i, ps := range sl {

															gMapToStrVal := ps.(map[string]interface{})
															geomatchInt[i] = &ves_io_schema.ObjectRefType{}

															geomatchInt[i].Kind = "uztna_flow"

															if v, ok := gMapToStrVal["name"]; ok && !isIntfNil(v) {
																geomatchInt[i].Name = v.(string)
															}

															if v, ok := gMapToStrVal["namespace"]; ok && !isIntfNil(v) {
																geomatchInt[i].Namespace = v.(string)
															}

															if v, ok := gMapToStrVal["tenant"]; ok && !isIntfNil(v) {
																geomatchInt[i].Tenant = v.(string)
															}

															if v, ok := gMapToStrVal["uid"]; ok && !isIntfNil(v) {
																geomatchInt[i].Uid = v.(string)
															}

														}

													}

												}
											}

										}

										if v, ok := flowsMapStrToI["saml_federation"]; ok && !isIntfNil(v) && !flowTypeTypeFound {

											flowTypeTypeFound = true
											flowTypeInt := &ves_io_schema_uztna_uztna_policy_template.Flows_SamlFederation{}
											flowTypeInt.SamlFederation = &ves_io_schema_uztna_uztna_policy_template.SAMLFederation{}
											flows[i].FlowType = flowTypeInt

											sl := v.([]interface{})
											for _, set := range sl {
												if set != nil {
													cs := set.(map[string]interface{})

													if v, ok := cs["saml"]; ok && !isIntfNil(v) {

														sl := v.([]interface{})
														samlInt := make([]*ves_io_schema.ObjectRefType, len(sl))
														flowTypeInt.SamlFederation.Saml = samlInt
														for i, ps := range sl {

															sMapToStrVal := ps.(map[string]interface{})
															samlInt[i] = &ves_io_schema.ObjectRefType{}

															samlInt[i].Kind = "uztna_flow"

															if v, ok := sMapToStrVal["name"]; ok && !isIntfNil(v) {
																samlInt[i].Name = v.(string)
															}

															if v, ok := sMapToStrVal["namespace"]; ok && !isIntfNil(v) {
																samlInt[i].Namespace = v.(string)
															}

															if v, ok := sMapToStrVal["tenant"]; ok && !isIntfNil(v) {
																samlInt[i].Tenant = v.(string)
															}

															if v, ok := sMapToStrVal["uid"]; ok && !isIntfNil(v) {
																samlInt[i].Uid = v.(string)
															}

														}

													}

												}
											}

										}

									}
								}

							}

						}
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Updating Volterra UztnaPolicyTemplate obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_uztna_uztna_policy_template.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating UztnaPolicyTemplate: %s", err)
	}

	return resourceVolterraUztnaPolicyTemplateRead(d, meta)
}

func resourceVolterraUztnaPolicyTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_uztna_uztna_policy_template.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaPolicyTemplate %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaPolicyTemplate before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra UztnaPolicyTemplate obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_uztna_uztna_policy_template.ObjectType, namespace, name, opts...)
}
