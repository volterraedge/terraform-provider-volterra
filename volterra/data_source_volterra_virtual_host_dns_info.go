//
// Copyright (c) 2023 F5 Inc. All rights reserved.
//

package volterra

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/codec"

	ves_io_schema_vh "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
	ves_io_schema_vh_dns_info "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host_dns_info"
)

const (
	vhDNSInfoRPCFQN = "ves.io.schema.virtual_host.CustomAPI"
)

// dataSourceVolterraVirtualHostDNSInfo is implementation of Volterra's Namespace resource read
func dataSourceVolterraVirtualHostDNSInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolterraVirtualHostGetDNSInfo,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// read volterra namespace
func dataSourceVolterraVirtualHostGetDNSInfo(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	vhGetDNSInfoReq := &ves_io_schema_vh.GetDnsInfoRequest{
		Name:      name,
		Namespace: namespace,
	}

	yamlReq, err := codec.ToYAML(vhGetDNSInfoReq)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}
	getURI := func(ns, name string) string {
		return fmt.Sprintf("/public/namespaces/%s/virtual_hosts/%s/get-dns-info", ns, name)
	}
	rspProto, err := client.CustomAPI(context.Background(), http.MethodGet,
		getURI(namespace, name), fmt.Sprintf("%s.%s", vhDNSInfoRPCFQN, "GetDnsInfo"), yamlReq)
	if err != nil {
		return fmt.Errorf("Error getting virtual host dns info: %s", err)
	}

	rspVHDnsInfo := rspProto.(*ves_io_schema_vh.GetDnsInfoResponse)
	d.SetId(name)
	d.Set("host_name", rspVHDnsInfo.GetDnsInfo().GetHostName())

	if dnsInfoList := rspVHDnsInfo.GetDnsInfo().GetDnsInfo(); dnsInfoList != nil {
		if err = d.Set("dns_info", flattenDNSInfo(dnsInfoList)); err != nil {
			return err
		}
	}
	return nil
}

func flattenDNSInfo(dnsInfoList []*ves_io_schema_vh_dns_info.DnsInfo) []interface{} {
	var out = make([]interface{}, 0, 0)
	m := make(map[string]interface{})
	for _, dnsInfo := range dnsInfoList {
		m["ip_address"] = dnsInfo.IpAddress
		out = append(out, m)
	}
	return out
}
