//
// Copyright (c) 2026 F5 Inc. All rights reserved.
//

package statemigration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSecureMeshSiteInstanceResourceV1() *schema.Resource {
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

			"bond_device_list": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bond_devices": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"devices": {
										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"active_backup": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"lacp": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"rate": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"link_polling_interval": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"link_up_delay": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"no_bond_devices": {

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

			"master_node_configuration": {

				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"public_ip": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"custom_network_config": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bgp_peer_address": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
						},

						"bgp_peer_address_v6": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
						},

						"bgp_router_id": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
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

						"default_interface_config": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"interface_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interfaces": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dc_cluster_group_connectivity_interface_disabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"dc_cluster_group_connectivity_interface_enabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"description": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"dedicated_interface": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"device": {
																Type:     schema.TypeString,
																Required: true,
															},

															"monitor": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{},
																},
															},

															"monitor_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"mtu": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"cluster": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"node": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"is_primary": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"not_primary": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},

												"dedicated_management_interface": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"device": {
																Type:     schema.TypeString,
																Required: true,
															},

															"mtu": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"cluster": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"node": {

																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"ethernet_interface": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"dhcp_server": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dhcp_networks": {

																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dns_address": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"same_as_dgw": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"dgw_address": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"first_address": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"last_address": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"network_prefix": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"network_prefix_allocator": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"name": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"namespace": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"tenant": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},

																					"pool_settings": {
																						Type:     schema.TypeString,
																						Required: true,
																					},

																					"pools": {

																						Type:     schema.TypeList,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"end_ip": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},

																								"exclude": {
																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},

																								"start_ip": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"dhcp_option82_tag": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"fixed_ip_map": {
																			Type:     schema.TypeMap,
																			Optional: true,
																		},

																		"automatic_from_end": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"automatic_from_start": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"interface_ip_map": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interface_ip_map": {
																						Type:     schema.TypeMap,
																						Optional: true,
																					},
																				},
																			},
																		},
																	},
																},
															},

															"static_ip": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"cluster_static_ip": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interface_ip_map": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"name": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"value": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Required: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"default_gw": {
																												Type:     schema.TypeString,
																												Optional: true,
																											},

																											"dns_server": {
																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"ip_address": {
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

																		"fleet_static_ip": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"default_gw": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"dns_server": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"network_prefix_allocator": {
																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"name": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"namespace": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"tenant": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"node_static_ip": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"default_gw": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"dns_server": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"ip_address": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																				},
																			},
																		},
																	},
																},
															},

															"device": {
																Type:     schema.TypeString,
																Required: true,
															},

															"ipv6_auto_config": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"host": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"router": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"network_prefix": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"stateful": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dhcp_networks": {

																									Type:     schema.TypeList,
																									Required: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"network_prefix": {

																												Type:     schema.TypeString,
																												Optional: true,
																											},

																											"network_prefix_allocator": {

																												Type:       schema.TypeList,
																												MaxItems:   1,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{

																														"name": {
																															Type:       schema.TypeString,
																															Optional:   true,
																															Deprecated: "This field is deprecated and will be removed in future release.",
																														},
																														"namespace": {
																															Type:       schema.TypeString,
																															Optional:   true,
																															Deprecated: "This field is deprecated and will be removed in future release.",
																														},
																														"tenant": {
																															Type:       schema.TypeString,
																															Optional:   true,
																															Deprecated: "This field is deprecated and will be removed in future release.",
																														},
																													},
																												},
																											},

																											"pool_settings": {
																												Type:     schema.TypeString,
																												Required: true,
																											},

																											"pools": {

																												Type:     schema.TypeList,
																												Optional: true,
																												Elem: &schema.Resource{
																													Schema: map[string]*schema.Schema{

																														"end_ip": {
																															Type:     schema.TypeString,
																															Optional: true,
																														},

																														"exclude": {
																															Type:       schema.TypeBool,
																															Optional:   true,
																															Deprecated: "This field is deprecated and will be removed in future release.",
																														},

																														"start_ip": {
																															Type:     schema.TypeString,
																															Optional: true,
																														},
																													},
																												},
																											},
																										},
																									},
																								},

																								"fixed_ip_map": {
																									Type:     schema.TypeMap,
																									Optional: true,
																								},

																								"automatic_from_end": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},

																								"automatic_from_start": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},

																								"interface_ip_map": {

																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"interface_ip_map": {
																												Type:     schema.TypeMap,
																												Optional: true,
																											},
																										},
																									},
																								},
																							},
																						},
																					},

																					"dns_config": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"configured_list": {

																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dns_list": {
																												Type: schema.TypeList,

																												Required: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},

																								"local_dns": {

																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"configured_address": {

																												Type:     schema.TypeString,
																												Optional: true,
																											},

																											"first_address": {

																												Type:     schema.TypeBool,
																												Optional: true,
																											},

																											"last_address": {

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
																	},
																},
															},

															"no_ipv6_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ipv6_address": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"cluster_static_ip": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interface_ip_map": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"name": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"value": {
																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Required: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"default_gw": {
																												Type:     schema.TypeString,
																												Optional: true,
																											},

																											"dns_server": {
																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"ip_address": {
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

																		"fleet_static_ip": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"default_gw": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"dns_server": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"network_prefix_allocator": {
																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"name": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"namespace": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"tenant": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"node_static_ip": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"default_gw": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"dns_server": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"ip_address": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																				},
																			},
																		},
																	},
																},
															},

															"monitor": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{},
																},
															},

															"monitor_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"mtu": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"inside_network": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"name": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"namespace": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"tenant": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"ip_fabric_network": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"segment_network": {

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

															"site_local_inside_network": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_local_network": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"srv6_network": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"name": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"namespace": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"tenant": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"storage_network": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"cluster": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"node": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"is_primary": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"not_primary": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"untagged": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"vlan_id": {

																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},

												"loopback_interface": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"dhcp_server": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dhcp_networks": {

																			Type:       schema.TypeList,
																			Required:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dns_address": {

																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"same_as_dgw": {

																						Type:       schema.TypeBool,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"dgw_address": {

																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"first_address": {

																						Type:       schema.TypeBool,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"last_address": {

																						Type:       schema.TypeBool,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"network_prefix": {

																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"network_prefix_allocator": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"name": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"namespace": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"tenant": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},

																					"pool_settings": {
																						Type:       schema.TypeString,
																						Required:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"pools": {

																						Type:       schema.TypeList,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"end_ip": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},

																								"exclude": {
																									Type:       schema.TypeBool,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},

																								"start_ip": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"dhcp_option82_tag": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"fixed_ip_map": {
																			Type:       schema.TypeMap,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"automatic_from_end": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"automatic_from_start": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"interface_ip_map": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interface_ip_map": {
																						Type:       schema.TypeMap,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},
																				},
																			},
																		},
																	},
																},
															},

															"static_ip": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"cluster_static_ip": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interface_ip_map": {

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
																								"value": {
																									Type:       schema.TypeList,
																									MaxItems:   1,
																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"default_gw": {
																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"dns_server": {
																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"ip_address": {
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

																		"fleet_static_ip": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"default_gw": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"dns_server": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"network_prefix_allocator": {
																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"name": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"namespace": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"tenant": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"node_static_ip": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"default_gw": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"dns_server": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"ip_address": {
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

															"device": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"no_ipv6_address": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"static_ipv6_address": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"cluster_static_ip": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"interface_ip_map": {

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
																								"value": {
																									Type:       schema.TypeList,
																									MaxItems:   1,
																									Required:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"default_gw": {
																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"dns_server": {
																												Type:       schema.TypeString,
																												Optional:   true,
																												Deprecated: "This field is deprecated and will be removed in future release.",
																											},

																											"ip_address": {
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

																		"fleet_static_ip": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"default_gw": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"dns_server": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"network_prefix_allocator": {
																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"name": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"namespace": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																								"tenant": {
																									Type:       schema.TypeString,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"node_static_ip": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"default_gw": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"dns_server": {
																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"ip_address": {
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

															"mtu": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"ip_fabric_network": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"site_local_inside_network": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"site_local_network": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"cluster": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"node": {

																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},

												"labels": {
													Type:     schema.TypeMap,
													Optional: true,
												},
											},
										},
									},
								},
							},
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

						"sm_connection_public_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"sm_connection_pvt_ip": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_sli_config": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"sli_config": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

									"dc_cluster_group_interface": {
										Type:       schema.TypeList,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},
												"namespace": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},
												"tenant": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},
											},
										},
									},

									"labels": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"nameserver": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"nameserver_v6": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"no_static_routes": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"static_routes": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"static_routes": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"ip_prefixes": {
																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"default_gateway": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"interface": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"name": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"namespace": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"tenant": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"ip_address": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"node_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"list": {

																			Type:     schema.TypeList,
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

																					"node": {
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

									"no_v6_static_routes": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"static_v6_routes": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"static_routes": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"ip_prefixes": {
																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"default_gateway": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"interface": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"name": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"namespace": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"tenant": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"ip_address": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"node_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"list": {

																			Type:     schema.TypeList,
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

																					"node": {
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

									"vip": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"vip_v6": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"default_config": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"slo_config": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

									"dc_cluster_group_interface": {
										Type:       schema.TypeList,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},
												"namespace": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},
												"tenant": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},
											},
										},
									},

									"labels": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"nameserver": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"nameserver_v6": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"no_static_routes": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"static_routes": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"static_routes": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"ip_prefixes": {
																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"default_gateway": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"interface": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"name": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"namespace": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"tenant": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"ip_address": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"node_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"list": {

																			Type:     schema.TypeList,
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

																					"node": {
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

									"no_v6_static_routes": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"static_v6_routes": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"static_routes": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attrs": {
																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"ip_prefixes": {
																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"default_gateway": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"interface": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"name": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"namespace": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																		"tenant": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"ip_address": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"node_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"list": {

																			Type:     schema.TypeList,
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

																					"node": {
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

									"vip": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"vip_v6": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"tunnel_dead_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"vip_vrrp_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"default_network_config": {

				Type:     schema.TypeBool,
				Optional: true,
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

			"volterra_certified_hw": {
				Type:     schema.TypeString,
				Required: true,
			},

			"worker_nodes": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func ResourceSecureMeshSiteInstanceStateUpgradeV1(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
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
		if boolVal, ok := value.(bool); ok {
			aa["perf_mode_l7_enhanced"] = []interface{}{map[string]interface{}{
				"jumbo_disabled": boolVal,
			}}
		}
	}
	return rawState, nil
}
