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
	ves_io_schema_malicious_user_mitigation "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/malicious_user_mitigation"
	ves_io_schema_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/policy"
	ves_io_schema_service_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy_rule"
	ves_io_schema_waf_rule_list "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/waf_rule_list"
)

// resourceVolterraServicePolicyRule is implementation of Volterra's ServicePolicyRule resources
func resourceVolterraServicePolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraServicePolicyRuleCreate,
		Read:   resourceVolterraServicePolicyRuleRead,
		Update: resourceVolterraServicePolicyRuleUpdate,
		Delete: resourceVolterraServicePolicyRuleDelete,

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

			"action": {
				Type:     schema.TypeString,
				Required: true,
			},

			"api_group_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"invert_matcher": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"match": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

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

							Type:     schema.TypeSet,
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

						"presence": {

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

			"any_asn": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"asn_list": {

				Type:     schema.TypeSet,
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

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"asn_sets": {

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

			"body_matcher": {

				Type:     schema.TypeSet,
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

			"any_client": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"client_name": {

				Type:     schema.TypeString,
				Optional: true,
			},

			"client_name_matcher": {

				Type:     schema.TypeSet,
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

			"client_selector": {

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

			"client_role": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"match": {
							Type:     schema.TypeString,
							Optional: true,
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

							Type:     schema.TypeSet,
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

						"presence": {

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

			"domain_matcher": {

				Type:     schema.TypeSet,
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

			"any_dst_asn": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"dst_asn_list": {

				Type:     schema.TypeSet,
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

			"dst_asn_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"asn_sets": {

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

			"any_dst_ip": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"dst_ip_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"invert_matcher": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"prefix_sets": {

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

			"dst_ip_prefix_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"invert_match": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"ip_prefixes": {

							Type: schema.TypeList,

							Required: true,
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

							Type:     schema.TypeSet,
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

						"presence": {

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

			"http_method": {

				Type:     schema.TypeSet,
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

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"invert_matcher": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"prefix_sets": {

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

			"ip_prefix_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"invert_match": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"ip_prefixes": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"l4_dest_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"invert_matcher": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"l4_dests": {

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

			"label_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"keys": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"malicious_user_mitigation": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"rules": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"mitigation_action": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"alert_only": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"block_temporarily": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"captcha_challenge": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"javascript_challenge": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"none": {

													Type:     schema.TypeBool,
													Optional: true,
												},
											},
										},
									},

									"threat_level": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"high": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"low": {

													Type:     schema.TypeBool,
													Optional: true,
												},

												"medium": {

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

			"malicious_user_mitigation_bypass": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"path": {

				Type:     schema.TypeSet,
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

							Type:     schema.TypeSet,
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

						"presence": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"rate_limiter": {

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

			"scheme": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"tls_fingerprint_matcher": {

				Type:     schema.TypeSet,
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

			"url_matcher": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"invert_matcher": {
							Type:     schema.TypeBool,
							Optional: true,
						},

						"url_items": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"domain_regex": {

										Type:     schema.TypeString,
										Optional: true,
									},

									"domain_value": {

										Type:     schema.TypeString,
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

									"path_value": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"virtual_host_matcher": {

				Type:     schema.TypeSet,
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

			"waf_action": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"none": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"waf_inline_rule_control": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"exclude_rule_ids": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"waf_rule_control": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"exclude_rule_ids": {

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

// resourceVolterraServicePolicyRuleCreate creates ServicePolicyRule resource
func resourceVolterraServicePolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_service_policy_rule.CreateSpecType{}
	createReq := &ves_io_schema_service_policy_rule.CreateRequest{
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

	if v, ok := d.GetOk("action"); ok && !isIntfNil(v) {

		createSpec.Action = ves_io_schema_policy.RuleAction(ves_io_schema_policy.RuleAction_value[v.(string)])

	}

	if v, ok := d.GetOk("api_group_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		apiGroupMatcher := &ves_io_schema_policy.StringMatcherType{}
		createSpec.ApiGroupMatcher = apiGroupMatcher
		for _, set := range sl {

			apiGroupMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := apiGroupMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				apiGroupMatcher.InvertMatcher = w.(bool)
			}

			if w, ok := apiGroupMatcherMapStrToI["match"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				apiGroupMatcher.Match = ls
			}

		}

	}

	if v, ok := d.GetOk("arg_matchers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		argMatchers := make([]*ves_io_schema_policy.ArgMatcherType, len(sl))
		createSpec.ArgMatchers = argMatchers
		for i, set := range sl {
			argMatchers[i] = &ves_io_schema_policy.ArgMatcherType{}

			argMatchersMapStrToI := set.(map[string]interface{})

			if w, ok := argMatchersMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				argMatchers[i].InvertMatcher = w.(bool)
			}

			matchTypeFound := false

			if v, ok := argMatchersMapStrToI["check_not_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.ArgMatcherType_CheckNotPresent{}
					matchInt.CheckNotPresent = &ves_io_schema.Empty{}
					argMatchers[i].Match = matchInt
				}

			}

			if v, ok := argMatchersMapStrToI["check_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.ArgMatcherType_CheckPresent{}
					matchInt.CheckPresent = &ves_io_schema.Empty{}
					argMatchers[i].Match = matchInt
				}

			}

			if v, ok := argMatchersMapStrToI["item"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.ArgMatcherType_Item{}
				matchInt.Item = &ves_io_schema_policy.MatcherType{}
				argMatchers[i].Match = matchInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.ExactValues = ls

					}

					if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.RegexValues = ls

					}

					if v, ok := cs["transformers"]; ok && !isIntfNil(v) {

						transformersList := []ves_io_schema_policy.Transformer{}
						for _, j := range v.([]interface{}) {
							transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
						}
						matchInt.Item.Transformers = transformersList

					}

				}

			}

			if v, ok := argMatchersMapStrToI["presence"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.ArgMatcherType_Presence{}

				argMatchers[i].Match = matchInt

				matchInt.Presence =
					v.(bool)

			}

			if w, ok := argMatchersMapStrToI["name"]; ok && !isIntfNil(w) {
				argMatchers[i].Name = w.(string)
			}

		}

	}

	asnChoiceTypeFound := false

	if v, ok := d.GetOk("any_asn"); ok && !asnChoiceTypeFound {

		asnChoiceTypeFound = true

		if v.(bool) {
			asnChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_AnyAsn{}
			asnChoiceInt.AnyAsn = &ves_io_schema.Empty{}
			createSpec.AsnChoice = asnChoiceInt
		}

	}

	if v, ok := d.GetOk("asn_list"); ok && !asnChoiceTypeFound {

		asnChoiceTypeFound = true
		asnChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_AsnList{}
		asnChoiceInt.AsnList = &ves_io_schema_policy.AsnMatchList{}
		createSpec.AsnChoice = asnChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["as_numbers"]; ok && !isIntfNil(v) {

				ls := make([]uint32, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {

					ls[i] = uint32(v.(int))
				}
				asnChoiceInt.AsnList.AsNumbers = ls

			}

		}

	}

	if v, ok := d.GetOk("asn_matcher"); ok && !asnChoiceTypeFound {

		asnChoiceTypeFound = true
		asnChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_AsnMatcher{}
		asnChoiceInt.AsnMatcher = &ves_io_schema_policy.AsnMatcherType{}
		createSpec.AsnChoice = asnChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["asn_sets"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				asnSetsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				asnChoiceInt.AsnMatcher.AsnSets = asnSetsInt
				for i, ps := range sl {

					asMapToStrVal := ps.(map[string]interface{})
					asnSetsInt[i] = &ves_io_schema.ObjectRefType{}

					asnSetsInt[i].Kind = "bgp_asn_set"

					if v, ok := asMapToStrVal["name"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Name = v.(string)
					}

					if v, ok := asMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Namespace = v.(string)
					}

					if v, ok := asMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Tenant = v.(string)
					}

					if v, ok := asMapToStrVal["uid"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("body_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		bodyMatcher := &ves_io_schema_policy.MatcherType{}
		createSpec.BodyMatcher = bodyMatcher
		for _, set := range sl {

			bodyMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := bodyMatcherMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				bodyMatcher.ExactValues = ls
			}

			if w, ok := bodyMatcherMapStrToI["regex_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				bodyMatcher.RegexValues = ls
			}

			if v, ok := bodyMatcherMapStrToI["transformers"]; ok && !isIntfNil(v) {

				transformersList := []ves_io_schema_policy.Transformer{}
				for _, j := range v.([]interface{}) {
					transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
				}
				bodyMatcher.Transformers = transformersList

			}

		}

	}

	clientChoiceTypeFound := false

	if v, ok := d.GetOk("any_client"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true

		if v.(bool) {
			clientChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_AnyClient{}
			clientChoiceInt.AnyClient = &ves_io_schema.Empty{}
			createSpec.ClientChoice = clientChoiceInt
		}

	}

	if v, ok := d.GetOk("client_name"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_ClientName{}

		createSpec.ClientChoice = clientChoiceInt

		clientChoiceInt.ClientName = v.(string)

	}

	if v, ok := d.GetOk("client_name_matcher"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_ClientNameMatcher{}
		clientChoiceInt.ClientNameMatcher = &ves_io_schema_policy.MatcherTypeBasic{}
		createSpec.ClientChoice = clientChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientNameMatcher.ExactValues = ls

			}

			if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientNameMatcher.RegexValues = ls

			}

		}

	}

	if v, ok := d.GetOk("client_selector"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_ClientSelector{}
		clientChoiceInt.ClientSelector = &ves_io_schema.LabelSelectorType{}
		createSpec.ClientChoice = clientChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientSelector.Expressions = ls

			}

		}

	}

	if v, ok := d.GetOk("client_role"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		clientRole := &ves_io_schema_policy.RoleMatcherType{}
		createSpec.ClientRole = clientRole
		for _, set := range sl {

			clientRoleMapStrToI := set.(map[string]interface{})

			if w, ok := clientRoleMapStrToI["match"]; ok && !isIntfNil(w) {
				clientRole.Match = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("cookie_matchers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		cookieMatchers := make([]*ves_io_schema_policy.CookieMatcherType, len(sl))
		createSpec.CookieMatchers = cookieMatchers
		for i, set := range sl {
			cookieMatchers[i] = &ves_io_schema_policy.CookieMatcherType{}

			cookieMatchersMapStrToI := set.(map[string]interface{})

			if w, ok := cookieMatchersMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				cookieMatchers[i].InvertMatcher = w.(bool)
			}

			matchTypeFound := false

			if v, ok := cookieMatchersMapStrToI["check_not_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.CookieMatcherType_CheckNotPresent{}
					matchInt.CheckNotPresent = &ves_io_schema.Empty{}
					cookieMatchers[i].Match = matchInt
				}

			}

			if v, ok := cookieMatchersMapStrToI["check_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.CookieMatcherType_CheckPresent{}
					matchInt.CheckPresent = &ves_io_schema.Empty{}
					cookieMatchers[i].Match = matchInt
				}

			}

			if v, ok := cookieMatchersMapStrToI["item"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.CookieMatcherType_Item{}
				matchInt.Item = &ves_io_schema_policy.MatcherType{}
				cookieMatchers[i].Match = matchInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.ExactValues = ls

					}

					if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.RegexValues = ls

					}

					if v, ok := cs["transformers"]; ok && !isIntfNil(v) {

						transformersList := []ves_io_schema_policy.Transformer{}
						for _, j := range v.([]interface{}) {
							transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
						}
						matchInt.Item.Transformers = transformersList

					}

				}

			}

			if v, ok := cookieMatchersMapStrToI["presence"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.CookieMatcherType_Presence{}

				cookieMatchers[i].Match = matchInt

				matchInt.Presence =
					v.(bool)

			}

			if w, ok := cookieMatchersMapStrToI["name"]; ok && !isIntfNil(w) {
				cookieMatchers[i].Name = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("domain_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		domainMatcher := &ves_io_schema_policy.MatcherTypeBasic{}
		createSpec.DomainMatcher = domainMatcher
		for _, set := range sl {

			domainMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := domainMatcherMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				domainMatcher.ExactValues = ls
			}

			if w, ok := domainMatcherMapStrToI["regex_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				domainMatcher.RegexValues = ls
			}

		}

	}

	dstAsnChoiceTypeFound := false

	if v, ok := d.GetOk("any_dst_asn"); ok && !dstAsnChoiceTypeFound {

		dstAsnChoiceTypeFound = true

		if v.(bool) {
			dstAsnChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_AnyDstAsn{}
			dstAsnChoiceInt.AnyDstAsn = &ves_io_schema.Empty{}
			createSpec.DstAsnChoice = dstAsnChoiceInt
		}

	}

	if v, ok := d.GetOk("dst_asn_list"); ok && !dstAsnChoiceTypeFound {

		dstAsnChoiceTypeFound = true
		dstAsnChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_DstAsnList{}
		dstAsnChoiceInt.DstAsnList = &ves_io_schema_policy.AsnMatchList{}
		createSpec.DstAsnChoice = dstAsnChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["as_numbers"]; ok && !isIntfNil(v) {

				ls := make([]uint32, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {

					ls[i] = uint32(v.(int))
				}
				dstAsnChoiceInt.DstAsnList.AsNumbers = ls

			}

		}

	}

	if v, ok := d.GetOk("dst_asn_matcher"); ok && !dstAsnChoiceTypeFound {

		dstAsnChoiceTypeFound = true
		dstAsnChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_DstAsnMatcher{}
		dstAsnChoiceInt.DstAsnMatcher = &ves_io_schema_policy.AsnMatcherType{}
		createSpec.DstAsnChoice = dstAsnChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["asn_sets"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				asnSetsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				dstAsnChoiceInt.DstAsnMatcher.AsnSets = asnSetsInt
				for i, ps := range sl {

					asMapToStrVal := ps.(map[string]interface{})
					asnSetsInt[i] = &ves_io_schema.ObjectRefType{}

					asnSetsInt[i].Kind = "bgp_asn_set"

					if v, ok := asMapToStrVal["name"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Name = v.(string)
					}

					if v, ok := asMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Namespace = v.(string)
					}

					if v, ok := asMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Tenant = v.(string)
					}

					if v, ok := asMapToStrVal["uid"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	dstIpChoiceTypeFound := false

	if v, ok := d.GetOk("any_dst_ip"); ok && !dstIpChoiceTypeFound {

		dstIpChoiceTypeFound = true

		if v.(bool) {
			dstIpChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_AnyDstIp{}
			dstIpChoiceInt.AnyDstIp = &ves_io_schema.Empty{}
			createSpec.DstIpChoice = dstIpChoiceInt
		}

	}

	if v, ok := d.GetOk("dst_ip_matcher"); ok && !dstIpChoiceTypeFound {

		dstIpChoiceTypeFound = true
		dstIpChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_DstIpMatcher{}
		dstIpChoiceInt.DstIpMatcher = &ves_io_schema_policy.IpMatcherType{}
		createSpec.DstIpChoice = dstIpChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["invert_matcher"]; ok && !isIntfNil(v) {

				dstIpChoiceInt.DstIpMatcher.InvertMatcher = v.(bool)
			}

			if v, ok := cs["prefix_sets"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				prefixSetsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				dstIpChoiceInt.DstIpMatcher.PrefixSets = prefixSetsInt
				for i, ps := range sl {

					psMapToStrVal := ps.(map[string]interface{})
					prefixSetsInt[i] = &ves_io_schema.ObjectRefType{}

					prefixSetsInt[i].Kind = "ip_prefix_set"

					if v, ok := psMapToStrVal["name"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Name = v.(string)
					}

					if v, ok := psMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Namespace = v.(string)
					}

					if v, ok := psMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Tenant = v.(string)
					}

					if v, ok := psMapToStrVal["uid"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("dst_ip_prefix_list"); ok && !dstIpChoiceTypeFound {

		dstIpChoiceTypeFound = true
		dstIpChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_DstIpPrefixList{}
		dstIpChoiceInt.DstIpPrefixList = &ves_io_schema_policy.PrefixMatchList{}
		createSpec.DstIpChoice = dstIpChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["invert_match"]; ok && !isIntfNil(v) {

				dstIpChoiceInt.DstIpPrefixList.InvertMatch = v.(bool)
			}

			if v, ok := cs["ip_prefixes"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				dstIpChoiceInt.DstIpPrefixList.IpPrefixes = ls

			}

		}

	}

	if v, ok := d.GetOk("expiration_timestamp"); ok && !isIntfNil(v) {

		ts, err := parseTime(v.(string))
		if err != nil {
			return fmt.Errorf("error creating ServicePolicyRule, timestamp format is wrong: %s", err)
		}
		createSpec.ExpirationTimestamp = ts

	}

	if v, ok := d.GetOk("headers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		headers := make([]*ves_io_schema_policy.HeaderMatcherType, len(sl))
		createSpec.Headers = headers
		for i, set := range sl {
			headers[i] = &ves_io_schema_policy.HeaderMatcherType{}

			headersMapStrToI := set.(map[string]interface{})

			if w, ok := headersMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				headers[i].InvertMatcher = w.(bool)
			}

			matchTypeFound := false

			if v, ok := headersMapStrToI["check_not_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.HeaderMatcherType_CheckNotPresent{}
					matchInt.CheckNotPresent = &ves_io_schema.Empty{}
					headers[i].Match = matchInt
				}

			}

			if v, ok := headersMapStrToI["check_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.HeaderMatcherType_CheckPresent{}
					matchInt.CheckPresent = &ves_io_schema.Empty{}
					headers[i].Match = matchInt
				}

			}

			if v, ok := headersMapStrToI["item"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.HeaderMatcherType_Item{}
				matchInt.Item = &ves_io_schema_policy.MatcherType{}
				headers[i].Match = matchInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.ExactValues = ls

					}

					if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.RegexValues = ls

					}

					if v, ok := cs["transformers"]; ok && !isIntfNil(v) {

						transformersList := []ves_io_schema_policy.Transformer{}
						for _, j := range v.([]interface{}) {
							transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
						}
						matchInt.Item.Transformers = transformersList

					}

				}

			}

			if v, ok := headersMapStrToI["presence"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.HeaderMatcherType_Presence{}

				headers[i].Match = matchInt

				matchInt.Presence =
					v.(bool)

			}

			if w, ok := headersMapStrToI["name"]; ok && !isIntfNil(w) {
				headers[i].Name = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("http_method"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		httpMethod := &ves_io_schema_policy.HttpMethodMatcherType{}
		createSpec.HttpMethod = httpMethod
		for _, set := range sl {

			httpMethodMapStrToI := set.(map[string]interface{})

			if w, ok := httpMethodMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				httpMethod.InvertMatcher = w.(bool)
			}

			if v, ok := httpMethodMapStrToI["methods"]; ok && !isIntfNil(v) {

				methodsList := []ves_io_schema.HttpMethod{}
				for _, j := range v.([]interface{}) {
					methodsList = append(methodsList, ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[j.(string)]))
				}
				httpMethod.Methods = methodsList

			}

		}

	}

	ipChoiceTypeFound := false

	if v, ok := d.GetOk("any_ip"); ok && !ipChoiceTypeFound {

		ipChoiceTypeFound = true

		if v.(bool) {
			ipChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_AnyIp{}
			ipChoiceInt.AnyIp = &ves_io_schema.Empty{}
			createSpec.IpChoice = ipChoiceInt
		}

	}

	if v, ok := d.GetOk("ip_matcher"); ok && !ipChoiceTypeFound {

		ipChoiceTypeFound = true
		ipChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_IpMatcher{}
		ipChoiceInt.IpMatcher = &ves_io_schema_policy.IpMatcherType{}
		createSpec.IpChoice = ipChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["invert_matcher"]; ok && !isIntfNil(v) {

				ipChoiceInt.IpMatcher.InvertMatcher = v.(bool)
			}

			if v, ok := cs["prefix_sets"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				prefixSetsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				ipChoiceInt.IpMatcher.PrefixSets = prefixSetsInt
				for i, ps := range sl {

					psMapToStrVal := ps.(map[string]interface{})
					prefixSetsInt[i] = &ves_io_schema.ObjectRefType{}

					prefixSetsInt[i].Kind = "ip_prefix_set"

					if v, ok := psMapToStrVal["name"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Name = v.(string)
					}

					if v, ok := psMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Namespace = v.(string)
					}

					if v, ok := psMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Tenant = v.(string)
					}

					if v, ok := psMapToStrVal["uid"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("ip_prefix_list"); ok && !ipChoiceTypeFound {

		ipChoiceTypeFound = true
		ipChoiceInt := &ves_io_schema_service_policy_rule.CreateSpecType_IpPrefixList{}
		ipChoiceInt.IpPrefixList = &ves_io_schema_policy.PrefixMatchList{}
		createSpec.IpChoice = ipChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["invert_match"]; ok && !isIntfNil(v) {

				ipChoiceInt.IpPrefixList.InvertMatch = v.(bool)
			}

			if v, ok := cs["ip_prefixes"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				ipChoiceInt.IpPrefixList.IpPrefixes = ls

			}

		}

	}

	if v, ok := d.GetOk("l4_dest_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		l4DestMatcher := &ves_io_schema_policy.L4DestMatcherType{}
		createSpec.L4DestMatcher = l4DestMatcher
		for _, set := range sl {

			l4DestMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := l4DestMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				l4DestMatcher.InvertMatcher = w.(bool)
			}

			if v, ok := l4DestMatcherMapStrToI["l4_dests"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				l4Dests := make([]*ves_io_schema.L4DestType, len(sl))
				l4DestMatcher.L4Dests = l4Dests
				for i, set := range sl {
					l4Dests[i] = &ves_io_schema.L4DestType{}

					l4DestsMapStrToI := set.(map[string]interface{})

					if w, ok := l4DestsMapStrToI["ports"]; ok && !isIntfNil(w) {
						l4Dests[i].Ports = w.(string)
					}

					if w, ok := l4DestsMapStrToI["prefixes"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						l4Dests[i].Prefixes = ls
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("label_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		labelMatcher := &ves_io_schema.LabelMatcherType{}
		createSpec.LabelMatcher = labelMatcher
		for _, set := range sl {

			labelMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := labelMatcherMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				labelMatcher.Keys = ls
			}

		}

	}

	if v, ok := d.GetOk("malicious_user_mitigation"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		maliciousUserMitigation := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationType{}
		createSpec.MaliciousUserMitigation = maliciousUserMitigation
		for _, set := range sl {

			maliciousUserMitigationMapStrToI := set.(map[string]interface{})

			if v, ok := maliciousUserMitigationMapStrToI["rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				rules := make([]*ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationRule, len(sl))
				maliciousUserMitigation.Rules = rules
				for i, set := range sl {
					rules[i] = &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationRule{}

					rulesMapStrToI := set.(map[string]interface{})

					if v, ok := rulesMapStrToI["mitigation_action"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						mitigationAction := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction{}
						rules[i].MitigationAction = mitigationAction
						for _, set := range sl {

							mitigationActionMapStrToI := set.(map[string]interface{})

							mitigationActionTypeFound := false

							if v, ok := mitigationActionMapStrToI["alert_only"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_AlertOnly{}
									mitigationActionInt.AlertOnly = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["block_temporarily"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_BlockTemporarily{}
									mitigationActionInt.BlockTemporarily = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["captcha_challenge"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_CaptchaChallenge{}
									mitigationActionInt.CaptchaChallenge = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["javascript_challenge"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_JavascriptChallenge{}
									mitigationActionInt.JavascriptChallenge = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["none"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_None{}
									mitigationActionInt.None = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

						}

					}

					if v, ok := rulesMapStrToI["threat_level"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						threatLevel := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel{}
						rules[i].ThreatLevel = threatLevel
						for _, set := range sl {

							threatLevelMapStrToI := set.(map[string]interface{})

							threatLevelTypeFound := false

							if v, ok := threatLevelMapStrToI["high"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_High{}
									threatLevelInt.High = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

							if v, ok := threatLevelMapStrToI["low"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_Low{}
									threatLevelInt.Low = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

							if v, ok := threatLevelMapStrToI["medium"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_Medium{}
									threatLevelInt.Medium = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("malicious_user_mitigation_bypass"); ok && !isIntfNil(v) {

		if v.(bool) {
			createSpec.MaliciousUserMitigationBypass = &ves_io_schema.Empty{}
		}

	}

	if v, ok := d.GetOk("path"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		path := &ves_io_schema_policy.PathMatcherType{}
		createSpec.Path = path
		for _, set := range sl {

			pathMapStrToI := set.(map[string]interface{})

			if w, ok := pathMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				path.ExactValues = ls
			}

			if w, ok := pathMapStrToI["prefix_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				path.PrefixValues = ls
			}

			if w, ok := pathMapStrToI["regex_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				path.RegexValues = ls
			}

			if v, ok := pathMapStrToI["transformers"]; ok && !isIntfNil(v) {

				transformersList := []ves_io_schema_policy.Transformer{}
				for _, j := range v.([]interface{}) {
					transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
				}
				path.Transformers = transformersList

			}

		}

	}

	if v, ok := d.GetOk("port_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		portMatcher := &ves_io_schema_policy.PortMatcherType{}
		createSpec.PortMatcher = portMatcher
		for _, set := range sl {

			portMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := portMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				portMatcher.InvertMatcher = w.(bool)
			}

			if w, ok := portMatcherMapStrToI["ports"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				portMatcher.Ports = ls
			}

		}

	}

	if v, ok := d.GetOk("query_params"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		queryParams := make([]*ves_io_schema_policy.QueryParameterMatcherType, len(sl))
		createSpec.QueryParams = queryParams
		for i, set := range sl {
			queryParams[i] = &ves_io_schema_policy.QueryParameterMatcherType{}

			queryParamsMapStrToI := set.(map[string]interface{})

			if w, ok := queryParamsMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				queryParams[i].InvertMatcher = w.(bool)
			}

			if w, ok := queryParamsMapStrToI["key"]; ok && !isIntfNil(w) {
				queryParams[i].Key = w.(string)
			}

			matchTypeFound := false

			if v, ok := queryParamsMapStrToI["check_not_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.QueryParameterMatcherType_CheckNotPresent{}
					matchInt.CheckNotPresent = &ves_io_schema.Empty{}
					queryParams[i].Match = matchInt
				}

			}

			if v, ok := queryParamsMapStrToI["check_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.QueryParameterMatcherType_CheckPresent{}
					matchInt.CheckPresent = &ves_io_schema.Empty{}
					queryParams[i].Match = matchInt
				}

			}

			if v, ok := queryParamsMapStrToI["item"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.QueryParameterMatcherType_Item{}
				matchInt.Item = &ves_io_schema_policy.MatcherType{}
				queryParams[i].Match = matchInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.ExactValues = ls

					}

					if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.RegexValues = ls

					}

					if v, ok := cs["transformers"]; ok && !isIntfNil(v) {

						transformersList := []ves_io_schema_policy.Transformer{}
						for _, j := range v.([]interface{}) {
							transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
						}
						matchInt.Item.Transformers = transformersList

					}

				}

			}

			if v, ok := queryParamsMapStrToI["presence"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.QueryParameterMatcherType_Presence{}

				queryParams[i].Match = matchInt

				matchInt.Presence =
					v.(bool)

			}

		}

	}

	if v, ok := d.GetOk("rate_limiter"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rateLimiterInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.RateLimiter = rateLimiterInt
		for i, ps := range sl {

			rlMapToStrVal := ps.(map[string]interface{})
			rateLimiterInt[i] = &ves_io_schema.ObjectRefType{}

			rateLimiterInt[i].Kind = "rate_limiter"

			if v, ok := rlMapToStrVal["name"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Name = v.(string)
			}

			if v, ok := rlMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Namespace = v.(string)
			}

			if v, ok := rlMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Tenant = v.(string)
			}

			if v, ok := rlMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("scheme"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		createSpec.Scheme = ls

	}

	if v, ok := d.GetOk("tls_fingerprint_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tlsFingerprintMatcher := &ves_io_schema_policy.TlsFingerprintMatcherType{}
		createSpec.TlsFingerprintMatcher = tlsFingerprintMatcher
		for _, set := range sl {

			tlsFingerprintMatcherMapStrToI := set.(map[string]interface{})

			if v, ok := tlsFingerprintMatcherMapStrToI["classes"]; ok && !isIntfNil(v) {

				classesList := []ves_io_schema_policy.KnownTlsFingerprintClass{}
				for _, j := range v.([]interface{}) {
					classesList = append(classesList, ves_io_schema_policy.KnownTlsFingerprintClass(ves_io_schema_policy.KnownTlsFingerprintClass_value[j.(string)]))
				}
				tlsFingerprintMatcher.Classes = classesList

			}

			if w, ok := tlsFingerprintMatcherMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				tlsFingerprintMatcher.ExactValues = ls
			}

			if w, ok := tlsFingerprintMatcherMapStrToI["excluded_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				tlsFingerprintMatcher.ExcludedValues = ls
			}

		}

	}

	if v, ok := d.GetOk("url_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		urlMatcher := &ves_io_schema_policy.URLMatcherType{}
		createSpec.UrlMatcher = urlMatcher
		for _, set := range sl {

			urlMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := urlMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				urlMatcher.InvertMatcher = w.(bool)
			}

			if v, ok := urlMatcherMapStrToI["url_items"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				urlItems := make([]*ves_io_schema_policy.URLItem, len(sl))
				urlMatcher.UrlItems = urlItems
				for i, set := range sl {
					urlItems[i] = &ves_io_schema_policy.URLItem{}

					urlItemsMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := urlItemsMapStrToI["domain_regex"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_policy.URLItem_DomainRegex{}

						urlItems[i].DomainChoice = domainChoiceInt

						domainChoiceInt.DomainRegex = v.(string)

					}

					if v, ok := urlItemsMapStrToI["domain_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_policy.URLItem_DomainValue{}

						urlItems[i].DomainChoice = domainChoiceInt

						domainChoiceInt.DomainValue = v.(string)

					}

					pathChoiceTypeFound := false

					if v, ok := urlItemsMapStrToI["path_prefix"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_policy.URLItem_PathPrefix{}

						urlItems[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathPrefix = v.(string)

					}

					if v, ok := urlItemsMapStrToI["path_regex"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_policy.URLItem_PathRegex{}

						urlItems[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathRegex = v.(string)

					}

					if v, ok := urlItemsMapStrToI["path_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_policy.URLItem_PathValue{}

						urlItems[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathValue = v.(string)

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("virtual_host_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		virtualHostMatcher := &ves_io_schema_policy.MatcherTypeBasic{}
		createSpec.VirtualHostMatcher = virtualHostMatcher
		for _, set := range sl {

			virtualHostMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := virtualHostMatcherMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				virtualHostMatcher.ExactValues = ls
			}

			if w, ok := virtualHostMatcherMapStrToI["regex_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				virtualHostMatcher.RegexValues = ls
			}

		}

	}

	if v, ok := d.GetOk("waf_action"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		wafAction := &ves_io_schema_policy.WafAction{}
		createSpec.WafAction = wafAction
		for _, set := range sl {

			wafActionMapStrToI := set.(map[string]interface{})

			actionTypeTypeFound := false

			if v, ok := wafActionMapStrToI["none"]; ok && !isIntfNil(v) && !actionTypeTypeFound {

				actionTypeTypeFound = true

				if v.(bool) {
					actionTypeInt := &ves_io_schema_policy.WafAction_None{}
					actionTypeInt.None = &ves_io_schema.Empty{}
					wafAction.ActionType = actionTypeInt
				}

			}

			if v, ok := wafActionMapStrToI["waf_inline_rule_control"]; ok && !isIntfNil(v) && !actionTypeTypeFound {

				actionTypeTypeFound = true
				actionTypeInt := &ves_io_schema_policy.WafAction_WafInlineRuleControl{}
				actionTypeInt.WafInlineRuleControl = &ves_io_schema_policy.WafInlineRuleControl{}
				wafAction.ActionType = actionTypeInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exclude_rule_ids"]; ok && !isIntfNil(v) {

						exclude_rule_idsList := []ves_io_schema_waf_rule_list.WafRuleID{}
						for _, j := range v.([]interface{}) {
							exclude_rule_idsList = append(exclude_rule_idsList, ves_io_schema_waf_rule_list.WafRuleID(ves_io_schema_waf_rule_list.WafRuleID_value[j.(string)]))
						}
						actionTypeInt.WafInlineRuleControl.ExcludeRuleIds = exclude_rule_idsList

					}

				}

			}

			if v, ok := wafActionMapStrToI["waf_rule_control"]; ok && !isIntfNil(v) && !actionTypeTypeFound {

				actionTypeTypeFound = true
				actionTypeInt := &ves_io_schema_policy.WafAction_WafRuleControl{}
				actionTypeInt.WafRuleControl = &ves_io_schema_policy.WafRuleControl{}
				wafAction.ActionType = actionTypeInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exclude_rule_ids"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						excludeRuleIdsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						actionTypeInt.WafRuleControl.ExcludeRuleIds = excludeRuleIdsInt
						for i, ps := range sl {

							eriMapToStrVal := ps.(map[string]interface{})
							excludeRuleIdsInt[i] = &ves_io_schema.ObjectRefType{}

							excludeRuleIdsInt[i].Kind = "waf_rule_list"

							if v, ok := eriMapToStrVal["name"]; ok && !isIntfNil(v) {
								excludeRuleIdsInt[i].Name = v.(string)
							}

							if v, ok := eriMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								excludeRuleIdsInt[i].Namespace = v.(string)
							}

							if v, ok := eriMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								excludeRuleIdsInt[i].Tenant = v.(string)
							}

							if v, ok := eriMapToStrVal["uid"]; ok && !isIntfNil(v) {
								excludeRuleIdsInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := wafActionMapStrToI["waf_skip_processing"]; ok && !isIntfNil(v) && !actionTypeTypeFound {

				actionTypeTypeFound = true

				if v.(bool) {
					actionTypeInt := &ves_io_schema_policy.WafAction_WafSkipProcessing{}
					actionTypeInt.WafSkipProcessing = &ves_io_schema.Empty{}
					wafAction.ActionType = actionTypeInt
				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra ServicePolicyRule object with struct: %+v", createReq)

	createServicePolicyRuleResp, err := client.CreateObject(context.Background(), ves_io_schema_service_policy_rule.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating ServicePolicyRule: %s", err)
	}
	d.SetId(createServicePolicyRuleResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraServicePolicyRuleRead(d, meta)
}

func resourceVolterraServicePolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_service_policy_rule.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ServicePolicyRule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ServicePolicyRule %q: %s", d.Id(), err)
	}
	return setServicePolicyRuleFields(client, d, resp)
}

func setServicePolicyRuleFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraServicePolicyRuleUpdate updates ServicePolicyRule resource
func resourceVolterraServicePolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_service_policy_rule.ReplaceSpecType{}
	updateReq := &ves_io_schema_service_policy_rule.ReplaceRequest{
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

	if v, ok := d.GetOk("action"); ok && !isIntfNil(v) {

		updateSpec.Action = ves_io_schema_policy.RuleAction(ves_io_schema_policy.RuleAction_value[v.(string)])

	}

	if v, ok := d.GetOk("api_group_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		apiGroupMatcher := &ves_io_schema_policy.StringMatcherType{}
		updateSpec.ApiGroupMatcher = apiGroupMatcher
		for _, set := range sl {

			apiGroupMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := apiGroupMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				apiGroupMatcher.InvertMatcher = w.(bool)
			}

			if w, ok := apiGroupMatcherMapStrToI["match"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				apiGroupMatcher.Match = ls
			}

		}

	}

	if v, ok := d.GetOk("arg_matchers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		argMatchers := make([]*ves_io_schema_policy.ArgMatcherType, len(sl))
		updateSpec.ArgMatchers = argMatchers
		for i, set := range sl {
			argMatchers[i] = &ves_io_schema_policy.ArgMatcherType{}

			argMatchersMapStrToI := set.(map[string]interface{})

			if w, ok := argMatchersMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				argMatchers[i].InvertMatcher = w.(bool)
			}

			matchTypeFound := false

			if v, ok := argMatchersMapStrToI["check_not_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.ArgMatcherType_CheckNotPresent{}
					matchInt.CheckNotPresent = &ves_io_schema.Empty{}
					argMatchers[i].Match = matchInt
				}

			}

			if v, ok := argMatchersMapStrToI["check_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.ArgMatcherType_CheckPresent{}
					matchInt.CheckPresent = &ves_io_schema.Empty{}
					argMatchers[i].Match = matchInt
				}

			}

			if v, ok := argMatchersMapStrToI["item"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.ArgMatcherType_Item{}
				matchInt.Item = &ves_io_schema_policy.MatcherType{}
				argMatchers[i].Match = matchInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.ExactValues = ls

					}

					if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.RegexValues = ls

					}

					if v, ok := cs["transformers"]; ok && !isIntfNil(v) {

						transformersList := []ves_io_schema_policy.Transformer{}
						for _, j := range v.([]interface{}) {
							transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
						}
						matchInt.Item.Transformers = transformersList

					}

				}

			}

			if v, ok := argMatchersMapStrToI["presence"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.ArgMatcherType_Presence{}

				argMatchers[i].Match = matchInt

				matchInt.Presence =
					v.(bool)

			}

			if w, ok := argMatchersMapStrToI["name"]; ok && !isIntfNil(w) {
				argMatchers[i].Name = w.(string)
			}

		}

	}

	asnChoiceTypeFound := false

	if v, ok := d.GetOk("any_asn"); ok && !asnChoiceTypeFound {

		asnChoiceTypeFound = true

		if v.(bool) {
			asnChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_AnyAsn{}
			asnChoiceInt.AnyAsn = &ves_io_schema.Empty{}
			updateSpec.AsnChoice = asnChoiceInt
		}

	}

	if v, ok := d.GetOk("asn_list"); ok && !asnChoiceTypeFound {

		asnChoiceTypeFound = true
		asnChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_AsnList{}
		asnChoiceInt.AsnList = &ves_io_schema_policy.AsnMatchList{}
		updateSpec.AsnChoice = asnChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["as_numbers"]; ok && !isIntfNil(v) {

				ls := make([]uint32, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {

					ls[i] = uint32(v.(int))
				}
				asnChoiceInt.AsnList.AsNumbers = ls

			}

		}

	}

	if v, ok := d.GetOk("asn_matcher"); ok && !asnChoiceTypeFound {

		asnChoiceTypeFound = true
		asnChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_AsnMatcher{}
		asnChoiceInt.AsnMatcher = &ves_io_schema_policy.AsnMatcherType{}
		updateSpec.AsnChoice = asnChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["asn_sets"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				asnSetsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				asnChoiceInt.AsnMatcher.AsnSets = asnSetsInt
				for i, ps := range sl {

					asMapToStrVal := ps.(map[string]interface{})
					asnSetsInt[i] = &ves_io_schema.ObjectRefType{}

					asnSetsInt[i].Kind = "bgp_asn_set"

					if v, ok := asMapToStrVal["name"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Name = v.(string)
					}

					if v, ok := asMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Namespace = v.(string)
					}

					if v, ok := asMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Tenant = v.(string)
					}

					if v, ok := asMapToStrVal["uid"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("body_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		bodyMatcher := &ves_io_schema_policy.MatcherType{}
		updateSpec.BodyMatcher = bodyMatcher
		for _, set := range sl {

			bodyMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := bodyMatcherMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				bodyMatcher.ExactValues = ls
			}

			if w, ok := bodyMatcherMapStrToI["regex_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				bodyMatcher.RegexValues = ls
			}

			if v, ok := bodyMatcherMapStrToI["transformers"]; ok && !isIntfNil(v) {

				transformersList := []ves_io_schema_policy.Transformer{}
				for _, j := range v.([]interface{}) {
					transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
				}
				bodyMatcher.Transformers = transformersList

			}

		}

	}

	clientChoiceTypeFound := false

	if v, ok := d.GetOk("any_client"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true

		if v.(bool) {
			clientChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_AnyClient{}
			clientChoiceInt.AnyClient = &ves_io_schema.Empty{}
			updateSpec.ClientChoice = clientChoiceInt
		}

	}

	if v, ok := d.GetOk("client_name"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_ClientName{}

		updateSpec.ClientChoice = clientChoiceInt

		clientChoiceInt.ClientName = v.(string)

	}

	if v, ok := d.GetOk("client_name_matcher"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_ClientNameMatcher{}
		clientChoiceInt.ClientNameMatcher = &ves_io_schema_policy.MatcherTypeBasic{}
		updateSpec.ClientChoice = clientChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientNameMatcher.ExactValues = ls

			}

			if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientNameMatcher.RegexValues = ls

			}

		}

	}

	if v, ok := d.GetOk("client_selector"); ok && !clientChoiceTypeFound {

		clientChoiceTypeFound = true
		clientChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_ClientSelector{}
		clientChoiceInt.ClientSelector = &ves_io_schema.LabelSelectorType{}
		updateSpec.ClientChoice = clientChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				clientChoiceInt.ClientSelector.Expressions = ls

			}

		}

	}

	if v, ok := d.GetOk("client_role"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		clientRole := &ves_io_schema_policy.RoleMatcherType{}
		updateSpec.ClientRole = clientRole
		for _, set := range sl {

			clientRoleMapStrToI := set.(map[string]interface{})

			if w, ok := clientRoleMapStrToI["match"]; ok && !isIntfNil(w) {
				clientRole.Match = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("cookie_matchers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		cookieMatchers := make([]*ves_io_schema_policy.CookieMatcherType, len(sl))
		updateSpec.CookieMatchers = cookieMatchers
		for i, set := range sl {
			cookieMatchers[i] = &ves_io_schema_policy.CookieMatcherType{}

			cookieMatchersMapStrToI := set.(map[string]interface{})

			if w, ok := cookieMatchersMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				cookieMatchers[i].InvertMatcher = w.(bool)
			}

			matchTypeFound := false

			if v, ok := cookieMatchersMapStrToI["check_not_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.CookieMatcherType_CheckNotPresent{}
					matchInt.CheckNotPresent = &ves_io_schema.Empty{}
					cookieMatchers[i].Match = matchInt
				}

			}

			if v, ok := cookieMatchersMapStrToI["check_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.CookieMatcherType_CheckPresent{}
					matchInt.CheckPresent = &ves_io_schema.Empty{}
					cookieMatchers[i].Match = matchInt
				}

			}

			if v, ok := cookieMatchersMapStrToI["item"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.CookieMatcherType_Item{}
				matchInt.Item = &ves_io_schema_policy.MatcherType{}
				cookieMatchers[i].Match = matchInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.ExactValues = ls

					}

					if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.RegexValues = ls

					}

					if v, ok := cs["transformers"]; ok && !isIntfNil(v) {

						transformersList := []ves_io_schema_policy.Transformer{}
						for _, j := range v.([]interface{}) {
							transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
						}
						matchInt.Item.Transformers = transformersList

					}

				}

			}

			if v, ok := cookieMatchersMapStrToI["presence"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.CookieMatcherType_Presence{}

				cookieMatchers[i].Match = matchInt

				matchInt.Presence =
					v.(bool)

			}

			if w, ok := cookieMatchersMapStrToI["name"]; ok && !isIntfNil(w) {
				cookieMatchers[i].Name = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("domain_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		domainMatcher := &ves_io_schema_policy.MatcherTypeBasic{}
		updateSpec.DomainMatcher = domainMatcher
		for _, set := range sl {

			domainMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := domainMatcherMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				domainMatcher.ExactValues = ls
			}

			if w, ok := domainMatcherMapStrToI["regex_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				domainMatcher.RegexValues = ls
			}

		}

	}

	dstAsnChoiceTypeFound := false

	if v, ok := d.GetOk("any_dst_asn"); ok && !dstAsnChoiceTypeFound {

		dstAsnChoiceTypeFound = true

		if v.(bool) {
			dstAsnChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_AnyDstAsn{}
			dstAsnChoiceInt.AnyDstAsn = &ves_io_schema.Empty{}
			updateSpec.DstAsnChoice = dstAsnChoiceInt
		}

	}

	if v, ok := d.GetOk("dst_asn_list"); ok && !dstAsnChoiceTypeFound {

		dstAsnChoiceTypeFound = true
		dstAsnChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_DstAsnList{}
		dstAsnChoiceInt.DstAsnList = &ves_io_schema_policy.AsnMatchList{}
		updateSpec.DstAsnChoice = dstAsnChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["as_numbers"]; ok && !isIntfNil(v) {

				ls := make([]uint32, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {

					ls[i] = uint32(v.(int))
				}
				dstAsnChoiceInt.DstAsnList.AsNumbers = ls

			}

		}

	}

	if v, ok := d.GetOk("dst_asn_matcher"); ok && !dstAsnChoiceTypeFound {

		dstAsnChoiceTypeFound = true
		dstAsnChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_DstAsnMatcher{}
		dstAsnChoiceInt.DstAsnMatcher = &ves_io_schema_policy.AsnMatcherType{}
		updateSpec.DstAsnChoice = dstAsnChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["asn_sets"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				asnSetsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				dstAsnChoiceInt.DstAsnMatcher.AsnSets = asnSetsInt
				for i, ps := range sl {

					asMapToStrVal := ps.(map[string]interface{})
					asnSetsInt[i] = &ves_io_schema.ObjectRefType{}

					asnSetsInt[i].Kind = "bgp_asn_set"

					if v, ok := asMapToStrVal["name"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Name = v.(string)
					}

					if v, ok := asMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Namespace = v.(string)
					}

					if v, ok := asMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Tenant = v.(string)
					}

					if v, ok := asMapToStrVal["uid"]; ok && !isIntfNil(v) {
						asnSetsInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	dstIpChoiceTypeFound := false

	if v, ok := d.GetOk("any_dst_ip"); ok && !dstIpChoiceTypeFound {

		dstIpChoiceTypeFound = true

		if v.(bool) {
			dstIpChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_AnyDstIp{}
			dstIpChoiceInt.AnyDstIp = &ves_io_schema.Empty{}
			updateSpec.DstIpChoice = dstIpChoiceInt
		}

	}

	if v, ok := d.GetOk("dst_ip_matcher"); ok && !dstIpChoiceTypeFound {

		dstIpChoiceTypeFound = true
		dstIpChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_DstIpMatcher{}
		dstIpChoiceInt.DstIpMatcher = &ves_io_schema_policy.IpMatcherType{}
		updateSpec.DstIpChoice = dstIpChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["invert_matcher"]; ok && !isIntfNil(v) {

				dstIpChoiceInt.DstIpMatcher.InvertMatcher = v.(bool)
			}

			if v, ok := cs["prefix_sets"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				prefixSetsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				dstIpChoiceInt.DstIpMatcher.PrefixSets = prefixSetsInt
				for i, ps := range sl {

					psMapToStrVal := ps.(map[string]interface{})
					prefixSetsInt[i] = &ves_io_schema.ObjectRefType{}

					prefixSetsInt[i].Kind = "ip_prefix_set"

					if v, ok := psMapToStrVal["name"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Name = v.(string)
					}

					if v, ok := psMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Namespace = v.(string)
					}

					if v, ok := psMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Tenant = v.(string)
					}

					if v, ok := psMapToStrVal["uid"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("dst_ip_prefix_list"); ok && !dstIpChoiceTypeFound {

		dstIpChoiceTypeFound = true
		dstIpChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_DstIpPrefixList{}
		dstIpChoiceInt.DstIpPrefixList = &ves_io_schema_policy.PrefixMatchList{}
		updateSpec.DstIpChoice = dstIpChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["invert_match"]; ok && !isIntfNil(v) {

				dstIpChoiceInt.DstIpPrefixList.InvertMatch = v.(bool)
			}

			if v, ok := cs["ip_prefixes"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				dstIpChoiceInt.DstIpPrefixList.IpPrefixes = ls

			}

		}

	}

	if v, ok := d.GetOk("expiration_timestamp"); ok && !isIntfNil(v) {

		ts, err := parseTime(v.(string))
		if err != nil {
			return fmt.Errorf("error creating ServicePolicyRule, timestamp format is wrong: %s", err)
		}
		updateSpec.ExpirationTimestamp = ts

	}

	if v, ok := d.GetOk("headers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		headers := make([]*ves_io_schema_policy.HeaderMatcherType, len(sl))
		updateSpec.Headers = headers
		for i, set := range sl {
			headers[i] = &ves_io_schema_policy.HeaderMatcherType{}

			headersMapStrToI := set.(map[string]interface{})

			if w, ok := headersMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				headers[i].InvertMatcher = w.(bool)
			}

			matchTypeFound := false

			if v, ok := headersMapStrToI["check_not_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.HeaderMatcherType_CheckNotPresent{}
					matchInt.CheckNotPresent = &ves_io_schema.Empty{}
					headers[i].Match = matchInt
				}

			}

			if v, ok := headersMapStrToI["check_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.HeaderMatcherType_CheckPresent{}
					matchInt.CheckPresent = &ves_io_schema.Empty{}
					headers[i].Match = matchInt
				}

			}

			if v, ok := headersMapStrToI["item"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.HeaderMatcherType_Item{}
				matchInt.Item = &ves_io_schema_policy.MatcherType{}
				headers[i].Match = matchInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.ExactValues = ls

					}

					if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.RegexValues = ls

					}

					if v, ok := cs["transformers"]; ok && !isIntfNil(v) {

						transformersList := []ves_io_schema_policy.Transformer{}
						for _, j := range v.([]interface{}) {
							transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
						}
						matchInt.Item.Transformers = transformersList

					}

				}

			}

			if v, ok := headersMapStrToI["presence"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.HeaderMatcherType_Presence{}

				headers[i].Match = matchInt

				matchInt.Presence =
					v.(bool)

			}

			if w, ok := headersMapStrToI["name"]; ok && !isIntfNil(w) {
				headers[i].Name = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("http_method"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		httpMethod := &ves_io_schema_policy.HttpMethodMatcherType{}
		updateSpec.HttpMethod = httpMethod
		for _, set := range sl {

			httpMethodMapStrToI := set.(map[string]interface{})

			if w, ok := httpMethodMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				httpMethod.InvertMatcher = w.(bool)
			}

			if v, ok := httpMethodMapStrToI["methods"]; ok && !isIntfNil(v) {

				methodsList := []ves_io_schema.HttpMethod{}
				for _, j := range v.([]interface{}) {
					methodsList = append(methodsList, ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[j.(string)]))
				}
				httpMethod.Methods = methodsList

			}

		}

	}

	ipChoiceTypeFound := false

	if v, ok := d.GetOk("any_ip"); ok && !ipChoiceTypeFound {

		ipChoiceTypeFound = true

		if v.(bool) {
			ipChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_AnyIp{}
			ipChoiceInt.AnyIp = &ves_io_schema.Empty{}
			updateSpec.IpChoice = ipChoiceInt
		}

	}

	if v, ok := d.GetOk("ip_matcher"); ok && !ipChoiceTypeFound {

		ipChoiceTypeFound = true
		ipChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_IpMatcher{}
		ipChoiceInt.IpMatcher = &ves_io_schema_policy.IpMatcherType{}
		updateSpec.IpChoice = ipChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["invert_matcher"]; ok && !isIntfNil(v) {

				ipChoiceInt.IpMatcher.InvertMatcher = v.(bool)
			}

			if v, ok := cs["prefix_sets"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				prefixSetsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				ipChoiceInt.IpMatcher.PrefixSets = prefixSetsInt
				for i, ps := range sl {

					psMapToStrVal := ps.(map[string]interface{})
					prefixSetsInt[i] = &ves_io_schema.ObjectRefType{}

					prefixSetsInt[i].Kind = "ip_prefix_set"

					if v, ok := psMapToStrVal["name"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Name = v.(string)
					}

					if v, ok := psMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Namespace = v.(string)
					}

					if v, ok := psMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Tenant = v.(string)
					}

					if v, ok := psMapToStrVal["uid"]; ok && !isIntfNil(v) {
						prefixSetsInt[i].Uid = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("ip_prefix_list"); ok && !ipChoiceTypeFound {

		ipChoiceTypeFound = true
		ipChoiceInt := &ves_io_schema_service_policy_rule.ReplaceSpecType_IpPrefixList{}
		ipChoiceInt.IpPrefixList = &ves_io_schema_policy.PrefixMatchList{}
		updateSpec.IpChoice = ipChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["invert_match"]; ok && !isIntfNil(v) {

				ipChoiceInt.IpPrefixList.InvertMatch = v.(bool)
			}

			if v, ok := cs["ip_prefixes"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				ipChoiceInt.IpPrefixList.IpPrefixes = ls

			}

		}

	}

	if v, ok := d.GetOk("l4_dest_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		l4DestMatcher := &ves_io_schema_policy.L4DestMatcherType{}
		updateSpec.L4DestMatcher = l4DestMatcher
		for _, set := range sl {

			l4DestMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := l4DestMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				l4DestMatcher.InvertMatcher = w.(bool)
			}

			if v, ok := l4DestMatcherMapStrToI["l4_dests"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				l4Dests := make([]*ves_io_schema.L4DestType, len(sl))
				l4DestMatcher.L4Dests = l4Dests
				for i, set := range sl {
					l4Dests[i] = &ves_io_schema.L4DestType{}

					l4DestsMapStrToI := set.(map[string]interface{})

					if w, ok := l4DestsMapStrToI["ports"]; ok && !isIntfNil(w) {
						l4Dests[i].Ports = w.(string)
					}

					if w, ok := l4DestsMapStrToI["prefixes"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						l4Dests[i].Prefixes = ls
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("label_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		labelMatcher := &ves_io_schema.LabelMatcherType{}
		updateSpec.LabelMatcher = labelMatcher
		for _, set := range sl {

			labelMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := labelMatcherMapStrToI["keys"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				labelMatcher.Keys = ls
			}

		}

	}

	if v, ok := d.GetOk("malicious_user_mitigation"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		maliciousUserMitigation := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationType{}
		updateSpec.MaliciousUserMitigation = maliciousUserMitigation
		for _, set := range sl {

			maliciousUserMitigationMapStrToI := set.(map[string]interface{})

			if v, ok := maliciousUserMitigationMapStrToI["rules"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				rules := make([]*ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationRule, len(sl))
				maliciousUserMitigation.Rules = rules
				for i, set := range sl {
					rules[i] = &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationRule{}

					rulesMapStrToI := set.(map[string]interface{})

					if v, ok := rulesMapStrToI["mitigation_action"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						mitigationAction := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction{}
						rules[i].MitigationAction = mitigationAction
						for _, set := range sl {

							mitigationActionMapStrToI := set.(map[string]interface{})

							mitigationActionTypeFound := false

							if v, ok := mitigationActionMapStrToI["alert_only"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_AlertOnly{}
									mitigationActionInt.AlertOnly = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["block_temporarily"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_BlockTemporarily{}
									mitigationActionInt.BlockTemporarily = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["captcha_challenge"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_CaptchaChallenge{}
									mitigationActionInt.CaptchaChallenge = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["javascript_challenge"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_JavascriptChallenge{}
									mitigationActionInt.JavascriptChallenge = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

							if v, ok := mitigationActionMapStrToI["none"]; ok && !isIntfNil(v) && !mitigationActionTypeFound {

								mitigationActionTypeFound = true

								if v.(bool) {
									mitigationActionInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserMitigationAction_None{}
									mitigationActionInt.None = &ves_io_schema.Empty{}
									mitigationAction.MitigationAction = mitigationActionInt
								}

							}

						}

					}

					if v, ok := rulesMapStrToI["threat_level"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						threatLevel := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel{}
						rules[i].ThreatLevel = threatLevel
						for _, set := range sl {

							threatLevelMapStrToI := set.(map[string]interface{})

							threatLevelTypeFound := false

							if v, ok := threatLevelMapStrToI["high"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_High{}
									threatLevelInt.High = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

							if v, ok := threatLevelMapStrToI["low"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_Low{}
									threatLevelInt.Low = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

							if v, ok := threatLevelMapStrToI["medium"]; ok && !isIntfNil(v) && !threatLevelTypeFound {

								threatLevelTypeFound = true

								if v.(bool) {
									threatLevelInt := &ves_io_schema_malicious_user_mitigation.MaliciousUserThreatLevel_Medium{}
									threatLevelInt.Medium = &ves_io_schema.Empty{}
									threatLevel.ThreatLevel = threatLevelInt
								}

							}

						}

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("malicious_user_mitigation_bypass"); ok && !isIntfNil(v) {

		if v.(bool) {
			updateSpec.MaliciousUserMitigationBypass = &ves_io_schema.Empty{}
		}

	}

	if v, ok := d.GetOk("path"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		path := &ves_io_schema_policy.PathMatcherType{}
		updateSpec.Path = path
		for _, set := range sl {

			pathMapStrToI := set.(map[string]interface{})

			if w, ok := pathMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				path.ExactValues = ls
			}

			if w, ok := pathMapStrToI["prefix_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				path.PrefixValues = ls
			}

			if w, ok := pathMapStrToI["regex_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				path.RegexValues = ls
			}

			if v, ok := pathMapStrToI["transformers"]; ok && !isIntfNil(v) {

				transformersList := []ves_io_schema_policy.Transformer{}
				for _, j := range v.([]interface{}) {
					transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
				}
				path.Transformers = transformersList

			}

		}

	}

	if v, ok := d.GetOk("port_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		portMatcher := &ves_io_schema_policy.PortMatcherType{}
		updateSpec.PortMatcher = portMatcher
		for _, set := range sl {

			portMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := portMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				portMatcher.InvertMatcher = w.(bool)
			}

			if w, ok := portMatcherMapStrToI["ports"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				portMatcher.Ports = ls
			}

		}

	}

	if v, ok := d.GetOk("query_params"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		queryParams := make([]*ves_io_schema_policy.QueryParameterMatcherType, len(sl))
		updateSpec.QueryParams = queryParams
		for i, set := range sl {
			queryParams[i] = &ves_io_schema_policy.QueryParameterMatcherType{}

			queryParamsMapStrToI := set.(map[string]interface{})

			if w, ok := queryParamsMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				queryParams[i].InvertMatcher = w.(bool)
			}

			if w, ok := queryParamsMapStrToI["key"]; ok && !isIntfNil(w) {
				queryParams[i].Key = w.(string)
			}

			matchTypeFound := false

			if v, ok := queryParamsMapStrToI["check_not_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.QueryParameterMatcherType_CheckNotPresent{}
					matchInt.CheckNotPresent = &ves_io_schema.Empty{}
					queryParams[i].Match = matchInt
				}

			}

			if v, ok := queryParamsMapStrToI["check_present"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true

				if v.(bool) {
					matchInt := &ves_io_schema_policy.QueryParameterMatcherType_CheckPresent{}
					matchInt.CheckPresent = &ves_io_schema.Empty{}
					queryParams[i].Match = matchInt
				}

			}

			if v, ok := queryParamsMapStrToI["item"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.QueryParameterMatcherType_Item{}
				matchInt.Item = &ves_io_schema_policy.MatcherType{}
				queryParams[i].Match = matchInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exact_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.ExactValues = ls

					}

					if v, ok := cs["regex_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						matchInt.Item.RegexValues = ls

					}

					if v, ok := cs["transformers"]; ok && !isIntfNil(v) {

						transformersList := []ves_io_schema_policy.Transformer{}
						for _, j := range v.([]interface{}) {
							transformersList = append(transformersList, ves_io_schema_policy.Transformer(ves_io_schema_policy.Transformer_value[j.(string)]))
						}
						matchInt.Item.Transformers = transformersList

					}

				}

			}

			if v, ok := queryParamsMapStrToI["presence"]; ok && !isIntfNil(v) && !matchTypeFound {

				matchTypeFound = true
				matchInt := &ves_io_schema_policy.QueryParameterMatcherType_Presence{}

				queryParams[i].Match = matchInt

				matchInt.Presence =
					v.(bool)

			}

		}

	}

	if v, ok := d.GetOk("rate_limiter"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		rateLimiterInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.RateLimiter = rateLimiterInt
		for i, ps := range sl {

			rlMapToStrVal := ps.(map[string]interface{})
			rateLimiterInt[i] = &ves_io_schema.ObjectRefType{}

			rateLimiterInt[i].Kind = "rate_limiter"

			if v, ok := rlMapToStrVal["name"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Name = v.(string)
			}

			if v, ok := rlMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Namespace = v.(string)
			}

			if v, ok := rlMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Tenant = v.(string)
			}

			if v, ok := rlMapToStrVal["uid"]; ok && !isIntfNil(v) {
				rateLimiterInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("scheme"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		updateSpec.Scheme = ls

	}

	if v, ok := d.GetOk("tls_fingerprint_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		tlsFingerprintMatcher := &ves_io_schema_policy.TlsFingerprintMatcherType{}
		updateSpec.TlsFingerprintMatcher = tlsFingerprintMatcher
		for _, set := range sl {

			tlsFingerprintMatcherMapStrToI := set.(map[string]interface{})

			if v, ok := tlsFingerprintMatcherMapStrToI["classes"]; ok && !isIntfNil(v) {

				classesList := []ves_io_schema_policy.KnownTlsFingerprintClass{}
				for _, j := range v.([]interface{}) {
					classesList = append(classesList, ves_io_schema_policy.KnownTlsFingerprintClass(ves_io_schema_policy.KnownTlsFingerprintClass_value[j.(string)]))
				}
				tlsFingerprintMatcher.Classes = classesList

			}

			if w, ok := tlsFingerprintMatcherMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				tlsFingerprintMatcher.ExactValues = ls
			}

			if w, ok := tlsFingerprintMatcherMapStrToI["excluded_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				tlsFingerprintMatcher.ExcludedValues = ls
			}

		}

	}

	if v, ok := d.GetOk("url_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		urlMatcher := &ves_io_schema_policy.URLMatcherType{}
		updateSpec.UrlMatcher = urlMatcher
		for _, set := range sl {

			urlMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := urlMatcherMapStrToI["invert_matcher"]; ok && !isIntfNil(w) {
				urlMatcher.InvertMatcher = w.(bool)
			}

			if v, ok := urlMatcherMapStrToI["url_items"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				urlItems := make([]*ves_io_schema_policy.URLItem, len(sl))
				urlMatcher.UrlItems = urlItems
				for i, set := range sl {
					urlItems[i] = &ves_io_schema_policy.URLItem{}

					urlItemsMapStrToI := set.(map[string]interface{})

					domainChoiceTypeFound := false

					if v, ok := urlItemsMapStrToI["domain_regex"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_policy.URLItem_DomainRegex{}

						urlItems[i].DomainChoice = domainChoiceInt

						domainChoiceInt.DomainRegex = v.(string)

					}

					if v, ok := urlItemsMapStrToI["domain_value"]; ok && !isIntfNil(v) && !domainChoiceTypeFound {

						domainChoiceTypeFound = true
						domainChoiceInt := &ves_io_schema_policy.URLItem_DomainValue{}

						urlItems[i].DomainChoice = domainChoiceInt

						domainChoiceInt.DomainValue = v.(string)

					}

					pathChoiceTypeFound := false

					if v, ok := urlItemsMapStrToI["path_prefix"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_policy.URLItem_PathPrefix{}

						urlItems[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathPrefix = v.(string)

					}

					if v, ok := urlItemsMapStrToI["path_regex"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_policy.URLItem_PathRegex{}

						urlItems[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathRegex = v.(string)

					}

					if v, ok := urlItemsMapStrToI["path_value"]; ok && !isIntfNil(v) && !pathChoiceTypeFound {

						pathChoiceTypeFound = true
						pathChoiceInt := &ves_io_schema_policy.URLItem_PathValue{}

						urlItems[i].PathChoice = pathChoiceInt

						pathChoiceInt.PathValue = v.(string)

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("virtual_host_matcher"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		virtualHostMatcher := &ves_io_schema_policy.MatcherTypeBasic{}
		updateSpec.VirtualHostMatcher = virtualHostMatcher
		for _, set := range sl {

			virtualHostMatcherMapStrToI := set.(map[string]interface{})

			if w, ok := virtualHostMatcherMapStrToI["exact_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				virtualHostMatcher.ExactValues = ls
			}

			if w, ok := virtualHostMatcherMapStrToI["regex_values"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				virtualHostMatcher.RegexValues = ls
			}

		}

	}

	if v, ok := d.GetOk("waf_action"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		wafAction := &ves_io_schema_policy.WafAction{}
		updateSpec.WafAction = wafAction
		for _, set := range sl {

			wafActionMapStrToI := set.(map[string]interface{})

			actionTypeTypeFound := false

			if v, ok := wafActionMapStrToI["none"]; ok && !isIntfNil(v) && !actionTypeTypeFound {

				actionTypeTypeFound = true

				if v.(bool) {
					actionTypeInt := &ves_io_schema_policy.WafAction_None{}
					actionTypeInt.None = &ves_io_schema.Empty{}
					wafAction.ActionType = actionTypeInt
				}

			}

			if v, ok := wafActionMapStrToI["waf_inline_rule_control"]; ok && !isIntfNil(v) && !actionTypeTypeFound {

				actionTypeTypeFound = true
				actionTypeInt := &ves_io_schema_policy.WafAction_WafInlineRuleControl{}
				actionTypeInt.WafInlineRuleControl = &ves_io_schema_policy.WafInlineRuleControl{}
				wafAction.ActionType = actionTypeInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exclude_rule_ids"]; ok && !isIntfNil(v) {

						exclude_rule_idsList := []ves_io_schema_waf_rule_list.WafRuleID{}
						for _, j := range v.([]interface{}) {
							exclude_rule_idsList = append(exclude_rule_idsList, ves_io_schema_waf_rule_list.WafRuleID(ves_io_schema_waf_rule_list.WafRuleID_value[j.(string)]))
						}
						actionTypeInt.WafInlineRuleControl.ExcludeRuleIds = exclude_rule_idsList

					}

				}

			}

			if v, ok := wafActionMapStrToI["waf_rule_control"]; ok && !isIntfNil(v) && !actionTypeTypeFound {

				actionTypeTypeFound = true
				actionTypeInt := &ves_io_schema_policy.WafAction_WafRuleControl{}
				actionTypeInt.WafRuleControl = &ves_io_schema_policy.WafRuleControl{}
				wafAction.ActionType = actionTypeInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["exclude_rule_ids"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						excludeRuleIdsInt := make([]*ves_io_schema.ObjectRefType, len(sl))
						actionTypeInt.WafRuleControl.ExcludeRuleIds = excludeRuleIdsInt
						for i, ps := range sl {

							eriMapToStrVal := ps.(map[string]interface{})
							excludeRuleIdsInt[i] = &ves_io_schema.ObjectRefType{}

							excludeRuleIdsInt[i].Kind = "waf_rule_list"

							if v, ok := eriMapToStrVal["name"]; ok && !isIntfNil(v) {
								excludeRuleIdsInt[i].Name = v.(string)
							}

							if v, ok := eriMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								excludeRuleIdsInt[i].Namespace = v.(string)
							}

							if v, ok := eriMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								excludeRuleIdsInt[i].Tenant = v.(string)
							}

							if v, ok := eriMapToStrVal["uid"]; ok && !isIntfNil(v) {
								excludeRuleIdsInt[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := wafActionMapStrToI["waf_skip_processing"]; ok && !isIntfNil(v) && !actionTypeTypeFound {

				actionTypeTypeFound = true

				if v.(bool) {
					actionTypeInt := &ves_io_schema_policy.WafAction_WafSkipProcessing{}
					actionTypeInt.WafSkipProcessing = &ves_io_schema.Empty{}
					wafAction.ActionType = actionTypeInt
				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra ServicePolicyRule obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_service_policy_rule.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating ServicePolicyRule: %s", err)
	}

	return resourceVolterraServicePolicyRuleRead(d, meta)
}

func resourceVolterraServicePolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_service_policy_rule.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ServicePolicyRule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ServicePolicyRule before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra ServicePolicyRule obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_service_policy_rule.ObjectType, namespace, name)
}
