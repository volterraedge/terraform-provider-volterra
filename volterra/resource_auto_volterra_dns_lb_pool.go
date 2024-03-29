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
	ves_io_schema_dns_lb_pool "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dns_lb_pool"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraDnsLbPool is implementation of Volterra's DnsLbPool resources
func resourceVolterraDnsLbPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraDnsLbPoolCreate,
		Read:   resourceVolterraDnsLbPoolRead,
		Update: resourceVolterraDnsLbPoolUpdate,
		Delete: resourceVolterraDnsLbPoolDelete,

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

			"load_balancing_mode": {
				Type:     schema.TypeString,
				Required: true,
			},

			"a_pool": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"disable_health_check": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"health_check": {

							Type:     schema.TypeSet,
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

						"max_answers": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"members": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"ip_endpoint": {
										Type:     schema.TypeString,
										Required: true,
									},

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"ratio": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"aaaa_pool": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"max_answers": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"members": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"disable": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"ip_endpoint": {
										Type:     schema.TypeString,
										Required: true,
									},

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"ratio": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"cname_pool": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"members": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"domain": {
										Type:     schema.TypeString,
										Required: true,
									},

									"final_translation": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"ratio": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"mx_pool": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"max_answers": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"members": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"domain": {
										Type:     schema.TypeString,
										Required: true,
									},

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"priority": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"ratio": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"srv_pool": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"max_answers": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"members": {

							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"final_translation": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"port": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"priority": {
										Type:     schema.TypeInt,
										Required: true,
									},

									"ratio": {
										Type:     schema.TypeInt,
										Optional: true,
									},

									"target": {
										Type:     schema.TypeString,
										Required: true,
									},

									"weight": {
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"ttl": {

				Type:     schema.TypeInt,
				Optional: true,
			},

			"use_rrset_ttl": {

				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

// resourceVolterraDnsLbPoolCreate creates DnsLbPool resource
func resourceVolterraDnsLbPoolCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_dns_lb_pool.CreateSpecType{}
	createReq := &ves_io_schema_dns_lb_pool.CreateRequest{
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

	//load_balancing_mode
	if v, ok := d.GetOk("load_balancing_mode"); ok && !isIntfNil(v) {

		createSpec.LoadBalancingMode = ves_io_schema_dns_lb_pool.LoadBalancingMode(ves_io_schema_dns_lb_pool.LoadBalancingMode_value[v.(string)])

	}

	//pool_type_choice

	poolTypeChoiceTypeFound := false

	if v, ok := d.GetOk("a_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.CreateSpecType_APool{}
		poolTypeChoiceInt.APool = &ves_io_schema_dns_lb_pool.APool{}
		createSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			healthCheckChoiceTypeFound := false

			if v, ok := cs["disable_health_check"]; ok && !isIntfNil(v) && !healthCheckChoiceTypeFound {

				healthCheckChoiceTypeFound = true

				if v.(bool) {
					healthCheckChoiceInt := &ves_io_schema_dns_lb_pool.APool_DisableHealthCheck{}
					healthCheckChoiceInt.DisableHealthCheck = &ves_io_schema.Empty{}
					poolTypeChoiceInt.APool.HealthCheckChoice = healthCheckChoiceInt
				}

			}

			if v, ok := cs["health_check"]; ok && !isIntfNil(v) && !healthCheckChoiceTypeFound {

				healthCheckChoiceTypeFound = true
				healthCheckChoiceInt := &ves_io_schema_dns_lb_pool.APool_HealthCheck{}
				healthCheckChoiceInt.HealthCheck = &ves_io_schema_views.ObjectRefType{}
				poolTypeChoiceInt.APool.HealthCheckChoice = healthCheckChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						healthCheckChoiceInt.HealthCheck.Name = v.(string)

					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						healthCheckChoiceInt.HealthCheck.Namespace = v.(string)

					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						healthCheckChoiceInt.HealthCheck.Tenant = v.(string)

					}

				}

			}

			if v, ok := cs["max_answers"]; ok && !isIntfNil(v) {

				poolTypeChoiceInt.APool.MaxAnswers = uint32(v.(int))

			}

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.AddressMember, len(sl))
				poolTypeChoiceInt.APool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.AddressMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["disable"]; ok && !isIntfNil(w) {
						members[i].Disable = w.(bool)
					}

					if w, ok := membersMapStrToI["ip_endpoint"]; ok && !isIntfNil(w) {
						members[i].IpEndpoint = w.(string)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["priority"]; ok && !isIntfNil(w) {
						members[i].Priority = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("aaaa_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.CreateSpecType_AaaaPool{}
		poolTypeChoiceInt.AaaaPool = &ves_io_schema_dns_lb_pool.AAAAPool{}
		createSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["max_answers"]; ok && !isIntfNil(v) {

				poolTypeChoiceInt.AaaaPool.MaxAnswers = uint32(v.(int))

			}

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.AddressMember, len(sl))
				poolTypeChoiceInt.AaaaPool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.AddressMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["disable"]; ok && !isIntfNil(w) {
						members[i].Disable = w.(bool)
					}

					if w, ok := membersMapStrToI["ip_endpoint"]; ok && !isIntfNil(w) {
						members[i].IpEndpoint = w.(string)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["priority"]; ok && !isIntfNil(w) {
						members[i].Priority = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("cname_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.CreateSpecType_CnamePool{}
		poolTypeChoiceInt.CnamePool = &ves_io_schema_dns_lb_pool.CNAMEPool{}
		createSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.CNAMEMember, len(sl))
				poolTypeChoiceInt.CnamePool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.CNAMEMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["domain"]; ok && !isIntfNil(w) {
						members[i].Domain = w.(string)
					}

					if w, ok := membersMapStrToI["final_translation"]; ok && !isIntfNil(w) {
						members[i].FinalTranslation = w.(bool)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("mx_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.CreateSpecType_MxPool{}
		poolTypeChoiceInt.MxPool = &ves_io_schema_dns_lb_pool.MXPool{}
		createSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["max_answers"]; ok && !isIntfNil(v) {

				poolTypeChoiceInt.MxPool.MaxAnswers = uint32(v.(int))

			}

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.MXMember, len(sl))
				poolTypeChoiceInt.MxPool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.MXMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["domain"]; ok && !isIntfNil(w) {
						members[i].Domain = w.(string)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["priority"]; ok && !isIntfNil(w) {
						members[i].Priority = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("srv_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.CreateSpecType_SrvPool{}
		poolTypeChoiceInt.SrvPool = &ves_io_schema_dns_lb_pool.SRVPool{}
		createSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["max_answers"]; ok && !isIntfNil(v) {

				poolTypeChoiceInt.SrvPool.MaxAnswers = uint32(v.(int))

			}

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.SRVMember, len(sl))
				poolTypeChoiceInt.SrvPool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.SRVMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["final_translation"]; ok && !isIntfNil(w) {
						members[i].FinalTranslation = w.(bool)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["port"]; ok && !isIntfNil(w) {
						members[i].Port = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["priority"]; ok && !isIntfNil(w) {
						members[i].Priority = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["target"]; ok && !isIntfNil(w) {
						members[i].Target = w.(string)
					}

					if w, ok := membersMapStrToI["weight"]; ok && !isIntfNil(w) {
						members[i].Weight = uint32(w.(int))
					}

				}

			}

		}

	}

	//ttl_choice

	ttlChoiceTypeFound := false

	if v, ok := d.GetOk("ttl"); ok && !ttlChoiceTypeFound {

		ttlChoiceTypeFound = true
		ttlChoiceInt := &ves_io_schema_dns_lb_pool.CreateSpecType_Ttl{}

		createSpec.TtlChoice = ttlChoiceInt

		ttlChoiceInt.Ttl = uint32(v.(int))

	}

	if v, ok := d.GetOk("use_rrset_ttl"); ok && !ttlChoiceTypeFound {

		ttlChoiceTypeFound = true

		if v.(bool) {
			ttlChoiceInt := &ves_io_schema_dns_lb_pool.CreateSpecType_UseRrsetTtl{}
			ttlChoiceInt.UseRrsetTtl = &ves_io_schema.Empty{}
			createSpec.TtlChoice = ttlChoiceInt
		}

	}

	log.Printf("[DEBUG] Creating Volterra DnsLbPool object with struct: %+v", createReq)

	createDnsLbPoolResp, err := client.CreateObject(context.Background(), ves_io_schema_dns_lb_pool.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating DnsLbPool: %s", err)
	}
	d.SetId(createDnsLbPoolResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraDnsLbPoolRead(d, meta)
}

func resourceVolterraDnsLbPoolRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_dns_lb_pool.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] DnsLbPool %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra DnsLbPool %q: %s", d.Id(), err)
	}
	return setDnsLbPoolFields(client, d, resp)
}

func setDnsLbPoolFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraDnsLbPoolUpdate updates DnsLbPool resource
func resourceVolterraDnsLbPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_dns_lb_pool.ReplaceSpecType{}
	updateReq := &ves_io_schema_dns_lb_pool.ReplaceRequest{
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

	if v, ok := d.GetOk("load_balancing_mode"); ok && !isIntfNil(v) {

		updateSpec.LoadBalancingMode = ves_io_schema_dns_lb_pool.LoadBalancingMode(ves_io_schema_dns_lb_pool.LoadBalancingMode_value[v.(string)])

	}

	poolTypeChoiceTypeFound := false

	if v, ok := d.GetOk("a_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.ReplaceSpecType_APool{}
		poolTypeChoiceInt.APool = &ves_io_schema_dns_lb_pool.APool{}
		updateSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			healthCheckChoiceTypeFound := false

			if v, ok := cs["disable_health_check"]; ok && !isIntfNil(v) && !healthCheckChoiceTypeFound {

				healthCheckChoiceTypeFound = true

				if v.(bool) {
					healthCheckChoiceInt := &ves_io_schema_dns_lb_pool.APool_DisableHealthCheck{}
					healthCheckChoiceInt.DisableHealthCheck = &ves_io_schema.Empty{}
					poolTypeChoiceInt.APool.HealthCheckChoice = healthCheckChoiceInt
				}

			}

			if v, ok := cs["health_check"]; ok && !isIntfNil(v) && !healthCheckChoiceTypeFound {

				healthCheckChoiceTypeFound = true
				healthCheckChoiceInt := &ves_io_schema_dns_lb_pool.APool_HealthCheck{}
				healthCheckChoiceInt.HealthCheck = &ves_io_schema_views.ObjectRefType{}
				poolTypeChoiceInt.APool.HealthCheckChoice = healthCheckChoiceInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						healthCheckChoiceInt.HealthCheck.Name = v.(string)

					}

					if v, ok := cs["namespace"]; ok && !isIntfNil(v) {

						healthCheckChoiceInt.HealthCheck.Namespace = v.(string)

					}

					if v, ok := cs["tenant"]; ok && !isIntfNil(v) {

						healthCheckChoiceInt.HealthCheck.Tenant = v.(string)

					}

				}

			}

			if v, ok := cs["max_answers"]; ok && !isIntfNil(v) {

				poolTypeChoiceInt.APool.MaxAnswers = uint32(v.(int))

			}

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.AddressMember, len(sl))
				poolTypeChoiceInt.APool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.AddressMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["disable"]; ok && !isIntfNil(w) {
						members[i].Disable = w.(bool)
					}

					if w, ok := membersMapStrToI["ip_endpoint"]; ok && !isIntfNil(w) {
						members[i].IpEndpoint = w.(string)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["priority"]; ok && !isIntfNil(w) {
						members[i].Priority = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("aaaa_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.ReplaceSpecType_AaaaPool{}
		poolTypeChoiceInt.AaaaPool = &ves_io_schema_dns_lb_pool.AAAAPool{}
		updateSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["max_answers"]; ok && !isIntfNil(v) {

				poolTypeChoiceInt.AaaaPool.MaxAnswers = uint32(v.(int))

			}

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.AddressMember, len(sl))
				poolTypeChoiceInt.AaaaPool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.AddressMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["disable"]; ok && !isIntfNil(w) {
						members[i].Disable = w.(bool)
					}

					if w, ok := membersMapStrToI["ip_endpoint"]; ok && !isIntfNil(w) {
						members[i].IpEndpoint = w.(string)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["priority"]; ok && !isIntfNil(w) {
						members[i].Priority = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("cname_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.ReplaceSpecType_CnamePool{}
		poolTypeChoiceInt.CnamePool = &ves_io_schema_dns_lb_pool.CNAMEPool{}
		updateSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.CNAMEMember, len(sl))
				poolTypeChoiceInt.CnamePool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.CNAMEMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["domain"]; ok && !isIntfNil(w) {
						members[i].Domain = w.(string)
					}

					if w, ok := membersMapStrToI["final_translation"]; ok && !isIntfNil(w) {
						members[i].FinalTranslation = w.(bool)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("mx_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.ReplaceSpecType_MxPool{}
		poolTypeChoiceInt.MxPool = &ves_io_schema_dns_lb_pool.MXPool{}
		updateSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["max_answers"]; ok && !isIntfNil(v) {

				poolTypeChoiceInt.MxPool.MaxAnswers = uint32(v.(int))

			}

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.MXMember, len(sl))
				poolTypeChoiceInt.MxPool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.MXMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["domain"]; ok && !isIntfNil(w) {
						members[i].Domain = w.(string)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["priority"]; ok && !isIntfNil(w) {
						members[i].Priority = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("srv_pool"); ok && !poolTypeChoiceTypeFound {

		poolTypeChoiceTypeFound = true
		poolTypeChoiceInt := &ves_io_schema_dns_lb_pool.ReplaceSpecType_SrvPool{}
		poolTypeChoiceInt.SrvPool = &ves_io_schema_dns_lb_pool.SRVPool{}
		updateSpec.PoolTypeChoice = poolTypeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["max_answers"]; ok && !isIntfNil(v) {

				poolTypeChoiceInt.SrvPool.MaxAnswers = uint32(v.(int))

			}

			if v, ok := cs["members"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				members := make([]*ves_io_schema_dns_lb_pool.SRVMember, len(sl))
				poolTypeChoiceInt.SrvPool.Members = members
				for i, set := range sl {
					members[i] = &ves_io_schema_dns_lb_pool.SRVMember{}
					membersMapStrToI := set.(map[string]interface{})

					if w, ok := membersMapStrToI["final_translation"]; ok && !isIntfNil(w) {
						members[i].FinalTranslation = w.(bool)
					}

					if w, ok := membersMapStrToI["name"]; ok && !isIntfNil(w) {
						members[i].Name = w.(string)
					}

					if w, ok := membersMapStrToI["port"]; ok && !isIntfNil(w) {
						members[i].Port = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["priority"]; ok && !isIntfNil(w) {
						members[i].Priority = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["ratio"]; ok && !isIntfNil(w) {
						members[i].Ratio = uint32(w.(int))
					}

					if w, ok := membersMapStrToI["target"]; ok && !isIntfNil(w) {
						members[i].Target = w.(string)
					}

					if w, ok := membersMapStrToI["weight"]; ok && !isIntfNil(w) {
						members[i].Weight = uint32(w.(int))
					}

				}

			}

		}

	}

	ttlChoiceTypeFound := false

	if v, ok := d.GetOk("ttl"); ok && !ttlChoiceTypeFound {

		ttlChoiceTypeFound = true
		ttlChoiceInt := &ves_io_schema_dns_lb_pool.ReplaceSpecType_Ttl{}

		updateSpec.TtlChoice = ttlChoiceInt

		ttlChoiceInt.Ttl = uint32(v.(int))

	}

	if v, ok := d.GetOk("use_rrset_ttl"); ok && !ttlChoiceTypeFound {

		ttlChoiceTypeFound = true

		if v.(bool) {
			ttlChoiceInt := &ves_io_schema_dns_lb_pool.ReplaceSpecType_UseRrsetTtl{}
			ttlChoiceInt.UseRrsetTtl = &ves_io_schema.Empty{}
			updateSpec.TtlChoice = ttlChoiceInt
		}

	}

	log.Printf("[DEBUG] Updating Volterra DnsLbPool obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_dns_lb_pool.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating DnsLbPool: %s", err)
	}

	return resourceVolterraDnsLbPoolRead(d, meta)
}

func resourceVolterraDnsLbPoolDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_dns_lb_pool.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] DnsLbPool %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra DnsLbPool before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra DnsLbPool obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_dns_lb_pool.ObjectType, namespace, name)
}
