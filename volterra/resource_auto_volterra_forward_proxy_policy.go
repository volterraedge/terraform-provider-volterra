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
	ves_io_schema_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policy"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_views_forward_proxy_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/forward_proxy_policy"
)

// resourceVolterraForwardProxyPolicy is implementation of Volterra's ForwardProxyPolicy resources
func resourceVolterraForwardProxyPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraForwardProxyPolicyCreate,
		Read:   resourceVolterraForwardProxyPolicyRead,
		Update: resourceVolterraForwardProxyPolicyUpdate,
		Delete: resourceVolterraForwardProxyPolicyDelete,

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

			"any_proxy": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"drp_http_connect": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"network_connector": {

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

			"proxy_label_selector": {

				Type:     schema.TypeSet,
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

			"allow_all": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"allow_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"default_action_allow": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_action_deny": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_action_next_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"dest_list": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ports": {
										Type:     schema.TypeString,
										Optional: true,
									},

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

						"http_list": {

							Type:     schema.TypeList,
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

									"any_path": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"path_exact_value": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"path_prefix_value": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"path_regex_value": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"rule_description": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"rule_name": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"tls_list": {

							Type:     schema.TypeList,
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

			"deny_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"default_action_allow": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_action_deny": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"default_action_next_policy": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"dest_list": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ports": {
										Type:     schema.TypeString,
										Optional: true,
									},

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

						"http_list": {

							Type:     schema.TypeList,
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

									"any_path": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"path_exact_value": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"path_prefix_value": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"path_regex_value": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"rule_description": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"rule_name": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"tls_list": {

							Type:     schema.TypeList,
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

			"rule_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"all_destinations": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"dest_list": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dest_list": {

													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ports": {
																Type:     schema.TypeString,
																Optional: true,
															},

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
											},
										},
									},

									"http_list": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"http_list": {

													Type:     schema.TypeList,
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

															"any_path": {

																Type:     schema.TypeBool,
																Optional: true,
															},

															"path_exact_value": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"path_prefix_value": {

																Type:     schema.TypeString,
																Optional: true,
															},

															"path_regex_value": {

																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},

									"tls_list": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"tls_list": {

													Type:     schema.TypeList,
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

									"no_http_connect_port": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"port_matcher": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"invert_matcher": {
													Type:     schema.TypeBool,
													Optional: true,
												},

												"ports": {

													Type: schema.TypeList,

													Required: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"rule_description": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"rule_name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"all_sources": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"inside_sources": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"interface": {

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

									"ip_prefix_set": {

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

									"label_selector": {

										Type:     schema.TypeSet,
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

									"namespace": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"prefix_list": {

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
								},
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraForwardProxyPolicyCreate creates ForwardProxyPolicy resource
func resourceVolterraForwardProxyPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_forward_proxy_policy.CreateSpecType{}
	createReq := &ves_io_schema_views_forward_proxy_policy.CreateRequest{
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

	proxyChoiceTypeFound := false

	if v, ok := d.GetOk("any_proxy"); ok && !proxyChoiceTypeFound {

		proxyChoiceTypeFound = true

		if v.(bool) {
			proxyChoiceInt := &ves_io_schema_views_forward_proxy_policy.CreateSpecType_AnyProxy{}
			proxyChoiceInt.AnyProxy = &ves_io_schema.Empty{}
			createSpec.ProxyChoice = proxyChoiceInt
		}

	}

	if v, ok := d.GetOk("drp_http_connect"); ok && !proxyChoiceTypeFound {

		proxyChoiceTypeFound = true

		if v.(bool) {
			proxyChoiceInt := &ves_io_schema_views_forward_proxy_policy.CreateSpecType_DrpHttpConnect{}
			proxyChoiceInt.DrpHttpConnect = &ves_io_schema.Empty{}
			createSpec.ProxyChoice = proxyChoiceInt
		}

	}

	if v, ok := d.GetOk("network_connector"); ok && !proxyChoiceTypeFound {

		proxyChoiceTypeFound = true
		proxyChoiceInt := &ves_io_schema_views_forward_proxy_policy.CreateSpecType_NetworkConnector{}
		proxyChoiceInt.NetworkConnector = &ves_io_schema_views.ObjectRefType{}
		createSpec.ProxyChoice = proxyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				proxyChoiceInt.NetworkConnector.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				proxyChoiceInt.NetworkConnector.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				proxyChoiceInt.NetworkConnector.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("proxy_label_selector"); ok && !proxyChoiceTypeFound {

		proxyChoiceTypeFound = true
		proxyChoiceInt := &ves_io_schema_views_forward_proxy_policy.CreateSpecType_ProxyLabelSelector{}
		proxyChoiceInt.ProxyLabelSelector = &ves_io_schema.LabelSelectorType{}
		createSpec.ProxyChoice = proxyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				proxyChoiceInt.ProxyLabelSelector.Expressions = ls

			}

		}

	}

	ruleChoiceTypeFound := false

	if v, ok := d.GetOk("allow_all"); ok && !ruleChoiceTypeFound {

		ruleChoiceTypeFound = true

		if v.(bool) {
			ruleChoiceInt := &ves_io_schema_views_forward_proxy_policy.CreateSpecType_AllowAll{}
			ruleChoiceInt.AllowAll = &ves_io_schema.Empty{}
			createSpec.RuleChoice = ruleChoiceInt
		}

	}

	if v, ok := d.GetOk("allow_list"); ok && !ruleChoiceTypeFound {

		ruleChoiceTypeFound = true
		ruleChoiceInt := &ves_io_schema_views_forward_proxy_policy.CreateSpecType_AllowList{}
		ruleChoiceInt.AllowList = &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType{}
		createSpec.RuleChoice = ruleChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			defaultActionChoiceTypeFound := false

			if v, ok := cs["default_action_allow"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionAllow{}
					defaultActionChoiceInt.DefaultActionAllow = &ves_io_schema.Empty{}
					ruleChoiceInt.AllowList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["default_action_deny"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionDeny{}
					defaultActionChoiceInt.DefaultActionDeny = &ves_io_schema.Empty{}
					ruleChoiceInt.AllowList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["default_action_next_policy"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionNextPolicy{}
					defaultActionChoiceInt.DefaultActionNextPolicy = &ves_io_schema.Empty{}
					ruleChoiceInt.AllowList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["dest_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				destList := make([]*ves_io_schema.L4DestType, len(sl))
				ruleChoiceInt.AllowList.DestList = destList
				for i, set := range sl {
					destList[i] = &ves_io_schema.L4DestType{}

					destListMapStrToI := set.(map[string]interface{})

					if w, ok := destListMapStrToI["ports"]; ok && !isIntfNil(w) {
						destList[i].Ports = w.(string)
					}

					if w, ok := destListMapStrToI["prefixes"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						destList[i].Prefixes = ls
					}

				}

			}

			if v, ok := cs["http_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				httpList := make([]*ves_io_schema_views_forward_proxy_policy.URLType, len(sl))
				ruleChoiceInt.AllowList.HttpList = httpList
				for i, set := range sl {
					httpList[i] = &ves_io_schema_views_forward_proxy_policy.URLType{}

					httpListMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := httpListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_ExactValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.ExactValue = v.(string)

					}

					if v, ok := httpListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_RegexValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.RegexValue = v.(string)

					}

					if v, ok := httpListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_SuffixValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.SuffixValue = v.(string)

					}

					pathChoiceTypeFound := false

					if v, ok := httpListMapStrToI["any_path"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true

						if v.(bool) {
							pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_AnyPath{}
							pathChoiceInt.AnyPath = &ves_io_schema.Empty{}
							httpList[i].PathChoice = pathChoiceInt
						}

					}

					if v, ok := httpListMapStrToI["path_exact_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathExactValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathExactValue = v.(string)

					}

					if v, ok := httpListMapStrToI["path_prefix_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathPrefixValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathPrefixValue = v.(string)

					}

					if v, ok := httpListMapStrToI["path_regex_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathRegexValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathRegexValue = v.(string)

					}

				}

			}

			if v, ok := cs["rule_description"]; ok && !isIntfNil(v) {

				ruleChoiceInt.AllowList.RuleDescription = v.(string)
			}

			if v, ok := cs["rule_name"]; ok && !isIntfNil(v) {

				ruleChoiceInt.AllowList.RuleName = v.(string)
			}

			if v, ok := cs["tls_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				tlsList := make([]*ves_io_schema.DomainType, len(sl))
				ruleChoiceInt.AllowList.TlsList = tlsList
				for i, set := range sl {
					tlsList[i] = &ves_io_schema.DomainType{}

					tlsListMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := tlsListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_ExactValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.ExactValue = v.(string)

					}

					if v, ok := tlsListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_RegexValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.RegexValue = v.(string)

					}

					if v, ok := tlsListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_SuffixValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.SuffixValue = v.(string)

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("deny_list"); ok && !ruleChoiceTypeFound {

		ruleChoiceTypeFound = true
		ruleChoiceInt := &ves_io_schema_views_forward_proxy_policy.CreateSpecType_DenyList{}
		ruleChoiceInt.DenyList = &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType{}
		createSpec.RuleChoice = ruleChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			defaultActionChoiceTypeFound := false

			if v, ok := cs["default_action_allow"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionAllow{}
					defaultActionChoiceInt.DefaultActionAllow = &ves_io_schema.Empty{}
					ruleChoiceInt.DenyList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["default_action_deny"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionDeny{}
					defaultActionChoiceInt.DefaultActionDeny = &ves_io_schema.Empty{}
					ruleChoiceInt.DenyList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["default_action_next_policy"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionNextPolicy{}
					defaultActionChoiceInt.DefaultActionNextPolicy = &ves_io_schema.Empty{}
					ruleChoiceInt.DenyList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["dest_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				destList := make([]*ves_io_schema.L4DestType, len(sl))
				ruleChoiceInt.DenyList.DestList = destList
				for i, set := range sl {
					destList[i] = &ves_io_schema.L4DestType{}

					destListMapStrToI := set.(map[string]interface{})

					if w, ok := destListMapStrToI["ports"]; ok && !isIntfNil(w) {
						destList[i].Ports = w.(string)
					}

					if w, ok := destListMapStrToI["prefixes"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						destList[i].Prefixes = ls
					}

				}

			}

			if v, ok := cs["http_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				httpList := make([]*ves_io_schema_views_forward_proxy_policy.URLType, len(sl))
				ruleChoiceInt.DenyList.HttpList = httpList
				for i, set := range sl {
					httpList[i] = &ves_io_schema_views_forward_proxy_policy.URLType{}

					httpListMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := httpListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_ExactValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.ExactValue = v.(string)

					}

					if v, ok := httpListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_RegexValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.RegexValue = v.(string)

					}

					if v, ok := httpListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_SuffixValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.SuffixValue = v.(string)

					}

					pathChoiceTypeFound := false

					if v, ok := httpListMapStrToI["any_path"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true

						if v.(bool) {
							pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_AnyPath{}
							pathChoiceInt.AnyPath = &ves_io_schema.Empty{}
							httpList[i].PathChoice = pathChoiceInt
						}

					}

					if v, ok := httpListMapStrToI["path_exact_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathExactValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathExactValue = v.(string)

					}

					if v, ok := httpListMapStrToI["path_prefix_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathPrefixValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathPrefixValue = v.(string)

					}

					if v, ok := httpListMapStrToI["path_regex_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathRegexValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathRegexValue = v.(string)

					}

				}

			}

			if v, ok := cs["rule_description"]; ok && !isIntfNil(v) {

				ruleChoiceInt.DenyList.RuleDescription = v.(string)
			}

			if v, ok := cs["rule_name"]; ok && !isIntfNil(v) {

				ruleChoiceInt.DenyList.RuleName = v.(string)
			}

			if v, ok := cs["tls_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				tlsList := make([]*ves_io_schema.DomainType, len(sl))
				ruleChoiceInt.DenyList.TlsList = tlsList
				for i, set := range sl {
					tlsList[i] = &ves_io_schema.DomainType{}

					tlsListMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := tlsListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_ExactValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.ExactValue = v.(string)

					}

					if v, ok := tlsListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_RegexValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.RegexValue = v.(string)

					}

					if v, ok := tlsListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_SuffixValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.SuffixValue = v.(string)

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("rule_list"); ok && !ruleChoiceTypeFound {

		ruleChoiceTypeFound = true
		ruleChoiceInt := &ves_io_schema_views_forward_proxy_policy.CreateSpecType_RuleList{}
		ruleChoiceInt.RuleList = &ves_io_schema_views_forward_proxy_policy.ForwardProxyRuleListType{}
		createSpec.RuleChoice = ruleChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				rules := make([]*ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType, len(sl))
				ruleChoiceInt.RuleList.Rules = rules
				for i, set := range sl {
					rules[i] = &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType{}

					rulesMapStrToI := set.(map[string]interface{})

					if v, ok := rulesMapStrToI["action"]; ok && !isIntfNil(v) {

						rules[i].Action = ves_io_schema_policy.RuleAction(ves_io_schema_policy.RuleAction_value[v.(string)])

					}

					destinationChoiceTypeFound := false

					if v, ok := rulesMapStrToI["all_destinations"]; ok && !isIntfNil(v) && !destinationChoiceTypeFound {

						destinationChoiceTypeFound = true

						if v.(bool) {
							destinationChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_AllDestinations{}
							destinationChoiceInt.AllDestinations = &ves_io_schema.Empty{}
							rules[i].DestinationChoice = destinationChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["dest_list"]; ok && !isIntfNil(v) && !destinationChoiceTypeFound {

						destinationChoiceTypeFound = true
						destinationChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_DestList{}
						destinationChoiceInt.DestList = &ves_io_schema_views_forward_proxy_policy.L4DestListType{}
						rules[i].DestinationChoice = destinationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["dest_list"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								destList := make([]*ves_io_schema.L4DestType, len(sl))
								destinationChoiceInt.DestList.DestList = destList
								for i, set := range sl {
									destList[i] = &ves_io_schema.L4DestType{}

									destListMapStrToI := set.(map[string]interface{})

									if w, ok := destListMapStrToI["ports"]; ok && !isIntfNil(w) {
										destList[i].Ports = w.(string)
									}

									if w, ok := destListMapStrToI["prefixes"]; ok && !isIntfNil(w) {
										ls := make([]string, len(w.([]interface{})))
										for i, v := range w.([]interface{}) {
											ls[i] = v.(string)
										}
										destList[i].Prefixes = ls
									}

								}

							}

						}

					}

					if v, ok := rulesMapStrToI["http_list"]; ok && !isIntfNil(v) && !destinationChoiceTypeFound {

						destinationChoiceTypeFound = true
						destinationChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_HttpList{}
						destinationChoiceInt.HttpList = &ves_io_schema_views_forward_proxy_policy.URLListType{}
						rules[i].DestinationChoice = destinationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["http_list"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								httpList := make([]*ves_io_schema_views_forward_proxy_policy.URLType, len(sl))
								destinationChoiceInt.HttpList.HttpList = httpList
								for i, set := range sl {
									httpList[i] = &ves_io_schema_views_forward_proxy_policy.URLType{}

									httpListMapStrToI := set.(map[string]interface{})

									domainChoiceTypeFound := false

									if v, ok := httpListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_ExactValue{}

										httpList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.ExactValue = v.(string)

									}

									if v, ok := httpListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_RegexValue{}

										httpList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.RegexValue = v.(string)

									}

									if v, ok := httpListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_SuffixValue{}

										httpList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.SuffixValue = v.(string)

									}

									pathChoiceTypeFound := false

									if v, ok := httpListMapStrToI["any_path"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

										pathChoiceTypeFound = true

										if v.(bool) {
											pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_AnyPath{}
											pathChoiceInt.AnyPath = &ves_io_schema.Empty{}
											httpList[i].PathChoice = pathChoiceInt
										}

									}

									if v, ok := httpListMapStrToI["path_exact_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

										pathChoiceTypeFound = true
										pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathExactValue{}

										httpList[i].PathChoice = pathChoiceInt

										pathChoiceInt.PathExactValue = v.(string)

									}

									if v, ok := httpListMapStrToI["path_prefix_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

										pathChoiceTypeFound = true
										pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathPrefixValue{}

										httpList[i].PathChoice = pathChoiceInt

										pathChoiceInt.PathPrefixValue = v.(string)

									}

									if v, ok := httpListMapStrToI["path_regex_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

										pathChoiceTypeFound = true
										pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathRegexValue{}

										httpList[i].PathChoice = pathChoiceInt

										pathChoiceInt.PathRegexValue = v.(string)

									}

								}

							}

						}

					}

					if v, ok := rulesMapStrToI["tls_list"]; ok && !isIntfNil(v) && !destinationChoiceTypeFound {

						destinationChoiceTypeFound = true
						destinationChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_TlsList{}
						destinationChoiceInt.TlsList = &ves_io_schema_views_forward_proxy_policy.DomainListType{}
						rules[i].DestinationChoice = destinationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["tls_list"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								tlsList := make([]*ves_io_schema.DomainType, len(sl))
								destinationChoiceInt.TlsList.TlsList = tlsList
								for i, set := range sl {
									tlsList[i] = &ves_io_schema.DomainType{}

									tlsListMapStrToI := set.(map[string]interface{})

									domainChoiceTypeFound := false

									if v, ok := tlsListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema.DomainType_ExactValue{}

										tlsList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.ExactValue = v.(string)

									}

									if v, ok := tlsListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema.DomainType_RegexValue{}

										tlsList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.RegexValue = v.(string)

									}

									if v, ok := tlsListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema.DomainType_SuffixValue{}

										tlsList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.SuffixValue = v.(string)

									}

								}

							}

						}

					}

					httpConnectChoiceTypeFound := false

					if v, ok := rulesMapStrToI["no_http_connect_port"]; ok && !isIntfNil(v) && !httpConnectChoiceTypeFound {

						httpConnectChoiceTypeFound = true

						if v.(bool) {
							httpConnectChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_NoHttpConnectPort{}
							httpConnectChoiceInt.NoHttpConnectPort = &ves_io_schema.Empty{}
							rules[i].HttpConnectChoice = httpConnectChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["port_matcher"]; ok && !isIntfNil(v) && !httpConnectChoiceTypeFound {

						httpConnectChoiceTypeFound = true
						httpConnectChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_PortMatcher{}
						httpConnectChoiceInt.PortMatcher = &ves_io_schema_policy.PortMatcherType{}
						rules[i].HttpConnectChoice = httpConnectChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["invert_matcher"]; ok && !isIntfNil(v) {

								httpConnectChoiceInt.PortMatcher.InvertMatcher = v.(bool)
							}

							if v, ok := cs["ports"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								httpConnectChoiceInt.PortMatcher.Ports = ls

							}

						}

					}

					if w, ok := rulesMapStrToI["rule_description"]; ok && !isIntfNil(w) {
						rules[i].RuleDescription = w.(string)
					}

					if w, ok := rulesMapStrToI["rule_name"]; ok && !isIntfNil(w) {
						rules[i].RuleName = w.(string)
					}

					sourceChoiceTypeFound := false

					if v, ok := rulesMapStrToI["all_sources"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true

						if v.(bool) {
							sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_AllSources{}
							sourceChoiceInt.AllSources = &ves_io_schema.Empty{}
							rules[i].SourceChoice = sourceChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["inside_sources"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true

						if v.(bool) {
							sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_InsideSources{}
							sourceChoiceInt.InsideSources = &ves_io_schema.Empty{}
							rules[i].SourceChoice = sourceChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["interface"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_Interface{}
						sourceChoiceInt.Interface = &ves_io_schema_views.ObjectRefType{}
						rules[i].SourceChoice = sourceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								sourceChoiceInt.Interface.Name = v.(string)
							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								sourceChoiceInt.Interface.Namespace = v.(string)
							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								sourceChoiceInt.Interface.Tenant = v.(string)
							}

						}

					}

					if v, ok := rulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_IpPrefixSet{}
						sourceChoiceInt.IpPrefixSet = &ves_io_schema_views.ObjectRefType{}
						rules[i].SourceChoice = sourceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								sourceChoiceInt.IpPrefixSet.Name = v.(string)
							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								sourceChoiceInt.IpPrefixSet.Namespace = v.(string)
							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								sourceChoiceInt.IpPrefixSet.Tenant = v.(string)
							}

						}

					}

					if v, ok := rulesMapStrToI["label_selector"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_LabelSelector{}
						sourceChoiceInt.LabelSelector = &ves_io_schema.LabelSelectorType{}
						rules[i].SourceChoice = sourceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								sourceChoiceInt.LabelSelector.Expressions = ls

							}

						}

					}

					if v, ok := rulesMapStrToI["namespace"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_Namespace{}

						rules[i].SourceChoice = sourceChoiceInt

						sourceChoiceInt.Namespace = v.(string)

					}

					if v, ok := rulesMapStrToI["prefix_list"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_PrefixList{}
						sourceChoiceInt.PrefixList = &ves_io_schema_views.PrefixStringListType{}
						rules[i].SourceChoice = sourceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								sourceChoiceInt.PrefixList.Prefixes = ls

							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra ForwardProxyPolicy object with struct: %+v", createReq)

	createForwardProxyPolicyResp, err := client.CreateObject(context.Background(), ves_io_schema_views_forward_proxy_policy.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating ForwardProxyPolicy: %s", err)
	}
	d.SetId(createForwardProxyPolicyResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraForwardProxyPolicyRead(d, meta)
}

func resourceVolterraForwardProxyPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_forward_proxy_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ForwardProxyPolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ForwardProxyPolicy %q: %s", d.Id(), err)
	}
	return setForwardProxyPolicyFields(client, d, resp)
}

func setForwardProxyPolicyFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraForwardProxyPolicyUpdate updates ForwardProxyPolicy resource
func resourceVolterraForwardProxyPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_forward_proxy_policy.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_forward_proxy_policy.ReplaceRequest{
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

	proxyChoiceTypeFound := false

	if v, ok := d.GetOk("any_proxy"); ok && !proxyChoiceTypeFound {

		proxyChoiceTypeFound = true

		if v.(bool) {
			proxyChoiceInt := &ves_io_schema_views_forward_proxy_policy.ReplaceSpecType_AnyProxy{}
			proxyChoiceInt.AnyProxy = &ves_io_schema.Empty{}
			updateSpec.ProxyChoice = proxyChoiceInt
		}

	}

	if v, ok := d.GetOk("drp_http_connect"); ok && !proxyChoiceTypeFound {

		proxyChoiceTypeFound = true

		if v.(bool) {
			proxyChoiceInt := &ves_io_schema_views_forward_proxy_policy.ReplaceSpecType_DrpHttpConnect{}
			proxyChoiceInt.DrpHttpConnect = &ves_io_schema.Empty{}
			updateSpec.ProxyChoice = proxyChoiceInt
		}

	}

	if v, ok := d.GetOk("network_connector"); ok && !proxyChoiceTypeFound {

		proxyChoiceTypeFound = true
		proxyChoiceInt := &ves_io_schema_views_forward_proxy_policy.ReplaceSpecType_NetworkConnector{}
		proxyChoiceInt.NetworkConnector = &ves_io_schema_views.ObjectRefType{}
		updateSpec.ProxyChoice = proxyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				proxyChoiceInt.NetworkConnector.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				proxyChoiceInt.NetworkConnector.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				proxyChoiceInt.NetworkConnector.Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("proxy_label_selector"); ok && !proxyChoiceTypeFound {

		proxyChoiceTypeFound = true
		proxyChoiceInt := &ves_io_schema_views_forward_proxy_policy.ReplaceSpecType_ProxyLabelSelector{}
		proxyChoiceInt.ProxyLabelSelector = &ves_io_schema.LabelSelectorType{}
		updateSpec.ProxyChoice = proxyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				proxyChoiceInt.ProxyLabelSelector.Expressions = ls

			}

		}

	}

	ruleChoiceTypeFound := false

	if v, ok := d.GetOk("allow_all"); ok && !ruleChoiceTypeFound {

		ruleChoiceTypeFound = true

		if v.(bool) {
			ruleChoiceInt := &ves_io_schema_views_forward_proxy_policy.ReplaceSpecType_AllowAll{}
			ruleChoiceInt.AllowAll = &ves_io_schema.Empty{}
			updateSpec.RuleChoice = ruleChoiceInt
		}

	}

	if v, ok := d.GetOk("allow_list"); ok && !ruleChoiceTypeFound {

		ruleChoiceTypeFound = true
		ruleChoiceInt := &ves_io_schema_views_forward_proxy_policy.ReplaceSpecType_AllowList{}
		ruleChoiceInt.AllowList = &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType{}
		updateSpec.RuleChoice = ruleChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			defaultActionChoiceTypeFound := false

			if v, ok := cs["default_action_allow"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionAllow{}
					defaultActionChoiceInt.DefaultActionAllow = &ves_io_schema.Empty{}
					ruleChoiceInt.AllowList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["default_action_deny"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionDeny{}
					defaultActionChoiceInt.DefaultActionDeny = &ves_io_schema.Empty{}
					ruleChoiceInt.AllowList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["default_action_next_policy"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionNextPolicy{}
					defaultActionChoiceInt.DefaultActionNextPolicy = &ves_io_schema.Empty{}
					ruleChoiceInt.AllowList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["dest_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				destList := make([]*ves_io_schema.L4DestType, len(sl))
				ruleChoiceInt.AllowList.DestList = destList
				for i, set := range sl {
					destList[i] = &ves_io_schema.L4DestType{}

					destListMapStrToI := set.(map[string]interface{})

					if w, ok := destListMapStrToI["ports"]; ok && !isIntfNil(w) {
						destList[i].Ports = w.(string)
					}

					if w, ok := destListMapStrToI["prefixes"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						destList[i].Prefixes = ls
					}

				}

			}

			if v, ok := cs["http_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				httpList := make([]*ves_io_schema_views_forward_proxy_policy.URLType, len(sl))
				ruleChoiceInt.AllowList.HttpList = httpList
				for i, set := range sl {
					httpList[i] = &ves_io_schema_views_forward_proxy_policy.URLType{}

					httpListMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := httpListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_ExactValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.ExactValue = v.(string)

					}

					if v, ok := httpListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_RegexValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.RegexValue = v.(string)

					}

					if v, ok := httpListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_SuffixValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.SuffixValue = v.(string)

					}

					pathChoiceTypeFound := false

					if v, ok := httpListMapStrToI["any_path"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true

						if v.(bool) {
							pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_AnyPath{}
							pathChoiceInt.AnyPath = &ves_io_schema.Empty{}
							httpList[i].PathChoice = pathChoiceInt
						}

					}

					if v, ok := httpListMapStrToI["path_exact_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathExactValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathExactValue = v.(string)

					}

					if v, ok := httpListMapStrToI["path_prefix_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathPrefixValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathPrefixValue = v.(string)

					}

					if v, ok := httpListMapStrToI["path_regex_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathRegexValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathRegexValue = v.(string)

					}

				}

			}

			if v, ok := cs["rule_description"]; ok && !isIntfNil(v) {

				ruleChoiceInt.AllowList.RuleDescription = v.(string)
			}

			if v, ok := cs["rule_name"]; ok && !isIntfNil(v) {

				ruleChoiceInt.AllowList.RuleName = v.(string)
			}

			if v, ok := cs["tls_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				tlsList := make([]*ves_io_schema.DomainType, len(sl))
				ruleChoiceInt.AllowList.TlsList = tlsList
				for i, set := range sl {
					tlsList[i] = &ves_io_schema.DomainType{}

					tlsListMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := tlsListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_ExactValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.ExactValue = v.(string)

					}

					if v, ok := tlsListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_RegexValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.RegexValue = v.(string)

					}

					if v, ok := tlsListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_SuffixValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.SuffixValue = v.(string)

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("deny_list"); ok && !ruleChoiceTypeFound {

		ruleChoiceTypeFound = true
		ruleChoiceInt := &ves_io_schema_views_forward_proxy_policy.ReplaceSpecType_DenyList{}
		ruleChoiceInt.DenyList = &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType{}
		updateSpec.RuleChoice = ruleChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			defaultActionChoiceTypeFound := false

			if v, ok := cs["default_action_allow"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionAllow{}
					defaultActionChoiceInt.DefaultActionAllow = &ves_io_schema.Empty{}
					ruleChoiceInt.DenyList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["default_action_deny"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionDeny{}
					defaultActionChoiceInt.DefaultActionDeny = &ves_io_schema.Empty{}
					ruleChoiceInt.DenyList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["default_action_next_policy"]; ok && !isIntfNil(v) && !defaultActionChoiceTypeFound {

				defaultActionChoiceTypeFound = true

				if v.(bool) {
					defaultActionChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxySimpleRuleType_DefaultActionNextPolicy{}
					defaultActionChoiceInt.DefaultActionNextPolicy = &ves_io_schema.Empty{}
					ruleChoiceInt.DenyList.DefaultActionChoice = defaultActionChoiceInt
				}

			}

			if v, ok := cs["dest_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				destList := make([]*ves_io_schema.L4DestType, len(sl))
				ruleChoiceInt.DenyList.DestList = destList
				for i, set := range sl {
					destList[i] = &ves_io_schema.L4DestType{}

					destListMapStrToI := set.(map[string]interface{})

					if w, ok := destListMapStrToI["ports"]; ok && !isIntfNil(w) {
						destList[i].Ports = w.(string)
					}

					if w, ok := destListMapStrToI["prefixes"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						destList[i].Prefixes = ls
					}

				}

			}

			if v, ok := cs["http_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				httpList := make([]*ves_io_schema_views_forward_proxy_policy.URLType, len(sl))
				ruleChoiceInt.DenyList.HttpList = httpList
				for i, set := range sl {
					httpList[i] = &ves_io_schema_views_forward_proxy_policy.URLType{}

					httpListMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := httpListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_ExactValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.ExactValue = v.(string)

					}

					if v, ok := httpListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_RegexValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.RegexValue = v.(string)

					}

					if v, ok := httpListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_SuffixValue{}

						httpList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.SuffixValue = v.(string)

					}

					pathChoiceTypeFound := false

					if v, ok := httpListMapStrToI["any_path"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true

						if v.(bool) {
							pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_AnyPath{}
							pathChoiceInt.AnyPath = &ves_io_schema.Empty{}
							httpList[i].PathChoice = pathChoiceInt
						}

					}

					if v, ok := httpListMapStrToI["path_exact_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathExactValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathExactValue = v.(string)

					}

					if v, ok := httpListMapStrToI["path_prefix_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathPrefixValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathPrefixValue = v.(string)

					}

					if v, ok := httpListMapStrToI["path_regex_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathRegexValue{}

						httpList[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathRegexValue = v.(string)

					}

				}

			}

			if v, ok := cs["rule_description"]; ok && !isIntfNil(v) {

				ruleChoiceInt.DenyList.RuleDescription = v.(string)
			}

			if v, ok := cs["rule_name"]; ok && !isIntfNil(v) {

				ruleChoiceInt.DenyList.RuleName = v.(string)
			}

			if v, ok := cs["tls_list"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				tlsList := make([]*ves_io_schema.DomainType, len(sl))
				ruleChoiceInt.DenyList.TlsList = tlsList
				for i, set := range sl {
					tlsList[i] = &ves_io_schema.DomainType{}

					tlsListMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := tlsListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_ExactValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.ExactValue = v.(string)

					}

					if v, ok := tlsListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_RegexValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.RegexValue = v.(string)

					}

					if v, ok := tlsListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema.DomainType_SuffixValue{}

						tlsList[i].DomainChoice = domainChoiceInt

						domainChoiceInt.SuffixValue = v.(string)

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("rule_list"); ok && !ruleChoiceTypeFound {

		ruleChoiceTypeFound = true
		ruleChoiceInt := &ves_io_schema_views_forward_proxy_policy.ReplaceSpecType_RuleList{}
		ruleChoiceInt.RuleList = &ves_io_schema_views_forward_proxy_policy.ForwardProxyRuleListType{}
		updateSpec.RuleChoice = ruleChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				rules := make([]*ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType, len(sl))
				ruleChoiceInt.RuleList.Rules = rules
				for i, set := range sl {
					rules[i] = &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType{}

					rulesMapStrToI := set.(map[string]interface{})

					if v, ok := rulesMapStrToI["action"]; ok && !isIntfNil(v) {

						rules[i].Action = ves_io_schema_policy.RuleAction(ves_io_schema_policy.RuleAction_value[v.(string)])

					}

					destinationChoiceTypeFound := false

					if v, ok := rulesMapStrToI["all_destinations"]; ok && !isIntfNil(v) && !destinationChoiceTypeFound {

						destinationChoiceTypeFound = true

						if v.(bool) {
							destinationChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_AllDestinations{}
							destinationChoiceInt.AllDestinations = &ves_io_schema.Empty{}
							rules[i].DestinationChoice = destinationChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["dest_list"]; ok && !isIntfNil(v) && !destinationChoiceTypeFound {

						destinationChoiceTypeFound = true
						destinationChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_DestList{}
						destinationChoiceInt.DestList = &ves_io_schema_views_forward_proxy_policy.L4DestListType{}
						rules[i].DestinationChoice = destinationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["dest_list"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								destList := make([]*ves_io_schema.L4DestType, len(sl))
								destinationChoiceInt.DestList.DestList = destList
								for i, set := range sl {
									destList[i] = &ves_io_schema.L4DestType{}

									destListMapStrToI := set.(map[string]interface{})

									if w, ok := destListMapStrToI["ports"]; ok && !isIntfNil(w) {
										destList[i].Ports = w.(string)
									}

									if w, ok := destListMapStrToI["prefixes"]; ok && !isIntfNil(w) {
										ls := make([]string, len(w.([]interface{})))
										for i, v := range w.([]interface{}) {
											ls[i] = v.(string)
										}
										destList[i].Prefixes = ls
									}

								}

							}

						}

					}

					if v, ok := rulesMapStrToI["http_list"]; ok && !isIntfNil(v) && !destinationChoiceTypeFound {

						destinationChoiceTypeFound = true
						destinationChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_HttpList{}
						destinationChoiceInt.HttpList = &ves_io_schema_views_forward_proxy_policy.URLListType{}
						rules[i].DestinationChoice = destinationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["http_list"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								httpList := make([]*ves_io_schema_views_forward_proxy_policy.URLType, len(sl))
								destinationChoiceInt.HttpList.HttpList = httpList
								for i, set := range sl {
									httpList[i] = &ves_io_schema_views_forward_proxy_policy.URLType{}

									httpListMapStrToI := set.(map[string]interface{})

									domainChoiceTypeFound := false

									if v, ok := httpListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_ExactValue{}

										httpList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.ExactValue = v.(string)

									}

									if v, ok := httpListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_RegexValue{}

										httpList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.RegexValue = v.(string)

									}

									if v, ok := httpListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_SuffixValue{}

										httpList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.SuffixValue = v.(string)

									}

									pathChoiceTypeFound := false

									if v, ok := httpListMapStrToI["any_path"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

										pathChoiceTypeFound = true

										if v.(bool) {
											pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_AnyPath{}
											pathChoiceInt.AnyPath = &ves_io_schema.Empty{}
											httpList[i].PathChoice = pathChoiceInt
										}

									}

									if v, ok := httpListMapStrToI["path_exact_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

										pathChoiceTypeFound = true
										pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathExactValue{}

										httpList[i].PathChoice = pathChoiceInt

										pathChoiceInt.PathExactValue = v.(string)

									}

									if v, ok := httpListMapStrToI["path_prefix_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

										pathChoiceTypeFound = true
										pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathPrefixValue{}

										httpList[i].PathChoice = pathChoiceInt

										pathChoiceInt.PathPrefixValue = v.(string)

									}

									if v, ok := httpListMapStrToI["path_regex_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

										pathChoiceTypeFound = true
										pathChoiceInt := &ves_io_schema_views_forward_proxy_policy.URLType_PathRegexValue{}

										httpList[i].PathChoice = pathChoiceInt

										pathChoiceInt.PathRegexValue = v.(string)

									}

								}

							}

						}

					}

					if v, ok := rulesMapStrToI["tls_list"]; ok && !isIntfNil(v) && !destinationChoiceTypeFound {

						destinationChoiceTypeFound = true
						destinationChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_TlsList{}
						destinationChoiceInt.TlsList = &ves_io_schema_views_forward_proxy_policy.DomainListType{}
						rules[i].DestinationChoice = destinationChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["tls_list"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								tlsList := make([]*ves_io_schema.DomainType, len(sl))
								destinationChoiceInt.TlsList.TlsList = tlsList
								for i, set := range sl {
									tlsList[i] = &ves_io_schema.DomainType{}

									tlsListMapStrToI := set.(map[string]interface{})

									domainChoiceTypeFound := false

									if v, ok := tlsListMapStrToI["exact_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema.DomainType_ExactValue{}

										tlsList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.ExactValue = v.(string)

									}

									if v, ok := tlsListMapStrToI["regex_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema.DomainType_RegexValue{}

										tlsList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.RegexValue = v.(string)

									}

									if v, ok := tlsListMapStrToI["suffix_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

										domainChoiceTypeFound = true
										domainChoiceInt := &ves_io_schema.DomainType_SuffixValue{}

										tlsList[i].DomainChoice = domainChoiceInt

										domainChoiceInt.SuffixValue = v.(string)

									}

								}

							}

						}

					}

					httpConnectChoiceTypeFound := false

					if v, ok := rulesMapStrToI["no_http_connect_port"]; ok && !isIntfNil(v) && !httpConnectChoiceTypeFound {

						httpConnectChoiceTypeFound = true

						if v.(bool) {
							httpConnectChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_NoHttpConnectPort{}
							httpConnectChoiceInt.NoHttpConnectPort = &ves_io_schema.Empty{}
							rules[i].HttpConnectChoice = httpConnectChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["port_matcher"]; ok && !isIntfNil(v) && !httpConnectChoiceTypeFound {

						httpConnectChoiceTypeFound = true
						httpConnectChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_PortMatcher{}
						httpConnectChoiceInt.PortMatcher = &ves_io_schema_policy.PortMatcherType{}
						rules[i].HttpConnectChoice = httpConnectChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["invert_matcher"]; ok && !isIntfNil(v) {

								httpConnectChoiceInt.PortMatcher.InvertMatcher = v.(bool)
							}

							if v, ok := cs["ports"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								httpConnectChoiceInt.PortMatcher.Ports = ls

							}

						}

					}

					if w, ok := rulesMapStrToI["rule_description"]; ok && !isIntfNil(w) {
						rules[i].RuleDescription = w.(string)
					}

					if w, ok := rulesMapStrToI["rule_name"]; ok && !isIntfNil(w) {
						rules[i].RuleName = w.(string)
					}

					sourceChoiceTypeFound := false

					if v, ok := rulesMapStrToI["all_sources"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true

						if v.(bool) {
							sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_AllSources{}
							sourceChoiceInt.AllSources = &ves_io_schema.Empty{}
							rules[i].SourceChoice = sourceChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["inside_sources"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true

						if v.(bool) {
							sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_InsideSources{}
							sourceChoiceInt.InsideSources = &ves_io_schema.Empty{}
							rules[i].SourceChoice = sourceChoiceInt
						}

					}

					if v, ok := rulesMapStrToI["interface"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_Interface{}
						sourceChoiceInt.Interface = &ves_io_schema_views.ObjectRefType{}
						rules[i].SourceChoice = sourceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								sourceChoiceInt.Interface.Name = v.(string)
							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								sourceChoiceInt.Interface.Namespace = v.(string)
							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								sourceChoiceInt.Interface.Tenant = v.(string)
							}

						}

					}

					if v, ok := rulesMapStrToI["ip_prefix_set"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_IpPrefixSet{}
						sourceChoiceInt.IpPrefixSet = &ves_io_schema_views.ObjectRefType{}
						rules[i].SourceChoice = sourceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["name"]; ok && !isIntfNil(v) {

								sourceChoiceInt.IpPrefixSet.Name = v.(string)
							}

							if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

								sourceChoiceInt.IpPrefixSet.Namespace = v.(string)
							}

							if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

								sourceChoiceInt.IpPrefixSet.Tenant = v.(string)
							}

						}

					}

					if v, ok := rulesMapStrToI["label_selector"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_LabelSelector{}
						sourceChoiceInt.LabelSelector = &ves_io_schema.LabelSelectorType{}
						rules[i].SourceChoice = sourceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								sourceChoiceInt.LabelSelector.Expressions = ls

							}

						}

					}

					if v, ok := rulesMapStrToI["namespace"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_Namespace{}

						rules[i].SourceChoice = sourceChoiceInt

						sourceChoiceInt.Namespace = v.(string)

					}

					if v, ok := rulesMapStrToI["prefix_list"]; ok && !isIntfNil(v) && !sourceChoiceTypeFound {

						sourceChoiceTypeFound = true
						sourceChoiceInt := &ves_io_schema_views_forward_proxy_policy.ForwardProxyAdvancedRuleType_PrefixList{}
						sourceChoiceInt.PrefixList = &ves_io_schema_views.PrefixStringListType{}
						rules[i].SourceChoice = sourceChoiceInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["prefixes"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								sourceChoiceInt.PrefixList.Prefixes = ls

							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra ForwardProxyPolicy obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_forward_proxy_policy.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating ForwardProxyPolicy: %s", err)
	}

	return resourceVolterraForwardProxyPolicyRead(d, meta)
}

func resourceVolterraForwardProxyPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_forward_proxy_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ForwardProxyPolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ForwardProxyPolicy before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra ForwardProxyPolicy obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_forward_proxy_policy.ObjectType, namespace, name)
}
