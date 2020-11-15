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
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_views_tcp_loadbalancer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/tcp_loadbalancer"
)

// resourceVolterraTcpLoadbalancer is implementation of Volterra's TcpLoadbalancer resources
func resourceVolterraTcpLoadbalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraTcpLoadbalancerCreate,
		Read:   resourceVolterraTcpLoadbalancerRead,
		Update: resourceVolterraTcpLoadbalancerUpdate,
		Delete: resourceVolterraTcpLoadbalancerDelete,

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

			"domains": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"listen_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"origin_pools_weights": {

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

			"with_sni": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

// resourceVolterraTcpLoadbalancerCreate creates TcpLoadbalancer resource
func resourceVolterraTcpLoadbalancerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_tcp_loadbalancer.CreateSpecType{}
	createReq := &ves_io_schema_views_tcp_loadbalancer.CreateRequest{
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

	advertiseChoiceTypeFound := false

	if v, ok := d.GetOk("advertise_custom"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true
		advertiseChoiceInt := &ves_io_schema_views_tcp_loadbalancer.CreateSpecType_AdvertiseCustom{}
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
		advertiseChoiceInt := &ves_io_schema_views_tcp_loadbalancer.CreateSpecType_AdvertiseOnPublic{}
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
			advertiseChoiceInt := &ves_io_schema_views_tcp_loadbalancer.CreateSpecType_AdvertiseOnPublicDefaultVip{}
			advertiseChoiceInt.AdvertiseOnPublicDefaultVip = &ves_io_schema.Empty{}
			createSpec.AdvertiseChoice = advertiseChoiceInt
		}

	}

	if v, ok := d.GetOk("do_not_advertise"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true

		if v.(bool) {
			advertiseChoiceInt := &ves_io_schema_views_tcp_loadbalancer.CreateSpecType_DoNotAdvertise{}
			advertiseChoiceInt.DoNotAdvertise = &ves_io_schema.Empty{}
			createSpec.AdvertiseChoice = advertiseChoiceInt
		}

	}

	if v, ok := d.GetOk("domains"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		createSpec.Domains = ls

	}

	if v, ok := d.GetOk("listen_port"); ok && !isIntfNil(v) {

		createSpec.ListenPort =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("origin_pools_weights"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		originPoolsWeights := make([]*ves_io_schema_views.OriginPoolWithWeight, len(sl))
		createSpec.OriginPoolsWeights = originPoolsWeights
		for i, set := range sl {
			originPoolsWeights[i] = &ves_io_schema_views.OriginPoolWithWeight{}

			originPoolsWeightsMapStrToI := set.(map[string]interface{})

			poolChoiceTypeFound := false

			if v, ok := originPoolsWeightsMapStrToI["cluster"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Cluster{}
				poolChoiceInt.Cluster = &ves_io_schema_views.ObjectRefType{}
				originPoolsWeights[i].PoolChoice = poolChoiceInt

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

			if v, ok := originPoolsWeightsMapStrToI["pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Pool{}
				poolChoiceInt.Pool = &ves_io_schema_views.ObjectRefType{}
				originPoolsWeights[i].PoolChoice = poolChoiceInt

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

			if w, ok := originPoolsWeightsMapStrToI["weight"]; ok && !isIntfNil(w) {
				originPoolsWeights[i].Weight = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("with_sni"); ok && !isIntfNil(v) {

		createSpec.WithSni =
			v.(bool)
	}

	log.Printf("[DEBUG] Creating Volterra TcpLoadbalancer object with struct: %+v", createReq)

	createTcpLoadbalancerResp, err := client.CreateObject(context.Background(), ves_io_schema_views_tcp_loadbalancer.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating TcpLoadbalancer: %s", err)
	}
	d.SetId(createTcpLoadbalancerResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraTcpLoadbalancerRead(d, meta)
}

func resourceVolterraTcpLoadbalancerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_tcp_loadbalancer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] TcpLoadbalancer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra TcpLoadbalancer %q: %s", d.Id(), err)
	}
	return setTcpLoadbalancerFields(client, d, resp)
}

func setTcpLoadbalancerFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraTcpLoadbalancerUpdate updates TcpLoadbalancer resource
func resourceVolterraTcpLoadbalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_tcp_loadbalancer.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_tcp_loadbalancer.ReplaceRequest{
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

	advertiseChoiceTypeFound := false

	if v, ok := d.GetOk("advertise_custom"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true
		advertiseChoiceInt := &ves_io_schema_views_tcp_loadbalancer.ReplaceSpecType_AdvertiseCustom{}
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
		advertiseChoiceInt := &ves_io_schema_views_tcp_loadbalancer.ReplaceSpecType_AdvertiseOnPublic{}
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
			advertiseChoiceInt := &ves_io_schema_views_tcp_loadbalancer.ReplaceSpecType_AdvertiseOnPublicDefaultVip{}
			advertiseChoiceInt.AdvertiseOnPublicDefaultVip = &ves_io_schema.Empty{}
			updateSpec.AdvertiseChoice = advertiseChoiceInt
		}

	}

	if v, ok := d.GetOk("do_not_advertise"); ok && !advertiseChoiceTypeFound {

		advertiseChoiceTypeFound = true

		if v.(bool) {
			advertiseChoiceInt := &ves_io_schema_views_tcp_loadbalancer.ReplaceSpecType_DoNotAdvertise{}
			advertiseChoiceInt.DoNotAdvertise = &ves_io_schema.Empty{}
			updateSpec.AdvertiseChoice = advertiseChoiceInt
		}

	}

	if v, ok := d.GetOk("domains"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			ls[i] = v.(string)
		}
		updateSpec.Domains = ls

	}

	if v, ok := d.GetOk("listen_port"); ok && !isIntfNil(v) {

		updateSpec.ListenPort =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("origin_pools"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		originPoolsInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
		updateSpec.OriginPools = originPoolsInt
		for i, ps := range sl {

			opMapToStrVal := ps.(map[string]interface{})
			originPoolsInt[i] = &ves_io_schema_views.ObjectRefType{}

			if v, ok := opMapToStrVal["name"]; ok && !isIntfNil(v) {
				originPoolsInt[i].Name = v.(string)
			}

			if v, ok := opMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				originPoolsInt[i].Namespace = v.(string)
			}

			if v, ok := opMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				originPoolsInt[i].Tenant = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("origin_pools_weights"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		originPoolsWeights := make([]*ves_io_schema_views.OriginPoolWithWeight, len(sl))
		updateSpec.OriginPoolsWeights = originPoolsWeights
		for i, set := range sl {
			originPoolsWeights[i] = &ves_io_schema_views.OriginPoolWithWeight{}

			originPoolsWeightsMapStrToI := set.(map[string]interface{})

			poolChoiceTypeFound := false

			if v, ok := originPoolsWeightsMapStrToI["cluster"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Cluster{}
				poolChoiceInt.Cluster = &ves_io_schema_views.ObjectRefType{}
				originPoolsWeights[i].PoolChoice = poolChoiceInt

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

			if v, ok := originPoolsWeightsMapStrToI["pool"]; ok && !isIntfNil(v) && !poolChoiceTypeFound {

				poolChoiceTypeFound = true
				poolChoiceInt := &ves_io_schema_views.OriginPoolWithWeight_Pool{}
				poolChoiceInt.Pool = &ves_io_schema_views.ObjectRefType{}
				originPoolsWeights[i].PoolChoice = poolChoiceInt

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

			if w, ok := originPoolsWeightsMapStrToI["weight"]; ok && !isIntfNil(w) {
				originPoolsWeights[i].Weight = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("with_sni"); ok && !isIntfNil(v) {

		updateSpec.WithSni =
			v.(bool)
	}

	log.Printf("[DEBUG] Updating Volterra TcpLoadbalancer obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_tcp_loadbalancer.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating TcpLoadbalancer: %s", err)
	}

	return resourceVolterraTcpLoadbalancerRead(d, meta)
}

func resourceVolterraTcpLoadbalancerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_tcp_loadbalancer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] TcpLoadbalancer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra TcpLoadbalancer before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra TcpLoadbalancer obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_tcp_loadbalancer.ObjectType, namespace, name)
}