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
	ves_io_schema_uztna_application_uztna_application "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/uztna/application/uztna_application"
)

// resourceVolterraUztnaApplication is implementation of Volterra's UztnaApplication resources
func resourceVolterraUztnaApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraUztnaApplicationCreate,
		Read:   resourceVolterraUztnaApplicationRead,
		Update: resourceVolterraUztnaApplicationUpdate,
		Delete: resourceVolterraUztnaApplicationDelete,

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

			"app_tags": {

				Type: schema.TypeList,

				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"application_vip": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ipv4": {

							Type:     schema.TypeList,
							MaxItems: 1,
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

							Type:     schema.TypeList,
							MaxItems: 1,
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

			"extended_app_tags": {

				Type: schema.TypeList,

				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"location": {

				Type:     schema.TypeList,
				Required: true,
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

			"origin": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"origin_pool": {

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

						"snat_pool": {

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

			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"protocol": {

				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"http": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"https": {

							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"certificate": {

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

						"tcp": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"udp": {

							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"url": {
				Type:     schema.TypeString,
				Required: true,
			},

			"uztna_domain_ref": {

				Type:     schema.TypeList,
				Required: true,
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

// resourceVolterraUztnaApplicationCreate creates UztnaApplication resource
func resourceVolterraUztnaApplicationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_uztna_application_uztna_application.CreateSpecType{}
	createReq := &ves_io_schema_uztna_application_uztna_application.CreateRequest{
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

	//app_tags
	if v, ok := d.GetOk("app_tags"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			if v == nil {
				return fmt.Errorf("please provide valid non-empty string value of field app_tags")
			}
			if str, ok := v.(string); ok {
				ls[i] = str
			}
		}
		createSpec.AppTags = ls

	}

	//application_vip
	if v, ok := d.GetOk("application_vip"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		applicationVip := &ves_io_schema.IpAddressType{}
		createSpec.ApplicationVip = applicationVip
		for _, set := range sl {
			if set != nil {
				applicationVipMapStrToI := set.(map[string]interface{})

				verTypeFound := false

				if v, ok := applicationVipMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

					verTypeFound = true
					verInt := &ves_io_schema.IpAddressType_Ipv4{}
					verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
					applicationVip.Ver = verInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv4.Addr = v.(string)

							}

						}
					}

				}

				if v, ok := applicationVipMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

					verTypeFound = true
					verInt := &ves_io_schema.IpAddressType_Ipv6{}
					verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
					applicationVip.Ver = verInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv6.Addr = v.(string)

							}

						}
					}

				}

			}
		}

	}

	//extended_app_tags
	if v, ok := d.GetOk("extended_app_tags"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			if v == nil {
				return fmt.Errorf("please provide valid non-empty string value of field extended_app_tags")
			}
			if str, ok := v.(string); ok {
				ls[i] = str
			}
		}
		createSpec.ExtendedAppTags = ls

	}

	//location
	if v, ok := d.GetOk("location"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		locationInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.Location = locationInt
		for i, ps := range sl {

			lMapToStrVal := ps.(map[string]interface{})
			locationInt[i] = &ves_io_schema.ObjectRefType{}

			locationInt[i].Kind = "site"

			if v, ok := lMapToStrVal["name"]; ok && !isIntfNil(v) {
				locationInt[i].Name = v.(string)
			}

			if v, ok := lMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				locationInt[i].Namespace = v.(string)
			}

			if v, ok := lMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				locationInt[i].Tenant = v.(string)
			}

			if v, ok := lMapToStrVal["uid"]; ok && !isIntfNil(v) {
				locationInt[i].Uid = v.(string)
			}

		}

	}

	//origin
	if v, ok := d.GetOk("origin"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		origin := &ves_io_schema_uztna_application_uztna_application.Origin{}
		createSpec.Origin = origin
		for _, set := range sl {
			if set != nil {
				originMapStrToI := set.(map[string]interface{})

				if v, ok := originMapStrToI["origin_pool"]; ok && !isIntfNil(v) {

					sl := v.([]interface{})
					originPoolInt := make([]*ves_io_schema.ObjectRefType, len(sl))
					origin.OriginPool = originPoolInt
					for i, ps := range sl {

						opMapToStrVal := ps.(map[string]interface{})
						originPoolInt[i] = &ves_io_schema.ObjectRefType{}

						originPoolInt[i].Kind = "uztna_origin_pool"

						if v, ok := opMapToStrVal["name"]; ok && !isIntfNil(v) {
							originPoolInt[i].Name = v.(string)
						}

						if v, ok := opMapToStrVal["namespace"]; ok && !isIntfNil(v) {
							originPoolInt[i].Namespace = v.(string)
						}

						if v, ok := opMapToStrVal["tenant"]; ok && !isIntfNil(v) {
							originPoolInt[i].Tenant = v.(string)
						}

						if v, ok := opMapToStrVal["uid"]; ok && !isIntfNil(v) {
							originPoolInt[i].Uid = v.(string)
						}

					}

				}

				if v, ok := originMapStrToI["snat_pool"]; ok && !isIntfNil(v) {

					sl := v.([]interface{})
					snatPoolInt := make([]*ves_io_schema.ObjectRefType, len(sl))
					origin.SnatPool = snatPoolInt
					for i, ps := range sl {

						spMapToStrVal := ps.(map[string]interface{})
						snatPoolInt[i] = &ves_io_schema.ObjectRefType{}

						snatPoolInt[i].Kind = "uztna_snat_pool"

						if v, ok := spMapToStrVal["name"]; ok && !isIntfNil(v) {
							snatPoolInt[i].Name = v.(string)
						}

						if v, ok := spMapToStrVal["namespace"]; ok && !isIntfNil(v) {
							snatPoolInt[i].Namespace = v.(string)
						}

						if v, ok := spMapToStrVal["tenant"]; ok && !isIntfNil(v) {
							snatPoolInt[i].Tenant = v.(string)
						}

						if v, ok := spMapToStrVal["uid"]; ok && !isIntfNil(v) {
							snatPoolInt[i].Uid = v.(string)
						}

					}

				}

			}
		}

	}

	//port
	if v, ok := d.GetOk("port"); ok && !isIntfNil(v) {

		createSpec.Port =
			uint32(v.(int))

	}

	//protocol
	if v, ok := d.GetOk("protocol"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		protocol := &ves_io_schema_uztna_application_uztna_application.Protocol{}
		createSpec.Protocol = protocol
		for _, set := range sl {
			if set != nil {
				protocolMapStrToI := set.(map[string]interface{})

				protocolChoiceTypeFound := false

				if v, ok := protocolMapStrToI["http"]; ok && !isIntfNil(v) && !protocolChoiceTypeFound {

					protocolChoiceTypeFound = true

					if v.(bool) {
						protocolChoiceInt := &ves_io_schema_uztna_application_uztna_application.Protocol_HTTP{}
						protocolChoiceInt.HTTP = &ves_io_schema.Empty{}
						protocol.ProtocolChoice = protocolChoiceInt
					}

				}

				if v, ok := protocolMapStrToI["https"]; ok && !isIntfNil(v) && !protocolChoiceTypeFound {

					protocolChoiceTypeFound = true
					protocolChoiceInt := &ves_io_schema_uztna_application_uztna_application.Protocol_HTTPS{}
					protocolChoiceInt.HTTPS = &ves_io_schema_uztna_application_uztna_application.AppCertificate{}
					protocol.ProtocolChoice = protocolChoiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["certificate"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								certificateInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								protocolChoiceInt.HTTPS.Certificate = certificateInt
								for i, ps := range sl {

									cMapToStrVal := ps.(map[string]interface{})
									certificateInt[i] = &ves_io_schema.ObjectRefType{}

									certificateInt[i].Kind = "certificate"

									if v, ok := cMapToStrVal["name"]; ok && !isIntfNil(v) {
										certificateInt[i].Name = v.(string)
									}

									if v, ok := cMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										certificateInt[i].Namespace = v.(string)
									}

									if v, ok := cMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										certificateInt[i].Tenant = v.(string)
									}

									if v, ok := cMapToStrVal["uid"]; ok && !isIntfNil(v) {
										certificateInt[i].Uid = v.(string)
									}

								}

							}

						}
					}

				}

				if v, ok := protocolMapStrToI["tcp"]; ok && !isIntfNil(v) && !protocolChoiceTypeFound {

					protocolChoiceTypeFound = true

					if v.(bool) {
						protocolChoiceInt := &ves_io_schema_uztna_application_uztna_application.Protocol_TCP{}
						protocolChoiceInt.TCP = &ves_io_schema.Empty{}
						protocol.ProtocolChoice = protocolChoiceInt
					}

				}

				if v, ok := protocolMapStrToI["udp"]; ok && !isIntfNil(v) && !protocolChoiceTypeFound {

					protocolChoiceTypeFound = true

					if v.(bool) {
						protocolChoiceInt := &ves_io_schema_uztna_application_uztna_application.Protocol_UDP{}
						protocolChoiceInt.UDP = &ves_io_schema.Empty{}
						protocol.ProtocolChoice = protocolChoiceInt
					}

				}

			}
		}

	}

	//url
	if v, ok := d.GetOk("url"); ok && !isIntfNil(v) {

		createSpec.Url =
			v.(string)

	}

	//uztna_domain_ref
	if v, ok := d.GetOk("uztna_domain_ref"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		uztnaDomainRefInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.UztnaDomainRef = uztnaDomainRefInt
		for i, ps := range sl {

			udrMapToStrVal := ps.(map[string]interface{})
			uztnaDomainRefInt[i] = &ves_io_schema.ObjectRefType{}

			uztnaDomainRefInt[i].Kind = "uztna_domain"

			if v, ok := udrMapToStrVal["name"]; ok && !isIntfNil(v) {
				uztnaDomainRefInt[i].Name = v.(string)
			}

			if v, ok := udrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				uztnaDomainRefInt[i].Namespace = v.(string)
			}

			if v, ok := udrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				uztnaDomainRefInt[i].Tenant = v.(string)
			}

			if v, ok := udrMapToStrVal["uid"]; ok && !isIntfNil(v) {
				uztnaDomainRefInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra UztnaApplication object with struct: %+v", createReq)

	createUztnaApplicationResp, err := client.CreateObject(context.Background(), ves_io_schema_uztna_application_uztna_application.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating UztnaApplication: %s", err)
	}
	d.SetId(createUztnaApplicationResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraUztnaApplicationRead(d, meta)
}

func resourceVolterraUztnaApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_uztna_application_uztna_application.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaApplication %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaApplication %q: %s", d.Id(), err)
	}
	return setUztnaApplicationFields(client, d, resp)
}

func setUztnaApplicationFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraUztnaApplicationUpdate updates UztnaApplication resource
func resourceVolterraUztnaApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_uztna_application_uztna_application.ReplaceSpecType{}
	updateReq := &ves_io_schema_uztna_application_uztna_application.ReplaceRequest{
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

	if v, ok := d.GetOk("app_tags"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			if v == nil {
				return fmt.Errorf("please provide valid non-empty string value of field app_tags")
			}
			if str, ok := v.(string); ok {
				ls[i] = str
			}
		}
		updateSpec.AppTags = ls

	}

	if v, ok := d.GetOk("application_vip"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		applicationVip := &ves_io_schema.IpAddressType{}
		updateSpec.ApplicationVip = applicationVip
		for _, set := range sl {
			if set != nil {
				applicationVipMapStrToI := set.(map[string]interface{})

				verTypeFound := false

				if v, ok := applicationVipMapStrToI["ipv4"]; ok && !isIntfNil(v) && !verTypeFound {

					verTypeFound = true
					verInt := &ves_io_schema.IpAddressType_Ipv4{}
					verInt.Ipv4 = &ves_io_schema.Ipv4AddressType{}
					applicationVip.Ver = verInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv4.Addr = v.(string)

							}

						}
					}

				}

				if v, ok := applicationVipMapStrToI["ipv6"]; ok && !isIntfNil(v) && !verTypeFound {

					verTypeFound = true
					verInt := &ves_io_schema.IpAddressType_Ipv6{}
					verInt.Ipv6 = &ves_io_schema.Ipv6AddressType{}
					applicationVip.Ver = verInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["addr"]; ok && !isIntfNil(v) {

								verInt.Ipv6.Addr = v.(string)

							}

						}
					}

				}

			}
		}

	}

	if v, ok := d.GetOk("extended_app_tags"); ok && !isIntfNil(v) {

		ls := make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			if v == nil {
				return fmt.Errorf("please provide valid non-empty string value of field extended_app_tags")
			}
			if str, ok := v.(string); ok {
				ls[i] = str
			}
		}
		updateSpec.ExtendedAppTags = ls

	}

	if v, ok := d.GetOk("location"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		locationInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.Location = locationInt
		for i, ps := range sl {

			lMapToStrVal := ps.(map[string]interface{})
			locationInt[i] = &ves_io_schema.ObjectRefType{}

			locationInt[i].Kind = "site"

			if v, ok := lMapToStrVal["name"]; ok && !isIntfNil(v) {
				locationInt[i].Name = v.(string)
			}

			if v, ok := lMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				locationInt[i].Namespace = v.(string)
			}

			if v, ok := lMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				locationInt[i].Tenant = v.(string)
			}

			if v, ok := lMapToStrVal["uid"]; ok && !isIntfNil(v) {
				locationInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("origin"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		origin := &ves_io_schema_uztna_application_uztna_application.Origin{}
		updateSpec.Origin = origin
		for _, set := range sl {
			if set != nil {
				originMapStrToI := set.(map[string]interface{})

				if v, ok := originMapStrToI["origin_pool"]; ok && !isIntfNil(v) {

					sl := v.([]interface{})
					originPoolInt := make([]*ves_io_schema.ObjectRefType, len(sl))
					origin.OriginPool = originPoolInt
					for i, ps := range sl {

						opMapToStrVal := ps.(map[string]interface{})
						originPoolInt[i] = &ves_io_schema.ObjectRefType{}

						originPoolInt[i].Kind = "uztna_origin_pool"

						if v, ok := opMapToStrVal["name"]; ok && !isIntfNil(v) {
							originPoolInt[i].Name = v.(string)
						}

						if v, ok := opMapToStrVal["namespace"]; ok && !isIntfNil(v) {
							originPoolInt[i].Namespace = v.(string)
						}

						if v, ok := opMapToStrVal["tenant"]; ok && !isIntfNil(v) {
							originPoolInt[i].Tenant = v.(string)
						}

						if v, ok := opMapToStrVal["uid"]; ok && !isIntfNil(v) {
							originPoolInt[i].Uid = v.(string)
						}

					}

				}

				if v, ok := originMapStrToI["snat_pool"]; ok && !isIntfNil(v) {

					sl := v.([]interface{})
					snatPoolInt := make([]*ves_io_schema.ObjectRefType, len(sl))
					origin.SnatPool = snatPoolInt
					for i, ps := range sl {

						spMapToStrVal := ps.(map[string]interface{})
						snatPoolInt[i] = &ves_io_schema.ObjectRefType{}

						snatPoolInt[i].Kind = "uztna_snat_pool"

						if v, ok := spMapToStrVal["name"]; ok && !isIntfNil(v) {
							snatPoolInt[i].Name = v.(string)
						}

						if v, ok := spMapToStrVal["namespace"]; ok && !isIntfNil(v) {
							snatPoolInt[i].Namespace = v.(string)
						}

						if v, ok := spMapToStrVal["tenant"]; ok && !isIntfNil(v) {
							snatPoolInt[i].Tenant = v.(string)
						}

						if v, ok := spMapToStrVal["uid"]; ok && !isIntfNil(v) {
							snatPoolInt[i].Uid = v.(string)
						}

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

		sl := v.([]interface{})
		protocol := &ves_io_schema_uztna_application_uztna_application.Protocol{}
		updateSpec.Protocol = protocol
		for _, set := range sl {
			if set != nil {
				protocolMapStrToI := set.(map[string]interface{})

				protocolChoiceTypeFound := false

				if v, ok := protocolMapStrToI["http"]; ok && !isIntfNil(v) && !protocolChoiceTypeFound {

					protocolChoiceTypeFound = true

					if v.(bool) {
						protocolChoiceInt := &ves_io_schema_uztna_application_uztna_application.Protocol_HTTP{}
						protocolChoiceInt.HTTP = &ves_io_schema.Empty{}
						protocol.ProtocolChoice = protocolChoiceInt
					}

				}

				if v, ok := protocolMapStrToI["https"]; ok && !isIntfNil(v) && !protocolChoiceTypeFound {

					protocolChoiceTypeFound = true
					protocolChoiceInt := &ves_io_schema_uztna_application_uztna_application.Protocol_HTTPS{}
					protocolChoiceInt.HTTPS = &ves_io_schema_uztna_application_uztna_application.AppCertificate{}
					protocol.ProtocolChoice = protocolChoiceInt

					sl := v.([]interface{})
					for _, set := range sl {
						if set != nil {
							cs := set.(map[string]interface{})

							if v, ok := cs["certificate"]; ok && !isIntfNil(v) {

								sl := v.([]interface{})
								certificateInt := make([]*ves_io_schema.ObjectRefType, len(sl))
								protocolChoiceInt.HTTPS.Certificate = certificateInt
								for i, ps := range sl {

									cMapToStrVal := ps.(map[string]interface{})
									certificateInt[i] = &ves_io_schema.ObjectRefType{}

									certificateInt[i].Kind = "certificate"

									if v, ok := cMapToStrVal["name"]; ok && !isIntfNil(v) {
										certificateInt[i].Name = v.(string)
									}

									if v, ok := cMapToStrVal["namespace"]; ok && !isIntfNil(v) {
										certificateInt[i].Namespace = v.(string)
									}

									if v, ok := cMapToStrVal["tenant"]; ok && !isIntfNil(v) {
										certificateInt[i].Tenant = v.(string)
									}

									if v, ok := cMapToStrVal["uid"]; ok && !isIntfNil(v) {
										certificateInt[i].Uid = v.(string)
									}

								}

							}

						}
					}

				}

				if v, ok := protocolMapStrToI["tcp"]; ok && !isIntfNil(v) && !protocolChoiceTypeFound {

					protocolChoiceTypeFound = true

					if v.(bool) {
						protocolChoiceInt := &ves_io_schema_uztna_application_uztna_application.Protocol_TCP{}
						protocolChoiceInt.TCP = &ves_io_schema.Empty{}
						protocol.ProtocolChoice = protocolChoiceInt
					}

				}

				if v, ok := protocolMapStrToI["udp"]; ok && !isIntfNil(v) && !protocolChoiceTypeFound {

					protocolChoiceTypeFound = true

					if v.(bool) {
						protocolChoiceInt := &ves_io_schema_uztna_application_uztna_application.Protocol_UDP{}
						protocolChoiceInt.UDP = &ves_io_schema.Empty{}
						protocol.ProtocolChoice = protocolChoiceInt
					}

				}

			}
		}

	}

	if v, ok := d.GetOk("url"); ok && !isIntfNil(v) {

		updateSpec.Url =
			v.(string)

	}

	if v, ok := d.GetOk("uztna_domain_ref"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		uztnaDomainRefInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.UztnaDomainRef = uztnaDomainRefInt
		for i, ps := range sl {

			udrMapToStrVal := ps.(map[string]interface{})
			uztnaDomainRefInt[i] = &ves_io_schema.ObjectRefType{}

			uztnaDomainRefInt[i].Kind = "uztna_domain"

			if v, ok := udrMapToStrVal["name"]; ok && !isIntfNil(v) {
				uztnaDomainRefInt[i].Name = v.(string)
			}

			if v, ok := udrMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				uztnaDomainRefInt[i].Namespace = v.(string)
			}

			if v, ok := udrMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				uztnaDomainRefInt[i].Tenant = v.(string)
			}

			if v, ok := udrMapToStrVal["uid"]; ok && !isIntfNil(v) {
				uztnaDomainRefInt[i].Uid = v.(string)
			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra UztnaApplication obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_uztna_application_uztna_application.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating UztnaApplication: %s", err)
	}

	return resourceVolterraUztnaApplicationRead(d, meta)
}

func resourceVolterraUztnaApplicationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_uztna_application_uztna_application.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] UztnaApplication %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra UztnaApplication before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra UztnaApplication obj with name %+v in namespace %+v", name, namespace)
	opts := []vesapi.CallOpt{
		vesapi.WithFailIfReferred(),
	}
	return client.DeleteObject(context.Background(), ves_io_schema_uztna_application_uztna_application.ObjectType, namespace, name, opts...)
}
