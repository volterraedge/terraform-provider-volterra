package statemigration

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceOriginPoolInstanceResourceV1() *schema.Resource {
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

			"advanced_options": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"circuit_breaker": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connection_limit": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_requests": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"pending_requests": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"priority": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"retries": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"default_circuit_breaker": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"disable_circuit_breaker": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"connection_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"http_idle_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"auto_http_config": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"http1_config": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"header_transformation": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_header_transformation": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"legacy_header_transformation": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"preserve_case_header_transformation": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"proper_case_header_transformation": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"http2_options": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"disable_lb_source_ip_persistance": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_lb_source_ip_persistance": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"disable_outlier_detection": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"outlier_detection": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"base_ejection_time": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"consecutive_5xx": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"consecutive_gateway_failure": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"interval": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_ejection_percent": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"no_panic_threshold": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"panic_threshold": {

							Type:     schema.TypeInt,
							Optional: true,
						},

						"disable_proxy_protocol": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"proxy_protocol_v1": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"proxy_protocol_v2": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"disable_subsets": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_subsets": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"endpoint_subsets": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"keys": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"any_endpoint": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"default_subset": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_subset": {
													Type:     schema.TypeMap,
													Optional: true,
												},
											},
										},
									},

									"fail_request": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"endpoint_selection": {
				Type:     schema.TypeString,
				Required: true,
			},

			"health_check_port": {

				Type:     schema.TypeInt,
				Optional: true,
			},

			"same_as_endpoint_port": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"healthcheck": {

				Type:     schema.TypeList,
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

			"loadbalancer_algorithm": {
				Type:     schema.TypeString,
				Required: true,
			},

			"origin_servers": {

				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cbip_service": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"service_name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"consul_service": {

							Type:     schema.TypeList,
							MaxItems: 1,
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
										Required: true,
									},

									"site_locator": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"site": {

													Type:     schema.TypeList,
													MaxItems: 1,
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

													Type:     schema.TypeList,
													MaxItems: 1,
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

									"snat_pool": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"no_snat_pool": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"snat_pool": {

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
											},
										},
									},
								},
							},
						},

						"custom_endpoint_object": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"endpoint": {

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
								},
							},
						},

						"k8s_service": {

							Type:     schema.TypeList,
							MaxItems: 1,
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

									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"service_name": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"site_locator": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"site": {

													Type:     schema.TypeList,
													MaxItems: 1,
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

													Type:     schema.TypeList,
													MaxItems: 1,
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

									"snat_pool": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"no_snat_pool": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"snat_pool": {

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
											},
										},
									},
								},
							},
						},

						"private_ip": {

							Type:     schema.TypeList,
							MaxItems: 1,
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

									"segment": {

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

									"ip": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"ipv6": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"site_locator": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"site": {

													Type:     schema.TypeList,
													MaxItems: 1,
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

													Type:     schema.TypeList,
													MaxItems: 1,
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

									"snat_pool": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"no_snat_pool": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"snat_pool": {

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
											},
										},
									},
								},
							},
						},

						"private_name": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"inside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"outside_network": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"segment": {

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

									"refresh_interval": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"site_locator": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"site": {

													Type:     schema.TypeList,
													MaxItems: 1,
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

													Type:     schema.TypeList,
													MaxItems: 1,
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

									"snat_pool": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"no_snat_pool": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"snat_pool": {

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
											},
										},
									},
								},
							},
						},

						"public_ip": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"ipv6": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"public_name": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"refresh_interval": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"vn_private_ip": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"virtual_network": {

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

									"ip": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"ipv6": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"vn_private_name": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"private_network": {

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

			"automatic_port": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"lb_port": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"port": {

				Type:     schema.TypeInt,
				Optional: true,
			},

			"no_tls": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"use_tls": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"default_session_key_caching": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"disable_session_key_caching": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"max_session_keys": {

							Type:     schema.TypeInt,
							Optional: true,
						},

						"no_mtls": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"use_mtls": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"tls_certificates": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate_url": {
													Type:     schema.TypeString,
													Required: true,
												},

												"description": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"custom_hash_algorithms": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"hash_algorithms": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"disable_ocsp_stapling": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
													},
												},

												"use_system_defaults": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
													},
												},

												"private_key": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

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
														},
													},
												},
											},
										},
									},
								},
							},
						},

						"use_mtls_obj": {

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

						"skip_server_verification": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"use_server_verification": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"trusted_ca": {

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

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_security": {

										Type:     schema.TypeList,
										MaxItems: 1,
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

			"upstream_conn_pool_reuse_type": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"disable_conn_pool_reuse": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_conn_pool_reuse": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func ResourceOriginPoolInstanceStateUpgradeV1(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	advancedOptions, ok := rawState["advanced_options"].([]interface{})
	if !ok {
		return rawState, nil
	}

	for _, va := range advancedOptions {
		aa, ok := va.(map[string]interface{})
		if !ok {
			continue
		}

		value, ok := aa["http1_config"]
		if ok && value != nil && reflect.TypeOf(value).Kind() == reflect.Bool {
			aa["http1_config"] = []interface{}{map[string]interface{}{
				"header_transformation": []interface{}{map[string]interface{}{
					"legacy_header_transformation": true,
				}},
			}}
		}
	}
	return rawState, nil
}
