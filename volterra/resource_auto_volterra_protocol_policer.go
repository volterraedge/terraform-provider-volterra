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
	ves_io_schema_protocol_policer "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/protocol_policer"
)

// resourceVolterraProtocolPolicer is implementation of Volterra's ProtocolPolicer resources
func resourceVolterraProtocolPolicer() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraProtocolPolicerCreate,
		Read:   resourceVolterraProtocolPolicerRead,
		Update: resourceVolterraProtocolPolicerUpdate,
		Delete: resourceVolterraProtocolPolicerDelete,

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

			"protocol_policer": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"policer": {

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

						"protocol": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dns": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
										},
									},

									"icmp": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"type": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"tcp": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"flags": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"udp": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{},
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

// resourceVolterraProtocolPolicerCreate creates ProtocolPolicer resource
func resourceVolterraProtocolPolicerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_protocol_policer.CreateSpecType{}
	createReq := &ves_io_schema_protocol_policer.CreateRequest{
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

	if v, ok := d.GetOk("protocol_policer"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		protocolPolicer := make([]*ves_io_schema_protocol_policer.ProtocolPolicerType, len(sl))
		createSpec.ProtocolPolicer = protocolPolicer
		for i, set := range sl {
			protocolPolicer[i] = &ves_io_schema_protocol_policer.ProtocolPolicerType{}

			protocolPolicerMapStrToI := set.(map[string]interface{})

			if v, ok := protocolPolicerMapStrToI["policer"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				policerInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				protocolPolicer[i].Policer = policerInt
				for i, ps := range sl {

					pMapToStrVal := ps.(map[string]interface{})
					policerInt[i] = &ves_io_schema.ObjectRefType{}

					policerInt[i].Kind = "policer"

					if v, ok := pMapToStrVal["name"]; ok && !isIntfNil(v) {
						policerInt[i].Name = v.(string)
					}

					if v, ok := pMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						policerInt[i].Namespace = v.(string)
					}

					if v, ok := pMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						policerInt[i].Tenant = v.(string)
					}

					if v, ok := pMapToStrVal["uid"]; ok && !isIntfNil(v) {
						policerInt[i].Uid = v.(string)
					}

				}

			}

			if v, ok := protocolPolicerMapStrToI["protocol"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				protocol := &ves_io_schema_protocol_policer.ProtocolType{}
				protocolPolicer[i].Protocol = protocol
				for _, set := range sl {

					protocolMapStrToI := set.(map[string]interface{})

					typeTypeFound := false

					if v, ok := protocolMapStrToI["dns"]; ok && !isIntfNil(v) && !typeTypeFound {

						typeTypeFound = true
						_ = v
					}

					if v, ok := protocolMapStrToI["icmp"]; ok && !isIntfNil(v) && !typeTypeFound {

						typeTypeFound = true
						typeInt := &ves_io_schema_protocol_policer.ProtocolType_Icmp{}
						typeInt.Icmp = &ves_io_schema_protocol_policer.IcmpType{}
						protocol.Type = typeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["type"]; ok && !isIntfNil(v) {

								typeList := []ves_io_schema_protocol_policer.IcmpMsgType{}
								for _, j := range v.([]interface{}) {
									typeList = append(typeList, ves_io_schema_protocol_policer.IcmpMsgType(ves_io_schema_protocol_policer.IcmpMsgType_value[j.(string)]))
								}
								typeInt.Icmp.Type = typeList

							}

						}

					}

					if v, ok := protocolMapStrToI["tcp"]; ok && !isIntfNil(v) && !typeTypeFound {

						typeTypeFound = true
						typeInt := &ves_io_schema_protocol_policer.ProtocolType_Tcp{}
						typeInt.Tcp = &ves_io_schema_protocol_policer.TcpType{}
						protocol.Type = typeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["flags"]; ok && !isIntfNil(v) {

								flagsList := []ves_io_schema_protocol_policer.TcpFlags{}
								for _, j := range v.([]interface{}) {
									flagsList = append(flagsList, ves_io_schema_protocol_policer.TcpFlags(ves_io_schema_protocol_policer.TcpFlags_value[j.(string)]))
								}
								typeInt.Tcp.Flags = flagsList

							}

						}

					}

					if v, ok := protocolMapStrToI["udp"]; ok && !isIntfNil(v) && !typeTypeFound {

						typeTypeFound = true
						_ = v
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra ProtocolPolicer object with struct: %+v", createReq)

	createProtocolPolicerResp, err := client.CreateObject(context.Background(), ves_io_schema_protocol_policer.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating ProtocolPolicer: %s", err)
	}
	d.SetId(createProtocolPolicerResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraProtocolPolicerRead(d, meta)
}

func resourceVolterraProtocolPolicerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_protocol_policer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ProtocolPolicer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ProtocolPolicer %q: %s", d.Id(), err)
	}
	return setProtocolPolicerFields(client, d, resp)
}

func setProtocolPolicerFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraProtocolPolicerUpdate updates ProtocolPolicer resource
func resourceVolterraProtocolPolicerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_protocol_policer.ReplaceSpecType{}
	updateReq := &ves_io_schema_protocol_policer.ReplaceRequest{
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

	if v, ok := d.GetOk("protocol_policer"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		protocolPolicer := make([]*ves_io_schema_protocol_policer.ProtocolPolicerType, len(sl))
		updateSpec.ProtocolPolicer = protocolPolicer
		for i, set := range sl {
			protocolPolicer[i] = &ves_io_schema_protocol_policer.ProtocolPolicerType{}

			protocolPolicerMapStrToI := set.(map[string]interface{})

			if v, ok := protocolPolicerMapStrToI["policer"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				policerInt := make([]*ves_io_schema.ObjectRefType, len(sl))
				protocolPolicer[i].Policer = policerInt
				for i, ps := range sl {

					pMapToStrVal := ps.(map[string]interface{})
					policerInt[i] = &ves_io_schema.ObjectRefType{}

					policerInt[i].Kind = "policer"

					if v, ok := pMapToStrVal["name"]; ok && !isIntfNil(v) {
						policerInt[i].Name = v.(string)
					}

					if v, ok := pMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						policerInt[i].Namespace = v.(string)
					}

					if v, ok := pMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						policerInt[i].Tenant = v.(string)
					}

					if v, ok := pMapToStrVal["uid"]; ok && !isIntfNil(v) {
						policerInt[i].Uid = v.(string)
					}

				}

			}

			if v, ok := protocolPolicerMapStrToI["protocol"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				protocol := &ves_io_schema_protocol_policer.ProtocolType{}
				protocolPolicer[i].Protocol = protocol
				for _, set := range sl {

					protocolMapStrToI := set.(map[string]interface{})

					typeTypeFound := false

					if v, ok := protocolMapStrToI["dns"]; ok && !isIntfNil(v) && !typeTypeFound {

						typeTypeFound = true
						_ = v
					}

					if v, ok := protocolMapStrToI["icmp"]; ok && !isIntfNil(v) && !typeTypeFound {

						typeTypeFound = true
						typeInt := &ves_io_schema_protocol_policer.ProtocolType_Icmp{}
						typeInt.Icmp = &ves_io_schema_protocol_policer.IcmpType{}
						protocol.Type = typeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["type"]; ok && !isIntfNil(v) {

								typeList := []ves_io_schema_protocol_policer.IcmpMsgType{}
								for _, j := range v.([]interface{}) {
									typeList = append(typeList, ves_io_schema_protocol_policer.IcmpMsgType(ves_io_schema_protocol_policer.IcmpMsgType_value[j.(string)]))
								}
								typeInt.Icmp.Type = typeList

							}

						}

					}

					if v, ok := protocolMapStrToI["tcp"]; ok && !isIntfNil(v) && !typeTypeFound {

						typeTypeFound = true
						typeInt := &ves_io_schema_protocol_policer.ProtocolType_Tcp{}
						typeInt.Tcp = &ves_io_schema_protocol_policer.TcpType{}
						protocol.Type = typeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["flags"]; ok && !isIntfNil(v) {

								flagsList := []ves_io_schema_protocol_policer.TcpFlags{}
								for _, j := range v.([]interface{}) {
									flagsList = append(flagsList, ves_io_schema_protocol_policer.TcpFlags(ves_io_schema_protocol_policer.TcpFlags_value[j.(string)]))
								}
								typeInt.Tcp.Flags = flagsList

							}

						}

					}

					if v, ok := protocolMapStrToI["udp"]; ok && !isIntfNil(v) && !typeTypeFound {

						typeTypeFound = true
						_ = v
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra ProtocolPolicer obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_protocol_policer.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating ProtocolPolicer: %s", err)
	}

	return resourceVolterraProtocolPolicerRead(d, meta)
}

func resourceVolterraProtocolPolicerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_protocol_policer.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] ProtocolPolicer %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra ProtocolPolicer before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra ProtocolPolicer obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_protocol_policer.ObjectType, namespace, name)
}
