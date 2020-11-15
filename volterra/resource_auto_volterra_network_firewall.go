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
)

// resourceVolterraNetworkFirewall is implementation of Volterra's NetworkFirewall resources
func resourceVolterraNetworkFirewall() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraNetworkFirewallCreate,
		Read:   resourceVolterraNetworkFirewallRead,
		Update: resourceVolterraNetworkFirewallUpdate,
		Delete: resourceVolterraNetworkFirewallDelete,

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

			"active_fast_acls": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fast_acls": {

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

			"disable_fast_acl": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"fast_acl_set": {

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

			"disable_forward_proxy_policy": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"forward_proxy_policy_set": {

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

			"disable_network_policy": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"network_policy_set": {

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
	}
}

// resourceVolterraNetworkFirewallCreate creates NetworkFirewall resource
func resourceVolterraNetworkFirewallCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_network_firewall.CreateSpecType{}
	createReq := &ves_io_schema_network_firewall.CreateRequest{
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

	fastAclChoiceTypeFound := false

	if v, ok := d.GetOk("active_fast_acls"); ok && !fastAclChoiceTypeFound {

		fastAclChoiceTypeFound = true
		fastAclChoiceInt := &ves_io_schema_network_firewall.CreateSpecType_ActiveFastAcls{}
		fastAclChoiceInt.ActiveFastAcls = &ves_io_schema_network_firewall.ActiveFastACLsType{}
		createSpec.FastAclChoice = fastAclChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["fast_acls"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				fastAclsInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				fastAclChoiceInt.ActiveFastAcls.FastAcls = fastAclsInt
				for i, ps := range sl {

					faMapToStrVal := ps.(map[string]interface{})
					fastAclsInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := faMapToStrVal["name"]; ok && !isIntfNil(v) {
						fastAclsInt[i].Name = v.(string)
					}

					if v, ok := faMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						fastAclsInt[i].Namespace = v.(string)
					}

					if v, ok := faMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						fastAclsInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("disable_fast_acl"); ok && !fastAclChoiceTypeFound {

		fastAclChoiceTypeFound = true

		if v.(bool) {
			fastAclChoiceInt := &ves_io_schema_network_firewall.CreateSpecType_DisableFastAcl{}
			fastAclChoiceInt.DisableFastAcl = &ves_io_schema.Empty{}
			createSpec.FastAclChoice = fastAclChoiceInt
		}

	}

	if v, ok := d.GetOk("fast_acl_set"); ok && !fastAclChoiceTypeFound {

		fastAclChoiceTypeFound = true
		fastAclChoiceInt := &ves_io_schema_network_firewall.CreateSpecType_FastAclSet{}
		fastAclChoiceInt.FastAclSet = &ves_io_schema_views.ObjectRefType{}
		createSpec.FastAclChoice = fastAclChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				fastAclChoiceInt.FastAclSet.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				fastAclChoiceInt.FastAclSet.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				fastAclChoiceInt.FastAclSet.Tenant = v.(string)
			}

		}

	}

	forwardProxyPolicyChoiceTypeFound := false

	if v, ok := d.GetOk("active_forward_proxy_policies"); ok && !forwardProxyPolicyChoiceTypeFound {

		forwardProxyPolicyChoiceTypeFound = true
		forwardProxyPolicyChoiceInt := &ves_io_schema_network_firewall.CreateSpecType_ActiveForwardProxyPolicies{}
		forwardProxyPolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
		createSpec.ForwardProxyPolicyChoice = forwardProxyPolicyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["forward_proxy_policies"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				forwardProxyPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				forwardProxyPolicyChoiceInt.ActiveForwardProxyPolicies.ForwardProxyPolicies = forwardProxyPoliciesInt
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

	if v, ok := d.GetOk("disable_forward_proxy_policy"); ok && !forwardProxyPolicyChoiceTypeFound {

		forwardProxyPolicyChoiceTypeFound = true

		if v.(bool) {
			forwardProxyPolicyChoiceInt := &ves_io_schema_network_firewall.CreateSpecType_DisableForwardProxyPolicy{}
			forwardProxyPolicyChoiceInt.DisableForwardProxyPolicy = &ves_io_schema.Empty{}
			createSpec.ForwardProxyPolicyChoice = forwardProxyPolicyChoiceInt
		}

	}

	if v, ok := d.GetOk("forward_proxy_policy_set"); ok && !forwardProxyPolicyChoiceTypeFound {

		forwardProxyPolicyChoiceTypeFound = true
		forwardProxyPolicyChoiceInt := &ves_io_schema_network_firewall.CreateSpecType_ForwardProxyPolicySet{}
		forwardProxyPolicyChoiceInt.ForwardProxyPolicySet = &ves_io_schema_views.ObjectRefType{}
		createSpec.ForwardProxyPolicyChoice = forwardProxyPolicyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				forwardProxyPolicyChoiceInt.ForwardProxyPolicySet.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				forwardProxyPolicyChoiceInt.ForwardProxyPolicySet.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				forwardProxyPolicyChoiceInt.ForwardProxyPolicySet.Tenant = v.(string)
			}

		}

	}

	networkPolicyChoiceTypeFound := false

	if v, ok := d.GetOk("active_network_policies"); ok && !networkPolicyChoiceTypeFound {

		networkPolicyChoiceTypeFound = true
		networkPolicyChoiceInt := &ves_io_schema_network_firewall.CreateSpecType_ActiveNetworkPolicies{}
		networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
		createSpec.NetworkPolicyChoice = networkPolicyChoiceInt

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

	if v, ok := d.GetOk("disable_network_policy"); ok && !networkPolicyChoiceTypeFound {

		networkPolicyChoiceTypeFound = true

		if v.(bool) {
			networkPolicyChoiceInt := &ves_io_schema_network_firewall.CreateSpecType_DisableNetworkPolicy{}
			networkPolicyChoiceInt.DisableNetworkPolicy = &ves_io_schema.Empty{}
			createSpec.NetworkPolicyChoice = networkPolicyChoiceInt
		}

	}

	if v, ok := d.GetOk("network_policy_set"); ok && !networkPolicyChoiceTypeFound {

		networkPolicyChoiceTypeFound = true
		networkPolicyChoiceInt := &ves_io_schema_network_firewall.CreateSpecType_NetworkPolicySet{}
		networkPolicyChoiceInt.NetworkPolicySet = &ves_io_schema_views.ObjectRefType{}
		createSpec.NetworkPolicyChoice = networkPolicyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				networkPolicyChoiceInt.NetworkPolicySet.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				networkPolicyChoiceInt.NetworkPolicySet.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				networkPolicyChoiceInt.NetworkPolicySet.Tenant = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra NetworkFirewall object with struct: %+v", createReq)

	createNetworkFirewallResp, err := client.CreateObject(context.Background(), ves_io_schema_network_firewall.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating NetworkFirewall: %s", err)
	}
	d.SetId(createNetworkFirewallResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraNetworkFirewallRead(d, meta)
}

func resourceVolterraNetworkFirewallRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_network_firewall.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkFirewall %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkFirewall %q: %s", d.Id(), err)
	}
	return setNetworkFirewallFields(client, d, resp)
}

func setNetworkFirewallFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraNetworkFirewallUpdate updates NetworkFirewall resource
func resourceVolterraNetworkFirewallUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_network_firewall.ReplaceSpecType{}
	updateReq := &ves_io_schema_network_firewall.ReplaceRequest{
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

	fastAclChoiceTypeFound := false

	if v, ok := d.GetOk("active_fast_acls"); ok && !fastAclChoiceTypeFound {

		fastAclChoiceTypeFound = true
		fastAclChoiceInt := &ves_io_schema_network_firewall.ReplaceSpecType_ActiveFastAcls{}
		fastAclChoiceInt.ActiveFastAcls = &ves_io_schema_network_firewall.ActiveFastACLsType{}
		updateSpec.FastAclChoice = fastAclChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["fast_acls"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				fastAclsInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				fastAclChoiceInt.ActiveFastAcls.FastAcls = fastAclsInt
				for i, ps := range sl {

					faMapToStrVal := ps.(map[string]interface{})
					fastAclsInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := faMapToStrVal["name"]; ok && !isIntfNil(v) {
						fastAclsInt[i].Name = v.(string)
					}

					if v, ok := faMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						fastAclsInt[i].Namespace = v.(string)
					}

					if v, ok := faMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						fastAclsInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("disable_fast_acl"); ok && !fastAclChoiceTypeFound {

		fastAclChoiceTypeFound = true

		if v.(bool) {
			fastAclChoiceInt := &ves_io_schema_network_firewall.ReplaceSpecType_DisableFastAcl{}
			fastAclChoiceInt.DisableFastAcl = &ves_io_schema.Empty{}
			updateSpec.FastAclChoice = fastAclChoiceInt
		}

	}

	if v, ok := d.GetOk("fast_acl_set"); ok && !fastAclChoiceTypeFound {

		fastAclChoiceTypeFound = true
		fastAclChoiceInt := &ves_io_schema_network_firewall.ReplaceSpecType_FastAclSet{}
		fastAclChoiceInt.FastAclSet = &ves_io_schema_views.ObjectRefType{}
		updateSpec.FastAclChoice = fastAclChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				fastAclChoiceInt.FastAclSet.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				fastAclChoiceInt.FastAclSet.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				fastAclChoiceInt.FastAclSet.Tenant = v.(string)
			}

		}

	}

	forwardProxyPolicyChoiceTypeFound := false

	if v, ok := d.GetOk("active_forward_proxy_policies"); ok && !forwardProxyPolicyChoiceTypeFound {

		forwardProxyPolicyChoiceTypeFound = true
		forwardProxyPolicyChoiceInt := &ves_io_schema_network_firewall.ReplaceSpecType_ActiveForwardProxyPolicies{}
		forwardProxyPolicyChoiceInt.ActiveForwardProxyPolicies = &ves_io_schema_network_firewall.ActiveForwardProxyPoliciesType{}
		updateSpec.ForwardProxyPolicyChoice = forwardProxyPolicyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["forward_proxy_policies"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				forwardProxyPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				forwardProxyPolicyChoiceInt.ActiveForwardProxyPolicies.ForwardProxyPolicies = forwardProxyPoliciesInt
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

	if v, ok := d.GetOk("disable_forward_proxy_policy"); ok && !forwardProxyPolicyChoiceTypeFound {

		forwardProxyPolicyChoiceTypeFound = true

		if v.(bool) {
			forwardProxyPolicyChoiceInt := &ves_io_schema_network_firewall.ReplaceSpecType_DisableForwardProxyPolicy{}
			forwardProxyPolicyChoiceInt.DisableForwardProxyPolicy = &ves_io_schema.Empty{}
			updateSpec.ForwardProxyPolicyChoice = forwardProxyPolicyChoiceInt
		}

	}

	if v, ok := d.GetOk("forward_proxy_policy_set"); ok && !forwardProxyPolicyChoiceTypeFound {

		forwardProxyPolicyChoiceTypeFound = true
		forwardProxyPolicyChoiceInt := &ves_io_schema_network_firewall.ReplaceSpecType_ForwardProxyPolicySet{}
		forwardProxyPolicyChoiceInt.ForwardProxyPolicySet = &ves_io_schema_views.ObjectRefType{}
		updateSpec.ForwardProxyPolicyChoice = forwardProxyPolicyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				forwardProxyPolicyChoiceInt.ForwardProxyPolicySet.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				forwardProxyPolicyChoiceInt.ForwardProxyPolicySet.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				forwardProxyPolicyChoiceInt.ForwardProxyPolicySet.Tenant = v.(string)
			}

		}

	}

	networkPolicyChoiceTypeFound := false

	if v, ok := d.GetOk("active_network_policies"); ok && !networkPolicyChoiceTypeFound {

		networkPolicyChoiceTypeFound = true
		networkPolicyChoiceInt := &ves_io_schema_network_firewall.ReplaceSpecType_ActiveNetworkPolicies{}
		networkPolicyChoiceInt.ActiveNetworkPolicies = &ves_io_schema_network_firewall.ActiveNetworkPoliciesType{}
		updateSpec.NetworkPolicyChoice = networkPolicyChoiceInt

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

	if v, ok := d.GetOk("disable_network_policy"); ok && !networkPolicyChoiceTypeFound {

		networkPolicyChoiceTypeFound = true

		if v.(bool) {
			networkPolicyChoiceInt := &ves_io_schema_network_firewall.ReplaceSpecType_DisableNetworkPolicy{}
			networkPolicyChoiceInt.DisableNetworkPolicy = &ves_io_schema.Empty{}
			updateSpec.NetworkPolicyChoice = networkPolicyChoiceInt
		}

	}

	if v, ok := d.GetOk("network_policy_set"); ok && !networkPolicyChoiceTypeFound {

		networkPolicyChoiceTypeFound = true
		networkPolicyChoiceInt := &ves_io_schema_network_firewall.ReplaceSpecType_NetworkPolicySet{}
		networkPolicyChoiceInt.NetworkPolicySet = &ves_io_schema_views.ObjectRefType{}
		updateSpec.NetworkPolicyChoice = networkPolicyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				networkPolicyChoiceInt.NetworkPolicySet.Name = v.(string)
			}

			if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

				networkPolicyChoiceInt.NetworkPolicySet.Namespace = v.(string)
			}

			if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

				networkPolicyChoiceInt.NetworkPolicySet.Tenant = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra NetworkFirewall obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_network_firewall.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating NetworkFirewall: %s", err)
	}

	return resourceVolterraNetworkFirewallRead(d, meta)
}

func resourceVolterraNetworkFirewallDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_network_firewall.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] NetworkFirewall %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra NetworkFirewall before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra NetworkFirewall obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_network_firewall.ObjectType, namespace, name)
}