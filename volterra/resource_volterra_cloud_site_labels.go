//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"gopkg.volterra.us/stdlib/client/vesapi"
	"gopkg.volterra.us/stdlib/server"

	"github.com/gogo/protobuf/proto"

	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/azure_vnet_site"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/gcp_vpc_site"
)

// resourceVolterraCloudSiteLabels provides a way to update cloud site labels
func resourceVolterraCloudSiteLabels() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraCloudSiteLabelsCreate,
		Read:   resourceVolterraCloudSiteLabelsRead,
		Update: resourceVolterraCloudSiteLabelsUpdate,
		Delete: resourceVolterraCloudSiteLabelsDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"site_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"aws_vpc_site",
					"aws_tgw_site",
					"azure_vnet_site",
					"gcp_vpc_site",
				}, false),
			},
			"labels": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"ignore_on_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

// resourceVolterraCloudSiteLabelsCreate creates Site resource
func resourceVolterraCloudSiteLabelsCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceVolterraCloudSiteLabelsUpdate(d, meta)
}

func resourceVolterraCloudSiteLabelsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var name, siteType string
	name = d.Get("name").(string)
	if v, ok := d.GetOk("site_type"); ok {
		siteType = v.(string)
	}
	labels := d.Get("labels").(map[string]interface{})

	_, resp, err := getVolterraCloudSite(client, name, getCloudObjType(siteType), server.GetSpecForm)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Site %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Site %q: %s", d.Id(), err)
	}
	// if expected labels are not set in the object, then it needs to be updated
	objLabels := resp.GetObjLabels()
	for k, v := range labels {
		objVal, ok := objLabels[k]
		if !ok {
			delete(labels, k)
		} else if objVal != v {
			labels[k] = objVal

		}
	}
	d.Set("labels", labels)
	d.SetId(resp.GetObjUid())
	return nil
}

func getCloudObjType(siteType string) string {
	switch siteType {
	case "aws_vpc_site":
		return aws_vpc_site.ObjectType
	case "aws_tgw_site":
		return aws_tgw_site.ObjectType
	case "gcp_vpc_site":
		return gcp_vpc_site.ObjectType
	case "azure_vnet_site":
		return azure_vnet_site.ObjectType
	}
	return ""
}

func getVolterraCloudSite(client *APIClient, name, siteObjectType string, rspFmt server.ObjectFormat) (*server.CallResponse,
	vesapi.GetObjectResponse, error) {
	msgFQN := strings.Replace(siteObjectType, "Object", "GetResponse", 1)
	protoMsgType := proto.MessageType(msgFQN).Elem()
	protoMsg := reflect.New(protoMsgType).Interface().(proto.Message)
	callRsp := &server.CallResponse{
		ProtoMsg: protoMsg,
	}
	opts := []vesapi.CallOpt{
		vesapi.WithResponseFormat(rspFmt),
		vesapi.WithOutCallResponse(callRsp),
	}

	objResp, err := client.GetObject(context.Background(), siteObjectType, "system", name, opts...)
	return callRsp, objResp, err
}

// resourceVolterraCloudSiteLabelsUpdate updates Site resource
func resourceVolterraCloudSiteLabelsUpdate(d *schema.ResourceData, meta interface{}) error {
	//fmt.Printf("Foo update\n")
	client := meta.(*APIClient)
	var name, siteType string
	var labels map[string]interface{}
	name = d.Get("name").(string)
	if v, ok := d.GetOk("site_type"); ok {
		siteType = v.(string)
	}
	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {
		labels = v.(map[string]interface{})
	}
	resp, _, err := getVolterraCloudSite(client, name, getCloudObjType(siteType), server.ReplaceRequestForm)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Println("[INFO] Site doesn't exist, retrying")
			d.SetId("")
			return nil
		}
		return err
	}
	//fmt.Printf("Foo update getting cloud replace\n")
	updateReq := getCloudSiteReplaceForm(siteType, labels, false, resp)
	log.Printf("[DEBUG] Updating Volterra %s obj with struct: %+v", siteType, updateReq)

	err = client.ReplaceObject(context.Background(), getCloudObjType(siteType), updateReq)
	if err != nil {
		return fmt.Errorf("error updating %s: %s", siteType, err.Error())
	}
	return resourceVolterraCloudSiteLabelsRead(d, meta)
}

func getCloudSiteReplaceForm(siteType string, labels map[string]interface{}, isDel bool,
	serverResp *server.CallResponse) vesapi.ReplaceObjectRequest {

	var replaceReq vesapi.ReplaceObjectRequest
	switch siteType {
	case "aws_vpc_site":
		resp := serverResp.ProtoMsg.(*aws_vpc_site.GetResponse)
		for k, v := range labels {
			resp.ReplaceForm.Metadata.Labels[k] = v.(string)
			if isDel {
				delete(resp.ReplaceForm.Metadata.Labels, v.(string))
			}
		}
		//fmt.Printf("Foo %# v", resp)
		replaceReq = resp.ReplaceForm
	case "aws_tgw_site":
		resp := serverResp.ProtoMsg.(*aws_tgw_site.GetResponse)
		for k, v := range labels {
			resp.ReplaceForm.Metadata.Labels[k] = v.(string)
			if isDel {
				delete(resp.ReplaceForm.Metadata.Labels, v.(string))
			}
		}
		replaceReq = resp.ReplaceForm
	case "gcp_vpc_site":
		resp := serverResp.ProtoMsg.(*gcp_vpc_site.GetResponse)
		for k, v := range labels {
			resp.ReplaceForm.Metadata.Labels[k] = v.(string)
			if isDel {
				delete(resp.ReplaceForm.Metadata.Labels, v.(string))
			}
		}
		replaceReq = resp.ReplaceForm
	case "azure_vnet_site":
		resp := serverResp.ProtoMsg.(*azure_vnet_site.GetResponse)
		for k, v := range labels {
			resp.ReplaceForm.Metadata.Labels[k] = v.(string)
			if isDel {
				delete(resp.ReplaceForm.Metadata.Labels, v.(string))
			}
		}
		replaceReq = resp.ReplaceForm
	}
	//fmt.Printf("Foo replace req %# v\n", replaceReq)
	return replaceReq
}

func resourceVolterraCloudSiteLabelsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var name, siteType string
	var labels map[string]interface{}
	var ignoreDel bool
	name = d.Get("name").(string)
	if v, ok := d.GetOk("site_type"); ok {
		siteType = v.(string)
	}
	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {
		labels = v.(map[string]interface{})
	}
	if v, ok := d.GetOk("ignore_on_delete"); ok {
		ignoreDel = v.(bool)
	}
	if ignoreDel {
		d.SetId("")
		return nil
	}
	resp, _, err := getVolterraCloudSite(client, name, getCloudObjType(siteType), server.ReplaceRequestForm)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Println("[INFO] Site doesn't exist, retrying")
			d.SetId("")
			return nil
		}
		return err
	}
	updateReq := getCloudSiteReplaceForm(siteType, labels, true, resp)
	log.Printf("[DEBUG] Updating Volterra Site obj with struct: %+v", updateReq)

	err = client.ReplaceObject(context.Background(), getCloudObjType(siteType), updateReq)
	if err != nil {
		return fmt.Errorf("error updating %s: %s", siteType, err.Error())
	}
	return nil
}
