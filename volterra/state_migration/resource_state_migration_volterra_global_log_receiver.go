//
// Copyright (c) 2026 F5 Inc. All rights reserved.
//

package statemigration

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceGlobalLogReceiverInstanceResourceV1() *schema.Resource {
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

			"ns_all": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"ns_current": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"ns_list": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"namespaces": {
							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"ns_system": {

				Type:       schema.TypeBool,
				Optional:   true,
				Deprecated: "This field is deprecated and will be removed in future release.",
			},

			"audit_logs": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"dns_logs": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"request_logs": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"sampled": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"unsampled": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"security_events": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"aws_cloud_watch_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aws_cred": {
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

						"batch": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_bytes_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_events": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_events_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"timeout_seconds": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"timeout_seconds_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"compression": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_gzip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"group_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"stream_name": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"azure_event_hubs_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connection_string": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
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

									"none": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
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

						"instance": {
							Type:     schema.TypeString,
							Required: true,
						},

						"namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"azure_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"batch": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_bytes_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_events": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_events_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"timeout_seconds": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"timeout_seconds_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"compression": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_gzip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"connection_string": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
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

									"none": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
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

						"container_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"filename_options": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_folder": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"log_type_folder": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"no_folder": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"datadog_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"batch": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_bytes_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_events": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_events_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"timeout_seconds": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"timeout_seconds_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"compression": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_gzip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"datadog_api_key": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
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

									"none": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
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

						"endpoint": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"site": {

							Type:     schema.TypeString,
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

									"no_ca": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"trusted_ca_url": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"mtls_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"mtls_enable": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"key_url": {

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

															"none": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{},
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

									"disable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"disable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"elastic_receiver": {

				Type:       schema.TypeList,
				MaxItems:   1,
				Optional:   true,
				Deprecated: "This field is deprecated and will be removed in future release.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auth_aws": {

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

						"auth_basic": {

							Type:       schema.TypeList,
							MaxItems:   1,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"password": {

										Type:       schema.TypeList,
										MaxItems:   1,
										Optional:   true,
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

												"none": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
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

									"user_name": {
										Type:       schema.TypeString,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},
								},
							},
						},

						"auth_none": {

							Type:       schema.TypeBool,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
						},

						"batch": {

							Type:       schema.TypeList,
							MaxItems:   1,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:       schema.TypeInt,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"max_bytes_disabled": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"max_events": {

										Type:       schema.TypeInt,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"max_events_disabled": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"timeout_seconds": {

										Type:       schema.TypeInt,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"timeout_seconds_default": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},
								},
							},
						},

						"compression": {

							Type:       schema.TypeList,
							MaxItems:   1,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"compression_gzip": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"compression_none": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},
								},
							},
						},

						"endpoint": {
							Type:       schema.TypeString,
							Required:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
						},

						"no_tls": {

							Type:       schema.TypeBool,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
						},

						"use_tls": {

							Type:       schema.TypeList,
							MaxItems:   1,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"no_ca": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"trusted_ca_url": {

										Type:       schema.TypeString,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"mtls_disabled": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"mtls_enable": {

										Type:       schema.TypeList,
										MaxItems:   1,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"key_url": {

													Type:       schema.TypeList,
													MaxItems:   1,
													Optional:   true,
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

															"none": {

																Type:       schema.TypeList,
																MaxItems:   1,
																Optional:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{},
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

									"disable_verify_certificate": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"enable_verify_certificate": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"disable_verify_hostname": {

										Type:       schema.TypeBool,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"enable_verify_hostname": {

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

			"file_receiver": {

				Type:       schema.TypeBool,
				Optional:   true,
				Deprecated: "This field is deprecated and will be removed in future release.",
			},

			"gcp_bucket_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"batch": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_bytes_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_events": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_events_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"timeout_seconds": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"timeout_seconds_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"bucket": {
							Type:     schema.TypeString,
							Required: true,
						},

						"compression": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_gzip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"filename_options": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_folder": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"log_type_folder": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"no_folder": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
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
					},
				},
			},

			"http_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auth_basic": {

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

												"none": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
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

									"user_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"auth_none": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"auth_token": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"token": {

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

												"none": {

													Type:     schema.TypeList,
													MaxItems: 1,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
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

						"batch": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_bytes_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_events": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_events_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"timeout_seconds": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"timeout_seconds_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"compression": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_gzip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
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

									"no_ca": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"trusted_ca_url": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"mtls_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"mtls_enable": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"key_url": {

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

															"none": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{},
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

									"disable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"disable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"uri": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"kafka_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"batch": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_bytes_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_events": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_events_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"timeout_seconds": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"timeout_seconds_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"bootstrap_servers": {
							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"compression": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_gzip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"kafka_topic": {
							Type:     schema.TypeString,
							Required: true,
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

									"no_ca": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"trusted_ca_url": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"mtls_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"mtls_enable": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"key_url": {

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

															"none": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{},
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

									"disable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"disable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"new_relic_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_key": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
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

									"none": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
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

						"eu": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"us": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"qradar_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"batch": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_bytes_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_events": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_events_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"timeout_seconds": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"timeout_seconds_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"compression": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_gzip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
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

									"no_ca": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"trusted_ca_url": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"mtls_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"mtls_enable": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"key_url": {

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

															"none": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{},
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

									"disable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"disable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"uri": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"s3_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aws_cred": {
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

						"batch": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_bytes_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_events": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_events_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"timeout_seconds": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"timeout_seconds_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"bucket": {
							Type:     schema.TypeString,
							Required: true,
						},

						"compression": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_gzip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"filename_options": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_folder": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"log_type_folder": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"no_folder": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"splunk_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"batch": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"max_bytes": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_bytes_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"max_events": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"max_events_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"timeout_seconds": {

										Type:     schema.TypeInt,
										Optional: true,
									},

									"timeout_seconds_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"compression": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"compression_default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_gzip": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"compression_none": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"endpoint": {
							Type:     schema.TypeString,
							Required: true,
						},

						"splunk_hec_token": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
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

									"none": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
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

									"no_ca": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"trusted_ca_url": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"mtls_disabled": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"mtls_enable": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"key_url": {

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

															"none": {

																Type:     schema.TypeList,
																MaxItems: 1,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{},
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

									"disable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_certificate": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"disable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"enable_verify_hostname": {

										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"sumo_logic_receiver": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"url": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
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

									"none": {

										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
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
		},
	}
}

func ResourceGlobalLogReceiverInstanceStateUpgradeV1(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	req_logs, ok := rawState["request_logs"]
	if ok && req_logs != nil && reflect.TypeOf(req_logs).Kind() == reflect.Bool {
		rawState["request_logs"] = []interface{}{map[string]interface{}{
			"sampled": true,
		}}
	}
	return rawState, nil
}
