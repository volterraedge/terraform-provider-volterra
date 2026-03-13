//
// Copyright (c) 2023 F5 Inc. All rights reserved.
//

package volterra

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"
	"gopkg.volterra.us/stdlib/server"

	http_lb "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/http_loadbalancer"
	vh "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
)

// dataSourceAddressAllocator implements read of given http loadbalancer state information
func dataSourceVolterraHttpLoadbalancerState() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolterraHttpLoadbalancerStateRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_cert_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_cert_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_cert_state": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"auto_cert_expiry": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_cert_subject": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auto_cert_issuer": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_records": {
							Computed: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
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

// read volterra http lb object and fill the state information
func dataSourceVolterraHttpLoadbalancerStateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	rspFmt := server.GetSpecForm
	msgFQN := strings.Replace(http_lb.ObjectType, "Object", "GetResponse", 1)
	protoMsgType := proto.MessageType(msgFQN).Elem()
	protoMsg := reflect.New(protoMsgType).Interface().(proto.Message)
	callRsp := &server.CallResponse{
		ProtoMsg: protoMsg,
	}
	opts := []vesapi.CallOpt{
		vesapi.WithResponseFormat(rspFmt),
		vesapi.WithOutCallResponse(callRsp),
	}
	resp, err := client.GetObject(context.Background(), http_lb.ObjectType, namespace, name, opts...)
	if err != nil {
		return fmt.Errorf("Error finding Volterra HTTP Loadbalancer %s: %s", name, err)
	}

	httpLbPath := fmt.Sprintf("%s/%s", resp.GetObjNamespace(), resp.GetObjName())

	lbResp := callRsp.ProtoMsg.(*http_lb.GetResponse)
	getSpec := lbResp.GetSpec()
	// state info
	if state, ok := vh.VirtualHostState_name[int32(getSpec.GetState())]; ok {
		if err := d.Set("state", state); err != nil {
			return fmt.Errorf("setting state for http lb %s failed: %s", httpLbPath, err.Error())
		}
	}

	// cname info
	if err := d.Set("cname", getSpec.GetHostName()); err != nil {
		return fmt.Errorf("setting cname for http lb %s failed: %s", httpLbPath, err.Error())
	}

	// ipaddress info
	for _, dnsInfo := range getSpec.GetDnsInfo() {
		if err := d.Set("ip_address", dnsInfo.GetIpAddress()); err != nil {
			return fmt.Errorf("setting ip_address for http lb %s failed: %s", httpLbPath, err.Error())
		}
		break
	}

	// auto_cert_state info
	if autoCertState, ok := vh.CertificationState_name[int32(getSpec.GetAutoCertInfo().GetAutoCertState())]; ok {
		if err := d.Set("auto_cert_state", autoCertState); err != nil {
			return fmt.Errorf("setting auto_cert_state for http lb %s failed: %s", httpLbPath, err.Error())
		}
	}

	// auto_cert_info
	flattenAutoCertInfo := make([]interface{}, 1, 1)
	autoCertInfo := make(map[string]interface{})
	autoCertInfo["auto_cert_state"] = getSpec.GetAutoCertInfo().GetAutoCertState()
	autoCertInfo["auto_cert_expiry"] = getSpec.GetAutoCertInfo().GetAutoCertExpiry().GoString()
	autoCertInfo["auto_cert_subject"] = getSpec.GetAutoCertInfo().GetAutoCertSubject()
	autoCertInfo["auto_cert_issuer"] = getSpec.GetAutoCertInfo().GetAutoCertIssuer()
	flattenDnsRecords := make([]interface{}, 0)
	for _, dnsInfo := range getSpec.GetAutoCertInfo().GetDnsRecords() {
		dnsRecords := make(map[string]interface{})
		dnsRecords["name"] = dnsInfo.GetName()
		dnsRecords["type"] = dnsInfo.GetType()
		dnsRecords["value"] = dnsInfo.GetValue()
		flattenDnsRecords = append(flattenDnsRecords, dnsRecords)
	}
	autoCertInfo["dns_records"] = flattenDnsRecords
	flattenAutoCertInfo[0] = autoCertInfo
	if err := d.Set("auto_cert_info", flattenAutoCertInfo); err != nil {
		return fmt.Errorf("setting auto_cert info for http lb %s failed: %s", httpLbPath, err.Error())
	}

	d.SetId(resp.GetObjSystemMetadata().GetUid())
	return nil
}
