//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/server"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_namespace "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
)

const (
	deleteNSRPCFQN = "ves.io.schema.namespace.CustomAPI.CascadeDelete"
	deleteNSURI    = "/public/namespaces/%s/cascade_delete"
)

// resourceVolterraNamespace is implementation of Volterra's Namespace resources
func resourceVolterraNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraNamespaceCreate,
		Read:   resourceVolterraNamespaceRead,
		Update: resourceVolterraNamespaceUpdate,
		Delete: resourceVolterraNamespaceDelete,

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
			"uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// resourceVolterraNamespaceCreate creates Namespace resource
func resourceVolterraNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_namespace.CreateSpecType{}
	createReq := &ves_io_schema_namespace.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}
	if v, ok := d.GetOk("annotations"); ok {
		ms := map[string]string{}
		for k, v := range v.(map[string]interface{}) {
			ms[k] = v.(string)
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok {
		createMeta.Description = v.(string)
	}

	if v, ok := d.GetOk("disable"); ok {
		createMeta.Disable = v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok {
		ms := map[string]string{}
		for k, v := range v.(map[string]interface{}) {
			ms[k] = v.(string)
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok {
		createMeta.Name = v.(string)
	}

	log.Printf("[DEBUG] Creating Volterra Namespace object with struct: %+v", createReq)

	createRspType := strings.Replace(ves_io_schema_namespace.ObjectType, "Object", "GetResponse", 1)
	protoMsgType := proto.MessageType(createRspType).Elem()
	protoMsg := reflect.New(protoMsgType).Interface().(proto.Message)
	callRsp := &server.CallResponse{
		ProtoMsg: protoMsg,
	}
	opts := []vesapi.CallOpt{
		vesapi.WithOutCallResponse(callRsp),
	}

	_, err := client.CreateObject(context.Background(), ves_io_schema_namespace.ObjectType, createReq, opts...)
	if err != nil {
		return fmt.Errorf("error creating Namespace: %s", err)
	}
	createNamespaceResp := callRsp.ProtoMsg.(*ves_io_schema_namespace.CreateResponse)

	d.SetId(createNamespaceResp.SystemMetadata.Uid)
	d.Set("uid", d.Id)
	return resourceVolterraNamespaceRead(d, meta)
}

func resourceVolterraNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_namespace.ObjectType, "", name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Namespace %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Namespace %q: %s", d.Id(), err)
	}
	return setUpdatedFields(client, d, resp)
}

func setUpdatedFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()
	d.Set("name", metadata.GetName())
	d.Set("tenant_name", resp.GetObjSystemMetadata().GetTenant())
	return nil
}

// resourceVolterraNamespaceUpdate updates Namespace resource
func resourceVolterraNamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_namespace.ReplaceSpecType{}
	updateReq := &ves_io_schema_namespace.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok {
		ms := map[string]string{}
		for k, v := range v.(map[string]interface{}) {
			ms[k] = v.(string)
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok {
		updateMeta.Description = v.(string)
	}

	if v, ok := d.GetOk("disable"); ok {
		updateMeta.Disable = v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok {
		ms := map[string]string{}
		for k, v := range v.(map[string]interface{}) {
			ms[k] = v.(string)
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok {
		updateMeta.Name = v.(string)
	}

	log.Printf("[DEBUG] Updating Volterra Namespace obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_namespace.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Namespace: %s", err)
	}
	// needed for eywa to sync with akar
	time.Sleep(5 * time.Millisecond)

	return resourceVolterraNamespaceRead(d, meta)
}

func resourceVolterraNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Executing cascade_delete Volterra Namespace obj %+v ", name)
	delReq := &ves_io_schema_namespace.CascadeDeleteRequest{
		Name: name,
	}
	yamlReq, err := codec.ToYAML(delReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(deleteNSURI, name), deleteNSRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Namespace %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
	}
	// needed for eywa to sync with akar
	time.Sleep(5 * time.Millisecond)
	return err
}
