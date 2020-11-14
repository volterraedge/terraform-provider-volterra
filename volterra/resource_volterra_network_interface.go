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
	ves_io_schema_network_interface "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_interface"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraNetworkInterface is implementation of Volterra's NetworkInterface resources
func resourceVolterraNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraNetworkInterfaceCreate,
		Read:   resourceVolterraNetworkInterfaceRead,
		Update: resourceVolterraNetworkInterfaceUpdate,
		Delete: resourceVolterraNetworkInterfaceDelete,

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

			"interface_choice": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"dedicated_interface": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"mtu": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"node_choice": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

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

									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"ethernet_interface": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_choice": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dhcp_client": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"dhcp_server": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dhcp_networks": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dns_choice": {

																			Type:     schema.TypeSet,
																			Optional: true,
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
																				},
																			},
																		},

																		"gateway_choice": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

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
																				},
																			},
																		},

																		"network_prefix_choice": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"network_prefix": {

																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"network_prefix_allocator": {

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

																		"pool_settings": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"pools": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"end_ip": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"exclude": {
																						Type:     schema.TypeBool,
																						Optional: true,
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
																Type:     schema.TypeString,
																Optional: true,
															},

															"fixed_ip_map": {
																Type:     schema.TypeMap,
																Optional: true,
															},

															"interfaces_addressing_choice": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"automatic_from_end": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"automatic_from_start": {

																			Type:     schema.TypeBool,
																			Optional: true,
																		},

																		"interface_ip_map": {

																			Type:     schema.TypeSet,
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
														},
													},
												},

												"static_ip": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"network_prefix_choice": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"cluster_static_ip": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"interface_ip_map": {

																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"name": {
																									Type:     schema.TypeString,
																									Required: true,
																								},
																								"value": {
																									Type:     schema.TypeSet,
																									Required: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{
																											"default_gw": {
																												Type:     schema.TypeString,
																												Optional: true,
																											},

																											"dns_server": {
																												Type:     schema.TypeString,
																												Optional: true,
																											},

																											"ip_address": {
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

																		"fleet_static_ip": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"default_gw": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"dns_server": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"network_prefix_allocator": {
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

																		"node_static_ip": {

																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"default_gw": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"dns_server": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},

																					"ip_address": {
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

									"device": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"mtu": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"network_choice": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"inside_network": {

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

									"node_choice": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

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

									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"vlan_choice": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

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
								},
							},
						},

						"legacy_interface": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dhcp_server": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"dns_server": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dns_mode": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"dns_server": {

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

									"address_allocator": {
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

									"default_gateway": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"default_gateway_address": {

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

												"default_gateway_mode": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"device_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"dhcp_address": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"mtu": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"static_addresses": {

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

									"tunnel": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tunnel": {
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

									"type": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"virtual_network": {
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

									"vlan_tag": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"vlan_tagging": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"tunnel_interface": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mtu": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"network_choice": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"inside_network": {

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

									"node_choice": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

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

									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"static_ip": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"network_prefix_choice": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"cluster_static_ip": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"interface_ip_map": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"name": {
																						Type:     schema.TypeString,
																						Required: true,
																					},
																					"value": {
																						Type:     schema.TypeSet,
																						Required: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"default_gw": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},

																								"dns_server": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},

																								"ip_address": {
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

															"fleet_static_ip": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"default_gw": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"dns_server": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"network_prefix_allocator": {
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

															"node_static_ip": {

																Type:     schema.TypeSet,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"default_gw": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"dns_server": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},

																		"ip_address": {
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

									"tunnel": {
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
		},
	}
}

// resourceVolterraNetworkInterfaceCreate creates NetworkInterface resource
func resourceVolterraNetworkInterfaceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_network_interface.CreateSpecType{}
	createReq := &ves_io_schema_network_interface.CreateRequest{
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

	if v, ok := d.GetOk("interface_choice"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			oneOfValueFound := false
			var foundKey string
			oneOfKey := set.(map[string]interface{})
			log.Printf("interfaceChoice createSpec")
			for k, v := range oneOfKey {
				if isIntfNil(v) {
					continue
				}
				if !oneOfValueFound {
					log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
					oneOfValueFound = true
					foundKey = k
					continue
				}
				return fmt.Errorf("Only one of interface_choice can be given, dupes found are %v and %v", k, foundKey)
			}
			switch foundKey {

			case "dedicated_interface":
				interfaceChoiceInt := &ves_io_schema_network_interface.CreateSpecType_DedicatedInterface{}
				interfaceChoiceInt.DedicatedInterface = &ves_io_schema_network_interface.DedicatedInterfaceType{}
				createSpec.InterfaceChoice = interfaceChoiceInt

				if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
					sl := v.(*schema.Set).List()
					for _, set := range sl {
						cs := set.(map[string]interface{})

						if v, ok := cs["device"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.DedicatedInterface.Device = v.(string)
						}

						if v, ok := cs["mtu"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.DedicatedInterface.Mtu = uint32(v.(int))
						}

						if v, ok := cs["node_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("nodeChoice interfaceChoiceInt.DedicatedInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of node_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "cluster":
									nodeChoiceInt := &ves_io_schema_network_interface.DedicatedInterfaceType_Cluster{}
									nodeChoiceInt.Cluster = &ves_io_schema.Empty{}
									interfaceChoiceInt.DedicatedInterface.NodeChoice = nodeChoiceInt

								case "node":
									nodeChoiceInt := &ves_io_schema_network_interface.DedicatedInterfaceType_Node{}

									interfaceChoiceInt.DedicatedInterface.NodeChoice = nodeChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										nodeChoiceInt.Node = v.(string)
									}

								}
							}

						}

						if v, ok := cs["priority"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.DedicatedInterface.Priority = uint32(v.(int))
						}

					}
				}

			case "ethernet_interface":
				interfaceChoiceInt := &ves_io_schema_network_interface.CreateSpecType_EthernetInterface{}
				interfaceChoiceInt.EthernetInterface = &ves_io_schema_network_interface.EthernetInterfaceType{}
				createSpec.InterfaceChoice = interfaceChoiceInt

				if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
					sl := v.(*schema.Set).List()
					for _, set := range sl {
						cs := set.(map[string]interface{})

						if v, ok := cs["address_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("addressChoice interfaceChoiceInt.EthernetInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of address_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "dhcp_client":
									addressChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_DhcpClient{}
									addressChoiceInt.DhcpClient = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.AddressChoice = addressChoiceInt

								case "dhcp_server":
									addressChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_DhcpServer{}
									addressChoiceInt.DhcpServer = &ves_io_schema_network_interface.DHCPServerParametersType{}
									interfaceChoiceInt.EthernetInterface.AddressChoice = addressChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["dhcp_networks"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()

												dhcpNetworks := make([]*ves_io_schema_network_interface.DHCPNetworkType, len(sl))
												addressChoiceInt.DhcpServer.DhcpNetworks = dhcpNetworks
												for i, set := range sl {
													dhcpNetworks[i] = &ves_io_schema_network_interface.DHCPNetworkType{}

													dhcpNetworksMapStrToI := set.(map[string]interface{})

													if v, ok := dhcpNetworksMapStrToI["dns_choice"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															oneOfValueFound := false
															var foundKey string
															oneOfKey := set.(map[string]interface{})
															log.Printf("dnsChoice dhcpNetworks[i]")
															for k, v := range oneOfKey {
																if isIntfNil(v) {
																	continue
																}
																if !oneOfValueFound {
																	log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
																	oneOfValueFound = true
																	foundKey = k
																	continue
																}
																return fmt.Errorf("Only one of dns_choice can be given, dupes found are %v and %v", k, foundKey)
															}
															switch foundKey {

															case "dns_address":
																dnsChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_DnsAddress{}

																dhcpNetworks[i].DnsChoice = dnsChoiceInt

																if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
																	dnsChoiceInt.DnsAddress = v.(string)
																}

															case "same_as_dgw":
																dnsChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_SameAsDgw{}
																dnsChoiceInt.SameAsDgw = &ves_io_schema.Empty{}
																dhcpNetworks[i].DnsChoice = dnsChoiceInt

															}
														}

													}

													if v, ok := dhcpNetworksMapStrToI["gateway_choice"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															oneOfValueFound := false
															var foundKey string
															oneOfKey := set.(map[string]interface{})
															log.Printf("gatewayChoice dhcpNetworks[i]")
															for k, v := range oneOfKey {
																if isIntfNil(v) {
																	continue
																}
																if !oneOfValueFound {
																	log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
																	oneOfValueFound = true
																	foundKey = k
																	continue
																}
																return fmt.Errorf("Only one of gateway_choice can be given, dupes found are %v and %v", k, foundKey)
															}
															switch foundKey {

															case "dgw_address":
																gatewayChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_DgwAddress{}

																dhcpNetworks[i].GatewayChoice = gatewayChoiceInt

																if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
																	gatewayChoiceInt.DgwAddress = v.(string)
																}

															case "first_address":
																gatewayChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_FirstAddress{}
																gatewayChoiceInt.FirstAddress = &ves_io_schema.Empty{}
																dhcpNetworks[i].GatewayChoice = gatewayChoiceInt

															case "last_address":
																gatewayChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_LastAddress{}
																gatewayChoiceInt.LastAddress = &ves_io_schema.Empty{}
																dhcpNetworks[i].GatewayChoice = gatewayChoiceInt

															}
														}

													}

													if v, ok := dhcpNetworksMapStrToI["network_prefix_choice"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															oneOfValueFound := false
															var foundKey string
															oneOfKey := set.(map[string]interface{})
															log.Printf("networkPrefixChoice dhcpNetworks[i]")
															for k, v := range oneOfKey {
																if isIntfNil(v) {
																	continue
																}
																if !oneOfValueFound {
																	log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
																	oneOfValueFound = true
																	foundKey = k
																	continue
																}
																return fmt.Errorf("Only one of network_prefix_choice can be given, dupes found are %v and %v", k, foundKey)
															}
															switch foundKey {

															case "network_prefix":
																networkPrefixChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_NetworkPrefix{}

																dhcpNetworks[i].NetworkPrefixChoice = networkPrefixChoiceInt

																if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
																	networkPrefixChoiceInt.NetworkPrefix = v.(string)
																}

															case "network_prefix_allocator":
																networkPrefixChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_NetworkPrefixAllocator{}
																networkPrefixChoiceInt.NetworkPrefixAllocator = &ves_io_schema_views.ObjectRefType{}
																dhcpNetworks[i].NetworkPrefixChoice = networkPrefixChoiceInt

																if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
																	sl := v.(*schema.Set).List()
																	for _, set := range sl {
																		cs := set.(map[string]interface{})

																		if v, ok := cs["name"]; ok && !isIntfNil(v) {

																			networkPrefixChoiceInt.NetworkPrefixAllocator.Name = v.(string)
																		}

																		if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

																			networkPrefixChoiceInt.NetworkPrefixAllocator.Namespace = v.(string)
																		}

																		if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

																			networkPrefixChoiceInt.NetworkPrefixAllocator.Tenant = v.(string)
																		}

																	}
																}

															}
														}

													}

													if v, ok := dhcpNetworksMapStrToI["pool_settings"]; ok && !isIntfNil(v) {

														dhcpNetworks[i].PoolSettings = ves_io_schema_network_interface.DHCPPoolSettingType(ves_io_schema_network_interface.DHCPPoolSettingType_value[v.(string)])

													}

													if v, ok := dhcpNetworksMapStrToI["pools"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()

														pools := make([]*ves_io_schema_network_interface.DHCPPoolType, len(sl))
														dhcpNetworks[i].Pools = pools
														for i, set := range sl {
															pools[i] = &ves_io_schema_network_interface.DHCPPoolType{}

															poolsMapStrToI := set.(map[string]interface{})

															if w, ok := poolsMapStrToI["end_ip"]; ok && !isIntfNil(w) {
																pools[i].EndIp = w.(string)
															}

															if w, ok := poolsMapStrToI["exclude"]; ok && !isIntfNil(w) {
																pools[i].Exclude = w.(bool)
															}

															if w, ok := poolsMapStrToI["start_ip"]; ok && !isIntfNil(w) {
																pools[i].StartIp = w.(string)
															}

														}

													}

												}

											}

											if v, ok := cs["dhcp_option82_tag"]; ok && !isIntfNil(v) {

												addressChoiceInt.DhcpServer.DhcpOption82Tag = v.(string)
											}

											if v, ok := cs["fixed_ip_map"]; ok && !isIntfNil(v) {

												ms := map[string]string{}
												for k, v := range v.(map[string]interface{}) {
													ms[k] = v.(string)
												}
												addressChoiceInt.DhcpServer.FixedIpMap = ms
											}

											if v, ok := cs["interfaces_addressing_choice"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													oneOfValueFound := false
													var foundKey string
													oneOfKey := set.(map[string]interface{})
													log.Printf("interfacesAddressingChoice addressChoiceInt.DhcpServer")
													for k, v := range oneOfKey {
														if isIntfNil(v) {
															continue
														}
														if !oneOfValueFound {
															log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
															oneOfValueFound = true
															foundKey = k
															continue
														}
														return fmt.Errorf("Only one of interfaces_addressing_choice can be given, dupes found are %v and %v", k, foundKey)
													}
													switch foundKey {

													case "automatic_from_end":
														interfacesAddressingChoiceInt := &ves_io_schema_network_interface.DHCPServerParametersType_AutomaticFromEnd{}
														interfacesAddressingChoiceInt.AutomaticFromEnd = &ves_io_schema.Empty{}
														addressChoiceInt.DhcpServer.InterfacesAddressingChoice = interfacesAddressingChoiceInt

													case "automatic_from_start":
														interfacesAddressingChoiceInt := &ves_io_schema_network_interface.DHCPServerParametersType_AutomaticFromStart{}
														interfacesAddressingChoiceInt.AutomaticFromStart = &ves_io_schema.Empty{}
														addressChoiceInt.DhcpServer.InterfacesAddressingChoice = interfacesAddressingChoiceInt

													case "interface_ip_map":
														interfacesAddressingChoiceInt := &ves_io_schema_network_interface.DHCPServerParametersType_InterfaceIpMap{}
														interfacesAddressingChoiceInt.InterfaceIpMap = &ves_io_schema_network_interface.DHCPInterfaceIPType{}
														addressChoiceInt.DhcpServer.InterfacesAddressingChoice = interfacesAddressingChoiceInt

														if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
															sl := v.(*schema.Set).List()
															for _, set := range sl {
																cs := set.(map[string]interface{})

																if v, ok := cs["interface_ip_map"]; ok && !isIntfNil(v) {

																	ms := map[string]string{}
																	for k, v := range v.(map[string]interface{}) {
																		ms[k] = v.(string)
																	}
																	interfacesAddressingChoiceInt.InterfaceIpMap.InterfaceIpMap = ms
																}

															}
														}

													}
												}

											}

										}
									}

								case "static_ip":
									addressChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_StaticIp{}
									addressChoiceInt.StaticIp = &ves_io_schema_network_interface.StaticIPParametersType{}
									interfaceChoiceInt.EthernetInterface.AddressChoice = addressChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["network_prefix_choice"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													oneOfValueFound := false
													var foundKey string
													oneOfKey := set.(map[string]interface{})
													log.Printf("networkPrefixChoice addressChoiceInt.StaticIp")
													for k, v := range oneOfKey {
														if isIntfNil(v) {
															continue
														}
														if !oneOfValueFound {
															log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
															oneOfValueFound = true
															foundKey = k
															continue
														}
														return fmt.Errorf("Only one of network_prefix_choice can be given, dupes found are %v and %v", k, foundKey)
													}
													switch foundKey {

													case "cluster_static_ip":
														networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_ClusterStaticIp{}
														networkPrefixChoiceInt.ClusterStaticIp = &ves_io_schema_network_interface.StaticIpParametersClusterType{}
														addressChoiceInt.StaticIp.NetworkPrefixChoice = networkPrefixChoiceInt

														if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
															sl := v.(*schema.Set).List()
															for _, set := range sl {
																cs := set.(map[string]interface{})

																if v, ok := cs["interface_ip_map"]; ok && !isIntfNil(v) {

																	sl := v.(*schema.Set).List()

																	interfaceIpMap := make(map[string]*ves_io_schema_network_interface.StaticIpParametersNodeType)
																	networkPrefixChoiceInt.ClusterStaticIp.InterfaceIpMap = interfaceIpMap
																	for _, set := range sl {

																		interfaceIpMapMapStrToI := set.(map[string]interface{})
																		key, ok := interfaceIpMapMapStrToI["name"]
																		if ok && !isIntfNil(key) {
																			interfaceIpMap[key.(string)] = &ves_io_schema_network_interface.StaticIpParametersNodeType{}
																			val, _ := interfaceIpMapMapStrToI["value"]
																			staticIPMapVals := val.(*schema.Set).List()
																			for _, staticIP := range staticIPMapVals {
																				staticIPMapStrToI := staticIP.(map[string]interface{})
																				if s, ok := staticIPMapStrToI["default_gw"]; ok && !isIntfNil(s) {
																					interfaceIpMap[key.(string)].DefaultGw = s.(string)
																				}
																				if s, ok := staticIPMapStrToI["dns_server"]; ok && !isIntfNil(s) {
																					interfaceIpMap[key.(string)].DnsServer = s.(string)
																				}

																				if s, ok := staticIPMapStrToI["ip_address"]; ok && !isIntfNil(s) {
																					interfaceIpMap[key.(string)].IpAddress = s.(string)
																				}
																				// break after one loop
																				break
																			}
																		}

																	}

																}

															}
														}

													case "fleet_static_ip":
														networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_FleetStaticIp{}
														networkPrefixChoiceInt.FleetStaticIp = &ves_io_schema_network_interface.StaticIpParametersFleetType{}
														addressChoiceInt.StaticIp.NetworkPrefixChoice = networkPrefixChoiceInt

														if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
															sl := v.(*schema.Set).List()
															for _, set := range sl {
																cs := set.(map[string]interface{})

																if v, ok := cs["default_gw"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.FleetStaticIp.DefaultGw = v.(string)
																}

																if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.FleetStaticIp.DnsServer = v.(string)
																}

																if v, ok := cs["network_prefix_allocator"]; ok && !isIntfNil(v) {

																	sl := v.(*schema.Set).List()

																	networkPrefixAllocator := &ves_io_schema_views.ObjectRefType{}
																	networkPrefixChoiceInt.FleetStaticIp.NetworkPrefixAllocator = networkPrefixAllocator
																	for _, set := range sl {

																		networkPrefixAllocatorMapStrToI := set.(map[string]interface{})

																		if w, ok := networkPrefixAllocatorMapStrToI["name"]; ok && !isIntfNil(w) {
																			networkPrefixAllocator.Name = w.(string)
																		}

																		if w, ok := networkPrefixAllocatorMapStrToI["namespace"]; ok && !isIntfNil(w) {
																			networkPrefixAllocator.Namespace = w.(string)
																		}

																		if w, ok := networkPrefixAllocatorMapStrToI["tenant"]; ok && !isIntfNil(w) {
																			networkPrefixAllocator.Tenant = w.(string)
																		}

																	}

																}

															}
														}

													case "node_static_ip":
														networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_NodeStaticIp{}
														networkPrefixChoiceInt.NodeStaticIp = &ves_io_schema_network_interface.StaticIpParametersNodeType{}
														addressChoiceInt.StaticIp.NetworkPrefixChoice = networkPrefixChoiceInt

														if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
															sl := v.(*schema.Set).List()
															for _, set := range sl {
																cs := set.(map[string]interface{})

																if v, ok := cs["default_gw"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.NodeStaticIp.DefaultGw = v.(string)
																}

																if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.NodeStaticIp.DnsServer = v.(string)
																}

																if v, ok := cs["ip_address"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.NodeStaticIp.IpAddress = v.(string)
																}

															}
														}

													}
												}

											}

										}
									}

								}
							}

						}

						if v, ok := cs["device"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.EthernetInterface.Device = v.(string)
						}

						if v, ok := cs["mtu"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.EthernetInterface.Mtu = uint32(v.(int))
						}

						if v, ok := cs["network_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("networkChoice interfaceChoiceInt.EthernetInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of network_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "inside_network":
									networkChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_InsideNetwork{}
									networkChoiceInt.InsideNetwork = &ves_io_schema_views.ObjectRefType{}
									interfaceChoiceInt.EthernetInterface.NetworkChoice = networkChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["name"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Name = v.(string)
											}

											if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Namespace = v.(string)
											}

											if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Tenant = v.(string)
											}

										}
									}

								case "site_local_inside_network":
									networkChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_SiteLocalInsideNetwork{}
									networkChoiceInt.SiteLocalInsideNetwork = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.NetworkChoice = networkChoiceInt

								case "site_local_network":
									networkChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_SiteLocalNetwork{}
									networkChoiceInt.SiteLocalNetwork = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.NetworkChoice = networkChoiceInt

								}
							}

						}

						if v, ok := cs["node_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("nodeChoice interfaceChoiceInt.EthernetInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of node_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "cluster":
									nodeChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_Cluster{}
									nodeChoiceInt.Cluster = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.NodeChoice = nodeChoiceInt

								case "node":
									nodeChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_Node{}

									interfaceChoiceInt.EthernetInterface.NodeChoice = nodeChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										nodeChoiceInt.Node = v.(string)
									}

								}
							}

						}

						if v, ok := cs["priority"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.EthernetInterface.Priority = uint32(v.(int))
						}

						if v, ok := cs["vlan_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("vlanChoice interfaceChoiceInt.EthernetInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of vlan_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "untagged":
									vlanChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_Untagged{}
									vlanChoiceInt.Untagged = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.VlanChoice = vlanChoiceInt

								case "vlan_id":
									vlanChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_VlanId{}

									interfaceChoiceInt.EthernetInterface.VlanChoice = vlanChoiceInt

								}
							}

						}

					}
				}

			case "legacy_interface":
				interfaceChoiceInt := &ves_io_schema_network_interface.CreateSpecType_LegacyInterface{}
				interfaceChoiceInt.LegacyInterface = &ves_io_schema_network_interface.LegacyInterfaceType{}
				createSpec.InterfaceChoice = interfaceChoiceInt

				if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
					sl := v.(*schema.Set).List()
					for _, set := range sl {
						cs := set.(map[string]interface{})

						if v, ok := cs["dhcp_server"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.DHCPServer = ves_io_schema_network_interface.NetworkInterfaceDHCPServer(ves_io_schema_network_interface.NetworkInterfaceDHCPServer_value[v.(string)])

						}

						if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							dnsServer := &ves_io_schema_network_interface.NetworkInterfaceDNS{}
							interfaceChoiceInt.LegacyInterface.DNSServer = dnsServer
							for _, set := range sl {

								dnsServerMapStrToI := set.(map[string]interface{})

								if v, ok := dnsServerMapStrToI["dns_mode"]; ok && !isIntfNil(v) {

									dnsServer.DnsMode = ves_io_schema_network_interface.NetworkInterfaceDNSMode(ves_io_schema_network_interface.NetworkInterfaceDNSMode_value[v.(string)])

								}

								if v, ok := dnsServerMapStrToI["dns_server"]; ok && !isIntfNil(v) {

									sl := v.(*schema.Set).List()

									dnsServerIpv4s := make([]*ves_io_schema.Ipv4AddressType, len(sl))
									dnsServer.DnsServer = dnsServerIpv4s
									for i, set := range sl {
										dnsServerIpv4s[i] = &ves_io_schema.Ipv4AddressType{}

										dnsServerMapStrToI := set.(map[string]interface{})

										if w, ok := dnsServerMapStrToI["addr"]; ok && !isIntfNil(w) {
											dnsServerIpv4s[i].Addr = w.(string)
										}

									}

								}

							}

						}

						if v, ok := cs["address_allocator"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							addressAllocatorInt := make([]*ves_io_schema.ObjectRefType, len(sl))
							interfaceChoiceInt.LegacyInterface.AddressAllocator = addressAllocatorInt
							for i, ps := range sl {

								aaMapToStrVal := ps.(map[string]interface{})
								addressAllocatorInt[i] = &ves_io_schema.ObjectRefType{}

								addressAllocatorInt[i].Kind = "address_allocator"

								if v, ok := aaMapToStrVal["name"]; ok && !isIntfNil(v) {
									addressAllocatorInt[i].Name = v.(string)
								}

								if v, ok := aaMapToStrVal["namespace"]; ok && !isIntfNil(v) {
									addressAllocatorInt[i].Namespace = v.(string)
								}

								if v, ok := aaMapToStrVal["tenant"]; ok && !isIntfNil(v) {
									addressAllocatorInt[i].Tenant = v.(string)
								}

								if v, ok := aaMapToStrVal["uid"]; ok && !isIntfNil(v) {
									addressAllocatorInt[i].Uid = v.(string)
								}

							}

						}

						if v, ok := cs["default_gateway"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							defaultGateway := &ves_io_schema_network_interface.NetworkInterfaceDFGW{}
							interfaceChoiceInt.LegacyInterface.DefaultGateway = defaultGateway
							for _, set := range sl {

								defaultGatewayMapStrToI := set.(map[string]interface{})

								if v, ok := defaultGatewayMapStrToI["default_gateway_address"]; ok && !isIntfNil(v) {

									sl := v.(*schema.Set).List()

									defaultGatewayAddress := &ves_io_schema.Ipv4AddressType{}
									defaultGateway.DefaultGatewayAddress = defaultGatewayAddress
									for _, set := range sl {

										defaultGatewayAddressMapStrToI := set.(map[string]interface{})

										if w, ok := defaultGatewayAddressMapStrToI["addr"]; ok && !isIntfNil(w) {
											defaultGatewayAddress.Addr = w.(string)
										}

									}

								}

								if v, ok := defaultGatewayMapStrToI["default_gateway_mode"]; ok && !isIntfNil(v) {

									defaultGateway.DefaultGatewayMode = ves_io_schema_network_interface.NetworkInterfaceGatewayMode(ves_io_schema_network_interface.NetworkInterfaceGatewayMode_value[v.(string)])

								}

							}

						}

						if v, ok := cs["device_name"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.DeviceName = v.(string)
						}

						if v, ok := cs["dhcp_address"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.DhcpAddress = ves_io_schema_network_interface.NetworkInterfaceDHCP(ves_io_schema_network_interface.NetworkInterfaceDHCP_value[v.(string)])

						}

						if v, ok := cs["mtu"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.Mtu = uint32(v.(int))
						}

						if v, ok := cs["priority"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.Priority = uint32(v.(int))
						}

						if v, ok := cs["static_addresses"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							staticAddresses := make([]*ves_io_schema.Ipv4SubnetType, len(sl))
							interfaceChoiceInt.LegacyInterface.StaticAddresses = staticAddresses
							for i, set := range sl {
								staticAddresses[i] = &ves_io_schema.Ipv4SubnetType{}

								staticAddressesMapStrToI := set.(map[string]interface{})

								if w, ok := staticAddressesMapStrToI["plen"]; ok && !isIntfNil(w) {
									staticAddresses[i].Plen = w.(uint32)
								}

								if w, ok := staticAddressesMapStrToI["prefix"]; ok && !isIntfNil(w) {
									staticAddresses[i].Prefix = w.(string)
								}

							}

						}

						if v, ok := cs["tunnel"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							tunnel := &ves_io_schema_network_interface.NetworkInterfaceTunnel{}
							interfaceChoiceInt.LegacyInterface.Tunnel = tunnel
							for _, set := range sl {

								tunnelMapStrToI := set.(map[string]interface{})

								if v, ok := tunnelMapStrToI["tunnel"]; ok && !isIntfNil(v) {

									sl := v.(*schema.Set).List()
									tunnelInt := make([]*ves_io_schema.ObjectRefType, len(sl))
									tunnel.Tunnel = tunnelInt
									for i, ps := range sl {

										tMapToStrVal := ps.(map[string]interface{})
										tunnelInt[i] = &ves_io_schema.ObjectRefType{}

										tunnelInt[i].Kind = "tunnel"

										if v, ok := tMapToStrVal["name"]; ok && !isIntfNil(v) {
											tunnelInt[i].Name = v.(string)
										}

										if v, ok := tMapToStrVal["namespace"]; ok && !isIntfNil(v) {
											tunnelInt[i].Namespace = v.(string)
										}

										if v, ok := tMapToStrVal["tenant"]; ok && !isIntfNil(v) {
											tunnelInt[i].Tenant = v.(string)
										}

										if v, ok := tMapToStrVal["uid"]; ok && !isIntfNil(v) {
											tunnelInt[i].Uid = v.(string)
										}

									}

								}

							}

						}

						if v, ok := cs["type"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.Type = ves_io_schema_network_interface.NetworkInterfaceType(ves_io_schema_network_interface.NetworkInterfaceType_value[v.(string)])

						}

						if v, ok := cs["virtual_network"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							virtualNetworkInt := make([]*ves_io_schema.ObjectRefType, len(sl))
							interfaceChoiceInt.LegacyInterface.VirtualNetwork = virtualNetworkInt
							for i, ps := range sl {

								vnMapToStrVal := ps.(map[string]interface{})
								virtualNetworkInt[i] = &ves_io_schema.ObjectRefType{}

								virtualNetworkInt[i].Kind = "virtual_network"

								if v, ok := vnMapToStrVal["name"]; ok && !isIntfNil(v) {
									virtualNetworkInt[i].Name = v.(string)
								}

								if v, ok := vnMapToStrVal["namespace"]; ok && !isIntfNil(v) {
									virtualNetworkInt[i].Namespace = v.(string)
								}

								if v, ok := vnMapToStrVal["tenant"]; ok && !isIntfNil(v) {
									virtualNetworkInt[i].Tenant = v.(string)
								}

								if v, ok := vnMapToStrVal["uid"]; ok && !isIntfNil(v) {
									virtualNetworkInt[i].Uid = v.(string)
								}

							}

						}

						if v, ok := cs["vlan_tag"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.VlanTag = uint32(v.(int))
						}

						if v, ok := cs["vlan_tagging"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.VlanTagging = ves_io_schema_network_interface.NetworkInterfaceVLANTagging(ves_io_schema_network_interface.NetworkInterfaceVLANTagging_value[v.(string)])

						}

					}
				}

			case "tunnel_interface":
				interfaceChoiceInt := &ves_io_schema_network_interface.CreateSpecType_TunnelInterface{}
				interfaceChoiceInt.TunnelInterface = &ves_io_schema_network_interface.TunnelInterfaceType{}
				createSpec.InterfaceChoice = interfaceChoiceInt

				if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
					sl := v.(*schema.Set).List()
					for _, set := range sl {
						cs := set.(map[string]interface{})

						if v, ok := cs["mtu"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.TunnelInterface.Mtu = uint32(v.(int))
						}

						if v, ok := cs["network_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("networkChoice interfaceChoiceInt.TunnelInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of network_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "inside_network":
									networkChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_InsideNetwork{}
									networkChoiceInt.InsideNetwork = &ves_io_schema_views.ObjectRefType{}
									interfaceChoiceInt.TunnelInterface.NetworkChoice = networkChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["name"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Name = v.(string)
											}

											if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Namespace = v.(string)
											}

											if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Tenant = v.(string)
											}

										}
									}

								case "site_local_inside_network":
									networkChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_SiteLocalInsideNetwork{}
									networkChoiceInt.SiteLocalInsideNetwork = &ves_io_schema.Empty{}
									interfaceChoiceInt.TunnelInterface.NetworkChoice = networkChoiceInt

								case "site_local_network":
									networkChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_SiteLocalNetwork{}
									networkChoiceInt.SiteLocalNetwork = &ves_io_schema.Empty{}
									interfaceChoiceInt.TunnelInterface.NetworkChoice = networkChoiceInt

								}
							}

						}

						if v, ok := cs["node_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("nodeChoice interfaceChoiceInt.TunnelInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of node_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "cluster":
									nodeChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_Cluster{}
									nodeChoiceInt.Cluster = &ves_io_schema.Empty{}
									interfaceChoiceInt.TunnelInterface.NodeChoice = nodeChoiceInt

								case "node":
									nodeChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_Node{}

									interfaceChoiceInt.TunnelInterface.NodeChoice = nodeChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										nodeChoiceInt.Node = v.(string)
									}

								}
							}

						}

						if v, ok := cs["priority"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.TunnelInterface.Priority = uint32(v.(int))
						}

						if v, ok := cs["static_ip"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							staticIp := &ves_io_schema_network_interface.StaticIPParametersType{}
							interfaceChoiceInt.TunnelInterface.StaticIp = staticIp
							for _, set := range sl {

								staticIpMapStrToI := set.(map[string]interface{})

								if v, ok := staticIpMapStrToI["network_prefix_choice"]; ok && !isIntfNil(v) {

									sl := v.(*schema.Set).List()
									for _, set := range sl {
										oneOfValueFound := false
										var foundKey string
										oneOfKey := set.(map[string]interface{})
										log.Printf("networkPrefixChoice staticIp")
										for k, v := range oneOfKey {
											if isIntfNil(v) {
												continue
											}
											if !oneOfValueFound {
												log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
												oneOfValueFound = true
												foundKey = k
												continue
											}
											return fmt.Errorf("Only one of network_prefix_choice can be given, dupes found are %v and %v", k, foundKey)
										}
										switch foundKey {

										case "cluster_static_ip":
											networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_ClusterStaticIp{}
											networkPrefixChoiceInt.ClusterStaticIp = &ves_io_schema_network_interface.StaticIpParametersClusterType{}
											staticIp.NetworkPrefixChoice = networkPrefixChoiceInt

											if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["interface_ip_map"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()

														interfaceIpMap := make(map[string]*ves_io_schema_network_interface.StaticIpParametersNodeType)
														networkPrefixChoiceInt.ClusterStaticIp.InterfaceIpMap = interfaceIpMap
														for _, set := range sl {

															interfaceIpMapMapStrToI := set.(map[string]interface{})
															key, ok := interfaceIpMapMapStrToI["name"]
															if ok && !isIntfNil(key) {
																interfaceIpMap[key.(string)] = &ves_io_schema_network_interface.StaticIpParametersNodeType{}
																val, _ := interfaceIpMapMapStrToI["value"]
																staticIPMapVals := val.(*schema.Set).List()
																for _, staticIP := range staticIPMapVals {
																	staticIPMapStrToI := staticIP.(map[string]interface{})
																	if s, ok := staticIPMapStrToI["default_gw"]; ok && !isIntfNil(s) {
																		interfaceIpMap[key.(string)].DefaultGw = s.(string)
																	}
																	if s, ok := staticIPMapStrToI["dns_server"]; ok && !isIntfNil(s) {
																		interfaceIpMap[key.(string)].DnsServer = s.(string)
																	}

																	if s, ok := staticIPMapStrToI["ip_address"]; ok && !isIntfNil(s) {
																		interfaceIpMap[key.(string)].IpAddress = s.(string)
																	}
																	// break after one loop
																	break
																}
															}
														}

													}

												}
											}

										case "fleet_static_ip":
											networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_FleetStaticIp{}
											networkPrefixChoiceInt.FleetStaticIp = &ves_io_schema_network_interface.StaticIpParametersFleetType{}
											staticIp.NetworkPrefixChoice = networkPrefixChoiceInt

											if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["default_gw"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.FleetStaticIp.DefaultGw = v.(string)
													}

													if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.FleetStaticIp.DnsServer = v.(string)
													}

													if v, ok := cs["network_prefix_allocator"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()

														networkPrefixAllocator := &ves_io_schema_views.ObjectRefType{}
														networkPrefixChoiceInt.FleetStaticIp.NetworkPrefixAllocator = networkPrefixAllocator
														for _, set := range sl {

															networkPrefixAllocatorMapStrToI := set.(map[string]interface{})

															if w, ok := networkPrefixAllocatorMapStrToI["name"]; ok && !isIntfNil(w) {
																networkPrefixAllocator.Name = w.(string)
															}

															if w, ok := networkPrefixAllocatorMapStrToI["namespace"]; ok && !isIntfNil(w) {
																networkPrefixAllocator.Namespace = w.(string)
															}

															if w, ok := networkPrefixAllocatorMapStrToI["tenant"]; ok && !isIntfNil(w) {
																networkPrefixAllocator.Tenant = w.(string)
															}

														}

													}

												}
											}

										case "node_static_ip":
											networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_NodeStaticIp{}
											networkPrefixChoiceInt.NodeStaticIp = &ves_io_schema_network_interface.StaticIpParametersNodeType{}
											staticIp.NetworkPrefixChoice = networkPrefixChoiceInt

											if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["default_gw"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.NodeStaticIp.DefaultGw = v.(string)
													}

													if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.NodeStaticIp.DnsServer = v.(string)
													}

													if v, ok := cs["ip_address"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.NodeStaticIp.IpAddress = v.(string)
													}

												}
											}

										}
									}

								}

							}

						}

						if v, ok := cs["tunnel"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							tunnel := &ves_io_schema_views.ObjectRefType{}
							interfaceChoiceInt.TunnelInterface.Tunnel = tunnel
							for _, set := range sl {

								tunnelMapStrToI := set.(map[string]interface{})

								if w, ok := tunnelMapStrToI["name"]; ok && !isIntfNil(w) {
									tunnel.Name = w.(string)
								}

								if w, ok := tunnelMapStrToI["namespace"]; ok && !isIntfNil(w) {
									tunnel.Namespace = w.(string)
								}

								if w, ok := tunnelMapStrToI["tenant"]; ok && !isIntfNil(w) {
									tunnel.Tenant = w.(string)
								}

							}

						}

					}
				}

			}
		}

	}

	log.Printf("[DEBUG] Creating Volterra NetworkInterface object with struct: %+v", createReq)

	createNetworkInterfaceResp, err := client.CreateObject(context.Background(), ves_io_schema_network_interface.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating NetworkInterface: %s", err)
	}
	d.SetId(createNetworkInterfaceResp.GetObjMetadata().GetUid())

	return resourceVolterraNetworkInterfaceRead(d, meta)
}

func resourceVolterraNetworkInterfaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_network_interface.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkInterface %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkInterface %q: %s", d.Id(), err)
	}
	return setNetworkInterfaceFields(client, d, resp)
}

func setNetworkInterfaceFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	d.SetId(metadata.GetUid())

	return nil
}

// resourceVolterraNetworkInterfaceUpdate updates NetworkInterface resource
func resourceVolterraNetworkInterfaceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_network_interface.ReplaceSpecType{}
	updateReq := &ves_io_schema_network_interface.ReplaceRequest{
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

	if v, ok := d.GetOk("interface_choice"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			oneOfValueFound := false
			var foundKey string
			oneOfKey := set.(map[string]interface{})
			log.Printf("interfaceChoice updateSpec")
			for k, v := range oneOfKey {
				if isIntfNil(v) {
					continue
				}
				if !oneOfValueFound {
					log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
					oneOfValueFound = true
					foundKey = k
					continue
				}
				return fmt.Errorf("Only one of interface_choice can be given, dupes found are %v and %v", k, foundKey)
			}
			switch foundKey {

			case "dedicated_interface":
				interfaceChoiceInt := &ves_io_schema_network_interface.ReplaceSpecType_DedicatedInterface{}
				interfaceChoiceInt.DedicatedInterface = &ves_io_schema_network_interface.DedicatedInterfaceType{}
				updateSpec.InterfaceChoice = interfaceChoiceInt

				if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
					sl := v.(*schema.Set).List()
					for _, set := range sl {
						cs := set.(map[string]interface{})

						if v, ok := cs["device"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.DedicatedInterface.Device = v.(string)
						}

						if v, ok := cs["mtu"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.DedicatedInterface.Mtu = uint32(v.(int))
						}

						if v, ok := cs["node_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("nodeChoice interfaceChoiceInt.DedicatedInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of node_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "cluster":
									nodeChoiceInt := &ves_io_schema_network_interface.DedicatedInterfaceType_Cluster{}
									nodeChoiceInt.Cluster = &ves_io_schema.Empty{}
									interfaceChoiceInt.DedicatedInterface.NodeChoice = nodeChoiceInt

								case "node":
									nodeChoiceInt := &ves_io_schema_network_interface.DedicatedInterfaceType_Node{}

									interfaceChoiceInt.DedicatedInterface.NodeChoice = nodeChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										nodeChoiceInt.Node = v.(string)
									}

								}
							}

						}

						if v, ok := cs["priority"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.DedicatedInterface.Priority = uint32(v.(int))
						}

					}
				}

			case "ethernet_interface":
				interfaceChoiceInt := &ves_io_schema_network_interface.ReplaceSpecType_EthernetInterface{}
				interfaceChoiceInt.EthernetInterface = &ves_io_schema_network_interface.EthernetInterfaceType{}
				updateSpec.InterfaceChoice = interfaceChoiceInt

				if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
					sl := v.(*schema.Set).List()
					for _, set := range sl {
						cs := set.(map[string]interface{})

						if v, ok := cs["address_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("addressChoice interfaceChoiceInt.EthernetInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of address_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "dhcp_client":
									addressChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_DhcpClient{}
									addressChoiceInt.DhcpClient = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.AddressChoice = addressChoiceInt

								case "dhcp_server":
									addressChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_DhcpServer{}
									addressChoiceInt.DhcpServer = &ves_io_schema_network_interface.DHCPServerParametersType{}
									interfaceChoiceInt.EthernetInterface.AddressChoice = addressChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["dhcp_networks"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()

												dhcpNetworks := make([]*ves_io_schema_network_interface.DHCPNetworkType, len(sl))
												addressChoiceInt.DhcpServer.DhcpNetworks = dhcpNetworks
												for i, set := range sl {
													dhcpNetworks[i] = &ves_io_schema_network_interface.DHCPNetworkType{}

													dhcpNetworksMapStrToI := set.(map[string]interface{})

													if v, ok := dhcpNetworksMapStrToI["dns_choice"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															oneOfValueFound := false
															var foundKey string
															oneOfKey := set.(map[string]interface{})
															log.Printf("dnsChoice dhcpNetworks[i]")
															for k, v := range oneOfKey {
																if isIntfNil(v) {
																	continue
																}
																if !oneOfValueFound {
																	log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
																	oneOfValueFound = true
																	foundKey = k
																	continue
																}
																return fmt.Errorf("Only one of dns_choice can be given, dupes found are %v and %v", k, foundKey)
															}
															switch foundKey {

															case "dns_address":
																dnsChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_DnsAddress{}

																dhcpNetworks[i].DnsChoice = dnsChoiceInt

																if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
																	dnsChoiceInt.DnsAddress = v.(string)
																}

															case "same_as_dgw":
																dnsChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_SameAsDgw{}
																dnsChoiceInt.SameAsDgw = &ves_io_schema.Empty{}
																dhcpNetworks[i].DnsChoice = dnsChoiceInt

															}
														}

													}

													if v, ok := dhcpNetworksMapStrToI["gateway_choice"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															oneOfValueFound := false
															var foundKey string
															oneOfKey := set.(map[string]interface{})
															log.Printf("gatewayChoice dhcpNetworks[i]")
															for k, v := range oneOfKey {
																if isIntfNil(v) {
																	continue
																}
																if !oneOfValueFound {
																	log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
																	oneOfValueFound = true
																	foundKey = k
																	continue
																}
																return fmt.Errorf("Only one of gateway_choice can be given, dupes found are %v and %v", k, foundKey)
															}
															switch foundKey {

															case "dgw_address":
																gatewayChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_DgwAddress{}

																dhcpNetworks[i].GatewayChoice = gatewayChoiceInt

																if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
																	gatewayChoiceInt.DgwAddress = v.(string)
																}

															case "first_address":
																gatewayChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_FirstAddress{}
																gatewayChoiceInt.FirstAddress = &ves_io_schema.Empty{}
																dhcpNetworks[i].GatewayChoice = gatewayChoiceInt

															case "last_address":
																gatewayChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_LastAddress{}
																gatewayChoiceInt.LastAddress = &ves_io_schema.Empty{}
																dhcpNetworks[i].GatewayChoice = gatewayChoiceInt

															}
														}

													}

													if v, ok := dhcpNetworksMapStrToI["network_prefix_choice"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()
														for _, set := range sl {
															oneOfValueFound := false
															var foundKey string
															oneOfKey := set.(map[string]interface{})
															log.Printf("networkPrefixChoice dhcpNetworks[i]")
															for k, v := range oneOfKey {
																if isIntfNil(v) {
																	continue
																}
																if !oneOfValueFound {
																	log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
																	oneOfValueFound = true
																	foundKey = k
																	continue
																}
																return fmt.Errorf("Only one of network_prefix_choice can be given, dupes found are %v and %v", k, foundKey)
															}
															switch foundKey {

															case "network_prefix":
																networkPrefixChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_NetworkPrefix{}

																dhcpNetworks[i].NetworkPrefixChoice = networkPrefixChoiceInt

																if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
																	networkPrefixChoiceInt.NetworkPrefix = v.(string)
																}

															case "network_prefix_allocator":
																networkPrefixChoiceInt := &ves_io_schema_network_interface.DHCPNetworkType_NetworkPrefixAllocator{}
																networkPrefixChoiceInt.NetworkPrefixAllocator = &ves_io_schema_views.ObjectRefType{}
																dhcpNetworks[i].NetworkPrefixChoice = networkPrefixChoiceInt

																if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
																	sl := v.(*schema.Set).List()
																	for _, set := range sl {
																		cs := set.(map[string]interface{})

																		if v, ok := cs["name"]; ok && !isIntfNil(v) {

																			networkPrefixChoiceInt.NetworkPrefixAllocator.Name = v.(string)
																		}

																		if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

																			networkPrefixChoiceInt.NetworkPrefixAllocator.Namespace = v.(string)
																		}

																		if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

																			networkPrefixChoiceInt.NetworkPrefixAllocator.Tenant = v.(string)
																		}

																	}
																}

															}
														}

													}

													if v, ok := dhcpNetworksMapStrToI["pool_settings"]; ok && !isIntfNil(v) {

														dhcpNetworks[i].PoolSettings = ves_io_schema_network_interface.DHCPPoolSettingType(ves_io_schema_network_interface.DHCPPoolSettingType_value[v.(string)])

													}

													if v, ok := dhcpNetworksMapStrToI["pools"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()

														pools := make([]*ves_io_schema_network_interface.DHCPPoolType, len(sl))
														dhcpNetworks[i].Pools = pools
														for i, set := range sl {
															pools[i] = &ves_io_schema_network_interface.DHCPPoolType{}

															poolsMapStrToI := set.(map[string]interface{})

															if w, ok := poolsMapStrToI["end_ip"]; ok && !isIntfNil(w) {
																pools[i].EndIp = w.(string)
															}

															if w, ok := poolsMapStrToI["exclude"]; ok && !isIntfNil(w) {
																pools[i].Exclude = w.(bool)
															}

															if w, ok := poolsMapStrToI["start_ip"]; ok && !isIntfNil(w) {
																pools[i].StartIp = w.(string)
															}

														}

													}

												}

											}

											if v, ok := cs["dhcp_option82_tag"]; ok && !isIntfNil(v) {

												addressChoiceInt.DhcpServer.DhcpOption82Tag = v.(string)
											}

											if v, ok := cs["fixed_ip_map"]; ok && !isIntfNil(v) {

												ms := map[string]string{}
												for k, v := range v.(map[string]interface{}) {
													ms[k] = v.(string)
												}
												addressChoiceInt.DhcpServer.FixedIpMap = ms
											}

											if v, ok := cs["interfaces_addressing_choice"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													oneOfValueFound := false
													var foundKey string
													oneOfKey := set.(map[string]interface{})
													log.Printf("interfacesAddressingChoice addressChoiceInt.DhcpServer")
													for k, v := range oneOfKey {
														if isIntfNil(v) {
															continue
														}
														if !oneOfValueFound {
															log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
															oneOfValueFound = true
															foundKey = k
															continue
														}
														return fmt.Errorf("Only one of interfaces_addressing_choice can be given, dupes found are %v and %v", k, foundKey)
													}
													switch foundKey {

													case "automatic_from_end":
														interfacesAddressingChoiceInt := &ves_io_schema_network_interface.DHCPServerParametersType_AutomaticFromEnd{}
														interfacesAddressingChoiceInt.AutomaticFromEnd = &ves_io_schema.Empty{}
														addressChoiceInt.DhcpServer.InterfacesAddressingChoice = interfacesAddressingChoiceInt

													case "automatic_from_start":
														interfacesAddressingChoiceInt := &ves_io_schema_network_interface.DHCPServerParametersType_AutomaticFromStart{}
														interfacesAddressingChoiceInt.AutomaticFromStart = &ves_io_schema.Empty{}
														addressChoiceInt.DhcpServer.InterfacesAddressingChoice = interfacesAddressingChoiceInt

													case "interface_ip_map":
														interfacesAddressingChoiceInt := &ves_io_schema_network_interface.DHCPServerParametersType_InterfaceIpMap{}
														interfacesAddressingChoiceInt.InterfaceIpMap = &ves_io_schema_network_interface.DHCPInterfaceIPType{}
														addressChoiceInt.DhcpServer.InterfacesAddressingChoice = interfacesAddressingChoiceInt

														if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
															sl := v.(*schema.Set).List()
															for _, set := range sl {
																cs := set.(map[string]interface{})

																if v, ok := cs["interface_ip_map"]; ok && !isIntfNil(v) {

																	ms := map[string]string{}
																	for k, v := range v.(map[string]interface{}) {
																		ms[k] = v.(string)
																	}
																	interfacesAddressingChoiceInt.InterfaceIpMap.InterfaceIpMap = ms
																}

															}
														}

													}
												}

											}

										}
									}

								case "static_ip":
									addressChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_StaticIp{}
									addressChoiceInt.StaticIp = &ves_io_schema_network_interface.StaticIPParametersType{}
									interfaceChoiceInt.EthernetInterface.AddressChoice = addressChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["network_prefix_choice"]; ok && !isIntfNil(v) {

												sl := v.(*schema.Set).List()
												for _, set := range sl {
													oneOfValueFound := false
													var foundKey string
													oneOfKey := set.(map[string]interface{})
													log.Printf("networkPrefixChoice addressChoiceInt.StaticIp")
													for k, v := range oneOfKey {
														if isIntfNil(v) {
															continue
														}
														if !oneOfValueFound {
															log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
															oneOfValueFound = true
															foundKey = k
															continue
														}
														return fmt.Errorf("Only one of network_prefix_choice can be given, dupes found are %v and %v", k, foundKey)
													}
													switch foundKey {

													case "cluster_static_ip":
														networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_ClusterStaticIp{}
														networkPrefixChoiceInt.ClusterStaticIp = &ves_io_schema_network_interface.StaticIpParametersClusterType{}
														addressChoiceInt.StaticIp.NetworkPrefixChoice = networkPrefixChoiceInt

														if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
															sl := v.(*schema.Set).List()
															for _, set := range sl {
																cs := set.(map[string]interface{})

																if v, ok := cs["interface_ip_map"]; ok && !isIntfNil(v) {

																	sl := v.(*schema.Set).List()

																	interfaceIpMap := make(map[string]*ves_io_schema_network_interface.StaticIpParametersNodeType)
																	networkPrefixChoiceInt.ClusterStaticIp.InterfaceIpMap = interfaceIpMap
																	for _, set := range sl {

																		interfaceIpMapMapStrToI := set.(map[string]interface{})
																		key, ok := interfaceIpMapMapStrToI["name"]
																		if ok && !isIntfNil(key) {
																			interfaceIpMap[key.(string)] = &ves_io_schema_network_interface.StaticIpParametersNodeType{}
																			val, _ := interfaceIpMapMapStrToI["value"]
																			staticIPMapVals := val.(*schema.Set).List()
																			for _, staticIP := range staticIPMapVals {
																				staticIPMapStrToI := staticIP.(map[string]interface{})
																				if s, ok := staticIPMapStrToI["default_gw"]; ok && !isIntfNil(s) {
																					interfaceIpMap[key.(string)].DefaultGw = s.(string)
																				}
																				if s, ok := staticIPMapStrToI["dns_server"]; ok && !isIntfNil(s) {
																					interfaceIpMap[key.(string)].DnsServer = s.(string)
																				}

																				if s, ok := staticIPMapStrToI["ip_address"]; ok && !isIntfNil(s) {
																					interfaceIpMap[key.(string)].IpAddress = s.(string)
																				}
																				// break after one loop
																				break
																			}
																		}

																	}

																}

															}
														}

													case "fleet_static_ip":
														networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_FleetStaticIp{}
														networkPrefixChoiceInt.FleetStaticIp = &ves_io_schema_network_interface.StaticIpParametersFleetType{}
														addressChoiceInt.StaticIp.NetworkPrefixChoice = networkPrefixChoiceInt

														if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
															sl := v.(*schema.Set).List()
															for _, set := range sl {
																cs := set.(map[string]interface{})

																if v, ok := cs["default_gw"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.FleetStaticIp.DefaultGw = v.(string)
																}

																if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.FleetStaticIp.DnsServer = v.(string)
																}

																if v, ok := cs["network_prefix_allocator"]; ok && !isIntfNil(v) {

																	sl := v.(*schema.Set).List()

																	networkPrefixAllocator := &ves_io_schema_views.ObjectRefType{}
																	networkPrefixChoiceInt.FleetStaticIp.NetworkPrefixAllocator = networkPrefixAllocator
																	for _, set := range sl {

																		networkPrefixAllocatorMapStrToI := set.(map[string]interface{})

																		if w, ok := networkPrefixAllocatorMapStrToI["name"]; ok && !isIntfNil(w) {
																			networkPrefixAllocator.Name = w.(string)
																		}

																		if w, ok := networkPrefixAllocatorMapStrToI["namespace"]; ok && !isIntfNil(w) {
																			networkPrefixAllocator.Namespace = w.(string)
																		}

																		if w, ok := networkPrefixAllocatorMapStrToI["tenant"]; ok && !isIntfNil(w) {
																			networkPrefixAllocator.Tenant = w.(string)
																		}

																	}

																}

															}
														}

													case "node_static_ip":
														networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_NodeStaticIp{}
														networkPrefixChoiceInt.NodeStaticIp = &ves_io_schema_network_interface.StaticIpParametersNodeType{}
														addressChoiceInt.StaticIp.NetworkPrefixChoice = networkPrefixChoiceInt

														if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
															sl := v.(*schema.Set).List()
															for _, set := range sl {
																cs := set.(map[string]interface{})

																if v, ok := cs["default_gw"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.NodeStaticIp.DefaultGw = v.(string)
																}

																if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.NodeStaticIp.DnsServer = v.(string)
																}

																if v, ok := cs["ip_address"]; ok && !isIntfNil(v) {

																	networkPrefixChoiceInt.NodeStaticIp.IpAddress = v.(string)
																}

															}
														}

													}
												}

											}

										}
									}

								}
							}

						}

						if v, ok := cs["device"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.EthernetInterface.Device = v.(string)
						}

						if v, ok := cs["mtu"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.EthernetInterface.Mtu = uint32(v.(int))
						}

						if v, ok := cs["network_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("networkChoice interfaceChoiceInt.EthernetInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of network_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "inside_network":
									networkChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_InsideNetwork{}
									networkChoiceInt.InsideNetwork = &ves_io_schema_views.ObjectRefType{}
									interfaceChoiceInt.EthernetInterface.NetworkChoice = networkChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["name"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Name = v.(string)
											}

											if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Namespace = v.(string)
											}

											if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Tenant = v.(string)
											}

										}
									}

								case "site_local_inside_network":
									networkChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_SiteLocalInsideNetwork{}
									networkChoiceInt.SiteLocalInsideNetwork = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.NetworkChoice = networkChoiceInt

								case "site_local_network":
									networkChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_SiteLocalNetwork{}
									networkChoiceInt.SiteLocalNetwork = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.NetworkChoice = networkChoiceInt

								}
							}

						}

						if v, ok := cs["node_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("nodeChoice interfaceChoiceInt.EthernetInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of node_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "cluster":
									nodeChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_Cluster{}
									nodeChoiceInt.Cluster = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.NodeChoice = nodeChoiceInt

								case "node":
									nodeChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_Node{}

									interfaceChoiceInt.EthernetInterface.NodeChoice = nodeChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										nodeChoiceInt.Node = v.(string)
									}

								}
							}

						}

						if v, ok := cs["priority"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.EthernetInterface.Priority = uint32(v.(int))
						}

						if v, ok := cs["vlan_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("vlanChoice interfaceChoiceInt.EthernetInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of vlan_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "untagged":
									vlanChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_Untagged{}
									vlanChoiceInt.Untagged = &ves_io_schema.Empty{}
									interfaceChoiceInt.EthernetInterface.VlanChoice = vlanChoiceInt

								case "vlan_id":
									vlanChoiceInt := &ves_io_schema_network_interface.EthernetInterfaceType_VlanId{}

									interfaceChoiceInt.EthernetInterface.VlanChoice = vlanChoiceInt

								}
							}

						}

					}
				}

			case "legacy_interface":
				interfaceChoiceInt := &ves_io_schema_network_interface.ReplaceSpecType_LegacyInterface{}
				interfaceChoiceInt.LegacyInterface = &ves_io_schema_network_interface.LegacyInterfaceType{}
				updateSpec.InterfaceChoice = interfaceChoiceInt

				if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
					sl := v.(*schema.Set).List()
					for _, set := range sl {
						cs := set.(map[string]interface{})

						if v, ok := cs["dhcp_server"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.DHCPServer = ves_io_schema_network_interface.NetworkInterfaceDHCPServer(ves_io_schema_network_interface.NetworkInterfaceDHCPServer_value[v.(string)])

						}

						if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							dnsServer := &ves_io_schema_network_interface.NetworkInterfaceDNS{}
							interfaceChoiceInt.LegacyInterface.DNSServer = dnsServer
							for _, set := range sl {

								dnsServerMapStrToI := set.(map[string]interface{})

								if v, ok := dnsServerMapStrToI["dns_mode"]; ok && !isIntfNil(v) {

									dnsServer.DnsMode = ves_io_schema_network_interface.NetworkInterfaceDNSMode(ves_io_schema_network_interface.NetworkInterfaceDNSMode_value[v.(string)])

								}

								if v, ok := dnsServerMapStrToI["dns_server"]; ok && !isIntfNil(v) {

									sl := v.(*schema.Set).List()

									dnsServerIpv4s := make([]*ves_io_schema.Ipv4AddressType, len(sl))
									dnsServer.DnsServer = dnsServerIpv4s
									for i, set := range sl {
										dnsServerIpv4s[i] = &ves_io_schema.Ipv4AddressType{}

										dnsServerMapStrToI := set.(map[string]interface{})

										if w, ok := dnsServerMapStrToI["addr"]; ok && !isIntfNil(w) {
											dnsServerIpv4s[i].Addr = w.(string)
										}

									}

								}

							}

						}

						if v, ok := cs["address_allocator"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							addressAllocatorInt := make([]*ves_io_schema.ObjectRefType, len(sl))
							interfaceChoiceInt.LegacyInterface.AddressAllocator = addressAllocatorInt
							for i, ps := range sl {

								aaMapToStrVal := ps.(map[string]interface{})
								addressAllocatorInt[i] = &ves_io_schema.ObjectRefType{}

								addressAllocatorInt[i].Kind = "address_allocator"

								if v, ok := aaMapToStrVal["name"]; ok && !isIntfNil(v) {
									addressAllocatorInt[i].Name = v.(string)
								}

								if v, ok := aaMapToStrVal["namespace"]; ok && !isIntfNil(v) {
									addressAllocatorInt[i].Namespace = v.(string)
								}

								if v, ok := aaMapToStrVal["tenant"]; ok && !isIntfNil(v) {
									addressAllocatorInt[i].Tenant = v.(string)
								}

								if v, ok := aaMapToStrVal["uid"]; ok && !isIntfNil(v) {
									addressAllocatorInt[i].Uid = v.(string)
								}

							}

						}

						if v, ok := cs["default_gateway"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							defaultGateway := &ves_io_schema_network_interface.NetworkInterfaceDFGW{}
							interfaceChoiceInt.LegacyInterface.DefaultGateway = defaultGateway
							for _, set := range sl {

								defaultGatewayMapStrToI := set.(map[string]interface{})

								if v, ok := defaultGatewayMapStrToI["default_gateway_address"]; ok && !isIntfNil(v) {

									sl := v.(*schema.Set).List()

									defaultGatewayAddress := &ves_io_schema.Ipv4AddressType{}
									defaultGateway.DefaultGatewayAddress = defaultGatewayAddress
									for _, set := range sl {

										defaultGatewayAddressMapStrToI := set.(map[string]interface{})

										if w, ok := defaultGatewayAddressMapStrToI["addr"]; ok && !isIntfNil(w) {
											defaultGatewayAddress.Addr = w.(string)
										}

									}

								}

								if v, ok := defaultGatewayMapStrToI["default_gateway_mode"]; ok && !isIntfNil(v) {

									defaultGateway.DefaultGatewayMode = ves_io_schema_network_interface.NetworkInterfaceGatewayMode(ves_io_schema_network_interface.NetworkInterfaceGatewayMode_value[v.(string)])

								}

							}

						}

						if v, ok := cs["device_name"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.DeviceName = v.(string)
						}

						if v, ok := cs["dhcp_address"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.DhcpAddress = ves_io_schema_network_interface.NetworkInterfaceDHCP(ves_io_schema_network_interface.NetworkInterfaceDHCP_value[v.(string)])

						}

						if v, ok := cs["mtu"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.Mtu = uint32(v.(int))
						}

						if v, ok := cs["priority"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.Priority = uint32(v.(int))
						}

						if v, ok := cs["static_addresses"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							staticAddresses := make([]*ves_io_schema.Ipv4SubnetType, len(sl))
							interfaceChoiceInt.LegacyInterface.StaticAddresses = staticAddresses
							for i, set := range sl {
								staticAddresses[i] = &ves_io_schema.Ipv4SubnetType{}

								staticAddressesMapStrToI := set.(map[string]interface{})

								if w, ok := staticAddressesMapStrToI["plen"]; ok && !isIntfNil(w) {
									staticAddresses[i].Plen = w.(uint32)
								}

								if w, ok := staticAddressesMapStrToI["prefix"]; ok && !isIntfNil(w) {
									staticAddresses[i].Prefix = w.(string)
								}

							}

						}

						if v, ok := cs["tunnel"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							tunnel := &ves_io_schema_network_interface.NetworkInterfaceTunnel{}
							interfaceChoiceInt.LegacyInterface.Tunnel = tunnel
							for _, set := range sl {

								tunnelMapStrToI := set.(map[string]interface{})

								if v, ok := tunnelMapStrToI["tunnel"]; ok && !isIntfNil(v) {

									sl := v.(*schema.Set).List()
									tunnelInt := make([]*ves_io_schema.ObjectRefType, len(sl))
									tunnel.Tunnel = tunnelInt
									for i, ps := range sl {

										tMapToStrVal := ps.(map[string]interface{})
										tunnelInt[i] = &ves_io_schema.ObjectRefType{}

										tunnelInt[i].Kind = "tunnel"

										if v, ok := tMapToStrVal["name"]; ok && !isIntfNil(v) {
											tunnelInt[i].Name = v.(string)
										}

										if v, ok := tMapToStrVal["namespace"]; ok && !isIntfNil(v) {
											tunnelInt[i].Namespace = v.(string)
										}

										if v, ok := tMapToStrVal["tenant"]; ok && !isIntfNil(v) {
											tunnelInt[i].Tenant = v.(string)
										}

										if v, ok := tMapToStrVal["uid"]; ok && !isIntfNil(v) {
											tunnelInt[i].Uid = v.(string)
										}

									}

								}

							}

						}

						if v, ok := cs["type"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.Type = ves_io_schema_network_interface.NetworkInterfaceType(ves_io_schema_network_interface.NetworkInterfaceType_value[v.(string)])

						}

						if v, ok := cs["virtual_network"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							virtualNetworkInt := make([]*ves_io_schema.ObjectRefType, len(sl))
							interfaceChoiceInt.LegacyInterface.VirtualNetwork = virtualNetworkInt
							for i, ps := range sl {

								vnMapToStrVal := ps.(map[string]interface{})
								virtualNetworkInt[i] = &ves_io_schema.ObjectRefType{}

								virtualNetworkInt[i].Kind = "virtual_network"

								if v, ok := vnMapToStrVal["name"]; ok && !isIntfNil(v) {
									virtualNetworkInt[i].Name = v.(string)
								}

								if v, ok := vnMapToStrVal["namespace"]; ok && !isIntfNil(v) {
									virtualNetworkInt[i].Namespace = v.(string)
								}

								if v, ok := vnMapToStrVal["tenant"]; ok && !isIntfNil(v) {
									virtualNetworkInt[i].Tenant = v.(string)
								}

								if v, ok := vnMapToStrVal["uid"]; ok && !isIntfNil(v) {
									virtualNetworkInt[i].Uid = v.(string)
								}

							}

						}

						if v, ok := cs["vlan_tag"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.VlanTag = uint32(v.(int))
						}

						if v, ok := cs["vlan_tagging"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.LegacyInterface.VlanTagging = ves_io_schema_network_interface.NetworkInterfaceVLANTagging(ves_io_schema_network_interface.NetworkInterfaceVLANTagging_value[v.(string)])

						}

					}
				}

			case "tunnel_interface":
				interfaceChoiceInt := &ves_io_schema_network_interface.ReplaceSpecType_TunnelInterface{}
				interfaceChoiceInt.TunnelInterface = &ves_io_schema_network_interface.TunnelInterfaceType{}
				updateSpec.InterfaceChoice = interfaceChoiceInt

				if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
					sl := v.(*schema.Set).List()
					for _, set := range sl {
						cs := set.(map[string]interface{})

						if v, ok := cs["mtu"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.TunnelInterface.Mtu = uint32(v.(int))
						}

						if v, ok := cs["network_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("networkChoice interfaceChoiceInt.TunnelInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of network_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "inside_network":
									networkChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_InsideNetwork{}
									networkChoiceInt.InsideNetwork = &ves_io_schema_views.ObjectRefType{}
									interfaceChoiceInt.TunnelInterface.NetworkChoice = networkChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										sl := v.(*schema.Set).List()
										for _, set := range sl {
											cs := set.(map[string]interface{})

											if v, ok := cs["name"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Name = v.(string)
											}

											if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Namespace = v.(string)
											}

											if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

												networkChoiceInt.InsideNetwork.Tenant = v.(string)
											}

										}
									}

								case "site_local_inside_network":
									networkChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_SiteLocalInsideNetwork{}
									networkChoiceInt.SiteLocalInsideNetwork = &ves_io_schema.Empty{}
									interfaceChoiceInt.TunnelInterface.NetworkChoice = networkChoiceInt

								case "site_local_network":
									networkChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_SiteLocalNetwork{}
									networkChoiceInt.SiteLocalNetwork = &ves_io_schema.Empty{}
									interfaceChoiceInt.TunnelInterface.NetworkChoice = networkChoiceInt

								}
							}

						}

						if v, ok := cs["node_choice"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()
							for _, set := range sl {
								oneOfValueFound := false
								var foundKey string
								oneOfKey := set.(map[string]interface{})
								log.Printf("nodeChoice interfaceChoiceInt.TunnelInterface")
								for k, v := range oneOfKey {
									if isIntfNil(v) {
										continue
									}
									if !oneOfValueFound {
										log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
										oneOfValueFound = true
										foundKey = k
										continue
									}
									return fmt.Errorf("Only one of node_choice can be given, dupes found are %v and %v", k, foundKey)
								}
								switch foundKey {

								case "cluster":
									nodeChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_Cluster{}
									nodeChoiceInt.Cluster = &ves_io_schema.Empty{}
									interfaceChoiceInt.TunnelInterface.NodeChoice = nodeChoiceInt

								case "node":
									nodeChoiceInt := &ves_io_schema_network_interface.TunnelInterfaceType_Node{}

									interfaceChoiceInt.TunnelInterface.NodeChoice = nodeChoiceInt

									if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
										nodeChoiceInt.Node = v.(string)
									}

								}
							}

						}

						if v, ok := cs["priority"]; ok && !isIntfNil(v) {

							interfaceChoiceInt.TunnelInterface.Priority = uint32(v.(int))
						}

						if v, ok := cs["static_ip"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							staticIp := &ves_io_schema_network_interface.StaticIPParametersType{}
							interfaceChoiceInt.TunnelInterface.StaticIp = staticIp
							for _, set := range sl {

								staticIpMapStrToI := set.(map[string]interface{})

								if v, ok := staticIpMapStrToI["network_prefix_choice"]; ok && !isIntfNil(v) {

									sl := v.(*schema.Set).List()
									for _, set := range sl {
										oneOfValueFound := false
										var foundKey string
										oneOfKey := set.(map[string]interface{})
										log.Printf("networkPrefixChoice staticIp")
										for k, v := range oneOfKey {
											if isIntfNil(v) {
												continue
											}
											if !oneOfValueFound {
												log.Printf("[DEBUG] key: %s defaultval: %+v ", k, v)
												oneOfValueFound = true
												foundKey = k
												continue
											}
											return fmt.Errorf("Only one of network_prefix_choice can be given, dupes found are %v and %v", k, foundKey)
										}
										switch foundKey {

										case "cluster_static_ip":
											networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_ClusterStaticIp{}
											networkPrefixChoiceInt.ClusterStaticIp = &ves_io_schema_network_interface.StaticIpParametersClusterType{}
											staticIp.NetworkPrefixChoice = networkPrefixChoiceInt

											if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["interface_ip_map"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()

														interfaceIpMap := make(map[string]*ves_io_schema_network_interface.StaticIpParametersNodeType)
														networkPrefixChoiceInt.ClusterStaticIp.InterfaceIpMap = interfaceIpMap
														for _, set := range sl {

															interfaceIpMapMapStrToI := set.(map[string]interface{})
															key, ok := interfaceIpMapMapStrToI["name"]
															if ok && !isIntfNil(key) {
																interfaceIpMap[key.(string)] = &ves_io_schema_network_interface.StaticIpParametersNodeType{}
																val, _ := interfaceIpMapMapStrToI["value"]
																staticIPMapVals := val.(*schema.Set).List()
																for _, staticIP := range staticIPMapVals {
																	staticIPMapStrToI := staticIP.(map[string]interface{})
																	if s, ok := staticIPMapStrToI["default_gw"]; ok && !isIntfNil(s) {
																		interfaceIpMap[key.(string)].DefaultGw = s.(string)
																	}
																	if s, ok := staticIPMapStrToI["dns_server"]; ok && !isIntfNil(s) {
																		interfaceIpMap[key.(string)].DnsServer = s.(string)
																	}

																	if s, ok := staticIPMapStrToI["ip_address"]; ok && !isIntfNil(s) {
																		interfaceIpMap[key.(string)].IpAddress = s.(string)
																	}
																	// break after one loop
																	break
																}
															}

														}

													}

												}
											}

										case "fleet_static_ip":
											networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_FleetStaticIp{}
											networkPrefixChoiceInt.FleetStaticIp = &ves_io_schema_network_interface.StaticIpParametersFleetType{}
											staticIp.NetworkPrefixChoice = networkPrefixChoiceInt

											if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["default_gw"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.FleetStaticIp.DefaultGw = v.(string)
													}

													if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.FleetStaticIp.DnsServer = v.(string)
													}

													if v, ok := cs["network_prefix_allocator"]; ok && !isIntfNil(v) {

														sl := v.(*schema.Set).List()

														networkPrefixAllocator := &ves_io_schema_views.ObjectRefType{}
														networkPrefixChoiceInt.FleetStaticIp.NetworkPrefixAllocator = networkPrefixAllocator
														for _, set := range sl {

															networkPrefixAllocatorMapStrToI := set.(map[string]interface{})

															if w, ok := networkPrefixAllocatorMapStrToI["name"]; ok && !isIntfNil(w) {
																networkPrefixAllocator.Name = w.(string)
															}

															if w, ok := networkPrefixAllocatorMapStrToI["namespace"]; ok && !isIntfNil(w) {
																networkPrefixAllocator.Namespace = w.(string)
															}

															if w, ok := networkPrefixAllocatorMapStrToI["tenant"]; ok && !isIntfNil(w) {
																networkPrefixAllocator.Tenant = w.(string)
															}

														}

													}

												}
											}

										case "node_static_ip":
											networkPrefixChoiceInt := &ves_io_schema_network_interface.StaticIPParametersType_NodeStaticIp{}
											networkPrefixChoiceInt.NodeStaticIp = &ves_io_schema_network_interface.StaticIpParametersNodeType{}
											staticIp.NetworkPrefixChoice = networkPrefixChoiceInt

											if v, ok := oneOfKey[foundKey]; ok && !isIntfNil(v) {
												sl := v.(*schema.Set).List()
												for _, set := range sl {
													cs := set.(map[string]interface{})

													if v, ok := cs["default_gw"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.NodeStaticIp.DefaultGw = v.(string)
													}

													if v, ok := cs["dns_server"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.NodeStaticIp.DnsServer = v.(string)
													}

													if v, ok := cs["ip_address"]; ok && !isIntfNil(v) {

														networkPrefixChoiceInt.NodeStaticIp.IpAddress = v.(string)
													}

												}
											}

										}
									}

								}

							}

						}

						if v, ok := cs["tunnel"]; ok && !isIntfNil(v) {

							sl := v.(*schema.Set).List()

							tunnel := &ves_io_schema_views.ObjectRefType{}
							interfaceChoiceInt.TunnelInterface.Tunnel = tunnel
							for _, set := range sl {

								tunnelMapStrToI := set.(map[string]interface{})

								if w, ok := tunnelMapStrToI["name"]; ok && !isIntfNil(w) {
									tunnel.Name = w.(string)
								}

								if w, ok := tunnelMapStrToI["namespace"]; ok && !isIntfNil(w) {
									tunnel.Namespace = w.(string)
								}

								if w, ok := tunnelMapStrToI["tenant"]; ok && !isIntfNil(w) {
									tunnel.Tenant = w.(string)
								}

							}

						}

					}
				}

			}
		}

	}

	log.Printf("[DEBUG] Updating Volterra NetworkInterface obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_network_interface.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating NetworkInterface: %s", err)
	}

	return resourceVolterraNetworkInterfaceRead(d, meta)
}

func resourceVolterraNetworkInterfaceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_network_interface.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkInterface %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkInterface before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra NetworkInterface obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_network_interface.ObjectType, namespace, name)
}
