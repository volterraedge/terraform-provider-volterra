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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"gopkg.volterra.us/stdlib/codec"

	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/azure_vnet_site"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/gcp_vpc_site"
)

const (
	setCloudInfoRPCFQN = "ves.io.schema.views.%s.CustomAPI.SetCloudSiteInfo"
	setCloudInfoURI    = "/public/namespaces/system/%s/%s/set_cloud_site_info"
)

// resourceVolterraSetCloudSiteInfo is implementation of SetCloudSiteInfo custom api
// for aws_vpc_site azure_vnet_site and gcp_vpc_site
func resourceVolterraSetCloudSiteInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraSetCloudSiteInfoCreate,
		Read:   resourceVolterraSetCloudSiteInfoRead,
		Update: resourceVolterraSetCloudSiteInfoUpdate,
		Delete: resourceVolterraSetCloudSiteInfoDelete,

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
					"azure_vnet_site",
					"gcp_vpc_site",
				}, false),
			},
			"public_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"private_ips": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"az": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "use new_attribute instead",
						},
						"outside_subnet_id": {
							Optional:   true,
							Type:       schema.TypeString,
							Deprecated: "use new_attribute instead",
						},
						"inside_subnet_id": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "use new_attribute instead",
						},
						"workload_subnet_id": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "use new_attribute instead",
						},
						"outside_subnet": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"az_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ipv4_prefix": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"inside_subnet": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"az_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ipv4_prefix": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"workload_subnet": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"az_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ipv4_prefix": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"direct_connect_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"direct_connect_gateway_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vgw_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"asn": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"spoke_vnet_prefix_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vnet_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Required: true,
						},
						"prefixes": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"express_route_info": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"route_server_ips": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"route_server_asn": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
			"vnet": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vnet_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"node_info": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_instance_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"node_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraSetCloudSiteInfoCreate applies cloud site information to the site
func resourceVolterraSetCloudSiteInfoCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	var name, siteType, yamlReq string
	var err error
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("site_type"); ok {
		siteType = v.(string)
	}
	var publicIps, privateIps []string
	if v, ok := d.GetOk("public_ips"); ok {
		publicIps = make([]string, len(v.([]interface{})))
		for j, v := range v.([]interface{}) {
			publicIps[j] = v.(string)
		}
	}
	if v, ok := d.GetOk("private_ips"); ok {
		privateIps = make([]string, len(v.([]interface{})))
		for j, v := range v.([]interface{}) {
			privateIps[j] = v.(string)
		}
	}

	switch siteType {
	case "aws_vpc_site":
		subnetIDs := getSubnetIDs(d)
		req := &aws_vpc_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AwsVpcInfo: &aws_vpc_site.AWSVPCSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
				SubnetIds:  subnetIDs,
			},
		}
		if v, ok := d.GetOk("vpc_id"); ok {
			req.AwsVpcInfo.VpcId = v.(string)
		}
		if dcxInfo := getDirectConnectInfo(d); dcxInfo != nil {
			req.DirectConnectInfo = dcxInfo
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
		}

	case "azure_vnet_site":
		req := &azure_vnet_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AzureVnetInfo: &azure_vnet_site.AzureVnetSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		if spokeVnetInfo := getSpokeVnetPrefixInfo(d); spokeVnetInfo != nil {
			req.AzureVnetInfo.SpokeVnetPrefixInfo = spokeVnetInfo
		}
		if expressRouteInfo := getExpressRouteInfo(d); expressRouteInfo != nil {
			req.AzureVnetInfo.ExpressRouteInfo = expressRouteInfo
		}
		if vnetInfo := getVnetInfo(d); vnetInfo != nil {
			req.AzureVnetInfo.Vnet = vnetInfo
		}
		if nodeInstanceNames := getNodeInstanceNames(d); nodeInstanceNames != nil {
			req.AzureVnetInfo.NodeInfo = nodeInstanceNames
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
		}
	case "gcp_vpc_site":
		req := &gcp_vpc_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AwsVpcInfo: &gcp_vpc_site.GCPVPCSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}
	}
	log.Printf("[DEBUG] Setting %s SetCloudSiteInfo struct: %+v", siteType, yamlReq)
	if _, err := client.CustomAPI(context.Background(), http.MethodPost,
		fmt.Sprintf(setCloudInfoURI, siteType, name),
		fmt.Sprintf(setCloudInfoRPCFQN, siteType), yamlReq); err != nil {
		return fmt.Errorf("Error Setting %s SetCloudSiteInfo: %s", siteType, err)
	}
	d.SetId(uuid.New().String())
	return nil
}

func resourceVolterraSetCloudSiteInfoRead(d *schema.ResourceData, meta interface{}) error {
	// Ignore dont do anything when refresh/get is trigerred
	return nil
}

// resourceVolterraSetCloudSiteInfoUpdate applies cloud site information to the site
func resourceVolterraSetCloudSiteInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceVolterraSetCloudSiteInfoCreate(d, meta)
}

func resourceVolterraSetCloudSiteInfoDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	var name, siteType, yamlReq string
	var err error
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("site_type"); ok {
		siteType = v.(string)
	}
	var publicIps, privateIps []string
	switch siteType {
	case "aws_vpc_site":
		req := &aws_vpc_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AwsVpcInfo: &aws_vpc_site.AWSVPCSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
		}

	case "azure_vnet_site":
		req := &azure_vnet_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AzureVnetInfo: &azure_vnet_site.AzureVnetSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
		}
	case "gcp_vpc_site":
		req := &gcp_vpc_site.SetCloudSiteInfoRequest{
			Name:      name,
			Namespace: systemNS,
			AwsVpcInfo: &gcp_vpc_site.GCPVPCSiteInfoType{
				PublicIps:  publicIps,
				PrivateIps: privateIps,
			},
		}
		yamlReq, err = codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("error marshalling rpc response to yaml: %s", err)
		}
	}
	log.Printf("[DEBUG] Setting %s SetCloudSiteInfo struct: %+v", siteType, yamlReq)
	if _, err := client.CustomAPI(context.Background(), http.MethodPost,
		fmt.Sprintf(setCloudInfoURI, siteType, name),
		fmt.Sprintf(setCloudInfoRPCFQN, siteType), yamlReq); err != nil {
		return fmt.Errorf("error Setting %s SetCloudSiteInfo: %s", siteType, err)
	}
	d.SetId("")
	return nil
}

func getSpokeVnetPrefixInfo(d *schema.ResourceData) []*azure_vnet_site.VnetIpPrefixesType {
	if v, ok := d.GetOk("spoke_vnet_prefix_info"); ok && !isIntfNil(v) {
		sl := v.([]interface{})
		spokeVnetList := []*azure_vnet_site.VnetIpPrefixesType{}
		for _, set := range sl {
			vnetPrefixConfig := &azure_vnet_site.VnetIpPrefixesType{
				Vnet: &views.AzureVnetType{},
			}
			vnetPrefixData := set.(map[string]interface{})
			if v, ok := vnetPrefixData["vnet_name"]; ok && !isIntfNil(v) {
				vnetPrefixConfig.Vnet.VnetName = v.(string)
			}
			if v, ok := vnetPrefixData["resource_group"]; ok && !isIntfNil(v) {
				vnetPrefixConfig.Vnet.ResourceGroup = v.(string)
			}
			if val, ok := vnetPrefixData["prefixes"]; ok && !isIntfNil(v) {
				for _, v := range val.([]interface{}) {
					vnetPrefixConfig.Prefixes = append(vnetPrefixConfig.Prefixes, v.(string))
				}
			}
			spokeVnetList = append(spokeVnetList, vnetPrefixConfig)
		}
		return spokeVnetList
	}
	return nil

}

func getExpressRouteInfo(d *schema.ResourceData) *azure_vnet_site.ExpressRouteInfo {
	if v, ok := d.GetOk("express_route_info"); ok && !isIntfNil(v) {
		sl := v.(*schema.Set).List()
		erInfo := &azure_vnet_site.ExpressRouteInfo{}
		for _, set := range sl {
			erMapStrToIntf := set.(map[string]interface{})
			if w, ok := erMapStrToIntf["route_server_ips"]; ok && !isIntfNil(w) {
				ls := make([]string, len(w.([]interface{})))
				for i, v := range w.([]interface{}) {
					ls[i] = v.(string)
				}
				erInfo.RouteServerIps = ls
			}
			if val, ok := erMapStrToIntf["route_server_asn"]; ok && !isIntfNil(val) {
				erInfo.RouteServerAsn = uint32(val.(int))
			}
		}
		return erInfo

	}
	return nil
}

func getVnetInfo(d *schema.ResourceData) *azure_vnet_site.VNETInfoType {
	if v, ok := d.GetOk("vnet"); ok && !isIntfNil(v) {
		sl := v.(*schema.Set).List()
		vnetInfo := &azure_vnet_site.VNETInfoType{}
		for _, set := range sl {
			mapStrToIntf := set.(map[string]interface{})
			if val, ok := mapStrToIntf["vnet_name"]; ok && !isIntfNil(val) {
				vnetInfo.VnetName = val.(string)
			}
			if val, ok := mapStrToIntf["resource_id"]; ok && !isIntfNil(val) {
				vnetInfo.ResourceId = val.(string)
			}
		}
		return vnetInfo

	}
	return nil
}

func getNodeInstanceNames(d *schema.ResourceData) []*azure_vnet_site.NodeInstanceNameType {
	if v, ok := d.GetOk("node_info"); ok && !isIntfNil(v) {
		sl := v.([]interface{})
		nodeInfo := []*azure_vnet_site.NodeInstanceNameType{}
		for _, set := range sl {
			instanceNameConfig := &azure_vnet_site.NodeInstanceNameType{}
			nodeInstanceData := set.(map[string]interface{})
			if v, ok := nodeInstanceData["node_instance_name"]; ok && !isIntfNil(v) {
				instanceNameConfig.NodeInstanceName = v.(string)
			}
			if v, ok := nodeInstanceData["node_id"]; ok && !isIntfNil(v) {
				instanceNameConfig.NodeId = v.(string)
			}
			nodeInfo = append(nodeInfo, instanceNameConfig)
		}
		return nodeInfo
	}
	return nil
}
