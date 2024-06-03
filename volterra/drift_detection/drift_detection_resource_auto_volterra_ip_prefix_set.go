package driftdetection

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ves_io_schema_ip_prefix_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ip_prefix_set"
	"gopkg.volterra.us/stdlib/client/vesapi"
)

func DriftDetectionSpecIpPrefixSet(d *schema.ResourceData, resp vesapi.GetObjectResponse) {
	spec := resp.GetObjSpec().(*ves_io_schema_ip_prefix_set.SpecType)
	d.Set("ipv6_prefix", spec.GcSpec.GetIpv6Prefix())
	d.Set("prefix", spec.GcSpec.GetPrefix())
}
