//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-tf-provider. DO NOT EDIT.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_network_policy_rule "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/network_policy_rule"
)

// resourceVolterraNetworkPolicyRule is implementation of Volterra's NetworkPolicyRule resources
func resourceVolterraNetworkPolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraNetworkPolicyRuleCreate,
		Read:   resourceVolterraNetworkPolicyRuleRead,
		Update: resourceVolterraNetworkPolicyRuleUpdate,
		Delete: resourceVolterraNetworkPolicyRuleDelete,

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

			"action": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"advanced_action": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"action": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"label_matcher": {

				Type:     schema.TypeList,
				MaxItems: 1,
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

			"ports": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"ip_prefix_set": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ref": {

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

			"prefix": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ipv6_prefix": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"prefix": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"prefix_selector": {

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
		},
	}
}

// resourceVolterraNetworkPolicyRuleCreate creates NetworkPolicyRule resource
func resourceVolterraNetworkPolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_network_policy_rule.CreateSpecType{}
	createReq := &ves_io_schema_network_policy_rule.CreateRequest{
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

	//action
	if v, ok := d.GetOk("action"); ok && !isIntfNil(v) {

		createSpec.Action = ves_io_schema_network_policy_rule.NetworkPolicyRuleAction(ves_io_schema_network_policy_rule.NetworkPolicyRuleAction_value[v.(string)])

	}

	//advanced_action
	if v, ok := d.GetOk("advanced_action"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		advancedAction := &ves_io_schema_network_policy_rule.NetworkPolicyRuleAdvancedAction{}
		createSpec.AdvancedAction = advancedAction
		for _, set := range sl {
			if set != nil {
				advancedActionMapStrToI := set.(map[string]interface{})

				if v, ok := advancedActionMapStrToI["action"]; ok && !isIntfNil(v) {

					advancedAction.Action = ves_io_schema_network_policy_rule.LogAction(ves_io_schema_network_policy_rule.LogAction_value[v.(string)])

				}

			}
		}

	}

	//label_matcher
	if v, ok := d.GetOk("label_matcher"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		labelMatcher := &ves_io_schema.LabelMatcherType{}
		createSpec.LabelMatcher = labelMatcher
		for _, set := range sl {
			if set != nil {
				labelMatcherMapStrToI := set.(map[string]interface{})

				if w, ok := labelMatcherMapStrToI["keys"]; ok && !isIntfNil(w) {
					ls := make([]string, len(w.([]interface{})))
					for i, v := range w.([]interface{}) {
						if v == nil {
							return fmt.Errorf("please provide valid non-empty string value of field keys")
						}
						if str, ok := v.(string); ok {
							ls[i] = str
						}
					}
					labelMatcher.Keys = ls
				}

			}
		}

	}

	//ports
	if v, ok := d.GetOk("ports"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			if v == nil {
				return fmt.Errorf("please provide valid non-empty string value of field ports")
			}
			if str, ok := v.(string); ok {
				ls[i] = str
			}
		}
		createSpec.Ports = ls

	}

	//protocol
	if v, ok := d.GetOk("protocol"); ok && !isIntfNil(v) {

		createSpec.Protocol =
			v.(string)

	}

	//remote_endpoint

	remoteEndpointTypeFound := false

	if v, ok := d.GetOk("ip_prefix_set"); ok && !isIntfNil(v) && !remoteEndpointTypeFound {

		remoteEndpointTypeFound = true
		remoteEndpointInt := &ves_io_schema_network_policy_rule.CreateSpecType_IpPrefixSet{}
		remoteEndpointInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
		createSpec.RemoteEndpoint = remoteEndpointInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["ref"]; ok && !isIntfNil(v) {

					sl := v.([]interface{})
					refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
					remoteEndpointInt.IpPrefixSet.Ref = refInt
					for i, ps := range sl {

						rMapToStrVal := ps.(map[string]interface{})
						refInt[i] = &ves_io_schema.ObjectRefType{}

						refInt[i].Kind = "ip_prefix_set"

						if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
							refInt[i].Name = v.(string)
						}

						if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
							refInt[i].Namespace = v.(string)
						}

						if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
							refInt[i].Tenant = v.(string)
						}

						if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
							refInt[i].Uid = v.(string)
						}

					}

				}

			}
		}

	}

	if v, ok := d.GetOk("prefix"); ok && !isIntfNil(v) && !remoteEndpointTypeFound {

		remoteEndpointTypeFound = true
		remoteEndpointInt := &ves_io_schema_network_policy_rule.CreateSpecType_Prefix{}
		remoteEndpointInt.Prefix = &ves_io_schema.PrefixListType{}
		createSpec.RemoteEndpoint = remoteEndpointInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["ipv6_prefix"]; ok && !isIntfNil(v) {

					ls := make([]string, len(v.([]interface{})))
					for i, v := range v.([]interface{}) {
						if v == nil {
							return fmt.Errorf("please provide valid non-empty string value of field ipv6_prefix")
						}
						if str, ok := v.(string); ok {
							ls[i] = str
						}
					}
					remoteEndpointInt.Prefix.Ipv6Prefix = ls

				}

				if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

					ls := make([]string, len(v.([]interface{})))
					for i, v := range v.([]interface{}) {
						if v == nil {
							return fmt.Errorf("please provide valid non-empty string value of field prefix")
						}
						if str, ok := v.(string); ok {
							ls[i] = str
						}
					}
					remoteEndpointInt.Prefix.Prefix = ls

				}

			}
		}

	}

	if v, ok := d.GetOk("prefix_selector"); ok && !isIntfNil(v) && !remoteEndpointTypeFound {

		remoteEndpointTypeFound = true
		remoteEndpointInt := &ves_io_schema_network_policy_rule.CreateSpecType_PrefixSelector{}
		remoteEndpointInt.PrefixSelector = &ves_io_schema.LabelSelectorType{}
		createSpec.RemoteEndpoint = remoteEndpointInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

					ls := make([]string, len(v.([]interface{})))
					for i, v := range v.([]interface{}) {
						if v == nil {
							return fmt.Errorf("please provide valid non-empty string value of field expressions")
						}
						if str, ok := v.(string); ok {
							ls[i] = str
						}
					}
					remoteEndpointInt.PrefixSelector.Expressions = ls

				}

			}
		}

	}

	log.Printf("[DEBUG] Creating Volterra NetworkPolicyRule object with struct: %+v", createReq)

	createNetworkPolicyRuleResp, err := client.CreateObject(context.Background(), ves_io_schema_network_policy_rule.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating NetworkPolicyRule: %s", err)
	}
	d.SetId(createNetworkPolicyRuleResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraNetworkPolicyRuleRead(d, meta)
}

func resourceVolterraNetworkPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_network_policy_rule.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkPolicyRule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkPolicyRule %q: %s", d.Id(), err)
	}
	return setNetworkPolicyRuleFields(client, d, resp)
}

func setNetworkPolicyRuleFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraNetworkPolicyRuleUpdate updates NetworkPolicyRule resource
func resourceVolterraNetworkPolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_network_policy_rule.ReplaceSpecType{}
	updateReq := &ves_io_schema_network_policy_rule.ReplaceRequest{
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

		updateSpec.Action = ves_io_schema_network_policy_rule.NetworkPolicyRuleAction(ves_io_schema_network_policy_rule.NetworkPolicyRuleAction_value[v.(string)])

	}

	if v, ok := d.GetOk("advanced_action"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		advancedAction := &ves_io_schema_network_policy_rule.NetworkPolicyRuleAdvancedAction{}
		updateSpec.AdvancedAction = advancedAction
		for _, set := range sl {
			if set != nil {
				advancedActionMapStrToI := set.(map[string]interface{})

				if v, ok := advancedActionMapStrToI["action"]; ok && !isIntfNil(v) {

					advancedAction.Action = ves_io_schema_network_policy_rule.LogAction(ves_io_schema_network_policy_rule.LogAction_value[v.(string)])

				}

			}
		}

	}

	if v, ok := d.GetOk("label_matcher"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		labelMatcher := &ves_io_schema.LabelMatcherType{}
		updateSpec.LabelMatcher = labelMatcher
		for _, set := range sl {
			if set != nil {
				labelMatcherMapStrToI := set.(map[string]interface{})

				if w, ok := labelMatcherMapStrToI["keys"]; ok && !isIntfNil(w) {
					ls := make([]string, len(w.([]interface{})))
					for i, v := range w.([]interface{}) {
						if v == nil {
							return fmt.Errorf("please provide valid non-empty string value of field keys")
						}
						if str, ok := v.(string); ok {
							ls[i] = str
						}
					}
					labelMatcher.Keys = ls
				}

			}
		}

	}

	if v, ok := d.GetOk("ports"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			if v == nil {
				return fmt.Errorf("please provide valid non-empty string value of field ports")
			}
			if str, ok := v.(string); ok {
				ls[i] = str
			}
		}
		updateSpec.Ports = ls

	}

	if v, ok := d.GetOk("protocol"); ok && !isIntfNil(v) {

		updateSpec.Protocol =
			v.(string)

	}

	remoteEndpointTypeFound := false

	if v, ok := d.GetOk("ip_prefix_set"); ok && !isIntfNil(v) && !remoteEndpointTypeFound {

		remoteEndpointTypeFound = true
		remoteEndpointInt := &ves_io_schema_network_policy_rule.ReplaceSpecType_IpPrefixSet{}
		remoteEndpointInt.IpPrefixSet = &ves_io_schema.IpPrefixSetRefType{}
		updateSpec.RemoteEndpoint = remoteEndpointInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["ref"]; ok && !isIntfNil(v) {

					sl := v.([]interface{})
					refInt := make([]*ves_io_schema.ObjectRefType, len(sl))
					remoteEndpointInt.IpPrefixSet.Ref = refInt
					for i, ps := range sl {

						rMapToStrVal := ps.(map[string]interface{})
						refInt[i] = &ves_io_schema.ObjectRefType{}

						refInt[i].Kind = "ip_prefix_set"

						if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
							refInt[i].Name = v.(string)
						}

						if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
							refInt[i].Namespace = v.(string)
						}

						if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
							refInt[i].Tenant = v.(string)
						}

						if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
							refInt[i].Uid = v.(string)
						}

					}

				}

			}
		}

	}

	if v, ok := d.GetOk("prefix"); ok && !isIntfNil(v) && !remoteEndpointTypeFound {

		remoteEndpointTypeFound = true
		remoteEndpointInt := &ves_io_schema_network_policy_rule.ReplaceSpecType_Prefix{}
		remoteEndpointInt.Prefix = &ves_io_schema.PrefixListType{}
		updateSpec.RemoteEndpoint = remoteEndpointInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["ipv6_prefix"]; ok && !isIntfNil(v) {

					ls := make([]string, len(v.([]interface{})))
					for i, v := range v.([]interface{}) {
						if v == nil {
							return fmt.Errorf("please provide valid non-empty string value of field ipv6_prefix")
						}
						if str, ok := v.(string); ok {
							ls[i] = str
						}
					}
					remoteEndpointInt.Prefix.Ipv6Prefix = ls

				}

				if v, ok := cs["prefix"]; ok && !isIntfNil(v) {

					ls := make([]string, len(v.([]interface{})))
					for i, v := range v.([]interface{}) {
						if v == nil {
							return fmt.Errorf("please provide valid non-empty string value of field prefix")
						}
						if str, ok := v.(string); ok {
							ls[i] = str
						}
					}
					remoteEndpointInt.Prefix.Prefix = ls

				}

			}
		}

	}

	if v, ok := d.GetOk("prefix_selector"); ok && !isIntfNil(v) && !remoteEndpointTypeFound {

		remoteEndpointTypeFound = true
		remoteEndpointInt := &ves_io_schema_network_policy_rule.ReplaceSpecType_PrefixSelector{}
		remoteEndpointInt.PrefixSelector = &ves_io_schema.LabelSelectorType{}
		updateSpec.RemoteEndpoint = remoteEndpointInt

		sl := v.([]interface{})
		for _, set := range sl {
			if set != nil {
				cs := set.(map[string]interface{})

				if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

					ls := make([]string, len(v.([]interface{})))
					for i, v := range v.([]interface{}) {
						if v == nil {
							return fmt.Errorf("please provide valid non-empty string value of field expressions")
						}
						if str, ok := v.(string); ok {
							ls[i] = str
						}
					}
					remoteEndpointInt.PrefixSelector.Expressions = ls

				}

			}
		}

	}

	log.Printf("[DEBUG] Updating Volterra NetworkPolicyRule obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_network_policy_rule.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating NetworkPolicyRule: %s", err)
	}

	return resourceVolterraNetworkPolicyRuleRead(d, meta)
}

func resourceVolterraNetworkPolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_network_policy_rule.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkPolicyRule %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkPolicyRule before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra NetworkPolicyRule obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_network_policy_rule.ObjectType, namespace, name, opts...)
}
