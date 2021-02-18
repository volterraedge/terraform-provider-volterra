package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/codec"

	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

const (
	setFastACLForInternetVIPsRPCFQN = "ves.io.schema.namespace.NamespaceCustomAPI.SetFastACLsForInternetVIPs"
	fastACLForInternetVIPsURI       = "/public/namespaces/%s/fast_acls_for_internet_vips"
	getFastACLForInternetVIPsRPCFQN = "ves.io.schema.namespace.NamespaceCustomAPI.GetFastACLsForInternetVIPs"
)

// resourceVolterraAFastACLsForInternetVIPs is implementation of activating the passed list of FastACLs for Internet VIPs
func resourceVolterraAFastACLsForInternetVIPs() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAFastACLsForInternetVIPsCreate,
		Read:   resourceVolterraAFastACLsForInternetVIPsRead,
		Update: resourceVolterraAFastACLsForInternetVIPsUpdate,
		Delete: resourceVolterraAFastACLsForInternetVIPsDelete,

		Schema: map[string]*schema.Schema{

			"fast_acls": {

				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

// resourceVolterraAFastACLsForInternetVIPsCreate activates the passed list of FastACLs for Internet VIPs
func resourceVolterraAFastACLsForInternetVIPsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req := &ves_io_schema_ns.SetFastACLsForInternetVIPsRequest{
		Namespace: systemNS,
	}

	if v, ok := d.GetOk("fast_acls"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		fastACLs := make([]*ves_io_schema_views.ObjectRefType, len(sl))
		req.FastAcls = fastACLs
		for i, ps := range sl {

			policyMapData := ps.(map[string]interface{})
			fastACLs[i] = &ves_io_schema_views.ObjectRefType{}

			if v, ok := policyMapData["name"]; ok && !isIntfNil(v) {
				fastACLs[i].Name = v.(string)
			}

			fastACLs[i].Namespace = systemNS

			if v, ok := policyMapData["tenant"]; ok && !isIntfNil(v) {
				fastACLs[i].Tenant = v.(string)
			}

		}

	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Activating the passed list of FastACLs for Internet VIPs, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(fastACLForInternetVIPsURI, req.Namespace), setFastACLForInternetVIPsRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error activating the passed list of FastACLs for Internet VIPs: %s", err)
	}

	d.SetId(uuid.New().String())

	return resourceVolterraAFastACLsForInternetVIPsRead(d, meta)
}

func resourceVolterraAFastACLsForInternetVIPsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	req := &ves_io_schema_ns.GetFastACLsForInternetVIPsRequest{
		Namespace: systemNS,
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Activating the passed list of FastACLs for Internet VIPs, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodGet, fmt.Sprintf(fastACLForInternetVIPsURI, req.Namespace), getFastACLForInternetVIPsRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Println("[INFO] FastACLs for Internet VIPs is not set")
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding FastACLs for Internet VIPs: %s", err)
	}
	return nil
}

// resourceVolterraAFastACLsForInternetVIPsUpdate updates passed list of FastACLs for Internet VIPs
func resourceVolterraAFastACLsForInternetVIPsUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceVolterraAFastACLsForInternetVIPsCreate(d, meta)
}

// resourceVolterraAFastACLsForInternetVIPsDelete remove passed list of FastACLs for Internet VIPs
func resourceVolterraAFastACLsForInternetVIPsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	req := &ves_io_schema_ns.SetFastACLsForInternetVIPsRequest{
		Namespace: systemNS,
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		req.Namespace =
			v.(string)
	}
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Activating the passed list of FastACLs for Internet VIPs, with req struct: %+v", req)

	_, err = client.CustomAPI(context.Background(), http.MethodPost, fmt.Sprintf(fastACLForInternetVIPsURI, req.Namespace), setFastACLForInternetVIPsRPCFQN, yamlReq)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Println("[INFO] FastACLs for Internet VIPs is not set")
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding FastACLs for Internet VIPs: %s", err)
	}
	return nil
}
