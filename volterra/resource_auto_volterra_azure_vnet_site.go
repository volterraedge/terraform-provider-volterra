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
	ves_io_schema_network_firewall "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_firewall"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_views_azure_vnet_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/azure_vnet_site"
)

// resourceVolterraAzureVnetSite is implementation of Volterra's AzureVnetSite resources
func resourceVolterraAzureVnetSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAzureVnetSiteCreate,
		Read:   resourceVolterraAzureVnetSiteRead,
		Update: resourceVolterraAzureVnetSiteUpdate,
		Delete: resourceVolterraAzureVnetSiteDelete,

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

			"azure_region": {
				Type:     schema.TypeString,
				Required: true,
			},

			"assisted": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"azure_cred": {

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

			"disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"machine_type": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"nodes_per_az": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"operating_system_version": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"resource_group": {
				Type:     schema.TypeString,
				Required: true,
			},

			"ingress_egress_gw": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"az_nodes": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"azure_az": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"disk_size": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"inside_subnet": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"subnet_resource_grp": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
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
											},
										},
									},

									"outside_subnet": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"subnet_resource_grp": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
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
											},
										},
									},
								},
							},
						},

						"azure_certified_hw": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"global_network_list": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_network_connections": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sli_to_global_dr": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {

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

												"slo_to_global_dr": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {

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

												"disable_forward_proxy": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"enable_forward_proxy": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connection_timeout": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"max_connect_attempts": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"white_listed_ports": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},

															"white_listed_prefixes": {

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

						"no_global_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"inside_static_routes": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeSet,
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

																Type:     schema.TypeSet,
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

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeSet,
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

																						Type:     schema.TypeSet,
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
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeSet,
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

																			Type:     schema.TypeSet,
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

						"active_network_policies": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_policies": {

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

						"no_network_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_outside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"outside_static_routes": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeSet,
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

																Type:     schema.TypeSet,
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

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeSet,
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

																						Type:     schema.TypeSet,
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
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeSet,
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

																			Type:     schema.TypeSet,
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

						"active_forward_proxy_policies": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"forward_proxy_policies": {

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

						"forward_proxy_allow_all": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_forward_proxy_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"ingress_gw": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"az_nodes": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"azure_az": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"disk_size": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"local_subnet": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"subnet_resource_grp": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
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
											},
										},
									},
								},
							},
						},

						"azure_certified_hw": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"voltstack_cluster": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"az_nodes": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"azure_az": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"disk_size": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"local_subnet": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"subnet": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"subnet_name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"subnet_resource_grp": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},

												"subnet_param": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4": {
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
											},
										},
									},
								},
							},
						},

						"azure_certified_hw": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"global_network_list": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"global_network_connections": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sli_to_global_dr": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {

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

												"slo_to_global_dr": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"global_vn": {

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

												"disable_forward_proxy": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"enable_forward_proxy": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connection_timeout": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"max_connect_attempts": {
																Type:     schema.TypeInt,
																Optional: true,
															},

															"white_listed_ports": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
															},

															"white_listed_prefixes": {

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

						"no_global_network": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"active_network_policies": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_policies": {

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

						"no_network_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_outside_static_routes": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"outside_static_routes": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static_route_list": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"custom_static_route": {

													Type:     schema.TypeSet,
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

																Type:     schema.TypeSet,
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

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ipv4": {

																						Type:     schema.TypeSet,
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

																						Type:     schema.TypeSet,
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
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ipv4": {

																			Type:     schema.TypeSet,
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

																			Type:     schema.TypeSet,
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

						"active_forward_proxy_policies": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"forward_proxy_policies": {

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

						"forward_proxy_allow_all": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"no_forward_proxy_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"ssh_key": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"vnet": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"existing_vnet": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"resource_group": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"vnet_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"new_vnet": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"primary_ipv4": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"volterra_software_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceVolterraAzureVnetSiteCreate creates AzureVnetSite resource
func resourceVolterraAzureVnetSiteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_azure_vnet_site.CreateSpecType{}
	createReq := &ves_io_schema_views_azure_vnet_site.CreateRequest{
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

	if v, ok := d.GetOk("azure_region"); ok && !isIntfNil(v) {

		createSpec.AzureRegion =
			v.(string)
	}

	deploymentTypeFound := false

	if v, ok := d.GetOk("assisted"); ok && !deploymentTypeFound {

		deploymentTypeFound = true

		if v.(bool) {
			deploymentInt := &ves_io_schema_views_azure_vnet_site.CreateSpecType_Assisted{}
			deploymentInt.Assisted = &ves_io_schema.Empty{}
			createSpec.Deployment = deploymentInt
		}

	}

	if v, ok := d.GetOk("azure_cred"); ok && !deploymentTypeFound {

		deploymentTypeFound = true
		deploymentInt := &ves_io_schema_views_azure_vnet_site.CreateSpecType_AzureCred{}
		deploymentInt.AzureCred = &ves_io_schema_views.ObjectRefType{}
		createSpec.Deployment = deploymentInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				deploymentInt.AzureCred.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				deploymentInt.AzureCred.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				deploymentInt.AzureCred.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("disk_size"); ok && !isIntfNil(v) {

		createSpec.DiskSize =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("machine_type"); ok && !isIntfNil(v) {

		createSpec.MachineType =
			v.(string)
	}

	if v, ok := d.GetOk("nodes_per_az"); ok && !isIntfNil(v) {

		createSpec.NodesPerAz =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("operating_system_version"); ok && !isIntfNil(v) {

		createSpec.OperatingSystemVersion =
			v.(string)
	}

	if v, ok := d.GetOk("resource_group"); ok && !isIntfNil(v) {

		createSpec.ResourceGroup =
			v.(string)
	}

	siteTypeTypeFound := false

	if v, ok := d.GetOk("ingress_egress_gw"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_azure_vnet_site.CreateSpecType_IngressEgressGw{}
		siteTypeInt.IngressEgressGw = &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType{}
		createSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.AzureVnetTwoInterfaceNodeType, len(sl))
				siteTypeInt.IngressEgressGw.AzNodes = azNodes
				for i, set := range sl {
					azNodes[i] = &ves_io_schema_views.AzureVnetTwoInterfaceNodeType{}

					azNodesMapStrToI := set.(map[string]interface{})

					if w, ok := azNodesMapStrToI["azure_az"]; ok && !isIntfNil(w) {
						azNodes[i].AzureAz = w.(string)
					}

					if w, ok := azNodesMapStrToI["disk_size"]; ok && !isIntfNil(w) {
						azNodes[i].DiskSize = w.(string)
					}

					if v, ok := azNodesMapStrToI["inside_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						insideSubnet := &ves_io_schema_views.AzureSubnetChoiceType{}
						azNodes[i].InsideSubnet = insideSubnet
						for _, set := range sl {

							insideSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := insideSubnetMapStrToI["subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.AzureSubnetChoiceType_Subnet{}
								choiceInt.Subnet = &ves_io_schema_views.AzureSubnetType{}
								insideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.Subnet.SubnetName = v.(string)
									}

									if v, ok := cs["subnet_resource_grp"]; ok && !isIntfNil(v) {

										choiceInt.Subnet.SubnetResourceGrp = v.(string)
									}

								}

							}

							if v, ok := insideSubnetMapStrToI["subnet_param"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.AzureSubnetChoiceType_SubnetParam{}
								choiceInt.SubnetParam = &ves_io_schema_views.CloudSubnetParamType{}
								insideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["ipv4"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv4 = v.(string)
									}

									if v, ok := cs["ipv6"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv6 = v.(string)
									}

								}

							}

						}

					}

					if v, ok := azNodesMapStrToI["outside_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						outsideSubnet := &ves_io_schema_views.AzureSubnetChoiceType{}
						azNodes[i].OutsideSubnet = outsideSubnet
						for _, set := range sl {

							outsideSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := outsideSubnetMapStrToI["subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.AzureSubnetChoiceType_Subnet{}
								choiceInt.Subnet = &ves_io_schema_views.AzureSubnetType{}
								outsideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.Subnet.SubnetName = v.(string)
									}

									if v, ok := cs["subnet_resource_grp"]; ok && !isIntfNil(v) {

										choiceInt.Subnet.SubnetResourceGrp = v.(string)
									}

								}

							}

							if v, ok := outsideSubnetMapStrToI["subnet_param"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.AzureSubnetChoiceType_SubnetParam{}
								choiceInt.SubnetParam = &ves_io_schema_views.CloudSubnetParamType{}
								outsideSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["ipv4"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv4 = v.(string)
									}

									if v, ok := cs["ipv6"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv6 = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["azure_certified_hw"]; ok && !isIntfNil(v) {

				siteTypeInt.IngressEgressGw.AzureCertifiedHw = v.(string)
			}

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["global_network_connections"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						globalNetworkConnections := make([]*ves_io_schema_views.GlobalNetworkConnectionType, len(sl))
						globalNetworkChoiceInt.GlobalNetworkList.GlobalNetworkConnections = globalNetworkConnections
						for i, set := range sl {
							globalNetworkConnections[i] = &ves_io_schema_views.GlobalNetworkConnectionType{}

							globalNetworkConnectionsMapStrToI := set.(map[string]interface{})

							connectionChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["sli_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SliToGlobalDr{}
								connectionChoiceInt.SliToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SliToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["slo_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SloToGlobalDr{}
								connectionChoiceInt.SloToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SloToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							forwardProxyChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["disable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true

								if v.(bool) {
									forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_DisableForwardProxy{}
									forwardProxyChoiceInt.DisableForwardProxy = &ves_io_schema.Empty{}
									globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt
								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["enable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true
								forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_EnableForwardProxy{}
								forwardProxyChoiceInt.EnableForwardProxy = &ves_io_schema.ForwardProxyConfigType{}
								globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["connection_timeout"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.ConnectionTimeout = uint32(v.(int))
									}

									if v, ok := cs["max_connect_attempts"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.MaxConnectAttempts = uint32(v.(int))
									}

									if v, ok := cs["white_listed_ports"]; ok && !isIntfNil(v) {

										ls := make([]uint32, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {

											ls[i] = uint32(v.(int))
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPorts = ls

									}

									if v, ok := cs["white_listed_prefixes"]; ok && !isIntfNil(v) {

										ls := make([]string, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {
											ls[i] = v.(string)
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPrefixes = ls

									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			insideStaticRouteChoiceTypeFound := false

			if v, ok := cs["inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true
				insideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_InsideStaticRoutes{}
				insideStaticRouteChoiceInt.InsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						insideStaticRouteChoiceInt.InsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

			if v, ok := cs["no_inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					insideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_NoInsideStaticRoutes{}
					insideStaticRouteChoiceInt.NoInsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						networkPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						networkPolicyChoiceInt.ActiveNetworkPolicies.NetworkPolicies = networkPoliciesInt
						for i, ps := range sl {

							npMapToStrVal := ps.(map[string]interface{})
							networkPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := npMapToStrVal["name"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Name = v.(string)
							}

							if v, ok := npMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := npMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						outsideStaticRouteChoiceInt.OutsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

			servicePolicyChoiceTypeFound := false

			if v, ok := cs["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["forward_proxy_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						forwardProxyPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						servicePolicyChoiceInt.ActiveForwardProxyPolicies.ForwardProxyPolicies = forwardProxyPoliciesInt
						for i, ps := range sl {

							fppMapToStrVal := ps.(map[string]interface{})
							forwardProxyPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := fppMapToStrVal["name"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Name = v.(string)
							}

							if v, ok := fppMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := fppMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("ingress_gw"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_azure_vnet_site.CreateSpecType_IngressGw{}
		siteTypeInt.IngressGw = &ves_io_schema_views_azure_vnet_site.AzureVnetIngressGwType{}
		createSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.AzureVnetOneInterfaceNodeType, len(sl))
				siteTypeInt.IngressGw.AzNodes = azNodes
				for i, set := range sl {
					azNodes[i] = &ves_io_schema_views.AzureVnetOneInterfaceNodeType{}

					azNodesMapStrToI := set.(map[string]interface{})

					if w, ok := azNodesMapStrToI["azure_az"]; ok && !isIntfNil(w) {
						azNodes[i].AzureAz = w.(string)
					}

					if w, ok := azNodesMapStrToI["disk_size"]; ok && !isIntfNil(w) {
						azNodes[i].DiskSize = w.(string)
					}

					if v, ok := azNodesMapStrToI["local_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						localSubnet := &ves_io_schema_views.AzureSubnetChoiceType{}
						azNodes[i].LocalSubnet = localSubnet
						for _, set := range sl {

							localSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := localSubnetMapStrToI["subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.AzureSubnetChoiceType_Subnet{}
								choiceInt.Subnet = &ves_io_schema_views.AzureSubnetType{}
								localSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.Subnet.SubnetName = v.(string)
									}

									if v, ok := cs["subnet_resource_grp"]; ok && !isIntfNil(v) {

										choiceInt.Subnet.SubnetResourceGrp = v.(string)
									}

								}

							}

							if v, ok := localSubnetMapStrToI["subnet_param"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.AzureSubnetChoiceType_SubnetParam{}
								choiceInt.SubnetParam = &ves_io_schema_views.CloudSubnetParamType{}
								localSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["ipv4"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv4 = v.(string)
									}

									if v, ok := cs["ipv6"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv6 = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["azure_certified_hw"]; ok && !isIntfNil(v) {

				siteTypeInt.IngressGw.AzureCertifiedHw = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("voltstack_cluster"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_azure_vnet_site.CreateSpecType_VoltstackCluster{}
		siteTypeInt.VoltstackCluster = &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType{}
		createSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["az_nodes"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				azNodes := make([]*ves_io_schema_views.AzureVnetOneInterfaceNodeType, len(sl))
				siteTypeInt.VoltstackCluster.AzNodes = azNodes
				for i, set := range sl {
					azNodes[i] = &ves_io_schema_views.AzureVnetOneInterfaceNodeType{}

					azNodesMapStrToI := set.(map[string]interface{})

					if w, ok := azNodesMapStrToI["azure_az"]; ok && !isIntfNil(w) {
						azNodes[i].AzureAz = w.(string)
					}

					if w, ok := azNodesMapStrToI["disk_size"]; ok && !isIntfNil(w) {
						azNodes[i].DiskSize = w.(string)
					}

					if v, ok := azNodesMapStrToI["local_subnet"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						localSubnet := &ves_io_schema_views.AzureSubnetChoiceType{}
						azNodes[i].LocalSubnet = localSubnet
						for _, set := range sl {

							localSubnetMapStrToI := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := localSubnetMapStrToI["subnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.AzureSubnetChoiceType_Subnet{}
								choiceInt.Subnet = &ves_io_schema_views.AzureSubnetType{}
								localSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["subnet_name"]; ok && !isIntfNil(v) {

										choiceInt.Subnet.SubnetName = v.(string)
									}

									if v, ok := cs["subnet_resource_grp"]; ok && !isIntfNil(v) {

										choiceInt.Subnet.SubnetResourceGrp = v.(string)
									}

								}

							}

							if v, ok := localSubnetMapStrToI["subnet_param"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceInt := &ves_io_schema_views.AzureSubnetChoiceType_SubnetParam{}
								choiceInt.SubnetParam = &ves_io_schema_views.CloudSubnetParamType{}
								localSubnet.Choice = choiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["ipv4"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv4 = v.(string)
									}

									if v, ok := cs["ipv6"]; ok && !isIntfNil(v) {

										choiceInt.SubnetParam.Ipv6 = v.(string)
									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["azure_certified_hw"]; ok && !isIntfNil(v) {

				siteTypeInt.VoltstackCluster.AzureCertifiedHw = v.(string)
			}

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["global_network_connections"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						globalNetworkConnections := make([]*ves_io_schema_views.GlobalNetworkConnectionType, len(sl))
						globalNetworkChoiceInt.GlobalNetworkList.GlobalNetworkConnections = globalNetworkConnections
						for i, set := range sl {
							globalNetworkConnections[i] = &ves_io_schema_views.GlobalNetworkConnectionType{}

							globalNetworkConnectionsMapStrToI := set.(map[string]interface{})

							connectionChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["sli_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SliToGlobalDr{}
								connectionChoiceInt.SliToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SliToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["slo_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SloToGlobalDr{}
								connectionChoiceInt.SloToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SloToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							forwardProxyChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["disable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true

								if v.(bool) {
									forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_DisableForwardProxy{}
									forwardProxyChoiceInt.DisableForwardProxy = &ves_io_schema.Empty{}
									globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt
								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["enable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true
								forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_EnableForwardProxy{}
								forwardProxyChoiceInt.EnableForwardProxy = &ves_io_schema.ForwardProxyConfigType{}
								globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["connection_timeout"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.ConnectionTimeout = uint32(v.(int))
									}

									if v, ok := cs["max_connect_attempts"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.MaxConnectAttempts = uint32(v.(int))
									}

									if v, ok := cs["white_listed_ports"]; ok && !isIntfNil(v) {

										ls := make([]uint32, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {

											ls[i] = uint32(v.(int))
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPorts = ls

									}

									if v, ok := cs["white_listed_prefixes"]; ok && !isIntfNil(v) {

										ls := make([]string, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {
											ls[i] = v.(string)
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPrefixes = ls

									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						networkPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						networkPolicyChoiceInt.ActiveNetworkPolicies.NetworkPolicies = networkPoliciesInt
						for i, ps := range sl {

							npMapToStrVal := ps.(map[string]interface{})
							networkPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := npMapToStrVal["name"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Name = v.(string)
							}

							if v, ok := npMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := npMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						outsideStaticRouteChoiceInt.OutsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

			servicePolicyChoiceTypeFound := false

			if v, ok := cs["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["forward_proxy_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						forwardProxyPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						servicePolicyChoiceInt.ActiveForwardProxyPolicies.ForwardProxyPolicies = forwardProxyPoliciesInt
						for i, ps := range sl {

							fppMapToStrVal := ps.(map[string]interface{})
							forwardProxyPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := fppMapToStrVal["name"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Name = v.(string)
							}

							if v, ok := fppMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := fppMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("ssh_key"); ok && !isIntfNil(v) {

		createSpec.SshKey =
			v.(string)
	}

	if v, ok := d.GetOk("vnet"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		vnet := &ves_io_schema_views.AzureVnetChoiceType{}
		createSpec.Vnet = vnet
		for _, set := range sl {

			vnetMapStrToI := set.(map[string]interface{})

			choiceTypeFound := false

			if v, ok := vnetMapStrToI["existing_vnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views.AzureVnetChoiceType_ExistingVnet{}
				choiceInt.ExistingVnet = &ves_io_schema_views.AzureVnetType{}
				vnet.Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["resource_group"]; ok && !isIntfNil(v) {

						choiceInt.ExistingVnet.ResourceGroup = v.(string)
					}

					if v, ok := cs["vnet_name"]; ok && !isIntfNil(v) {

						choiceInt.ExistingVnet.VnetName = v.(string)
					}

				}

			}

			if v, ok := vnetMapStrToI["new_vnet"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views.AzureVnetChoiceType_NewVnet{}
				choiceInt.NewVnet = &ves_io_schema_views.AzureVnetParamsType{}
				vnet.Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						choiceInt.NewVnet.Name = v.(string)
					}

					if v, ok := cs["primary_ipv4"]; ok && !isIntfNil(v) {

						choiceInt.NewVnet.PrimaryIpv4 = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("volterra_software_version"); ok && !isIntfNil(v) {

		createSpec.VolterraSoftwareVersion =
			v.(string)
	}

	log.Printf("[DEBUG] Creating Volterra AzureVnetSite object with struct: %+v", createReq)

	createAzureVnetSiteResp, err := client.CreateObject(context.Background(), ves_io_schema_views_azure_vnet_site.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating AzureVnetSite: %s", err)
	}
	d.SetId(createAzureVnetSiteResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraAzureVnetSiteRead(d, meta)
}

func resourceVolterraAzureVnetSiteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_azure_vnet_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AzureVnetSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AzureVnetSite %q: %s", d.Id(), err)
	}
	return setAzureVnetSiteFields(client, d, resp)
}

func setAzureVnetSiteFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraAzureVnetSiteUpdate updates AzureVnetSite resource
func resourceVolterraAzureVnetSiteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_azure_vnet_site.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_azure_vnet_site.ReplaceRequest{
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

	if v, ok := d.GetOk("nodes_per_az"); ok && !isIntfNil(v) {

		updateSpec.NodesPerAz =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("operating_system_version"); ok && !isIntfNil(v) {

		updateSpec.OperatingSystemVersion =
			v.(string)
	}

	siteTypeTypeFound := false

	if v, ok := d.GetOk("ingress_egress_gw"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_azure_vnet_site.ReplaceSpecType_IngressEgressGw{}
		siteTypeInt.IngressEgressGw = &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType{}
		updateSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["global_network_connections"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						globalNetworkConnections := make([]*ves_io_schema_views.GlobalNetworkConnectionType, len(sl))
						globalNetworkChoiceInt.GlobalNetworkList.GlobalNetworkConnections = globalNetworkConnections
						for i, set := range sl {
							globalNetworkConnections[i] = &ves_io_schema_views.GlobalNetworkConnectionType{}

							globalNetworkConnectionsMapStrToI := set.(map[string]interface{})

							connectionChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["sli_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SliToGlobalDr{}
								connectionChoiceInt.SliToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SliToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["slo_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SloToGlobalDr{}
								connectionChoiceInt.SloToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SloToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							forwardProxyChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["disable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true

								if v.(bool) {
									forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_DisableForwardProxy{}
									forwardProxyChoiceInt.DisableForwardProxy = &ves_io_schema.Empty{}
									globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt
								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["enable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true
								forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_EnableForwardProxy{}
								forwardProxyChoiceInt.EnableForwardProxy = &ves_io_schema.ForwardProxyConfigType{}
								globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["connection_timeout"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.ConnectionTimeout = uint32(v.(int))
									}

									if v, ok := cs["max_connect_attempts"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.MaxConnectAttempts = uint32(v.(int))
									}

									if v, ok := cs["white_listed_ports"]; ok && !isIntfNil(v) {

										ls := make([]uint32, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {

											ls[i] = uint32(v.(int))
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPorts = ls

									}

									if v, ok := cs["white_listed_prefixes"]; ok && !isIntfNil(v) {

										ls := make([]string, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {
											ls[i] = v.(string)
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPrefixes = ls

									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			insideStaticRouteChoiceTypeFound := false

			if v, ok := cs["inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true
				insideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_InsideStaticRoutes{}
				insideStaticRouteChoiceInt.InsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						insideStaticRouteChoiceInt.InsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

			if v, ok := cs["no_inside_static_routes"]; ok && !isIntfNil(v) && !insideStaticRouteChoiceTypeFound {

				insideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					insideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_NoInsideStaticRoutes{}
					insideStaticRouteChoiceInt.NoInsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.InsideStaticRouteChoice = insideStaticRouteChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						networkPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						networkPolicyChoiceInt.ActiveNetworkPolicies.NetworkPolicies = networkPoliciesInt
						for i, ps := range sl {

							npMapToStrVal := ps.(map[string]interface{})
							networkPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := npMapToStrVal["name"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Name = v.(string)
							}

							if v, ok := npMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := npMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.IngressEgressGw.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						outsideStaticRouteChoiceInt.OutsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

			servicePolicyChoiceTypeFound := false

			if v, ok := cs["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["forward_proxy_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						forwardProxyPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						servicePolicyChoiceInt.ActiveForwardProxyPolicies.ForwardProxyPolicies = forwardProxyPoliciesInt
						for i, ps := range sl {

							fppMapToStrVal := ps.(map[string]interface{})
							forwardProxyPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := fppMapToStrVal["name"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Name = v.(string)
							}

							if v, ok := fppMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := fppMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetIngressEgressGwReplaceType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.IngressEgressGw.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("ingress_gw"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		_ = v
	}

	if v, ok := d.GetOk("voltstack_cluster"); ok && !siteTypeTypeFound {

		siteTypeTypeFound = true
		siteTypeInt := &ves_io_schema_views_azure_vnet_site.ReplaceSpecType_VoltstackCluster{}
		siteTypeInt.VoltstackCluster = &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType{}
		updateSpec.SiteType = siteTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			globalNetworkChoiceTypeFound := false

			if v, ok := cs["global_network_list"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true
				globalNetworkChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType_GlobalNetworkList{}
				globalNetworkChoiceInt.GlobalNetworkList = &ves_io_schema_views.GlobalNetworkConnectionListType{}
				siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["global_network_connections"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						globalNetworkConnections := make([]*ves_io_schema_views.GlobalNetworkConnectionType, len(sl))
						globalNetworkChoiceInt.GlobalNetworkList.GlobalNetworkConnections = globalNetworkConnections
						for i, set := range sl {
							globalNetworkConnections[i] = &ves_io_schema_views.GlobalNetworkConnectionType{}

							globalNetworkConnectionsMapStrToI := set.(map[string]interface{})

							connectionChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["sli_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SliToGlobalDr{}
								connectionChoiceInt.SliToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SliToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["slo_to_global_dr"]; ok && !isIntfNil(v) && !connectionChoiceTypeFound {

								connectionChoiceTypeFound = true
								connectionChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_SloToGlobalDr{}
								connectionChoiceInt.SloToGlobalDr = &ves_io_schema_views.GlobalConnectorType{}
								globalNetworkConnections[i].ConnectionChoice = connectionChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["global_vn"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										globalVn := &ves_io_schema_views.ObjectRefType{}
										connectionChoiceInt.SloToGlobalDr.GlobalVn = globalVn
										for _, set := range sl {

											globalVnMapStrToI := set.(map[string]interface{})

											if w, ok := globalVnMapStrToI["name"]; ok && !isIntfNil(w) {
												globalVn.Name = w.(string)
											}

											if w, ok := globalVnMapStrToI["namespace"]; ok && !isIntfNil(w) {
												globalVn.Namespace = w.(string)
											}

											if w, ok := globalVnMapStrToI["tenant"]; ok && !isIntfNil(w) {
												globalVn.Tenant = w.(string)
											}

										}

									}

								}

							}

							forwardProxyChoiceTypeFound := false

							if v, ok := globalNetworkConnectionsMapStrToI["disable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true

								if v.(bool) {
									forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_DisableForwardProxy{}
									forwardProxyChoiceInt.DisableForwardProxy = &ves_io_schema.Empty{}
									globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt
								}

							}

							if v, ok := globalNetworkConnectionsMapStrToI["enable_forward_proxy"]; ok && !isIntfNil(v) && !forwardProxyChoiceTypeFound {

								forwardProxyChoiceTypeFound = true
								forwardProxyChoiceInt := &ves_io_schema_views.GlobalNetworkConnectionType_EnableForwardProxy{}
								forwardProxyChoiceInt.EnableForwardProxy = &ves_io_schema.ForwardProxyConfigType{}
								globalNetworkConnections[i].ForwardProxyChoice = forwardProxyChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["connection_timeout"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.ConnectionTimeout = uint32(v.(int))
									}

									if v, ok := cs["max_connect_attempts"]; ok && !isIntfNil(v) {

										forwardProxyChoiceInt.EnableForwardProxy.MaxConnectAttempts = uint32(v.(int))
									}

									if v, ok := cs["white_listed_ports"]; ok && !isIntfNil(v) {

										ls := make([]uint32, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {

											ls[i] = uint32(v.(int))
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPorts = ls

									}

									if v, ok := cs["white_listed_prefixes"]; ok && !isIntfNil(v) {

										ls := make([]string, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {
											ls[i] = v.(string)
										}
										forwardProxyChoiceInt.EnableForwardProxy.WhiteListedPrefixes = ls

									}

								}

							}

						}

					}

				}

			}

			if v, ok := cs["no_global_network"]; ok && !isIntfNil(v) && !globalNetworkChoiceTypeFound {

				globalNetworkChoiceTypeFound = true

				if v.(bool) {
					globalNetworkChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType_NoGlobalNetwork{}
					globalNetworkChoiceInt.NoGlobalNetwork = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.GlobalNetworkChoice = globalNetworkChoiceInt
				}

			}

			networkPolicyChoiceTypeFound := false

			if v, ok := cs["active_network_policies"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true
				networkPolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType_ActiveNetworkPolicies{}
				networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
				siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						networkPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						networkPolicyChoiceInt.ActiveNetworkPolicies.NetworkPolicies = networkPoliciesInt
						for i, ps := range sl {

							npMapToStrVal := ps.(map[string]interface{})
							networkPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := npMapToStrVal["name"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Name = v.(string)
							}

							if v, ok := npMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := npMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								networkPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["no_network_policy"]; ok && !isIntfNil(v) && !networkPolicyChoiceTypeFound {

				networkPolicyChoiceTypeFound = true

				if v.(bool) {
					networkPolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType_NoNetworkPolicy{}
					networkPolicyChoiceInt.NoNetworkPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.NetworkPolicyChoice = networkPolicyChoiceInt
				}

			}

			outsideStaticRouteChoiceTypeFound := false

			if v, ok := cs["no_outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true

				if v.(bool) {
					outsideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType_NoOutsideStaticRoutes{}
					outsideStaticRouteChoiceInt.NoOutsideStaticRoutes = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt
				}

			}

			if v, ok := cs["outside_static_routes"]; ok && !isIntfNil(v) && !outsideStaticRouteChoiceTypeFound {

				outsideStaticRouteChoiceTypeFound = true
				outsideStaticRouteChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType_OutsideStaticRoutes{}
				outsideStaticRouteChoiceInt.OutsideStaticRoutes = &ves_io_schema_views.SiteStaticRoutesListType{}
				siteTypeInt.VoltstackCluster.OutsideStaticRouteChoice = outsideStaticRouteChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["static_route_list"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						staticRouteList := make([]*ves_io_schema_views.SiteStaticRoutesType, len(sl))
						outsideStaticRouteChoiceInt.OutsideStaticRoutes.StaticRouteList = staticRouteList
						for i, set := range sl {
							staticRouteList[i] = &ves_io_schema_views.SiteStaticRoutesType{}

							staticRouteListMapStrToI := set.(map[string]interface{})

							configModeChoiceTypeFound := false

							if v, ok := staticRouteListMapStrToI["custom_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_CustomStaticRoute{}
								configModeChoiceInt.CustomStaticRoute = &ves_io_schema.StaticRouteType{}
								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["attrs"]; ok && !isIntfNil(v) {

										attrsList := []ves_io_schema.RouteAttrType{}
										for _, j := range v.([]interface{}) {
											attrsList = append(attrsList, ves_io_schema.RouteAttrType(ves_io_schema.RouteAttrType_value[j.(string)]))
										}
										configModeChoiceInt.CustomStaticRoute.Attrs = attrsList

									}

									if v, ok := cs["labels"]; ok && !isIntfNil(v) {

										ms := map[string]string{}
										for k, v := range v.(map[string]interface{}) {
											ms[k] = v.(string)
										}
										configModeChoiceInt.CustomStaticRoute.Labels = ms
									}

									if v, ok := cs["nexthop"]; ok && !isIntfNil(v) {

										sl := v.(*schema.Set).List()
										nexthop := &ves_io_schema.NextHopType{}
										configModeChoiceInt.CustomStaticRoute.Nexthop = nexthop
										for _, set := range sl {

											nexthopMapStrToI := set.(map[string]interface{})

											if v, ok := nexthopMapStrToI["interface"]; ok && !isIntfNil(v) {

												sl := v.([]interface{})
												intfInt := make([]*ves_io_schema.ObjectRefType, len(sl))
												nexthop.Interface = intfInt
												for i, ps := range sl {

													iMapToStrVal := ps.(map[string]interface{})
													intfInt[i] = &ves_io_schema.ObjectRefType{}

													intfInt[i].Kind = "network_interface"

													if v, ok := iMapToStrVal["name"]; ok && !isIntfNil(v) {
														intfInt[i].Name = v.(string)
													}

													if v, ok := iMapToStrVal["namespace"]; ok && !isIntfNil(v) {
														intfInt[i].Namespace = v.(string)
													}

													if v, ok := iMapToStrVal["tenant"]; ok && !isIntfNil(v) {
														intfInt[i].Tenant = v.(string)
													}

													if v, ok := iMapToStrVal["uid"]; ok && !isIntfNil(v) {
														intfInt[i].Uid = v.(string)
													}

												}

											}

											if v, ok := nexthopMapStrToI["nexthop_address"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												nexthopAddress := &ves_io_schema.IpAddressType{}
												nexthop.NexthopAddress = nexthopAddress
												for _, set := range sl {

													nexthopAddressMapStrToI := set.(map[string]interface{})

													verTypeFound := false

													if v, ok := nexthopAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv4{}
														verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv4.Addr = v.(string)
															}

														}

													}

													if v, ok := nexthopAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

														verTypeFound = true
														verInt := &ves_io_schema.IpAddressType_Ipv6{}
														verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
														nexthopAddress.Ver = verInt

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															cs := set.(map[string]interface{})

															if v, ok := cs["addr"]; ok && !isIntfNil(v) {

																verInt.Ipv6.Addr = v.(string)
															}

														}

													}

												}

											}

											if v, ok := nexthopMapStrToI["type"]; ok && !isIntfNil(v) {

												nexthop.Type = ves_io_schema.NextHopTypes(ves_io_schema.NextHopTypes_value[v.(string)])

											}

										}

									}

									if v, ok := cs["subnets"]; ok && !isIntfNil(v) {

										sl := v.([]interface{})
										subnets := make([]*ves_io_schema.IpSubnetType, len(sl))
										configModeChoiceInt.CustomStaticRoute.Subnets = subnets
										for i, set := range sl {
											subnets[i] = &ves_io_schema.IpSubnetType{}

											subnetsMapStrToI := set.(map[string]interface{})

											verTypeFound := false

											if v, ok := subnetsMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv4{}
												verInt.Ipv4 = &ves_io_schema.Ipv4SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv4.Prefix = v.(string)
													}

												}

											}

											if v, ok := subnetsMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

												verTypeFound = true
												verInt := &ves_io_schema.IpSubnetType_Ipv6{}
												verInt.Ipv6 = &ves_io_schema.Ipv6SubnetType{}
												subnets[i].Ver = verInt

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["plen"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Plen = uint32(v.(int))
													}

													if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

														verInt.Ipv6.Prefix = v.(string)
													}

												}

											}

										}

									}

								}

							}

							if v, ok := staticRouteListMapStrToI["simple_static_route"]; ok && !isIntfNil(v) && !configModeChoiceTypeFound {

								configModeChoiceTypeFound = true
								configModeChoiceInt := &ves_io_schema_views.SiteStaticRoutesType_SimpleStaticRoute{}

								staticRouteList[i].ConfigModeChoice = configModeChoiceInt

								configModeChoiceInt.SimpleStaticRoute = v.(string)

							}

						}

					}

				}

			}

			servicePolicyChoiceTypeFound := false

			if v, ok := cs["active_forward_proxy_policies"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true
				servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType_ActiveForwardProxyPolicies{}
				servicePolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
				siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["forward_proxy_policies"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						forwardProxyPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						servicePolicyChoiceInt.ActiveForwardProxyPolicies.ForwardProxyPolicies = forwardProxyPoliciesInt
						for i, ps := range sl {

							fppMapToStrVal := ps.(map[string]interface{})
							forwardProxyPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := fppMapToStrVal["name"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Name = v.(string)
							}

							if v, ok := fppMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Namespace = v.(string)
							}

							if v, ok := fppMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								forwardProxyPoliciesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["forward_proxy_allow_all"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType_ForwardProxyAllowAll{}
					servicePolicyChoiceInt.ForwardProxyAllowAll = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

			if v, ok := cs["no_forward_proxy_policy"]; ok && !isIntfNil(v) && !servicePolicyChoiceTypeFound {

				servicePolicyChoiceTypeFound = true

				if v.(bool) {
					servicePolicyChoiceInt := &ves_io_schema_views_azure_vnet_site.AzureVnetVoltstackClusterReplaceType_NoForwardProxyPolicy{}
					servicePolicyChoiceInt.NoForwardProxyPolicy = &ves_io_schema.Empty{}
					siteTypeInt.VoltstackCluster.ServicePolicyChoice = servicePolicyChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("volterra_software_version"); ok && !isIntfNil(v) {

		updateSpec.VolterraSoftwareVersion =
			v.(string)
	}

	log.Printf("[DEBUG] Updating Volterra AzureVnetSite obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_azure_vnet_site.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating AzureVnetSite: %s", err)
	}

	return resourceVolterraAzureVnetSiteRead(d, meta)
}

func resourceVolterraAzureVnetSiteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_azure_vnet_site.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AzureVnetSite %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AzureVnetSite before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra AzureVnetSite obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_azure_vnet_site.ObjectType, namespace, name)
}
