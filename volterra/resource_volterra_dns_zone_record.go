//
// Copyright (c) 2023 F5 Inc. All rights reserved.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ves_io_schema_dns_zone "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dns_zone"
	ves_io_schema_dns_zone_rrset "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dns_zone/rrset"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	"gopkg.volterra.us/stdlib/codec"
)

const (
	setRRSETRPCFQN = "ves.io.schema.dns_zone.rrset.CustomAPI"
	urirrset       = "/public/namespaces/system/dns_zones/%s/rrsets/%s"
	urirrsetall    = "/public/namespaces/system/dns_zones/%s/rrsets/%s/%s/%s"
)

// resourceVolterraSetKnownLabel is implementation of Volterra's known_label resources
func resourceVolterraSetRRSETRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSetRRSETRecordCreate,
		Read:   resourceVolterraSetRRSETRecordRead,
		Update: resourceVolterraSetRRSETRecordUpdate,
		Delete: resourceVolterraSetRRSETRecordDelete,

		Schema: map[string]*schema.Schema{
			"dns_zone_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rrset": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"ttl": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						"a_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"aaaa_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"afsdb_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": {
													Type:     schema.TypeString,
													Required: true,
												},

												"subtype": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
								},
							},
						},

						"alias_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:       schema.TypeString,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"caa_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"flags": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"tag": {
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
								},
							},
						},

						"cds_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sha1_digest": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"digest": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"sha256_digest": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"digest": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"sha384_digest": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"digest": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"ds_key_algorithm": {
													Type:     schema.TypeString,
													Required: true,
												},

												"key_tag": {
													Type:     schema.TypeInt,
													Required: true,
												},
											},
										},
									},
								},
							},
						},

						"cert_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"algorithm": {
													Type:     schema.TypeString,
													Required: true,
												},

												"cert_key_tag": {
													Type:     schema.TypeInt,
													Required: true,
												},

												"cert_type": {
													Type:     schema.TypeString,
													Required: true,
												},

												"certificate": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
								},
							},
						},

						"cname_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"dlv_record": {

							Type:       schema.TypeSet,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:       schema.TypeString,
										Optional:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
									},

									"values": {

										Type:       schema.TypeList,
										Required:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sha1_digest": {

													Type:       schema.TypeSet,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"digest": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},

												"sha256_digest": {

													Type:       schema.TypeSet,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"digest": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},

												"sha384_digest": {

													Type:       schema.TypeSet,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"digest": {
																Type:       schema.TypeString,
																Required:   true,
																Deprecated: "This field is deprecated and will be removed in future release.",
															},
														},
													},
												},

												"ds_key_algorithm": {
													Type:       schema.TypeString,
													Required:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"key_tag": {
													Type:       schema.TypeInt,
													Required:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},
											},
										},
									},
								},
							},
						},

						"ds_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"sha1_digest": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"digest": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"sha256_digest": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"digest": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"sha384_digest": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"digest": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"ds_key_algorithm": {
													Type:     schema.TypeString,
													Required: true,
												},

												"key_tag": {
													Type:     schema.TypeInt,
													Required: true,
												},
											},
										},
									},
								},
							},
						},

						"eui48_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"eui64_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"lb_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"value": {

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
								},
							},
						},

						"loc_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"altitude": {
													Type:     schema.TypeFloat,
													Required: true,
												},

												"horizontal_precision": {
													Type:     schema.TypeFloat,
													Optional: true,
												},

												"latitude_degree": {
													Type:     schema.TypeInt,
													Required: true,
												},

												"latitude_hemisphere": {
													Type:     schema.TypeString,
													Required: true,
												},

												"latitude_minute": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"latitude_second": {
													Type:     schema.TypeFloat,
													Optional: true,
												},

												"location_diameter": {
													Type:     schema.TypeFloat,
													Optional: true,
												},

												"longitude_degree": {
													Type:     schema.TypeInt,
													Required: true,
												},

												"longitude_hemisphere": {
													Type:     schema.TypeString,
													Required: true,
												},

												"longitude_minute": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"longitude_second": {
													Type:     schema.TypeFloat,
													Optional: true,
												},

												"vertical_precision": {
													Type:     schema.TypeFloat,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"mx_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"domain": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"priority": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"naptr_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"flags": {
													Type:     schema.TypeString,
													Required: true,
												},

												"order": {
													Type:     schema.TypeInt,
													Required: true,
												},

												"preference": {
													Type:     schema.TypeInt,
													Required: true,
												},

												"regexp": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"replacement": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"service": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"ns_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"ptr_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"srv_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"port": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"priority": {
													Type:     schema.TypeInt,
													Optional: true,
												},

												"target": {
													Type:     schema.TypeString,
													Optional: true,
												},

												"weight": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"sshfp_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"algorithm": {
													Type:     schema.TypeString,
													Required: true,
												},

												"fingerprint": {
													Type:       schema.TypeString,
													Optional:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},

												"sha1_fingerprint": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fingerprint": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"sha256_fingerprint": {

													Type:     schema.TypeSet,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fingerprint": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},

												"fingerprinttype": {
													Type:       schema.TypeString,
													Required:   true,
													Deprecated: "This field is deprecated and will be removed in future release.",
												},
											},
										},
									},
								},
							},
						},

						"tlsa_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"certificate_association_data": {
													Type:     schema.TypeString,
													Required: true,
												},

												"certificate_usage": {
													Type:     schema.TypeString,
													Required: true,
												},

												"matching_type": {
													Type:     schema.TypeString,
													Required: true,
												},

												"selector": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
								},
							},
						},

						"txt_record": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"values": {

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
	}
}

// resourceVolterraSetRRSETRecordCreate creates dns_zone_record resource
func resourceVolterraSetRRSETRecordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var dnsZoneName, groupName string
	var RrSetGroup *ves_io_schema_dns_zone.RRSet
	if v, ok := d.GetOk("dns_zone_name"); ok {
		dnsZoneName = v.(string)
	}
	if v, ok := d.GetOk("group_name"); ok {
		groupName = v.(string)
	}
	if v, ok := d.GetOk("rrset"); ok {
		sl := v.([]interface{})
		for _, set := range sl {
			RrSetGroup = &ves_io_schema_dns_zone.RRSet{}
			RrSetGroupMapStrToI := set.(map[string]interface{})

			if w, ok := RrSetGroupMapStrToI["description"]; ok && !isIntfNil(w) {
				RrSetGroup.Description = w.(string)
			}

			if w, ok := RrSetGroupMapStrToI["ttl"]; ok && !isIntfNil(w) {
				RrSetGroup.Ttl = uint32(w.(int))
			}

			typeRecordSetTypeFound := false

			if v, ok := RrSetGroupMapStrToI["a_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_ARecord{}
				typeRecordSetInt.ARecord = &ves_io_schema_dns_zone.DNSAResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.ARecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.ARecord.Values = ls

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["aaaa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_AaaaRecord{}
				typeRecordSetInt.AaaaRecord = &ves_io_schema_dns_zone.DNSAAAAResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.AaaaRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.AaaaRecord.Values = ls

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["afsdb_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_AfsdbRecord{}
				typeRecordSetInt.AfsdbRecord = &ves_io_schema_dns_zone.DNSAFSDBRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.AfsdbRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.AFSDBRecordValue, len(sl))
						typeRecordSetInt.AfsdbRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.AFSDBRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["hostname"]; ok && !isIntfNil(w) {
								values[i].Hostname = w.(string)
							}

							if v, ok := valuesMapStrToI["subtype"]; ok && !isIntfNil(v) {

								values[i].Subtype = ves_io_schema_dns_zone.AFSDBRecordSubtype(ves_io_schema_dns_zone.AFSDBRecordSubtype_value[v.(string)])

							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["alias_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_AliasRecord{}
				typeRecordSetInt.AliasRecord = &ves_io_schema_dns_zone.DNSAliasResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.AliasRecord.Name = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						typeRecordSetInt.AliasRecord.Value = v.(string)

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["caa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_CaaRecord{}
				typeRecordSetInt.CaaRecord = &ves_io_schema_dns_zone.DNSCAAResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CaaRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.CertificationAuthorityAuthorization, len(sl))
						typeRecordSetInt.CaaRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.CertificationAuthorityAuthorization{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["flags"]; ok && !isIntfNil(w) {
								values[i].Flags = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["tag"]; ok && !isIntfNil(w) {
								values[i].Tag = w.(string)
							}

							if w, ok := valuesMapStrToI["value"]; ok && !isIntfNil(w) {
								values[i].Value = w.(string)
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["cds_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_CdsRecord{}
				typeRecordSetInt.CdsRecord = &ves_io_schema_dns_zone.DNSCDSRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CdsRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.DSRecordValue, len(sl))
						typeRecordSetInt.CdsRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.DSRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							digestChoiceTypeFound := false

							if v, ok := valuesMapStrToI["sha1_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha1Digest{}
								digestChoiceInt.Sha1Digest = &ves_io_schema_dns_zone.SHA1Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha1Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha256_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha256Digest{}
								digestChoiceInt.Sha256Digest = &ves_io_schema_dns_zone.SHA256Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha256Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha384_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha384Digest{}
								digestChoiceInt.Sha384Digest = &ves_io_schema_dns_zone.SHA384Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha384Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["ds_key_algorithm"]; ok && !isIntfNil(v) {

								values[i].DsKeyAlgorithm = ves_io_schema_dns_zone.DSKeyAlgorithm(ves_io_schema_dns_zone.DSKeyAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["key_tag"]; ok && !isIntfNil(w) {
								values[i].KeyTag = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["cert_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_CertRecord{}
				typeRecordSetInt.CertRecord = &ves_io_schema_dns_zone.CERTResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CertRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.CERTRecordValue, len(sl))
						typeRecordSetInt.CertRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.CERTRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if v, ok := valuesMapStrToI["algorithm"]; ok && !isIntfNil(v) {

								values[i].Algorithm = ves_io_schema_dns_zone.CERTAlgorithm(ves_io_schema_dns_zone.CERTAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["cert_key_tag"]; ok && !isIntfNil(w) {
								values[i].CertKeyTag = uint32(w.(int))
							}

							if v, ok := valuesMapStrToI["cert_type"]; ok && !isIntfNil(v) {

								values[i].CertType = ves_io_schema_dns_zone.CERTType(ves_io_schema_dns_zone.CERTType_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["certificate"]; ok && !isIntfNil(w) {
								values[i].Certificate = w.(string)
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["cname_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_CnameRecord{}
				typeRecordSetInt.CnameRecord = &ves_io_schema_dns_zone.DNSCNAMEResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CnameRecord.Name = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CnameRecord.Value = v.(string)

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["dlv_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_DlvRecord{}
				typeRecordSetInt.DlvRecord = &ves_io_schema_dns_zone.DLVResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.DlvRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.DSRecordValue, len(sl))
						typeRecordSetInt.DlvRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.DSRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							digestChoiceTypeFound := false

							if v, ok := valuesMapStrToI["sha1_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha1Digest{}
								digestChoiceInt.Sha1Digest = &ves_io_schema_dns_zone.SHA1Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha1Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha256_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha256Digest{}
								digestChoiceInt.Sha256Digest = &ves_io_schema_dns_zone.SHA256Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha256Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha384_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha384Digest{}
								digestChoiceInt.Sha384Digest = &ves_io_schema_dns_zone.SHA384Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha384Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["ds_key_algorithm"]; ok && !isIntfNil(v) {

								values[i].DsKeyAlgorithm = ves_io_schema_dns_zone.DSKeyAlgorithm(ves_io_schema_dns_zone.DSKeyAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["key_tag"]; ok && !isIntfNil(w) {
								values[i].KeyTag = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["ds_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_DsRecord{}
				typeRecordSetInt.DsRecord = &ves_io_schema_dns_zone.DNSDSRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.DsRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.DSRecordValue, len(sl))
						typeRecordSetInt.DsRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.DSRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							digestChoiceTypeFound := false

							if v, ok := valuesMapStrToI["sha1_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha1Digest{}
								digestChoiceInt.Sha1Digest = &ves_io_schema_dns_zone.SHA1Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha1Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha256_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha256Digest{}
								digestChoiceInt.Sha256Digest = &ves_io_schema_dns_zone.SHA256Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha256Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha384_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha384Digest{}
								digestChoiceInt.Sha384Digest = &ves_io_schema_dns_zone.SHA384Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha384Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["ds_key_algorithm"]; ok && !isIntfNil(v) {

								values[i].DsKeyAlgorithm = ves_io_schema_dns_zone.DSKeyAlgorithm(ves_io_schema_dns_zone.DSKeyAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["key_tag"]; ok && !isIntfNil(w) {
								values[i].KeyTag = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["eui48_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_Eui48Record{}
				typeRecordSetInt.Eui48Record = &ves_io_schema_dns_zone.DNSEUI48ResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.Eui48Record.Name = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						typeRecordSetInt.Eui48Record.Value = v.(string)

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["eui64_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_Eui64Record{}
				typeRecordSetInt.Eui64Record = &ves_io_schema_dns_zone.DNSEUI64ResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.Eui64Record.Name = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						typeRecordSetInt.Eui64Record.Value = v.(string)

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["lb_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_LbRecord{}
				typeRecordSetInt.LbRecord = &ves_io_schema_dns_zone.DNSLBResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.LbRecord.Name = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						valueInt := &ves_io_schema_views.ObjectRefType{}
						typeRecordSetInt.LbRecord.Value = valueInt

						for _, set := range sl {
							vMapToStrVal := set.(map[string]interface{})
							if val, ok := vMapToStrVal["name"]; ok && !isIntfNil(v) {
								valueInt.Name = val.(string)
							}
							if val, ok := vMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								valueInt.Namespace = val.(string)
							}

							if val, ok := vMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								valueInt.Tenant = val.(string)
							}
						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["loc_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_LocRecord{}
				typeRecordSetInt.LocRecord = &ves_io_schema_dns_zone.DNSLOCResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.LocRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.LOCValue, len(sl))
						typeRecordSetInt.LocRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.LOCValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["altitude"]; ok && !isIntfNil(w) {
								values[i].Altitude = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["horizontal_precision"]; ok && !isIntfNil(w) {
								values[i].HorizontalPrecision = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["latitude_degree"]; ok && !isIntfNil(w) {
								values[i].LatitudeDegree = int32(w.(int))
							}

							if v, ok := valuesMapStrToI["latitude_hemisphere"]; ok && !isIntfNil(v) {

								values[i].LatitudeHemisphere = ves_io_schema_dns_zone.LatitudeHemisphere(ves_io_schema_dns_zone.LatitudeHemisphere_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["latitude_minute"]; ok && !isIntfNil(w) {
								values[i].LatitudeMinute = int32(w.(int))
							}

							if w, ok := valuesMapStrToI["latitude_second"]; ok && !isIntfNil(w) {
								values[i].LatitudeSecond = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["location_diameter"]; ok && !isIntfNil(w) {
								values[i].LocationDiameter = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["longitude_degree"]; ok && !isIntfNil(w) {
								values[i].LongitudeDegree = int32(w.(int))
							}

							if v, ok := valuesMapStrToI["longitude_hemisphere"]; ok && !isIntfNil(v) {

								values[i].LongitudeHemisphere = ves_io_schema_dns_zone.LongitudeHemisphere(ves_io_schema_dns_zone.LongitudeHemisphere_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["longitude_minute"]; ok && !isIntfNil(w) {
								values[i].LongitudeMinute = int32(w.(int))
							}

							if w, ok := valuesMapStrToI["longitude_second"]; ok && !isIntfNil(w) {
								values[i].LongitudeSecond = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["vertical_precision"]; ok && !isIntfNil(w) {
								values[i].VerticalPrecision = float32(w.(float64))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["mx_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_MxRecord{}
				typeRecordSetInt.MxRecord = &ves_io_schema_dns_zone.DNSMXResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.MxRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.MailExchanger, len(sl))
						typeRecordSetInt.MxRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.MailExchanger{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["domain"]; ok && !isIntfNil(w) {
								values[i].Domain = w.(string)
							}

							if w, ok := valuesMapStrToI["priority"]; ok && !isIntfNil(w) {
								values[i].Priority = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["naptr_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_NaptrRecord{}
				typeRecordSetInt.NaptrRecord = &ves_io_schema_dns_zone.DNSNAPTRResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.NaptrRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.NAPTRValue, len(sl))
						typeRecordSetInt.NaptrRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.NAPTRValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["flags"]; ok && !isIntfNil(w) {
								values[i].Flags = w.(string)
							}

							if w, ok := valuesMapStrToI["order"]; ok && !isIntfNil(w) {
								values[i].Order = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["preference"]; ok && !isIntfNil(w) {
								values[i].Preference = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["regexp"]; ok && !isIntfNil(w) {
								values[i].Regexp = w.(string)
							}

							if w, ok := valuesMapStrToI["replacement"]; ok && !isIntfNil(w) {
								values[i].Replacement = w.(string)
							}

							if w, ok := valuesMapStrToI["service"]; ok && !isIntfNil(w) {
								values[i].Service = w.(string)
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["ns_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_NsRecord{}
				typeRecordSetInt.NsRecord = &ves_io_schema_dns_zone.DNSNSResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.NsRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.NsRecord.Values = ls

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["ptr_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_PtrRecord{}
				typeRecordSetInt.PtrRecord = &ves_io_schema_dns_zone.DNSPTRResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.PtrRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.PtrRecord.Values = ls

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["srv_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_SrvRecord{}
				typeRecordSetInt.SrvRecord = &ves_io_schema_dns_zone.DNSSRVResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.SrvRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.SRVService, len(sl))
						typeRecordSetInt.SrvRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.SRVService{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["port"]; ok && !isIntfNil(w) {
								values[i].Port = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["priority"]; ok && !isIntfNil(w) {
								values[i].Priority = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["target"]; ok && !isIntfNil(w) {
								values[i].Target = w.(string)
							}

							if w, ok := valuesMapStrToI["weight"]; ok && !isIntfNil(w) {
								values[i].Weight = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["sshfp_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_SshfpRecord{}
				typeRecordSetInt.SshfpRecord = &ves_io_schema_dns_zone.SSHFPResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.SshfpRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.SSHFPRecordValue, len(sl))
						typeRecordSetInt.SshfpRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.SSHFPRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if v, ok := valuesMapStrToI["algorithm"]; ok && !isIntfNil(v) {

								values[i].Algorithm = ves_io_schema_dns_zone.SSHFPAlgorithm(ves_io_schema_dns_zone.SSHFPAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["fingerprint"]; ok && !isIntfNil(w) {
								values[i].Fingerprint = w.(string)
							}

							fingerprintTypeTypeFound := false

							if v, ok := valuesMapStrToI["sha1_fingerprint"]; ok && !isIntfNil(v) && !fingerprintTypeTypeFound {

								fingerprintTypeTypeFound = true
								fingerprintTypeInt := &ves_io_schema_dns_zone.SSHFPRecordValue_Sha1Fingerprint{}
								fingerprintTypeInt.Sha1Fingerprint = &ves_io_schema_dns_zone.SHA1Fingerprint{}
								values[i].FingerprintType = fingerprintTypeInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["fingerprint"]; ok && !isIntfNil(v) {

										fingerprintTypeInt.Sha1Fingerprint.Fingerprint = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha256_fingerprint"]; ok && !isIntfNil(v) && !fingerprintTypeTypeFound {

								fingerprintTypeTypeFound = true
								fingerprintTypeInt := &ves_io_schema_dns_zone.SSHFPRecordValue_Sha256Fingerprint{}
								fingerprintTypeInt.Sha256Fingerprint = &ves_io_schema_dns_zone.SHA256Fingerprint{}
								values[i].FingerprintType = fingerprintTypeInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["fingerprint"]; ok && !isIntfNil(v) {

										fingerprintTypeInt.Sha256Fingerprint.Fingerprint = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["fingerprinttype"]; ok && !isIntfNil(v) {

								values[i].Fingerprinttype = ves_io_schema_dns_zone.SSHFPFingerprintType(ves_io_schema_dns_zone.SSHFPFingerprintType_value[v.(string)])

							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["tlsa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_TlsaRecord{}
				typeRecordSetInt.TlsaRecord = &ves_io_schema_dns_zone.TLSAResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.TlsaRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.TLSARecordValue, len(sl))
						typeRecordSetInt.TlsaRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.TLSARecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["certificate_association_data"]; ok && !isIntfNil(w) {
								values[i].CertificateAssociationData = w.(string)
							}

							if v, ok := valuesMapStrToI["certificate_usage"]; ok && !isIntfNil(v) {

								values[i].CertificateUsage = ves_io_schema_dns_zone.TLSARecordCertificateUsage(ves_io_schema_dns_zone.TLSARecordCertificateUsage_value[v.(string)])

							}

							if v, ok := valuesMapStrToI["matching_type"]; ok && !isIntfNil(v) {

								values[i].MatchingType = ves_io_schema_dns_zone.TLSARecordMatchingType(ves_io_schema_dns_zone.TLSARecordMatchingType_value[v.(string)])

							}

							if v, ok := valuesMapStrToI["selector"]; ok && !isIntfNil(v) {

								values[i].Selector = ves_io_schema_dns_zone.TLSARecordCSelector(ves_io_schema_dns_zone.TLSARecordCSelector_value[v.(string)])

							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["txt_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_TxtRecord{}
				typeRecordSetInt.TxtRecord = &ves_io_schema_dns_zone.DNSTXTResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.TxtRecord.Name = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.TxtRecord.Values = ls

					}

				}

			}

		}

	}
	req := &ves_io_schema_dns_zone_rrset.CreateRequest{
		DnsZoneName: dnsZoneName,
		GroupName:   groupName,
		Rrset:       RrSetGroup,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] setting volterra record create request: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(urirrset, dnsZoneName, groupName), fmt.Sprintf("%s.%s", setRRSETRPCFQN, "Create"), yamlReq)
	if err != nil {
		return fmt.Errorf("Error encountered while creating a record. Error: %s", err)
	}
	d.SetId(uuid.New().String())
	return resourceVolterraSetRRSETRecordRead(d, meta)
}

// resourceVolterraSetKnownLabelRead read dns_zone_record resource
func resourceVolterraSetRRSETRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var dnsZoneName, groupName, recordName, types string
	if v, ok := d.GetOk("dns_zone_name"); ok {
		dnsZoneName = v.(string)
	}
	if v, ok := d.GetOk("group_name"); ok {
		groupName = v.(string)
	}
	if v, ok := d.GetOk("rrset"); ok {
		sl := v.([]interface{})
		for _, set := range sl {
			RrSetGroupMapStrToI := set.(map[string]interface{})
			typeRecordSetTypeFound := false
			if v, ok := RrSetGroupMapStrToI["a_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "A"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}
			if v, ok := RrSetGroupMapStrToI["aaaa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "AAAA"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["afsdb_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "AFSDB"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["alias_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "ALIAS"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["caa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "CAA"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["cds_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "CDS"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["cert_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "CERT"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["cname_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "CNAME"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["dlv_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "DLV"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["ds_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "DS"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["eui48_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "EUI48"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["eui64_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "EUI64"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["lb_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "LB"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["loc_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "LOC"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["mx_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "MX"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["naptr_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "NAPTR"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["ns_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "NS"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["ptr_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "PTR"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["srv_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "SRV"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["sshfp_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "SSHFP"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["tlsa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "TLSA"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["txt_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "TXT"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

		}

	}
	req := &ves_io_schema_dns_zone_rrset.GetRequest{
		DnsZoneName: dnsZoneName,
		GroupName:   groupName,
		RecordName:  recordName,
		Type:        types,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] setting volterra dns_zone_record resource get request value: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(urirrsetall, dnsZoneName, groupName, recordName, types), fmt.Sprintf("%s.%s", setRRSETRPCFQN, "Get"), yamlReq)
	if err != nil {
		fmt.Printf("Error encountered while fetching dns_zone_record resource info. Error: %s", err)
	}
	return nil
}

// resourceVolterraSetKnownLabelUpdate updates dns_zone_record resource
func resourceVolterraSetRRSETRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var dnsZoneName, groupName, recordName, types string
	var RrSetGroup *ves_io_schema_dns_zone.RRSet
	if v, ok := d.GetOk("dns_zone_name"); ok {
		dnsZoneName = v.(string)
	}
	if v, ok := d.GetOk("group_name"); ok {
		groupName = v.(string)
	}
	if v, ok := d.GetOk("rrset"); ok {
		sl := v.([]interface{})
		for _, set := range sl {
			RrSetGroup = &ves_io_schema_dns_zone.RRSet{}
			RrSetGroupMapStrToI := set.(map[string]interface{})

			if w, ok := RrSetGroupMapStrToI["description"]; ok && !isIntfNil(w) {
				RrSetGroup.Description = w.(string)
			}

			if w, ok := RrSetGroupMapStrToI["ttl"]; ok && !isIntfNil(w) {
				RrSetGroup.Ttl = uint32(w.(int))
			}

			typeRecordSetTypeFound := false

			if v, ok := RrSetGroupMapStrToI["a_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_ARecord{}
				typeRecordSetInt.ARecord = &ves_io_schema_dns_zone.DNSAResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "A"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.ARecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.ARecord.Values = ls

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["aaaa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_AaaaRecord{}
				typeRecordSetInt.AaaaRecord = &ves_io_schema_dns_zone.DNSAAAAResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "AAAA"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.AaaaRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.AaaaRecord.Values = ls

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["afsdb_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_AfsdbRecord{}
				typeRecordSetInt.AfsdbRecord = &ves_io_schema_dns_zone.DNSAFSDBRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "AFSDB"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.AfsdbRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.AFSDBRecordValue, len(sl))
						typeRecordSetInt.AfsdbRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.AFSDBRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["hostname"]; ok && !isIntfNil(w) {
								values[i].Hostname = w.(string)
							}

							if v, ok := valuesMapStrToI["subtype"]; ok && !isIntfNil(v) {

								values[i].Subtype = ves_io_schema_dns_zone.AFSDBRecordSubtype(ves_io_schema_dns_zone.AFSDBRecordSubtype_value[v.(string)])

							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["alias_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_AliasRecord{}
				typeRecordSetInt.AliasRecord = &ves_io_schema_dns_zone.DNSAliasResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "ALIAS"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.AliasRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						typeRecordSetInt.AliasRecord.Value = v.(string)

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["caa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_CaaRecord{}
				typeRecordSetInt.CaaRecord = &ves_io_schema_dns_zone.DNSCAAResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "CAA"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CaaRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.CertificationAuthorityAuthorization, len(sl))
						typeRecordSetInt.CaaRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.CertificationAuthorityAuthorization{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["flags"]; ok && !isIntfNil(w) {
								values[i].Flags = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["tag"]; ok && !isIntfNil(w) {
								values[i].Tag = w.(string)
							}

							if w, ok := valuesMapStrToI["value"]; ok && !isIntfNil(w) {
								values[i].Value = w.(string)
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["cds_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_CdsRecord{}
				typeRecordSetInt.CdsRecord = &ves_io_schema_dns_zone.DNSCDSRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "CDS"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CdsRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.DSRecordValue, len(sl))
						typeRecordSetInt.CdsRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.DSRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							digestChoiceTypeFound := false

							if v, ok := valuesMapStrToI["sha1_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha1Digest{}
								digestChoiceInt.Sha1Digest = &ves_io_schema_dns_zone.SHA1Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha1Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha256_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha256Digest{}
								digestChoiceInt.Sha256Digest = &ves_io_schema_dns_zone.SHA256Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha256Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha384_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha384Digest{}
								digestChoiceInt.Sha384Digest = &ves_io_schema_dns_zone.SHA384Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha384Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["ds_key_algorithm"]; ok && !isIntfNil(v) {

								values[i].DsKeyAlgorithm = ves_io_schema_dns_zone.DSKeyAlgorithm(ves_io_schema_dns_zone.DSKeyAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["key_tag"]; ok && !isIntfNil(w) {
								values[i].KeyTag = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["cert_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_CertRecord{}
				typeRecordSetInt.CertRecord = &ves_io_schema_dns_zone.CERTResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "CERT"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CertRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.CERTRecordValue, len(sl))
						typeRecordSetInt.CertRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.CERTRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if v, ok := valuesMapStrToI["algorithm"]; ok && !isIntfNil(v) {

								values[i].Algorithm = ves_io_schema_dns_zone.CERTAlgorithm(ves_io_schema_dns_zone.CERTAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["cert_key_tag"]; ok && !isIntfNil(w) {
								values[i].CertKeyTag = uint32(w.(int))
							}

							if v, ok := valuesMapStrToI["cert_type"]; ok && !isIntfNil(v) {

								values[i].CertType = ves_io_schema_dns_zone.CERTType(ves_io_schema_dns_zone.CERTType_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["certificate"]; ok && !isIntfNil(w) {
								values[i].Certificate = w.(string)
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["cname_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_CnameRecord{}
				typeRecordSetInt.CnameRecord = &ves_io_schema_dns_zone.DNSCNAMEResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "CNAME"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CnameRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						typeRecordSetInt.CnameRecord.Value = v.(string)

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["dlv_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_DlvRecord{}
				typeRecordSetInt.DlvRecord = &ves_io_schema_dns_zone.DLVResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "DLV"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.DlvRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.DSRecordValue, len(sl))
						typeRecordSetInt.DlvRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.DSRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							digestChoiceTypeFound := false

							if v, ok := valuesMapStrToI["sha1_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha1Digest{}
								digestChoiceInt.Sha1Digest = &ves_io_schema_dns_zone.SHA1Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha1Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha256_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha256Digest{}
								digestChoiceInt.Sha256Digest = &ves_io_schema_dns_zone.SHA256Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha256Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha384_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha384Digest{}
								digestChoiceInt.Sha384Digest = &ves_io_schema_dns_zone.SHA384Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha384Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["ds_key_algorithm"]; ok && !isIntfNil(v) {

								values[i].DsKeyAlgorithm = ves_io_schema_dns_zone.DSKeyAlgorithm(ves_io_schema_dns_zone.DSKeyAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["key_tag"]; ok && !isIntfNil(w) {
								values[i].KeyTag = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["ds_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_DsRecord{}
				typeRecordSetInt.DsRecord = &ves_io_schema_dns_zone.DNSDSRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "DS"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.DsRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.DSRecordValue, len(sl))
						typeRecordSetInt.DsRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.DSRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							digestChoiceTypeFound := false

							if v, ok := valuesMapStrToI["sha1_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha1Digest{}
								digestChoiceInt.Sha1Digest = &ves_io_schema_dns_zone.SHA1Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha1Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha256_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha256Digest{}
								digestChoiceInt.Sha256Digest = &ves_io_schema_dns_zone.SHA256Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha256Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha384_digest"]; ok && !isIntfNil(v) && !digestChoiceTypeFound {

								digestChoiceTypeFound = true
								digestChoiceInt := &ves_io_schema_dns_zone.DSRecordValue_Sha384Digest{}
								digestChoiceInt.Sha384Digest = &ves_io_schema_dns_zone.SHA384Digest{}
								values[i].DigestChoice = digestChoiceInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["digest"]; ok && !isIntfNil(v) {

										digestChoiceInt.Sha384Digest.Digest = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["ds_key_algorithm"]; ok && !isIntfNil(v) {

								values[i].DsKeyAlgorithm = ves_io_schema_dns_zone.DSKeyAlgorithm(ves_io_schema_dns_zone.DSKeyAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["key_tag"]; ok && !isIntfNil(w) {
								values[i].KeyTag = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["eui48_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_Eui48Record{}
				typeRecordSetInt.Eui48Record = &ves_io_schema_dns_zone.DNSEUI48ResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "EUI48"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.Eui48Record.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						typeRecordSetInt.Eui48Record.Value = v.(string)

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["eui64_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_Eui64Record{}
				typeRecordSetInt.Eui64Record = &ves_io_schema_dns_zone.DNSEUI64ResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "EUI64"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.Eui64Record.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						typeRecordSetInt.Eui64Record.Value = v.(string)

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["lb_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_LbRecord{}
				typeRecordSetInt.LbRecord = &ves_io_schema_dns_zone.DNSLBResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "LB"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.LbRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["value"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						valueInt := &ves_io_schema_views.ObjectRefType{}
						typeRecordSetInt.LbRecord.Value = valueInt

						for _, set := range sl {
							vMapToStrVal := set.(map[string]interface{})
							if val, ok := vMapToStrVal["name"]; ok && !isIntfNil(v) {
								valueInt.Name = val.(string)
							}
							if val, ok := vMapToStrVal["namespace"]; ok && !isIntfNil(v) {
								valueInt.Namespace = val.(string)
							}

							if val, ok := vMapToStrVal["tenant"]; ok && !isIntfNil(v) {
								valueInt.Tenant = val.(string)
							}
						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["loc_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_LocRecord{}
				typeRecordSetInt.LocRecord = &ves_io_schema_dns_zone.DNSLOCResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "LOC"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.LocRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.LOCValue, len(sl))
						typeRecordSetInt.LocRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.LOCValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["altitude"]; ok && !isIntfNil(w) {
								values[i].Altitude = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["horizontal_precision"]; ok && !isIntfNil(w) {
								values[i].HorizontalPrecision = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["latitude_degree"]; ok && !isIntfNil(w) {
								values[i].LatitudeDegree = int32(w.(int))
							}

							if v, ok := valuesMapStrToI["latitude_hemisphere"]; ok && !isIntfNil(v) {

								values[i].LatitudeHemisphere = ves_io_schema_dns_zone.LatitudeHemisphere(ves_io_schema_dns_zone.LatitudeHemisphere_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["latitude_minute"]; ok && !isIntfNil(w) {
								values[i].LatitudeMinute = int32(w.(int))
							}

							if w, ok := valuesMapStrToI["latitude_second"]; ok && !isIntfNil(w) {
								values[i].LatitudeSecond = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["location_diameter"]; ok && !isIntfNil(w) {
								values[i].LocationDiameter = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["longitude_degree"]; ok && !isIntfNil(w) {
								values[i].LongitudeDegree = int32(w.(int))
							}

							if v, ok := valuesMapStrToI["longitude_hemisphere"]; ok && !isIntfNil(v) {

								values[i].LongitudeHemisphere = ves_io_schema_dns_zone.LongitudeHemisphere(ves_io_schema_dns_zone.LongitudeHemisphere_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["longitude_minute"]; ok && !isIntfNil(w) {
								values[i].LongitudeMinute = int32(w.(int))
							}

							if w, ok := valuesMapStrToI["longitude_second"]; ok && !isIntfNil(w) {
								values[i].LongitudeSecond = float32(w.(float64))
							}

							if w, ok := valuesMapStrToI["vertical_precision"]; ok && !isIntfNil(w) {
								values[i].VerticalPrecision = float32(w.(float64))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["mx_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_MxRecord{}
				typeRecordSetInt.MxRecord = &ves_io_schema_dns_zone.DNSMXResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "MX"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.MxRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.MailExchanger, len(sl))
						typeRecordSetInt.MxRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.MailExchanger{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["domain"]; ok && !isIntfNil(w) {
								values[i].Domain = w.(string)
							}

							if w, ok := valuesMapStrToI["priority"]; ok && !isIntfNil(w) {
								values[i].Priority = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["naptr_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_NaptrRecord{}
				typeRecordSetInt.NaptrRecord = &ves_io_schema_dns_zone.DNSNAPTRResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "NAPTR"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.NaptrRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.NAPTRValue, len(sl))
						typeRecordSetInt.NaptrRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.NAPTRValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["flags"]; ok && !isIntfNil(w) {
								values[i].Flags = w.(string)
							}

							if w, ok := valuesMapStrToI["order"]; ok && !isIntfNil(w) {
								values[i].Order = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["preference"]; ok && !isIntfNil(w) {
								values[i].Preference = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["regexp"]; ok && !isIntfNil(w) {
								values[i].Regexp = w.(string)
							}

							if w, ok := valuesMapStrToI["replacement"]; ok && !isIntfNil(w) {
								values[i].Replacement = w.(string)
							}

							if w, ok := valuesMapStrToI["service"]; ok && !isIntfNil(w) {
								values[i].Service = w.(string)
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["ns_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_NsRecord{}
				typeRecordSetInt.NsRecord = &ves_io_schema_dns_zone.DNSNSResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "NS"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.NsRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.NsRecord.Values = ls

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["ptr_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_PtrRecord{}
				typeRecordSetInt.PtrRecord = &ves_io_schema_dns_zone.DNSPTRResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "PTR"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.PtrRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.PtrRecord.Values = ls

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["srv_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_SrvRecord{}
				typeRecordSetInt.SrvRecord = &ves_io_schema_dns_zone.DNSSRVResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "SRV"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.SrvRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.SRVService, len(sl))
						typeRecordSetInt.SrvRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.SRVService{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["port"]; ok && !isIntfNil(w) {
								values[i].Port = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["priority"]; ok && !isIntfNil(w) {
								values[i].Priority = uint32(w.(int))
							}

							if w, ok := valuesMapStrToI["target"]; ok && !isIntfNil(w) {
								values[i].Target = w.(string)
							}

							if w, ok := valuesMapStrToI["weight"]; ok && !isIntfNil(w) {
								values[i].Weight = uint32(w.(int))
							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["sshfp_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_SshfpRecord{}
				typeRecordSetInt.SshfpRecord = &ves_io_schema_dns_zone.SSHFPResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "SSHFP"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.SshfpRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.SSHFPRecordValue, len(sl))
						typeRecordSetInt.SshfpRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.SSHFPRecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if v, ok := valuesMapStrToI["algorithm"]; ok && !isIntfNil(v) {

								values[i].Algorithm = ves_io_schema_dns_zone.SSHFPAlgorithm(ves_io_schema_dns_zone.SSHFPAlgorithm_value[v.(string)])

							}

							if w, ok := valuesMapStrToI["fingerprint"]; ok && !isIntfNil(w) {
								values[i].Fingerprint = w.(string)
							}

							fingerprintTypeTypeFound := false

							if v, ok := valuesMapStrToI["sha1_fingerprint"]; ok && !isIntfNil(v) && !fingerprintTypeTypeFound {

								fingerprintTypeTypeFound = true
								fingerprintTypeInt := &ves_io_schema_dns_zone.SSHFPRecordValue_Sha1Fingerprint{}
								fingerprintTypeInt.Sha1Fingerprint = &ves_io_schema_dns_zone.SHA1Fingerprint{}
								values[i].FingerprintType = fingerprintTypeInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["fingerprint"]; ok && !isIntfNil(v) {

										fingerprintTypeInt.Sha1Fingerprint.Fingerprint = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["sha256_fingerprint"]; ok && !isIntfNil(v) && !fingerprintTypeTypeFound {

								fingerprintTypeTypeFound = true
								fingerprintTypeInt := &ves_io_schema_dns_zone.SSHFPRecordValue_Sha256Fingerprint{}
								fingerprintTypeInt.Sha256Fingerprint = &ves_io_schema_dns_zone.SHA256Fingerprint{}
								values[i].FingerprintType = fingerprintTypeInt

								sl := v.(*schema.Set).List()
								for _, set := range sl {
									cs := set.(map[string]interface{})

									if v, ok := cs["fingerprint"]; ok && !isIntfNil(v) {

										fingerprintTypeInt.Sha256Fingerprint.Fingerprint = v.(string)

									}

								}

							}

							if v, ok := valuesMapStrToI["fingerprinttype"]; ok && !isIntfNil(v) {

								values[i].Fingerprinttype = ves_io_schema_dns_zone.SSHFPFingerprintType(ves_io_schema_dns_zone.SSHFPFingerprintType_value[v.(string)])

							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["tlsa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_TlsaRecord{}
				typeRecordSetInt.TlsaRecord = &ves_io_schema_dns_zone.TLSAResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "TLSA"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.TlsaRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						sl := v.([]interface{})
						values := make([]*ves_io_schema_dns_zone.TLSARecordValue, len(sl))
						typeRecordSetInt.TlsaRecord.Values = values
						for i, set := range sl {
							values[i] = &ves_io_schema_dns_zone.TLSARecordValue{}
							valuesMapStrToI := set.(map[string]interface{})

							if w, ok := valuesMapStrToI["certificate_association_data"]; ok && !isIntfNil(w) {
								values[i].CertificateAssociationData = w.(string)
							}

							if v, ok := valuesMapStrToI["certificate_usage"]; ok && !isIntfNil(v) {

								values[i].CertificateUsage = ves_io_schema_dns_zone.TLSARecordCertificateUsage(ves_io_schema_dns_zone.TLSARecordCertificateUsage_value[v.(string)])

							}

							if v, ok := valuesMapStrToI["matching_type"]; ok && !isIntfNil(v) {

								values[i].MatchingType = ves_io_schema_dns_zone.TLSARecordMatchingType(ves_io_schema_dns_zone.TLSARecordMatchingType_value[v.(string)])

							}

							if v, ok := valuesMapStrToI["selector"]; ok && !isIntfNil(v) {

								values[i].Selector = ves_io_schema_dns_zone.TLSARecordCSelector(ves_io_schema_dns_zone.TLSARecordCSelector_value[v.(string)])

							}

						}

					}

				}

			}

			if v, ok := RrSetGroupMapStrToI["txt_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {

				typeRecordSetTypeFound = true
				typeRecordSetInt := &ves_io_schema_dns_zone.RRSet_TxtRecord{}
				typeRecordSetInt.TxtRecord = &ves_io_schema_dns_zone.DNSTXTResourceRecord{}
				RrSetGroup.TypeRecordSet = typeRecordSetInt
				types = "TXT"

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						typeRecordSetInt.TxtRecord.Name = v.(string)
						recordName = v.(string)

					}

					if v, ok := cs["values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						typeRecordSetInt.TxtRecord.Values = ls

					}

				}

			}

		}

	}
	reqUp := &ves_io_schema_dns_zone_rrset.ReplaceRequest{
		DnsZoneName: dnsZoneName,
		GroupName:   groupName,
		RecordName:  recordName,
		Type:        types,
		Rrset:       RrSetGroup,
	}
	yamlReqUp, err := codec.ToYAML(reqUp)
	if err != nil {
		log.Printf("Ignored: Error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] setting volterra dns_zone_record resource update request value: %+v", reqUp)
	_, err = client.CustomAPI(context.Background(), http.MethodPut, fmt.Sprintf(urirrsetall, dnsZoneName, groupName, recordName, types), fmt.Sprintf("%s.%s", setRRSETRPCFQN, "Replace"), yamlReqUp)
	if err != nil {
		log.Printf("Update: Ignored: Error encountered while updating dns_zone_record resource. Error: : %s", err)
	}
	return nil
}

// resourceVolterraSetRRSETRecordDelete delete dns_zone_record resource
func resourceVolterraSetRRSETRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var dnsZoneName, groupName, recordName, types string
	if v, ok := d.GetOk("dns_zone_name"); ok {
		dnsZoneName = v.(string)
	}
	if v, ok := d.GetOk("group_name"); ok {
		groupName = v.(string)
	}
	if v, ok := d.GetOk("rrset"); ok {
		sl := v.([]interface{})
		for _, set := range sl {
			RrSetGroupMapStrToI := set.(map[string]interface{})
			typeRecordSetTypeFound := false
			if v, ok := RrSetGroupMapStrToI["a_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "A"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}
			if v, ok := RrSetGroupMapStrToI["aaaa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "AAAA"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["afsdb_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "AFSDB"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["alias_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "ALIAS"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["caa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "CAA"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["cds_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "CDS"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["cert_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "CERT"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["cname_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "CNAME"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["dlv_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "DLV"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["ds_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "DS"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["eui48_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "EUI48"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["eui64_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "EUI64"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["lb_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "LB"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["loc_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "LOC"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["mx_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "MX"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["naptr_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "NAPTR"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["ns_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "NS"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["ptr_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "PTR"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["srv_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "SRV"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["sshfp_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "SSHFP"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["tlsa_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "TLSA"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

			if v, ok := RrSetGroupMapStrToI["txt_record"]; ok && !isIntfNil(v) && !typeRecordSetTypeFound {
				typeRecordSetTypeFound = true
				types = "TXT"
				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})
					if v, ok := cs["name"]; ok && !isIntfNil(v) {
						recordName = v.(string)
					}
				}
			}

		}

	}
	reqDel := &ves_io_schema_dns_zone_rrset.DeleteRequest{
		DnsZoneName: dnsZoneName,
		GroupName:   groupName,
		RecordName:  recordName,
		Type:        types,
	}
	yamlReqDel, err := codec.ToYAML(reqDel)
	if err != nil {
		log.Printf("Ignored: Error marshalling rpc response to yaml: %s", err)
	}
	log.Printf("[DEBUG] setting volterra dns_zone_record resource delete request value: %+v", reqDel)
	_, err = client.CustomAPI(context.Background(), http.MethodDelete, fmt.Sprintf(urirrsetall, dnsZoneName, groupName, recordName, types), fmt.Sprintf("%s.%s", setRRSETRPCFQN, "Delete"), yamlReqDel)
	if err != nil {
		log.Printf("Update: Ignored: Error encountered while deleting dns_zone_record resource. Error: : %s", err)
	}
	return nil
}
