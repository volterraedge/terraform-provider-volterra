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
	ves_io_schema_advertise_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/advertise_policy"
)

// resourceVolterraAdvertisePolicy is implementation of Volterra's AdvertisePolicy resources
func resourceVolterraAdvertisePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAdvertisePolicyCreate,
		Read:   resourceVolterraAdvertisePolicyRead,
		Update: resourceVolterraAdvertisePolicyUpdate,
		Delete: resourceVolterraAdvertisePolicyDelete,

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

			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"public_ip": {

				Type:     schema.TypeList,
				Optional: true,
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

			"skip_xff_append": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"tls_parameters": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"common_params": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cipher_suites": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"maximum_protocol_version": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"minimum_protocol_version": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"tls_certificates": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate_url": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"description": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"private_key": {

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

									"trusted_ca_url": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"validation_params": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"skip_hostname_verification": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"trusted_ca_url": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"use_volterra_trusted_ca_url": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"verify_subject_alt_names": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},

						"require_client_certificate": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"where": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"site": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"ref": {

										Type:     schema.TypeList,
										Optional: true,
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

						"virtual_network": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ref": {

										Type:     schema.TypeList,
										Optional: true,
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

						"virtual_site": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"ref": {

										Type:     schema.TypeList,
										Optional: true,
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
	}
}

// resourceVolterraAdvertisePolicyCreate creates AdvertisePolicy resource
func resourceVolterraAdvertisePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_advertise_policy.CreateSpecType{}
	createReq := &ves_io_schema_advertise_policy.CreateRequest{
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

	if v, ok := d.GetOk("address"); ok && !isIntfNil(v) {

		createSpec.Address =
			v.(string)
	}

	if v, ok := d.GetOk("port"); ok && !isIntfNil(v) {

		createSpec.Port =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("protocol"); ok && !isIntfNil(v) {

		createSpec.Protocol =
			v.(string)
	}

	if v, ok := d.GetOk("public_ip"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		publicIpInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.PublicIp = publicIpInt
		for i, ps := range sl {

			piMapToStrVal := ps.(map[string]interface{})
			publicIpInt[i] = &ves_io_schema.ObjectRefType{}

			publicIpInt[i].Kind = "public_ip"

			if v, ok := piMapToStrVal["name"]; ok && !isIntfNil(v) {
				publicIpInt[i].Name = v.(string)
			}

			if v, ok := piMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				publicIpInt[i].Namespace = v.(string)
			}

			if v, ok := piMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				publicIpInt[i].Tenant = v.(string)
			}

			if v, ok := piMapToStrVal["uid"]; ok && !isIntfNil(v) {
				publicIpInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("skip_xff_append"); ok && !isIntfNil(v) {

		createSpec.SkipXffAppend =
			v.(bool)
	}

	if v, ok := d.GetOk("tls_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tlsParameters := &ves_io_schema.DownstreamTlsParamsType{}
		createSpec.TlsParameters = tlsParameters
		for _, set := range sl {

			tlsParametersMapStrToI := set.(map[string]interface{})

			if v, ok := tlsParametersMapStrToI["common_params"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				commonParams := &ves_io_schema.TlsParamsType{}
				tlsParameters.CommonParams = commonParams
				for _, set := range sl {

					commonParamsMapStrToI := set.(map[string]interface{})

					if w, ok := commonParamsMapStrToI["cipher_suites"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						commonParams.CipherSuites = ls
					}

					if v, ok := commonParamsMapStrToI["maximum_protocol_version"]; ok && !isIntfNil(v) {

						commonParams.MaximumProtocolVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

					}

					if v, ok := commonParamsMapStrToI["minimum_protocol_version"]; ok && !isIntfNil(v) {

						commonParams.MinimumProtocolVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

					}

					if v, ok := commonParamsMapStrToI["tls_certificates"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						tlsCertificates := make([]*ves_io_schema.TlsCertificateType, len(sl))
						commonParams.TlsCertificates = tlsCertificates
						for i, set := range sl {
							tlsCertificates[i] = &ves_io_schema.TlsCertificateType{}

							tlsCertificatesMapStrToI := set.(map[string]interface{})

							if w, ok := tlsCertificatesMapStrToI["certificate_url"]; ok && !isIntfNil(w) {
								tlsCertificates[i].CertificateUrl = w.(string)
							}

							if w, ok := tlsCertificatesMapStrToI["description"]; ok && !isIntfNil(w) {
								tlsCertificates[i].Description = w.(string)
							}

							if v, ok := tlsCertificatesMapStrToI["private_key"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								privateKey := &ves_io_schema.SecretType{}
								tlsCertificates[i].PrivateKey = privateKey
								for _, set := range sl {

									privateKeyMapStrToI := set.(map[string]interface{})

									if v, ok := privateKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										privateKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := privateKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := privateKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := privateKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := privateKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

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

					if w, ok := commonParamsMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
						commonParams.TrustedCaUrl = w.(string)
					}

					if v, ok := commonParamsMapStrToI["validation_params"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						validationParams := &ves_io_schema.TlsValidationParamsType{}
						commonParams.ValidationParams = validationParams
						for _, set := range sl {

							validationParamsMapStrToI := set.(map[string]interface{})

							if w, ok := validationParamsMapStrToI["skip_hostname_verification"]; ok && !isIntfNil(w) {
								validationParams.SkipHostnameVerification = w.(bool)
							}

							if w, ok := validationParamsMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
								validationParams.TrustedCaUrl = w.(string)
							}

							if w, ok := validationParamsMapStrToI["use_volterra_trusted_ca_url"]; ok && !isIntfNil(w) {
								validationParams.UseVolterraTrustedCaUrl = w.(bool)
							}

							if w, ok := validationParamsMapStrToI["verify_subject_alt_names"]; ok && !isIntfNil(w) {
								ls := make([]string, len(w.([]interface{})))
								for i, v := range w.([]interface{}) {
									ls[i] = v.(string)
								}
								validationParams.VerifySubjectAltNames = ls
							}

						}

					}

				}

			}

			if w, ok := tlsParametersMapStrToI["require_client_certificate"]; ok && !isIntfNil(w) {
				tlsParameters.RequireClientCertificate = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("where"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		where := &ves_io_schema.NetworkSiteRefSelector{}
		createSpec.Where = where
		for _, set := range sl {

			whereMapStrToI := set.(map[string]interface{})

			refOrSelectorTypeFound := false

			if v, ok := whereMapStrToI["site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_Site{}
				refOrSelectorInt.Site = &ves_io_schema.SiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.Site.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.Site.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_network"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualNetwork{}
				refOrSelectorInt.VirtualNetwork = &ves_io_schema.NetworkRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualNetwork.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_network"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualSite{}
				refOrSelectorInt.VirtualSite = &ves_io_schema.VSiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.VirtualSite.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualSite.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra AdvertisePolicy object with struct: %+v", createReq)

	createAdvertisePolicyResp, err := client.CreateObject(context.Background(), ves_io_schema_advertise_policy.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating AdvertisePolicy: %s", err)
	}
	d.SetId(createAdvertisePolicyResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraAdvertisePolicyRead(d, meta)
}

func resourceVolterraAdvertisePolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_advertise_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AdvertisePolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AdvertisePolicy %q: %s", d.Id(), err)
	}
	return setAdvertisePolicyFields(client, d, resp)
}

func setAdvertisePolicyFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraAdvertisePolicyUpdate updates AdvertisePolicy resource
func resourceVolterraAdvertisePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_advertise_policy.ReplaceSpecType{}
	updateReq := &ves_io_schema_advertise_policy.ReplaceRequest{
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

	if v, ok := d.GetOk("address"); ok && !isIntfNil(v) {

		updateSpec.Address =
			v.(string)
	}

	if v, ok := d.GetOk("port"); ok && !isIntfNil(v) {

		updateSpec.Port =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("protocol"); ok && !isIntfNil(v) {

		updateSpec.Protocol =
			v.(string)
	}

	if v, ok := d.GetOk("public_ip"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		publicIpInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.PublicIp = publicIpInt
		for i, ps := range sl {

			piMapToStrVal := ps.(map[string]interface{})
			publicIpInt[i] = &ves_io_schema.ObjectRefType{}

			publicIpInt[i].Kind = "public_ip"

			if v, ok := piMapToStrVal["name"]; ok && !isIntfNil(v) {
				publicIpInt[i].Name = v.(string)
			}

			if v, ok := piMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				publicIpInt[i].Namespace = v.(string)
			}

			if v, ok := piMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				publicIpInt[i].Tenant = v.(string)
			}

			if v, ok := piMapToStrVal["uid"]; ok && !isIntfNil(v) {
				publicIpInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("skip_xff_append"); ok && !isIntfNil(v) {

		updateSpec.SkipXffAppend =
			v.(bool)
	}

	if v, ok := d.GetOk("tls_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tlsParameters := &ves_io_schema.DownstreamTlsParamsType{}
		updateSpec.TlsParameters = tlsParameters
		for _, set := range sl {

			tlsParametersMapStrToI := set.(map[string]interface{})

			if v, ok := tlsParametersMapStrToI["common_params"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				commonParams := &ves_io_schema.TlsParamsType{}
				tlsParameters.CommonParams = commonParams
				for _, set := range sl {

					commonParamsMapStrToI := set.(map[string]interface{})

					if w, ok := commonParamsMapStrToI["cipher_suites"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						commonParams.CipherSuites = ls
					}

					if v, ok := commonParamsMapStrToI["maximum_protocol_version"]; ok && !isIntfNil(v) {

						commonParams.MaximumProtocolVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

					}

					if v, ok := commonParamsMapStrToI["minimum_protocol_version"]; ok && !isIntfNil(v) {

						commonParams.MinimumProtocolVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

					}

					if v, ok := commonParamsMapStrToI["tls_certificates"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						tlsCertificates := make([]*ves_io_schema.TlsCertificateType, len(sl))
						commonParams.TlsCertificates = tlsCertificates
						for i, set := range sl {
							tlsCertificates[i] = &ves_io_schema.TlsCertificateType{}

							tlsCertificatesMapStrToI := set.(map[string]interface{})

							if w, ok := tlsCertificatesMapStrToI["certificate_url"]; ok && !isIntfNil(w) {
								tlsCertificates[i].CertificateUrl = w.(string)
							}

							if w, ok := tlsCertificatesMapStrToI["description"]; ok && !isIntfNil(w) {
								tlsCertificates[i].Description = w.(string)
							}

							if v, ok := tlsCertificatesMapStrToI["private_key"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								privateKey := &ves_io_schema.SecretType{}
								tlsCertificates[i].PrivateKey = privateKey
								for _, set := range sl {

									privateKeyMapStrToI := set.(map[string]interface{})

									if v, ok := privateKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

										privateKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

									}

									secretInfoOneofTypeFound := false

									if v, ok := privateKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
										secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := privateKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
										secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := privateKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
										secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

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

									if v, ok := privateKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

										secretInfoOneofTypeFound = true
										secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
										secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
										privateKey.SecretInfoOneof = secretInfoOneofInt

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

					if w, ok := commonParamsMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
						commonParams.TrustedCaUrl = w.(string)
					}

					if v, ok := commonParamsMapStrToI["validation_params"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						validationParams := &ves_io_schema.TlsValidationParamsType{}
						commonParams.ValidationParams = validationParams
						for _, set := range sl {

							validationParamsMapStrToI := set.(map[string]interface{})

							if w, ok := validationParamsMapStrToI["skip_hostname_verification"]; ok && !isIntfNil(w) {
								validationParams.SkipHostnameVerification = w.(bool)
							}

							if w, ok := validationParamsMapStrToI["trusted_ca_url"]; ok && !isIntfNil(w) {
								validationParams.TrustedCaUrl = w.(string)
							}

							if w, ok := validationParamsMapStrToI["use_volterra_trusted_ca_url"]; ok && !isIntfNil(w) {
								validationParams.UseVolterraTrustedCaUrl = w.(bool)
							}

							if w, ok := validationParamsMapStrToI["verify_subject_alt_names"]; ok && !isIntfNil(w) {
								ls := make([]string, len(w.([]interface{})))
								for i, v := range w.([]interface{}) {
									ls[i] = v.(string)
								}
								validationParams.VerifySubjectAltNames = ls
							}

						}

					}

				}

			}

			if w, ok := tlsParametersMapStrToI["require_client_certificate"]; ok && !isIntfNil(w) {
				tlsParameters.RequireClientCertificate = w.(bool)
			}

		}

	}

	if v, ok := d.GetOk("where"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		where := &ves_io_schema.NetworkSiteRefSelector{}
		updateSpec.Where = where
		for _, set := range sl {

			whereMapStrToI := set.(map[string]interface{})

			refOrSelectorTypeFound := false

			if v, ok := whereMapStrToI["site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_Site{}
				refOrSelectorInt.Site = &ves_io_schema.SiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.Site.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.Site.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_network"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualNetwork{}
				refOrSelectorInt.VirtualNetwork = &ves_io_schema.NetworkRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualNetwork.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_network"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualSite{}
				refOrSelectorInt.VirtualSite = &ves_io_schema.VSiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.VirtualSite.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualSite.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra AdvertisePolicy obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_advertise_policy.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating AdvertisePolicy: %s", err)
	}

	return resourceVolterraAdvertisePolicyRead(d, meta)
}

func resourceVolterraAdvertisePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_advertise_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AdvertisePolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AdvertisePolicy before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra AdvertisePolicy obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_advertise_policy.ObjectType, namespace, name)
}