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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"
	"gopkg.volterra.us/stdlib/server"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_tunnel "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/tunnel"
)

// dataSourceTunnel is implementation of Volterra's Namespace resource read
func dataSourceVolterraTunnel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTunnelRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipsec_psk_b64": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// read volterra namespace
func dataSourceTunnelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	rspFmt := server.GetSpecForm
	msgFQN := strings.Replace(ves_io_schema_tunnel.ObjectType, "Object", "GetResponse", 1)
	protoMsgType := proto.MessageType(msgFQN).Elem()
	protoMsg := reflect.New(protoMsgType).Interface().(proto.Message)
	callRsp := &server.CallResponse{
		ProtoMsg: protoMsg,
	}
	opts := []vesapi.CallOpt{
		vesapi.WithResponseFormat(rspFmt),
		vesapi.WithOutCallResponse(callRsp),
	}
	resp, err := client.GetObject(context.Background(), ves_io_schema_tunnel.ObjectType, namespace, name, opts...)
	if err != nil {
		return fmt.Errorf("Error finding Volterra Tunnel %s: %s", name, err)
	}

	specResp := callRsp.ProtoMsg.(*ves_io_schema_tunnel.GetResponse)
	tunType := ves_io_schema_tunnel.TunnelType_name[int32(specResp.Spec.TunnelType)]

	var ipsecPSK string
	if specResp.Spec.Params != nil {
		tunParams, ok := specResp.Spec.Params.Type.(*ves_io_schema_tunnel.TunnelParams_Ipsec)
		if !ok {
			return fmt.Errorf("Expected Tunnel type should be Ipsec")
		}
		ipsecPSKClear, ok := tunParams.Ipsec.IpsecPsk.SecretInfoOneof.(*ves_io_schema.SecretType_ClearSecretInfo)
		if !ok {
			return fmt.Errorf("Can only read psk in clear secret format")
		}
		ipsecPSK = strings.TrimPrefix(ipsecPSKClear.ClearSecretInfo.Url, "string:///")
	}

	d.Set("type", tunType)
	d.Set("ipsec_psk_b64", ipsecPSK)
	d.SetId(resp.GetObjSystemMetadata().GetUid())
	return nil
}
