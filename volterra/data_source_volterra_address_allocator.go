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

	ves_io_schema_addr_allocator "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/address_allocator"
)

// dataSourceAddressAllocator is implementation of Volterra's Namespace resource read
func dataSourceVolterraAddressAllocator() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAddressAllocatorRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allocation_map": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

// read volterra address allocator object
func dataSourceAddressAllocatorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	rspFmt := server.GetSpecForm
	msgFQN := strings.Replace(ves_io_schema_addr_allocator.ObjectType, "Object", "GetResponse", 1)
	protoMsgType := proto.MessageType(msgFQN).Elem()
	protoMsg := reflect.New(protoMsgType).Interface().(proto.Message)
	callRsp := &server.CallResponse{
		ProtoMsg: protoMsg,
	}
	opts := []vesapi.CallOpt{
		vesapi.WithResponseFormat(rspFmt),
		vesapi.WithOutCallResponse(callRsp),
	}
	resp, err := client.GetObject(context.Background(), ves_io_schema_addr_allocator.ObjectType, namespace, name, opts...)
	if err != nil {
		return fmt.Errorf("Error finding Volterra Address Allocator %s: %s", name, err)
	}

	specResp := callRsp.ProtoMsg.(*ves_io_schema_addr_allocator.GetResponse)
	mode := ves_io_schema_addr_allocator.AllocatorMode_name[int32(specResp.Spec.Mode)]

	d.Set("mode", mode)
	d.Set("allocation_map", make(map[string]string))
	// TODO set allocation map with right values, once schema is updated
	d.SetId(resp.GetObjSystemMetadata().GetUid())
	return nil
}
