package statemigration

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceHttpLoadbalancerInstanceResourceV1() *schema.Resource {
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

			"cname": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"add_location": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"advertise_custom": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"advertise_where": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advertise_on_public": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"public_ip": {

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

									"site": {

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

												"network": {
													Type:     schema.TypeString,
													Required: true,
												},

												"site": {

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

									"site_segment": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip": {
													Type:     schema.TypeString,
													Required: true,
												},

												"ipv6": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"segment": {

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

												"site": {

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

									"virtual_network": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_v6_vip": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"specific_v6_vip": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"default_vip": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"specific_vip": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"virtual_network": {

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

									"virtual_site": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"network": {
													Type:     schema.TypeString,
													Required: true,
												},

												"virtual_site": {

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

									"virtual_site_segment": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip": {
													Type:     schema.TypeString,
													Required: true,
												},

												"ipv6": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"segment": {

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

												"virtual_site": {

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

									"virtual_site_with_vip": {

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

												"network": {
													Type:     schema.TypeString,
													Required: true,
												},

												"virtual_site": {

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

									"vk8s_service": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
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

									"port": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"port_ranges": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"use_default_port": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"advertise_on_public": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"public_ip": {

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

			"advertise_on_public_default_vip": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"do_not_advertise": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"api_specification": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_definition": {

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

						"validation_all_spec_endpoints": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fall_through_mode": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fall_through_mode_allow": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"fall_through_mode_custom": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"open_api_validation_rules": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"action_block": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"action_report": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"action_skip": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"api_endpoint": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"methods": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"path": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																				},
																			},
																		},

																		"api_group": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"base_path": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

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
														},
													},
												},
											},
										},
									},

									"settings": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"oversized_body_fail_validation": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"oversized_body_skip_validation": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"property_validation_settings_custom": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"query_parameters": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"allow_additional_parameters": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"disallow_additional_parameters": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"property_validation_settings_default": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"validation_mode": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"response_validation_mode_active": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"response_validation_properties": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"enforcement_block": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"enforcement_report": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"skip_response_validation": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"skip_validation": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"validation_mode_active": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"request_validation_properties": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"enforcement_block": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"enforcement_report": {

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

						"validation_custom_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fall_through_mode": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fall_through_mode_allow": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"fall_through_mode_custom": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"open_api_validation_rules": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"action_block": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"action_report": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"action_skip": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"api_endpoint": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"methods": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"path": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																				},
																			},
																		},

																		"api_group": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"base_path": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

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
														},
													},
												},
											},
										},
									},

									"open_api_validation_rules": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"api_endpoint": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"methods": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"path": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"api_group": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"base_path": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"any_domain": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"specific_domain": {

													Type:     schema.TypeString,
													Optional: true,
												},

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

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"validation_mode": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"response_validation_mode_active": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"response_validation_properties": {

																			Type: schema.TypeList,

																			Required: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"enforcement_block": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"enforcement_report": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"skip_response_validation": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"skip_validation": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"validation_mode_active": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"request_validation_properties": {

																			Type: schema.TypeList,

																			Required: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"enforcement_block": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"enforcement_report": {

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

									"settings": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"oversized_body_fail_validation": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"oversized_body_skip_validation": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"property_validation_settings_custom": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"query_parameters": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"allow_additional_parameters": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"disallow_additional_parameters": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"property_validation_settings_default": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"validation_disabled": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"disable_api_definition": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_api_discovery": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_api_discovery": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_crawler": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"api_crawler_config": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"domains": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"domain": {
																Type:     schema.TypeString,
																Required: true,
															},

															"simple_login": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"password": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
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

																		"user": {
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

									"disable_api_crawler": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"api_discovery_from_code_scan": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"code_base_integrations": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"all_repos": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"selected_repos": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"api_code_repo": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"code_base_integration": {

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
								},
							},
						},

						"custom_api_auth_discovery": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"api_discovery_ref": {

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

						"default_api_auth_discovery": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"discovered_api_settings": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"purge_duration_for_inactive_discovered_apis": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"disable_learn_from_redirect_traffic": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_learn_from_redirect_traffic": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"api_protection_rules": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_endpoint_rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"deny": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"api_endpoint_method": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_matcher": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"methods": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"api_endpoint_path": {
										Type:     schema.TypeString,
										Required: true,
									},

									"client_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"any_client": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"client_selector": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"expressions": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"ip_threat_category_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ip_threat_categories": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"any_ip": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"asn_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"as_numbers": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},
														},
													},
												},

												"asn_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"asn_sets": {

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

												"ip_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"prefix_sets": {

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

												"ip_prefix_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_match": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"ip_prefixes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"ipv6_prefixes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"tls_fingerprint_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"classes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"exact_values": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"excluded_values": {

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

									"any_domain": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"specific_domain": {

										Type:     schema.TypeString,
										Optional: true,
									},

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

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"request_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cookie_matchers": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"headers": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"jwt_claims": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"query_params": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"key": {
																Type:     schema.TypeString,
																Required: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

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
								},
							},
						},

						"api_groups_rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"allow": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"deny": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"api_group": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"base_path": {
										Type:     schema.TypeString,
										Required: true,
									},

									"client_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"any_client": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"client_selector": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"expressions": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"ip_threat_category_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ip_threat_categories": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"any_ip": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"asn_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"as_numbers": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},
														},
													},
												},

												"asn_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"asn_sets": {

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

												"ip_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"prefix_sets": {

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

												"ip_prefix_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_match": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"ip_prefixes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"ipv6_prefixes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"tls_fingerprint_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"classes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"exact_values": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"excluded_values": {

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

									"any_domain": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"specific_domain": {

										Type:     schema.TypeString,
										Optional: true,
									},

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

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"request_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cookie_matchers": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"headers": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"jwt_claims": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"query_params": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"key": {
																Type:     schema.TypeString,
																Required: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

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
								},
							},
						},
					},
				},
			},

			"api_testing": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_header_value": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"domains": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"allow_destructive_methods": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"credentials": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"credential_name": {
													Type:     schema.TypeString,
													Required: true,
												},

												"api_key": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": {
																Type:     schema.TypeString,
																Required: true,
															},

															"value": {

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

												"basic_auth": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"password": {

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

															"user": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"bearer_token": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"token": {

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

												"login_endpoint": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"json_payload": {

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

															"method": {
																Type:     schema.TypeString,
																Required: true,
															},

															"path": {
																Type:     schema.TypeString,
																Required: true,
															},

															"token_response_key": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"admin": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"standard": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"domain": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"every_day": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"every_month": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"every_week": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"disable_api_testing": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"blocked_clients": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"actions": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"as_number": {

							Type:     schema.TypeInt,
							Optional: true,
						},

						"http_header": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"headers": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_match": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},

												"exact": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"presence": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"regex": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"ip_prefix": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"ipv6_prefix": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"user_identifier": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"expiration_timestamp": {
							Type:     schema.TypeString,
							Optional: true,
						},

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

			"bot_defense": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"policy": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable_js_insert": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"js_insert_all_pages": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"javascript_location": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"js_insert_all_pages_except": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"exclude_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"any_domain": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"domain": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"suffix_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

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

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"path": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"path": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"prefix": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"javascript_location": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"js_insertion_rules": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"exclude_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"any_domain": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"domain": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"suffix_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

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

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"path": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"path": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"prefix": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"rules": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"any_domain": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"domain": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"suffix_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"javascript_location": {
																Type:     schema.TypeString,
																Optional: true,
															},

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

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"path": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"path": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"prefix": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex": {

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

									"javascript_mode": {
										Type:     schema.TypeString,
										Required: true,
									},

									"js_download_path": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"disable_mobile_sdk": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"mobile_sdk_config": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"mobile_identifier": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"headers": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"check_not_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"item": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"exact_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"regex_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"transformers": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
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
											},
										},
									},

									"protected_app_endpoints": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"mobile": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"web": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"web_mobile": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"mobile_identifier": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"any_domain": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"domain": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"exact_value": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"regex_value": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"suffix_value": {

																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"flow_label": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"account_management": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"create": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"password_reset": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"authentication": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"login": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{},
																			},
																		},

																		"login_mfa": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"login_partner": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"logout": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"token_refresh": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"financial_services": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"apply": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"money_transfer": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"flight": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"checkin": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"profile_management": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"create": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"update": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"view": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"search": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"flight_search": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"product_search": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"reservation_search": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"room_search": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"shopping_gift_cards": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"gift_card_make_purchase_with_gift_card": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"gift_card_validation": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_add_to_cart": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_checkout": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_choose_seat": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_enter_drawing_submission": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_make_payment": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_order": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_price_inquiry": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_promo_code_validation": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_purchase_gift_card": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"shop_update_quantity": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"undefined_flow_label": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"allow_good_bots": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"mitigate_good_bots": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"headers": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"http_methods": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

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

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"mitigation": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"block": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"body": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"status": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"flag": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"append_headers": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"auto_type_header_name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},

																					"inference_header_name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																				},
																			},
																		},

																		"no_headers": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"redirect": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"uri": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},

												"path": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"path": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"prefix": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"regex": {

																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"protocol": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"query_params": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"key": {
																Type:     schema.TypeString,
																Required: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

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
								},
							},
						},

						"regional_endpoint": {
							Type:     schema.TypeString,
							Required: true,
						},

						"timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"disable_bot_defense": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"captcha_challenge": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cookie_expiry": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"custom_page": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"enable_challenge": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"captcha_challenge_parameters": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cookie_expiry": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"custom_page": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"default_captcha_challenge_parameters": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_js_challenge_parameters": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"js_challenge_parameters": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cookie_expiry": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"custom_page": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"js_script_delay": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"default_mitigation_settings": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"malicious_user_mitigation": {

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
					},
				},
			},

			"js_challenge": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cookie_expiry": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"custom_page": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"js_script_delay": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"no_challenge": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"policy_based_challenge": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"captcha_challenge_parameters": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cookie_expiry": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"custom_page": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"default_captcha_challenge_parameters": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"always_enable_captcha_challenge": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"always_enable_js_challenge": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_challenge": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_js_challenge_parameters": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"js_challenge_parameters": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cookie_expiry": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"custom_page": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"js_script_delay": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"default_mitigation_settings": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"malicious_user_mitigation": {

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

						"rule_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rules": {

										Type:     schema.TypeList,
										Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"spec": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"arg_matchers": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_not_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"item": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"exact_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"regex_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"transformers": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"any_asn": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"asn_list": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"as_numbers": {

																			Type: schema.TypeList,

																			Required: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeInt,
																			},
																		},
																	},
																},
															},

															"asn_matcher": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"asn_sets": {

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

															"body_matcher": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"disable_challenge": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"enable_captcha_challenge": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"enable_javascript_challenge": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"any_client": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"client_selector": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"expressions": {

																			Type: schema.TypeList,

																			Required: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"cookie_matchers": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_not_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"item": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"exact_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"regex_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"transformers": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"domain_matcher": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"expiration_timestamp": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"headers": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_not_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"item": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"exact_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"regex_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"transformers": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"http_method": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"methods": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"any_ip": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"ip_matcher": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"prefix_sets": {

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

															"ip_prefix_list": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_match": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"ip_prefixes": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"ipv6_prefixes": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"path": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"prefix_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"suffix_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"query_params": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"key": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"check_not_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"item": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"exact_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"regex_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"transformers": {

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

															"tls_fingerprint_matcher": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"classes": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"excluded_values": {

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
								},
							},
						},
					},
				},
			},

			"client_side_defense": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"policy": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable_js_insert": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"js_insert_all_pages": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"js_insert_all_pages_except": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"exclude_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"any_domain": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"domain": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"suffix_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

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

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"path": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"path": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"prefix": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex": {

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

									"js_insertion_rules": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"exclude_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"any_domain": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"domain": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"suffix_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

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

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"path": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"path": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"prefix": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"rules": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"any_domain": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"domain": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"suffix_value": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

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

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"path": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"path": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"prefix": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"regex": {

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

			"disable_client_side_defense": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"cors_policy": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allow_credentials": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"allow_headers": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"allow_methods": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"allow_origin": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"allow_origin_regex": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"disabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"expose_headers": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"maximum_age": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"csrf_policy": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"all_load_balancer_domains": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"custom_domain_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"domains": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"disabled": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"data_guard_rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"apply_data_guard": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"skip_data_guard": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"any_domain": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"exact_value": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"suffix_value": {

							Type:     schema.TypeString,
							Optional: true,
						},

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

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"path": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"path": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"prefix": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"regex": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"ddos_mitigation_rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"expiration_timestamp": {
							Type:     schema.TypeString,
							Optional: true,
						},

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

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"block": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"ddos_client_source": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"asn_list": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"as_numbers": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
											},
										},
									},

									"country_list": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ja4_tls_fingerprint_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"exact_values": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"tls_fingerprint_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"classes": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"exact_values": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"excluded_values": {

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

						"ip_prefix_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"invert_match": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"ip_prefixes": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ipv6_prefixes": {

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

			"default_route_pools": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"endpoint_subsets": {
							Type:     schema.TypeMap,
							Optional: true,
						},

						"cluster": {

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

						"pool": {

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

						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"domains": {

				Type: schema.TypeList,

				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"graphql_rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"any_domain": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"exact_value": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"suffix_value": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"exact_path": {
							Type:     schema.TypeString,
							Required: true,
						},

						"graphql_settings": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable_introspection": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_introspection": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_batched_queries": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"max_depth": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"max_total_length": {
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
						},

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

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"method_get": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"method_post": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"cookie_stickiness": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"add_httponly": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"ignore_httponly": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"path": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"ignore_samesite": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"samesite_lax": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"samesite_none": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"samesite_strict": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"add_secure": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"ignore_secure": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"ttl": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"least_active": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"random": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"ring_hash": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"hash_policy": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cookie": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"add_httponly": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"ignore_httponly": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},

												"path": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"ignore_samesite": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"samesite_lax": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"samesite_none": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"samesite_strict": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"add_secure": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"ignore_secure": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"ttl": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"header_name": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"source_ip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"terminal": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"round_robin": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"source_ip_stickiness": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_ip_reputation": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_ip_reputation": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ip_threat_categories": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"jwt_validation": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"action": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"block": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"report": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"jwks_config": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cleartext": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"mandatory_claims": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"claim_names": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"reserved_claims": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"audience": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"audiences": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"audience_disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"issuer": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"issuer_disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"validate_period_disable": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"validate_period_enable": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"target": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_endpoint": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"api_groups": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"api_groups": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"base_paths": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"base_paths": {

													Type: schema.TypeList,

													Required: true,
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

						"token_location": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bearer_token": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"l7_ddos_protection": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ddos_policy_custom": {

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

						"ddos_policy_none": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"mitigation_block": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"mitigation_js_challenge": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cookie_expiry": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"custom_page": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"js_script_delay": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"http": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"dns_volterra_managed": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"port": {

							Type:     schema.TypeInt,
							Optional: true,
						},

						"port_ranges": {

							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"https": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"add_hsts": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"coalescing_options": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default_coalescing": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"strict_coalescing": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"connection_idle_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"default_loadbalancer": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"non_default_loadbalancer": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"http_protocol_options": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"http_protocol_enable_v1_only": {

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

									"http_protocol_enable_v1_v2": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"http_protocol_enable_v2_only": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"http_redirect": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"disable_path_normalize": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_path_normalize": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"port": {

							Type:     schema.TypeInt,
							Optional: true,
						},

						"port_ranges": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"append_server_name": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"default_header": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"pass_through": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"server_name": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"tls_cert_params": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"certificates": {

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

												"client_certificate_optional": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"crl": {

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

												"no_crl": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"trusted_ca": {

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

												"trusted_ca_url": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"xfcc_disabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"xfcc_options": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"xfcc_header_elements": {

																Type: schema.TypeList,

																Required: true,
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

									"tls_config": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
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

						"tls_parameters": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

												"client_certificate_optional": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"crl": {

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

												"no_crl": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"trusted_ca": {

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

												"trusted_ca_url": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"xfcc_disabled": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"xfcc_options": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"xfcc_header_elements": {

																Type: schema.TypeList,

																Required: true,
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

									"tls_config": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
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
					},
				},
			},

			"https_auto_cert": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"add_hsts": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"coalescing_options": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default_coalescing": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"strict_coalescing": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"connection_idle_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"default_loadbalancer": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"non_default_loadbalancer": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"http_protocol_options": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"http_protocol_enable_v1_only": {

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

									"http_protocol_enable_v1_v2": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"http_protocol_enable_v2_only": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"http_redirect": {
							Type:     schema.TypeBool,
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

									"client_certificate_optional": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"crl": {

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

									"no_crl": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"trusted_ca": {

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

									"trusted_ca_url": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"xfcc_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"xfcc_options": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"xfcc_header_elements": {

													Type: schema.TypeList,

													Required: true,
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

						"disable_path_normalize": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_path_normalize": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"port": {

							Type:     schema.TypeInt,
							Optional: true,
						},

						"port_ranges": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"append_server_name": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"default_header": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"pass_through": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"server_name": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"tls_config": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
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

			"disable_malicious_user_detection": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_malicious_user_detection": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_malware_protection": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"malware_protection_settings": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"malware_protection_rules": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"block": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"report": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"domain": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"any_domain": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
													},
												},

												"domain": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"exact_value": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"regex_value": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"suffix_value": {

																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"http_methods": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

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

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"path": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"path": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"prefix": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"regex": {

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

			"more_option": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"buffer_policy": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_request_bytes": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"compression_params": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"content_length": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"content_type": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"disable_on_etag_header": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"remove_accept_encoding_header": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"custom_errors": {
							Type:     schema.TypeMap,
							Optional: true,
						},

						"disable_default_error_pages": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"idle_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"max_request_header_size": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"request_cookies_to_add": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"overwrite": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"secret_value": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
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

									"value": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"request_cookies_to_remove": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"request_headers_to_add": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"append": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"secret_value": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
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

									"value": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"request_headers_to_remove": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"response_cookies_to_add": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"add_domain": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"ignore_domain": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"add_expiry": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"ignore_expiry": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"add_httponly": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"ignore_httponly": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"ignore_max_age": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_age_value": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"overwrite": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"add_partitioned": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"ignore_partitioned": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"add_path": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"ignore_path": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"ignore_samesite": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"samesite_lax": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"samesite_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"samesite_strict": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"add_secure": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"ignore_secure": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"ignore_value": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"secret_value": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
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

									"value": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"response_cookies_to_remove": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"response_headers_to_add": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"append": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"secret_value": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
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

									"value": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"response_headers_to_remove": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"origin_server_subset_rule_list": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"origin_server_subset_rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"any_asn": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"asn_list": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"as_numbers": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
											},
										},
									},

									"asn_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"asn_sets": {

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

									"country_codes": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"any_ip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"ip_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_matcher": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"prefix_sets": {

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

									"ip_prefix_list": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_match": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"ip_prefixes": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"ipv6_prefixes": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

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

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"origin_server_subsets_action": {
										Type:     schema.TypeMap,
										Required: true,
									},

									"re_name_list": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"client_selector": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"expressions": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"protected_cookies": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"disable_tampering_protection": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_tampering_protection": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"add_httponly": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"ignore_httponly": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"ignore_samesite": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"samesite_lax": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"samesite_none": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"samesite_strict": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"add_secure": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"ignore_secure": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"api_rate_limit": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_endpoint_rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"api_endpoint_method": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_matcher": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"methods": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"api_endpoint_path": {
										Type:     schema.TypeString,
										Required: true,
									},

									"client_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"any_client": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"client_selector": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"expressions": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"ip_threat_category_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ip_threat_categories": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"any_ip": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"asn_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"as_numbers": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},
														},
													},
												},

												"asn_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"asn_sets": {

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

												"ip_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"prefix_sets": {

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

												"ip_prefix_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_match": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"ip_prefixes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"ipv6_prefixes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"tls_fingerprint_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"classes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"exact_values": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"excluded_values": {

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

									"any_domain": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"specific_domain": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"inline_rate_limiter": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ref_user_id": {

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

												"use_http_lb_user_id": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"threshold": {
													Type:     schema.TypeInt,
													Required: true,
												},

												"unit": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"ref_rate_limiter": {

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

									"request_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cookie_matchers": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"headers": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"jwt_claims": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"query_params": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"key": {
																Type:     schema.TypeString,
																Required: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

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
								},
							},
						},

						"bypass_rate_limiting_rules": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bypass_rate_limiting_rules": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"any_client": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"client_selector": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"expressions": {

																			Type: schema.TypeList,

																			Required: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"ip_threat_category_list": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ip_threat_categories": {

																			Type: schema.TypeList,

																			Required: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"any_ip": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"asn_list": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"as_numbers": {

																			Type: schema.TypeList,

																			Required: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeInt,
																			},
																		},
																	},
																},
															},

															"asn_matcher": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"asn_sets": {

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

															"ip_matcher": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"prefix_sets": {

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

															"ip_prefix_list": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_match": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"ip_prefixes": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"ipv6_prefixes": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"tls_fingerprint_matcher": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"classes": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"excluded_values": {

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

												"any_url": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"api_endpoint": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"methods": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"path": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"api_groups": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"api_groups": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"base_path": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"any_domain": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"specific_domain": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"request_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"cookie_matchers": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_not_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"item": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"exact_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"regex_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"transformers": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"headers": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_not_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"item": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"exact_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"regex_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"transformers": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"jwt_claims": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_not_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"item": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"exact_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"regex_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"transformers": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},

																		"name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"query_params": {

																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"invert_matcher": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"key": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"check_not_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"check_present": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"item": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"exact_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"regex_values": {

																						Type: schema.TypeList,

																						Optional: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},

																					"transformers": {

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
											},
										},
									},
								},
							},
						},

						"custom_ip_allowed_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rate_limiter_allowed_prefixes": {

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

						"ip_allowed_list": {

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

						"no_ip_allowed_list": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"server_url_rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"api_group": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"base_path": {
										Type:     schema.TypeString,
										Required: true,
									},

									"client_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"any_client": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"client_selector": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"expressions": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"ip_threat_category_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ip_threat_categories": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"any_ip": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"asn_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"as_numbers": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},
														},
													},
												},

												"asn_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"asn_sets": {

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

												"ip_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"prefix_sets": {

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

												"ip_prefix_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_match": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"ip_prefixes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"ipv6_prefixes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"tls_fingerprint_matcher": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"classes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"exact_values": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"excluded_values": {

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

									"any_domain": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"specific_domain": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"inline_rate_limiter": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ref_user_id": {

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

												"use_http_lb_user_id": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"threshold": {
													Type:     schema.TypeInt,
													Required: true,
												},

												"unit": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"ref_rate_limiter": {

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

									"request_matcher": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cookie_matchers": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"headers": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"jwt_claims": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"query_params": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"invert_matcher": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"key": {
																Type:     schema.TypeString,
																Required: true,
															},

															"check_not_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"check_present": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"item": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"exact_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"regex_values": {

																			Type: schema.TypeList,

																			Optional: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"transformers": {

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
								},
							},
						},
					},
				},
			},

			"disable_rate_limit": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"rate_limit": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_ip_allowed_list": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rate_limiter_allowed_prefixes": {

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

						"ip_allowed_list": {

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

						"no_ip_allowed_list": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_policies": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"policies": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"policies": {

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

						"rate_limiter": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action_block": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hours": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"duration": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},

												"minutes": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"duration": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},

												"seconds": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"duration": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"burst_multiplier": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"period_multiplier": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"total_number": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"unit": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"routes": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_route_object": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"route_ref": {

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
								},
							},
						},

						"direct_response_route": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"headers": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_match": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},

												"exact": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"presence": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"regex": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"http_method": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"incoming_port": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"no_port_match": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"port": {

													Type:     schema.TypeInt,
													Optional: true,
												},

												"port_ranges": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"path": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"path": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"prefix": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"regex": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"route_direct_response": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"response_body": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"response_code": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"redirect_route": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"headers": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_match": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},

												"exact": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"presence": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"regex": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"http_method": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"incoming_port": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"no_port_match": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"port": {

													Type:     schema.TypeInt,
													Optional: true,
												},

												"port_ranges": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"path": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"path": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"prefix": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"regex": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"route_redirect": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"host_redirect": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"proto_redirect": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"remove_all_params": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"replace_params": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"retain_all_params": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"path_redirect": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"prefix_rewrite": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"response_code": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"simple_route": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_options": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"bot_defense_javascript_injection": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"javascript_location": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"javascript_tags": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"javascript_url": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"tag_attributes": {

																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"javascript_tag": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"tag_value": {
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

												"inherited_bot_defense_javascript_injection": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"buffer_policy": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"disabled": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"max_request_bytes": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},

												"common_buffering": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"do_not_retract_cluster": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"retract_cluster": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"cors_policy": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"allow_credentials": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"allow_headers": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"allow_methods": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"allow_origin": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"allow_origin_regex": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"disabled": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"expose_headers": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"maximum_age": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},

												"csrf_policy": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"all_load_balancer_domains": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"custom_domain_list": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"domains": {

																			Type: schema.TypeList,

																			Required: true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},

															"disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"disable_location_add": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"endpoint_subsets": {
													Type:     schema.TypeMap,
													Optional: true,
												},

												"common_hash_policy": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"specific_hash_policy": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"hash_policy": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"cookie": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"add_httponly": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"ignore_httponly": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},

																					"path": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"ignore_samesite": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"samesite_lax": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"samesite_none": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"samesite_strict": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"add_secure": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"ignore_secure": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"ttl": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"header_name": {

																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"source_ip": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"terminal": {
																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},

												"disable_mirroring": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"mirror_policy": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"origin_pool": {

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

															"percent": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"denominator": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"numerator": {
																			Type:     schema.TypeInt,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},

												"priority": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"request_cookies_to_add": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},

															"overwrite": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"secret_value": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"value": {

																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"request_cookies_to_remove": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"request_headers_to_add": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"append": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},

															"secret_value": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"value": {

																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"request_headers_to_remove": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"response_cookies_to_add": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"add_domain": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"ignore_domain": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"add_expiry": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"ignore_expiry": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"add_httponly": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"ignore_httponly": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"ignore_max_age": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"max_age_value": {

																Type:     schema.TypeInt,
																Optional: true,
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},

															"overwrite": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"add_partitioned": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"ignore_partitioned": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"add_path": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"ignore_path": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"ignore_samesite": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"samesite_lax": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"samesite_none": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"samesite_strict": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"add_secure": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"ignore_secure": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"ignore_value": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"secret_value": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"value": {

																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"response_cookies_to_remove": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"response_headers_to_add": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"append": {
																Type:     schema.TypeBool,
																Optional: true,
															},

															"name": {
																Type:     schema.TypeString,
																Required: true,
															},

															"secret_value": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"value": {

																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"response_headers_to_remove": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"default_retry_policy": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"no_retry_policy": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"retry_policy": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"back_off": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"base_interval": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},

																		"max_interval": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},

															"num_retries": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"per_try_timeout": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"retriable_status_codes": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},

															"retry_condition": {

																Type: schema.TypeList,

																Required: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"disable_prefix_rewrite": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"prefix_rewrite": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"regex_rewrite": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"pattern": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"substitution": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"disable_spdy": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"enable_spdy": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"timeout": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"app_firewall": {

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

												"disable_waf": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"inherited_waf": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"disable_web_socket_config": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"web_socket_config": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"use_websocket": {
																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"headers": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_match": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},

												"exact": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"presence": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"regex": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"auto_host_rewrite": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"disable_host_rewrite": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"host_rewrite": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"http_method": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"incoming_port": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"no_port_match": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"port": {

													Type:     schema.TypeInt,
													Optional: true,
												},

												"port_ranges": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"origin_pools": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"endpoint_subsets": {
													Type:     schema.TypeMap,
													Optional: true,
												},

												"cluster": {

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

												"pool": {

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

												"priority": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"weight": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"path": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"path": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"prefix": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"regex": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"query_params": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"remove_all_params": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"replace_params": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"retain_all_params": {

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

			"sensitive_data_disclosure_rules": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"sensitive_data_types_in_response": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"body": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fields": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"api_endpoint": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"methods": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"path": {
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

			"default_sensitive_data_policy": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"sensitive_data_policy": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"sensitive_data_policy_ref": {

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

			"active_service_policies": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"policies": {

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

			"no_service_policies": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"service_policies_from_namespace": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"slow_ddos_mitigation": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"request_headers_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"disable_request_timeout": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"request_timeout": {

							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"system_default_timeouts": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_threat_mesh": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_threat_mesh": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_trust_client_ip_headers": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_trust_client_ip_headers": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"client_ip_headers": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"trusted_clients": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"actions": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"as_number": {

							Type:     schema.TypeInt,
							Optional: true,
						},

						"http_header": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"headers": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_match": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"name": {
													Type:     schema.TypeString,
													Required: true,
												},

												"exact": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"presence": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"regex": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"ip_prefix": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"ipv6_prefix": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"user_identifier": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"expiration_timestamp": {
							Type:     schema.TypeString,
							Optional: true,
						},

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

			"user_id_client_ip": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"user_identification": {

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

			"app_firewall": {

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

			"disable_waf": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"waf_exclusion_rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"any_domain": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"exact_value": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"suffix_value": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"expiration_timestamp": {
							Type:     schema.TypeString,
							Optional: true,
						},

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

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"methods": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"any_path": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"path_prefix": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"path_regex": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"app_firewall_detection_control": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"exclude_attack_type_contexts": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"context": {
													Type:     schema.TypeString,
													Required: true,
												},

												"context_name": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"exclude_attack_type": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"exclude_bot_name_contexts": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"bot_name": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"exclude_signature_contexts": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"context": {
													Type:     schema.TypeString,
													Required: true,
												},

												"context_name": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"signature_id": {
													Type:     schema.TypeInt,
													Required: true,
												},
											},
										},
									},

									"exclude_violation_contexts": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"context": {
													Type:     schema.TypeString,
													Required: true,
												},

												"context_name": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"exclude_violation": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
								},
							},
						},

						"waf_skip_processing": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func ResourceHttpLoadbalancerInstanceStateUpgradeV1(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	httpsAutoCert, ok := rawState["https_auto_cert"].([]interface{})
	if !ok {
		return rawState, nil
	}

	for _, v := range httpsAutoCert {
		a, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		httpProtocolOptions, ok := a["http_protocol_options"].([]interface{})
		if !ok {
			continue
		}

		for _, va := range httpProtocolOptions {
			aa, ok := va.(map[string]interface{})
			if !ok {
				continue
			}

			value, ok := aa["http_protocol_enable_v1_only"]
			if ok && reflect.TypeOf(value).Kind() == reflect.Bool {
				aa["http_protocol_enable_v1_only"] = []interface{}{map[string]interface{}{
					"header_transformation": []interface{}{map[string]interface{}{
						"legacy_header_transformation": true,
					}},
				}}
			}
		}
	}

	https, ok := rawState["https"].([]interface{})
	if !ok {
		return rawState, nil
	}

	for _, v := range https {
		a, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		httpProtocolOptions, ok := a["http_protocol_options"].([]interface{})
		if !ok {
			continue
		}

		for _, va := range httpProtocolOptions {
			aa, ok := va.(map[string]interface{})
			if !ok {
				continue
			}

			value, ok := aa["http_protocol_enable_v1_only"]
			if ok && reflect.TypeOf(value).Kind() == reflect.Bool {
				aa["http_protocol_enable_v1_only"] = []interface{}{map[string]interface{}{
					"header_transformation": []interface{}{map[string]interface{}{
						"legacy_header_transformation": true,
					}},
				}}
			}
		}
	}

	defaultPool, ok := rawState["default_pool"].([]interface{})
	if !ok {
		return rawState, nil
	}

	for _, v := range defaultPool {
		a, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		advancedOptions, ok := a["advanced_options"].([]interface{})
		if !ok {
			continue
		}

		for _, va := range advancedOptions {
			aa, ok := va.(map[string]interface{})
			if !ok {
				continue
			}

			value, ok := aa["http1_config"]
			if ok && reflect.TypeOf(value).Kind() == reflect.Bool {
				aa["http1_config"] = []interface{}{map[string]interface{}{
					"header_transformation": []interface{}{map[string]interface{}{
						"legacy_header_transformation": true,
					}},
				}}
			}
		}
	}
	return rawState, nil
}
