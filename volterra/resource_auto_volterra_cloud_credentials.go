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
	ves_io_schema_cloud_credentials "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cloud_credentials"
)

// resourceVolterraCloudCredentials is implementation of Volterra's CloudCredentials resources
func resourceVolterraCloudCredentials() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraCloudCredentialsCreate,
		Read:   resourceVolterraCloudCredentialsRead,
		Update: resourceVolterraCloudCredentialsUpdate,
		Delete: resourceVolterraCloudCredentialsDelete,

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

			"aws_secret_key": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_key": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"secret_key": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"secret_encoding_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"blindfold_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"decryption_provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"location": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"store_provider": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"clear_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"url": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"vault_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"location": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"secret_encoding": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"version": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"wingman_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
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

			"azure_client_secret": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"client_id": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"client_secret": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"secret_encoding_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"blindfold_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"decryption_provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"location": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"store_provider": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"clear_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"url": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"vault_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"location": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"secret_encoding": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"version": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"wingman_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"subscription_id": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"tenant_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"azure_pfx_certificate": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"certificate_url": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"client_id": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"password": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"secret_encoding_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"blindfold_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"decryption_provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"location": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"store_provider": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"clear_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"url": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"vault_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"location": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"secret_encoding": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"version": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"wingman_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"subscription_id": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"tenant_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"gcp_cred_file": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"credential_file": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"secret_encoding_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"blindfold_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"decryption_provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"location": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"store_provider": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"clear_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"url": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"vault_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"location": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"provider": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"secret_encoding": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"version": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"wingman_secret_info": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
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
	}
}

// resourceVolterraCloudCredentialsCreate creates CloudCredentials resource
func resourceVolterraCloudCredentialsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_cloud_credentials.CreateSpecType{}
	createReq := &ves_io_schema_cloud_credentials.CreateRequest{
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

	cloudTypeFound := false

	if v, ok := d.GetOk("aws_secret_key"); ok && !cloudTypeFound {

		cloudTypeFound = true
		cloudInt := &ves_io_schema_cloud_credentials.CreateSpecType_AwsSecretKey{}
		cloudInt.AwsSecretKey = &ves_io_schema_cloud_credentials.AWSSecretType{}
		createSpec.Cloud = cloudInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["access_key"]; ok && !isIntfNil(v) {

				cloudInt.AwsSecretKey.AccessKey = v.(string)
			}

			if v, ok := cs["secret_key"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				secretKey := &ves_io_schema.SecretType{}
				cloudInt.AwsSecretKey.SecretKey = secretKey
				for _, set := range sl {

					secretKeyMapStrToI := set.(map[string]interface{})

					if v, ok := secretKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						secretKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := secretKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						secretKey.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := secretKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						secretKey.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
							}

						}

					}

					if v, ok := secretKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						secretKey.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := secretKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						secretKey.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

	if v, ok := d.GetOk("azure_client_secret"); ok && !cloudTypeFound {

		cloudTypeFound = true
		cloudInt := &ves_io_schema_cloud_credentials.CreateSpecType_AzureClientSecret{}
		cloudInt.AzureClientSecret = &ves_io_schema_cloud_credentials.AzureSecretType{}
		createSpec.Cloud = cloudInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["client_id"]; ok && !isIntfNil(v) {

				cloudInt.AzureClientSecret.ClientId = v.(string)
			}

			if v, ok := cs["client_secret"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				clientSecret := &ves_io_schema.SecretType{}
				cloudInt.AzureClientSecret.ClientSecret = clientSecret
				for _, set := range sl {

					clientSecretMapStrToI := set.(map[string]interface{})

					if v, ok := clientSecretMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						clientSecret.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := clientSecretMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						clientSecret.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := clientSecretMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						clientSecret.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
							}

						}

					}

					if v, ok := clientSecretMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						clientSecret.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := clientSecretMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						clientSecret.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["subscription_id"]; ok && !isIntfNil(v) {

				cloudInt.AzureClientSecret.SubscriptionId = v.(string)
			}

			if v, ok := cs["tenant_id"]; ok && !isIntfNil(v) {

				cloudInt.AzureClientSecret.TenantId = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("azure_pfx_certificate"); ok && !cloudTypeFound {

		cloudTypeFound = true
		cloudInt := &ves_io_schema_cloud_credentials.CreateSpecType_AzurePfxCertificate{}
		cloudInt.AzurePfxCertificate = &ves_io_schema_cloud_credentials.AzurePfxType{}
		createSpec.Cloud = cloudInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["certificate_url"]; ok && !isIntfNil(v) {

				cloudInt.AzurePfxCertificate.CertificateUrl = v.(string)
			}

			if v, ok := cs["client_id"]; ok && !isIntfNil(v) {

				cloudInt.AzurePfxCertificate.ClientId = v.(string)
			}

			if v, ok := cs["password"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				password := &ves_io_schema.SecretType{}
				cloudInt.AzurePfxCertificate.Password = password
				for _, set := range sl {

					passwordMapStrToI := set.(map[string]interface{})

					if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						password.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						password.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
							}

						}

					}

					if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						password.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						password.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["subscription_id"]; ok && !isIntfNil(v) {

				cloudInt.AzurePfxCertificate.SubscriptionId = v.(string)
			}

			if v, ok := cs["tenant_id"]; ok && !isIntfNil(v) {

				cloudInt.AzurePfxCertificate.TenantId = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("gcp_cred_file"); ok && !cloudTypeFound {

		cloudTypeFound = true
		cloudInt := &ves_io_schema_cloud_credentials.CreateSpecType_GcpCredFile{}
		cloudInt.GcpCredFile = &ves_io_schema_cloud_credentials.GCPCredFileType{}
		createSpec.Cloud = cloudInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["credential_file"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				credentialFile := &ves_io_schema.SecretType{}
				cloudInt.GcpCredFile.CredentialFile = credentialFile
				for _, set := range sl {

					credentialFileMapStrToI := set.(map[string]interface{})

					if v, ok := credentialFileMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						credentialFile.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := credentialFileMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						credentialFile.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := credentialFileMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						credentialFile.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
							}

						}

					}

					if v, ok := credentialFileMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						credentialFile.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := credentialFileMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						credentialFile.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

	log.Printf("[DEBUG] Creating Volterra CloudCredentials object with struct: %+v", createReq)

	createCloudCredentialsResp, err := client.CreateObject(context.Background(), ves_io_schema_cloud_credentials.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating CloudCredentials: %s", err)
	}
	d.SetId(createCloudCredentialsResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraCloudCredentialsRead(d, meta)
}

func resourceVolterraCloudCredentialsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_cloud_credentials.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] CloudCredentials %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra CloudCredentials %q: %s", d.Id(), err)
	}
	return setCloudCredentialsFields(client, d, resp)
}

func setCloudCredentialsFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraCloudCredentialsUpdate updates CloudCredentials resource
func resourceVolterraCloudCredentialsUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_cloud_credentials.ReplaceSpecType{}
	updateReq := &ves_io_schema_cloud_credentials.ReplaceRequest{
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

	cloudTypeFound := false

	if v, ok := d.GetOk("aws_secret_key"); ok && !cloudTypeFound {

		cloudTypeFound = true
		cloudInt := &ves_io_schema_cloud_credentials.ReplaceSpecType_AwsSecretKey{}
		cloudInt.AwsSecretKey = &ves_io_schema_cloud_credentials.AWSSecretType{}
		updateSpec.Cloud = cloudInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["access_key"]; ok && !isIntfNil(v) {

				cloudInt.AwsSecretKey.AccessKey = v.(string)
			}

			if v, ok := cs["secret_key"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				secretKey := &ves_io_schema.SecretType{}
				cloudInt.AwsSecretKey.SecretKey = secretKey
				for _, set := range sl {

					secretKeyMapStrToI := set.(map[string]interface{})

					if v, ok := secretKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						secretKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := secretKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						secretKey.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := secretKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						secretKey.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
							}

						}

					}

					if v, ok := secretKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						secretKey.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := secretKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						secretKey.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

	if v, ok := d.GetOk("azure_client_secret"); ok && !cloudTypeFound {

		cloudTypeFound = true
		cloudInt := &ves_io_schema_cloud_credentials.ReplaceSpecType_AzureClientSecret{}
		cloudInt.AzureClientSecret = &ves_io_schema_cloud_credentials.AzureSecretType{}
		updateSpec.Cloud = cloudInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["client_id"]; ok && !isIntfNil(v) {

				cloudInt.AzureClientSecret.ClientId = v.(string)
			}

			if v, ok := cs["client_secret"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				clientSecret := &ves_io_schema.SecretType{}
				cloudInt.AzureClientSecret.ClientSecret = clientSecret
				for _, set := range sl {

					clientSecretMapStrToI := set.(map[string]interface{})

					if v, ok := clientSecretMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						clientSecret.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := clientSecretMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						clientSecret.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := clientSecretMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						clientSecret.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
							}

						}

					}

					if v, ok := clientSecretMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						clientSecret.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := clientSecretMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						clientSecret.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["subscription_id"]; ok && !isIntfNil(v) {

				cloudInt.AzureClientSecret.SubscriptionId = v.(string)
			}

			if v, ok := cs["tenant_id"]; ok && !isIntfNil(v) {

				cloudInt.AzureClientSecret.TenantId = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("azure_pfx_certificate"); ok && !cloudTypeFound {

		cloudTypeFound = true
		cloudInt := &ves_io_schema_cloud_credentials.ReplaceSpecType_AzurePfxCertificate{}
		cloudInt.AzurePfxCertificate = &ves_io_schema_cloud_credentials.AzurePfxType{}
		updateSpec.Cloud = cloudInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["certificate_url"]; ok && !isIntfNil(v) {

				cloudInt.AzurePfxCertificate.CertificateUrl = v.(string)
			}

			if v, ok := cs["client_id"]; ok && !isIntfNil(v) {

				cloudInt.AzurePfxCertificate.ClientId = v.(string)
			}

			if v, ok := cs["password"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				password := &ves_io_schema.SecretType{}
				cloudInt.AzurePfxCertificate.Password = password
				for _, set := range sl {

					passwordMapStrToI := set.(map[string]interface{})

					if v, ok := passwordMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						password.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := passwordMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						password.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := passwordMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						password.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
							}

						}

					}

					if v, ok := passwordMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						password.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := passwordMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						password.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["subscription_id"]; ok && !isIntfNil(v) {

				cloudInt.AzurePfxCertificate.SubscriptionId = v.(string)
			}

			if v, ok := cs["tenant_id"]; ok && !isIntfNil(v) {

				cloudInt.AzurePfxCertificate.TenantId = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("gcp_cred_file"); ok && !cloudTypeFound {

		cloudTypeFound = true
		cloudInt := &ves_io_schema_cloud_credentials.ReplaceSpecType_GcpCredFile{}
		cloudInt.GcpCredFile = &ves_io_schema_cloud_credentials.GCPCredFileType{}
		updateSpec.Cloud = cloudInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["credential_file"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				credentialFile := &ves_io_schema.SecretType{}
				cloudInt.GcpCredFile.CredentialFile = credentialFile
				for _, set := range sl {

					credentialFileMapStrToI := set.(map[string]interface{})

					if v, ok := credentialFileMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

						credentialFile.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					secretInfoOneofTypeFound := false

					if v, ok := credentialFileMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
						secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
						credentialFile.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := credentialFileMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
						secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
						credentialFile.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["provider"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)
							}

							if v, ok := cs["url"]; ok && !isIntfNil(v) {

								secretInfoOneofInt.ClearSecretInfo.Url = v.(string)
							}

						}

					}

					if v, ok := credentialFileMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
						secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
						credentialFile.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

					if v, ok := credentialFileMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

						secretInfoOneofTypeFound = true
						secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
						secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
						credentialFile.SecretInfoOneof = secretInfoOneofInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
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

	log.Printf("[DEBUG] Updating Volterra CloudCredentials obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_cloud_credentials.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating CloudCredentials: %s", err)
	}

	return resourceVolterraCloudCredentialsRead(d, meta)
}

func resourceVolterraCloudCredentialsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_cloud_credentials.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] CloudCredentials %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra CloudCredentials before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra CloudCredentials obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_cloud_credentials.ObjectType, namespace, name)
}
