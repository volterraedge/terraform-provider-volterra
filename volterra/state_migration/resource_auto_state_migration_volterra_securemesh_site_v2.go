//
// Copyright (c) 2026 F5 Inc. All rights reserved.
//

package statemigration

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSecureMeshSiteV2InstanceResourceV1() *schema.Resource {
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

			"admin_user_credentials": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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

						"ssh_key": {
							Type:     schema.TypeString,
							Optional: true,
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

			"dns_ntp_config": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_dns": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns_servers": {
										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"f5_dns_default": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"custom_ntp": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ntp_servers": {
										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"f5_ntp_default": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"custom_proxy": {

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

						"proxy_ip_address": {
							Type:     schema.TypeString,
							Required: true,
						},

						"proxy_port": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"disable_re_tunnel": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"enable_re_tunnel": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"username": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"f5_proxy": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"private_adn": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"private_adn": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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

			"no_forward_proxy": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"load_balancing": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"vip_vrrp_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"local_vrf": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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

									"secondary_nameserver": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"secondary_nameserver_v6": {
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

									"secondary_nameserver": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"secondary_nameserver_v6": {
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
					},
				},
			},

			"log_receiver": {

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

			"log_receiver_with_net": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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

						"use_management_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"use_slo_sli": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"logs_streaming_disabled": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_management_network": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_management_network": {

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

			"no_network_policy": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"disable_ha": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_ha": {

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

			"aws": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aws_cloud_user_account": {
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

									"aws_region": {
										Type:     schema.TypeString,
										Required: true,
									},

									"aws_resource_mapping_list": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aws_resource_mappings": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"aws_resources": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"availability_zone": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"security_group": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"subnet_id": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

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

									"disable_cloud_connect": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_cloud_connect": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"tgw_id": {
													Type:     schema.TypeString,
													Required: true,
												},

												"volterra_site_asn": {
													Type:     schema.TypeInt,
													Required: true,
												},
											},
										},
									},

									"disable_disk_encryption": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"disk_encryption_key": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key_id": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"disk_size": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"egress_igw_gw": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"igw_gw_id": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"egress_nat_gw": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"nat_gw_id": {
													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"no_egress": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"private_adn": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"instance_type": {
										Type:     schema.TypeString,
										Required: true,
									},

									"node_list": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node_list": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"aws_az_name": {
																Type:     schema.TypeString,
																Required: true,
															},

															"hostname": {
																Type:     schema.TypeString,
																Required: true,
															},

															"interface_list": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"aws_node_interface_configuration": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"inherit_aws_node_interface_configuration": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"override_aws_node_interface_configuration": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"security_group": {
																									Type:     schema.TypeString,
																									Required: true,
																								},

																								"subnet_id": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},

																		"mtu": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},

																		"network_option": {

																			Type:     schema.TypeList,
																			MaxItems: 1,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"segment_network": {

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

																					"site_local_inside_network": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"site_local_network": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},
																				},
																			},
																		},

																		"site_to_site_connectivity_interface_disabled": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_to_site_connectivity_interface_enabled": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"type": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},
											},
										},
									},

									"cloud_link_config": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cloud_link": {
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

												"inside": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"outside": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"vgw_id": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"private_connectivity_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"disable_private_workload_routing_to_ce": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_private_workload_routing_list": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_private_workload_routing_to_ce": {

													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"subnet_id": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},
											},
										},
									},

									"tags": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"vpc_id": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"azure": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"managed": {

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

									"azure_cred": {
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

									"azure_region": {
										Type:     schema.TypeString,
										Required: true,
									},

									"disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"disk_size": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"machine_type": {
										Type:     schema.TypeString,
										Required: true,
									},

									"resource_group": {
										Type:     schema.TypeString,
										Required: true,
									},

									"multiple_interface": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"node_list": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"azure_az": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"hostname": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"interface_list": {

																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"mtu": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"network_option": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Required: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"segment_network": {

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

																								"site_local_inside_network": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},

																								"site_local_network": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"site_to_site_connectivity_interface_disabled": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"site_to_site_connectivity_interface_enabled": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"subnet": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"existing_subnet": {

																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

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

																											"autogenerate": {

																												Type:     schema.TypeBool,
																												Optional: true,
																											},

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

																		"type": {
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

									"single_interface": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"node_list": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"azure_az": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"hostname": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"interface_list": {

																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"mtu": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"network_option": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Required: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"segment_network": {

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

																								"site_local_inside_network": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},

																								"site_local_network": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"site_to_site_connectivity_interface_disabled": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"site_to_site_connectivity_interface_enabled": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"subnet": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"existing_subnet": {

																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

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

																											"autogenerate": {

																												Type:     schema.TypeBool,
																												Optional: true,
																											},

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

																		"type": {
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
								},
							},
						},

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"baremetal": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"equinix": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"gcp": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disk_size": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"gcp_cred": {
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

									"gcp_region": {
										Type:     schema.TypeString,
										Required: true,
									},

									"instance_type": {
										Type:     schema.TypeString,
										Required: true,
									},

									"private_connectivity": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cloud_link": {
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

												"inside": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"outside": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"private_connectivity_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"multiple_interface": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"node_list": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"gcp_az_name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"hostname": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"interface_list": {

																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"mtu": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"network_option": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Required: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"segment_network": {

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

																								"site_local_inside_network": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},

																								"site_local_network": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"site_to_site_connectivity_interface_disabled": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"site_to_site_connectivity_interface_enabled": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"subnet": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"existing_subnet_id": {

																									Type:     schema.TypeString,
																									Optional: true,
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

																											"autogenerate": {

																												Type:     schema.TypeBool,
																												Optional: true,
																											},

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

																					"vpc": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"existing_vpc_id": {

																									Type:     schema.TypeString,
																									Optional: true,
																								},

																								"new_vpc": {

																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"autogenerate": {

																												Type:     schema.TypeBool,
																												Optional: true,
																											},

																											"name_tag": {

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
										},
									},

									"single_interface": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"node_list": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"node_list": {

																Type:     schema.TypeList,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"gcp_az_name": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"hostname": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"interface_list": {

																			Type:     schema.TypeList,
																			Required: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"mtu": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},

																					"network_option": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Required: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"segment_network": {

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

																								"site_local_inside_network": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},

																								"site_local_network": {

																									Type:     schema.TypeBool,
																									Optional: true,
																								},
																							},
																						},
																					},

																					"site_to_site_connectivity_interface_disabled": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"site_to_site_connectivity_interface_enabled": {

																						Type:     schema.TypeBool,
																						Optional: true,
																					},

																					"subnet": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"existing_subnet_id": {

																									Type:     schema.TypeString,
																									Optional: true,
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

																											"autogenerate": {

																												Type:     schema.TypeBool,
																												Optional: true,
																											},

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

																					"vpc": {

																						Type:     schema.TypeList,
																						MaxItems: 1,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"existing_vpc_id": {

																									Type:     schema.TypeString,
																									Optional: true,
																								},

																								"new_vpc": {

																									Type:     schema.TypeList,
																									MaxItems: 1,
																									Optional: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"autogenerate": {

																												Type:     schema.TypeBool,
																												Optional: true,
																											},

																											"name_tag": {

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
										},
									},

									"tags": {
										Type:     schema.TypeMap,
										Optional: true,
									},
								},
							},
						},

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"kvm": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"nutanix": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"oci": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"openshift_virtualization": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"openstack": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"rseries": {

				Type:       schema.TypeList,
				MaxItems:   1,
				Optional:   true,
				Deprecated: "This field is deprecated and will be removed in future release.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"not_managed": {

							Type:       schema.TypeList,
							MaxItems:   1,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:       schema.TypeList,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"interface_list": {

													Type:       schema.TypeList,
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

															"no_ipv4_address": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"static_ip": {

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

															"description": {
																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"bond_interface": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"devices": {
																			Type: schema.TypeList,

																			Required:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"active_backup": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"lacp": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"rate": {
																						Type:       schema.TypeInt,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},
																				},
																			},
																		},

																		"link_polling_interval": {
																			Type:       schema.TypeInt,
																			Required:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"link_up_delay": {
																			Type:       schema.TypeInt,
																			Required:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"name": {
																			Type:       schema.TypeString,
																			Required:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"ethernet_interface": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:       schema.TypeString,
																			Required:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"mac": {
																			Type:       schema.TypeString,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:       schema.TypeString,
																			Required:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"vlan_id": {
																			Type:       schema.TypeInt,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},
																	},
																},
															},

															"ipv6_auto_config": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"host": {

																			Type:       schema.TypeBool,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																		},

																		"router": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"network_prefix": {

																						Type:       schema.TypeString,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

																					"stateful": {

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

																					"dns_config": {

																						Type:       schema.TypeList,
																						MaxItems:   1,
																						Optional:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"configured_list": {

																									Type:       schema.TypeList,
																									MaxItems:   1,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dns_list": {
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

																								"local_dns": {

																									Type:       schema.TypeList,
																									MaxItems:   1,
																									Optional:   true,
																									Deprecated: "This field is deprecated and will be removed in future release.",
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"configured_address": {

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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:       schema.TypeMap,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"monitor": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{},
																},
															},

															"monitor_disabled": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"mtu": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"name": {
																Type:       schema.TypeString,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"network_option": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

																			Type:       schema.TypeList,
																			MaxItems:   1,
																			Optional:   true,
																			Deprecated: "This field is deprecated and will be removed in future release.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"kind": {
																						Type:       schema.TypeString,
																						Computed:   true,
																						Deprecated: "This field is deprecated and will be removed in future release.",
																					},

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
																	},
																},
															},

															"priority": {
																Type:       schema.TypeInt,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},

												"public_ip": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"type": {
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
					},
				},
			},

			"vmware": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"not_managed": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"node_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"interface_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dhcp_client": {

																Type:     schema.TypeBool,
																Optional: true,
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

															"no_ipv4_address": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"static_ip": {

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

															"description": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"bond_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
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

															"ethernet_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"mac": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},

															"vlan_interface": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"device": {
																			Type:     schema.TypeString,
																			Required: true,
																		},

																		"vlan_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
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

															"is_management": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"is_primary": {
																Type:       schema.TypeBool,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},

															"labels": {
																Type:     schema.TypeMap,
																Optional: true,
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

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"network_option": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Required: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"segment_network": {

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

																		"site_local_inside_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"site_local_network": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},
																	},
																},
															},

															"priority": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"site_to_site_connectivity_interface_disabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"site_to_site_connectivity_interface_enabled": {

																Type:     schema.TypeBool,
																Optional: true,
															},
														},
													},
												},

												"public_ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"type": {
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

			"custom_proxy_bypass": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"proxy_bypass": {
							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"no_proxy_bypass": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"re_select": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"geo_proximity": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"specific_geography": {

							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
						},

						"specific_re": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"backup_re": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"primary_re": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"dc_cluster_group_sli": {

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

			"no_s2s_connectivity_sli": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"dc_cluster_group_slo": {

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

			"no_s2s_connectivity_slo": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"site_mesh_group_on_slo": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"no_site_mesh_group": {

							Type:       schema.TypeBool,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
						},

						"site_mesh_group": {

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

			"segment_vrf": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"segment_config": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"nameserver": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"nameserver_v6": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"secondary_nameserver": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"secondary_nameserver_v6": {
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
								},
							},
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
					},
				},
			},

			"software_settings": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
					},
				},
			},

			"tunnel_dead_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"tunnel_type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"upgrade_settings": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
					},
				},
			},

			"disable_url_categorization": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"enable_url_categorization": {

				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func ResourceSecureMeshSiteV2InstanceStateUpgradeV1(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
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
