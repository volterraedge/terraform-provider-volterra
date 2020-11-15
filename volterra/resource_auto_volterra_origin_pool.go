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
	ves_io_schema_cluster "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/cluster"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_views_origin_pool "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/origin_pool"
)

// resourceVolterraOriginPool is implementation of Volterra's OriginPool resources
func resourceVolterraOriginPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraOriginPoolCreate,
		Read:   resourceVolterraOriginPoolRead,
		Update: resourceVolterraOriginPoolUpdate,
		Delete: resourceVolterraOriginPoolDelete,

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

			"endpoint_selection": {
				Type:     schema.TypeString,
				Required: true,
			},

			"healthcheck": {

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

			"loadbalancer_algorithm": {
				Type:     schema.TypeString,
				Required: true,
			},

			"origin_servers": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"consul_service": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"inside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"outside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"service_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"site_locator": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"site": {

													Type:     schema.TypeSet,
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

												"virtual_site": {

													Type:     schema.TypeSet,
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

						"custom_endpoint_object": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"endpoint": {

										Type:     schema.TypeSet,
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

						"k8s_service": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"inside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"outside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"vk8s_networks": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"service_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"site_locator": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"site": {

													Type:     schema.TypeSet,
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

												"virtual_site": {

													Type:     schema.TypeSet,
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

						"private_ip": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"inside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"outside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"site_locator": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"site": {

													Type:     schema.TypeSet,
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

												"virtual_site": {

													Type:     schema.TypeSet,
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

						"private_name": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"inside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"outside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"site_locator": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"site": {

													Type:     schema.TypeSet,
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

												"virtual_site": {

													Type:     schema.TypeSet,
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

						"public_ip": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"public_name": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"no_tls": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"use_tls": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"no_mtls": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"use_mtls": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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
								},
							},
						},

						"skip_server_verification": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"use_server_verification": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"trusted_ca_url": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"volterra_trusted_ca": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"disable_sni": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"sni": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"use_host_header_as_sni": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"tls_config": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_security": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cipher_suites": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"max_version": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"min_version": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"default_security": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"low_security": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"medium_security": {

										Type:     schema.TypeBool,
										Optional: true,
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

// resourceVolterraOriginPoolCreate creates OriginPool resource
func resourceVolterraOriginPoolCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_origin_pool.CreateSpecType{}
	createReq := &ves_io_schema_views_origin_pool.CreateRequest{
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

	if v, ok := d.GetOk("endpoint_selection"); ok && !isIntfNil(v) {

		createSpec.EndpointSelection = ves_io_schema_cluster.EndpointSelectionPolicy(ves_io_schema_cluster.EndpointSelectionPolicy_value[v.(string)])

	}

	if v, ok := d.GetOk("healthcheck"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		healthcheckInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
		createSpec.Healthcheck = healthcheckInt
		for i, ps := range sl {

			hMapToStrVal := ps.(map[string]interface{})
			healthcheckInt[i] = &ves_io_schema_views.ObjectRefType{}

			if v, ok := hMapToStrVal["name"]; ok && !isIntfNil(v) {
				healthcheckInt[i].Name = v.(string)
			}

			if v, ok := hMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				healthcheckInt[i].Namespace = v.(string)
			}

			if v, ok := hMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				healthcheckInt[i].Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("loadbalancer_algorithm"); ok && !isIntfNil(v) {

		createSpec.LoadbalancerAlgorithm = ves_io_schema_cluster.LoadbalancerAlgorithm(ves_io_schema_cluster.LoadbalancerAlgorithm_value[v.(string)])

	}

	if v, ok := d.GetOk("origin_servers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		originServers := make([]*ves_io_schema_views_origin_pool.OriginServerType, len(sl))
		createSpec.OriginServers = originServers
		for i, set := range sl {
			originServers[i] = &ves_io_schema_views_origin_pool.OriginServerType{}

			originServersMapStrToI := set.(map[string]interface{})

			choiceTypeFound := false

			if v, ok := originServersMapStrToI["consul_service"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_ConsulService{}
				choiceInt.ConsulService = &ves_io_schema_views_origin_pool.OriginServerConsulService{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					networkChoiceTypeFound := false

					if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerConsulService_InsideNetwork{}
							networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
							choiceInt.ConsulService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerConsulService_OutsideNetwork{}
							networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
							choiceInt.ConsulService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["service_name"]; ok && !isIntfNil(v) {

						choiceInt.ConsulService.ServiceName = v.(string)
					}

					if v, ok := cs["site_locator"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						siteLocator := &ves_io_schema_views.SiteLocator{}
						choiceInt.ConsulService.SiteLocator = siteLocator
						for _, set := range sl {

							siteLocatorMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := siteLocatorMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_Site{}
								choiceInt.Site = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := siteLocatorMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_VirtualSite{}
								choiceInt.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["custom_endpoint_object"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_CustomEndpointObject{}
				choiceInt.CustomEndpointObject = &ves_io_schema_views_origin_pool.OriginServerCustomEndpoint{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["endpoint"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						endpoint := &ves_io_schema_views.ObjectRefType{}
						choiceInt.CustomEndpointObject.Endpoint = endpoint
						for _, set := range sl {

							endpointMapStrToI := set.(map[string]interface{})

							if w, ok := endpointMapStrToI["name"]; ok && !isIntfNil(w) {
								endpoint.Name = w.(string)
							}

							if w, ok := endpointMapStrToI["namespace"]; ok && !isIntfNil(w) {
								endpoint.Namespace = w.(string)
							}

							if w, ok := endpointMapStrToI["tenant"]; ok && !isIntfNil(w) {
								endpoint.Tenant = w.(string)
							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["k8s_service"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_K8SService{}
				choiceInt.K8SService = &ves_io_schema_views_origin_pool.OriginServerK8SService{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					networkChoiceTypeFound := false

					if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerK8SService_InsideNetwork{}
							networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
							choiceInt.K8SService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerK8SService_OutsideNetwork{}
							networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
							choiceInt.K8SService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["vk8s_networks"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerK8SService_Vk8SNetworks{}
							networkChoiceInt.Vk8SNetworks = &ves_io_schema.Empty{}
							choiceInt.K8SService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["service_name"]; ok && !isIntfNil(v) {

						choiceInt.K8SService.ServiceName = v.(string)
					}

					if v, ok := cs["site_locator"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						siteLocator := &ves_io_schema_views.SiteLocator{}
						choiceInt.K8SService.SiteLocator = siteLocator
						for _, set := range sl {

							siteLocatorMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := siteLocatorMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_Site{}
								choiceInt.Site = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := siteLocatorMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_VirtualSite{}
								choiceInt.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["private_ip"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_PrivateIp{}
				choiceInt.PrivateIp = &ves_io_schema_views_origin_pool.OriginServerPrivateIP{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ip"]; ok && !isIntfNil(v) {

						choiceInt.PrivateIp.Ip = v.(string)
					}

					networkChoiceTypeFound := false

					if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerPrivateIP_InsideNetwork{}
							networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
							choiceInt.PrivateIp.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerPrivateIP_OutsideNetwork{}
							networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
							choiceInt.PrivateIp.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["site_locator"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						siteLocator := &ves_io_schema_views.SiteLocator{}
						choiceInt.PrivateIp.SiteLocator = siteLocator
						for _, set := range sl {

							siteLocatorMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := siteLocatorMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_Site{}
								choiceInt.Site = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := siteLocatorMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_VirtualSite{}
								choiceInt.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["private_name"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_PrivateName{}
				choiceInt.PrivateName = &ves_io_schema_views_origin_pool.OriginServerPrivateName{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["dns_name"]; ok && !isIntfNil(v) {

						choiceInt.PrivateName.DnsName = v.(string)
					}

					networkChoiceTypeFound := false

					if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerPrivateName_InsideNetwork{}
							networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
							choiceInt.PrivateName.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerPrivateName_OutsideNetwork{}
							networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
							choiceInt.PrivateName.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["site_locator"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						siteLocator := &ves_io_schema_views.SiteLocator{}
						choiceInt.PrivateName.SiteLocator = siteLocator
						for _, set := range sl {

							siteLocatorMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := siteLocatorMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_Site{}
								choiceInt.Site = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := siteLocatorMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_VirtualSite{}
								choiceInt.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["public_ip"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_PublicIp{}
				choiceInt.PublicIp = &ves_io_schema_views_origin_pool.OriginServerPublicIP{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ip"]; ok && !isIntfNil(v) {

						choiceInt.PublicIp.Ip = v.(string)
					}

				}

			}

			if v, ok := originServersMapStrToI["public_name"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_PublicName{}
				choiceInt.PublicName = &ves_io_schema_views_origin_pool.OriginServerPublicName{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["dns_name"]; ok && !isIntfNil(v) {

						choiceInt.PublicName.DnsName = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("port"); ok && !isIntfNil(v) {

		createSpec.Port =
			uint32(v.(int))
	}

	tlsChoiceTypeFound := false

	if v, ok := d.GetOk("no_tls"); ok && !tlsChoiceTypeFound {

		tlsChoiceTypeFound = true

		if v.(bool) {
			tlsChoiceInt := &ves_io_schema_views_origin_pool.CreateSpecType_NoTls{}
			tlsChoiceInt.NoTls = &ves_io_schema.Empty{}
			createSpec.TlsChoice = tlsChoiceInt
		}

	}

	if v, ok := d.GetOk("use_tls"); ok && !tlsChoiceTypeFound {

		tlsChoiceTypeFound = true
		tlsChoiceInt := &ves_io_schema_views_origin_pool.CreateSpecType_UseTls{}
		tlsChoiceInt.UseTls = &ves_io_schema_views_origin_pool.UpstreamTlsParameters{}
		createSpec.TlsChoice = tlsChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			mtlsChoiceTypeFound := false

			if v, ok := cs["no_mtls"]; ok && !isIntfNil(v) && !mtlsChoiceTypeFound {

				mtlsChoiceTypeFound = true

				if v.(bool) {
					mtlsChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_NoMtls{}
					mtlsChoiceInt.NoMtls = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.MtlsChoice = mtlsChoiceInt
				}

			}

			if v, ok := cs["use_mtls"]; ok && !isIntfNil(v) && !mtlsChoiceTypeFound {

				mtlsChoiceTypeFound = true
				mtlsChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_UseMtls{}
				mtlsChoiceInt.UseMtls = &ves_io_schema_views_origin_pool.TlsCertificatesType{}
				tlsChoiceInt.UseTls.MtlsChoice = mtlsChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["tls_certificates"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						tlsCertificates := make([]*ves_io_schema.TlsCertificateType, len(sl))
						mtlsChoiceInt.UseMtls.TlsCertificates = tlsCertificates
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

				}

			}

			serverValidationChoiceTypeFound := false

			if v, ok := cs["skip_server_verification"]; ok && !isIntfNil(v) && !serverValidationChoiceTypeFound {

				serverValidationChoiceTypeFound = true

				if v.(bool) {
					serverValidationChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_SkipServerVerification{}
					serverValidationChoiceInt.SkipServerVerification = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.ServerValidationChoice = serverValidationChoiceInt
				}

			}

			if v, ok := cs["use_server_verification"]; ok && !isIntfNil(v) && !serverValidationChoiceTypeFound {

				serverValidationChoiceTypeFound = true
				serverValidationChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_UseServerVerification{}
				serverValidationChoiceInt.UseServerVerification = &ves_io_schema_views_origin_pool.UpstreamTlsValidationContext{}
				tlsChoiceInt.UseTls.ServerValidationChoice = serverValidationChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["trusted_ca_url"]; ok && !isIntfNil(v) {

						serverValidationChoiceInt.UseServerVerification.TrustedCaUrl = v.(string)
					}

				}

			}

			if v, ok := cs["volterra_trusted_ca"]; ok && !isIntfNil(v) && !serverValidationChoiceTypeFound {

				serverValidationChoiceTypeFound = true

				if v.(bool) {
					serverValidationChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_VolterraTrustedCa{}
					serverValidationChoiceInt.VolterraTrustedCa = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.ServerValidationChoice = serverValidationChoiceInt
				}

			}

			sniChoiceTypeFound := false

			if v, ok := cs["disable_sni"]; ok && !isIntfNil(v) && !sniChoiceTypeFound {

				sniChoiceTypeFound = true

				if v.(bool) {
					sniChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_DisableSni{}
					sniChoiceInt.DisableSni = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.SniChoice = sniChoiceInt
				}

			}

			if v, ok := cs["sni"]; ok && !isIntfNil(v) && !sniChoiceTypeFound {

				sniChoiceTypeFound = true
				sniChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_Sni{}

				tlsChoiceInt.UseTls.SniChoice = sniChoiceInt

				sniChoiceInt.Sni = v.(string)

			}

			if v, ok := cs["use_host_header_as_sni"]; ok && !isIntfNil(v) && !sniChoiceTypeFound {

				sniChoiceTypeFound = true

				if v.(bool) {
					sniChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_UseHostHeaderAsSni{}
					sniChoiceInt.UseHostHeaderAsSni = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.SniChoice = sniChoiceInt
				}

			}

			if v, ok := cs["tls_config"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				tlsConfig := &ves_io_schema_views.TlsConfig{}
				tlsChoiceInt.UseTls.TlsConfig = tlsConfig
				for _, set := range sl {

					tlsConfigMapStrToI := set.(map[string]interface{})

					choiceTypeFound := false

					if v, ok := tlsConfigMapStrToI["custom_security"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.TlsConfig_CustomSecurity{}
						choiceInt.CustomSecurity = &ves_io_schema_views.CustomCiphers{}
						tlsConfig.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["cipher_suites"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								choiceInt.CustomSecurity.CipherSuites = ls

							}

							if v, ok := cs["max_version"]; ok && !isIntfNil(v) {

								choiceInt.CustomSecurity.MaxVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

							}

							if v, ok := cs["min_version"]; ok && !isIntfNil(v) {

								choiceInt.CustomSecurity.MinVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

							}

						}

					}

					if v, ok := tlsConfigMapStrToI["default_security"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true

						if v.(bool) {
							choiceInt := &ves_io_schema_views.TlsConfig_DefaultSecurity{}
							choiceInt.DefaultSecurity = &ves_io_schema.Empty{}
							tlsConfig.Choice = choiceInt
						}

					}

					if v, ok := tlsConfigMapStrToI["low_security"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true

						if v.(bool) {
							choiceInt := &ves_io_schema_views.TlsConfig_LowSecurity{}
							choiceInt.LowSecurity = &ves_io_schema.Empty{}
							tlsConfig.Choice = choiceInt
						}

					}

					if v, ok := tlsConfigMapStrToI["medium_security"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true

						if v.(bool) {
							choiceInt := &ves_io_schema_views.TlsConfig_MediumSecurity{}
							choiceInt.MediumSecurity = &ves_io_schema.Empty{}
							tlsConfig.Choice = choiceInt
						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra OriginPool object with struct: %+v", createReq)

	createOriginPoolResp, err := client.CreateObject(context.Background(), ves_io_schema_views_origin_pool.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating OriginPool: %s", err)
	}
	d.SetId(createOriginPoolResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraOriginPoolRead(d, meta)
}

func resourceVolterraOriginPoolRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_origin_pool.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] OriginPool %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra OriginPool %q: %s", d.Id(), err)
	}
	return setOriginPoolFields(client, d, resp)
}

func setOriginPoolFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraOriginPoolUpdate updates OriginPool resource
func resourceVolterraOriginPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_origin_pool.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_origin_pool.ReplaceRequest{
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

	if v, ok := d.GetOk("endpoint_selection"); ok && !isIntfNil(v) {

		updateSpec.EndpointSelection = ves_io_schema_cluster.EndpointSelectionPolicy(ves_io_schema_cluster.EndpointSelectionPolicy_value[v.(string)])

	}

	if v, ok := d.GetOk("healthcheck"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		healthcheckInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
		updateSpec.Healthcheck = healthcheckInt
		for i, ps := range sl {

			hMapToStrVal := ps.(map[string]interface{})
			healthcheckInt[i] = &ves_io_schema_views.ObjectRefType{}

			if v, ok := hMapToStrVal["name"]; ok && !isIntfNil(v) {
				healthcheckInt[i].Name = v.(string)
			}

			if v, ok := hMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				healthcheckInt[i].Namespace = v.(string)
			}

			if v, ok := hMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				healthcheckInt[i].Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("loadbalancer_algorithm"); ok && !isIntfNil(v) {

		updateSpec.LoadbalancerAlgorithm = ves_io_schema_cluster.LoadbalancerAlgorithm(ves_io_schema_cluster.LoadbalancerAlgorithm_value[v.(string)])

	}

	if v, ok := d.GetOk("origin_servers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		originServers := make([]*ves_io_schema_views_origin_pool.OriginServerType, len(sl))
		updateSpec.OriginServers = originServers
		for i, set := range sl {
			originServers[i] = &ves_io_schema_views_origin_pool.OriginServerType{}

			originServersMapStrToI := set.(map[string]interface{})

			choiceTypeFound := false

			if v, ok := originServersMapStrToI["consul_service"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_ConsulService{}
				choiceInt.ConsulService = &ves_io_schema_views_origin_pool.OriginServerConsulService{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					networkChoiceTypeFound := false

					if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerConsulService_InsideNetwork{}
							networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
							choiceInt.ConsulService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerConsulService_OutsideNetwork{}
							networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
							choiceInt.ConsulService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["service_name"]; ok && !isIntfNil(v) {

						choiceInt.ConsulService.ServiceName = v.(string)
					}

					if v, ok := cs["site_locator"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						siteLocator := &ves_io_schema_views.SiteLocator{}
						choiceInt.ConsulService.SiteLocator = siteLocator
						for _, set := range sl {

							siteLocatorMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := siteLocatorMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_Site{}
								choiceInt.Site = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := siteLocatorMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_VirtualSite{}
								choiceInt.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["custom_endpoint_object"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_CustomEndpointObject{}
				choiceInt.CustomEndpointObject = &ves_io_schema_views_origin_pool.OriginServerCustomEndpoint{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["endpoint"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						endpoint := &ves_io_schema_views.ObjectRefType{}
						choiceInt.CustomEndpointObject.Endpoint = endpoint
						for _, set := range sl {

							endpointMapStrToI := set.(map[string]interface{})

							if w, ok := endpointMapStrToI["name"]; ok && !isIntfNil(w) {
								endpoint.Name = w.(string)
							}

							if w, ok := endpointMapStrToI["namespace"]; ok && !isIntfNil(w) {
								endpoint.Namespace = w.(string)
							}

							if w, ok := endpointMapStrToI["tenant"]; ok && !isIntfNil(w) {
								endpoint.Tenant = w.(string)
							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["k8s_service"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_K8SService{}
				choiceInt.K8SService = &ves_io_schema_views_origin_pool.OriginServerK8SService{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					networkChoiceTypeFound := false

					if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerK8SService_InsideNetwork{}
							networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
							choiceInt.K8SService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerK8SService_OutsideNetwork{}
							networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
							choiceInt.K8SService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["vk8s_networks"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerK8SService_Vk8SNetworks{}
							networkChoiceInt.Vk8SNetworks = &ves_io_schema.Empty{}
							choiceInt.K8SService.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["service_name"]; ok && !isIntfNil(v) {

						choiceInt.K8SService.ServiceName = v.(string)
					}

					if v, ok := cs["site_locator"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						siteLocator := &ves_io_schema_views.SiteLocator{}
						choiceInt.K8SService.SiteLocator = siteLocator
						for _, set := range sl {

							siteLocatorMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := siteLocatorMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_Site{}
								choiceInt.Site = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := siteLocatorMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_VirtualSite{}
								choiceInt.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["private_ip"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_PrivateIp{}
				choiceInt.PrivateIp = &ves_io_schema_views_origin_pool.OriginServerPrivateIP{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ip"]; ok && !isIntfNil(v) {

						choiceInt.PrivateIp.Ip = v.(string)
					}

					networkChoiceTypeFound := false

					if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerPrivateIP_InsideNetwork{}
							networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
							choiceInt.PrivateIp.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerPrivateIP_OutsideNetwork{}
							networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
							choiceInt.PrivateIp.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["site_locator"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						siteLocator := &ves_io_schema_views.SiteLocator{}
						choiceInt.PrivateIp.SiteLocator = siteLocator
						for _, set := range sl {

							siteLocatorMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := siteLocatorMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_Site{}
								choiceInt.Site = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := siteLocatorMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_VirtualSite{}
								choiceInt.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["private_name"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_PrivateName{}
				choiceInt.PrivateName = &ves_io_schema_views_origin_pool.OriginServerPrivateName{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["dns_name"]; ok && !isIntfNil(v) {

						choiceInt.PrivateName.DnsName = v.(string)
					}

					networkChoiceTypeFound := false

					if v, ok := cs["inside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerPrivateName_InsideNetwork{}
							networkChoiceInt.InsideNetwork = &ves_io_schema.Empty{}
							choiceInt.PrivateName.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["outside_network"]; ok && !isIntfNil(v) && !networkChoiceTypeFound {

						networkChoiceTypeFound = true

						if v.(bool) {
							networkChoiceInt := &ves_io_schema_views_origin_pool.OriginServerPrivateName_OutsideNetwork{}
							networkChoiceInt.OutsideNetwork = &ves_io_schema.Empty{}
							choiceInt.PrivateName.NetworkChoice = networkChoiceInt
						}

					}

					if v, ok := cs["site_locator"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						siteLocator := &ves_io_schema_views.SiteLocator{}
						choiceInt.PrivateName.SiteLocator = siteLocator
						for _, set := range sl {

							siteLocatorMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := siteLocatorMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_Site{}
								choiceInt.Site = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := siteLocatorMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.SiteLocator_VirtualSite{}
								choiceInt.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								siteLocator.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceInt.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := originServersMapStrToI["public_ip"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_PublicIp{}
				choiceInt.PublicIp = &ves_io_schema_views_origin_pool.OriginServerPublicIP{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ip"]; ok && !isIntfNil(v) {

						choiceInt.PublicIp.Ip = v.(string)
					}

				}

			}

			if v, ok := originServersMapStrToI["public_name"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_origin_pool.OriginServerType_PublicName{}
				choiceInt.PublicName = &ves_io_schema_views_origin_pool.OriginServerPublicName{}
				originServers[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["dns_name"]; ok && !isIntfNil(v) {

						choiceInt.PublicName.DnsName = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("port"); ok && !isIntfNil(v) {

		updateSpec.Port =
			uint32(v.(int))
	}

	tlsChoiceTypeFound := false

	if v, ok := d.GetOk("no_tls"); ok && !tlsChoiceTypeFound {

		tlsChoiceTypeFound = true

		if v.(bool) {
			tlsChoiceInt := &ves_io_schema_views_origin_pool.ReplaceSpecType_NoTls{}
			tlsChoiceInt.NoTls = &ves_io_schema.Empty{}
			updateSpec.TlsChoice = tlsChoiceInt
		}

	}

	if v, ok := d.GetOk("use_tls"); ok && !tlsChoiceTypeFound {

		tlsChoiceTypeFound = true
		tlsChoiceInt := &ves_io_schema_views_origin_pool.ReplaceSpecType_UseTls{}
		tlsChoiceInt.UseTls = &ves_io_schema_views_origin_pool.UpstreamTlsParameters{}
		updateSpec.TlsChoice = tlsChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			mtlsChoiceTypeFound := false

			if v, ok := cs["no_mtls"]; ok && !isIntfNil(v) && !mtlsChoiceTypeFound {

				mtlsChoiceTypeFound = true

				if v.(bool) {
					mtlsChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_NoMtls{}
					mtlsChoiceInt.NoMtls = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.MtlsChoice = mtlsChoiceInt
				}

			}

			if v, ok := cs["use_mtls"]; ok && !isIntfNil(v) && !mtlsChoiceTypeFound {

				mtlsChoiceTypeFound = true
				mtlsChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_UseMtls{}
				mtlsChoiceInt.UseMtls = &ves_io_schema_views_origin_pool.TlsCertificatesType{}
				tlsChoiceInt.UseTls.MtlsChoice = mtlsChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["tls_certificates"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						tlsCertificates := make([]*ves_io_schema.TlsCertificateType, len(sl))
						mtlsChoiceInt.UseMtls.TlsCertificates = tlsCertificates
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

				}

			}

			serverValidationChoiceTypeFound := false

			if v, ok := cs["skip_server_verification"]; ok && !isIntfNil(v) && !serverValidationChoiceTypeFound {

				serverValidationChoiceTypeFound = true

				if v.(bool) {
					serverValidationChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_SkipServerVerification{}
					serverValidationChoiceInt.SkipServerVerification = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.ServerValidationChoice = serverValidationChoiceInt
				}

			}

			if v, ok := cs["use_server_verification"]; ok && !isIntfNil(v) && !serverValidationChoiceTypeFound {

				serverValidationChoiceTypeFound = true
				serverValidationChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_UseServerVerification{}
				serverValidationChoiceInt.UseServerVerification = &ves_io_schema_views_origin_pool.UpstreamTlsValidationContext{}
				tlsChoiceInt.UseTls.ServerValidationChoice = serverValidationChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["trusted_ca_url"]; ok && !isIntfNil(v) {

						serverValidationChoiceInt.UseServerVerification.TrustedCaUrl = v.(string)
					}

				}

			}

			if v, ok := cs["volterra_trusted_ca"]; ok && !isIntfNil(v) && !serverValidationChoiceTypeFound {

				serverValidationChoiceTypeFound = true

				if v.(bool) {
					serverValidationChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_VolterraTrustedCa{}
					serverValidationChoiceInt.VolterraTrustedCa = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.ServerValidationChoice = serverValidationChoiceInt
				}

			}

			sniChoiceTypeFound := false

			if v, ok := cs["disable_sni"]; ok && !isIntfNil(v) && !sniChoiceTypeFound {

				sniChoiceTypeFound = true

				if v.(bool) {
					sniChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_DisableSni{}
					sniChoiceInt.DisableSni = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.SniChoice = sniChoiceInt
				}

			}

			if v, ok := cs["sni"]; ok && !isIntfNil(v) && !sniChoiceTypeFound {

				sniChoiceTypeFound = true
				sniChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_Sni{}

				tlsChoiceInt.UseTls.SniChoice = sniChoiceInt

				sniChoiceInt.Sni = v.(string)

			}

			if v, ok := cs["use_host_header_as_sni"]; ok && !isIntfNil(v) && !sniChoiceTypeFound {

				sniChoiceTypeFound = true

				if v.(bool) {
					sniChoiceInt := &ves_io_schema_views_origin_pool.UpstreamTlsParameters_UseHostHeaderAsSni{}
					sniChoiceInt.UseHostHeaderAsSni = &ves_io_schema.Empty{}
					tlsChoiceInt.UseTls.SniChoice = sniChoiceInt
				}

			}

			if v, ok := cs["tls_config"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				tlsConfig := &ves_io_schema_views.TlsConfig{}
				tlsChoiceInt.UseTls.TlsConfig = tlsConfig
				for _, set := range sl {

					tlsConfigMapStrToI := set.(map[string]interface{})

					choiceTypeFound := false

					if v, ok := tlsConfigMapStrToI["custom_security"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.TlsConfig_CustomSecurity{}
						choiceInt.CustomSecurity = &ves_io_schema_views.CustomCiphers{}
						tlsConfig.Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["cipher_suites"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								choiceInt.CustomSecurity.CipherSuites = ls

							}

							if v, ok := cs["max_version"]; ok && !isIntfNil(v) {

								choiceInt.CustomSecurity.MaxVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

							}

							if v, ok := cs["min_version"]; ok && !isIntfNil(v) {

								choiceInt.CustomSecurity.MinVersion = ves_io_schema.TlsProtocol(ves_io_schema.TlsProtocol_value[v.(string)])

							}

						}

					}

					if v, ok := tlsConfigMapStrToI["default_security"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true

						if v.(bool) {
							choiceInt := &ves_io_schema_views.TlsConfig_DefaultSecurity{}
							choiceInt.DefaultSecurity = &ves_io_schema.Empty{}
							tlsConfig.Choice = choiceInt
						}

					}

					if v, ok := tlsConfigMapStrToI["low_security"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true

						if v.(bool) {
							choiceInt := &ves_io_schema_views.TlsConfig_LowSecurity{}
							choiceInt.LowSecurity = &ves_io_schema.Empty{}
							tlsConfig.Choice = choiceInt
						}

					}

					if v, ok := tlsConfigMapStrToI["medium_security"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true

						if v.(bool) {
							choiceInt := &ves_io_schema_views.TlsConfig_MediumSecurity{}
							choiceInt.MediumSecurity = &ves_io_schema.Empty{}
							tlsConfig.Choice = choiceInt
						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra OriginPool obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_origin_pool.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating OriginPool: %s", err)
	}

	return resourceVolterraOriginPoolRead(d, meta)
}

func resourceVolterraOriginPoolDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_origin_pool.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] OriginPool %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra OriginPool before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra OriginPool obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_origin_pool.ObjectType, namespace, name)
}
