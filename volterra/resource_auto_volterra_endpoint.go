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
	ves_io_schema_endpoint "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/endpoint"
)

// resourceVolterraEndpoint is implementation of Volterra's Endpoint resources
func resourceVolterraEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraEndpointCreate,
		Read:   resourceVolterraEndpointRead,
		Update: resourceVolterraEndpointUpdate,
		Delete: resourceVolterraEndpointDelete,

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

			"dns_name": {

				Type:     schema.TypeString,
				Optional: true,
			},

			"dns_name_advance": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"refresh_interval": {

							Type:     schema.TypeInt,
							Optional: true,
						},

						"strict_ttl": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"ip": {

				Type:     schema.TypeString,
				Optional: true,
			},

			"service_info": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"discovery_type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"service_name": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"service_selector": {

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
					},
				},
			},

			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"where": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"site": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

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

						"virtual_network": {

							Type:     schema.TypeSet,
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

						"virtual_site": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"network_type": {
										Type:     schema.TypeString,
										Optional: true,
									},

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
					},
				},
			},
		},
	}
}

// resourceVolterraEndpointCreate creates Endpoint resource
func resourceVolterraEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_endpoint.CreateSpecType{}
	createReq := &ves_io_schema_endpoint.CreateRequest{
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

	endpointAddressTypeFound := false

	if v, ok := d.GetOk("dns_name"); ok && !endpointAddressTypeFound {

		endpointAddressTypeFound = true
		endpointAddressInt := &ves_io_schema_endpoint.CreateSpecType_DnsName{}

		createSpec.EndpointAddress = endpointAddressInt

		endpointAddressInt.DnsName = v.(string)

	}

	if v, ok := d.GetOk("dns_name_advance"); ok && !endpointAddressTypeFound {

		endpointAddressTypeFound = true
		endpointAddressInt := &ves_io_schema_endpoint.CreateSpecType_DnsNameAdvance{}
		endpointAddressInt.DnsNameAdvance = &ves_io_schema_endpoint.DnsNameAdvanceType{}
		createSpec.EndpointAddress = endpointAddressInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				endpointAddressInt.DnsNameAdvance.Name = v.(string)
			}

			ttlChoiceTypeFound := false

			if v, ok := cs["refresh_interval"]; ok && !isIntfNil(v) && !ttlChoiceTypeFound {

				ttlChoiceTypeFound = true
				ttlChoiceInt := &ves_io_schema_endpoint.DnsNameAdvanceType_RefreshInterval{}

				endpointAddressInt.DnsNameAdvance.TtlChoice = ttlChoiceInt

				ttlChoiceInt.RefreshInterval =
					uint32(v.(int))

			}

			if v, ok := cs["strict_ttl"]; ok && !isIntfNil(v) && !ttlChoiceTypeFound {

				ttlChoiceTypeFound = true

				if v.(bool) {
					ttlChoiceInt := &ves_io_schema_endpoint.DnsNameAdvanceType_StrictTtl{}
					ttlChoiceInt.StrictTtl = &ves_io_schema.Empty{}
					endpointAddressInt.DnsNameAdvance.TtlChoice = ttlChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("ip"); ok && !endpointAddressTypeFound {

		endpointAddressTypeFound = true
		endpointAddressInt := &ves_io_schema_endpoint.CreateSpecType_Ip{}

		createSpec.EndpointAddress = endpointAddressInt

		endpointAddressInt.Ip = v.(string)

	}

	if v, ok := d.GetOk("service_info"); ok && !endpointAddressTypeFound {

		endpointAddressTypeFound = true
		endpointAddressInt := &ves_io_schema_endpoint.CreateSpecType_ServiceInfo{}
		endpointAddressInt.ServiceInfo = &ves_io_schema_endpoint.ServiceInfoType{}
		createSpec.EndpointAddress = endpointAddressInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["discovery_type"]; ok && !isIntfNil(v) {

				endpointAddressInt.ServiceInfo.DiscoveryType = ves_io_schema.DiscoveryType(ves_io_schema.DiscoveryType_value[v.(string)])

			}

			serviceInfoTypeFound := false

			if v, ok := cs["service_name"]; ok && !isIntfNil(v) && !serviceInfoTypeFound {

				serviceInfoTypeFound = true
				serviceInfoInt := &ves_io_schema_endpoint.ServiceInfoType_ServiceName{}

				endpointAddressInt.ServiceInfo.ServiceInfo = serviceInfoInt

				serviceInfoInt.ServiceName = v.(string)

			}

			if v, ok := cs["service_selector"]; ok && !isIntfNil(v) && !serviceInfoTypeFound {

				serviceInfoTypeFound = true
				serviceInfoInt := &ves_io_schema_endpoint.ServiceInfoType_ServiceSelector{}
				serviceInfoInt.ServiceSelector = &ves_io_schema.LabelSelectorType{}
				endpointAddressInt.ServiceInfo.ServiceInfo = serviceInfoInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						serviceInfoInt.ServiceSelector.Expressions = ls

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("port"); ok && !isIntfNil(v) {

		createSpec.Port =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("protocol"); ok && !isIntfNil(v) {

		createSpec.Protocol =
			v.(string)
	}

	if v, ok := d.GetOk("where"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		where := &ves_io_schema.NetworkSiteRefSelector{}
		createSpec.Where = where
		for _, set := range sl {

			whereMapStrToI := set.(map[string]interface{})

			refOrSelectorTypeFound := false

			if v, ok := whereMapStrToI["site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_Site{}
				refOrSelectorInt.Site = &ves_io_schema.SiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.Site.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.Site.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_network"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualNetwork{}
				refOrSelectorInt.VirtualNetwork = &ves_io_schema.NetworkRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualNetwork.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_network"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualSite{}
				refOrSelectorInt.VirtualSite = &ves_io_schema.VSiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.VirtualSite.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualSite.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra Endpoint object with struct: %+v", createReq)

	createEndpointResp, err := client.CreateObject(context.Background(), ves_io_schema_endpoint.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Endpoint: %s", err)
	}
	d.SetId(createEndpointResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraEndpointRead(d, meta)
}

func resourceVolterraEndpointRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_endpoint.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Endpoint %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Endpoint %q: %s", d.Id(), err)
	}
	return setEndpointFields(client, d, resp)
}

func setEndpointFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraEndpointUpdate updates Endpoint resource
func resourceVolterraEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_endpoint.ReplaceSpecType{}
	updateReq := &ves_io_schema_endpoint.ReplaceRequest{
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

	endpointAddressTypeFound := false

	if v, ok := d.GetOk("dns_name"); ok && !endpointAddressTypeFound {

		endpointAddressTypeFound = true
		endpointAddressInt := &ves_io_schema_endpoint.ReplaceSpecType_DnsName{}

		updateSpec.EndpointAddress = endpointAddressInt

		endpointAddressInt.DnsName = v.(string)

	}

	if v, ok := d.GetOk("dns_name_advance"); ok && !endpointAddressTypeFound {

		endpointAddressTypeFound = true
		endpointAddressInt := &ves_io_schema_endpoint.ReplaceSpecType_DnsNameAdvance{}
		endpointAddressInt.DnsNameAdvance = &ves_io_schema_endpoint.DnsNameAdvanceType{}
		updateSpec.EndpointAddress = endpointAddressInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["name"]; ok && !isIntfNil(v) {

				endpointAddressInt.DnsNameAdvance.Name = v.(string)
			}

			ttlChoiceTypeFound := false

			if v, ok := cs["refresh_interval"]; ok && !isIntfNil(v) && !ttlChoiceTypeFound {

				ttlChoiceTypeFound = true
				ttlChoiceInt := &ves_io_schema_endpoint.DnsNameAdvanceType_RefreshInterval{}

				endpointAddressInt.DnsNameAdvance.TtlChoice = ttlChoiceInt

				ttlChoiceInt.RefreshInterval =
					uint32(v.(int))

			}

			if v, ok := cs["strict_ttl"]; ok && !isIntfNil(v) && !ttlChoiceTypeFound {

				ttlChoiceTypeFound = true

				if v.(bool) {
					ttlChoiceInt := &ves_io_schema_endpoint.DnsNameAdvanceType_StrictTtl{}
					ttlChoiceInt.StrictTtl = &ves_io_schema.Empty{}
					endpointAddressInt.DnsNameAdvance.TtlChoice = ttlChoiceInt
				}

			}

		}

	}

	if v, ok := d.GetOk("ip"); ok && !endpointAddressTypeFound {

		endpointAddressTypeFound = true
		endpointAddressInt := &ves_io_schema_endpoint.ReplaceSpecType_Ip{}

		updateSpec.EndpointAddress = endpointAddressInt

		endpointAddressInt.Ip = v.(string)

	}

	if v, ok := d.GetOk("service_info"); ok && !endpointAddressTypeFound {

		endpointAddressTypeFound = true
		endpointAddressInt := &ves_io_schema_endpoint.ReplaceSpecType_ServiceInfo{}
		endpointAddressInt.ServiceInfo = &ves_io_schema_endpoint.ServiceInfoType{}
		updateSpec.EndpointAddress = endpointAddressInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["discovery_type"]; ok && !isIntfNil(v) {

				endpointAddressInt.ServiceInfo.DiscoveryType = ves_io_schema.DiscoveryType(ves_io_schema.DiscoveryType_value[v.(string)])

			}

			serviceInfoTypeFound := false

			if v, ok := cs["service_name"]; ok && !isIntfNil(v) && !serviceInfoTypeFound {

				serviceInfoTypeFound = true
				serviceInfoInt := &ves_io_schema_endpoint.ServiceInfoType_ServiceName{}

				endpointAddressInt.ServiceInfo.ServiceInfo = serviceInfoInt

				serviceInfoInt.ServiceName = v.(string)

			}

			if v, ok := cs["service_selector"]; ok && !isIntfNil(v) && !serviceInfoTypeFound {

				serviceInfoTypeFound = true
				serviceInfoInt := &ves_io_schema_endpoint.ServiceInfoType_ServiceSelector{}
				serviceInfoInt.ServiceSelector = &ves_io_schema.LabelSelectorType{}
				endpointAddressInt.ServiceInfo.ServiceInfo = serviceInfoInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						serviceInfoInt.ServiceSelector.Expressions = ls

					}

				}

			}

		}

	}

	if v, ok := d.GetOk("port"); ok && !isIntfNil(v) {

		updateSpec.Port =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("protocol"); ok && !isIntfNil(v) {

		updateSpec.Protocol =
			v.(string)
	}

	if v, ok := d.GetOk("where"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		where := &ves_io_schema.NetworkSiteRefSelector{}
		updateSpec.Where = where
		for _, set := range sl {

			whereMapStrToI := set.(map[string]interface{})

			refOrSelectorTypeFound := false

			if v, ok := whereMapStrToI["site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_Site{}
				refOrSelectorInt.Site = &ves_io_schema.SiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.Site.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.Site.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_network"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualNetwork{}
				refOrSelectorInt.VirtualNetwork = &ves_io_schema.NetworkRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualNetwork.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_network"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

			if v, ok := whereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.NetworkSiteRefSelector_VirtualSite{}
				refOrSelectorInt.VirtualSite = &ves_io_schema.VSiteRefType{}
				where.RefOrSelector = refOrSelectorInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["network_type"]; ok && !isIntfNil(v) {

						refOrSelectorInt.VirtualSite.NetworkType = ves_io_schema.VirtualNetworkType(ves_io_schema.VirtualNetworkType_value[v.(string)])

					}

					if v, ok := cs["ref"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						refIntNew := make([]*ves_io_schema.ObjectRefType, len(sl))
						refOrSelectorInt.VirtualSite.Ref = refIntNew
						for i, ps := range sl {

							rMapToStrVal := ps.(map[string]interface{})
							refIntNew[i] = &ves_io_schema.ObjectRefType{}

							refIntNew[i].Kind = "virtual_site"

							if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
								refIntNew[i].Name = v.(string)
							}

							if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								refIntNew[i].Namespace = v.(string)
							}

							if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								refIntNew[i].Tenant = v.(string)
							}

							if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
								refIntNew[i].Uid = v.(string)
							}

						}

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra Endpoint obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_endpoint.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Endpoint: %s", err)
	}

	return resourceVolterraEndpointRead(d, meta)
}

func resourceVolterraEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_endpoint.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Endpoint %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Endpoint before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Endpoint obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_endpoint.ObjectType, namespace, name)
}
