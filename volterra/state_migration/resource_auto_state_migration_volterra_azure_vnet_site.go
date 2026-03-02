//
// Copyright (c) 2026 F5 Inc. All rights reserved.
//

package statemigration

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAzureVnetSiteInstanceResourceV1() *schema.Resource {
	return &schema.Resource{
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

			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"admin_password": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
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

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"decryption_provider": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"location": {
										Type:     schema.TypeString,
										Required: true,
									},

									"store_provider": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"clear_secret_info": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"provider": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"url": {
										Type:     schema.TypeString,
										Required: true,
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

			"block_all_services": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"blocked_services": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"blocked_sevice": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"ssh": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"web_user_interface": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"network_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"default_blocked_services": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"coordinates": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"latitude": {
							Type:     schema.TypeFloat,
							Optional: true,
						},

						"longitude": {
							Type:     schema.TypeFloat,
							Optional: true,
						},
					},
				},
			},

			"custom_dns": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"inside_nameserver": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"inside_nameserver_v6": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"outside_nameserver": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"outside_nameserver_v6": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"azure_cred": {

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

			"disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"disable_encryption": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_encryption": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"disk_encryption_set_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						"resource_group": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"kubernetes_upgrade_drain": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"disable_upgrade_drain": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_upgrade_drain": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"drain_max_unavailable_node_count": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"drain_max_unavailable_node_percentage": {

										Type:       schema.TypeInt,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"drain_node_timeout": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"disable_vega_upgrade_mode": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"enable_vega_upgrade_mode": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},
								},
							},
						},
					},
				},
			},

			"log_receiver": {

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

			"logs_streaming_disabled": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"machine_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"offline_survivability_mode": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_offline_survivability_mode": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_offline_survivability_mode": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"os": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"default_os_version": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"operating_system_version": {

							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"alternate_region": {

				Type:     schema.TypeString,
				Optional: true,
			},

			"azure_region": {

				Type:     schema.TypeString,
				Optional: true,
			},

			"resource_group": {
				Type:     schema.TypeString,
				Required: true,
			},

			"ingress_egress_gw": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accelerated_networking": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"az_nodes": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"azure_az": {
										Type:     schema.TypeString,
										Required: true,
									},

									"disk_size": {
										Type:       schema.TypeInt,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"inside_subnet": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_resource_grp": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"vnet_resource_group": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Required: true,
															},

															"ipv6": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"outside_subnet": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_resource_grp": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"vnet_resource_group": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Required: true,
															},

															"ipv6": {
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

						"azure_certified_hw": {
							Type:     schema.TypeString,
							Required: true,
						},

						"dc_cluster_group_inside_vn": {

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

						"dc_cluster_group_outside_vn": {

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

						"no_dc_cluster_group": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_forward_proxy_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"forward_proxy_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"forward_proxy_allow_all": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_forward_proxy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"global_network_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_network_connections": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sli_to_global_dr": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {
																Type:     schema.TypeList,
																MaxItems: 1,
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

												"slo_to_global_dr": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {
																Type:     schema.TypeList,
																MaxItems: 1,
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

												"disable_forward_proxy": {

													Type:       schema.TypeBool,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"enable_forward_proxy": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connection_timeout": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"max_connect_attempts": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"no_interception": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"tls_intercept": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"enable_for_all_domains": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"policy": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interception_rules": {

																						Type:       schema.TypeList,
																						Required:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"domain_match": {

																									Type:       schema.TypeList,
																									MaxItems:   1,
																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"exact_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"regex_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"suffix_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},
																										},
																									},
																								},

																								"disable_interception": {

																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},

																								"enable_interception": {

																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"custom_certificate": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"certificate_url": {
																						Type:       schema.TypeString,
																						Required:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"description": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"custom_hash_algorithms": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"hash_algorithms": {
																									Type: schema.TypeList,

																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"disable_ocsp_stapling": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{},
																						},
																					},

																					"use_system_defaults": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{},
																						},
																					},

																					"private_key": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Required:   true,
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
																				},
																			},
																		},

																		"volterra_certificate": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"trusted_ca_url": {

																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"volterra_trusted_ca": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"white_listed_ports": {
																Type: schema.TypeList,

																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},

															"white_listed_prefixes": {
																Type: schema.TypeList,

																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
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
								},
							},
						},

						"no_global_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"hub": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"express_route_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"express_route_enabled": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"auto_asn": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"custom_asn": {

													Type:     schema.TypeInt,
													Optional: true,
												},

												"connections": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"metadata": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"description": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"disable": {
																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"circuit_id": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"other_subscription": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"authorized_key": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
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

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"decryption_provider": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},

																								"location": {
																									Type:     schema.TypeString,
																									Required: true,
																								},

																								"store_provider": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"clear_secret_info": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"provider": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},

																								"url": {
																									Type:     schema.TypeString,
																									Required: true,
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

																		"circuit_id": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"weight": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},

												"site_registration_over_express_route": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"cloudlink_network_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"site_registration_over_internet": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"gateway_subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"auto": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"subnet_resource_grp": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"vnet_resource_group": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"subnet_param": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"ipv6": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"route_server_subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"auto": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"subnet_resource_grp": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"vnet_resource_group": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"subnet_param": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"ipv6": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"sku_ergw1az": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"sku_ergw2az": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"sku_high_perf": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"sku_standard": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"advertise_to_route_server": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"do_not_advertise_to_route_server": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"spoke_vnets": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"labels": {
													Type:     schema.TypeMap,
													Optional: true,
												},

												"auto": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"manual": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"vnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"resource_group": {
																Type:     schema.TypeString,
																Required: true,
															},

															"f5_orchestrated_routing": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"manual_routing": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"vnet_name": {
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
							},
						},

						"not_hub": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"inside_static_routes": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"nexthop": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"interface": {
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

																		"nexthop_address": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"ipv6": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"subnets": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"ipv6": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
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

												"simple_static_route": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"no_inside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_enhanced_firewall_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enhanced_firewall_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"active_network_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"no_network_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_outside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"outside_static_routes": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"nexthop": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"interface": {
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

																		"nexthop_address": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"ipv6": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"subnets": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"ipv6": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
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

												"simple_static_route": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"performance_enhancement_mode": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"perf_mode_l3_enhanced": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"jumbo": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"no_jumbo": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"perf_mode_l7_enhanced": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"jumbo_disabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"jumbo_enabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"sm_connection_public_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"sm_connection_pvt_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"ingress_egress_gw_ar": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accelerated_networking": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"azure_certified_hw": {
							Type:     schema.TypeString,
							Required: true,
						},

						"dc_cluster_group_inside_vn": {

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

						"dc_cluster_group_outside_vn": {

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

						"no_dc_cluster_group": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_forward_proxy_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"forward_proxy_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"forward_proxy_allow_all": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_forward_proxy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"global_network_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_network_connections": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sli_to_global_dr": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {
																Type:     schema.TypeList,
																MaxItems: 1,
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

												"slo_to_global_dr": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {
																Type:     schema.TypeList,
																MaxItems: 1,
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

												"disable_forward_proxy": {

													Type:       schema.TypeBool,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"enable_forward_proxy": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connection_timeout": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"max_connect_attempts": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"no_interception": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"tls_intercept": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"enable_for_all_domains": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"policy": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interception_rules": {

																						Type:       schema.TypeList,
																						Required:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"domain_match": {

																									Type:       schema.TypeList,
																									MaxItems:   1,
																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"exact_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"regex_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"suffix_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},
																										},
																									},
																								},

																								"disable_interception": {

																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},

																								"enable_interception": {

																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"custom_certificate": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"certificate_url": {
																						Type:       schema.TypeString,
																						Required:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"description": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"custom_hash_algorithms": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"hash_algorithms": {
																									Type: schema.TypeList,

																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"disable_ocsp_stapling": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{},
																						},
																					},

																					"use_system_defaults": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{},
																						},
																					},

																					"private_key": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Required:   true,
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
																				},
																			},
																		},

																		"volterra_certificate": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"trusted_ca_url": {

																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"volterra_trusted_ca": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"white_listed_ports": {
																Type: schema.TypeList,

																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},

															"white_listed_prefixes": {
																Type: schema.TypeList,

																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
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
								},
							},
						},

						"no_global_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"hub": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"express_route_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"express_route_enabled": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"auto_asn": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"custom_asn": {

													Type:     schema.TypeInt,
													Optional: true,
												},

												"connections": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"metadata": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"description": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"disable": {
																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"circuit_id": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"other_subscription": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"authorized_key": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
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

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"decryption_provider": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},

																								"location": {
																									Type:     schema.TypeString,
																									Required: true,
																								},

																								"store_provider": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"clear_secret_info": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"provider": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},

																								"url": {
																									Type:     schema.TypeString,
																									Required: true,
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

																		"circuit_id": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"weight": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},

												"site_registration_over_express_route": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"cloudlink_network_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"site_registration_over_internet": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"gateway_subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"auto": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"subnet_resource_grp": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"vnet_resource_group": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"subnet_param": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"ipv6": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"route_server_subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"auto": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"subnet_resource_grp": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"vnet_resource_group": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"subnet_param": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"ipv6": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"sku_ergw1az": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"sku_ergw2az": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"sku_high_perf": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"sku_standard": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"advertise_to_route_server": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"do_not_advertise_to_route_server": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"spoke_vnets": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"labels": {
													Type:     schema.TypeMap,
													Optional: true,
												},

												"auto": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"manual": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"vnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"resource_group": {
																Type:     schema.TypeString,
																Required: true,
															},

															"f5_orchestrated_routing": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"manual_routing": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"vnet_name": {
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
							},
						},

						"not_hub": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"inside_static_routes": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"nexthop": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"interface": {
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

																		"nexthop_address": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"ipv6": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"subnets": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"ipv6": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
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

												"simple_static_route": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"no_inside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_enhanced_firewall_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enhanced_firewall_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"active_network_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"no_network_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"node": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fault_domain": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"inside_subnet": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_resource_grp": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"vnet_resource_group": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Required: true,
															},

															"ipv6": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"node_number": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"outside_subnet": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_resource_grp": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"vnet_resource_group": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Required: true,
															},

															"ipv6": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"update_domain": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"no_outside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"outside_static_routes": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"nexthop": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"interface": {
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

																		"nexthop_address": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"ipv6": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"subnets": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"ipv6": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
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

												"simple_static_route": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"performance_enhancement_mode": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"perf_mode_l3_enhanced": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"jumbo": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"no_jumbo": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"perf_mode_l7_enhanced": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"jumbo_disabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"jumbo_enabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"sm_connection_public_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"sm_connection_pvt_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"ingress_gw": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accelerated_networking": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"az_nodes": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"azure_az": {
										Type:     schema.TypeString,
										Required: true,
									},

									"disk_size": {
										Type:       schema.TypeInt,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"local_subnet": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_resource_grp": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"vnet_resource_group": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Required: true,
															},

															"ipv6": {
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

						"azure_certified_hw": {
							Type:     schema.TypeString,
							Required: true,
						},

						"performance_enhancement_mode": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"perf_mode_l3_enhanced": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"jumbo": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"no_jumbo": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"perf_mode_l7_enhanced": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"jumbo_disabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"jumbo_enabled": {

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
				},
			},

			"ingress_gw_ar": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accelerated_networking": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"azure_certified_hw": {
							Type:     schema.TypeString,
							Required: true,
						},

						"node": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fault_domain": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"local_subnet": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_resource_grp": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"vnet_resource_group": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Required: true,
															},

															"ipv6": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"node_number": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"update_domain": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"performance_enhancement_mode": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"perf_mode_l3_enhanced": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"jumbo": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"no_jumbo": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"perf_mode_l7_enhanced": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"jumbo_disabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"jumbo_enabled": {

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
				},
			},

			"voltstack_cluster": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accelerated_networking": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"az_nodes": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"azure_az": {
										Type:     schema.TypeString,
										Required: true,
									},

									"disk_size": {
										Type:       schema.TypeInt,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"local_subnet": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_resource_grp": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"vnet_resource_group": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Required: true,
															},

															"ipv6": {
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

						"azure_certified_hw": {
							Type:     schema.TypeString,
							Required: true,
						},

						"dc_cluster_group": {

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

						"no_dc_cluster_group": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_forward_proxy_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"forward_proxy_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"forward_proxy_allow_all": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_forward_proxy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"global_network_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_network_connections": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sli_to_global_dr": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {
																Type:     schema.TypeList,
																MaxItems: 1,
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

												"slo_to_global_dr": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {
																Type:     schema.TypeList,
																MaxItems: 1,
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

												"disable_forward_proxy": {

													Type:       schema.TypeBool,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"enable_forward_proxy": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connection_timeout": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"max_connect_attempts": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"no_interception": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"tls_intercept": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"enable_for_all_domains": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"policy": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interception_rules": {

																						Type:       schema.TypeList,
																						Required:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"domain_match": {

																									Type:       schema.TypeList,
																									MaxItems:   1,
																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"exact_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"regex_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"suffix_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},
																										},
																									},
																								},

																								"disable_interception": {

																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},

																								"enable_interception": {

																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"custom_certificate": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"certificate_url": {
																						Type:       schema.TypeString,
																						Required:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"description": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"custom_hash_algorithms": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"hash_algorithms": {
																									Type: schema.TypeList,

																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"disable_ocsp_stapling": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{},
																						},
																					},

																					"use_system_defaults": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{},
																						},
																					},

																					"private_key": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Required:   true,
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
																				},
																			},
																		},

																		"volterra_certificate": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"trusted_ca_url": {

																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"volterra_trusted_ca": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"white_listed_ports": {
																Type: schema.TypeList,

																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},

															"white_listed_prefixes": {
																Type: schema.TypeList,

																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
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
								},
							},
						},

						"no_global_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"k8s_cluster": {

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

						"no_k8s_cluster": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_enhanced_firewall_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enhanced_firewall_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"active_network_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"no_network_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_outside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"outside_static_routes": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"nexthop": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"interface": {
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

																		"nexthop_address": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"ipv6": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"subnets": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"ipv6": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
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

												"simple_static_route": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"sm_connection_public_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"sm_connection_pvt_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_storage": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"storage_class_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"storage_classes": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_storage_class": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"storage_class_name": {
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
				},
			},

			"voltstack_cluster_ar": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"accelerated_networking": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"azure_certified_hw": {
							Type:     schema.TypeString,
							Required: true,
						},

						"dc_cluster_group": {

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

						"no_dc_cluster_group": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_forward_proxy_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"forward_proxy_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"forward_proxy_allow_all": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_forward_proxy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"global_network_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_network_connections": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sli_to_global_dr": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {
																Type:     schema.TypeList,
																MaxItems: 1,
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

												"slo_to_global_dr": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {
																Type:     schema.TypeList,
																MaxItems: 1,
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

												"disable_forward_proxy": {

													Type:       schema.TypeBool,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"enable_forward_proxy": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connection_timeout": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"max_connect_attempts": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"no_interception": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"tls_intercept": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"enable_for_all_domains": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"policy": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interception_rules": {

																						Type:       schema.TypeList,
																						Required:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"domain_match": {

																									Type:       schema.TypeList,
																									MaxItems:   1,
																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"exact_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"regex_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"suffix_value": {

																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},
																										},
																									},
																								},

																								"disable_interception": {

																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},

																								"enable_interception": {

																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"custom_certificate": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"certificate_url": {
																						Type:       schema.TypeString,
																						Required:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"description": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"custom_hash_algorithms": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"hash_algorithms": {
																									Type: schema.TypeList,

																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"disable_ocsp_stapling": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{},
																						},
																					},

																					"use_system_defaults": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{},
																						},
																					},

																					"private_key": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Required:   true,
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
																				},
																			},
																		},

																		"volterra_certificate": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"trusted_ca_url": {

																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"volterra_trusted_ca": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"white_listed_ports": {
																Type: schema.TypeList,

																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},

															"white_listed_prefixes": {
																Type: schema.TypeList,

																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
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
								},
							},
						},

						"no_global_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"k8s_cluster": {

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

						"no_k8s_cluster": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_enhanced_firewall_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enhanced_firewall_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"active_network_policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_policies": {
										Type:     schema.TypeList,
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
								},
							},
						},

						"no_network_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"node": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fault_domain": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"local_subnet": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_resource_grp": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"vnet_resource_group": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"subnet_name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
																Type:     schema.TypeString,
																Required: true,
															},

															"ipv6": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"node_number": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"update_domain": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"no_outside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"outside_static_routes": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"nexthop": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"interface": {
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

																		"nexthop_address": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"ipv6": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"addr": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"type": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"subnets": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"ipv6": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"plen": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"prefix": {
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

												"simple_static_route": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"sm_connection_public_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"sm_connection_pvt_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_storage": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"storage_class_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"storage_classes": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_storage_class": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"storage_class_name": {
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
				},
			},

			"ssh_key": {
				Type:     schema.TypeString,
				Required: true,
			},

			"sw": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"default_sw_version": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"volterra_software_version": {

							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"vnet": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"existing_vnet": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"resource_group": {
										Type:     schema.TypeString,
										Required: true,
									},

									"f5_orchestrated_routing": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"manual_routing": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"vnet_name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"new_vnet": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"autogenerate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"name": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"primary_ipv4": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"no_worker_nodes": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"nodes_per_az": {

				Type:     schema.TypeInt,
				Optional: true,
			},

			"total_nodes": {

				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func ResourceAzureVnetSiteInstanceStateUpgradeV1(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	advancedOptions, ok := rawState["performance_enhancement_mode"].([]interface{})
	if !ok {
		return rawState, nil
	}

	for _, va := range advancedOptions {
		aa, ok := va.(map[string]interface{})
		if !ok {
			continue
		}

		value, ok := aa["perf_mode_l7_enhanced"]
		if ok && value != nil && reflect.TypeOf(value).Kind() == reflect.Bool {
			aa["perf_mode_l7_enhanced"] = []interface{}{map[string]interface{}{
				"jumbo_enabled": true,
			}}
		}
	}
	return rawState, nil
}
