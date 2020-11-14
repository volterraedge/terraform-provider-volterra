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
	ves_io_schema_route "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/route"
)

// resourceVolterraRoute is implementation of Volterra's Route resources
func resourceVolterraRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraRouteCreate,
		Read:   resourceVolterraRouteRead,
		Update: resourceVolterraRouteUpdate,
		Delete: resourceVolterraRouteDelete,

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

			"routes": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"route_destination": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auto_host_rewrite": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"host_rewrite": {

										Type:     schema.TypeString,
										Optional: true,
									},

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

									"destinations": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cluster": {

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

												"endpoint_subsets": {
													Type:     schema.TypeMap,
													Optional: true,
												},

												"weight": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},

									"endpoint_subsets": {
										Type:     schema.TypeMap,
										Optional: true,
									},

									"hash_policy": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cookie": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"name": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"path": {
																Type:     schema.TypeString,
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

									"mirror_policy": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cluster": {

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

												"percent": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"denominator": {
																Type:     schema.TypeString,
																Optional: true,
															},

															"numerator": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"prefix_rewrite": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"priority": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"retry_policy": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"back_off": {

													Type:     schema.TypeSet,
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

												"retry_on": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"spdy_config": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"use_spdy": {
													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"timeout": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"web_socket_config": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"idle_timeout": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"max_connect_attempts": {
													Type:     schema.TypeInt,
													Optional: true,
												},

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

						"disable_custom_script": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"disable_location_add": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"match": {

							Type:     schema.TypeList,
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
													Optional: true,
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

									"query_params": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"exact": {

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

						"service_policy": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},

						"waf_type": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"waf": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"waf": {

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

									"waf_rules": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"waf_rules": {

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
								},
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraRouteCreate creates Route resource
func resourceVolterraRouteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_route.CreateSpecType{}
	createReq := &ves_io_schema_route.CreateRequest{
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

	if v, ok := d.GetOk("routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		routes := make([]*ves_io_schema_route.RouteType, len(sl))
		createSpec.Routes = routes
		for i, set := range sl {
			routes[i] = &ves_io_schema_route.RouteType{}

			routesMapStrToI := set.(map[string]interface{})

			routeActionTypeFound := false

			if v, ok := routesMapStrToI["route_destination"]; ok && !isIntfNil(v) && !routeActionTypeFound {

				routeActionTypeFound = true
				routeActionInt := &ves_io_schema_route.RouteType_RouteDestination{}
				routeActionInt.RouteDestination = &ves_io_schema_route.RouteDestinationList{}
				routes[i].RouteAction = routeActionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					hostRewriteParamsTypeFound := false

					if v, ok := cs["auto_host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true
						hostRewriteParamsInt := &ves_io_schema_route.RouteDestinationList_AutoHostRewrite{}

						routeActionInt.RouteDestination.HostRewriteParams = hostRewriteParamsInt

						hostRewriteParamsInt.AutoHostRewrite =
							v.(bool)

					}

					if v, ok := cs["host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true
						hostRewriteParamsInt := &ves_io_schema_route.RouteDestinationList_HostRewrite{}

						routeActionInt.RouteDestination.HostRewriteParams = hostRewriteParamsInt

						hostRewriteParamsInt.HostRewrite = v.(string)

					}

					if v, ok := cs["buffer_policy"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						bufferPolicy := &ves_io_schema.BufferConfigType{}
						routeActionInt.RouteDestination.BufferPolicy = bufferPolicy
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

					if v, ok := cs["cors_policy"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						corsPolicy := &ves_io_schema.CorsPolicy{}
						routeActionInt.RouteDestination.CorsPolicy = corsPolicy
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

					if v, ok := cs["destinations"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						destinations := make([]*ves_io_schema_route.RouteDestination, len(sl))
						routeActionInt.RouteDestination.Destinations = destinations
						for i, set := range sl {
							destinations[i] = &ves_io_schema_route.RouteDestination{}

							destinationsMapStrToI := set.(map[string]interface{})

							if v, ok := destinationsMapStrToI["cluster"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								clusterInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								destinations[i].Cluster = clusterInt
								for i, ps := range sl {

									cMapToStrVal := ps.(map[string]interface{})
									clusterInt[i] = &ves_io_schema.ObjectRefType{}

									clusterInt[i].Kind = "cluster"

									if v, ok := cMapToStrVal["name"]; ok && !isIntfNil(v) {
										clusterInt[i].Name = v.(string)
									}

									if v, ok := cMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										clusterInt[i].Namespace = v.(string)
									}

									if v, ok := cMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										clusterInt[i].Tenant = v.(string)
									}

									if v, ok := cMapToStrVal["uid"]; ok && !isIntfNil(v) {
										clusterInt[i].Uid = v.(string)
									}

								}

							}

							if w, ok := destinationsMapStrToI["endpoint_subsets"]; ok && !isIntfNil(w) {
								ms := map[string]string{}
								for k, v := range w.(map[string]interface{}) {
									ms[k] = v.(string)
								}
								destinations[i].EndpointSubsets = ms
							}

							if w, ok := destinationsMapStrToI["weight"]; ok && !isIntfNil(w) {
								destinations[i].Weight = w.(uint32)
							}

						}

					}

					if v, ok := cs["endpoint_subsets"]; ok && !isIntfNil(v) {

						ms := map[string]string{}
						for k, v := range v.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						routeActionInt.RouteDestination.EndpointSubsets = ms
					}

					if v, ok := cs["hash_policy"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						hashPolicy := make([]*ves_io_schema_route.HashPolicyType, len(sl))
						routeActionInt.RouteDestination.HashPolicy = hashPolicy
						for i, set := range sl {
							hashPolicy[i] = &ves_io_schema_route.HashPolicyType{}

							hashPolicyMapStrToI := set.(map[string]interface{})

							policySpecifierTypeFound := false

							if v, ok := hashPolicyMapStrToI["cookie"]; ok && !isIntfNil(v) && !policySpecifierTypeFound {

								policySpecifierTypeFound = true
								policySpecifierInt := &ves_io_schema_route.HashPolicyType_Cookie{}
								policySpecifierInt.Cookie = &ves_io_schema_route.CookieForHashing{}
								hashPolicy[i].PolicySpecifier = policySpecifierInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										policySpecifierInt.Cookie.Name = v.(string)
									}

									if v, ok := cs["path"]; ok && !isIntfNil(v) {

										policySpecifierInt.Cookie.Path = v.(string)
									}

									if v, ok := cs["ttl"]; ok && !isIntfNil(v) {

										policySpecifierInt.Cookie.Ttl = uint32(v.(int))
									}

								}

							}

							if v, ok := hashPolicyMapStrToI["header_name"]; ok && !isIntfNil(v) && !policySpecifierTypeFound {

								policySpecifierTypeFound = true
								policySpecifierInt := &ves_io_schema_route.HashPolicyType_HeaderName{}

								hashPolicy[i].PolicySpecifier = policySpecifierInt

								policySpecifierInt.HeaderName = v.(string)

							}

							if v, ok := hashPolicyMapStrToI["source_ip"]; ok && !isIntfNil(v) && !policySpecifierTypeFound {

								policySpecifierTypeFound = true
								policySpecifierInt := &ves_io_schema_route.HashPolicyType_SourceIp{}

								hashPolicy[i].PolicySpecifier = policySpecifierInt

								policySpecifierInt.SourceIp =
									v.(bool)

							}

							if w, ok := hashPolicyMapStrToI["terminal"]; ok && !isIntfNil(w) {
								hashPolicy[i].Terminal = w.(bool)
							}

						}

					}

					if v, ok := cs["mirror_policy"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						mirrorPolicy := &ves_io_schema_route.MirrorPolicyType{}
						routeActionInt.RouteDestination.MirrorPolicy = mirrorPolicy
						for _, set := range sl {

							mirrorPolicyMapStrToI := set.(map[string]interface{})

							if v, ok := mirrorPolicyMapStrToI["cluster"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								clusterInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								mirrorPolicy.Cluster = clusterInt
								for i, ps := range sl {

									cMapToStrVal := ps.(map[string]interface{})
									clusterInt[i] = &ves_io_schema.ObjectRefType{}

									clusterInt[i].Kind = "cluster"

									if v, ok := cMapToStrVal["name"]; ok && !isIntfNil(v) {
										clusterInt[i].Name = v.(string)
									}

									if v, ok := cMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										clusterInt[i].Namespace = v.(string)
									}

									if v, ok := cMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										clusterInt[i].Tenant = v.(string)
									}

									if v, ok := cMapToStrVal["uid"]; ok && !isIntfNil(v) {
										clusterInt[i].Uid = v.(string)
									}

								}

							}

							if v, ok := mirrorPolicyMapStrToI["percent"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								percent := &ves_io_schema.FractionalPercent{}
								mirrorPolicy.Percent = percent
								for _, set := range sl {

									percentMapStrToI := set.(map[string]interface{})

									if v, ok := percentMapStrToI["denominator"]; ok && !isIntfNil(v) {

										percent.Denominator = ves_io_schema.DenominatorType(ves_io_schema.DenominatorType_value[v.(string)])

									}

									if w, ok := percentMapStrToI["numerator"]; ok && !isIntfNil(w) {
										percent.Numerator = w.(uint32)
									}

								}

							}

						}

					}

					if v, ok := cs["prefix_rewrite"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDestination.PrefixRewrite = v.(string)
					}

					if v, ok := cs["priority"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDestination.Priority = ves_io_schema.RoutingPriority(ves_io_schema.RoutingPriority_value[v.(string)])

					}

					if v, ok := cs["retry_policy"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						retryPolicy := &ves_io_schema.RetryPolicyType{}
						routeActionInt.RouteDestination.RetryPolicy = retryPolicy
						for _, set := range sl {

							retryPolicyMapStrToI := set.(map[string]interface{})

							if v, ok := retryPolicyMapStrToI["back_off"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								backOff := &ves_io_schema.RetryBackOff{}
								retryPolicy.BackOff = backOff
								for _, set := range sl {

									backOffMapStrToI := set.(map[string]interface{})

									if w, ok := backOffMapStrToI["base_interval"]; ok && !isIntfNil(w) {
										backOff.BaseInterval = w.(uint32)
									}

									if w, ok := backOffMapStrToI["max_interval"]; ok && !isIntfNil(w) {
										backOff.MaxInterval = w.(uint32)
									}

								}

							}

							if w, ok := retryPolicyMapStrToI["num_retries"]; ok && !isIntfNil(w) {
								retryPolicy.NumRetries = w.(uint32)
							}

							if w, ok := retryPolicyMapStrToI["per_try_timeout"]; ok && !isIntfNil(w) {
								retryPolicy.PerTryTimeout = w.(uint32)
							}

							if w, ok := retryPolicyMapStrToI["retriable_status_codes"]; ok && !isIntfNil(w) {
								ls := make([]uint32, len(w.([]interface{})))
								for i, v := range w.([]interface{}) {

									ls[i] = uint32(v.(int))
								}
								retryPolicy.RetriableStatusCodes = ls
							}

							if w, ok := retryPolicyMapStrToI["retry_on"]; ok && !isIntfNil(w) {
								retryPolicy.RetryOn = w.(string)
							}

						}

					}

					if v, ok := cs["spdy_config"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						spdyConfig := &ves_io_schema_route.SpdyConfigType{}
						routeActionInt.RouteDestination.SpdyConfig = spdyConfig
						for _, set := range sl {

							spdyConfigMapStrToI := set.(map[string]interface{})

							if w, ok := spdyConfigMapStrToI["use_spdy"]; ok && !isIntfNil(w) {
								spdyConfig.UseSpdy = w.(bool)
							}

						}

					}

					if v, ok := cs["timeout"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDestination.Timeout = uint32(v.(int))
					}

					if v, ok := cs["web_socket_config"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						webSocketConfig := &ves_io_schema_route.WebsocketConfigType{}
						routeActionInt.RouteDestination.WebSocketConfig = webSocketConfig
						for _, set := range sl {

							webSocketConfigMapStrToI := set.(map[string]interface{})

							if w, ok := webSocketConfigMapStrToI["idle_timeout"]; ok && !isIntfNil(w) {
								webSocketConfig.IdleTimeout = w.(uint32)
							}

							if w, ok := webSocketConfigMapStrToI["max_connect_attempts"]; ok && !isIntfNil(w) {
								webSocketConfig.MaxConnectAttempts = w.(uint32)
							}

							if w, ok := webSocketConfigMapStrToI["use_websocket"]; ok && !isIntfNil(w) {
								webSocketConfig.UseWebsocket = w.(bool)
							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["route_direct_response"]; ok && !isIntfNil(v) && !routeActionTypeFound {

				routeActionTypeFound = true
				routeActionInt := &ves_io_schema_route.RouteType_RouteDirectResponse{}
				routeActionInt.RouteDirectResponse = &ves_io_schema_route.RouteDirectResponse{}
				routes[i].RouteAction = routeActionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["response_body"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDirectResponse.ResponseBody = v.(string)
					}

					if v, ok := cs["response_code"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDirectResponse.ResponseCode = uint32(v.(int))
					}

				}

			}

			if v, ok := routesMapStrToI["route_redirect"]; ok && !isIntfNil(v) && !routeActionTypeFound {

				routeActionTypeFound = true
				routeActionInt := &ves_io_schema_route.RouteType_RouteRedirect{}
				routeActionInt.RouteRedirect = &ves_io_schema_route.RouteRedirect{}
				routes[i].RouteAction = routeActionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					queryParamsTypeFound := false

					if v, ok := cs["all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

						queryParamsTypeFound = true
						queryParamsInt := &ves_io_schema_route.RouteRedirect_AllParams{}

						routeActionInt.RouteRedirect.QueryParams = queryParamsInt

						queryParamsInt.AllParams =
							v.(bool)

					}

					if v, ok := cs["remove_all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

						queryParamsTypeFound = true

						if v.(bool) {
							queryParamsInt := &ves_io_schema_route.RouteRedirect_RemoveAllParams{}
							queryParamsInt.RemoveAllParams = &ves_io_schema.Empty{}
							routeActionInt.RouteRedirect.QueryParams = queryParamsInt
						}

					}

					if v, ok := cs["retain_all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

						queryParamsTypeFound = true

						if v.(bool) {
							queryParamsInt := &ves_io_schema_route.RouteRedirect_RetainAllParams{}
							queryParamsInt.RetainAllParams = &ves_io_schema.Empty{}
							routeActionInt.RouteRedirect.QueryParams = queryParamsInt
						}

					}

					if v, ok := cs["strip_query_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

						queryParamsTypeFound = true
						queryParamsInt := &ves_io_schema_route.RouteRedirect_StripQueryParams{}
						queryParamsInt.StripQueryParams = &ves_io_schema_route.RouteQueryParams{}
						routeActionInt.RouteRedirect.QueryParams = queryParamsInt

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

					if v, ok := cs["host_redirect"]; ok && !isIntfNil(v) {

						routeActionInt.RouteRedirect.HostRedirect = v.(string)
					}

					if v, ok := cs["path_redirect"]; ok && !isIntfNil(v) {

						routeActionInt.RouteRedirect.PathRedirect = v.(string)
					}

					if v, ok := cs["proto_redirect"]; ok && !isIntfNil(v) {

						routeActionInt.RouteRedirect.ProtoRedirect = v.(string)
					}

					if v, ok := cs["response_code"]; ok && !isIntfNil(v) {

						routeActionInt.RouteRedirect.ResponseCode = uint32(v.(int))
					}

				}

			}

			if w, ok := routesMapStrToI["disable_custom_script"]; ok && !isIntfNil(w) {
				routes[i].DisableCustomScript = w.(bool)
			}

			if w, ok := routesMapStrToI["disable_location_add"]; ok && !isIntfNil(w) {
				routes[i].DisableLocationAdd = w.(bool)
			}

			if v, ok := routesMapStrToI["match"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				match := make([]*ves_io_schema.RouteMatch, len(sl))
				routes[i].Match = match
				for i, set := range sl {
					match[i] = &ves_io_schema.RouteMatch{}

					matchMapStrToI := set.(map[string]interface{})

					if v, ok := matchMapStrToI["headers"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						headers := make([]*ves_io_schema.HeaderMatcherType, len(sl))
						match[i].Headers = headers
						for i, set := range sl {
							headers[i] = &ves_io_schema.HeaderMatcherType{}

							headersMapStrToI := set.(map[string]interface{})

							if w, ok := headersMapStrToI["invert_match"]; ok && !isIntfNil(w) {
								headers[i].InvertMatch = w.(bool)
							}

							if w, ok := headersMapStrToI["name"]; ok && !isIntfNil(w) {
								headers[i].Name = w.(string)
							}

							valueMatchTypeFound := false

							if v, ok := headersMapStrToI["exact"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.HeaderMatcherType_Exact{}

								headers[i].ValueMatch = valueMatchInt

								valueMatchInt.Exact = v.(string)

							}

							if v, ok := headersMapStrToI["presence"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.HeaderMatcherType_Presence{}

								headers[i].ValueMatch = valueMatchInt

								valueMatchInt.Presence =
									v.(bool)

							}

							if v, ok := headersMapStrToI["regex"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.HeaderMatcherType_Regex{}

								headers[i].ValueMatch = valueMatchInt

								valueMatchInt.Regex = v.(string)

							}

						}

					}

					if v, ok := matchMapStrToI["http_method"]; ok && !isIntfNil(v) {

						match[i].HttpMethod = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if v, ok := matchMapStrToI["path"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						path := &ves_io_schema.PathMatcherType{}
						match[i].Path = path
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

					if v, ok := matchMapStrToI["query_params"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						queryParams := make([]*ves_io_schema.QueryParameterMatcherType, len(sl))
						match[i].QueryParams = queryParams
						for i, set := range sl {
							queryParams[i] = &ves_io_schema.QueryParameterMatcherType{}

							queryParamsMapStrToI := set.(map[string]interface{})

							if w, ok := queryParamsMapStrToI["key"]; ok && !isIntfNil(w) {
								queryParams[i].Key = w.(string)
							}

							valueMatchTypeFound := false

							if v, ok := queryParamsMapStrToI["exact"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.QueryParameterMatcherType_Exact{}

								queryParams[i].ValueMatch = valueMatchInt

								valueMatchInt.Exact = v.(string)

							}

							if v, ok := queryParamsMapStrToI["regex"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.QueryParameterMatcherType_Regex{}

								queryParams[i].ValueMatch = valueMatchInt

								valueMatchInt.Regex = v.(string)

							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["request_headers_to_add"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				requestHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
				routes[i].RequestHeadersToAdd = requestHeadersToAdd
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

			if w, ok := routesMapStrToI["request_headers_to_remove"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				routes[i].RequestHeadersToRemove = ls
			}

			if v, ok := routesMapStrToI["response_headers_to_add"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				responseHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
				routes[i].ResponseHeadersToAdd = responseHeadersToAdd
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

			if w, ok := routesMapStrToI["response_headers_to_remove"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				routes[i].ResponseHeadersToRemove = ls
			}

			if v, ok := routesMapStrToI["service_policy"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				servicePolicy := &ves_io_schema_route.ServicePolicyInfo{}
				routes[i].ServicePolicy = servicePolicy
				for _, set := range sl {

					servicePolicyMapStrToI := set.(map[string]interface{})

					if w, ok := servicePolicyMapStrToI["disable"]; ok && !isIntfNil(w) {
						servicePolicy.Disable = w.(bool)
					}

				}

			}

			if v, ok := routesMapStrToI["waf_type"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				wafType := &ves_io_schema.WafType{}
				routes[i].WafType = wafType
				for _, set := range sl {

					wafTypeMapStrToI := set.(map[string]interface{})

					refTypeTypeFound := false

					if v, ok := wafTypeMapStrToI["waf"]; ok && !isIntfNil(v) && !refTypeTypeFound {

						refTypeTypeFound = true
						refTypeInt := &ves_io_schema.WafType_Waf{}
						refTypeInt.Waf = &ves_io_schema.WafRefType{}
						wafType.RefType = refTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["waf"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								wafInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								refTypeInt.Waf.Waf = wafInt
								for i, ps := range sl {

									wMapToStrVal := ps.(map[string]interface{})
									wafInt[i] = &ves_io_schema.ObjectRefType{}

									wafInt[i].Kind = "waf"

									if v, ok := wMapToStrVal["name"]; ok && !isIntfNil(v) {
										wafInt[i].Name = v.(string)
									}

									if v, ok := wMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										wafInt[i].Namespace = v.(string)
									}

									if v, ok := wMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										wafInt[i].Tenant = v.(string)
									}

									if v, ok := wMapToStrVal["uid"]; ok && !isIntfNil(v) {
										wafInt[i].Uid = v.(string)
									}

								}

							}

						}

					}

					if v, ok := wafTypeMapStrToI["waf_rules"]; ok && !isIntfNil(v) && !refTypeTypeFound {

						refTypeTypeFound = true
						refTypeInt := &ves_io_schema.WafType_WafRules{}
						refTypeInt.WafRules = &ves_io_schema.WafRulesRefType{}
						wafType.RefType = refTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["waf_rules"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								wafRulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								refTypeInt.WafRules.WafRules = wafRulesInt
								for i, ps := range sl {

									wrMapToStrVal := ps.(map[string]interface{})
									wafRulesInt[i] = &ves_io_schema.ObjectRefType{}

									wafRulesInt[i].Kind = "waf_rules"

									if v, ok := wrMapToStrVal["name"]; ok && !isIntfNil(v) {
										wafRulesInt[i].Name = v.(string)
									}

									if v, ok := wrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										wafRulesInt[i].Namespace = v.(string)
									}

									if v, ok := wrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										wafRulesInt[i].Tenant = v.(string)
									}

									if v, ok := wrMapToStrVal["uid"]; ok && !isIntfNil(v) {
										wafRulesInt[i].Uid = v.(string)
									}

								}

							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra Route object with struct: %+v", createReq)

	createRouteResp, err := client.CreateObject(context.Background(), ves_io_schema_route.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Route: %s", err)
	}
	d.SetId(createRouteResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraRouteRead(d, meta)
}

func resourceVolterraRouteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_route.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Route %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Route %q: %s", d.Id(), err)
	}
	return setRouteFields(client, d, resp)
}

func setRouteFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraRouteUpdate updates Route resource
func resourceVolterraRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_route.ReplaceSpecType{}
	updateReq := &ves_io_schema_route.ReplaceRequest{
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

	if v, ok := d.GetOk("routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		routes := make([]*ves_io_schema_route.RouteType, len(sl))
		updateSpec.Routes = routes
		for i, set := range sl {
			routes[i] = &ves_io_schema_route.RouteType{}

			routesMapStrToI := set.(map[string]interface{})

			routeActionTypeFound := false

			if v, ok := routesMapStrToI["route_destination"]; ok && !isIntfNil(v) && !routeActionTypeFound {

				routeActionTypeFound = true
				routeActionInt := &ves_io_schema_route.RouteType_RouteDestination{}
				routeActionInt.RouteDestination = &ves_io_schema_route.RouteDestinationList{}
				routes[i].RouteAction = routeActionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					hostRewriteParamsTypeFound := false

					if v, ok := cs["auto_host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true
						hostRewriteParamsInt := &ves_io_schema_route.RouteDestinationList_AutoHostRewrite{}

						routeActionInt.RouteDestination.HostRewriteParams = hostRewriteParamsInt

						hostRewriteParamsInt.AutoHostRewrite =
							v.(bool)

					}

					if v, ok := cs["host_rewrite"]; ok && !isIntfNil(v) && !hostRewriteParamsTypeFound {

						hostRewriteParamsTypeFound = true
						hostRewriteParamsInt := &ves_io_schema_route.RouteDestinationList_HostRewrite{}

						routeActionInt.RouteDestination.HostRewriteParams = hostRewriteParamsInt

						hostRewriteParamsInt.HostRewrite = v.(string)

					}

					if v, ok := cs["buffer_policy"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						bufferPolicy := &ves_io_schema.BufferConfigType{}
						routeActionInt.RouteDestination.BufferPolicy = bufferPolicy
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

					if v, ok := cs["cors_policy"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						corsPolicy := &ves_io_schema.CorsPolicy{}
						routeActionInt.RouteDestination.CorsPolicy = corsPolicy
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

					if v, ok := cs["destinations"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						destinations := make([]*ves_io_schema_route.RouteDestination, len(sl))
						routeActionInt.RouteDestination.Destinations = destinations
						for i, set := range sl {
							destinations[i] = &ves_io_schema_route.RouteDestination{}

							destinationsMapStrToI := set.(map[string]interface{})

							if v, ok := destinationsMapStrToI["cluster"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								clusterInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								destinations[i].Cluster = clusterInt
								for i, ps := range sl {

									cMapToStrVal := ps.(map[string]interface{})
									clusterInt[i] = &ves_io_schema.ObjectRefType{}

									clusterInt[i].Kind = "cluster"

									if v, ok := cMapToStrVal["name"]; ok && !isIntfNil(v) {
										clusterInt[i].Name = v.(string)
									}

									if v, ok := cMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										clusterInt[i].Namespace = v.(string)
									}

									if v, ok := cMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										clusterInt[i].Tenant = v.(string)
									}

									if v, ok := cMapToStrVal["uid"]; ok && !isIntfNil(v) {
										clusterInt[i].Uid = v.(string)
									}

								}

							}

							if w, ok := destinationsMapStrToI["endpoint_subsets"]; ok && !isIntfNil(w) {
								ms := map[string]string{}
								for k, v := range w.(map[string]interface{}) {
									ms[k] = v.(string)
								}
								destinations[i].EndpointSubsets = ms
							}

							if w, ok := destinationsMapStrToI["weight"]; ok && !isIntfNil(w) {
								destinations[i].Weight = w.(uint32)
							}

						}

					}

					if v, ok := cs["endpoint_subsets"]; ok && !isIntfNil(v) {

						ms := map[string]string{}
						for k, v := range v.(map[string]interface{}) {
							ms[k] = v.(string)
						}
						routeActionInt.RouteDestination.EndpointSubsets = ms
					}

					if v, ok := cs["hash_policy"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						hashPolicy := make([]*ves_io_schema_route.HashPolicyType, len(sl))
						routeActionInt.RouteDestination.HashPolicy = hashPolicy
						for i, set := range sl {
							hashPolicy[i] = &ves_io_schema_route.HashPolicyType{}

							hashPolicyMapStrToI := set.(map[string]interface{})

							policySpecifierTypeFound := false

							if v, ok := hashPolicyMapStrToI["cookie"]; ok && !isIntfNil(v) && !policySpecifierTypeFound {

								policySpecifierTypeFound = true
								policySpecifierInt := &ves_io_schema_route.HashPolicyType_Cookie{}
								policySpecifierInt.Cookie = &ves_io_schema_route.CookieForHashing{}
								hashPolicy[i].PolicySpecifier = policySpecifierInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["name"]; ok && !isIntfNil(v) {

										policySpecifierInt.Cookie.Name = v.(string)
									}

									if v, ok := cs["path"]; ok && !isIntfNil(v) {

										policySpecifierInt.Cookie.Path = v.(string)
									}

									if v, ok := cs["ttl"]; ok && !isIntfNil(v) {

										policySpecifierInt.Cookie.Ttl = uint32(v.(int))
									}

								}

							}

							if v, ok := hashPolicyMapStrToI["header_name"]; ok && !isIntfNil(v) && !policySpecifierTypeFound {

								policySpecifierTypeFound = true
								policySpecifierInt := &ves_io_schema_route.HashPolicyType_HeaderName{}

								hashPolicy[i].PolicySpecifier = policySpecifierInt

								policySpecifierInt.HeaderName = v.(string)

							}

							if v, ok := hashPolicyMapStrToI["source_ip"]; ok && !isIntfNil(v) && !policySpecifierTypeFound {

								policySpecifierTypeFound = true
								policySpecifierInt := &ves_io_schema_route.HashPolicyType_SourceIp{}

								hashPolicy[i].PolicySpecifier = policySpecifierInt

								policySpecifierInt.SourceIp =
									v.(bool)

							}

							if w, ok := hashPolicyMapStrToI["terminal"]; ok && !isIntfNil(w) {
								hashPolicy[i].Terminal = w.(bool)
							}

						}

					}

					if v, ok := cs["mirror_policy"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						mirrorPolicy := &ves_io_schema_route.MirrorPolicyType{}
						routeActionInt.RouteDestination.MirrorPolicy = mirrorPolicy
						for _, set := range sl {

							mirrorPolicyMapStrToI := set.(map[string]interface{})

							if v, ok := mirrorPolicyMapStrToI["cluster"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								clusterInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								mirrorPolicy.Cluster = clusterInt
								for i, ps := range sl {

									cMapToStrVal := ps.(map[string]interface{})
									clusterInt[i] = &ves_io_schema.ObjectRefType{}

									clusterInt[i].Kind = "cluster"

									if v, ok := cMapToStrVal["name"]; ok && !isIntfNil(v) {
										clusterInt[i].Name = v.(string)
									}

									if v, ok := cMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										clusterInt[i].Namespace = v.(string)
									}

									if v, ok := cMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										clusterInt[i].Tenant = v.(string)
									}

									if v, ok := cMapToStrVal["uid"]; ok && !isIntfNil(v) {
										clusterInt[i].Uid = v.(string)
									}

								}

							}

							if v, ok := mirrorPolicyMapStrToI["percent"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								percent := &ves_io_schema.FractionalPercent{}
								mirrorPolicy.Percent = percent
								for _, set := range sl {

									percentMapStrToI := set.(map[string]interface{})

									if v, ok := percentMapStrToI["denominator"]; ok && !isIntfNil(v) {

										percent.Denominator = ves_io_schema.DenominatorType(ves_io_schema.DenominatorType_value[v.(string)])

									}

									if w, ok := percentMapStrToI["numerator"]; ok && !isIntfNil(w) {
										percent.Numerator = w.(uint32)
									}

								}

							}

						}

					}

					if v, ok := cs["prefix_rewrite"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDestination.PrefixRewrite = v.(string)
					}

					if v, ok := cs["priority"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDestination.Priority = ves_io_schema.RoutingPriority(ves_io_schema.RoutingPriority_value[v.(string)])

					}

					if v, ok := cs["retry_policy"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						retryPolicy := &ves_io_schema.RetryPolicyType{}
						routeActionInt.RouteDestination.RetryPolicy = retryPolicy
						for _, set := range sl {

							retryPolicyMapStrToI := set.(map[string]interface{})

							if v, ok := retryPolicyMapStrToI["back_off"]; ok && !isIntfNil(v) {

								sl := v.(*schema.Set).List()
								backOff := &ves_io_schema.RetryBackOff{}
								retryPolicy.BackOff = backOff
								for _, set := range sl {

									backOffMapStrToI := set.(map[string]interface{})

									if w, ok := backOffMapStrToI["base_interval"]; ok && !isIntfNil(w) {
										backOff.BaseInterval = w.(uint32)
									}

									if w, ok := backOffMapStrToI["max_interval"]; ok && !isIntfNil(w) {
										backOff.MaxInterval = w.(uint32)
									}

								}

							}

							if w, ok := retryPolicyMapStrToI["num_retries"]; ok && !isIntfNil(w) {
								retryPolicy.NumRetries = w.(uint32)
							}

							if w, ok := retryPolicyMapStrToI["per_try_timeout"]; ok && !isIntfNil(w) {
								retryPolicy.PerTryTimeout = w.(uint32)
							}

							if w, ok := retryPolicyMapStrToI["retriable_status_codes"]; ok && !isIntfNil(w) {
								ls := make([]uint32, len(w.([]interface{})))
								for i, v := range w.([]interface{}) {

									ls[i] = uint32(v.(int))
								}
								retryPolicy.RetriableStatusCodes = ls
							}

							if w, ok := retryPolicyMapStrToI["retry_on"]; ok && !isIntfNil(w) {
								retryPolicy.RetryOn = w.(string)
							}

						}

					}

					if v, ok := cs["spdy_config"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						spdyConfig := &ves_io_schema_route.SpdyConfigType{}
						routeActionInt.RouteDestination.SpdyConfig = spdyConfig
						for _, set := range sl {

							spdyConfigMapStrToI := set.(map[string]interface{})

							if w, ok := spdyConfigMapStrToI["use_spdy"]; ok && !isIntfNil(w) {
								spdyConfig.UseSpdy = w.(bool)
							}

						}

					}

					if v, ok := cs["timeout"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDestination.Timeout = uint32(v.(int))
					}

					if v, ok := cs["web_socket_config"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						webSocketConfig := &ves_io_schema_route.WebsocketConfigType{}
						routeActionInt.RouteDestination.WebSocketConfig = webSocketConfig
						for _, set := range sl {

							webSocketConfigMapStrToI := set.(map[string]interface{})

							if w, ok := webSocketConfigMapStrToI["idle_timeout"]; ok && !isIntfNil(w) {
								webSocketConfig.IdleTimeout = w.(uint32)
							}

							if w, ok := webSocketConfigMapStrToI["max_connect_attempts"]; ok && !isIntfNil(w) {
								webSocketConfig.MaxConnectAttempts = w.(uint32)
							}

							if w, ok := webSocketConfigMapStrToI["use_websocket"]; ok && !isIntfNil(w) {
								webSocketConfig.UseWebsocket = w.(bool)
							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["route_direct_response"]; ok && !isIntfNil(v) && !routeActionTypeFound {

				routeActionTypeFound = true
				routeActionInt := &ves_io_schema_route.RouteType_RouteDirectResponse{}
				routeActionInt.RouteDirectResponse = &ves_io_schema_route.RouteDirectResponse{}
				routes[i].RouteAction = routeActionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["response_body"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDirectResponse.ResponseBody = v.(string)
					}

					if v, ok := cs["response_code"]; ok && !isIntfNil(v) {

						routeActionInt.RouteDirectResponse.ResponseCode = uint32(v.(int))
					}

				}

			}

			if v, ok := routesMapStrToI["route_redirect"]; ok && !isIntfNil(v) && !routeActionTypeFound {

				routeActionTypeFound = true
				routeActionInt := &ves_io_schema_route.RouteType_RouteRedirect{}
				routeActionInt.RouteRedirect = &ves_io_schema_route.RouteRedirect{}
				routes[i].RouteAction = routeActionInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					queryParamsTypeFound := false

					if v, ok := cs["all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

						queryParamsTypeFound = true
						queryParamsInt := &ves_io_schema_route.RouteRedirect_AllParams{}

						routeActionInt.RouteRedirect.QueryParams = queryParamsInt

						queryParamsInt.AllParams =
							v.(bool)

					}

					if v, ok := cs["remove_all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

						queryParamsTypeFound = true

						if v.(bool) {
							queryParamsInt := &ves_io_schema_route.RouteRedirect_RemoveAllParams{}
							queryParamsInt.RemoveAllParams = &ves_io_schema.Empty{}
							routeActionInt.RouteRedirect.QueryParams = queryParamsInt
						}

					}

					if v, ok := cs["retain_all_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

						queryParamsTypeFound = true

						if v.(bool) {
							queryParamsInt := &ves_io_schema_route.RouteRedirect_RetainAllParams{}
							queryParamsInt.RetainAllParams = &ves_io_schema.Empty{}
							routeActionInt.RouteRedirect.QueryParams = queryParamsInt
						}

					}

					if v, ok := cs["strip_query_params"]; ok && !isIntfNil(v) && !queryParamsTypeFound {

						queryParamsTypeFound = true
						queryParamsInt := &ves_io_schema_route.RouteRedirect_StripQueryParams{}
						queryParamsInt.StripQueryParams = &ves_io_schema_route.RouteQueryParams{}
						routeActionInt.RouteRedirect.QueryParams = queryParamsInt

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

					if v, ok := cs["host_redirect"]; ok && !isIntfNil(v) {

						routeActionInt.RouteRedirect.HostRedirect = v.(string)
					}

					if v, ok := cs["path_redirect"]; ok && !isIntfNil(v) {

						routeActionInt.RouteRedirect.PathRedirect = v.(string)
					}

					if v, ok := cs["proto_redirect"]; ok && !isIntfNil(v) {

						routeActionInt.RouteRedirect.ProtoRedirect = v.(string)
					}

					if v, ok := cs["response_code"]; ok && !isIntfNil(v) {

						routeActionInt.RouteRedirect.ResponseCode = uint32(v.(int))
					}

				}

			}

			if w, ok := routesMapStrToI["disable_custom_script"]; ok && !isIntfNil(w) {
				routes[i].DisableCustomScript = w.(bool)
			}

			if w, ok := routesMapStrToI["disable_location_add"]; ok && !isIntfNil(w) {
				routes[i].DisableLocationAdd = w.(bool)
			}

			if v, ok := routesMapStrToI["match"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				match := make([]*ves_io_schema.RouteMatch, len(sl))
				routes[i].Match = match
				for i, set := range sl {
					match[i] = &ves_io_schema.RouteMatch{}

					matchMapStrToI := set.(map[string]interface{})

					if v, ok := matchMapStrToI["headers"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						headers := make([]*ves_io_schema.HeaderMatcherType, len(sl))
						match[i].Headers = headers
						for i, set := range sl {
							headers[i] = &ves_io_schema.HeaderMatcherType{}

							headersMapStrToI := set.(map[string]interface{})

							if w, ok := headersMapStrToI["invert_match"]; ok && !isIntfNil(w) {
								headers[i].InvertMatch = w.(bool)
							}

							if w, ok := headersMapStrToI["name"]; ok && !isIntfNil(w) {
								headers[i].Name = w.(string)
							}

							valueMatchTypeFound := false

							if v, ok := headersMapStrToI["exact"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.HeaderMatcherType_Exact{}

								headers[i].ValueMatch = valueMatchInt

								valueMatchInt.Exact = v.(string)

							}

							if v, ok := headersMapStrToI["presence"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.HeaderMatcherType_Presence{}

								headers[i].ValueMatch = valueMatchInt

								valueMatchInt.Presence =
									v.(bool)

							}

							if v, ok := headersMapStrToI["regex"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.HeaderMatcherType_Regex{}

								headers[i].ValueMatch = valueMatchInt

								valueMatchInt.Regex = v.(string)

							}

						}

					}

					if v, ok := matchMapStrToI["http_method"]; ok && !isIntfNil(v) {

						match[i].HttpMethod = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if v, ok := matchMapStrToI["path"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						path := &ves_io_schema.PathMatcherType{}
						match[i].Path = path
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

					if v, ok := matchMapStrToI["query_params"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						queryParams := make([]*ves_io_schema.QueryParameterMatcherType, len(sl))
						match[i].QueryParams = queryParams
						for i, set := range sl {
							queryParams[i] = &ves_io_schema.QueryParameterMatcherType{}

							queryParamsMapStrToI := set.(map[string]interface{})

							if w, ok := queryParamsMapStrToI["key"]; ok && !isIntfNil(w) {
								queryParams[i].Key = w.(string)
							}

							valueMatchTypeFound := false

							if v, ok := queryParamsMapStrToI["exact"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.QueryParameterMatcherType_Exact{}

								queryParams[i].ValueMatch = valueMatchInt

								valueMatchInt.Exact = v.(string)

							}

							if v, ok := queryParamsMapStrToI["regex"]; ok && !isIntfNil(v) && !valueMatchTypeFound {

								valueMatchTypeFound = true
								valueMatchInt := &ves_io_schema.QueryParameterMatcherType_Regex{}

								queryParams[i].ValueMatch = valueMatchInt

								valueMatchInt.Regex = v.(string)

							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["request_headers_to_add"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				requestHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
				routes[i].RequestHeadersToAdd = requestHeadersToAdd
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

			if w, ok := routesMapStrToI["request_headers_to_remove"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				routes[i].RequestHeadersToRemove = ls
			}

			if v, ok := routesMapStrToI["response_headers_to_add"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				responseHeadersToAdd := make([]*ves_io_schema.HeaderManipulationOptionType, len(sl))
				routes[i].ResponseHeadersToAdd = responseHeadersToAdd
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

			if w, ok := routesMapStrToI["response_headers_to_remove"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				routes[i].ResponseHeadersToRemove = ls
			}

			if v, ok := routesMapStrToI["service_policy"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				servicePolicy := &ves_io_schema_route.ServicePolicyInfo{}
				routes[i].ServicePolicy = servicePolicy
				for _, set := range sl {

					servicePolicyMapStrToI := set.(map[string]interface{})

					if w, ok := servicePolicyMapStrToI["disable"]; ok && !isIntfNil(w) {
						servicePolicy.Disable = w.(bool)
					}

				}

			}

			if v, ok := routesMapStrToI["waf_type"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				wafType := &ves_io_schema.WafType{}
				routes[i].WafType = wafType
				for _, set := range sl {

					wafTypeMapStrToI := set.(map[string]interface{})

					refTypeTypeFound := false

					if v, ok := wafTypeMapStrToI["waf"]; ok && !isIntfNil(v) && !refTypeTypeFound {

						refTypeTypeFound = true
						refTypeInt := &ves_io_schema.WafType_Waf{}
						refTypeInt.Waf = &ves_io_schema.WafRefType{}
						wafType.RefType = refTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["waf"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								wafInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								refTypeInt.Waf.Waf = wafInt
								for i, ps := range sl {

									wMapToStrVal := ps.(map[string]interface{})
									wafInt[i] = &ves_io_schema.ObjectRefType{}

									wafInt[i].Kind = "waf"

									if v, ok := wMapToStrVal["name"]; ok && !isIntfNil(v) {
										wafInt[i].Name = v.(string)
									}

									if v, ok := wMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										wafInt[i].Namespace = v.(string)
									}

									if v, ok := wMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										wafInt[i].Tenant = v.(string)
									}

									if v, ok := wMapToStrVal["uid"]; ok && !isIntfNil(v) {
										wafInt[i].Uid = v.(string)
									}

								}

							}

						}

					}

					if v, ok := wafTypeMapStrToI["waf_rules"]; ok && !isIntfNil(v) && !refTypeTypeFound {

						refTypeTypeFound = true
						refTypeInt := &ves_io_schema.WafType_WafRules{}
						refTypeInt.WafRules = &ves_io_schema.WafRulesRefType{}
						wafType.RefType = refTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["waf_rules"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								wafRulesInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								refTypeInt.WafRules.WafRules = wafRulesInt
								for i, ps := range sl {

									wrMapToStrVal := ps.(map[string]interface{})
									wafRulesInt[i] = &ves_io_schema.ObjectRefType{}

									wafRulesInt[i].Kind = "waf_rules"

									if v, ok := wrMapToStrVal["name"]; ok && !isIntfNil(v) {
										wafRulesInt[i].Name = v.(string)
									}

									if v, ok := wrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										wafRulesInt[i].Namespace = v.(string)
									}

									if v, ok := wrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										wafRulesInt[i].Tenant = v.(string)
									}

									if v, ok := wrMapToStrVal["uid"]; ok && !isIntfNil(v) {
										wafRulesInt[i].Uid = v.(string)
									}

								}

							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra Route obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_route.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Route: %s", err)
	}

	return resourceVolterraRouteRead(d, meta)
}

func resourceVolterraRouteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_route.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Route %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Route before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Route obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_route.ObjectType, namespace, name)
}
