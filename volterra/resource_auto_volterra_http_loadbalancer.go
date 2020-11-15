//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policy"
	ves_io_schema_rate_limiter "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/rate_limiter"
	ves_io_schema_route "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/route"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_views_http_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
	ves_io_schema_virtual_host "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
	ves_io_schema_waf_rule_list "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf_rule_list"
)

// resourceVolterraHttpLoadbalancer is implementation of Volterra's HttpLoadbalancer resources
func resourceVolterraHttpLoadbalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraHttpLoadbalancerCreate,
		Read:   resourceVolterraHttpLoadbalancerRead,
		Update: resourceVolterraHttpLoadbalancerUpdate,
		Delete: resourceVolterraHttpLoadbalancerDelete,

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

			"add_location": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"advertise_custom": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"advertise_where": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"site": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"network": {
													Type:     schema.TypeString,
													Optional: true,
												},

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
											},
										},
									},

									"virtual_site": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"network": {
													Type:     schema.TypeString,
													Optional: true,
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

									"vk8s_service": {

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

									"port": {

										Type:     schema.TypeInt,
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

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"public_ip": {

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

			"advertise_on_public_default_vip": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"do_not_advertise": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"captcha_challenge": {

				Type:     schema.TypeSet,
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

						"enable_captcha_challenge": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"js_challenge": {

				Type:     schema.TypeSet,
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

						"enable_js_challenge": {
							Type:     schema.TypeBool,
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

			"cors_policy": {

				Type:     schema.TypeSet,
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

						"max_age": {
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

			"default_route_pools": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cluster": {

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

						"pool": {

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

			"http": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"https": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"add_hsts": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"http_redirect": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"tls_parameters": {

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

												"trusted_ca_url": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
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
				},
			},

			"https_auto_cert": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"add_hsts": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"http_redirect": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"malicious_user_mitigation": {

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

			"more_option": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"buffer_policy": {

							Type:     schema.TypeSet,
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

									"max_request_time": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"compression_params": {

							Type:     schema.TypeSet,
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

						"idle_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"javascript_info": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cache_prefix": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"custom_script_url": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"script_config": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"jwt": {

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

						"max_request_header_size": {
							Type:     schema.TypeInt,
							Optional: true,
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
										Optional: true,
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
										Optional: true,
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

			"disable_rate_limit": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"rate_limit": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_ip_allowed_list": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rate_limiter_allowed_prefixes": {

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

						"ip_allowed_list": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"prefixes": {

										Type: schema.TypeList,

										Required: true,
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

						"rate_limiter": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"total_number": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"unit": {
										Type:     schema.TypeString,
										Optional: true,
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

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"route_ref": {

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

						"direct_response_route": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"http_method": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"path": {

										Type:     schema.TypeSet,
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

									"route_direct_response": {

										Type:     schema.TypeSet,
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

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"http_method": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"path": {

										Type:     schema.TypeSet,
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

									"route_redirect": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"all_params": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"remove_all_params": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"retain_all_params": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"strip_query_params": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"query_params": {

																Type: schema.TypeList,

																Optional: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"host_redirect": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"path_redirect": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"proto_redirect": {
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

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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

									"origin_pools": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cluster": {

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

												"pool": {

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

												"weight": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"path": {

										Type:     schema.TypeSet,
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

			"user_identification": {

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

			"disable_waf": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"waf": {

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

			"waf_rule": {

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

			"waf_exclusion_rules": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"any_domain": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"domain_regex": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"exclude_rule_ids": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"expiration_timestamp": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"methods": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"path_regex": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraHttpLoadbalancerCreate creates HttpLoadbalancer resource
func resourceVolterraHttpLoadbalancerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_http_loadbalancer.CreateSpecType{}
	createReq := &ves_io_schema_views_http_loadbalancer.CreateRequest{
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

	if v, ok := d.GetOk("add_location"); ok && !isIntfNil(v) {

		createSpec.AddLocation =
			v.(bool)
	}

	advertiseChoiceTypeFound := false

	if v, ok := d.GetOk("advertise_custom"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true
		advertiseChoiceInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_AdvertiseCustom{}
		advertiseChoiceInt.AdvertiseCustom = &ves_io_schema_views.AdvertiseCustom{}
		createSpec.AdvertiseChoice = advertiseChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["advertise_where"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				advertiseWhere := make([]*ves_io_schema_views.WhereType, len(sl))
				advertiseChoiceInt.AdvertiseCustom.AdvertiseWhere = advertiseWhere
				for i, set := range sl {
					advertiseWhere[i] = &ves_io_schema_views.WhereType{}

					advertiseWhereMapStrToI := set.(map[string]interface{})

					choiceTypeFound := false

					if v, ok := advertiseWhereMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.WhereType_Site{}
						choiceInt.Site = &ves_io_schema_views.WhereSite{}
						advertiseWhere[i].Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["ip"]; ok && !isIntfNil(v) {

								choiceInt.Site.Ip = v.(string)
							}

							if v, ok := cs["network"]; ok && !isIntfNil(v) {

								choiceInt.Site.Network = ves_io_schema_views.SiteNetwork(ves_io_schema_views.SiteNetwork_value[v.(string)])

							}

							if v, ok := cs["site"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								site := &ves_io_schema_views.ObjectRefType{}
								choiceInt.Site.Site = site
								for _, set := range sl {

									siteMapStrToI := set.(map[string]interface{})

									if w, ok := siteMapStrToI["name"]; ok && !isIntfNil(w) {
										site.Name = w.(string)
									}

									if w, ok := siteMapStrToI["namespace"]; ok && !isIntfNil(w) {
										site.Namespace = w.(string)
									}

									if w, ok := siteMapStrToI["tenant"]; ok && !isIntfNil(w) {
										site.Tenant = w.(string)
									}

								}

							}

						}

					}

					if v, ok := advertiseWhereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.WhereType_VirtualSite{}
						choiceInt.VirtualSite = &ves_io_schema_views.WhereVirtualSite{}
						advertiseWhere[i].Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["network"]; ok && !isIntfNil(v) {

								choiceInt.VirtualSite.Network = ves_io_schema_views.SiteNetwork(ves_io_schema_views.SiteNetwork_value[v.(string)])

							}

							if v, ok := cs["virtual_site"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								virtualSite := &ves_io_schema_views.ObjectRefType{}
								choiceInt.VirtualSite.VirtualSite = virtualSite
								for _, set := range sl {

									virtualSiteMapStrToI := set.(map[string]interface{})

									if w, ok := virtualSiteMapStrToI["name"]; ok && !isIntfNil(w) {
										virtualSite.Name = w.(string)
									}

									if w, ok := virtualSiteMapStrToI["namespace"]; ok && !isIntfNil(w) {
										virtualSite.Namespace = w.(string)
									}

									if w, ok := virtualSiteMapStrToI["tenant"]; ok && !isIntfNil(w) {
										virtualSite.Tenant = w.(string)
									}

								}

							}

						}

					}

					if v, ok := advertiseWhereMapStrToI["vk8s_service"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.WhereType_Vk8SService{}
						choiceInt.Vk8SService = &ves_io_schema_views.WhereVK8SService{}
						advertiseWhere[i].Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := cs["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceIntNew := &ves_io_schema_views.WhereVK8SService_Site{}
								choiceIntNew.Site = &ves_io_schema_views.ObjectRefType{}
								choiceInt.Vk8SService.Choice = choiceIntNew

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceIntNew.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceIntNew.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceIntNew.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := cs["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceIntNew := &ves_io_schema_views.WhereVK8SService_VirtualSite{}
								choiceIntNew.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								choiceInt.Vk8SService.Choice = choiceIntNew

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceIntNew.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceIntNew.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceIntNew.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

					portChoiceTypeFound := false

					if v, ok := advertiseWhereMapStrToI["port"]; ok && !isIntfNil(v) && !portChoiceTypeFound {

						portChoiceTypeFound = true
						portChoiceInt := &ves_io_schema_views.WhereType_Port{}

						advertiseWhere[i].PortChoice = portChoiceInt

						portChoiceInt.Port =
							uint32(v.(int))

					}

					if v, ok := advertiseWhereMapStrToI["use_default_port"]; ok && !isIntfNil(v) && !portChoiceTypeFound {

						portChoiceTypeFound = true

						if v.(bool) {
							portChoiceInt := &ves_io_schema_views.WhereType_UseDefaultPort{}
							portChoiceInt.UseDefaultPort = &ves_io_schema.Empty{}
							advertiseWhere[i].PortChoice = portChoiceInt
						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("advertise_on_public"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true
		advertiseChoiceInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_AdvertiseOnPublic{}
		advertiseChoiceInt.AdvertiseOnPublic = &ves_io_schema_views.AdvertisePublic{}
		createSpec.AdvertiseChoice = advertiseChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["public_ip"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				publicIp := &ves_io_schema_views.ObjectRefType{}
				advertiseChoiceInt.AdvertiseOnPublic.PublicIp = publicIp
				for _, set := range sl {

					publicIpMapStrToI := set.(map[string]interface{})

					if w, ok := publicIpMapStrToI["name"]; ok && !isIntfNil(w) {
						publicIp.Name = w.(string)
					}

					if w, ok := publicIpMapStrToI["namespace"]; ok && !isIntfNil(w) {
						publicIp.Namespace = w.(string)
					}

					if w, ok := publicIpMapStrToI["tenant"]; ok && !isIntfNil(w) {
						publicIp.Tenant = w.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("advertise_on_public_default_vip"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true

		if v.(bool) {
			advertiseChoiceInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_AdvertiseOnPublicDefaultVip{}
			advertiseChoiceInt.AdvertiseOnPublicDefaultVip = &ves_io_schema.Empty{}
			createSpec.AdvertiseChoice = advertiseChoiceInt
		}

	}

	if v, ok := d.GetOk("do_not_advertise"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true

		if v.(bool) {
			advertiseChoiceInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_DoNotAdvertise{}
			advertiseChoiceInt.DoNotAdvertise = &ves_io_schema.Empty{}
			createSpec.AdvertiseChoice = advertiseChoiceInt
		}

	}

	challengeTypeTypeFound := false

	if v, ok := d.GetOk("captcha_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true
		challengeTypeInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_CaptchaChallenge{}
		challengeTypeInt.CaptchaChallenge = &ves_io_schema_virtual_host.CaptchaChallengeType{}
		createSpec.ChallengeType = challengeTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cookie_expiry"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.CookieExpiry = uint32(v.(int))
			}

			if v, ok := cs["custom_page"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.CustomPage = v.(string)
			}

			if v, ok := cs["enable_captcha_challenge"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.EnableCaptchaChallenge = v.(bool)
			}

		}

	}

	if v, ok := d.GetOk("js_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true
		challengeTypeInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_JsChallenge{}
		challengeTypeInt.JsChallenge = &ves_io_schema_virtual_host.JavascriptChallengeType{}
		createSpec.ChallengeType = challengeTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cookie_expiry"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.CookieExpiry = uint32(v.(int))
			}

			if v, ok := cs["custom_page"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.CustomPage = v.(string)
			}

			if v, ok := cs["enable_js_challenge"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.EnableJsChallenge = v.(bool)
			}

			if v, ok := cs["js_script_delay"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.JsScriptDelay = uint32(v.(int))
			}

		}

	}

	if v, ok := d.GetOk("no_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true

		if v.(bool) {
			challengeTypeInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_NoChallenge{}
			challengeTypeInt.NoChallenge = &ves_io_schema.Empty{}
			createSpec.ChallengeType = challengeTypeInt
		}

	}

	if v, ok := d.GetOk("cors_policy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		corsPolicy := &ves_io_schema.CorsPolicy{}
		createSpec.CorsPolicy = corsPolicy
		for _, set := range sl {

			corsPolicyMapStrToI := set.(map[string]interface{})

			if w, ok := corsPolicyMapStrToI["allow_credentials"]; ok && !isIntfNil(w) {
				corsPolicy.AllowCredentials = w.(bool)
			}

			if w, ok := corsPolicyMapStrToI["allow_headers"]; ok && !isIntfNil(w) {
				corsPolicy.AllowHeaders = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["allow_methods"]; ok && !isIntfNil(w) {
				corsPolicy.AllowMethods = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["allow_origin"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				corsPolicy.AllowOrigin = ls
			}

			if w, ok := corsPolicyMapStrToI["allow_origin_regex"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				corsPolicy.AllowOriginRegex = ls
			}

			if w, ok := corsPolicyMapStrToI["disabled"]; ok && !isIntfNil(w) {
				corsPolicy.Disabled = w.(bool)
			}

			if w, ok := corsPolicyMapStrToI["expose_headers"]; ok && !isIntfNil(w) {
				corsPolicy.ExposeHeaders = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["max_age"]; ok && !isIntfNil(w) {
				corsPolicy.MaxAge = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["maximum_age"]; ok && !isIntfNil(w) {
				corsPolicy.MaximumAge = w.(int32)
			}

		}

	}

	if v, ok := d.GetOk("default_route_pools"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		defaultRoutePools := make([]*ves_io_schema_views.OriginPoolWithWeight, len(sl))
		createSpec.DefaultRoutePools = defaultRoutePools
		for i, set := range sl {
			defaultRoutePools[i] = &ves_io_schema_views.OriginPoolWithWeight{}

			defaultRoutePoolsMapStrToI := set.(map[string]interface{})

			poolChoiceTypeFound := false

			if v, ok := defaultRoutePoolsMapStrToI["cluster"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Cluster{}
				poolChoiceInt.Cluster = &ves_io_schema_views.ObjectRefType{}
				defaultRoutePools[i].PoolChoice = poolChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						poolChoiceInt.Cluster.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						poolChoiceInt.Cluster.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						poolChoiceInt.Cluster.Tenant = v.(string)
					}

				}

			}

			if v, ok := defaultRoutePoolsMapStrToI["pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Pool{}
				poolChoiceInt.Pool = &ves_io_schema_views.ObjectRefType{}
				defaultRoutePools[i].PoolChoice = poolChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						poolChoiceInt.Pool.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						poolChoiceInt.Pool.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						poolChoiceInt.Pool.Tenant = v.(string)
					}

				}

			}

			if w, ok := defaultRoutePoolsMapStrToI["weight"]; ok && !isIntfNil(w) {
				defaultRoutePools[i].Weight = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("domains"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		createSpec.Domains = ls

	}

	loadbalancerTypeTypeFound := false

	if v, ok := d.GetOk("http"); ok && !loadbalancerTypeTypeFound {

		loadbalancerTypeTypeFound = true

		if v.(bool) {
			loadbalancerTypeInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_Http{}
			loadbalancerTypeInt.Http = &ves_io_schema.Empty{}
			createSpec.LoadbalancerType = loadbalancerTypeInt
		}

	}

	if v, ok := d.GetOk("https"); ok && !loadbalancerTypeTypeFound {

		loadbalancerTypeTypeFound = true
		loadbalancerTypeInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_Https{}
		loadbalancerTypeInt.Https = &ves_io_schema_views_http_loadbalancer.ProxyTypeHttps{}
		createSpec.LoadbalancerType = loadbalancerTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["add_hsts"]; ok && !isIntfNil(v) {

				loadbalancerTypeInt.Https.AddHsts = v.(bool)
			}

			if v, ok := cs["http_redirect"]; ok && !isIntfNil(v) {

				loadbalancerTypeInt.Https.HttpRedirect = v.(bool)
			}

			if v, ok := cs["tls_parameters"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				tlsParameters := &ves_io_schema_views_http_loadbalancer.DownstreamTlsParamsType{}
				loadbalancerTypeInt.Https.TlsParameters = tlsParameters
				for _, set := range sl {

					tlsParametersMapStrToI := set.(map[string]interface{})

					mtlsChoiceTypeFound := false

					if v, ok := tlsParametersMapStrToI["no_mtls"]; ok && !isIntfNil(v) && !mtlsChoiceTypeFound {

						mtlsChoiceTypeFound = true

						if v.(bool) {
							mtlsChoiceInt := &ves_io_schema_views_http_loadbalancer.DownstreamTlsParamsType_NoMtls{}
							mtlsChoiceInt.NoMtls = &ves_io_schema.Empty{}
							tlsParameters.MtlsChoice = mtlsChoiceInt
						}

					}

					if v, ok := tlsParametersMapStrToI["use_mtls"]; ok && !isIntfNil(v) && !mtlsChoiceTypeFound {

						mtlsChoiceTypeFound = true
						mtlsChoiceInt := &ves_io_schema_views_http_loadbalancer.DownstreamTlsParamsType_UseMtls{}
						mtlsChoiceInt.UseMtls = &ves_io_schema_views_http_loadbalancer.DownstreamTlsValidationContext{}
						tlsParameters.MtlsChoice = mtlsChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["trusted_ca_url"]; ok && !isIntfNil(v) {

								mtlsChoiceInt.UseMtls.TrustedCaUrl = v.(string)
							}

						}

					}

					if v, ok := tlsParametersMapStrToI["tls_certificates"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						tlsCertificates := make([]*ves_io_schema.TlsCertificateType, len(sl))
						tlsParameters.TlsCertificates = tlsCertificates
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

					if v, ok := tlsParametersMapStrToI["tls_config"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						tlsConfig := &ves_io_schema_views.TlsConfig{}
						tlsParameters.TlsConfig = tlsConfig
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

		}

	}

	if v, ok := d.GetOk("https_auto_cert"); ok && !loadbalancerTypeTypeFound {

		loadbalancerTypeTypeFound = true
		loadbalancerTypeInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_HttpsAutoCert{}
		loadbalancerTypeInt.HttpsAutoCert = &ves_io_schema_views_http_loadbalancer.ProxyTypeHttpsAutoCerts{}
		createSpec.LoadbalancerType = loadbalancerTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["add_hsts"]; ok && !isIntfNil(v) {

				loadbalancerTypeInt.HttpsAutoCert.AddHsts = v.(bool)
			}

			if v, ok := cs["http_redirect"]; ok && !isIntfNil(v) {

				loadbalancerTypeInt.HttpsAutoCert.HttpRedirect = v.(bool)
			}

		}

	}

	if v, ok := d.GetOk("malicious_user_mitigation"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		maliciousUserMitigation := &ves_io_schema_views.ObjectRefType{}
		createSpec.MaliciousUserMitigation = maliciousUserMitigation
		for _, set := range sl {

			maliciousUserMitigationMapStrToI := set.(map[string]interface{})

			if w, ok := maliciousUserMitigationMapStrToI["name"]; ok && !isIntfNil(w) {
				maliciousUserMitigation.Name = w.(string)
			}

			if w, ok := maliciousUserMitigationMapStrToI["namespace"]; ok && !isIntfNil(w) {
				maliciousUserMitigation.Namespace = w.(string)
			}

			if w, ok := maliciousUserMitigationMapStrToI["tenant"]; ok && !isIntfNil(w) {
				maliciousUserMitigation.Tenant = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("more_option"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		moreOption := &ves_io_schema_views_http_loadbalancer.AdvancedOptionsType{}
		createSpec.MoreOption = moreOption
		for _, set := range sl {

			moreOptionMapStrToI := set.(map[string]interface{})

			if v, ok := moreOptionMapStrToI["buffer_policy"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				bufferPolicy := &ves_io_schema.BufferConfigType{}
				moreOption.BufferPolicy = bufferPolicy
				for _, set := range sl {

					bufferPolicyMapStrToI := set.(map[string]interface{})

					if w, ok := bufferPolicyMapStrToI["disabled"]; ok && !isIntfNil(w) {
						bufferPolicy.Disabled = w.(bool)
					}

					if w, ok := bufferPolicyMapStrToI["max_request_bytes"]; ok && !isIntfNil(w) {
						bufferPolicy.MaxRequestBytes = w.(uint32)
					}

					if w, ok := bufferPolicyMapStrToI["max_request_time"]; ok && !isIntfNil(w) {
						bufferPolicy.MaxRequestTime = w.(uint32)
					}

				}

			}

			if v, ok := moreOptionMapStrToI["compression_params"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				compressionParams := &ves_io_schema_virtual_host.CompressionType{}
				moreOption.CompressionParams = compressionParams
				for _, set := range sl {

					compressionParamsMapStrToI := set.(map[string]interface{})

					if w, ok := compressionParamsMapStrToI["content_length"]; ok && !isIntfNil(w) {
						compressionParams.ContentLength = w.(uint32)
					}

					if w, ok := compressionParamsMapStrToI["content_type"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						compressionParams.ContentType = ls
					}

					if w, ok := compressionParamsMapStrToI["disable_on_etag_header"]; ok && !isIntfNil(w) {
						compressionParams.DisableOnEtagHeader = w.(bool)
					}

					if w, ok := compressionParamsMapStrToI["remove_accept_encoding_header"]; ok && !isIntfNil(w) {
						compressionParams.RemoveAcceptEncodingHeader = w.(bool)
					}

				}

			}

			if w, ok := moreOptionMapStrToI["custom_errors"]; ok && !isIntfNil(w) {
				ms := map[uint32]string{}
				for k, v := range w.(map[string]interface{}) {
					val := v.(string)

					s, err := strconv.ParseUint(k, 10, 32)
					if err != nil {
						return fmt.Errorf("Error while decrypting custom_errors: %s", err)
					}
					key := uint32(s)
					ms[key] = val
				}
				moreOption.CustomErrors = ms
			}

			if w, ok := moreOptionMapStrToI["idle_timeout"]; ok && !isIntfNil(w) {
				moreOption.IdleTimeout = w.(uint32)
			}

			if v, ok := moreOptionMapStrToI["javascript_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				javascriptInfo := &ves_io_schema_virtual_host.JavaScriptConfigType{}
				moreOption.JavascriptInfo = javascriptInfo
				for _, set := range sl {

					javascriptInfoMapStrToI := set.(map[string]interface{})

					if w, ok := javascriptInfoMapStrToI["cache_prefix"]; ok && !isIntfNil(w) {
						javascriptInfo.CachePrefix = w.(string)
					}

					if w, ok := javascriptInfoMapStrToI["custom_script_url"]; ok && !isIntfNil(w) {
						javascriptInfo.CustomScriptUrl = w.(string)
					}

					jsonFmtString := v.(string)
					jsm := jsonpb.Unmarshaler{}
					if err := jsm.Unmarshal(strings.NewReader(jsonFmtString), javascriptInfo.ScriptConfig); err != nil {
						return err
					}

				}

			}

			if v, ok := moreOptionMapStrToI["jwt"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				jwtInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				moreOption.Jwt = jwtInt
				for i, ps := range sl {

					jMapToStrVal := ps.(map[string]interface{})
					jwtInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := jMapToStrVal["name"]; ok && !isIntfNil(v) {
						jwtInt[i].Name = v.(string)
					}

					if v, ok := jMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						jwtInt[i].Namespace = v.(string)
					}

					if v, ok := jMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						jwtInt[i].Tenant = v.(string)
					}

				}

			}

			if w, ok := moreOptionMapStrToI["max_request_header_size"]; ok && !isIntfNil(w) {
				moreOption.MaxRequestHeaderSize = w.(uint32)
			}

			if v, ok := moreOptionMapStrToI["request_headers_to_add"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				requestHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
				moreOption.RequestHeadersToAdd = requestHeadersToAdd
				for i, set := range sl {
					requestHeadersToAdd[i] = &ves_io_schema.HeaderManipulationOptionType{}

					requestHeadersToAddMapStrToI := set.(map[string]interface{})

					if w, ok := requestHeadersToAddMapStrToI["append"]; ok && !isIntfNil(w) {
						requestHeadersToAdd[i].Append = w.(bool)
					}

					if w, ok := requestHeadersToAddMapStrToI["name"]; ok && !isIntfNil(w) {
						requestHeadersToAdd[i].Name = w.(string)
					}

					if w, ok := requestHeadersToAddMapStrToI["value"]; ok && !isIntfNil(w) {
						requestHeadersToAdd[i].Value = w.(string)
					}

				}

			}

			if w, ok := moreOptionMapStrToI["request_headers_to_remove"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				moreOption.RequestHeadersToRemove = ls
			}

			if v, ok := moreOptionMapStrToI["response_headers_to_add"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				responseHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
				moreOption.ResponseHeadersToAdd = responseHeadersToAdd
				for i, set := range sl {
					responseHeadersToAdd[i] = &ves_io_schema.HeaderManipulationOptionType{}

					responseHeadersToAddMapStrToI := set.(map[string]interface{})

					if w, ok := responseHeadersToAddMapStrToI["append"]; ok && !isIntfNil(w) {
						responseHeadersToAdd[i].Append = w.(bool)
					}

					if w, ok := responseHeadersToAddMapStrToI["name"]; ok && !isIntfNil(w) {
						responseHeadersToAdd[i].Name = w.(string)
					}

					if w, ok := responseHeadersToAddMapStrToI["value"]; ok && !isIntfNil(w) {
						responseHeadersToAdd[i].Value = w.(string)
					}

				}

			}

			if w, ok := moreOptionMapStrToI["response_headers_to_remove"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				moreOption.ResponseHeadersToRemove = ls
			}

		}

	}

	rateLimitChoiceTypeFound := false

	if v, ok := d.GetOk("disable_rate_limit"); ok && !rateLimitChoiceTypeFound {

		rateLimitChoiceTypeFound = true

		if v.(bool) {
			rateLimitChoiceInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_DisableRateLimit{}
			rateLimitChoiceInt.DisableRateLimit = &ves_io_schema.Empty{}
			createSpec.RateLimitChoice = rateLimitChoiceInt
		}

	}

	if v, ok := d.GetOk("rate_limit"); ok && !rateLimitChoiceTypeFound {

		rateLimitChoiceTypeFound = true
		rateLimitChoiceInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_RateLimit{}
		rateLimitChoiceInt.RateLimit = &ves_io_schema_views_http_loadbalancer.RateLimitConfigType{}
		createSpec.RateLimitChoice = rateLimitChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			ipAllowedListChoiceTypeFound := false

			if v, ok := cs["custom_ip_allowed_list"]; ok && !isIntfNil(v) && !ipAllowedListChoiceTypeFound {

				ipAllowedListChoiceTypeFound = true
				ipAllowedListChoiceInt := &ves_io_schema_views_http_loadbalancer.RateLimitConfigType_CustomIpAllowedList{}
				ipAllowedListChoiceInt.CustomIpAllowedList = &ves_io_schema_views_http_loadbalancer.CustomIpAllowedList{}
				rateLimitChoiceInt.RateLimit.IpAllowedListChoice = ipAllowedListChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["rate_limiter_allowed_prefixes"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						rateLimiterAllowedPrefixesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						ipAllowedListChoiceInt.CustomIpAllowedList.RateLimiterAllowedPrefixes = rateLimiterAllowedPrefixesInt
						for i, ps := range sl {

							rlapMapToStrVal := ps.(map[string]interface{})
							rateLimiterAllowedPrefixesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := rlapMapToStrVal["name"]; ok && !isIntfNil(v) {
								rateLimiterAllowedPrefixesInt[i].Name = v.(string)
							}

							if v, ok := rlapMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								rateLimiterAllowedPrefixesInt[i].Namespace = v.(string)
							}

							if v, ok := rlapMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								rateLimiterAllowedPrefixesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["ip_allowed_list"]; ok && !isIntfNil(v) && !ipAllowedListChoiceTypeFound {

				ipAllowedListChoiceTypeFound = true
				ipAllowedListChoiceInt := &ves_io_schema_views_http_loadbalancer.RateLimitConfigType_IpAllowedList{}
				ipAllowedListChoiceInt.IpAllowedList = &ves_io_schema_views.PrefixStringListType{}
				rateLimitChoiceInt.RateLimit.IpAllowedListChoice = ipAllowedListChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						ipAllowedListChoiceInt.IpAllowedList.Prefixes = ls

					}

				}

			}

			if v, ok := cs["no_ip_allowed_list"]; ok && !isIntfNil(v) && !ipAllowedListChoiceTypeFound {

				ipAllowedListChoiceTypeFound = true

				if v.(bool) {
					ipAllowedListChoiceInt := &ves_io_schema_views_http_loadbalancer.RateLimitConfigType_NoIpAllowedList{}
					ipAllowedListChoiceInt.NoIpAllowedList = &ves_io_schema.Empty{}
					rateLimitChoiceInt.RateLimit.IpAllowedListChoice = ipAllowedListChoiceInt
				}

			}

			if v, ok := cs["rate_limiter"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				rateLimiter := &ves_io_schema_rate_limiter.RateLimitValue{}
				rateLimitChoiceInt.RateLimit.RateLimiter = rateLimiter
				for _, set := range sl {

					rateLimiterMapStrToI := set.(map[string]interface{})

					if w, ok := rateLimiterMapStrToI["total_number"]; ok && !isIntfNil(w) {
						rateLimiter.TotalNumber = w.(uint32)
					}

					if v, ok := rateLimiterMapStrToI["unit"]; ok && !isIntfNil(v) {

						rateLimiter.Unit = ves_io_schema_rate_limiter.RateLimitPeriodUnit(ves_io_schema_rate_limiter.RateLimitPeriodUnit_value[v.(string)])

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		routes := make([]*ves_io_schema_views_http_loadbalancer.RouteType, len(sl))
		createSpec.Routes = routes
		for i, set := range sl {
			routes[i] = &ves_io_schema_views_http_loadbalancer.RouteType{}

			routesMapStrToI := set.(map[string]interface{})

			choiceTypeFound := false

			if v, ok := routesMapStrToI["custom_route_object"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_http_loadbalancer.RouteType_CustomRouteObject{}
				choiceInt.CustomRouteObject = &ves_io_schema_views_http_loadbalancer.RouteTypeCustomRoute{}
				routes[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["route_ref"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						routeRef := &ves_io_schema_views.ObjectRefType{}
						choiceInt.CustomRouteObject.RouteRef = routeRef
						for _, set := range sl {

							routeRefMapStrToI := set.(map[string]interface{})

							if w, ok := routeRefMapStrToI["name"]; ok && !isIntfNil(w) {
								routeRef.Name = w.(string)
							}

							if w, ok := routeRefMapStrToI["namespace"]; ok && !isIntfNil(w) {
								routeRef.Namespace = w.(string)
							}

							if w, ok := routeRefMapStrToI["tenant"]; ok && !isIntfNil(w) {
								routeRef.Tenant = w.(string)
							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["direct_response_route"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_http_loadbalancer.RouteType_DirectResponseRoute{}
				choiceInt.DirectResponseRoute = &ves_io_schema_views_http_loadbalancer.RouteTypeDirectResponse{}
				routes[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["http_method"]; ok && !isIntfNil(v) {

						choiceInt.DirectResponseRoute.HttpMethod = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if v, ok := cs["path"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						path := &ves_io_schema.PathMatcherType{}
						choiceInt.DirectResponseRoute.Path = path
						for _, set := range sl {

							pathMapStrToI := set.(map[string]interface{})

							pathMatchTypeFound := false

							if v, ok := pathMapStrToI["path"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Path{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Path = v.(string)

							}

							if v, ok := pathMapStrToI["prefix"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Prefix{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Prefix = v.(string)

							}

							if v, ok := pathMapStrToI["regex"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Regex{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Regex = v.(string)

							}

						}

					}

					if v, ok := cs["route_direct_response"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						routeDirectResponse := &ves_io_schema_route.RouteDirectResponse{}
						choiceInt.DirectResponseRoute.RouteDirectResponse = routeDirectResponse
						for _, set := range sl {

							routeDirectResponseMapStrToI := set.(map[string]interface{})

							if w, ok := routeDirectResponseMapStrToI["response_body"]; ok && !isIntfNil(w) {
								routeDirectResponse.ResponseBody = w.(string)
							}

							if w, ok := routeDirectResponseMapStrToI["response_code"]; ok && !isIntfNil(w) {
								routeDirectResponse.ResponseCode = w.(uint32)
							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["redirect_route"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_http_loadbalancer.RouteType_RedirectRoute{}
				choiceInt.RedirectRoute = &ves_io_schema_views_http_loadbalancer.RouteTypeRedirect{}
				routes[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["http_method"]; ok && !isIntfNil(v) {

						choiceInt.RedirectRoute.HttpMethod = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if v, ok := cs["path"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						path := &ves_io_schema.PathMatcherType{}
						choiceInt.RedirectRoute.Path = path
						for _, set := range sl {

							pathMapStrToI := set.(map[string]interface{})

							pathMatchTypeFound := false

							if v, ok := pathMapStrToI["path"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Path{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Path = v.(string)

							}

							if v, ok := pathMapStrToI["prefix"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Prefix{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Prefix = v.(string)

							}

							if v, ok := pathMapStrToI["regex"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Regex{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Regex = v.(string)

							}

						}

					}

					if v, ok := cs["route_redirect"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						routeRedirect := &ves_io_schema_route.RouteRedirect{}
						choiceInt.RedirectRoute.RouteRedirect = routeRedirect
						for _, set := range sl {

							routeRedirectMapStrToI := set.(map[string]interface{})

							queryParamsTypeFound := false

							if v, ok := routeRedirectMapStrToI["all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

								queryParamsTypeFound = true
								queryParamsInt := &ves_io_schema_route.RouteRedirect_AllParams{}

								routeRedirect.QueryParams = queryParamsInt

								queryParamsInt.AllParams =
									v.(bool)

							}

							if v, ok := routeRedirectMapStrToI["remove_all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

								queryParamsTypeFound = true

								if v.(bool) {
									queryParamsInt := &ves_io_schema_route.RouteRedirect_RemoveAllParams{}
									queryParamsInt.RemoveAllParams = &ves_io_schema.Empty{}
									routeRedirect.QueryParams = queryParamsInt
								}

							}

							if v, ok := routeRedirectMapStrToI["retain_all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

								queryParamsTypeFound = true

								if v.(bool) {
									queryParamsInt := &ves_io_schema_route.RouteRedirect_RetainAllParams{}
									queryParamsInt.RetainAllParams = &ves_io_schema.Empty{}
									routeRedirect.QueryParams = queryParamsInt
								}

							}

							if v, ok := routeRedirectMapStrToI["strip_query_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

								queryParamsTypeFound = true
								queryParamsInt := &ves_io_schema_route.RouteRedirect_StripQueryParams{}
								queryParamsInt.StripQueryParams = &ves_io_schema_route.RouteQueryParams{}
								routeRedirect.QueryParams = queryParamsInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["query_params"]; ok && !isIntfNil(v) {

										ls := make([]string, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {
											ls[i] = v.(string)
										}
										queryParamsInt.StripQueryParams.QueryParams = ls

									}

								}

							}

							if w, ok := routeRedirectMapStrToI["host_redirect"]; ok && !isIntfNil(w) {
								routeRedirect.HostRedirect = w.(string)
							}

							if w, ok := routeRedirectMapStrToI["path_redirect"]; ok && !isIntfNil(w) {
								routeRedirect.PathRedirect = w.(string)
							}

							if w, ok := routeRedirectMapStrToI["proto_redirect"]; ok && !isIntfNil(w) {
								routeRedirect.ProtoRedirect = w.(string)
							}

							if w, ok := routeRedirectMapStrToI["response_code"]; ok && !isIntfNil(w) {
								routeRedirect.ResponseCode = w.(uint32)
							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["simple_route"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_http_loadbalancer.RouteType_SimpleRoute{}
				choiceInt.SimpleRoute = &ves_io_schema_views_http_loadbalancer.RouteTypeSimple{}
				routes[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					hostRewriteParamsTypeFound := false

					if v, ok := cs["auto_host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true

						if v.(bool) {
							hostRewriteParamsInt := &ves_io_schema_views_http_loadbalancer.RouteTypeSimple_AutoHostRewrite{}
							hostRewriteParamsInt.AutoHostRewrite = &ves_io_schema.Empty{}
							choiceInt.SimpleRoute.HostRewriteParams = hostRewriteParamsInt
						}

					}

					if v, ok := cs["disable_host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true

						if v.(bool) {
							hostRewriteParamsInt := &ves_io_schema_views_http_loadbalancer.RouteTypeSimple_DisableHostRewrite{}
							hostRewriteParamsInt.DisableHostRewrite = &ves_io_schema.Empty{}
							choiceInt.SimpleRoute.HostRewriteParams = hostRewriteParamsInt
						}

					}

					if v, ok := cs["host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true
						hostRewriteParamsInt := &ves_io_schema_views_http_loadbalancer.RouteTypeSimple_HostRewrite{}

						choiceInt.SimpleRoute.HostRewriteParams = hostRewriteParamsInt

						hostRewriteParamsInt.HostRewrite = v.(string)

					}

					if v, ok := cs["http_method"]; ok && !isIntfNil(v) {

						choiceInt.SimpleRoute.HttpMethod = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if v, ok := cs["origin_pools"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						originPools := make([]*ves_io_schema_views.OriginPoolWithWeight, len(sl))
						choiceInt.SimpleRoute.OriginPools = originPools
						for i, set := range sl {
							originPools[i] = &ves_io_schema_views.OriginPoolWithWeight{}

							originPoolsMapStrToI := set.(map[string]interface{})

							poolChoiceTypeFound := false

							if v, ok := originPoolsMapStrToI["cluster"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

								poolChoiceTypeFound = true
								poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Cluster{}
								poolChoiceInt.Cluster = &ves_io_schema_views.ObjectRefType{}
								originPools[i].PoolChoice = poolChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										poolChoiceInt.Cluster.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										poolChoiceInt.Cluster.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										poolChoiceInt.Cluster.Tenant = v.(string)
									}

								}

							}

							if v, ok := originPoolsMapStrToI["pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

								poolChoiceTypeFound = true
								poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Pool{}
								poolChoiceInt.Pool = &ves_io_schema_views.ObjectRefType{}
								originPools[i].PoolChoice = poolChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										poolChoiceInt.Pool.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										poolChoiceInt.Pool.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										poolChoiceInt.Pool.Tenant = v.(string)
									}

								}

							}

							if w, ok := originPoolsMapStrToI["weight"]; ok && !isIntfNil(w) {
								originPools[i].Weight = w.(uint32)
							}

						}

					}

					if v, ok := cs["path"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						path := &ves_io_schema.PathMatcherType{}
						choiceInt.SimpleRoute.Path = path
						for _, set := range sl {

							pathMapStrToI := set.(map[string]interface{})

							pathMatchTypeFound := false

							if v, ok := pathMapStrToI["path"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Path{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Path = v.(string)

							}

							if v, ok := pathMapStrToI["prefix"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Prefix{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Prefix = v.(string)

							}

							if v, ok := pathMapStrToI["regex"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Regex{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Regex = v.(string)

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("user_identification"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		userIdentification := &ves_io_schema_views.ObjectRefType{}
		createSpec.UserIdentification = userIdentification
		for _, set := range sl {

			userIdentificationMapStrToI := set.(map[string]interface{})

			if w, ok := userIdentificationMapStrToI["name"]; ok && !isIntfNil(w) {
				userIdentification.Name = w.(string)
			}

			if w, ok := userIdentificationMapStrToI["namespace"]; ok && !isIntfNil(w) {
				userIdentification.Namespace = w.(string)
			}

			if w, ok := userIdentificationMapStrToI["tenant"]; ok && !isIntfNil(w) {
				userIdentification.Tenant = w.(string)
			}

		}

	}

	wafChoiceTypeFound := false

	if v, ok := d.GetOk("disable_waf"); ok && !wafChoiceTypeFound {

		wafChoiceTypeFound = true

		if v.(bool) {
			wafChoiceInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_DisableWaf{}
			wafChoiceInt.DisableWaf = &ves_io_schema.Empty{}
			createSpec.WafChoice = wafChoiceInt
		}

	}

	if v, ok := d.GetOk("waf"); ok && !wafChoiceTypeFound {

		wafChoiceTypeFound = true
		wafChoiceInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_Waf{}
		wafChoiceInt.Waf = &ves_io_schema_views.ObjectRefType{}
		createSpec.WafChoice = wafChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				wafChoiceInt.Waf.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				wafChoiceInt.Waf.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				wafChoiceInt.Waf.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("waf_rule"); ok && !wafChoiceTypeFound {

		wafChoiceTypeFound = true
		wafChoiceInt := &ves_io_schema_views_http_loadbalancer.CreateSpecType_WafRule{}
		wafChoiceInt.WafRule = &ves_io_schema_views.ObjectRefType{}
		createSpec.WafChoice = wafChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				wafChoiceInt.WafRule.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				wafChoiceInt.WafRule.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				wafChoiceInt.WafRule.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("waf_exclusion_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		wafExclusionRules := make([]*ves_io_schema_policy.SimpleWafExclusionRule, len(sl))
		createSpec.WafExclusionRules = wafExclusionRules
		for i, set := range sl {
			wafExclusionRules[i] = &ves_io_schema_policy.SimpleWafExclusionRule{}

			wafExclusionRulesMapStrToI := set.(map[string]interface{})

			if w, ok := wafExclusionRulesMapStrToI["description"]; ok && !isIntfNil(w) {
				wafExclusionRules[i].Description = w.(string)
			}

			domainChoiceTypeFound := false

			if v, ok := wafExclusionRulesMapStrToI["any_domain"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

				domainChoiceTypeFound = true

				if v.(bool) {
					domainChoiceInt := &ves_io_schema_policy.SimpleWafExclusionRule_AnyDomain{}
					domainChoiceInt.AnyDomain = &ves_io_schema.Empty{}
					wafExclusionRules[i].DomainChoice = domainChoiceInt
				}

			}

			if v, ok := wafExclusionRulesMapStrToI["domain_regex"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

				domainChoiceTypeFound = true
				domainChoiceInt := &ves_io_schema_policy.SimpleWafExclusionRule_DomainRegex{}

				wafExclusionRules[i].DomainChoice = domainChoiceInt

				domainChoiceInt.DomainRegex = v.(string)

			}

			if v, ok := wafExclusionRulesMapStrToI["exclude_rule_ids"]; ok && !isIntfNil(v) {

				exclude_rule_idsList := []ves_io_schema_waf_rule_list.WafRuleID{}
				for _, j := range v.([]interface{}) {
					exclude_rule_idsList = append(exclude_rule_idsList, ves_io_schema_waf_rule_list.WafRuleID(ves_io_schema_waf_rule_list.WafRuleID_value[j.(string)]))
				}
				wafExclusionRules[i].ExcludeRuleIds = exclude_rule_idsList

			}

			if w, ok := wafExclusionRulesMapStrToI["expiration_timestamp"]; ok && !isIntfNil(w) {
				ts, err := parseTime(w.(string))
				if err != nil {
					return fmt.Errorf("error creating ExpirationTimestamp, timestamp format is wrong: %s", err)
				}
				wafExclusionRules[i].ExpirationTimestamp = ts
			}

			if v, ok := wafExclusionRulesMapStrToI["methods"]; ok && !isIntfNil(v) {

				methodsList := []ves_io_schema.HttpMethod{}
				for _, j := range v.([]interface{}) {
					methodsList = append(methodsList, ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[j.(string)]))
				}
				wafExclusionRules[i].Methods = methodsList

			}

			if w, ok := wafExclusionRulesMapStrToI["name"]; ok && !isIntfNil(w) {
				wafExclusionRules[i].Name = w.(string)
			}

			if w, ok := wafExclusionRulesMapStrToI["path_regex"]; ok && !isIntfNil(w) {
				wafExclusionRules[i].PathRegex = w.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra HttpLoadbalancer object with struct: %+v", createReq)

	createHttpLoadbalancerResp, err := client.CreateObject(context.Background(), ves_io_schema_views_http_loadbalancer.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating HttpLoadbalancer: %s", err)
	}
	d.SetId(createHttpLoadbalancerResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraHttpLoadbalancerRead(d, meta)
}

func resourceVolterraHttpLoadbalancerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_http_loadbalancer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] HttpLoadbalancer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra HttpLoadbalancer %q: %s", d.Id(), err)
	}
	return setHttpLoadbalancerFields(client, d, resp)
}

func setHttpLoadbalancerFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraHttpLoadbalancerUpdate updates HttpLoadbalancer resource
func resourceVolterraHttpLoadbalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_http_loadbalancer.ReplaceRequest{
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

	if v, ok := d.GetOk("add_location"); ok && !isIntfNil(v) {

		updateSpec.AddLocation =
			v.(bool)
	}

	advertiseChoiceTypeFound := false

	if v, ok := d.GetOk("advertise_custom"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true
		advertiseChoiceInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_AdvertiseCustom{}
		advertiseChoiceInt.AdvertiseCustom = &ves_io_schema_views.AdvertiseCustom{}
		updateSpec.AdvertiseChoice = advertiseChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["advertise_where"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				advertiseWhere := make([]*ves_io_schema_views.WhereType, len(sl))
				advertiseChoiceInt.AdvertiseCustom.AdvertiseWhere = advertiseWhere
				for i, set := range sl {
					advertiseWhere[i] = &ves_io_schema_views.WhereType{}

					advertiseWhereMapStrToI := set.(map[string]interface{})

					choiceTypeFound := false

					if v, ok := advertiseWhereMapStrToI["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.WhereType_Site{}
						choiceInt.Site = &ves_io_schema_views.WhereSite{}
						advertiseWhere[i].Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["ip"]; ok && !isIntfNil(v) {

								choiceInt.Site.Ip = v.(string)
							}

							if v, ok := cs["network"]; ok && !isIntfNil(v) {

								choiceInt.Site.Network = ves_io_schema_views.SiteNetwork(ves_io_schema_views.SiteNetwork_value[v.(string)])

							}

							if v, ok := cs["site"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								site := &ves_io_schema_views.ObjectRefType{}
								choiceInt.Site.Site = site
								for _, set := range sl {

									siteMapStrToI := set.(map[string]interface{})

									if w, ok := siteMapStrToI["name"]; ok && !isIntfNil(w) {
										site.Name = w.(string)
									}

									if w, ok := siteMapStrToI["namespace"]; ok && !isIntfNil(w) {
										site.Namespace = w.(string)
									}

									if w, ok := siteMapStrToI["tenant"]; ok && !isIntfNil(w) {
										site.Tenant = w.(string)
									}

								}

							}

						}

					}

					if v, ok := advertiseWhereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.WhereType_VirtualSite{}
						choiceInt.VirtualSite = &ves_io_schema_views.WhereVirtualSite{}
						advertiseWhere[i].Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["network"]; ok && !isIntfNil(v) {

								choiceInt.VirtualSite.Network = ves_io_schema_views.SiteNetwork(ves_io_schema_views.SiteNetwork_value[v.(string)])

							}

							if v, ok := cs["virtual_site"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								virtualSite := &ves_io_schema_views.ObjectRefType{}
								choiceInt.VirtualSite.VirtualSite = virtualSite
								for _, set := range sl {

									virtualSiteMapStrToI := set.(map[string]interface{})

									if w, ok := virtualSiteMapStrToI["name"]; ok && !isIntfNil(w) {
										virtualSite.Name = w.(string)
									}

									if w, ok := virtualSiteMapStrToI["namespace"]; ok && !isIntfNil(w) {
										virtualSite.Namespace = w.(string)
									}

									if w, ok := virtualSiteMapStrToI["tenant"]; ok && !isIntfNil(w) {
										virtualSite.Tenant = w.(string)
									}

								}

							}

						}

					}

					if v, ok := advertiseWhereMapStrToI["vk8s_service"]; ok && !isIntfNil(v) && !choiceTypeFound {

						choiceTypeFound = true
						choiceInt := &ves_io_schema_views.WhereType_Vk8SService{}
						choiceInt.Vk8SService = &ves_io_schema_views.WhereVK8SService{}
						advertiseWhere[i].Choice = choiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							choiceTypeFound := false

							if v, ok := cs["site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceIntNew := &ves_io_schema_views.WhereVK8SService_Site{}
								choiceIntNew.Site = &ves_io_schema_views.ObjectRefType{}
								choiceInt.Vk8SService.Choice = choiceIntNew

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceIntNew.Site.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceIntNew.Site.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceIntNew.Site.Tenant = v.(string)
									}

								}

							}

							if v, ok := cs["virtual_site"]; ok && !isIntfNil(v) && !choiceTypeFound {

								choiceTypeFound = true
								choiceIntNew := &ves_io_schema_views.WhereVK8SService_VirtualSite{}
								choiceIntNew.VirtualSite = &ves_io_schema_views.ObjectRefType{}
								choiceInt.Vk8SService.Choice = choiceIntNew

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										choiceIntNew.VirtualSite.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										choiceIntNew.VirtualSite.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										choiceIntNew.VirtualSite.Tenant = v.(string)
									}

								}

							}

						}

					}

					portChoiceTypeFound := false

					if v, ok := advertiseWhereMapStrToI["port"]; ok && !isIntfNil(v) && !portChoiceTypeFound {

						portChoiceTypeFound = true
						portChoiceInt := &ves_io_schema_views.WhereType_Port{}

						advertiseWhere[i].PortChoice = portChoiceInt

						portChoiceInt.Port =
							uint32(v.(int))

					}

					if v, ok := advertiseWhereMapStrToI["use_default_port"]; ok && !isIntfNil(v) && !portChoiceTypeFound {

						portChoiceTypeFound = true

						if v.(bool) {
							portChoiceInt := &ves_io_schema_views.WhereType_UseDefaultPort{}
							portChoiceInt.UseDefaultPort = &ves_io_schema.Empty{}
							advertiseWhere[i].PortChoice = portChoiceInt
						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("advertise_on_public"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true
		advertiseChoiceInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_AdvertiseOnPublic{}
		advertiseChoiceInt.AdvertiseOnPublic = &ves_io_schema_views.AdvertisePublic{}
		updateSpec.AdvertiseChoice = advertiseChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["public_ip"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				publicIp := &ves_io_schema_views.ObjectRefType{}
				advertiseChoiceInt.AdvertiseOnPublic.PublicIp = publicIp
				for _, set := range sl {

					publicIpMapStrToI := set.(map[string]interface{})

					if w, ok := publicIpMapStrToI["name"]; ok && !isIntfNil(w) {
						publicIp.Name = w.(string)
					}

					if w, ok := publicIpMapStrToI["namespace"]; ok && !isIntfNil(w) {
						publicIp.Namespace = w.(string)
					}

					if w, ok := publicIpMapStrToI["tenant"]; ok && !isIntfNil(w) {
						publicIp.Tenant = w.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("advertise_on_public_default_vip"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true

		if v.(bool) {
			advertiseChoiceInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_AdvertiseOnPublicDefaultVip{}
			advertiseChoiceInt.AdvertiseOnPublicDefaultVip = &ves_io_schema.Empty{}
			updateSpec.AdvertiseChoice = advertiseChoiceInt
		}

	}

	if v, ok := d.GetOk("do_not_advertise"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true

		if v.(bool) {
			advertiseChoiceInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_DoNotAdvertise{}
			advertiseChoiceInt.DoNotAdvertise = &ves_io_schema.Empty{}
			updateSpec.AdvertiseChoice = advertiseChoiceInt
		}

	}

	challengeTypeTypeFound := false

	if v, ok := d.GetOk("captcha_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true
		challengeTypeInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_CaptchaChallenge{}
		challengeTypeInt.CaptchaChallenge = &ves_io_schema_virtual_host.CaptchaChallengeType{}
		updateSpec.ChallengeType = challengeTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cookie_expiry"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.CookieExpiry = uint32(v.(int))
			}

			if v, ok := cs["custom_page"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.CustomPage = v.(string)
			}

			if v, ok := cs["enable_captcha_challenge"]; ok && !isIntfNil(v) {

				challengeTypeInt.CaptchaChallenge.EnableCaptchaChallenge = v.(bool)
			}

		}

	}

	if v, ok := d.GetOk("js_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true
		challengeTypeInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_JsChallenge{}
		challengeTypeInt.JsChallenge = &ves_io_schema_virtual_host.JavascriptChallengeType{}
		updateSpec.ChallengeType = challengeTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cookie_expiry"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.CookieExpiry = uint32(v.(int))
			}

			if v, ok := cs["custom_page"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.CustomPage = v.(string)
			}

			if v, ok := cs["enable_js_challenge"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.EnableJsChallenge = v.(bool)
			}

			if v, ok := cs["js_script_delay"]; ok && !isIntfNil(v) {

				challengeTypeInt.JsChallenge.JsScriptDelay = uint32(v.(int))
			}

		}

	}

	if v, ok := d.GetOk("no_challenge"); ok && !challengeTypeTypeFound {

		challengeTypeTypeFound = true

		if v.(bool) {
			challengeTypeInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_NoChallenge{}
			challengeTypeInt.NoChallenge = &ves_io_schema.Empty{}
			updateSpec.ChallengeType = challengeTypeInt
		}

	}

	if v, ok := d.GetOk("cors_policy"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		corsPolicy := &ves_io_schema.CorsPolicy{}
		updateSpec.CorsPolicy = corsPolicy
		for _, set := range sl {

			corsPolicyMapStrToI := set.(map[string]interface{})

			if w, ok := corsPolicyMapStrToI["allow_credentials"]; ok && !isIntfNil(w) {
				corsPolicy.AllowCredentials = w.(bool)
			}

			if w, ok := corsPolicyMapStrToI["allow_headers"]; ok && !isIntfNil(w) {
				corsPolicy.AllowHeaders = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["allow_methods"]; ok && !isIntfNil(w) {
				corsPolicy.AllowMethods = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["allow_origin"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				corsPolicy.AllowOrigin = ls
			}

			if w, ok := corsPolicyMapStrToI["allow_origin_regex"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				corsPolicy.AllowOriginRegex = ls
			}

			if w, ok := corsPolicyMapStrToI["disabled"]; ok && !isIntfNil(w) {
				corsPolicy.Disabled = w.(bool)
			}

			if w, ok := corsPolicyMapStrToI["expose_headers"]; ok && !isIntfNil(w) {
				corsPolicy.ExposeHeaders = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["max_age"]; ok && !isIntfNil(w) {
				corsPolicy.MaxAge = w.(string)
			}

			if w, ok := corsPolicyMapStrToI["maximum_age"]; ok && !isIntfNil(w) {
				corsPolicy.MaximumAge = w.(int32)
			}

		}

	}

	if v, ok := d.GetOk("default_route_pools"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		defaultRoutePools := make([]*ves_io_schema_views.OriginPoolWithWeight, len(sl))
		updateSpec.DefaultRoutePools = defaultRoutePools
		for i, set := range sl {
			defaultRoutePools[i] = &ves_io_schema_views.OriginPoolWithWeight{}

			defaultRoutePoolsMapStrToI := set.(map[string]interface{})

			poolChoiceTypeFound := false

			if v, ok := defaultRoutePoolsMapStrToI["cluster"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Cluster{}
				poolChoiceInt.Cluster = &ves_io_schema_views.ObjectRefType{}
				defaultRoutePools[i].PoolChoice = poolChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						poolChoiceInt.Cluster.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						poolChoiceInt.Cluster.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						poolChoiceInt.Cluster.Tenant = v.(string)
					}

				}

			}

			if v, ok := defaultRoutePoolsMapStrToI["pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Pool{}
				poolChoiceInt.Pool = &ves_io_schema_views.ObjectRefType{}
				defaultRoutePools[i].PoolChoice = poolChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						poolChoiceInt.Pool.Name = v.(string)
					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						poolChoiceInt.Pool.Namespace = v.(string)
					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						poolChoiceInt.Pool.Tenant = v.(string)
					}

				}

			}

			if w, ok := defaultRoutePoolsMapStrToI["weight"]; ok && !isIntfNil(w) {
				defaultRoutePools[i].Weight = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("domains"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		updateSpec.Domains = ls

	}

	loadbalancerTypeTypeFound := false

	if v, ok := d.GetOk("http"); ok && !loadbalancerTypeTypeFound {

		loadbalancerTypeTypeFound = true

		if v.(bool) {
			loadbalancerTypeInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_Http{}
			loadbalancerTypeInt.Http = &ves_io_schema.Empty{}
			updateSpec.LoadbalancerType = loadbalancerTypeInt
		}

	}

	if v, ok := d.GetOk("https"); ok && !loadbalancerTypeTypeFound {

		loadbalancerTypeTypeFound = true
		loadbalancerTypeInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_Https{}
		loadbalancerTypeInt.Https = &ves_io_schema_views_http_loadbalancer.ProxyTypeHttps{}
		updateSpec.LoadbalancerType = loadbalancerTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["add_hsts"]; ok && !isIntfNil(v) {

				loadbalancerTypeInt.Https.AddHsts = v.(bool)
			}

			if v, ok := cs["http_redirect"]; ok && !isIntfNil(v) {

				loadbalancerTypeInt.Https.HttpRedirect = v.(bool)
			}

			if v, ok := cs["tls_parameters"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				tlsParameters := &ves_io_schema_views_http_loadbalancer.DownstreamTlsParamsType{}
				loadbalancerTypeInt.Https.TlsParameters = tlsParameters
				for _, set := range sl {

					tlsParametersMapStrToI := set.(map[string]interface{})

					mtlsChoiceTypeFound := false

					if v, ok := tlsParametersMapStrToI["no_mtls"]; ok && !isIntfNil(v) && !mtlsChoiceTypeFound {

						mtlsChoiceTypeFound = true

						if v.(bool) {
							mtlsChoiceInt := &ves_io_schema_views_http_loadbalancer.DownstreamTlsParamsType_NoMtls{}
							mtlsChoiceInt.NoMtls = &ves_io_schema.Empty{}
							tlsParameters.MtlsChoice = mtlsChoiceInt
						}

					}

					if v, ok := tlsParametersMapStrToI["use_mtls"]; ok && !isIntfNil(v) && !mtlsChoiceTypeFound {

						mtlsChoiceTypeFound = true
						mtlsChoiceInt := &ves_io_schema_views_http_loadbalancer.DownstreamTlsParamsType_UseMtls{}
						mtlsChoiceInt.UseMtls = &ves_io_schema_views_http_loadbalancer.DownstreamTlsValidationContext{}
						tlsParameters.MtlsChoice = mtlsChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["trusted_ca_url"]; ok && !isIntfNil(v) {

								mtlsChoiceInt.UseMtls.TrustedCaUrl = v.(string)
							}

						}

					}

					if v, ok := tlsParametersMapStrToI["tls_certificates"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						tlsCertificates := make([]*ves_io_schema.TlsCertificateType, len(sl))
						tlsParameters.TlsCertificates = tlsCertificates
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

					if v, ok := tlsParametersMapStrToI["tls_config"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						tlsConfig := &ves_io_schema_views.TlsConfig{}
						tlsParameters.TlsConfig = tlsConfig
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

		}

	}

	if v, ok := d.GetOk("https_auto_cert"); ok && !loadbalancerTypeTypeFound {

		loadbalancerTypeTypeFound = true
		loadbalancerTypeInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_HttpsAutoCert{}
		loadbalancerTypeInt.HttpsAutoCert = &ves_io_schema_views_http_loadbalancer.ProxyTypeHttpsAutoCerts{}
		updateSpec.LoadbalancerType = loadbalancerTypeInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["add_hsts"]; ok && !isIntfNil(v) {

				loadbalancerTypeInt.HttpsAutoCert.AddHsts = v.(bool)
			}

			if v, ok := cs["http_redirect"]; ok && !isIntfNil(v) {

				loadbalancerTypeInt.HttpsAutoCert.HttpRedirect = v.(bool)
			}

		}

	}

	if v, ok := d.GetOk("malicious_user_mitigation"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		maliciousUserMitigation := &ves_io_schema_views.ObjectRefType{}
		updateSpec.MaliciousUserMitigation = maliciousUserMitigation
		for _, set := range sl {

			maliciousUserMitigationMapStrToI := set.(map[string]interface{})

			if w, ok := maliciousUserMitigationMapStrToI["name"]; ok && !isIntfNil(w) {
				maliciousUserMitigation.Name = w.(string)
			}

			if w, ok := maliciousUserMitigationMapStrToI["namespace"]; ok && !isIntfNil(w) {
				maliciousUserMitigation.Namespace = w.(string)
			}

			if w, ok := maliciousUserMitigationMapStrToI["tenant"]; ok && !isIntfNil(w) {
				maliciousUserMitigation.Tenant = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("more_option"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		moreOption := &ves_io_schema_views_http_loadbalancer.AdvancedOptionsType{}
		updateSpec.MoreOption = moreOption
		for _, set := range sl {

			moreOptionMapStrToI := set.(map[string]interface{})

			if v, ok := moreOptionMapStrToI["buffer_policy"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				bufferPolicy := &ves_io_schema.BufferConfigType{}
				moreOption.BufferPolicy = bufferPolicy
				for _, set := range sl {

					bufferPolicyMapStrToI := set.(map[string]interface{})

					if w, ok := bufferPolicyMapStrToI["disabled"]; ok && !isIntfNil(w) {
						bufferPolicy.Disabled = w.(bool)
					}

					if w, ok := bufferPolicyMapStrToI["max_request_bytes"]; ok && !isIntfNil(w) {
						bufferPolicy.MaxRequestBytes = w.(uint32)
					}

					if w, ok := bufferPolicyMapStrToI["max_request_time"]; ok && !isIntfNil(w) {
						bufferPolicy.MaxRequestTime = w.(uint32)
					}

				}

			}

			if v, ok := moreOptionMapStrToI["compression_params"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				compressionParams := &ves_io_schema_virtual_host.CompressionType{}
				moreOption.CompressionParams = compressionParams
				for _, set := range sl {

					compressionParamsMapStrToI := set.(map[string]interface{})

					if w, ok := compressionParamsMapStrToI["content_length"]; ok && !isIntfNil(w) {
						compressionParams.ContentLength = w.(uint32)
					}

					if w, ok := compressionParamsMapStrToI["content_type"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						compressionParams.ContentType = ls
					}

					if w, ok := compressionParamsMapStrToI["disable_on_etag_header"]; ok && !isIntfNil(w) {
						compressionParams.DisableOnEtagHeader = w.(bool)
					}

					if w, ok := compressionParamsMapStrToI["remove_accept_encoding_header"]; ok && !isIntfNil(w) {
						compressionParams.RemoveAcceptEncodingHeader = w.(bool)
					}

				}

			}

			if w, ok := moreOptionMapStrToI["custom_errors"]; ok && !isIntfNil(w) {
				ms := map[uint32]string{}
				for k, v := range w.(map[string]interface{}) {
					val := v.(string)

					s, err := strconv.ParseUint(k, 10, 32)
					if err != nil {
						return fmt.Errorf("Error while decrypting custom_errors: %s", err)
					}
					key := uint32(s)
					ms[key] = val
				}
				moreOption.CustomErrors = ms
			}

			if w, ok := moreOptionMapStrToI["idle_timeout"]; ok && !isIntfNil(w) {
				moreOption.IdleTimeout = w.(uint32)
			}

			if v, ok := moreOptionMapStrToI["javascript_info"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				javascriptInfo := &ves_io_schema_virtual_host.JavaScriptConfigType{}
				moreOption.JavascriptInfo = javascriptInfo
				for _, set := range sl {

					javascriptInfoMapStrToI := set.(map[string]interface{})

					if w, ok := javascriptInfoMapStrToI["cache_prefix"]; ok && !isIntfNil(w) {
						javascriptInfo.CachePrefix = w.(string)
					}

					if w, ok := javascriptInfoMapStrToI["custom_script_url"]; ok && !isIntfNil(w) {
						javascriptInfo.CustomScriptUrl = w.(string)
					}

					jsonFmtString := v.(string)
					jsm := jsonpb.Unmarshaler{}
					if err := jsm.Unmarshal(strings.NewReader(jsonFmtString), javascriptInfo.ScriptConfig); err != nil {
						return err
					}

				}

			}

			if v, ok := moreOptionMapStrToI["jwt"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				jwtInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				moreOption.Jwt = jwtInt
				for i, ps := range sl {

					jMapToStrVal := ps.(map[string]interface{})
					jwtInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := jMapToStrVal["name"]; ok && !isIntfNil(v) {
						jwtInt[i].Name = v.(string)
					}

					if v, ok := jMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						jwtInt[i].Namespace = v.(string)
					}

					if v, ok := jMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						jwtInt[i].Tenant = v.(string)
					}

				}

			}

			if w, ok := moreOptionMapStrToI["max_request_header_size"]; ok && !isIntfNil(w) {
				moreOption.MaxRequestHeaderSize = w.(uint32)
			}

			if v, ok := moreOptionMapStrToI["request_headers_to_add"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				requestHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
				moreOption.RequestHeadersToAdd = requestHeadersToAdd
				for i, set := range sl {
					requestHeadersToAdd[i] = &ves_io_schema.HeaderManipulationOptionType{}

					requestHeadersToAddMapStrToI := set.(map[string]interface{})

					if w, ok := requestHeadersToAddMapStrToI["append"]; ok && !isIntfNil(w) {
						requestHeadersToAdd[i].Append = w.(bool)
					}

					if w, ok := requestHeadersToAddMapStrToI["name"]; ok && !isIntfNil(w) {
						requestHeadersToAdd[i].Name = w.(string)
					}

					if w, ok := requestHeadersToAddMapStrToI["value"]; ok && !isIntfNil(w) {
						requestHeadersToAdd[i].Value = w.(string)
					}

				}

			}

			if w, ok := moreOptionMapStrToI["request_headers_to_remove"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				moreOption.RequestHeadersToRemove = ls
			}

			if v, ok := moreOptionMapStrToI["response_headers_to_add"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				responseHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
				moreOption.ResponseHeadersToAdd = responseHeadersToAdd
				for i, set := range sl {
					responseHeadersToAdd[i] = &ves_io_schema.HeaderManipulationOptionType{}

					responseHeadersToAddMapStrToI := set.(map[string]interface{})

					if w, ok := responseHeadersToAddMapStrToI["append"]; ok && !isIntfNil(w) {
						responseHeadersToAdd[i].Append = w.(bool)
					}

					if w, ok := responseHeadersToAddMapStrToI["name"]; ok && !isIntfNil(w) {
						responseHeadersToAdd[i].Name = w.(string)
					}

					if w, ok := responseHeadersToAddMapStrToI["value"]; ok && !isIntfNil(w) {
						responseHeadersToAdd[i].Value = w.(string)
					}

				}

			}

			if w, ok := moreOptionMapStrToI["response_headers_to_remove"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				moreOption.ResponseHeadersToRemove = ls
			}

		}

	}

	rateLimitChoiceTypeFound := false

	if v, ok := d.GetOk("disable_rate_limit"); ok && !rateLimitChoiceTypeFound {

		rateLimitChoiceTypeFound = true

		if v.(bool) {
			rateLimitChoiceInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_DisableRateLimit{}
			rateLimitChoiceInt.DisableRateLimit = &ves_io_schema.Empty{}
			updateSpec.RateLimitChoice = rateLimitChoiceInt
		}

	}

	if v, ok := d.GetOk("rate_limit"); ok && !rateLimitChoiceTypeFound {

		rateLimitChoiceTypeFound = true
		rateLimitChoiceInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_RateLimit{}
		rateLimitChoiceInt.RateLimit = &ves_io_schema_views_http_loadbalancer.RateLimitConfigType{}
		updateSpec.RateLimitChoice = rateLimitChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			ipAllowedListChoiceTypeFound := false

			if v, ok := cs["custom_ip_allowed_list"]; ok && !isIntfNil(v) && !ipAllowedListChoiceTypeFound {

				ipAllowedListChoiceTypeFound = true
				ipAllowedListChoiceInt := &ves_io_schema_views_http_loadbalancer.RateLimitConfigType_CustomIpAllowedList{}
				ipAllowedListChoiceInt.CustomIpAllowedList = &ves_io_schema_views_http_loadbalancer.CustomIpAllowedList{}
				rateLimitChoiceInt.RateLimit.IpAllowedListChoice = ipAllowedListChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["rate_limiter_allowed_prefixes"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						rateLimiterAllowedPrefixesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
						ipAllowedListChoiceInt.CustomIpAllowedList.RateLimiterAllowedPrefixes = rateLimiterAllowedPrefixesInt
						for i, ps := range sl {

							rlapMapToStrVal := ps.(map[string]interface{})
							rateLimiterAllowedPrefixesInt[i] = &ves_io_schema_views.ObjectRefType{}

							if v, ok := rlapMapToStrVal["name"]; ok && !isIntfNil(v) {
								rateLimiterAllowedPrefixesInt[i].Name = v.(string)
							}

							if v, ok := rlapMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								rateLimiterAllowedPrefixesInt[i].Namespace = v.(string)
							}

							if v, ok := rlapMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								rateLimiterAllowedPrefixesInt[i].Tenant = v.(string)
							}

						}

					}

				}

			}

			if v, ok := cs["ip_allowed_list"]; ok && !isIntfNil(v) && !ipAllowedListChoiceTypeFound {

				ipAllowedListChoiceTypeFound = true
				ipAllowedListChoiceInt := &ves_io_schema_views_http_loadbalancer.RateLimitConfigType_IpAllowedList{}
				ipAllowedListChoiceInt.IpAllowedList = &ves_io_schema_views.PrefixStringListType{}
				rateLimitChoiceInt.RateLimit.IpAllowedListChoice = ipAllowedListChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						ipAllowedListChoiceInt.IpAllowedList.Prefixes = ls

					}

				}

			}

			if v, ok := cs["no_ip_allowed_list"]; ok && !isIntfNil(v) && !ipAllowedListChoiceTypeFound {

				ipAllowedListChoiceTypeFound = true

				if v.(bool) {
					ipAllowedListChoiceInt := &ves_io_schema_views_http_loadbalancer.RateLimitConfigType_NoIpAllowedList{}
					ipAllowedListChoiceInt.NoIpAllowedList = &ves_io_schema.Empty{}
					rateLimitChoiceInt.RateLimit.IpAllowedListChoice = ipAllowedListChoiceInt
				}

			}

			if v, ok := cs["rate_limiter"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				rateLimiter := &ves_io_schema_rate_limiter.RateLimitValue{}
				rateLimitChoiceInt.RateLimit.RateLimiter = rateLimiter
				for _, set := range sl {

					rateLimiterMapStrToI := set.(map[string]interface{})

					if w, ok := rateLimiterMapStrToI["total_number"]; ok && !isIntfNil(w) {
						rateLimiter.TotalNumber = w.(uint32)
					}

					if v, ok := rateLimiterMapStrToI["unit"]; ok && !isIntfNil(v) {

						rateLimiter.Unit = ves_io_schema_rate_limiter.RateLimitPeriodUnit(ves_io_schema_rate_limiter.RateLimitPeriodUnit_value[v.(string)])

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		routes := make([]*ves_io_schema_views_http_loadbalancer.RouteType, len(sl))
		updateSpec.Routes = routes
		for i, set := range sl {
			routes[i] = &ves_io_schema_views_http_loadbalancer.RouteType{}

			routesMapStrToI := set.(map[string]interface{})

			choiceTypeFound := false

			if v, ok := routesMapStrToI["custom_route_object"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_http_loadbalancer.RouteType_CustomRouteObject{}
				choiceInt.CustomRouteObject = &ves_io_schema_views_http_loadbalancer.RouteTypeCustomRoute{}
				routes[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["route_ref"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						routeRef := &ves_io_schema_views.ObjectRefType{}
						choiceInt.CustomRouteObject.RouteRef = routeRef
						for _, set := range sl {

							routeRefMapStrToI := set.(map[string]interface{})

							if w, ok := routeRefMapStrToI["name"]; ok && !isIntfNil(w) {
								routeRef.Name = w.(string)
							}

							if w, ok := routeRefMapStrToI["namespace"]; ok && !isIntfNil(w) {
								routeRef.Namespace = w.(string)
							}

							if w, ok := routeRefMapStrToI["tenant"]; ok && !isIntfNil(w) {
								routeRef.Tenant = w.(string)
							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["direct_response_route"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_http_loadbalancer.RouteType_DirectResponseRoute{}
				choiceInt.DirectResponseRoute = &ves_io_schema_views_http_loadbalancer.RouteTypeDirectResponse{}
				routes[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["http_method"]; ok && !isIntfNil(v) {

						choiceInt.DirectResponseRoute.HttpMethod = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if v, ok := cs["path"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						path := &ves_io_schema.PathMatcherType{}
						choiceInt.DirectResponseRoute.Path = path
						for _, set := range sl {

							pathMapStrToI := set.(map[string]interface{})

							pathMatchTypeFound := false

							if v, ok := pathMapStrToI["path"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Path{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Path = v.(string)

							}

							if v, ok := pathMapStrToI["prefix"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Prefix{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Prefix = v.(string)

							}

							if v, ok := pathMapStrToI["regex"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Regex{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Regex = v.(string)

							}

						}

					}

					if v, ok := cs["route_direct_response"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						routeDirectResponse := &ves_io_schema_route.RouteDirectResponse{}
						choiceInt.DirectResponseRoute.RouteDirectResponse = routeDirectResponse
						for _, set := range sl {

							routeDirectResponseMapStrToI := set.(map[string]interface{})

							if w, ok := routeDirectResponseMapStrToI["response_body"]; ok && !isIntfNil(w) {
								routeDirectResponse.ResponseBody = w.(string)
							}

							if w, ok := routeDirectResponseMapStrToI["response_code"]; ok && !isIntfNil(w) {
								routeDirectResponse.ResponseCode = w.(uint32)
							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["redirect_route"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_http_loadbalancer.RouteType_RedirectRoute{}
				choiceInt.RedirectRoute = &ves_io_schema_views_http_loadbalancer.RouteTypeRedirect{}
				routes[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["http_method"]; ok && !isIntfNil(v) {

						choiceInt.RedirectRoute.HttpMethod = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if v, ok := cs["path"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						path := &ves_io_schema.PathMatcherType{}
						choiceInt.RedirectRoute.Path = path
						for _, set := range sl {

							pathMapStrToI := set.(map[string]interface{})

							pathMatchTypeFound := false

							if v, ok := pathMapStrToI["path"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Path{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Path = v.(string)

							}

							if v, ok := pathMapStrToI["prefix"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Prefix{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Prefix = v.(string)

							}

							if v, ok := pathMapStrToI["regex"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Regex{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Regex = v.(string)

							}

						}

					}

					if v, ok := cs["route_redirect"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						routeRedirect := &ves_io_schema_route.RouteRedirect{}
						choiceInt.RedirectRoute.RouteRedirect = routeRedirect
						for _, set := range sl {

							routeRedirectMapStrToI := set.(map[string]interface{})

							queryParamsTypeFound := false

							if v, ok := routeRedirectMapStrToI["all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

								queryParamsTypeFound = true
								queryParamsInt := &ves_io_schema_route.RouteRedirect_AllParams{}

								routeRedirect.QueryParams = queryParamsInt

								queryParamsInt.AllParams =
									v.(bool)

							}

							if v, ok := routeRedirectMapStrToI["remove_all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

								queryParamsTypeFound = true

								if v.(bool) {
									queryParamsInt := &ves_io_schema_route.RouteRedirect_RemoveAllParams{}
									queryParamsInt.RemoveAllParams = &ves_io_schema.Empty{}
									routeRedirect.QueryParams = queryParamsInt
								}

							}

							if v, ok := routeRedirectMapStrToI["retain_all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

								queryParamsTypeFound = true

								if v.(bool) {
									queryParamsInt := &ves_io_schema_route.RouteRedirect_RetainAllParams{}
									queryParamsInt.RetainAllParams = &ves_io_schema.Empty{}
									routeRedirect.QueryParams = queryParamsInt
								}

							}

							if v, ok := routeRedirectMapStrToI["strip_query_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

								queryParamsTypeFound = true
								queryParamsInt := &ves_io_schema_route.RouteRedirect_StripQueryParams{}
								queryParamsInt.StripQueryParams = &ves_io_schema_route.RouteQueryParams{}
								routeRedirect.QueryParams = queryParamsInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["query_params"]; ok && !isIntfNil(v) {

										ls := make([]string, len(v.([]interface{})))
										for i, v := range v.([]interface{}) {
											ls[i] = v.(string)
										}
										queryParamsInt.StripQueryParams.QueryParams = ls

									}

								}

							}

							if w, ok := routeRedirectMapStrToI["host_redirect"]; ok && !isIntfNil(w) {
								routeRedirect.HostRedirect = w.(string)
							}

							if w, ok := routeRedirectMapStrToI["path_redirect"]; ok && !isIntfNil(w) {
								routeRedirect.PathRedirect = w.(string)
							}

							if w, ok := routeRedirectMapStrToI["proto_redirect"]; ok && !isIntfNil(w) {
								routeRedirect.ProtoRedirect = w.(string)
							}

							if w, ok := routeRedirectMapStrToI["response_code"]; ok && !isIntfNil(w) {
								routeRedirect.ResponseCode = w.(uint32)
							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["simple_route"]; ok && !isIntfNil(v) && !choiceTypeFound {

				choiceTypeFound = true
				choiceInt := &ves_io_schema_views_http_loadbalancer.RouteType_SimpleRoute{}
				choiceInt.SimpleRoute = &ves_io_schema_views_http_loadbalancer.RouteTypeSimple{}
				routes[i].Choice = choiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					hostRewriteParamsTypeFound := false

					if v, ok := cs["auto_host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true

						if v.(bool) {
							hostRewriteParamsInt := &ves_io_schema_views_http_loadbalancer.RouteTypeSimple_AutoHostRewrite{}
							hostRewriteParamsInt.AutoHostRewrite = &ves_io_schema.Empty{}
							choiceInt.SimpleRoute.HostRewriteParams = hostRewriteParamsInt
						}

					}

					if v, ok := cs["disable_host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true

						if v.(bool) {
							hostRewriteParamsInt := &ves_io_schema_views_http_loadbalancer.RouteTypeSimple_DisableHostRewrite{}
							hostRewriteParamsInt.DisableHostRewrite = &ves_io_schema.Empty{}
							choiceInt.SimpleRoute.HostRewriteParams = hostRewriteParamsInt
						}

					}

					if v, ok := cs["host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true
						hostRewriteParamsInt := &ves_io_schema_views_http_loadbalancer.RouteTypeSimple_HostRewrite{}

						choiceInt.SimpleRoute.HostRewriteParams = hostRewriteParamsInt

						hostRewriteParamsInt.HostRewrite = v.(string)

					}

					if v, ok := cs["http_method"]; ok && !isIntfNil(v) {

						choiceInt.SimpleRoute.HttpMethod = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if v, ok := cs["origin_pools"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						originPools := make([]*ves_io_schema_views.OriginPoolWithWeight, len(sl))
						choiceInt.SimpleRoute.OriginPools = originPools
						for i, set := range sl {
							originPools[i] = &ves_io_schema_views.OriginPoolWithWeight{}

							originPoolsMapStrToI := set.(map[string]interface{})

							poolChoiceTypeFound := false

							if v, ok := originPoolsMapStrToI["cluster"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

								poolChoiceTypeFound = true
								poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Cluster{}
								poolChoiceInt.Cluster = &ves_io_schema_views.ObjectRefType{}
								originPools[i].PoolChoice = poolChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										poolChoiceInt.Cluster.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										poolChoiceInt.Cluster.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										poolChoiceInt.Cluster.Tenant = v.(string)
									}

								}

							}

							if v, ok := originPoolsMapStrToI["pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

								poolChoiceTypeFound = true
								poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Pool{}
								poolChoiceInt.Pool = &ves_io_schema_views.ObjectRefType{}
								originPools[i].PoolChoice = poolChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										poolChoiceInt.Pool.Name = v.(string)
									}

									if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

										poolChoiceInt.Pool.Namespace = v.(string)
									}

									if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

										poolChoiceInt.Pool.Tenant = v.(string)
									}

								}

							}

							if w, ok := originPoolsMapStrToI["weight"]; ok && !isIntfNil(w) {
								originPools[i].Weight = w.(uint32)
							}

						}

					}

					if v, ok := cs["path"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						path := &ves_io_schema.PathMatcherType{}
						choiceInt.SimpleRoute.Path = path
						for _, set := range sl {

							pathMapStrToI := set.(map[string]interface{})

							pathMatchTypeFound := false

							if v, ok := pathMapStrToI["path"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Path{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Path = v.(string)

							}

							if v, ok := pathMapStrToI["prefix"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Prefix{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Prefix = v.(string)

							}

							if v, ok := pathMapStrToI["regex"]; ok && !isIntfNil(v) && !pathMatchTypeFound {

								pathMatchTypeFound = true
								pathMatchInt := &ves_io_schema.PathMatcherType_Regex{}

								path.PathMatch = pathMatchInt

								pathMatchInt.Regex = v.(string)

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("user_identification"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		userIdentification := &ves_io_schema_views.ObjectRefType{}
		updateSpec.UserIdentification = userIdentification
		for _, set := range sl {

			userIdentificationMapStrToI := set.(map[string]interface{})

			if w, ok := userIdentificationMapStrToI["name"]; ok && !isIntfNil(w) {
				userIdentification.Name = w.(string)
			}

			if w, ok := userIdentificationMapStrToI["namespace"]; ok && !isIntfNil(w) {
				userIdentification.Namespace = w.(string)
			}

			if w, ok := userIdentificationMapStrToI["tenant"]; ok && !isIntfNil(w) {
				userIdentification.Tenant = w.(string)
			}

		}

	}

	wafChoiceTypeFound := false

	if v, ok := d.GetOk("disable_waf"); ok && !wafChoiceTypeFound {

		wafChoiceTypeFound = true

		if v.(bool) {
			wafChoiceInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_DisableWaf{}
			wafChoiceInt.DisableWaf = &ves_io_schema.Empty{}
			updateSpec.WafChoice = wafChoiceInt
		}

	}

	if v, ok := d.GetOk("waf"); ok && !wafChoiceTypeFound {

		wafChoiceTypeFound = true
		wafChoiceInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_Waf{}
		wafChoiceInt.Waf = &ves_io_schema_views.ObjectRefType{}
		updateSpec.WafChoice = wafChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				wafChoiceInt.Waf.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				wafChoiceInt.Waf.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				wafChoiceInt.Waf.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("waf_rule"); ok && !wafChoiceTypeFound {

		wafChoiceTypeFound = true
		wafChoiceInt := &ves_io_schema_views_http_loadbalancer.ReplaceSpecType_WafRule{}
		wafChoiceInt.WafRule = &ves_io_schema_views.ObjectRefType{}
		updateSpec.WafChoice = wafChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				wafChoiceInt.WafRule.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				wafChoiceInt.WafRule.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				wafChoiceInt.WafRule.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("waf_exclusion_rules"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		wafExclusionRules := make([]*ves_io_schema_policy.SimpleWafExclusionRule, len(sl))
		updateSpec.WafExclusionRules = wafExclusionRules
		for i, set := range sl {
			wafExclusionRules[i] = &ves_io_schema_policy.SimpleWafExclusionRule{}

			wafExclusionRulesMapStrToI := set.(map[string]interface{})

			if w, ok := wafExclusionRulesMapStrToI["description"]; ok && !isIntfNil(w) {
				wafExclusionRules[i].Description = w.(string)
			}

			domainChoiceTypeFound := false

			if v, ok := wafExclusionRulesMapStrToI["any_domain"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

				domainChoiceTypeFound = true

				if v.(bool) {
					domainChoiceInt := &ves_io_schema_policy.SimpleWafExclusionRule_AnyDomain{}
					domainChoiceInt.AnyDomain = &ves_io_schema.Empty{}
					wafExclusionRules[i].DomainChoice = domainChoiceInt
				}

			}

			if v, ok := wafExclusionRulesMapStrToI["domain_regex"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

				domainChoiceTypeFound = true
				domainChoiceInt := &ves_io_schema_policy.SimpleWafExclusionRule_DomainRegex{}

				wafExclusionRules[i].DomainChoice = domainChoiceInt

				domainChoiceInt.DomainRegex = v.(string)

			}

			if v, ok := wafExclusionRulesMapStrToI["exclude_rule_ids"]; ok && !isIntfNil(v) {

				exclude_rule_idsList := []ves_io_schema_waf_rule_list.WafRuleID{}
				for _, j := range v.([]interface{}) {
					exclude_rule_idsList = append(exclude_rule_idsList, ves_io_schema_waf_rule_list.WafRuleID(ves_io_schema_waf_rule_list.WafRuleID_value[j.(string)]))
				}
				wafExclusionRules[i].ExcludeRuleIds = exclude_rule_idsList

			}

			if w, ok := wafExclusionRulesMapStrToI["expiration_timestamp"]; ok && !isIntfNil(w) {
				ts, err := parseTime(w.(string))
				if err != nil {
					return fmt.Errorf("error creating ExpirationTimestamp, timestamp format is wrong: %s", err)
				}
				wafExclusionRules[i].ExpirationTimestamp = ts
			}

			if v, ok := wafExclusionRulesMapStrToI["methods"]; ok && !isIntfNil(v) {

				methodsList := []ves_io_schema.HttpMethod{}
				for _, j := range v.([]interface{}) {
					methodsList = append(methodsList, ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[j.(string)]))
				}
				wafExclusionRules[i].Methods = methodsList

			}

			if w, ok := wafExclusionRulesMapStrToI["name"]; ok && !isIntfNil(w) {
				wafExclusionRules[i].Name = w.(string)
			}

			if w, ok := wafExclusionRulesMapStrToI["path_regex"]; ok && !isIntfNil(w) {
				wafExclusionRules[i].PathRegex = w.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra HttpLoadbalancer obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_http_loadbalancer.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating HttpLoadbalancer: %s", err)
	}

	return resourceVolterraHttpLoadbalancerRead(d, meta)
}

func resourceVolterraHttpLoadbalancerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_http_loadbalancer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] HttpLoadbalancer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra HttpLoadbalancer before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra HttpLoadbalancer obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_http_loadbalancer.ObjectType, namespace, name)
}
