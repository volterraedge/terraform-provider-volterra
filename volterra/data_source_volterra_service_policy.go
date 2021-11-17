//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
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

	ves_io_schema_service_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/service_policy"
)

// dataSourceVolterraServicePolicy is implementation of Volterra's Service Policy resource read
func dataSourceVolterraServicePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolterraServicePolicyRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// read volterra service_policy
func dataSourceVolterraServicePolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	rspFmt := server.GetSpecForm
	msgFQN := strings.Replace(ves_io_schema_service_policy.ObjectType, "Object", "GetResponse", 1)
	protoMsgType := proto.MessageType(msgFQN).Elem()
	protoMsg := reflect.New(protoMsgType).Interface().(proto.Message)
	callRsp := &server.CallResponse{
		ProtoMsg: protoMsg,
	}
	opts := []vesapi.CallOpt{
		vesapi.WithResponseFormat(rspFmt),
		vesapi.WithOutCallResponse(callRsp),
	}
	resp, err := client.GetObject(context.Background(), ves_io_schema_service_policy.ObjectType, namespace, name, opts...)
	if err != nil {
		return fmt.Errorf("Error finding Volterra Service Policy %s: %s", name, err)
	}

	d.SetId(resp.GetObjSystemMetadata().GetUid())
	return nil
}
