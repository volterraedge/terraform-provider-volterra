//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

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
			"tenant_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// read volterra service_policy
func dataSourceVolterraServicePolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_service_policy.ObjectType, "", namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Service Policy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Service Policy %q: %s", d.Id(), err)
	}
	d.Set("tenant_name", resp.GetObjSystemMetadata().GetTenant())
	d.Set("name", resp.GetObjName)
	d.SetId(resp.GetObjUid())
	return nil
}
