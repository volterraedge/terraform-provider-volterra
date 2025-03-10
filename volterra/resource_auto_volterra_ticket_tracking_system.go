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
	ves_io_schema_ticket_management_ticket_tracking_system "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ticket_management/ticket_tracking_system"
)

// resourceVolterraTicketTrackingSystem is implementation of Volterra's TicketTrackingSystem resources
func resourceVolterraTicketTrackingSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraTicketTrackingSystemCreate,
		Read:   resourceVolterraTicketTrackingSystemRead,
		Update: resourceVolterraTicketTrackingSystemUpdate,
		Delete: resourceVolterraTicketTrackingSystemDelete,

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

			"jira_config": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"adhoc_rest_api": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"account_email": {
										Type:     schema.TypeString,
										Required: true,
									},

									"api_token": {
										Type:     schema.TypeString,
										Required: true,
									},

									"encrypted_api_token": {

										Type:       schema.TypeList,
										MaxItems:   1,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"blindfold_secret_info_internal": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"decryption_provider": {
																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"location": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"store_provider": {
																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},

												"secret_encoding_type": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"blindfold_secret_info": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"decryption_provider": {
																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"location": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"store_provider": {
																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},

												"clear_secret_info": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"provider": {
																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"url": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},

												"vault_secret_info": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": {
																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"location": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"provider": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"secret_encoding": {
																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"version": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},

												"wingman_secret_info": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"name": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},
											},
										},
									},

									"organization_domain": {
										Type:     schema.TypeString,
										Required: true,
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

// resourceVolterraTicketTrackingSystemCreate creates TicketTrackingSystem resource
func resourceVolterraTicketTrackingSystemCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_ticket_management_ticket_tracking_system.CreateSpecType{}
	createReq := &ves_io_schema_ticket_management_ticket_tracking_system.CreateRequest{
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

	//provider_config

	providerConfigTypeFound := false

	if v, ok := d.GetOk("jira_config"); ok && !providerConfigTypeFound {

		providerConfigTypeFound = true
		providerConfigInt := &ves_io_schema_ticket_management_ticket_tracking_system.CreateSpecType_JiraConfig{}
		providerConfigInt.JiraConfig = &ves_io_schema_ticket_management_ticket_tracking_system.JiraConfigurationType{}
		createSpec.ProviderConfig = providerConfigInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				configTypeFound := false

				if v, ok := cs["adhoc_rest_api"]; ok && !isIntfNil(v) && !configTypeFound {

					configTypeFound = true
					configInt := &ves_io_schema_ticket_management_ticket_tracking_system.JiraConfigurationType_AdhocRestApi{}
					configInt.AdhocRestApi = &ves_io_schema_ticket_management_ticket_tracking_system.JiraAdhocRestApiConfigurationType{}
					providerConfigInt.JiraConfig.Config = configInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["account_email"]; ok && !isIntfNil(v) {

								configInt.AdhocRestApi.AccountEmail = v.(string)

							}

							if v, ok := cs["api_token"]; ok && !isIntfNil(v) {

								configInt.AdhocRestApi.ApiToken = v.(string)

							}

							if v, ok := cs["encrypted_api_token"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								encryptedApiToken := &ves_io_schema.SecretType{}
								configInt.AdhocRestApi.EncryptedApiToken = encryptedApiToken
								for _, set := range sl {
									if set != nil {
										encryptedApiTokenMapStrToI := set.(map[string]interface{})

										if v, ok := encryptedApiTokenMapStrToI["blindfold_secret_info_internal"]; ok && !isIntfNil(v) {

											sl := v.([]interface{})
											blindfoldSecretInfoInternal := &ves_io_schema.BlindfoldSecretInfoType{}
											encryptedApiToken.BlindfoldSecretInfoInternal = blindfoldSecretInfoInternal
											for _, set := range sl {
												if set != nil {
													blindfoldSecretInfoInternalMapStrToI := set.(map[string]interface{})

													if w, ok := blindfoldSecretInfoInternalMapStrToI["decryption_provider"]; ok && !isIntfNil(w) {
														blindfoldSecretInfoInternal.DecryptionProvider = w.(string)
													}

													if w, ok := blindfoldSecretInfoInternalMapStrToI["location"]; ok && !isIntfNil(w) {
														blindfoldSecretInfoInternal.Location = w.(string)
													}

													if w, ok := blindfoldSecretInfoInternalMapStrToI["store_provider"]; ok && !isIntfNil(w) {
														blindfoldSecretInfoInternal.StoreProvider = w.(string)
													}

												}
											}

										}

										if v, ok := encryptedApiTokenMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

											encryptedApiToken.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

										}

										secretInfoOneofTypeFound := false

										if v, ok := encryptedApiTokenMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

											secretInfoOneofTypeFound = true
											secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
											secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
											encryptedApiToken.SecretInfoOneof = secretInfoOneofInt

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

										if v, ok := encryptedApiTokenMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

											secretInfoOneofTypeFound = true
											secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
											secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
											encryptedApiToken.SecretInfoOneof = secretInfoOneofInt

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

										if v, ok := encryptedApiTokenMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

											secretInfoOneofTypeFound = true
											secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
											secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
											encryptedApiToken.SecretInfoOneof = secretInfoOneofInt

											sl := v.([]interface{})
											for _, set := range sl {
												if set != nil {
													cs := set.(map[string]interface{})

													if v, ok := cs["key"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.Key = v.(string)

													}

													if v, ok := cs["location"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.Location = v.(string)

													}

													if v, ok := cs["provider"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.Provider = v.(string)

													}

													if v, ok := cs["secret_encoding"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.SecretEncoding = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

													}

													if v, ok := cs["version"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.Version = uint32(v.(int))

													}

												}
											}

										}

										if v, ok := encryptedApiTokenMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

											secretInfoOneofTypeFound = true
											secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
											secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
											encryptedApiToken.SecretInfoOneof = secretInfoOneofInt

											sl := v.([]interface{})
											for _, set := range sl {
												if set != nil {
													cs := set.(map[string]interface{})

													if v, ok := cs["name"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)

													}

												}
											}

										}

									}
								}

							}

							if v, ok := cs["organization_domain"]; ok && !isIntfNil(v) {

								configInt.AdhocRestApi.OrganizationDomain = v.(string)

							}

						}
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Creating Volterra TicketTrackingSystem object with struct: %+v", createReq)

	createTicketTrackingSystemResp, err := client.CreateObject(context.Background(), ves_io_schema_ticket_management_ticket_tracking_system.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating TicketTrackingSystem: %s", err)
	}
	d.SetId(createTicketTrackingSystemResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraTicketTrackingSystemRead(d, meta)
}

func resourceVolterraTicketTrackingSystemRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_ticket_management_ticket_tracking_system.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] TicketTrackingSystem %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra TicketTrackingSystem %q: %s", d.Id(), err)
	}
	return setTicketTrackingSystemFields(client, d, resp)
}

func setTicketTrackingSystemFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraTicketTrackingSystemUpdate updates TicketTrackingSystem resource
func resourceVolterraTicketTrackingSystemUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_ticket_management_ticket_tracking_system.ReplaceSpecType{}
	updateReq := &ves_io_schema_ticket_management_ticket_tracking_system.ReplaceRequest{
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

	providerConfigTypeFound := false

	if v, ok := d.GetOk("jira_config"); ok && !providerConfigTypeFound {

		providerConfigTypeFound = true
		providerConfigInt := &ves_io_schema_ticket_management_ticket_tracking_system.ReplaceSpecType_JiraConfig{}
		providerConfigInt.JiraConfig = &ves_io_schema_ticket_management_ticket_tracking_system.JiraConfigurationType{}
		updateSpec.ProviderConfig = providerConfigInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				configTypeFound := false

				if v, ok := cs["adhoc_rest_api"]; ok && !isIntfNil(v) && !configTypeFound {

					configTypeFound = true
					configInt := &ves_io_schema_ticket_management_ticket_tracking_system.JiraConfigurationType_AdhocRestApi{}
					configInt.AdhocRestApi = &ves_io_schema_ticket_management_ticket_tracking_system.JiraAdhocRestApiConfigurationType{}
					providerConfigInt.JiraConfig.Config = configInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["account_email"]; ok && !isIntfNil(v) {

								configInt.AdhocRestApi.AccountEmail = v.(string)

							}

							if v, ok := cs["api_token"]; ok && !isIntfNil(v) {

								configInt.AdhocRestApi.ApiToken = v.(string)

							}

							if v, ok := cs["encrypted_api_token"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								encryptedApiToken := &ves_io_schema.SecretType{}
								configInt.AdhocRestApi.EncryptedApiToken = encryptedApiToken
								for _, set := range sl {
									if set != nil {
										encryptedApiTokenMapStrToI := set.(map[string]interface{})

										if v, ok := encryptedApiTokenMapStrToI["blindfold_secret_info_internal"]; ok && !isIntfNil(v) {

											sl := v.([]interface{})
											blindfoldSecretInfoInternal := &ves_io_schema.BlindfoldSecretInfoType{}
											encryptedApiToken.BlindfoldSecretInfoInternal = blindfoldSecretInfoInternal
											for _, set := range sl {
												if set != nil {
													blindfoldSecretInfoInternalMapStrToI := set.(map[string]interface{})

													if w, ok := blindfoldSecretInfoInternalMapStrToI["decryption_provider"]; ok && !isIntfNil(w) {
														blindfoldSecretInfoInternal.DecryptionProvider = w.(string)
													}

													if w, ok := blindfoldSecretInfoInternalMapStrToI["location"]; ok && !isIntfNil(w) {
														blindfoldSecretInfoInternal.Location = w.(string)
													}

													if w, ok := blindfoldSecretInfoInternalMapStrToI["store_provider"]; ok && !isIntfNil(w) {
														blindfoldSecretInfoInternal.StoreProvider = w.(string)
													}

												}
											}

										}

										if v, ok := encryptedApiTokenMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

											encryptedApiToken.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

										}

										secretInfoOneofTypeFound := false

										if v, ok := encryptedApiTokenMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

											secretInfoOneofTypeFound = true
											secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
											secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
											encryptedApiToken.SecretInfoOneof = secretInfoOneofInt

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

										if v, ok := encryptedApiTokenMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

											secretInfoOneofTypeFound = true
											secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
											secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
											encryptedApiToken.SecretInfoOneof = secretInfoOneofInt

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

										if v, ok := encryptedApiTokenMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

											secretInfoOneofTypeFound = true
											secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
											secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
											encryptedApiToken.SecretInfoOneof = secretInfoOneofInt

											sl := v.([]interface{})
											for _, set := range sl {
												if set != nil {
													cs := set.(map[string]interface{})

													if v, ok := cs["key"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.Key = v.(string)

													}

													if v, ok := cs["location"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.Location = v.(string)

													}

													if v, ok := cs["provider"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.Provider = v.(string)

													}

													if v, ok := cs["secret_encoding"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.SecretEncoding = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

													}

													if v, ok := cs["version"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.VaultSecretInfo.Version = uint32(v.(int))

													}

												}
											}

										}

										if v, ok := encryptedApiTokenMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

											secretInfoOneofTypeFound = true
											secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
											secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
											encryptedApiToken.SecretInfoOneof = secretInfoOneofInt

											sl := v.([]interface{})
											for _, set := range sl {
												if set != nil {
													cs := set.(map[string]interface{})

													if v, ok := cs["name"]; ok && !isIntfNil(v) {

														secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)

													}

												}
											}

										}

									}
								}

							}

							if v, ok := cs["organization_domain"]; ok && !isIntfNil(v) {

								configInt.AdhocRestApi.OrganizationDomain = v.(string)

							}

						}
					}

				}

			}
		}

	}

	log.Printf("[DEBUG] Updating Volterra TicketTrackingSystem obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_ticket_management_ticket_tracking_system.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating TicketTrackingSystem: %s", err)
	}

	return resourceVolterraTicketTrackingSystemRead(d, meta)
}

func resourceVolterraTicketTrackingSystemDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_ticket_management_ticket_tracking_system.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] TicketTrackingSystem %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra TicketTrackingSystem before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra TicketTrackingSystem obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_ticket_management_ticket_tracking_system.ObjectType, namespace, name)
}
