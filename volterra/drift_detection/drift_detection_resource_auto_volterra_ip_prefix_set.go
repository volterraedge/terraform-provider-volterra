package driftdetection

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ves_io_schema_ip_prefix_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ip_prefix_set"
	"gopkg.volterra.us/stdlib/client/vesapi"
)

func FlattenIpv6Prefixes(x []*ves_io_schema_ip_prefix_set.Ipv6Prefix) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		rs := map[string]interface{}{
			"description": v.GetDescription(),
			"ipv4_prefix": v.GetIpv6Prefix(),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func FlattenIpv4Prefixes(x []*ves_io_schema_ip_prefix_set.Ipv4Prefix) []interface{} {
	rslt := make([]interface{}, 0)
	for _, v := range x {
		rs := map[string]interface{}{
			"description": v.GetDescription(),
			"ipv4_prefix": v.GetIpv4Prefix(),
		}
		rslt = append(rslt, rs)
	}
	return rslt
}

func DriftDetectionSpecIpPrefixSet(d *schema.ResourceData, resp vesapi.GetObjectResponse) {
	spec := resp.GetObjSpec().(*ves_io_schema_ip_prefix_set.SpecType)
	d.Set("ipv4_prefixes", FlattenIpv4Prefixes(spec.GcSpec.GetIpv4Prefixes()))
	d.Set("ipv6_prefix", spec.GcSpec.GetIpv6Prefix())
	d.Set("ipv6_prefixes", FlattenIpv6Prefixes(spec.GcSpec.GetIpv6Prefixes()))
	d.Set("prefix", spec.GcSpec.GetPrefix())
}
