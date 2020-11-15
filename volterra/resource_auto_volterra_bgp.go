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
	ves_io_schema_bgp "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/bgp"
)

// resourceVolterraBgp is implementation of Volterra's Bgp resources
func resourceVolterraBgp() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraBgpCreate,
		Read:   resourceVolterraBgpRead,
		Update: resourceVolterraBgpUpdate,
		Delete: resourceVolterraBgpDelete,

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

			"bgp_parameters": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"asn": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"bgp_router_id": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipv4": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"addr": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"ipv6": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"addr": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"bgp_router_id_key": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"bgp_router_id_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"bgp_peers": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"asn": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"bgp_peer_address": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipv4": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"addr": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"ipv6": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"addr": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"bgp_peer_address_key": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"bgp_peer_address_type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"bgp_peer_subnet_offset": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"network_interface": {

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

// resourceVolterraBgpCreate creates Bgp resource
func resourceVolterraBgpCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_bgp.CreateSpecType{}
	createReq := &ves_io_schema_bgp.CreateRequest{
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

	if v, ok := d.GetOk("bgp_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		bgpParameters := &ves_io_schema_bgp.BgpParameters{}
		createSpec.BgpParameters = bgpParameters
		for _, set := range sl {

			bgpParametersMapStrToI := set.(map[string]interface{})

			if w, ok := bgpParametersMapStrToI["asn"]; ok && !isIntfNil(w) {
				bgpParameters.Asn = w.(uint32)
			}

			if v, ok := bgpParametersMapStrToI["bgp_router_id"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				bgpRouterId := &ves_io_schema.IpAddressType{}
				bgpParameters.BgpRouterId = bgpRouterId
				for _, set := range sl {

					bgpRouterIdMapStrToI := set.(map[string]interface{})

					verTypeFound := false

					if v, ok := bgpRouterIdMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

						verTypeFound = true
						verInt := &ves_io_schema.IpAddressType_Ipv4{}
						verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
						bgpRouterId.Ver = verInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv4.Addr = v.(string)
							}

						}

					}

					if v, ok := bgpRouterIdMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

						verTypeFound = true
						verInt := &ves_io_schema.IpAddressType_Ipv6{}
						verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
						bgpRouterId.Ver = verInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv6.Addr = v.(string)
							}

						}

					}

				}

			}

			if w, ok := bgpParametersMapStrToI["bgp_router_id_key"]; ok && !isIntfNil(w) {
				bgpParameters.BgpRouterIdKey = w.(string)
			}

			if v, ok := bgpParametersMapStrToI["bgp_router_id_type"]; ok && !isIntfNil(v) {

				bgpParameters.BgpRouterIdType = ves_io_schema_bgp.BgpRouterIdType(ves_io_schema_bgp.BgpRouterIdType_value[v.(string)])

			}

		}

	}

	if v, ok := d.GetOk("bgp_peers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		bgpPeers := make([]*ves_io_schema_bgp.BgpPeer, len(sl))
		createSpec.BgpPeers = bgpPeers
		for i, set := range sl {
			bgpPeers[i] = &ves_io_schema_bgp.BgpPeer{}

			bgpPeersMapStrToI := set.(map[string]interface{})

			if w, ok := bgpPeersMapStrToI["asn"]; ok && !isIntfNil(w) {
				bgpPeers[i].Asn = w.(uint32)
			}

			if v, ok := bgpPeersMapStrToI["bgp_peer_address"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				bgpPeerAddress := &ves_io_schema.IpAddressType{}
				bgpPeers[i].BgpPeerAddress = bgpPeerAddress
				for _, set := range sl {

					bgpPeerAddressMapStrToI := set.(map[string]interface{})

					verTypeFound := false

					if v, ok := bgpPeerAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

						verTypeFound = true
						verInt := &ves_io_schema.IpAddressType_Ipv4{}
						verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
						bgpPeerAddress.Ver = verInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv4.Addr = v.(string)
							}

						}

					}

					if v, ok := bgpPeerAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

						verTypeFound = true
						verInt := &ves_io_schema.IpAddressType_Ipv6{}
						verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
						bgpPeerAddress.Ver = verInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv6.Addr = v.(string)
							}

						}

					}

				}

			}

			if w, ok := bgpPeersMapStrToI["bgp_peer_address_key"]; ok && !isIntfNil(w) {
				bgpPeers[i].BgpPeerAddressKey = w.(string)
			}

			if v, ok := bgpPeersMapStrToI["bgp_peer_address_type"]; ok && !isIntfNil(v) {

				bgpPeers[i].BgpPeerAddressType = ves_io_schema_bgp.BgpPeerAddressType(ves_io_schema_bgp.BgpPeerAddressType_value[v.(string)])

			}

			if w, ok := bgpPeersMapStrToI["bgp_peer_subnet_offset"]; ok && !isIntfNil(w) {
				bgpPeers[i].BgpPeerSubnetOffset = w.(uint32)
			}

			if w, ok := bgpPeersMapStrToI["port"]; ok && !isIntfNil(w) {
				bgpPeers[i].Port = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("network_interface"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		networkInterfaceInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.NetworkInterface = networkInterfaceInt
		for i, ps := range sl {

			niMapToStrVal := ps.(map[string]interface{})
			networkInterfaceInt[i] = &ves_io_schema.ObjectRefType{}

			networkInterfaceInt[i].Kind = "network_interface"

			if v, ok := niMapToStrVal["name"]; ok && !isIntfNil(v) {
				networkInterfaceInt[i].Name = v.(string)
			}

			if v, ok := niMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				networkInterfaceInt[i].Namespace = v.(string)
			}

			if v, ok := niMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				networkInterfaceInt[i].Tenant = v.(string)
			}

			if v, ok := niMapToStrVal["uid"]; ok && !isIntfNil(v) {
				networkInterfaceInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("where"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		where := &ves_io_schema.SiteVirtualSiteRefSelector{}
		createSpec.Where = where
		for _, set := range sl {

			whereMapStrToI := set.(map[string]interface{})

			refOrSelectorTypeFound := false

			if v, ok := whereMapStrToI["site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.SiteVirtualSiteRefSelector_Site{}
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

			if v, ok := whereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.SiteVirtualSiteRefSelector_VirtualSite{}
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

	log.Printf("[DEBUG] Creating Volterra Bgp object with struct: %+v", createReq)

	createBgpResp, err := client.CreateObject(context.Background(), ves_io_schema_bgp.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Bgp: %s", err)
	}
	d.SetId(createBgpResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraBgpRead(d, meta)
}

func resourceVolterraBgpRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_bgp.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Bgp %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Bgp %q: %s", d.Id(), err)
	}
	return setBgpFields(client, d, resp)
}

func setBgpFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraBgpUpdate updates Bgp resource
func resourceVolterraBgpUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_bgp.ReplaceSpecType{}
	updateReq := &ves_io_schema_bgp.ReplaceRequest{
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

	if v, ok := d.GetOk("bgp_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		bgpParameters := &ves_io_schema_bgp.BgpParameters{}
		updateSpec.BgpParameters = bgpParameters
		for _, set := range sl {

			bgpParametersMapStrToI := set.(map[string]interface{})

			if w, ok := bgpParametersMapStrToI["asn"]; ok && !isIntfNil(w) {
				bgpParameters.Asn = w.(uint32)
			}

			if v, ok := bgpParametersMapStrToI["bgp_router_id"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				bgpRouterId := &ves_io_schema.IpAddressType{}
				bgpParameters.BgpRouterId = bgpRouterId
				for _, set := range sl {

					bgpRouterIdMapStrToI := set.(map[string]interface{})

					verTypeFound := false

					if v, ok := bgpRouterIdMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

						verTypeFound = true
						verInt := &ves_io_schema.IpAddressType_Ipv4{}
						verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
						bgpRouterId.Ver = verInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv4.Addr = v.(string)
							}

						}

					}

					if v, ok := bgpRouterIdMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

						verTypeFound = true
						verInt := &ves_io_schema.IpAddressType_Ipv6{}
						verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
						bgpRouterId.Ver = verInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv6.Addr = v.(string)
							}

						}

					}

				}

			}

			if w, ok := bgpParametersMapStrToI["bgp_router_id_key"]; ok && !isIntfNil(w) {
				bgpParameters.BgpRouterIdKey = w.(string)
			}

			if v, ok := bgpParametersMapStrToI["bgp_router_id_type"]; ok && !isIntfNil(v) {

				bgpParameters.BgpRouterIdType = ves_io_schema_bgp.BgpRouterIdType(ves_io_schema_bgp.BgpRouterIdType_value[v.(string)])

			}

		}

	}

	if v, ok := d.GetOk("bgp_peers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		bgpPeers := make([]*ves_io_schema_bgp.BgpPeer, len(sl))
		updateSpec.BgpPeers = bgpPeers
		for i, set := range sl {
			bgpPeers[i] = &ves_io_schema_bgp.BgpPeer{}

			bgpPeersMapStrToI := set.(map[string]interface{})

			if w, ok := bgpPeersMapStrToI["asn"]; ok && !isIntfNil(w) {
				bgpPeers[i].Asn = w.(uint32)
			}

			if v, ok := bgpPeersMapStrToI["bgp_peer_address"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				bgpPeerAddress := &ves_io_schema.IpAddressType{}
				bgpPeers[i].BgpPeerAddress = bgpPeerAddress
				for _, set := range sl {

					bgpPeerAddressMapStrToI := set.(map[string]interface{})

					verTypeFound := false

					if v, ok := bgpPeerAddressMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

						verTypeFound = true
						verInt := &ves_io_schema.IpAddressType_Ipv4{}
						verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
						bgpPeerAddress.Ver = verInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv4.Addr = v.(string)
							}

						}

					}

					if v, ok := bgpPeerAddressMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

						verTypeFound = true
						verInt := &ves_io_schema.IpAddressType_Ipv6{}
						verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
						bgpPeerAddress.Ver = verInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv6.Addr = v.(string)
							}

						}

					}

				}

			}

			if w, ok := bgpPeersMapStrToI["bgp_peer_address_key"]; ok && !isIntfNil(w) {
				bgpPeers[i].BgpPeerAddressKey = w.(string)
			}

			if v, ok := bgpPeersMapStrToI["bgp_peer_address_type"]; ok && !isIntfNil(v) {

				bgpPeers[i].BgpPeerAddressType = ves_io_schema_bgp.BgpPeerAddressType(ves_io_schema_bgp.BgpPeerAddressType_value[v.(string)])

			}

			if w, ok := bgpPeersMapStrToI["bgp_peer_subnet_offset"]; ok && !isIntfNil(w) {
				bgpPeers[i].BgpPeerSubnetOffset = w.(uint32)
			}

			if w, ok := bgpPeersMapStrToI["port"]; ok && !isIntfNil(w) {
				bgpPeers[i].Port = w.(uint32)
			}

		}

	}

	if v, ok := d.GetOk("network_interface"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		networkInterfaceInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.NetworkInterface = networkInterfaceInt
		for i, ps := range sl {

			niMapToStrVal := ps.(map[string]interface{})
			networkInterfaceInt[i] = &ves_io_schema.ObjectRefType{}

			networkInterfaceInt[i].Kind = "network_interface"

			if v, ok := niMapToStrVal["name"]; ok && !isIntfNil(v) {
				networkInterfaceInt[i].Name = v.(string)
			}

			if v, ok := niMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				networkInterfaceInt[i].Namespace = v.(string)
			}

			if v, ok := niMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				networkInterfaceInt[i].Tenant = v.(string)
			}

			if v, ok := niMapToStrVal["uid"]; ok && !isIntfNil(v) {
				networkInterfaceInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("where"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		where := &ves_io_schema.SiteVirtualSiteRefSelector{}
		updateSpec.Where = where
		for _, set := range sl {

			whereMapStrToI := set.(map[string]interface{})

			refOrSelectorTypeFound := false

			if v, ok := whereMapStrToI["site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.SiteVirtualSiteRefSelector_Site{}
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

			if v, ok := whereMapStrToI["virtual_site"]; ok && !isIntfNil(v) && !refOrSelectorTypeFound {

				refOrSelectorTypeFound = true
				refOrSelectorInt := &ves_io_schema.SiteVirtualSiteRefSelector_VirtualSite{}
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

	log.Printf("[DEBUG] Updating Volterra Bgp obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_bgp.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Bgp: %s", err)
	}

	return resourceVolterraBgpRead(d, meta)
}

func resourceVolterraBgpDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_bgp.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Bgp %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Bgp before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Bgp obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_bgp.ObjectType, namespace, name)
}