//
// Copyright (c) 2026 F5 Inc. All rights reserved.
//

package statemigration

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFleetInstanceResourceV1() *schema.Resource {
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

			"blocked_services": {

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

			"dc_cluster_group_inside": {

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

			"enable_default_fleet_config_download": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"fleet_label": {
				Type:     schema.TypeString,
				Required: true,
			},

			"disable_gpu": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_gpu": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_vgpu": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"feature_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						"server_address": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"server_port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"inside_virtual_network": {
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

			"default_config": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"device_list": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"devices": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_device": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": {
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

												"use": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"owner": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
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

			"network_connectors": {
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

			"network_firewall": {
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

			"operating_system_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"outside_virtual_network": {
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

			"default_sriov_interface": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"sriov_interfaces": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"sriov_interface": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"number_of_vfio_vfs": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"number_of_vfs": {
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"default_storage_class": {

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

									"advanced_storage_parameters": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"allow_volume_expansion": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"default_storage_class": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"custom_storage": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"yaml": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"hpe_storage": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow_mutations": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"allow_overrides": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"dedupe_enabled": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"description": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"destroy_on_delete": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"encrypted": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"folder": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"limit_iops": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"limit_mbps": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"performance_policy": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"pool": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"protection_template": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"secret_name": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"secret_namespace": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"sync_on_detach": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"thick": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"netapp_trident": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"selector": {
													Type:     schema.TypeMap,
													Optional: true,
												},

												"storage_pools": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"pure_service_orchestrator": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"backend": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"bandwidth_limit": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"iops_limit": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"reclaim_policy": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"storage_class_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"storage_device": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"no_storage_device": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"storage_device_list": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"storage_devices": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_advanced_parameters": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"custom_storage": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"hpe_storage": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"api_server_port": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"csi_version": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"iscsi_chap_password": {

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

												"iscsi_chap_user": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"log_level": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"password": {

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

												"storage_server_ip_address": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"storage_server_name": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"username": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"netapp_trident": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"netapp_backend_ontap_nas": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"auto_export_cidrs": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv6_prefixes": {
																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"prefixes": {
																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"auto_export_policy": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"backend_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"client_certificate": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"client_private_key": {

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

															"data_lif_dns_name": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"data_lif_ip": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"limit_aggregate_usage": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"limit_volume_size": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"management_lif_dns_name": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"management_lif_ip": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"nfs_mount_options": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"password": {

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

															"region": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"storage": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"labels": {
																			Type:     schema.TypeMap,
																			Optional: true,
																		},

																		"volume_defaults": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"encryption": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"export_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"adaptive_qos_policy": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"no_qos": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"qos_policy": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"security_style": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"snapshot_dir": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"snapshot_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"snapshot_reserve": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"space_reserve": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"split_on_clone": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"tiering_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"unix_permissions": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"zone": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"storage_driver_name": {
																Type:     schema.TypeString,
																Required: true,
															},

															"storage_prefix": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"svm": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"trusted_ca_certificate": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"username": {
																Type:     schema.TypeString,
																Required: true,
															},

															"volume_defaults": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"encryption": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"export_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"adaptive_qos_policy": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"no_qos": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"qos_policy": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"security_style": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"snapshot_dir": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"snapshot_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"snapshot_reserve": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"space_reserve": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"split_on_clone": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"tiering_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"unix_permissions": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"netapp_backend_ontap_san": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"no_chap": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"use_chap": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"chap_initiator_secret": {

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

																		"chap_target_initiator_secret": {

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

																		"chap_target_username": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"chap_username": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"client_certificate": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"client_private_key": {

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

															"data_lif_dns_name": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"data_lif_ip": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"igroup_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"limit_aggregate_usage": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"limit_volume_size": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"management_lif_dns_name": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"management_lif_ip": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"password": {

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

															"region": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"storage": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"labels": {
																			Type:     schema.TypeMap,
																			Optional: true,
																		},

																		"volume_defaults": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"encryption": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"export_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"adaptive_qos_policy": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"no_qos": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"qos_policy": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"security_style": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"snapshot_dir": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"snapshot_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"snapshot_reserve": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"space_reserve": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"split_on_clone": {
																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"tiering_policy": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"unix_permissions": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"zone": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"storage_driver_name": {
																Type:     schema.TypeString,
																Required: true,
															},

															"storage_prefix": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"svm": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"trusted_ca_certificate": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"username": {
																Type:     schema.TypeString,
																Required: true,
															},

															"volume_defaults": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"encryption": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"export_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"adaptive_qos_policy": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"no_qos": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"qos_policy": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"security_style": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"snapshot_dir": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"snapshot_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"snapshot_reserve": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"space_reserve": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"split_on_clone": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"tiering_policy": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"unix_permissions": {
																			Type:     schema.TypeInt,
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

									"pure_service_orchestrator": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"arrays": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"flash_array": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"default_fs_opt": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"default_fs_type": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"default_mount_opts": {
																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"disable_preempt_attachments": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"flash_arrays": {

																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"api_token": {

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

																					"labels": {
																						Type:     schema.TypeMap,
																						Optional: true,
																					},

																					"mgmt_dns_name": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"mgmt_ip": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"iscsi_login_timeout": {
																			Type:     schema.TypeInt,
																			Required: true,
																		},

																		"san_type": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"flash_blade": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"enable_snapshot_directory": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"export_rules": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"flash_blades": {

																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"api_token": {

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

																					"lables": {
																						Type:     schema.TypeMap,
																						Optional: true,
																					},

																					"mgmt_dns_name": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"mgmt_ip": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"nfs_endpoint_dns_name": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"nfs_endpoint_ip": {

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

												"cluster_id": {
													Type:     schema.TypeString,
													Required: true,
												},

												"enable_storage_topology": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"enable_strict_topology": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"storage_device": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"no_storage_interfaces": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"storage_interface_list": {

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

			"no_storage_static_routes": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"storage_static_routes": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"storage_routes": {

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
					},
				},
			},

			"allow_all_usb": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"deny_all_usb": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"usb_policy": {

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

			"disable_vm": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_vm": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},

			"volterra_software_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func ResourceFleetInstanceStateUpgradeV1(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
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
