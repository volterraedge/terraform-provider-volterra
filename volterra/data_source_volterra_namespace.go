//
// Copyright (c) 2023 F5 Inc. All rights reserved.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ves_io_schema_namespace "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
)

// dataSourceVolterraNamespace is implementation of Volterra's Namespace resource read
func dataSourceVolterraNamespace() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolterraNamespaceRead,

		Schema: map[string]*schema.Schema{
			"name": {
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

// read volterra namespace
func dataSourceVolterraNamespaceRead(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("tenant_name", resp.GetObjSystemMetadata().GetTenant())
	d.Set("name", resp.GetObjName())
	d.SetId(resp.GetObjUid())
	return nil
}
